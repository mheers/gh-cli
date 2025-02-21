package authheader

import (
	"fmt"
	"net/http"
	"os"
)

// CustomTransport wraps an existing http.RoundTripper and adds a header
type CustomTransport struct {
	Transport http.RoundTripper
	HeaderKey string
	HeaderVal string
}

// RoundTrip executes the request with the additional header
func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original
	reqClone := req.Clone(req.Context())

	fmt.Println("Adding custom header to request")

	// Add the custom header
	reqClone.Header.Set(c.HeaderKey, c.HeaderVal)

	// Use the wrapped transport
	return c.Transport.RoundTrip(reqClone)
}

func Transport(t http.RoundTripper) *CustomTransport {
	tr := t
	if tr == nil {
		tr = http.DefaultTransport
	}

	headers := Headers()
	value := headers["MFA"]

	return &CustomTransport{
		Transport: tr,
		HeaderKey: "MFA",
		HeaderVal: value,
	}
}

func Headers() map[string]string {
	return map[string]string{
		"MFA": fmt.Sprintf("bearer %s", os.Getenv("MFA_TOKEN")),
	}
}

func ExtendHeaders(headers map[string]string) map[string]string {
	for k, v := range Headers() {
		headers[k] = v
	}
	return headers
}
