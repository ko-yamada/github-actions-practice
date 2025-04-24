package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	// ランダムな秒数でsleepする
	println("Build started")

	// コマンドライン引数からsleep時間を取得
	sleepTime, _ := strconv.Atoi(os.Args[2])
	time.Sleep(time.Duration(sleepTime) * time.Second)
	println(os.Args[1], " successfully")
}
