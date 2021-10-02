package commands

import (
	"fmt"
	"github.com/Summer2021-Project/api-skeleton/di"
	"runtime"
	"strings"
)

const logo = `
		Dubbo-go
`

func welcome() {
	fmt.Println(strings.Replace(logo, "*", "`", -1))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("Server      Name:      %s", "Dubbo-go-api"))
	fmt.Println(fmt.Sprintf("System      Name:      %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:   %s", runtime.Version()[2:]))
	fmt.Println(fmt.Sprintf("Listen      Addr:      %s", di.Server().Addr))
}
