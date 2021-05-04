package main

import (
	"flag"
	"fmt"
	"github.com/xluohome/phonedata"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generatePhoneNumber(first string) string {
	second := fmt.Sprintf("%.8d ", rand.Int31()%10000000)
	phone := first + second
	phone = strings.Replace(phone, " ", "", -1)
	return phone
}

func wirterPhoneNumber(phone string) {
	file, err := os.OpenFile(fmt.Sprintf("./%s.txt", phone[:3]), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		file.WriteString(phone + "\n")
		file.Close()
	}
}

func checkPhone(phone string) bool {
	ok, err := phonedata.Find(phone)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if ok.City == "北京" {
		return true
	} else {
		return false
	}
}

func batchGenerate(first string, count int64) {
	var i int64
	for i = 0; i < count; i++ {
		phone := generatePhoneNumber(first)
		if checkPhone(phone) {
			wirterPhoneNumber(phone)
		}
		info := fmt.Sprintf("[Info] %s	进度%.2f%%	已完成第 %d 个", time.Now().Format("2006-01-02 15:04:05"), (float64(i+1) / float64(count) * 100), (i + 1))
		fmt.Println(info)
	}
}

func main() {
	var first string
	var count int64
	flag.StringVar(&first, "f", "", "Top three mobile phone numbers, for example 135 136")
	flag.Int64Var(&count, "c", 10000000, "The count of generate phone number, default 10000000")
	flag.Parse()
	if len(first) == 0 {
		info := "Usage of gen_phone_windos.exe:\n  -c int\n        The count of generate phone number, default 10000000 (default 10000000)\n  -f string\n        Top three mobile phone numbers, for example 135 136"
		fmt.Println(info)
		os.Exit(1)
	}
	beginTime := time.Now()
	batchGenerate(first, count)
	endTime := time.Now()
	result := fmt.Sprintf("[End] 共生成 %d 条，耗时 %2f 秒", count, endTime.Sub(beginTime).Seconds())
	fmt.Println(result)
}
