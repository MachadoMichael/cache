package main

import (
	"github.com/MachadoMichael/cache/infra"
	"github.com/MachadoMichael/cache/infra/database"
	"github.com/MachadoMichael/cache/route"
)

func main() {
	infra.Init()
	database.Init()
	route.Init()
}
