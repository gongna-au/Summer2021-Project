package argv

import (
	"os"
	"path/filepath"
	"regexp"
)

// 命令行信息
var (
	prog program
	cmd  string
)

// 命令行程序信息
type program struct {
	Path    string
	AbsPath string
	Dir     string
	File    string
}

// Program 返回命令行程序信息
func Program() *program {
	return &prog
}

// Command 返回当前命令信息
func Command() string {
	return cmd
}

// 创建命令行程序信息
func newProgram() program {
	abspath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	dir, file := filepath.Split(abspath)
	return program{
		Path:    os.Args[0],
		AbsPath: abspath,
		Dir:     dir[:len(dir)-1],
		File:    file,
	}
}

// 创建当前命令信息
func newCommand(singleton bool) string {
	if len(os.Args) <= 1 || singleton {
		return ""
	}
	cmd := ""
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9_\-:]+$`, os.Args[1]); ok {
		cmd = os.Args[1]
		if cmd[:1] == "-" {
			cmd = ""
		}
	}
	return cmd
}
