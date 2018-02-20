# go-simpro

Generated, designed to be extended upon later.

Usage:
```
package main

import (
	"./generated"
	"github.com/arran4/go-rfc5849-hmac"
	"log"
)

func main() {
	go_rfc5849_hmac.PublicKey = 
	go_rfc5849_hmac.SecretKey = 
	r, err := simpro.CompanySearch("")
	if err != nil {
		log.Panicf("Error :%v", err)
	}
	log.Printf("%#v", r)
}
```
