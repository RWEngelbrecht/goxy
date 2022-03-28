package internal

import (
	"net/http/httputil"
	"encoding/json"
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	t "github.com/RWEngelbrecht/jwter/types"
)

var (
	spJwt string
)

func getProxyUrl() string {
	return os.Getenv("APP_URL")
}

func getSigningCert(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func getPublicCert(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func GetCertificate(res http.ResponseWriter, req *http.Request) {

	certificate, err := getPublicCert(os.Getenv("PUBLIC_CERT"))

	if err != nil {
		log.Fatalln("createSPToken: getPublicCert: ", err)
	}

	data_to_return := t.Certs{
		Dev: string(certificate),
	}

	response := t.Payload{data_to_return}

	LogRequestPayload(req)
	json.NewEncoder(res).Encode(response)
}

func CreateSPToken() {
	// create signer for RS256
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	var signKey *rsa.PrivateKey

	token.Header["kid"] = "dev"
	// set claims
	token.Claims = &t.SPClaims{
		&jwt.StandardClaims{
			//set expire time
			ExpiresAt: time.Now().Add(time.Minute * 43200).Unix(),
		},
		os.Getenv("USER_ID"),
		"8745aff3-ffdd-4f9d-b967-e9a31b7e67ad",
		"real_user_id",
		"real_customer_id",
		[]string{"group1", "group2"},
		false, // user_impersonation
	}

	signCert, err := getSigningCert(os.Getenv("SIGNING_CERT"))
	if err != nil {
		log.Fatalln("createSPToken: getSigningCert: ", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signCert)
	if err != nil {
		log.Fatalln("createSPToken: ParseRSAPrivateKeyFromPEM: ", err)
	}

	spJwt, err = token.SignedString(signKey)
	if err != nil {
		log.Fatalln("createSPToken: SignedString: ", err)
	}
}

// serve reverse proxy for given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// update req headers
	req.Header.Set("X-Insight-Token", spJwt)
	LogRequestPayload(req)
	proxy.ServeHTTP(res, req)
}

func HandleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyUrl()

	serveReverseProxy(url, res, req)
}
