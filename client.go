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

var (
	// Only reason we have this as a var is to be able to change it during
	// tests. I admit this is a tad lazy and the baseEndpoint should live on
	// the Client struct :))
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
	c      *http.Client
	secret string
}

func New(opts ...Option) (*Client, error) {

	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if IsStringEmpty(c.secret) {
		return nil, errors.New("please provide your secret key")
	}

	if c.c == nil {
		c.c = &http.Client{
			Transport: &basicAuthransport{
				originalTransport: http.DefaultTransport,
				secret:            c.secret,
			},
			Timeout: time.Second * 5,
		}
	}

	c.secret = ""

	return c, nil
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

	if IsStringEmpty(opts.Reference) {
		return IndexedUser{}, errors.New("please provide the user reference")
	}

	var p = new(indexedUserResponse)

	var buf = new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(&opts); err != nil {
		return IndexedUser{}, err
	}

	r, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("%s/users/update/%s", baseEndpoint, opts.Reference), buf)
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
