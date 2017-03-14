package dictionary

//Dictionary is the external source from where data can be extracted
type Dictionary interface {
	Print(word string) error
}
