package structs

import (
	"testing"
)

func TestSQLArgSequence(t *testing.T) {

	expected := "type TestStruct struct {\n  Id *lib.NotNull[int]\n  Test *sql.Null[string]\n}"

	columns := []ColumnDefinition{
		ColumnDefinition{ "id","integer",false,""},
		ColumnDefinition{ "test","text",true,""},
	}

	tableSchema := TableSchema{ Name: "TestStruct", Columns: columns}

	string := tableSchema.GetStructString()

	if string != expected {
		t.Log("got:\n", string)
		t.Log("expected:\n", expected)
		t.Fatal("incorrect result")
	}

}