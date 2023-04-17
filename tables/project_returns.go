package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"strings"
)

func GetProjectReturnsTable(ctx *context.Context) table.Table {

	projectReturns := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("sqlite", "dogxss"))

	info := projectReturns.GetInfo().HideFilterArea()
	info.SetDefaultPageSize(10) //只查询前30
	info.WhereRaw("deleted_at IS NULL")
	info.HideNewButton()
	info.AddButton("返回列表", icon.Home, action.Jump("/admin/info/projects"))

	info.HideNewButton()

	info.HideEditButton()

	info.HideRowSelector()

	info.AddField("Id", "id", db.Integer).
		FieldSortable().
		FieldHide()
	info.AddField("Created_at", "created_at", db.Datetime).
		FieldHide()
	info.AddField("Updated_at", "updated_at", db.Datetime).
		FieldHide()
	info.AddField("Deleted_at", "deleted_at", db.Datetime).
		FieldHide()
	//info.AddField("项目名", "project_name", db.Varchar)
	info.AddField("项目id", "project_id", db.Varchar).
		FieldHide()
	info.AddField("来自", "return_url", db.Varchar).
		FieldFilterable().FieldWidth(300)
	info.AddField("Cookie", "return_cookie", db.Varchar).
		FieldHide()
	info.AddField("Return_referer", "return_referer", db.Varchar).
		FieldHide()
	info.AddField("Return_user_agent", "return_user_agent", db.Varchar).
		FieldHide()
	info.AddField("时间", "return_time", db.Varchar).
		FieldFilterable().
		FieldSortable().FieldWidth(300)
	info.AddField("Return_origin", "return_origin", db.Varchar).
		FieldHide()
	info.AddField("Return_dom", "return_dom", db.Varchar).
		FieldHide()
	info.AddField("Return_session_storage", "return_session_storage", db.Varchar).
		FieldHide()
	info.AddField("Return_screen_short", "return_screen_short", db.Varchar).
		FieldHide()
	info.AddField("Return_local_storge", "return_local_storge", db.Varchar).
		FieldHide()
	info.ExportValue()
	info.SetTable("project_returns").SetTitle("返回列表").SetDescription("cookie..")

	formList := projectReturns.GetForm()
	formList.AddField("Id", "id", db.Integer, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenInsert()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenUpdate()
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("项目名称", "project_name", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("项目id", "project_id", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("来自", "return_url", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Cookie", "return_cookie", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Referer", "return_referer", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("User_Agent", "return_user_agent", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("请求时间", "return_time", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Origin", "return_origin", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("页面源码", "return_dom", db.Varchar, form.Code).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("SessionStorage", "return_session_storage", db.Varchar, form.Code).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("ScreenShort", "return_screen_short", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("localStorge", "return_local_storge", db.Varchar, form.Code).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.HideContinueEditCheckBox()
	formList.HideContinueNewCheckBox()
	formList.HideResetButton()

	formList.SetTable("project_returns").SetTitle("返回详细").SetDescription("列表")
	detail := projectReturns.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField("Ip", "ip", db.Varchar).FieldXssFilter()
	detail.AddField("请求地址", "return_url", db.Varchar).FieldXssFilter()
	detail.AddField("UserAgent", "return_user_agent", db.Varchar).FieldXssFilter()
	detail.AddField("请求时间", "return_time", db.Varchar)
	detail.AddField("Cookie", "return_cookie", db.Varchar).FieldXssFilter().FieldWidth(200).FieldCopyable("Copy")
	detail.AddField("Referer", "return_referer", db.Varchar).FieldWidth(300).FieldXssFilter()
	detail.AddField("Origin", "return_origin", db.Varchar).FieldXssFilter()
	detail.AddField("截图", "return_screen_short", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "\n", "<br>")
		//str:=strings.Replace(replacer.Replace(model.Value),"&lt;br&gt;","<br>",-1)
		var sc string
		if model.Value!="" {
			sc = "<img src=\""+replacer.Replace(model.Value)+"\" width=\"100%\" height=\"100%\">"
		}else {
			sc = "无"
		}
		return sc
	})
	detail.AddField("SessionStorage", "return_session_storage", db.Varchar).FieldXssFilter().FieldWidth(200).FieldCopyable("Copy")
	detail.AddField("LocalStorage", "return_local_storage", db.Varchar).FieldXssFilter().FieldWidth(200).FieldCopyable("Copy")
	detail.AddField("项目标识", "project_id", db.Varchar).FieldXssFilter()
	detail.AddField("页面源码", "return_dom", db.Varchar).FieldXssFilter().FieldWidth(200).FieldCopyable("Copy")
	//detail.AddField("url", "project_url", db.Varchar).FieldHide()
	//detail.AddField("name", "project_file_name", db.Varchar).FieldXssFilter().FieldHide()
	return projectReturns
}
