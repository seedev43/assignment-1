package main

import (
	"assignment-1/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ConvertJSONToObject(filename string) []model.Data {
	// path file
	jsonPath := filepath.Join("./", filename)

	// membaca isi dari file JSON
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	var participants model.Participants

	// mengubah data json menjadi object
	if err := json.Unmarshal(data, &participants); err != nil {
		log.Fatal(err)
	}

	return participants.Participants

}

func main() {
	data := ConvertJSONToObject("participants.json")
	participantOfMap := map[string]model.Data{}

	for _, participant := range data {
		participantOfMap[participant.Code] = participant
	}

	// strings.ToUpper untuk mengubah inputan menjadi huruf besar
	input := strings.ToUpper(os.Args[1])

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
