package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	log.SetFlags(12)
	f, err := excelize.OpenFile("12.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	rows := f.GetRows("Sheet1")
	var result []map[string]interface{}
	for _, row := range rows {
		id, errId := strconv.Atoi(row[0])
		age, errAge := strconv.Atoi(row[2])
		if errId != nil || errAge != nil {
			log.Println(errId, errAge)
			continue
		}

		student := map[string]interface{}{
			"id":   id,
			"name": row[1],
			"age":  age,
		}
		result = append(result, student)
	}
	students, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	// file write
	err = ioutil.WriteFile("test.json", students, 0644)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf(string(students))
}
