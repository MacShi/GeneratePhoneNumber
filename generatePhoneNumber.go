package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)
import "math/rand"

func generatePhoneNumber(first string) (string) {
	second := fmt.Sprintf("%.8d ", rand.Int31()%10000000)
	phone := first+second
	phone = strings.Replace(phone," ","",-1)
	return phone
}

func wirterPhoneNumber(phone string)  {
	file,err := os.OpenFile(fmt.Sprintf("./%s.txt",phone[:3]),os.O_WRONLY| os.O_CREATE|os.O_APPEND,0666)
	if err==nil{
		file.WriteString(phone+"\n")
		file.Close()
	}
}
func batchGenerate(first string,count int64)  {
	var i int64
	for i =0; i<count;i++{
		phone := generatePhoneNumber(first)
		wirterPhoneNumber(phone)
		//rate :=fmt.Sprintf("%.2f", (float64(i+1)/10000000))
		info := fmt.Sprintf("[Info] %s	进度%.2f%%	已完成%d", time.Now().Format("2006-01-02 15:04:05"), (float64((i+1)/count) *100 ),(i+1))
		fmt.Println(info)
	}
}

func main()  {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(generatePhoneNumber("134"))
	var first string
	var count int64
	flag.StringVar(&first,"f","","Top three mobile phone numbers, for example 135 136")
	flag.Int64Var(&count,"c",10000000,"The count of generate phone number, default 10000000")
	//flag.i
	flag.Parse()
	if len(first)==0{
		fmt.Println("please use -h")
		os.Exit(1)
	}
	batchGenerate(first,count)
}