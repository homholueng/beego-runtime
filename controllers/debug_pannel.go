package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DebugPannelController struct {
	beego.Controller
}

func (c *DebugPannelController) Get() {
	c.Data["BK_STATIC_URL"] = "/static"
	c.Data["SITE_URL"] = "/bk_plugin"
	c.TplName = "debug_panel.tpl"
}
