package generation

import (
	"encoding/json"
	"log"
	"os"

	"codeberg.org/neurocollective/wuxia/structs"
)

const CREATE = "CREATE"
const TABLE = "TABLE"

func ReadDump() (Schema, error) {
	fileBytes, err := os.ReadFile("dumps/dump.sql")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	schema := Schema{}

	fileString := string(fileBytes)

	for _, line := range strings.Split(fileString, "\n") {
		for _, token := range strings.Split(line, " ") {
			if token == CREATE {

			}
		}
	}

	return schema, nil
}
