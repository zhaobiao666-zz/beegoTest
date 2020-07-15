package controllers

import (
	"beegoTest/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)


var StudentsInfoMap = make(map[string]models.StudentInfo)
type MainController struct {
	beego.Controller
}

type StudentsController struct {
	beego.Controller
}

func (this *StudentsController) Get() {
	//c.Ctx.WriteString("zhaobiao")
	fmt.Println("Get require: ")
	id := this.GetString("id")
	if studentInfo, ok := StudentsInfoMap[id]; ok != false {
		this.Data["json"] = models.RespStudentInfo{
			StudentInfo: studentInfo,
			Code:        200,
			Desc:        "Success",
		}
	} else {
			this.Data["json"] = models.RespStudentInfo{
				Code:        200,
				Desc:        "Have no this student",
			}
		}
	this.ServeJSON()
	//this.Data["json"] = err.Error()
}


func (this *StudentsController) Post() {
	//c.Ctx.WriteString("zhaobiao")
	fmt.Println("Post require")
	var stuInfo models.StudentInfo
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &stuInfo); err == nil {
		StudentsInfoMap[stuInfo.Id] = stuInfo
		this.Data["json"] = models.RespStudentInfo{
			StudentInfo : stuInfo,
			Code:200,
			Desc: "Success",
		}
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()

}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
