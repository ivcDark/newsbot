package main

import (
	"github.com/ivcDark/newsbot/cmd"
	_ "github.com/mattn/go-sqlite3" // <-- обязательно, чтобы зарегистрировать драйвер
)

func main() {
	cmd.Execute()
}
