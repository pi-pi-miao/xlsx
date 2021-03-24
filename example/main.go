package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main(){
	file := xlsx.NewFile()

	file.SetProperties().
		SetApplicationName("demo").
		SetTitle("message").
		SetSubject("message list").
		SetKeywords("key").
		SetCategory("test").
		SetDescription("just test message").
		SetCreator("pi-pi-miao").
		SetLastModifiedBy("pi-pi-miao")

	sheet, err := file.AddSheet("person list")
	if err != nil {
		fmt.Printf("add sheet err:%v",err)
		return
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	//cell.SetString("hello ")
	cell.Value = "name"
	cell = row.AddCell()
	cell.Value = "gender"
	cell = row.AddCell()
	cell.Value = "message"

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "alice"
	cell = row.AddCell()
	cell.Value = "male"
	cell = row.AddCell()
	cell.Value = "github.com/pi-pi-miao"

	///home/pi-pi/project/template/example2/github.com/tealeg/xlsx/example/
	err = file.Save("pi-pi-miao.xlsx")
	if err != nil {
		fmt.Printf("save err :%v",err)
		return
	}
	fmt.Println(file.Properties)
}
