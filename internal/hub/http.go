package hub

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
)

// client which uses the Docker Desktop HTTP proxy if it is available, or falls back to the default.
func client() *http.Client {
	c, err := dialDesktopHTTPProxy()
	if err != nil {
		log.Println("using the default HTTP client")
		return http.DefaultClient
	}
	_ = c.Close()
	log.Println("using the Docker Desktop HTTP proxy")
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
			}),
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return dialDesktopHTTPProxy()
			},
		},
	}
}
