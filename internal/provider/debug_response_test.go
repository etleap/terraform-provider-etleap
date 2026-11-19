package provider

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDebugResponseLeavesOutCredentials(t *testing.T) {
	request, _ := http.NewRequest("POST", "https://api.etleap.com/api/v2/connections", strings.NewReader(`{"password":"database-password"}`))
	request.SetBasicAuth("access-key", "secret-key")
	request.Header.Set("Cookie", "PLAY_SESSION=request-session")
	response := &http.Response{
		Status:     "400 Bad Request",
		StatusCode: 400,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{"Set-Cookie": {"PLAY_SESSION=response-session"}},
		Body:       io.NopCloser(strings.NewReader(`{"errors":["That table name is in use"]}`)),
		Request:    request,
	}

	debug := debugResponse(response)

	for _, secret := range []string{"database-password", request.Header.Get("Authorization"), "request-session", "response-session"} {
		if strings.Contains(debug, secret) {
			t.Errorf("debug output contains %q:\n%s", secret, debug)
		}
	}
	for _, expected := range []string{"POST /api/v2/connections", "Authorization: [REDACTED]", "Set-Cookie: [REDACTED]", "That table name is in use"} {
		if !strings.Contains(debug, expected) {
			t.Errorf("debug output does not contain %q:\n%s", expected, debug)
		}
	}
	body, _ := io.ReadAll(response.Body)
	if string(body) != `{"errors":["That table name is in use"]}` {
		t.Errorf("response body changed to %q", body)
	}
}
