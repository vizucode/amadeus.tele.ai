package translate

import (
	gt "github.com/bas24/googletranslatefree"
)

func Translate(msg string, sourceLang string, toLang string) (string, error) {
	result, err := gt.Translate(msg, sourceLang, toLang)
	return result, err
}
