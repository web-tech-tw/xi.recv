// xi.recv: The opensource direct message service for LINE OpenChat (LINE Square).
// License: Apache License 2.0
// (c) 2021 Star Inc. and its contributors.
package Web

import (
	"encoding/json"
	"github.com/star-inc/xi.recv/internal/Kernel/Config"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func checkFromLINE(idToken string) (int, interface{}) {
	var err error
	var result interface{}
	client := http.Client{}
	formValues := url.Values{
		"client_id": {os.Getenv(Config.LineChannelID)},
		"id_token":  {idToken},
	}
	response, err := client.PostForm(
		"https://api.line.me/oauth2/v2.1/verify",
		formValues,
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if response.StatusCode == 200 {
		result = new(successResponse)
	} else {
		result = new(errorResponse)
	}
	return response.StatusCode, json.Unmarshal(body, result)
}

type successResponse struct {
	Iss     string `json:"iss"`
	Sub     string `json:"sub"`
	Aud     string `json:"aud"`
	Exp     string `json:"exp"`
	Iat     string `json:"iat"`
	Nonce   string `json:"nonce"`
	Amr     string `json:"amr"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}

type errorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
