package util

import (
	"io/ioutil"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) []byte {

	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + path)

	Check(err)

	return dat
}

type Pair struct {
	a, b interface{}
}
