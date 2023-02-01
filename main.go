package main

import (
	"redisKeyspace/db"
)

func main() {
	db.CheckForEvents()
}