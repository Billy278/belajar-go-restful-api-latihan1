package helper

import "fmt"

func PanicHandler(err error) {
	if err != nil {
		fmt.Println("Cek error", err)
		panic(err)
	}
}
