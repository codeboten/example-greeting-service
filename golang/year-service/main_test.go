package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/year", nil)
	w := httptest.NewRecorder()
	getHandler().ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	// defer otelconf.Shutdown(t.Context())
}
