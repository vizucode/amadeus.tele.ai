package api

type Restchai interface {
	// response is data that has been parsed into string
	GetChat(msg string) string
}
