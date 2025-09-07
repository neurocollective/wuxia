package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"codeberg.org/neurocollective/wuxia/structs"
	"codeberg.org/neurocollective/wuxia/lib"
)

func BuildPostgresClient(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connectionString)
	if err != nil {

		log.Println("ERROR opening postgres connection with github.com/neurocollective/go_utils.BuildPostgresClient() ->")
		log.Println(err.Error())

		return nil, err
	}

	return db, nil
}

func main() {

	client, err := BuildPostgresClient("user=postgres password=postgres dbname=postgres sslmode=disable")

	if err != nil {
		fmt.Println("error connecting to be", err)
		os.Exit(1)
	}

	// begin insert

	expenditure := structs.Expenditure{
		UserId: &lib.NotNull[int]{1},
		CategoryId: &sql.Null[int]{0,false},
		Value: &lib.NotNull[float32]{45.99},
		Description: &lib.NotNull[string]{"stuff"},
	}

	err = structs.InsertExpenditure(client, expenditure)

	if err != nil {
		fmt.Println("error inserting", err)
		os.Exit(1)
	}

	log.Println("inserted??")

	// end insert

	var e structs.Expenditure

	queryString := "select " + e.ColumnsString() + " from " + e.TableName() + ";"
	log.Println("queryString:", queryString)

	//args := []any{}

	rows, err := client.Query(queryString)

	if err != nil {
		fmt.Println("error connecting to be", err)
		os.Exit(1)
	}

	expenditures, err := structs.ReceiveExpenditures(rows)

	if err != nil {
		fmt.Println("error receiving rows", err)
		os.Exit(1)
	}

	for _, expenditure := range expenditures {
		expenditure.Print()
	}

}
