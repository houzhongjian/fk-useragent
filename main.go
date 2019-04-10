package main

import (
	"log"

	"uaparse.com/uaparse/ua"
)

func main() {
	uas := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	rs := ua.Parse(uas)
	log.Printf("rs:%+v\n", rs)
}
