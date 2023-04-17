package dog

import (
	"dogxss/payload"
	"github.com/gin-gonic/gin"
	"os"
)

func GetPayLoad(c *gin.Context)  {
	expname := c.Param("exp")
	if payload.CheckFileExist("./exp/"+expname){
		if payload.CheckProJect(expname) {
			exp,err:=os.ReadFile("./exp/"+expname)
			if err != nil {
				c.String(404,"404 page not found")
				return
			}else {
				c.String(200, string(exp))
				return
			}
		}

	}
	c.String(404,"404 page not found")
	return
}
