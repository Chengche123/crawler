package main

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
)

func main() {
	u := "http://images.dmzj1.com/webpic/19/jiubaotongxuebufangguowojsi.jpg"
	p, err := url.Parse(u)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Printf("%#v\n", p)
	fmt.Println(filepath.Base(p.Path))
}
