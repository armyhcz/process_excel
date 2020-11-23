package main

import (
	"flag"
	"fmt"
	"log"
	"modules/process_excel/internal/database"
	"modules/process_excel/pkg/model"
)

type Config struct {
	env string
}

var cfg Config

func init() {
	flag.StringVar(&cfg.env, "env", "test", "")
}

func main() {
	log.SetFlags(12)
	flag.Parse()
	if cfg.env == "" {
		log.Fatal("empty env")
	}
	log.Println(cfg.env)
	if err := database.Init(cfg.env); err != nil {
		log.Fatal(err)
	}

	db := database.Get()
	var result []model.Student
	rows, err := db.Query("select `id`, `name`, `age` from student;")
	log.Println(err)
	if err == nil {
		for rows.Next() {
			var student model.Student
			err := rows.Scan(&student.Id, &student.Name, &student.Age)
			if err != nil {
				continue
			}

			result = append(result, student)
		}
		database.CloseRow(rows)
	}

	fmt.Println(result)
}
