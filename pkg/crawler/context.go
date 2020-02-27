package crawler

import (
	"context"
	"goutils/model"
)

type Context struct{
	context.Context
	Task model.TaskInfo
	DataList []model.DataInfo
}

func (self *Context) AddJsData(jq map[string]interface{}) *Context{
	var data model.DataInfo
	u, ok := jq["Url"].(string)
	if !ok {
		return self
	}
	data.Url = u
	data.Date = jq["Date"].(string)
	data.Title = jq["Title"].(string)
	self.DataList = append(self.DataList, data)
	return self
}

