package internal

import (
	"net/http"
	"log"
	"os"
)

func GetListenAddress() string {
	return ":" + os.Getenv("PROXY_PORT")
}

func LogSetup() {
	proxy_to := os.Getenv("APP_URL")

	log.Printf("Proxy server will run on :%s\n", GetListenAddress())
	log.Printf("Server will proxy to %s\n", proxy_to)
}

func LogRequestPayload(req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.String())
	log.Printf("[%s] %s", req.Method, req.Header.Get("X-Insight-Token"))
	// for head, val := range req.Header {
	// 	log.Printf("%s: %s", head, val)
	// }
}
