package shadow_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jianjustin/shadow.web.framework"
)

func TestForGetHandler(t *testing.T) {
	r := shadow.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.Run(":9999")
}

func TestForPostHandler(t *testing.T) {
	r := shadow.New()
	r.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
