package main

import (

	"net/http"
	"log"

	"github.com/joho/godotenv"

	in "github.com/RWEngelbrecht/jwter/internal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Fatal error (probably your fault): %s", err)
	}

	//log .env values
	in.LogSetup()

	in.CreateSPToken()
	if err != nil {
		log.Fatalln("main: createSPToken: ", err)
	}

	//start servers
	// public cert download endpoint
	http.HandleFunc("/l/api/jwt/certs", in.GetCertificate)
	// redirect anything else from root
	http.HandleFunc("/", in.HandleRequestAndRedirect)

	if err := http.ListenAndServe(in.GetListenAddress(), nil); err != nil {
		log.Fatalln("main: ", err)
	}
}
