package main

import (
	"dogxss/config"
	"dogxss/dog"
	"dogxss/pages"
	"github.com/unrolled/secure"
	"os"
	"os/signal"

	"flag"
	"github.com/GoAdminGroup/components/login"
	_ "github.com/GoAdminGroup/components/login/theme2"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	//"github.com/unrolled/secure"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                       // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"dogxss/models"
	"dogxss/tables"
)

var (
	Host    string
	SSlPort string
	WebPort string
	SSLKey  string //私钥
	SSLPem  string //pem
)

func init() {
	flag.StringVar(&Host, "d", "", "domain")
	flag.StringVar(&SSlPort, "ssl", "443", "https端口") //ssl端口
	flag.StringVar(&WebPort, "web", "80", "web端口")  //web端口
	flag.StringVar(&SSLKey, "key", "", "ssl私钥") //ssl私钥
	flag.StringVar(&SSLPem, "pem", "", "ssl公钥") //ssl私钥
	flag.Parse()
	if Host == "" {
		flag.Usage()
		os.Exit(0)
	}
}
func main() {
	config.SetPayloadUrl(Host)
	startServer()
}
func Cors() gin.HandlerFunc { //跨域问题
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		//c.Header("Waf", "DogDog")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		if c.Request.URL.Port() == SSlPort {
			secureMiddleware := secure.New(secure.Options{
				SSLRedirect: true,
				SSLHost:     Host + ":" + SSlPort,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)
			if err != nil {
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	r.Use(Cors())

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// 使用登录页面组件
	login.Init(login.Config{
		Theme: "theme2",
	})

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		//panic(err)
	}

	r.GET("/"+config.Payload+"/*exp", dog.GetPayLoad) //返回payload
	r.POST("/callback", dog.ReturnPayLoad)            //返回的信息

	eng.HTML("GET", "/admin/dogxss", pages.GetDashBoard)
	models.Init(eng.SqliteConnection())
	go r.Run(":" + WebPort)
	_ = r.RunTLS(":"+SSlPort, SSLPem, SSLKey)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.SqliteConnection().Close()
}
