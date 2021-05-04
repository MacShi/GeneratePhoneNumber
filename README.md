## 介绍
指定手机号前三位和电话号码归属地，随机生成手机号，并写入文本中
## 使用方法
```
Windows.exe -h
Usage of gen_phone_windos.exe:
  -C string
        set the city of phone number
  -c int
        The count of generate phone number, default 10000000 (default 10000000)
  -f string
        Top three mobile phone numbers, for example 135 136


For Example:
//在当前目录下生成135.txt文件，共有1000行135开头的手机号
Windows.exe -f 135 -c 1000
```

