// error interfaceを満たすMyError型を定義

package main

import (
	"fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

// MyError型をerror型として返す
func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// error型を型アサーションしてMyError構造体を取り出す
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
