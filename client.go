package thepeer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	baseEndpoint = "https://api.thepeer.co"
)

type basicAuthransport struct {
	originalTransport http.RoundTripper
	secret            string
}

func (c *basicAuthransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("x-api-key", c.secret)
	return c.originalTransport.RoundTrip(r)
}

type Client struct {
	c *http.Client
}

func New(c *http.Client) *Client {

	if c == nil {
		c = &http.Client{
			Transport: &basicAuthransport{
				originalTransport: http.DefaultTransport,
			},
			Timeout: time.Second * 5,
		}
	}

	return &Client{c: c}
}

func (c *Client) SendReceipt(receipt string) (*Transaction, error) {
	var p = new(Transaction)

	r, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/send/%s", baseEndpoint, receipt),
		strings.NewReader("{}"))

	if err != nil {
		return p, err
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {

		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return nil, err
		}

		return p, errors.New(s.Message)
	}

	return p, json.NewDecoder(resp.Body).Decode(p)
}

func (c *Client) DeIndexUser(opts *DeIndexUserOptions) error {

	var buf = new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(&opts); err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("%s/users/delete/%s", baseEndpoint, opts.UserReference), buf)
	if err != nil {
		return err
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {
		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return err
		}

		return errors.New(s.Message)
	}

	return nil
}

func (c *Client) UpdateUser(opts *UpdateUserOptions) (IndexedUser, error) {

	var p = new(indexedUserResponse)

	var buf = new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(&opts); err != nil {
		return IndexedUser{}, err
	}

	r, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/users/update/%s", baseEndpoint, opts.Identifier), buf)
	if err != nil {
		return IndexedUser{}, err
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return IndexedUser{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {
		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return IndexedUser{}, err
		}

		return IndexedUser{}, errors.New(s.Message)
	}

	return p.IndexedUser, json.NewDecoder(resp.Body).Decode(p)
}

func (c *Client) IndexUser(opts *IndexUserOptions) (IndexedUser, error) {

	var p = new(indexedUserResponse)

	var buf = new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(&opts); err != nil {
		return IndexedUser{}, err
	}

	r, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/users/index", baseEndpoint), buf)
	if err != nil {
		return IndexedUser{}, err
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return IndexedUser{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {
		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return IndexedUser{}, err
		}

		return IndexedUser{}, errors.New(s.Message)
	}

	return p.IndexedUser, json.NewDecoder(resp.Body).Decode(p)
}

func (c *Client) FetchReceipt(id string) (*Receipt, error) {

	var p = new(receiptResponse)

	r, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/verify/%s", baseEndpoint, id), nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusCreated {
		var s struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return nil, err
		}

		return nil, errors.New(s.Message)
	}

	return &p.Receipt, json.NewDecoder(resp.Body).Decode(p)
}
