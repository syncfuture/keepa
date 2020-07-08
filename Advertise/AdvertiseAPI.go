package advertise

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/syncfuture/go/u"
)

type SponsoredProductsQuery struct {
	ProfileID string
	Code      string
}

type RequestTokenQuery struct {
	GrantType    string
	AuthCode     string
	RedirectURI  string
	AccessToken  string
	RefreshToken string
}

type AuthorizationResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func (x *AdvertiseClient) RequestToken(query *RequestTokenQuery) (r *AuthorizationResult, err error) {
	r = new(AuthorizationResult)

	contentType := "application/x-www-form-urlencoded;charset=UTF-8"

	var url string
	if query.AuthCode != "" {
		url = fmt.Sprintf("%s/auth/o2/token?grant_type=%s&code=%s&redirect_uri=%s&client_id=%s&client_secret=%s",
			x.BaseURL, query.GrantType, query.AuthCode, query.RedirectURI, x.ClientID, x.ClientSecret)
	} else {
		panic("Code cannot both be empty")
	}

	// if query.Update != "" {
	// 	url += "&update=" + query.Update
	// }

	return client.Post()
}
