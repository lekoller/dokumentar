package components

import "github.com/lekoller/dokumentar/core/templates"

func mountRenderCallback(c *ControlPanel) func() {
	return func() {
		var builderList []templates.BuilderListItem

		for _, item := range c.InputList.Items {
			if !item.Deleted {
				entity, err := item.Entity.Get()
				if err != nil {
					panic(err)
				}
				endpoint, err := item.Endpoint.Get()
				if err != nil {
					panic(err)
				}
				json, err := item.JsonEntry.Get()
				if err != nil {
					panic(err)
				}
				comment, err := item.CommentEntry.Get()
				if err != nil {
					panic(err)
				}
				builderList = append(builderList, templates.BuilderListItem{
					Entity:       entity,
					ConnType:     *item.ConnType,
					Method:       *item.Method,
					Endpoint:     endpoint,
					JsonEntry:    json,
					CommentEntry: comment,
				})
			}
		}

		pro, err := c.ProjectInfo.ProjectName.Get()
		if err != nil {
			panic(err)
		}
		con, err := c.ProjectInfo.ContainerName.Get()
		if err != nil {
			panic(err)
		}
		mod, err := c.ProjectInfo.ModuleName.Get()
		if err != nil {
			panic(err)
		}

		templates.BuildTemplate(templates.BuilderDTO{
			ProjectName:   pro,
			ContainerName: con,
			ModuleName:    mod,
			List:          builderList,
		})
	}
}
