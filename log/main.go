package main

import (
	"github.com/aprendagolang/log/logs"
)

func main() {
	// escreve no log ola.log
	logs.Write("olá mundo!", "./ola.log")
	// escreve no log /tmp/my-app.log
	logs.Write("funciona!!!", "")
}
