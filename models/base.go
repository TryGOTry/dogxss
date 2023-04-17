package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open("sqlite", c.GetDB("default"))
	orm.Debug()
	if err != nil {
		panic("initialize orm failed")
	}
}
