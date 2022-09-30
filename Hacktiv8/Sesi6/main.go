package main

import "belajargolang/Hacktiv8/Sesi6/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
