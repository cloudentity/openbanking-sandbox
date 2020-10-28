package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	os.Setenv("CLIENT_ID", "123")
	os.Setenv("CLIENT_SECRET", "secret")
	os.Setenv("TOKEN_URL", "http://example.com")

	cfg, err := LoadConfig()
	require.NoError(t, err)
	require.Equal(t, "123", cfg.ClientID)
	require.Equal(t, "secret", cfg.ClientSecret)
	require.Equal(t, "http://example.com", cfg.TokenURL.String())
	require.Equal(t, 5*time.Second, cfg.Timeout)
}
