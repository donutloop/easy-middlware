package easy_middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		panic("test")
		w.WriteHeader(http.StatusOK)
	}
	testHandler := http.HandlerFunc(handler)

	b := new(bytes.Buffer)
	logger := log.New(b, "", 0)

	dumper := func(requestDump []byte, stackDump []byte) {
		logger.Println(string(requestDump))
	}

	server := httptest.NewServer(Recovery(dumper)(testHandler))
	defer server.Close()

	response, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Recovery middleware request: %v", err)
		return
	}
	defer response.Body.Close()

	if !strings.Contains(b.String(), "User-Agent") {
		t.Errorf("Format of request is diffrent (%s)", b.String())
	}
}
