package main

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	//"bytes"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type SPJwtInfo struct {
	UserId string
}

type SPClaims struct {
	*jwt.StandardClaims
	SPJwtInfo SPJwtInfo
}

func logSetup() {
	proxy_to := os.Getenv("URL")

	log.Printf("Proxy server will run on :%s\n", getListenAddress())
	log.Printf("Server will proxy to %s\n", proxy_to)
}

func logRequestPayload(req *http.Request) {
	log.Printf("[%s] %s", req.Method, req.URL.String())
	for head, val := range req.Header {
		log.Printf("%s: %s", head, val)
	}
}

func getListenAddress() string {
	return ":" + os.Getenv("PORT")
}

func getProxyUrl() string {
	return os.Getenv("URL")
}

func getSigningCert(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func createSPToken() (string, error) {
	log.Println("Creating SP JWT")
	// create signer for RS256
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	var signKey *rsa.PrivateKey
	// set claims
	token.Claims = &SPClaims{
		&jwt.StandardClaims{
			//set expire time
			ExpiresAt: time.Now().Add(time.Minute * 43200).Unix(),
		},
		SPJwtInfo{
			UserId: os.Getenv("USER_ID"),
		},
	}

	signCert, err := getSigningCert(os.Getenv("SIGNING_CERT"))
	if err != nil {
		log.Fatalln("createSPToken: getSigningCert: ", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signCert)
	if err != nil {
		log.Fatalln("createSPToken: ParseRSAPrivateKeyFromPEM: ", err)
	}

	return token.SignedString(signKey)
}

// serve reverse proxy for given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// update req headers
	req.Header.Set("X-Insight-Token", spJwt)
	logRequestPayload(req)

	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyUrl()

	serveReverseProxy(url, res, req)
}

var (
	spJwt string
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Fatal error (probably your fault): %s", err)
	}

	//log .env values
	logSetup()

	spJwt, err = createSPToken()
	if err != nil {
		log.Fatalln("main: createSPToken: ", err)
	}

	//start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		log.Fatalln("main: ", err)
	}
}
