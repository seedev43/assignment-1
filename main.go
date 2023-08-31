package main

import (
	"assignment-1/lib"
	"assignment-1/model"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := lib.ReadJSON("participants.json")
	participantOfMap := map[string]model.Data{}

	for _, participant := range data {
		participantOfMap[participant.Code] = participant
	}

	// strings.ToUpper untuk mengubah inputan menjadi huruf besar
	input := strings.ToUpper(os.Args[1])
	fmt.Println("=== CARI DATA PESERTA DENGAN KODE PESERTA ===")
	if dataExist, found := participantOfMap[input]; found {
		fmt.Println("ID:", dataExist.Id)
		fmt.Println("Nama:", dataExist.Name)
		fmt.Println("Alamat:", dataExist.Address)
		fmt.Println("Pekerjaan:", dataExist.Occupation)
		fmt.Println("Alasan:", dataExist.Reason)
	} else {
		fmt.Println("Data dengan kode " + input + " tidak ditemukan")
	}
}
