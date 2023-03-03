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
		fmt.Println("Error: Invalid input file")
		return
	}
	if *csvFile == "" {
		fmt.Println("Error: Invalid output file")
		return
	}

	jsonData, err := os.ReadFile(*jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	header := make([]string, 0)
	for k := range data[0] {
		header = append(header, k)
	}

	rows := make([][]string, len(data)+1)
	rows[0] = header
	for i, d := range data {
		row := make([]string, len(header))
		for j, h := range header {
			if reflect.TypeOf(d[h]).Kind() == reflect.Slice {
				row[j] = strings.Join(d[h].([]string), ", ")
			} else {
				row[j] = fmt.Sprintf("%v", d[h])
			}
		}
		rows[i+1] = row
	}

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
