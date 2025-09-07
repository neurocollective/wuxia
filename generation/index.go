package generation

import (
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"codeberg.org/neurocollective/wuxia/structs"
)

const TEXT = "text"
const VARCHAR = "varchar"
const INTEGER = "integer"
const NUMERIC = "numeric"
const TIMESTAMP = "timestamp"

const SPACE = " "
const CREATE = "CREATE"
const TABLE = "TABLE"
const END = ");"
const NOT = "NOT"
const NULL = "NULL"

func CleanToken(token string) string {
	tokenNoSpaces := strings.TrimSpace(token)
	cleanedToken := strings.Replace(tokenNoSpaces, "\t", "", -1)
	return cleanedToken
}

func isAType(token string) bool {
	return token == TEXT || token == VARCHAR || token == INTEGER || token == NUMERIC || token == TIMESTAMP
}

func ReadDump() (structs.Schema, error) {
	fileBytes, err := os.ReadFile("dumps/schema.sql")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// columns := make([]structs.ColumnDefinition, 0, 100)

	var columns []structs.ColumnDefinition
	tables := make([]structs.TableSchema, 0, 100)
	schema := structs.Schema{}

	fileString := string(fileBytes)

	//fmt.Println("fileString: ", fileString)

	var scanningTable bool
	var columnIndex int = -1
	var tableIndex int = -1
	var currentColumn structs.ColumnDefinition
	var currentTable structs.TableSchema

	for _, line := range strings.Split(fileString, "\n") {

		tokens := strings.Split(line, " ")

		for index, token := range tokens {

			cleanedToken := CleanToken(token)

            if cleanedToken == "" {
                continue
            }

			var nextToken string
			if index < len(tokens)-1 {
				nextToken = CleanToken(tokens[index+1])
			}

			if !scanningTable && index == 0 && cleanedToken == CREATE {

				if nextToken == TABLE {
                    fmt.Println("NEW TABLE, token:", cleanedToken)
					// this is a table creation
					scanningTable = true
					columns = make([]structs.ColumnDefinition, 0, 100)
					currentColumn = structs.ColumnDefinition{Nullable: true}
					tableIndex += 1
					columnIndex = -1
					currentTable = structs.TableSchema{}
					tables = append(tables, currentTable)
                    continue
				} else {
                    fmt.Println("what??")
                }
			} else if scanningTable {
                fmt.Println("Scanning table, token:", cleanedToken)  
                if cleanedToken == END {
                    fmt.Println("END OF COLUMN, token:", cleanedToken, "index:", index)               
                    scanningTable = false
                    columns = append(columns, currentColumn)
                    currentTable.Columns = columns[:columnIndex+1]
                    columnIndex = -1
                } else if index == 0 {
                    fmt.Println("NEW COLUMN, token:", cleanedToken, "index:", index)
					columnIndex += 1
					currentColumn.Name = cleanedToken
				} else if isAType(cleanedToken) {
					currentColumn.Type = cleanedToken
				} else if cleanedToken == NOT && nextToken == NULL {
					currentColumn.Nullable = false
				}
			} else {
                fmt.Println("WTF, token:", cleanedToken)    
            }
		}
	}

	// schema.Columns = columns[:currentColumn]
	schema.Tables = tables[:len(tables)]
	return schema, nil
}

func WriteStructsToGeneratedFolder() {

}
