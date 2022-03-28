package types

import "github.com/dgrijalva/jwt-go"

type Payload struct {
	Certs Certs `json:"certs"`
}

type Certs struct {
	Dev string `json:"dev"`
}

type SPClaims struct {
	*jwt.StandardClaims
	User_id            string   `json:"user_id"`
	Customer_id        string   `json:"customer_id"`        // TODO: add to .env
	Real_user_id       string   `json:"real_user_id"`       // TODO: add to .env
	Real_customer_id   string   `json:"real_customer_id"`   // TODO: add to .env
	Groups             []string `json:"groups"`             // TODO: add to .env
	User_impersonation bool     `json:"user_impersonation"` // TODO: add to .env
}
