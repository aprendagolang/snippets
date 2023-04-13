package main

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (u *UserController) Get() {
	u.Ctx.WriteString("hello world")
}

func main() {
	web.Router("/", &UserController{})
	web.Run()
}
