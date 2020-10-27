package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudentity/acp/pkg/swagger/client"
	"github.com/cloudentity/acp/pkg/swagger/client/logins"
	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/cloudentity/go-tools/checkx"
	"github.com/cloudentity/go-tools/httpx"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	tenant := "default"

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	hc, err := httpx.NewClient(httpx.ClientConfig{
		Timeout:            time.Second * 10,
		InsecureSkipVerify: viper.GetBool("client.insecureSkipVerify"),
		RootCA:             viper.GetString("client.rootCA"),
	})
	checkx.NoError(err, "failed to create http client")

	cc := clientcredentials.Config{
		ClientID:     "buc3b1hhuc714r78env0",
		ClientSecret: "PBV7q0akoP603rZbU0EFdxbhZ-djxF7FIVwyKaLnBYU",
		Scopes:       []string{"manage_scope_grants"},
		TokenURL:     "https://authorization.cloudentity.com:8443/default/system/oauth2/token",
	}

	acpClient := client.New(httptransport.NewWithClient(
		"authorization.cloudentity.com:8443",
		"/",
		[]string{"https"},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	r.GET("/", func(c *gin.Context) {

		var (
			loginID       = c.Query("login_id")
			loginState    = c.Query("login_state")
			grantResponse *logins.GetScopeGrantRequestOK
			err           error
		)

		if loginID == "" {
			c.String(http.StatusBadRequest, "missing login_id param")
			return
		}

		if grantResponse, err = acpClient.Logins.GetScopeGrantRequest(logins.NewGetScopeGrantRequestParams().
			WithTid(tenant).
			WithLoginID(loginID), nil); err != nil {

			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get scope grant: %+v", err))
			return
		}

		logrus.Infof("get response payload: %+v", grantResponse.Payload)

		c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"login_id":         loginID,
			"login_state":      loginState,
			"client_id":        grantResponse.Payload.ClientID,
			"requested_scopes": grantResponse.Payload.RequestedScopes,
		})

	})

	r.POST("/", func(c *gin.Context) {
		var (
			loginID    = c.Query("login_id")
			loginState = c.Query("login_state")
			redirect   string
			err        error
		)

		if loginID == "" || loginState == "" {
			c.String(http.StatusBadRequest, "missing login_id or login_state param missing")
			return
		}

		switch c.PostForm("action") {
		case "Approve":
			var resp *logins.AcceptScopeGrantRequestOK

			acceptScopeGrant := &models.AcceptScopeGrant{
				GrantedScopes: c.PostFormArray("scopes"),
				ID:            loginID,
				LoginState:    loginState,
			}

			if resp, err = acpClient.Logins.AcceptScopeGrantRequest(logins.NewAcceptScopeGrantRequestParams().
				WithTid(tenant).WithLoginID(loginID).WithAcceptScopeGrant(acceptScopeGrant), nil); err != nil {

				c.String(http.StatusBadRequest, fmt.Sprintf("failed to accept scope grant: %+v", err))
				return
			}

			redirect = resp.Payload.RedirectTo
			logrus.Infof("scope grant accepted, redirect to: %s", redirect)
		case "Reject":
			var resp *logins.RejectScopeGrantRequestOK

			rejectScopeGrant := &models.RejectScopeGrant{
				ID:         loginID,
				LoginState: loginState,
				Error:      "rejected",
				StatusCode: 403,
			}

			if resp, err = acpClient.Logins.RejectScopeGrantRequest(logins.NewRejectScopeGrantRequestParams().
				WithTid(tenant).WithLoginID(loginID).WithRejectScopeGrant(rejectScopeGrant), nil); err != nil {

				c.String(http.StatusBadRequest, fmt.Sprintf("failed to reject scope grant: %+v", err))
				return
			}
			redirect = resp.Payload.RedirectTo
			logrus.Infof("scope grant denied, redirect to: %s", redirect)

		default:
			c.String(http.StatusBadRequest, "invalid action")
			return
		}

		c.Redirect(http.StatusFound, redirect)
	})

	r.Run() // nolint
}
