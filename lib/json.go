package lib

import (
	"assignment-1/model"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func ReadJSON(filename string) map[string]model.Data {
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

	participantOfMap := map[string]model.Data{}

	for _, participant := range participants.Participants {
		participantOfMap[participant.Code] = participant
	}

	return participantOfMap

}