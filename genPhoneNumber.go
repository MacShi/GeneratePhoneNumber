package main

import (
	"bufio"
	"fmt"
	"github.com/xluohome/phonedata"
	"io"
	"os"
	"strconv"
	"time"
)

func wirterPhoneNumber1(phone string) {
	file, err := os.OpenFile("./phone.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		file.WriteString(phone + "\n")
		file.Close()
	}
}


func readHlr()  {
	file, err := os.OpenFile("ZZ-HLR.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		//return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		genPhoneNumber(string(a))
	}

}


func genPhoneNumber(first string)  {
	for g:=0;g<=9;g++{
		for s:=0;s<=9;s++{
			for b:=0;b<=9;b++{
				for q:=0;q<=9;q++{
					num := strconv.Itoa(g)+strconv.Itoa(s)+strconv.Itoa(s)+strconv.Itoa(q)
					phone :=first+num
					_,err:=phonedata.Find(phone)
					if err==nil{
						wirterPhoneNumber1(phone)
					}
				}
			}
		}
	}
}

func main()  {
	beginTime := time.Now()
	readHlr()
	endTime := time.Now()
	result := fmt.Sprintf("[End] 耗时 %2f 秒", endTime.Sub(beginTime).Seconds())
	fmt.Println(result)
}