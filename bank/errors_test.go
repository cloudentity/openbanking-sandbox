package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	e := ErrNotFound{resourceName: "test"}
	require.Equal(t, "unable to find resource test", e.Error())
}
