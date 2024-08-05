package main

import (
	"github.com/MachadoMichael/cache/infra"
	"github.com/MachadoMichael/cache/infra/database"
)

func main() {
	infra.Init()
	database.Init()
}
