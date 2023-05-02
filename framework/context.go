package framework

import (
	"context"
	"net/http"
	"time"
)

// Context is a more powerful struct base on the official Context
// Context can handle request and send response
type Context struct {
	request *http.Request
}

func (c *Context) BaseContext() context.Context {
	return c.request.Context()
}

// implement  context.Context

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.BaseContext().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.BaseContext().Done()
}

func (c *Context) Err() error {
	return c.BaseContext().Err()
}

func (c *Context) Value(key any) any {
	return c.BaseContext().Value(key)
}
