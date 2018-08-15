package main

import (
	_ "xiangmu/ihome/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"


	"strings"
	"net/http"
)

func main() {
	beego.SetStaticPath("/static","static")
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath() {
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	///index.html --->/static/html/index.html
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}



