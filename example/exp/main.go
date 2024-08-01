package main

import (
	"fmt"
	"log"
	"strings"

	link "github.com/Omkar-Patil16/HTML-Link-Parser"
)

var basicHtml = `
<a href="https://www.w3schools.com">
<img border="0" alt="W3Schools" src="logo_w3s.gif" width="100" height="100">
</a`

func main() {
	//converting string into a reader for Parser
	r := strings.NewReader(basicHtml)
	doc, err := link.Parser(r)
	if err != nil {
		log.Printf("%+v", err)
	}
	fmt.Println(doc)
}
