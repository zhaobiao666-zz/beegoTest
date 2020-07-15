package main

import (
	"beegoTest/models"
	_ "beegoTest/routers"
	"github.com/astaxie/beego"
)

var StudentsInfoMap map[string]models.StudentInfo

func main() {
	beego.Run()
}

