package rocket

import (
	"fmt"
	"net/http"
)

type ErrHTTP struct {
	Request  *http.Request
	Response *http.Response
}

func (e ErrHTTP) Error() string {
	return fmt.Sprintf("%s %s returned %s", e.Request.Method, e.Request.URL, e.Response.Status)
}
