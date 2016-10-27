package middleware

import (
	"context"
	"net/http"
)

type Adapter struct {
	ctx     context.Context
	handler Handler
}

func (this *Adapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) { // {{{
	this.handler.ServeHTTPCtx(this.ctx, rw, req)
} // }}}
