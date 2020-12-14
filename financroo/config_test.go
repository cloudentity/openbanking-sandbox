package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	for k, v := range map[string]string{
		"ACP_HOST":          "localhost",
		"ACP_INTERNAL_HOST": "acp",
		"APP_HOST":          "localhost",
		"CERT_FILE":         "cert.pem",
		"KEY_FILE":          "key.pem",
	} {
		os.Setenv(k, v)
	}

	config, err := LoadConfig()
	require.NoError(t, err)

	require.Equal(t, 8091, config.Port)
	require.Equal(t, "https://localhost:8443", config.AcpURL())

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
