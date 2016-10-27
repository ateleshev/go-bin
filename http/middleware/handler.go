package middleware

import (
	"context"
	"net/http"
)

type Handler interface {
	ServeHTTPCtx(context.Context, http.ResponseWriter, *http.Request)
}

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (self HandlerFunc) ServeHTTPCtx(ctx context.Context, rw http.ResponseWriter, req *http.Request) { // {{{
	self(ctx, rw, req)
} // }}}
