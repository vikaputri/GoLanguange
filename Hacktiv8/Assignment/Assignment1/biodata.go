package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		i, err := strconv.Atoi(arg)
		fmt.Println("Nomer Absen : ", i)
		if err != nil {
			panic(err)
		}
		data_orang(i - 1)
	} else {
		fmt.Println("Nomor absen tidak dikenali")
	}

}

func data_orang(peserta int) {
	data := []Student{
		{"Vika", "Pekalongan", "Mahasiswa", "Ketertarikan dengan Backend"},
		{"Putri", "Depok", "Mahasiswa", "Ingin belajar hal baru"},
		{"Ariyanti", "Jakarta Selatan", "Mahasiswa", "Menambah wawasan"},
		{"Della", "Batang", "Mahasiswa", "Menambah ilmu baru"},
		{"Oktavia", "Yogyakarta", "Mahasiswa", "Menambah pengetahuan"},
		{"Hanif", "Jambi", "Mahasiswa", "Ketertarikan dengan backend"},
		{"Akbar", "Lombok", "Mahasiswa", "Menambah ilmu baru"},
		{"Pranaja", "Bali", "Mahasiswa", "Ingin belajar hal baru"},
		{"Maira", "Semarang", "Mahasiswa", "Menambah pengetahuan"},
		{"Yaya", "Padang", "Mahasiswa", "Menambah wawasan"},
		{"Aas", "Palembang", "Mahasiswa", "Ketertarikan dengan backend"},
		{"Datul", "Makasar", "Mahasiswa", "Ingin belajar hal baru"},
		{"Naris", "Jakarta Barat", "Mahasiswa", "Menambah wawasan"},
		{"Cici", "Jakarta Timur", "Mahasiswa", "Menambah ilmu baru"},
		{"Lia", "Jakarta Pusat", "Mahasiswa", "Menambah pengetahuan"},
		{"Linda", "Jakarta Utara", "Mahasiswa", "Ketertarikan dengan backend"},
		{"Faeza", "Pemalang", "Mahasiswa", "Menambah ilmu baru"},
		{"Asca", "Solo", "Mahasiswa", "Ingin belajar hal baru"},
		{"Arjuna", "Cirebon", "Mahasiswa", "Menambah pengetahuan"},
		{"Humaira", "Bogor", "Mahasiswa", "Menambah wawasan"},
	}

	if peserta >= 0 && peserta < len(data) {
		fmt.Printf("Absen %d yaitu %s, dia adalah seorang %s yang bertempat tinggal di %s, memilih kelas golang karena: %s",
			peserta+1, data[peserta].name, data[peserta].job, data[peserta].address, data[peserta].reason)
	} else {
		fmt.Printf("Tidak ada nomer absen %d", peserta+1)
	}

}
