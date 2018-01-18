package main

import (
	"encoding/base64"
	"fmt"
)

func Encode(d []byte) []byte {
	sEnc := base64.StdEncoding.EncodeToString(d)
	a := Bar(sEnc)
	return a // do it!
}

func Decode(d []byte) []byte {
	a := Foo(d)
	sDec, _ := base64.StdEncoding.DecodeString(a)
	return sDec // do it!
}

func Bar(a string) []byte {
	return []byte(a)
}

func Foo(s []byte) string {
	return string(s)
}

func main() {
	fmt.Println(Encode([]byte("test")))
	fmt.Println(Decode([]byte("fPNKd")))

}
