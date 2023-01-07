package api

import "context"

type Restchai interface {
	// response is data that has been parsed into string
	GetChat(ctx context.Context, msg string) string
}
