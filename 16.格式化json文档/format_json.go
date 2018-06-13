package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"io"
	"encoding/json"
	"log"
	"strconv"
)
var Ac Account
var FilePath string

type Account struct {
	Car_ID string `json:"car_ID"`
	Colors [][3]string `json:"colors"`
	Wheels [][2]string `json:"wheels"`
	Explains []Exp  `json:"explains"`
}
type Exp struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Title string `json:"title"`
	Des string `json:"des"`
}

func main() {
	fmt.Printf("请输入需要规范化的car.json文件路径: ")
	fmt.Scanln(&FilePath)
	fmt.Println(FilePath)

	jsonFile := FilePath
	b, e := ioutil.ReadFile(jsonFile)
	if e != nil {
		fmt.Printf("Error: %s\n", e)
		return
	}
	str := string(b)
	str1 := strings.Replace(str, "“","\"",-1)
	str2 := strings.Replace(str1, "”","\"",-1)
	str3 := strings.Replace(str2, "'","\"",-1)
	str4 := strings.Replace(str3, "，",",",-1)
	str5 := strings.Replace(str4, "：",":",-1)
	str6 := strings.Replace(str5, "\t","",-1)
	str7 := strings.Replace(str6, "\n","",-1)
	f, err1 := os.OpenFile(jsonFile, os.O_WRONLY, 0644)
	if err1 != nil {
		fmt.Printf("Error: %s\n", err1)
		return
	}
	n, err1 := io.WriteString(f, str7)
	fmt.Println(n)

	err := json.Unmarshal([]byte(str7), &Ac)
	if err !=nil{
		log.Fatal(err)
	}
	for i,_ := range Ac.Colors{
		if i <= 9{
			Ac.Colors[i][0] = "Color_0" + strconv.Itoa(i)
		}else {
			Ac.Colors[i][0] = "Color_" +strconv.Itoa(i)
		}
	}

	for i,_ := range Ac.Wheels{
		if i+1 <= 9{
			Ac.Wheels[i][0] = "Wheel_0" + strconv.Itoa(i+1)
		}else {
			Ac.Wheels[i][0] = "Wheel_" +strconv.Itoa(i+1)
		}
	}
	deal_file,_ := json.Marshal(Ac)
	res := string(deal_file)
	fmt.Println(res)
	WriteWithFileWrite(jsonFile, res)
}

func WriteWithFileWrite(name,content string){
	fileObj,err := os.OpenFile(name,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _,err := fileObj.WriteString(content);err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile " +
			"and *File.WriteString method.",content)
	}
}

