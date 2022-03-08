package main

import (
	"net/http/httputil"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  //"bytes"
  "log"
  "os"

  "github.com/joho/godotenv"
)

type SPCert struct {
	Cert string `json:"certs"`
}

func getListenAddress() string {
	return ":" + os.Getenv("PORT")
}

func getProxyUrl() string {
	return os.Getenv("URL")
}


func getSPCert() {
	//TODO: make work w HTTPS
	var spCert SPCert

	endpoint := os.Getenv("CERT_URL")
	req, _ := http.NewRequest("GET", endpoint, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(err1)
	}

	err2 := json.Unmarshal(body, &spCert)
	if err2 != nil {
		panic(err2)
	}
	log.Printf(spCert.Cert)
}

func logSetup() {
	proxy_to := os.Getenv("URL")

	log.Printf("Proxy server will run on :%s\n", getListenAddress())
	log.Printf("Server will proxy to %s\n", proxy_to)
}

func logRequestPayload(req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.String())
//	for head, val := range req.Header {
//		log.Printf("%s: %s", head, val)
//	}
}

// serve reverse proxy for given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// update req headers
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Insight-Token", "TODO:GET INSIGHT TOKEN")
	req.Host = url.Host
	logRequestPayload(req)
//	getSPCert()
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyUrl()
	serveReverseProxy(url, res, req)
}

func main() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Fatal error (probably your fault): %s", err)
  }

	//log .env values
	logSetup()

	//start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}

