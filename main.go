package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/scritchley/orc"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printHelp()
		os.Exit(0)
	}

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Missing arguments!")
		fmt.Fprintln(os.Stderr, args)
		printHelp()

		os.Exit(1)
	}

	action := strings.ToLower(args[0])
	inputfilepath := args[1]

	orcReader, err := orc.Open(inputfilepath)
	if err != nil {
		fmt.Fprint(os.Stderr, "Could not read file: "+inputfilepath+" "+err.Error())
		os.Exit(2)
	}
	defer orcReader.Close()

	switch action {
	case "read":
		readRows(*orcReader)
		break
	case "readjson":
		readRowsJson(*orcReader)
		break
	case "schema":
		printSchema(*orcReader)
		break
	case "schemajson":
		printSchemaJson(*orcReader)
		break
	case "schemasql":
		printSchemaSQL(*orcReader)
		break
	case "meta", "metadata":
		printMetadata(*orcReader)
		break
	case "count":
		printCount(*orcReader)
		break
	}

	// TODO SCHEMA SQL
}

func printHelp() {
	fmt.Println("-- O R C R e a d e r 1337 -- ")
	fmt.Println("Usage: ./orcreader <action> <inputfile>")
	fmt.Println("Available actions: read, schema, schemajson, schemasql, readjson, metadata, count")
	fmt.Println("Example: ./orcreader read myfile.orc")
}

// Prints the schema as sql create table statement
func printSchemaSQL(orcReader orc.Reader) {

	fmt.Print("CREATE TABLE mytable (\n")

	for index, value := range orcReader.Schema().Columns() {
		switch orcReader.Schema().Types()[index+1].Kind.String() {
		case "BOOLEAN":
			fmt.Print("\t", value, " ", "BOOLEAN")
			break
		case "SHORT":
			fmt.Print("\t", value, " ", "INTEGER")
			break
		case "INT":
			fmt.Print("\t", value, " ", "INTEGER")
			break
		case "LONG":
			fmt.Print("\t", value, " ", "BIGINT")
			break
		case "FLOAT":
			fmt.Print("\t", value, " ", "DOUBLE")
			break
		case "DOUBLE":
			fmt.Print("\t", value, " ", "DOUBLE")
			break
		case "TIMESTAMP":
			fmt.Print("\t", value, " ", "TIMESTAMP")
			break
		case "DATE":
			fmt.Print("\t", value, " ", "DATE")
			break
		case "STRING":
			fmt.Print("\t", value, " ", "VARCHAR")
			break
		case "BINARY":
			fmt.Print("\t", value, " ", "VARBINARY")
			break
		case "VARCHAR":
			fmt.Print("\t", value, " ", "VARCHAR")
			break
		default:
			fmt.Print("\t", value, " ", orcReader.Schema().Types()[index+1].Kind.String())
			break
		}

		if index != len(orcReader.Schema().Columns())-1 {
			fmt.Println(", ")
		}
	}

	fmt.Print("\n)\nWITH (format = 'ORC')")
}

// Prints metadata aobut the orc file
func printMetadata(orcReader orc.Reader) {
	fmt.Print(strings.ReplaceAll(strings.ReplaceAll(orcReader.Metadata().String(), ">", ">\n"), "<", "\n<"))
}

// prints the amount of rows in the file
func printCount(orcReader orc.Reader) {
	fmt.Print(orcReader.NumRows())
}

// Prints the schema as orc schema
func printSchema(orcReader orc.Reader) {
	schema := orcReader.Schema()
	fmt.Print(schema)
}

// Prints the schema as JSON
func printSchemaJson(orcReader orc.Reader) {
	schema := orcReader.Schema()
	fmt.Print(schema.ToJSON())
}

// Prints rows from input orc file
func readRows(orcReader orc.Reader) {

	cursor := orcReader.Select(orcReader.Schema().Columns()...)

	for cursor.Stripes() {

		for cursor.Next() {
			fmt.Println(cursor.Row())
		}

	}
	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(3)
	}
}

// Prints rows from input orc file
// TODO: Handle rows and arrays
func readRowsJson(orcReader orc.Reader) {

	cursor := orcReader.Select(orcReader.Schema().Columns()...)

	for cursor.Stripes() {

		for cursor.Next() {

			fmt.Print("{")
			for index, value := range cursor.Row() {

				fmt.Print("\"" + orcReader.Schema().Columns()[index] + "\": ")

				// Pluss en fordi kind på plass 0 er struct
				switch orcReader.Schema().Types()[index+1].Kind.String() {
				case "SHORT", "INT", "LONG", "FLOAT", "DOUBLE":
					fmt.Print(value)
					break
				default:
					fmt.Print("\"", value, "\"")
					break
				}

				if index != len(orcReader.Schema().Columns())-1 {
					fmt.Print(", ")
				}

			}
			fmt.Print("}\n")
		}
	}

	if err := cursor.Err(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(3)
	}
}
