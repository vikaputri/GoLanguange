package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	msgl1 := <-c
	fmt.Println(msgl1)

	msgl2 := <-c
	fmt.Println(msgl2)

	msgl3 := <-c
	fmt.Println(msgl3)
	close(c)

}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
