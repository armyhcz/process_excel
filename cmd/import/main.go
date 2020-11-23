package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"modules/process_excel/pkg/model"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	file := flag.String("file", "test.json", "json file name")

	jsonFile, err := os.Open(*file)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var student []model.Student
	err = json.Unmarshal([]byte(byteValue), &student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)
	fmt.Println(student[0].Name)

	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	for k, v := range student {
		f.SetSheetRow("Sheet2", fmt.Sprint("A", k+1), &[]string{strconv.Itoa(v.Id), v.Name, strconv.Itoa(v.Age)})
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	if err := f.SaveAs("student.xlsx"); err != nil {
		println(err.Error())
	}
}
