package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Version       = "1.0.0"
	UserAgent     = "WooCommerce API Client/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

// Http client
type Client struct {
	storeURL  *url.URL
	apiConfig *ApiConfig
	rawClient *http.Client
}

type Interface interface {
	//Post() *Client
	//Put() *Client
	//	Patch() *Client
	Get(ctx context.Context, endpoint string, params url.Values, result interface{}) error
	///Delete() *Client
}

// NewClient creates a new httpClient for the given apiConfig.
func NewClient(store string, apiConfig *ApiConfig) (*Client, error) {
	storeURL, err := url.Parse(store)
	if err != nil {
		return nil, err
	}

	if apiConfig == nil {
		apiConfig = &ApiConfig{}
	}
	if apiConfig.OauthTimestamp.IsZero() {
		apiConfig.OauthTimestamp = time.Now()
	}

	if apiConfig.Version == "" {
		apiConfig.Version = "v2"
	}
	path := "/wp-json/wc"
	if apiConfig.API {
		path = apiConfig.APIPrefix
	}

	storeURL.Path = path + "/" + apiConfig.Version

	rawClient := http.DefaultClient

	// Set client's transport to default when nil
	if rawClient.Transport == nil {
		rawClient.Transport = http.DefaultTransport
	}

	if transport, ok := rawClient.Transport.(*http.Transport); ok && transport != nil {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: apiConfig.VerifySSL}
	}

	c := &Client{
		storeURL: storeURL,

		apiConfig: apiConfig,
		rawClient: rawClient,
	}

	return c, nil
}

func (c *Client) basicAuth(params url.Values) string {
	params.Add("consumer_key", c.apiConfig.ConsumerKey)
	params.Add("consumer_secret", c.apiConfig.ConsumerSecret)
	return params.Encode()
}

func (c *Client) oauth(method, urlStr string, params url.Values) string {
	params.Add("oauth_consumer_key", c.apiConfig.ConsumerKey)
	params.Add("oauth_timestamp", strconv.Itoa(int(c.apiConfig.OauthTimestamp.Unix())))
	nonce := make([]byte, 16)
	rand.Read(nonce)
	sha1Nonce := fmt.Sprintf("%x", sha1.Sum(nonce))
	params.Add("oauth_nonce", sha1Nonce)
	params.Add("oauth_signature_method", HashAlgorithm)
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paramStrs []string
	for _, key := range keys {
		paramStrs = append(paramStrs, fmt.Sprintf("%s=%s", key, params.Get(key)))
	}
	paramStr := strings.Join(paramStrs, "&")
	params.Add("oauth_signature", c.oauthSign(method, urlStr, paramStr))
	return params.Encode()
}

func (c *Client) oauthSign(method, endpoint, params string) string {
	signingKey := c.apiConfig.ConsumerSecret
	if c.apiConfig.Version != "v1" && c.apiConfig.Version != "v2" {
		signingKey = signingKey + "&"
	}

	a := strings.Join([]string{method, url.QueryEscape(endpoint), url.QueryEscape(params)}, "&")
	mac := hmac.New(sha256.New, []byte(signingKey))
	mac.Write([]byte(a))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes)
}

func (c *Client) newRequest(method, endpoint string, params url.Values, data interface{}) (*http.Request, error) {
	urlstr := c.storeURL.String() + "/" + strings.TrimLeft(endpoint, "/")
	if params == nil {
		params = make(url.Values)
	}

	var dlmtr = "?"
	if strings.ContainsAny(urlstr, "?&") {
		dlmtr = "&"
	}

	if c.storeURL.Scheme == "https" {
		urlstr += dlmtr + c.basicAuth(params)
	} else {
		urlstr += dlmtr + c.oauth(method, urlstr, params)
	}

	switch method {
	case http.MethodGet:
		return http.NewRequest(method, urlstr, nil)

	case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions:
	default:
		return nil, fmt.Errorf("method is not recognised: %s", method)
	}

	body := new(bytes.Buffer)
	encoder := json.NewEncoder(body)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	return http.NewRequest(method, urlstr, body)
}

func (c *Client) request(ctx context.Context, method, endpoint string, params url.Values, data interface{}) (io.ReadCloser, error) {
	req, err := c.newRequest(method, endpoint, params, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.rawClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusBadRequest ||
		resp.StatusCode == http.StatusUnauthorized ||
		resp.StatusCode == http.StatusNotFound ||
		resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("Request failed: %s", resp.Status)
	}
	return resp.Body, nil
}

func (c *Client) Post(ctx context.Context, endpoint string, data interface{}) (io.ReadCloser, error) {
	return c.request(ctx, "POST", endpoint, nil, data)
}

func (c *Client) Put(ctx context.Context, endpoint string, data interface{}) (io.ReadCloser, error) {
	return c.request(ctx, "PUT", endpoint, nil, data)
}

func (c *Client) Get(ctx context.Context, endpoint string, params url.Values, result interface{}) error {

	response, err := c.request(ctx, "GET", endpoint, params, nil)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return err

	} else {

		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}

	}
	return nil
}

func (c *Client) Delete(ctx context.Context, endpoint string, params url.Values) (io.ReadCloser, error) {
	return c.request(ctx, "POST", endpoint, params, nil)
}

func (c *Client) Options(ctx context.Context, endpoint string) (io.ReadCloser, error) {
	return c.request(ctx, "OPTIONS", endpoint, nil, nil)
}
