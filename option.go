package thepeer

import "net/http"

type Option func(c *Client)

func WithHTTPClient(cl *http.Client) Option {
	return func(c *Client) {
		c.c = cl
	}
}

func WithSecretKey(s string) Option {
	return func(c *Client) {
		c.secret = s
	}
}
