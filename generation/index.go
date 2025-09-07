package generation

import (
	"os"
	"log"
	"encoding/json"
)

 func ReadDump() {
 	fileBytes, err := os.ReadFile("dumps/dump.sql")

 	if err != nil {
 		log.Println(err)
 		os.Exit(1)
 	}

}