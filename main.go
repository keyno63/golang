package main

import (
	"./sqlquery"
	//"github.com/keyno63/golang/sqlquery"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Go言語はじめました！")
	sqlquery.Query()
}

// 52:54:00:b9:72:d0
