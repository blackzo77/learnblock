package main // 의미 작성

import (
	"github.com/blackzo77/learnblock/cli"
	"github.com/blackzo77/learnblock/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
