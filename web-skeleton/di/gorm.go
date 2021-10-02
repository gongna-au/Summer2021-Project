package di

import (
	"github.com/Summer2021-Project/di"
	"github.com/Summer2021-Project/dotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	obj := di.Object{
		Name: "gorm",
		New: func() (i interface{}, e error) {
			return gorm.Open(mysql.Open(dotenv.Getenv("DATABASE_DSN").String()))
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}

func Gorm() (db *gorm.DB) {
	if err := di.Populate("gorm", &db); err != nil {
		panic(err)
	}
	return
}
