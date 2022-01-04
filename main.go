package main

import (
	"github.com/DazSanchez/go-dynamodb-practice/db"
	"github.com/DazSanchez/go-dynamodb-practice/server"
)

func main() {
	db.Init()
	server.Init()
}
