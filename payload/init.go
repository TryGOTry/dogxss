package payload

import (
	"fmt"
	"os"
)
func CheckFileExist(fileName string) bool { //检查配置文件是否存在
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func SqlInit() {
	//Db.Create("dogxss.db")
	err := Db.CreateTable(Project{}).Error
	if err != nil {
		fmt.Println("[Cat]  数据库初始化失败!错误信息:", err)
		os.Exit(0)
	}
	Db.CreateTable(ProjectReturn{})
}
