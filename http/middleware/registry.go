package middleware

type Registry interface {
	Append(Middleware)
}
