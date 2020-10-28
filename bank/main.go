package main

import (
	"github.com/cloudentity/go-tools/checkx"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	server, err := NewServer()
	checkx.NoError(err, "failed to init server")

	err = server.Start()
	checkx.NoError(err, "failed to start server")

	/*

		r.GET("/", func(c *gin.Context) {
			var (
				loginRequest  = NewLoginRequest(c)
				grantResponse *models.ScopeGrantSessionResponse
				err           error
			)

			if err = loginRequest.Validate(); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			if grantResponse, err = server.Client.GetScopeGrant(loginRequest); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to get scope grant: %+v", err))
				return
			}

			logrus.Infof("get response payload: %+v", grantResponse)

			c.HTML(http.StatusOK, "consent.tmpl", gin.H{
				"login_request":    loginRequest,
				"client_id":        grantResponse.ClientID,
				"requested_scopes": grantResponse.RequestedScopes,
			})
		})

		r.POST("/", func(c *gin.Context) {
			var (
				loginRequest = NewLoginRequest(c)
				redirect     string
				err          error
			)

			if err = loginRequest.Validate(); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			switch c.PostForm("action") {
			case "Approve":
				if redirect, err = server.Client.ApproveScopeGrant(loginRequest, c.PostFormArray("scopes")); err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("failed to accept scope grant: %+v", err))
					return
				}

				logrus.Infof("scope grant accepted, redirect to: %s", redirect)
			case "Reject":
				if redirect, err = server.Client.RejectScopeGrant(loginRequest, "rejected", 403); err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("failed to reject scope grant: %+v", err))
					return
				}

				logrus.Infof("scope grant denied, redirect to: %s", redirect)
			default:
				c.String(http.StatusBadRequest, "invalid action")
				return
			}

			c.Redirect(http.StatusFound, redirect)
		})

		r.Run() // nolint
	*/
}
