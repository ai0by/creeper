package runner

import (
	//加载api文档使用
	_ "creeper/creeper_http_docs"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

//路由
type creeperApiEngineRouter struct {
	routerPath  string
	method      []string
	handlerFunc gin.HandlerFunc
}

var creeperApiEngine *gin.Engine

//存储路由
var creeperApiEngineRouters []*creeperApiEngineRouter

func CreeperApiRunner() {
	//启动设置端口
	cfg, err := goconfig.LoadConfigFile("etc/creeper.ini")
	if err != nil {
		panic(err)
	}
	mode, err := cfg.GetValue("web", "mode")
	if err != nil {
		panic(err)
	}
	gin.SetMode(mode)
	creeperApiEngine = gin.New()
	//允许使用跨域请求,全局中间件
	creeperApiEngine.Use(cors())
	httpPort, err := cfg.GetValue("web", "http_port")
	if err != nil {
		panic(err)
	}
	//路由加载
	loadCreeperApiEngineRouter()
	if mode == "debug" {
		//swagger
		url := ginSwagger.URL(fmt.Sprintf("http://127.0.0.1:%s/swagger/doc.json", httpPort)) // The url pointing to API definition
		creeperApiEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
	//启动
	err = creeperApiEngine.Run(fmt.Sprintf(":%s", httpPort))
	if err != nil {
		panic(err)
	}
}

//给控制器注册路由使用
func RegisterCreeperApiRunner(routerPath string, method []string, handlerFunc gin.HandlerFunc) {
	creeperApiEngineRouters = append(creeperApiEngineRouters, &creeperApiEngineRouter{
		routerPath:  routerPath,
		method:      method,
		handlerFunc: handlerFunc})
	logrus.Info("路由长度：", len(creeperApiEngineRouters))
}

//加载已经注册的路由
func loadCreeperApiEngineRouter() {
	for _, router := range creeperApiEngineRouters {
		//method空就是所有
		if len(router.method) == 0 {
			creeperApiEngine.Any(router.routerPath, router.handlerFunc)
		} else {
			for _, m := range router.method {
				creeperApiEngine.Handle(m, router.routerPath, router.handlerFunc)
			}
		}
	}
}
