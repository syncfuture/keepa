package advertise

type RequestTokenQuery struct {
	GrantType string
	AuthCode  string
}

type RefreshTokenQuery struct {
	GrantType    string
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

	x.BaseURL = x.BaseURL + "/auth/o2/token"
	x.SetParameter("grant_type", query.GrantType)
	x.SetParameter("code", query.AuthCode)

	return x.Post()
}

func (x *AdvertiseClient) RefreshToken(query *RefreshTokenQuery) (r *AuthorizationResult, err error) {
	r = new(AuthorizationResult)

	x.BaseURL = x.BaseURL + "/auth/o2/token"
	x.SetParameter("grant_type", query.GrantType)
	x.SetParameter("refresh_token", query.RefreshToken)

	return x.Post()
}
