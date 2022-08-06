package gorounties

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Vika Putri" //hang kalo ada yang ngirim tapi tidak diterima
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel //deadlock kalo ada yang nerima tapi tidak ada yang ngirim
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
