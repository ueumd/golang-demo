package main

/*
 * 加密敏感配置参数
 *
 * wencan
 * 2018-08-31
 */

import (
	"fmt"
	"os"
	"myapiserver/enciphertex"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("no arguments")
		return
	}

	plain := os.Args[1]
	crypted, err := cryptex.Cryptex.Encode(plain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(crypted)
}
