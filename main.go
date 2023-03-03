package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// Define command line flags
	jsonFile := flag.String("json", "people.json", "the path to the input JSON file")
	csvFile := flag.String("csv", "people.csv", "the path to the output CSV file")
	flag.Parse()

	// Validate command line arguments
	if *jsonFile == "" {
		fmt.Println("Error: You must specify an input JSON file with the -json flag")
		return
	}
	if *csvFile == "" {
		fmt.Println("Error: You must specify an output CSV file with the -csv flag")
		return
	}

	// Read the JSON data from the file
	jsonData, err := os.ReadFile(*jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON data into a slice of maps
	var data []map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract the header row from the data
	header := make([]string, 0)
	for k := range data[0] {
		header = append(header, k)
	}

	// Convert the data to a 2D slice of strings
	rows := make([][]string, len(data)+1)
	rows[0] = header
	for i, d := range data {
		row := make([]string, len(header))
		for j, h := range header {
			if reflect.TypeOf(d[h]).Kind() == reflect.Slice {
				// If the value is a slice, join it with commas
				row[j] = strings.Join(d[h].([]string), ", ")
			} else {
				// Otherwise, just use the string representation of the value
				row[j] = fmt.Sprintf("%v", d[h])
			}
		}
		rows[i+1] = row
	}

	// Write the data to the CSV file
	file, err := os.Create(*csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range rows {
		err := writer.Write(row)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
