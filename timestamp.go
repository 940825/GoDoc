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
