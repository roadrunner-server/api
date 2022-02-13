package middleware

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// Middleware represents http stdlib middleware interface
type Middleware interface {
	Middleware(f http.Handler) http.Handler
}

// FastHTTPMiddleware represents fasthttp middleware interface
type FastHTTPMiddleware interface {
	FastHTTPMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler
}
