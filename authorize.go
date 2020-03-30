package lixinsso

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// Restore the authorize type
const (
	Token = 1
	Code  = 2
)

func init() {
	// The admin config the server incorrectly. So we need to skip the ssl check.
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: cfg,
	}

	// Set cookiejar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Panic(err)
	}
	http.DefaultClient.Jar = jar

	// Capture 302
	//http.DefaultClient.CheckRedirect = checkRedirect
}

/*
func checkRedirect(req *http.Request, via *http.Request) error {

}
*/

// Authorize struct defines the basic parameter which needed in authorize url.
type Authorize struct {
	clientId     string
	responseType int
	redirectURI  string
}

// NewAuthorize function builds a new authorize object.
func NewAuthorize(clientId string, responseType int, redirectURI string) *Authorize {
	auth := new(Authorize)
	auth.clientId = clientId
	auth.responseType = responseType
	auth.redirectURI = redirectURI

	return auth
}

func ping() bool {
	resp, err := http.Get(pingURL)
	if err != nil {
		log.Panic(err)
	}

	var b []byte = make([]byte, 64)
	_, err = resp.Body.Read(b)
	if err.Error() != "EOF" {
		log.Panic(err)
	}

	if strings.Index(string(b), "not_logged_in") >= 0 {
		return false
	}

	return true
}

// Authorize function make a request to the authorize url, according to the parameters in the Authorize struct.
// The function not only can authorize token type request, but also code one.
// If the responseType is "code", the parameter user can be nil. The code authorize use session to authorize.
func (auth *Authorize) Authorize(u *User) bool {
	var responseType string
	switch auth.responseType {
	case Token:
		responseType = "token"
	case Code:
		responseType = "code"
	default:
		responseType = "unknown"
	}

	// If the responseType is "token", username and password is needed.
	var postParam map[string][]string
	if auth.responseType == Token {
		postParam = url.Values{"username": {u.username}, "password": {u.password}}
	}

	requestURL, err := url.Parse(authorizeURL)
	if err != nil {
		log.Panic(err)
	}
	getParam := url.Values{"client_id": {auth.clientId}, "response_type": {responseType}, "redirect_uri": {auth.redirectURI}}
	requestURL.RawQuery = getParam.Encode()

	_, err = http.PostForm(requestURL.String(), postParam)
	if err != nil {
		log.Panic(err)
	}

	return ping()
}

// AuthorizeApp function authorize a specific APP. The usage can infer the example.
// checkFunc parameter is used to whether the specific app is authorized. The parameter can be nil.
func (auth *Authorize) AuthorizeApp(checkFunc func() bool) bool {
	// If ping function shows that user not login, authorize failed.
	if !auth.Authorize(nil) {
		return false
	}

	if checkFunc == nil {
		return true
	}

	return checkFunc()
}
