package strest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nodeselector/go-stuff/sthooks"
	"github.com/nodeselector/go-stuff/sthttp"
)

type Client struct {
	c     sthttp.TestableClient
	hooks *sthooks.ClientHooks
}

type defaultAttempter struct{}

func (d defaultAttempter) Attempt(a sthttp.Attempt) error {
	return a()
}

func (c *Client) Do(ctx context.Context, req *http.Request, dst interface{}, opts ...Option) error {
	ctx = c.hooks.OnBegin(ctx)
	defer func() {
		c.hooks.OnEnd(ctx)
	}()

	do := &Options{
		Attempter: &defaultAttempter{},
	}

	for _, opt := range opts {
		opt(do)
	}

	attempt := func() error {
		resp, err := c.handleRequest(ctx, do, req)
		if err != nil {
			ctx = sthooks.WithError(ctx, err)
			return err
		}

		err = c.handleResp(ctx, do, resp, dst)
		if err != nil {
			ctx = sthooks.WithError(ctx, err)
			return err
		}

		return nil
	}

	err := do.Attempter.Attempt(attempt)
	if err != nil {
		ctx = sthooks.WithError(ctx, err)
		return err
	}

	return nil
}

func (c *Client) handleRequest(ctx context.Context, opts *Options, req *http.Request) (*http.Response, error) {
	ctx = c.hooks.OnStartPerformRequest(ctx)
	defer func() {
		c.hooks.OnDonePerformRequest(ctx)
	}()

	resp, err := c.c.Do(req)
	if err != nil {
		ctx = sthooks.WithError(ctx, err)
		return nil, err
	}

	return resp, err
}

func (c *Client) handleResp(ctx context.Context, opts *Options, resp *http.Response, dst interface{}) error {
	ctx = c.hooks.OnStartHandleResponse(ctx)
	defer func() {
		c.hooks.OnDoneHandleResponse(ctx)
	}()

	if dst != nil {
		err := json.NewDecoder(resp.Body).Decode(dst)
		if err != nil {
			ctx = sthooks.WithError(ctx, err)
			return err
		}
	}

	return nil
}
