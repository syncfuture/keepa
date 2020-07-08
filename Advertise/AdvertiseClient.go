package advertise

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/syncfuture/advertise/core"
	"github.com/syncfuture/go/u"
)

type AdvertiseClient struct {
	// Config *adconfig.ADConfig
	BaseURL      string
	RedirectURI  string
	ClientID     string
	ClientSecret string
	Parameters   map[string]string
}

func NewClient() (r *AdvertiseClient) {
	r = new(AdvertiseClient)
	// r.Config = config

	r.BaseURL = core.ConfigProvider.GetString("AD.BaseURL")
	r.RedirectURI = core.ConfigProvider.GetString("AD.RedirectURI")
	r.ClientID = core.ConfigProvider.GetString("AD.ClientID")
	r.ClientSecret = core.ConfigProvider.GetString("AD.ClientSecret")
	r.Parameters = make(map[string]string)

	return r
}

func (x *AdvertiseClient) Post() (r *AuthorizationResult, err error) {
	r = new(AuthorizationResult)

	x.SetParameter("redirect_uri", x.RedirectURI)
	x.SetParameter("client_id", x.ClientID)
	x.SetParameter("client_secret", x.ClientSecret)

	// contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	// url := x.generateURL()
	values := url.Values{}
	for k, v := range x.Parameters {
		values.Set(k, v)
	}

	resp, err := http.PostForm(x.BaseURL, values)
	if u.LogError(err) {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if u.LogError(err) {
		return
	}

	err = json.Unmarshal(bytes, r)
	if u.LogError(err) {
		return
	}

	return
}

func (x *AdvertiseClient) generateURL() string {
	uuu, _ := url.Parse(x.BaseURL)
	values := url.Values{}

	for k, v := range x.Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	uuu.RawQuery = params

	return uuu.String()
}

func (x *AdvertiseClient) SetParameter(key, value string) {
	if key != "" && value != "" {
		x.Parameters[key] = value
	} else if value == "" {
		delete(x.Parameters, key)
	}
}
