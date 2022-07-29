package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your database number")

	flag.Parse()
	fmt.Println("Host:", *host)
	fmt.Println("Username", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Number:", *number) //go run packageFlag.go -number=nama
}
