package routers

import (
	"Beescan/controller"
	"Beescan/core/log"
	"Beescan/middleware"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

/*
创建人员：云深不知处
创建时间：2022/1/3
程序功能：路由
*/

//go:embed static templates
var content embed.FS

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(log.LoggerToFile())
	r.Use(middleware.AuthMiddleWare())

	// 告诉gin框架模板文件引用的静态文件去哪里找
	//r.Static("/public/static", "../routers/static")
	r.StaticFS("/public", http.FS(content))

	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("../routers/templates/*")
	t := template.Must(template.New("").ParseFS(content, "templates/*"))
	r.SetHTMLTemplate(t)

	r.NoRoute(controller.Error)
	r.GET("/", controller.LoginGet)

	// 初始登录
	r.GET("/login", controller.LoginGet)
	r.POST("/login", controller.LoginPost)

	// 首页
	r.GET("/info", controller.InfoGet)
	r.POST("/info", controller.InfoGet)

	// 资产展示
	r.GET("/assets", controller.AssetsGet)
	r.POST("/assets", controller.AssetsPost)

	// 资产导出
	r.GET("/csv", controller.AssetsExport)

	// 资产探测
	r.GET("/scan", controller.ScanGet)
	r.POST("/scan", controller.ScanPost)
	r.GET("/task", controller.TaskDelete)

	// 资产详细页面
	r.GET("/ipdetail", controller.SingleAssetsDetail)

	// 漏洞检测
	r.GET("/vul", controller.VulGet)
	r.POST("/vul", controller.VulPost)

	// 漏洞详情页面
	r.GET("/vuldetail", controller.VulDetail)

	// 日志管理
	r.GET("/logs", controller.LogsGet)
	r.GET("/nodelog", controller.NodeLog)

	return r
}
