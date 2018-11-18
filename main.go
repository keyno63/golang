package main

import (
	"./httpServer"
	"./sqlquery"
	//"github.com/keyno63/golang/sqlquery"
	"fmt"
)

func main() {
	fmt.Println("Go言語はじめました！")
	fmt.Println("SQL query, sample.")
	sqlquery.Query()
	fmt.Println("http server, sample.")
	httpServer.Server()

}

// 52:54:00:b9:72:d0
