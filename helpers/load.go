package helpers

import (
	"io/ioutil"
)

func MustLoadFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
