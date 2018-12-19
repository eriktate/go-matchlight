package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const matchlightAuthKey = "X-Matchlight-Auth"
const baseURL = "https://api.terbiumlabs.com/api/v2/"

type apiResult struct {
	Count uint            `json:"count,omitempty"`
	Data  json.RawMessage `json:"data"`
}

// A Client knows how to make requests to the Matchlight API.
type Client struct {
	Key    string
	Secret string

	http http.Client
}

// NewClient returns a new Client instance.
func NewClient(key, secret string) Client {
	return Client{
		Key:    key,
		Secret: secret,
	}
}

func (c Client) get(url string, out interface{}) error {
	res, err := c.getRaw(url)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(res, out); err != nil {
		return errors.Wrap(err, "failed to unmarshal result data")
	}

	return nil
}

func (c Client) getRaw(url string) ([]byte, error) {
	url = fmt.Sprintf("%s%s", baseURL, url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(matchlightAuthKey, fmt.Sprintf("%s:%s", c.Key, c.Secret))
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, string(body))
	}

	return body, nil
}

func (c Client) post(url string, in interface{}, out interface{}) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return err
	}

	res, err := c.postRaw(url, payload)
	if err != nil {
		return err
	}

	var result apiResult
	if err := json.Unmarshal(res, &result); err != nil {
		return err
	}

	if err := json.Unmarshal(result.Data, out); err != nil {
		return err
	}

	return nil
}

func (c Client) postRaw(url string, payload []byte) ([]byte, error) {
	url = fmt.Sprintf("%s%s", baseURL, url)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add(matchlightAuthKey, fmt.Sprintf("%s:%s", c.Key, c.Secret))
	req.Header.Add("Content-Type", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, string(body))
	}

	return body, nil
}

func (c Client) delete(url string) ([]byte, error) {
	url = fmt.Sprintf("%s%s", baseURL, url)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(matchlightAuthKey, fmt.Sprintf("%s:%s", c.Key, c.Secret))
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, string(body))
	}

	return body, nil
}

// Matchlight aggregates all of the disparate clients into a single implementation.
type Matchlight struct {
	ProjectClient
}

// NewMatchlight returns a new instance of Matchlight with an API key and secret.
func NewMatchlight(key, secret string) Matchlight {
	c := Client{
		Key:    key,
		Secret: secret,
	}

	return Matchlight{
		ProjectClient{c},
	}
}
