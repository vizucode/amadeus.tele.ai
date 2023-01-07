package localstorage

type Localstorage interface {
	Read(fileName string) string
	Write(fileName string, data string)
}
