package di

import (
	"github.com/Summer2021-Project/di"
	"github.com/Summer2021-Project/dotenv"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func init() {
	obj := di.Object{
		Name: "xorm",
		New: func() (i interface{}, e error) {
			return xorm.NewEngine("mysql", dotenv.Getenv("DATABASE_DSN").String())
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}

func Xorm() (db *xorm.Engine) {
	if err := di.Populate("xorm", &db); err != nil {
		panic(err)
	}
	return
}
