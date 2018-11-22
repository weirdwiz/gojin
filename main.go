package main

import (
	"encoding/json"
	"fmt"

	goose "github.com/advancedlogic/GoOse"
)

func main() {
	g := goose.New()

	article, _ := g.ExtractFromURL("https://www.inkandswitch.com/slow-software.html")
	b, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
