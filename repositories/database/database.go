package database

type Database interface {
	Read(fileName string) string
	Write(fileName string, data string)
}
