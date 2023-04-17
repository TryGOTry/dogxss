package payload

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

var Db *gorm.DB

func init() {

	Db = GetDb()
	//SqlInit()
}
func GetDb() *gorm.DB {
	var err error
	//pwd, _ := os.Getwd()
	//s := strings.Replace(pwd, "\\", "/", -1)
	//fmt.Println(s+"/db/team.db")

	Db, err := gorm.Open("sqlite3", "./dogxss.db")
	//defer Db.Close()
	if err != nil {
		log.Println("[Info] 数据库连接失败！", err)

	}
	if err != nil {
		//panic(err)
		print(err)
	}
	//Db.LogMode(true) //sql调试模式
	return Db
}
func AuthProJect(id string) bool { //验证项目id
	var h Project
	err := Db.Table("projects").Where("project_id = ? ", id).First(&h).Error
	if err != nil {
		fmt.Println("sql:", err)
		return false
	}
	return true
}
func AddInfo(i ProjectReturn) {
	Db.Table("project_returns").Create(&i)
}
func CountProJect() int { //全部项目
	count := 0
	Db.Table("projects").Count(&count)
	return count
}
func CountReturn() int { //全部文件数量
	count := 0
	Db.Table("project_returns").Count(&count)
	return count
}
func CheckProJect(projectname string) bool {
	var h Project
	projectname = strings.Replace(projectname,"/","",1)
	//fmt.Println(projectname)
	err := Db.Table("projects").Where("project_name = ? ", projectname).First(&h).Error
	if err != nil {
		fmt.Println("sql:", err)
		return false
	}
	return true
}
type Project struct {
	//gorm.Model     `json:"gorm_model"`
	Id int `json:"id",gorm:"primary_key",gorm:"default:1"`
	ProjectName    string `json:"project_name" gorm:"primary_key"`
	ProjectInfo    string `json:"project_info"`
	ProjectTime    string `json:"project_time"`
	ProjectId      string `json:"project_id,omitempty"`
	ProjectUrl     string `json:"project_url,omitempty"`
	ProjectPayload string `json:"project_payload,omitempty"`
}

type ProjectReturn struct {
	gorm.Model
	//ProjectName          string
	ProjectId            string
	ReturnUrl            string
	ReturnCookie         string
	ReturnReferer        string
	ReturnUserAgent      string
	ReturnTime           string
	ReturnOrigin         string
	ReturnDom            string //源码
	ReturnSessionStorage string //sessionstorage
	Ip                   string
	ReturnScreenShort    string //base64编码的
	ReturnLocalStorge    string
}
