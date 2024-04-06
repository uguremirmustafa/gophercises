package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/uguremirmustafa/task/cmd"
	"github.com/uguremirmustafa/task/db"
)

func main() {

	homePath, _ := os.UserHomeDir()
	dbPath := filepath.Join(homePath, ".bolt.tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
