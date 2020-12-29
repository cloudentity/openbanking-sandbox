package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"time"
)

func TestConfig(t *testing.T) {
	for k, v := range map[string]string{
		"ACPURL":         "https://localhost:8443",
		"ACPINTERNALURL": "https://acp:8443",
		"APPHOST":        "localhost",
		"UIURL":          "https://localhost:8091",
		"CERTFILE":       "cert.pem",
		"KEYFILE":        "key.pem",
	} {
		os.Setenv(k, v)
	}

	config, err := LoadConfig()
	require.NoError(t, err)

	require.Equal(t, 8091, config.Port)
	require.Equal(t, "./my.db", config.DBFile)
	require.Equal(t, "https://localhost:8443", config.ACPURL)
	require.Equal(t, "https://acp:8443", config.ACPInternalURL)
	require.Equal(t, "localhost", config.AppHost)
	require.Equal(t, "https://localhost:8091", config.UIURL)
	require.Equal(t, "cert.pem", config.CertFile)
	require.Equal(t, "key.pem", config.KeyFile)

	require.NotEmpty(t, config.Login.ClientID)
	require.NotEmpty(t, config.Login.ServerID)
	require.NotEmpty(t, config.Login.TenantID)

	require.NotEmpty(t, config.Banks[0].ID)
	require.NotEmpty(t, config.Banks[0].URL)
	require.NotEmpty(t, config.Banks[0].AcpClient.TenantID)
	require.NotEmpty(t, config.Banks[0].AcpClient.ServerID)
	require.NotEmpty(t, config.Banks[0].AcpClient.ClientID)
	require.NotEmpty(t, config.Banks[0].AcpClient.CertFile)
	require.NotEmpty(t, config.Banks[0].AcpClient.KeyFile)
	require.NotEmpty(t, config.Banks[0].AcpClient.RootCA)
	require.Equal(t, 5*time.Second, config.Banks[0].AcpClient.Timeout)
}
