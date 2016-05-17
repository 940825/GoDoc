package main

import(
    "fmt"
    "time"
)

func get_time()(string) {
    timenow := time.Now().Unix()
    tm := time.Unix(timenow, 0)
    timestamp := tm.Format("2006-01-02 15:04:05")
    return timestamp
}

func main() {
    time := get_time()
    fmt.Println(time)
}

/*
看了上面的代码，你可能会好奇，为什么格式字符串的时候，用的是2006-01-02这种格式。其实在Go语言里，这些数字都是有特殊函义的，不是随便指定的数字，见下面列表：

月份 1,01,Jan,January

日　 2,02,_2

时　 3,03,15,PM,pm,AM,am

分　 4,04

秒　 5,05

年　 06,2006

周几 Mon,Monday

时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST

时区字母缩写 MST

*/
