package zendesk

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURLFormat = "https://%s.zendesk.com/api/v2"
)

//Client Zendesk Client
type Client struct {
	baseURL    string // ....zendesk.com/api/v2
	credential string
	client     http.Client
}

//NewClientByPassword Return Client With Information by Password Auth
func NewClientByPassword(baseURL, userAgent, passWD string) (*Client, error) {

	baseURLString := fmt.Sprintf(baseURLFormat, baseURL)
	u, err := url.Parse(baseURLString)
	if err != nil {
		return nil, err
	}
	baseURL = u.String()

	credential := base64.StdEncoding.EncodeToString([]byte(userAgent + ":" + passWD))

	return &Client{
		baseURL:    baseURL,
		credential: credential,
		client:     http.Client{},
	}, nil
}

//NewClientByToken Return Client With Information by Token Auth
func NewClientByToken(baseURL, userAgent, token string) (*Client, error) {

	baseURLString := fmt.Sprintf(baseURLFormat, baseURL)
	u, err := url.Parse(baseURLString)
	if err != nil {
		return nil, err
	}
	baseURL = u.String()

	credential := base64.StdEncoding.EncodeToString([]byte(userAgent + "/token:" + token))

	return &Client{
		baseURL:    baseURL,
		credential: credential,
		client:     http.Client{},
	}, nil
}

//Get Get Request to Url
func (c *Client) Get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", "Basic "+c.credential)

	resp, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Bad response from api")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
