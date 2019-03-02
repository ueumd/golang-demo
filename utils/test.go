package main

import (
	"fmt"
	"log"
	"myapiserver/utils/encipher/cryptex"
)

func main()  {

	token, _ := cryptex.Cryptex.Encode("abcd")
	log.Println(token)
	fmt.Println(token)

	xtoken, _ := cryptex.Cryptex.Decode(token)
	fmt.Println(xtoken)


}