package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nomor     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	biodatas := []Biodata{
		{1, "Riyadi F", "Banjarmasin", "Programmer", "Saya ingin menjadi seorang programmer."},
		{2, "Elvin Natsir", "Palangkaraya", "Desain grafis", "Saya ingin programmer."},
		{3, "Lissa A", "Pontianak", "Programmer", "Saya ingin menjadi programmer."},
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go [nomor]")
		return
	}

	nomor, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Nomor harus berupa bilangan bulat.")
		return
	}

	var biodata Biodata
	for _, b := range biodatas {
		if b.Nomor == nomor {
			biodata = b
			break
		}
	}

	if biodata.Nomor != 0 {
		fmt.Printf("Nomor: %d\n", biodata.Nomor)
		fmt.Printf("Nama: %s\n", biodata.Nama)
		fmt.Printf("Alamat: %s\n", biodata.Alamat)
		fmt.Printf("Pekerjaan: %s\n", biodata.Pekerjaan)
		fmt.Printf("Alasan: %s\n", biodata.Alasan)
	} else {
		fmt.Println("Biodata tidak ditemukan.")
	}
}
