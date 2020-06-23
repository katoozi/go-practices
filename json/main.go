package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Person is simple structure
type Person struct {
	Firstname, Lastname, Tel string
	Age                      int
}

func loadFromJSON(filename string, v interface{}) error {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	decodeJSON := json.NewDecoder(file)
	err = decodeJSON.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func saveToJSON(file *os.File, v interface{}) error {
	encodeJSON := json.NewEncoder(file)
	err := encodeJSON.Encode(v)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var persons []Person
	err := loadFromJSON("data.json", &persons)
	if err != nil {
		log.Fatalf("Error while loading json file %v", err)
	}
	fmt.Println(persons)

	data := []Person{
		{
			Firstname: "test",
			Lastname:  "test",
			Age:       100,
			Tel:       "+1234678765",
		},
		{
			Firstname: "test2",
			Lastname:  "test2",
			Age:       101,
			Tel:       "+1234678765",
		},
	}
	if err != nil {
		log.Fatalf("Error while openning new file: %v", err)
	}
	err = saveToJSON(os.Stdout, &data)
	if err != nil {
		log.Fatalf("error while write new file: %v", err)
	}

	fileData, err := ioutil.ReadFile("data_map.json")
	if err != nil {
		log.Fatalf("Error while read map file: %v", err)
	}
	var personMap map[string]interface{}
	json.Unmarshal([]byte(fileData), &personMap)
	fmt.Println(personMap["persons"])
	fmt.Println(personMap["other_persons"])
}
