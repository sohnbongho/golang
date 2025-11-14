package main

import (
	"example03/src/string_util"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(string_util.Reverse("abcde"))
}
