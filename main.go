package main

import (
	"assignment-1/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := lib.ReadJSON("participants.json")
	// strings.ToUpper untuk mengubah inputan menjadi huruf besar
	input := strings.ToUpper(os.Args[1])
	fmt.Println("=== CARI DATA PESERTA DENGAN KODE PESERTA ===")
	if dataExist, found := data[input]; found {
		fmt.Println("ID:", dataExist.Id)
		fmt.Println("Nama:", dataExist.Name)
		fmt.Println("Alamat:", dataExist.Address)
		fmt.Println("Pekerjaan:", dataExist.Occupation)
		fmt.Println("Alasan:", dataExist.Reason)
	} else {
		fmt.Println("Data dengan kode " + input + " tidak ditemukan")
	}
}
