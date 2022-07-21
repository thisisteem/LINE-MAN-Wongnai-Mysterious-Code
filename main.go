package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	fmt.Println(string(sd))
}

// After decode the secret string we will got => iangnoW:NAM:ENIL:ta:su:nioJ
// And then if we reverse the text we will got => Join:us:at:LINE:MAN:Wongnai
// So the answer is "Join:us:at:LINE:MAN:Wongnai" !!!