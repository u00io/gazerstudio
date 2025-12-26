package project

import (
	"encoding/json"
	"os"

	"github.com/u00io/gazerstudio/utils"
)

type Project struct {
	Id   string
	Name string
}

const (
	ProjectFileName  = "project.json"
	DataItemFileName = "data_item.json"
)

func NewProject() *Project {
	var c Project
	c.Id = utils.GenerateId()
	c.Name = "Noname"
	return &c
}

func (c *Project) Load(id string) error {
	c.Id = id
	path := utils.ProjectDirPath(id)
	bs, err := os.ReadFile(path + "/" + ProjectFileName)
	if err != nil {
		c.Save()
		return err
	}
	err = json.Unmarshal(bs, c)
	if err != nil {
		c.Save()
		return err
	}
	return err
}

func (c *Project) Save() {
	bs, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return
	}

	dir := utils.ProjectDirPath(c.Id)

	_ = os.MkdirAll(dir, 0755)
	_ = os.WriteFile(dir+"/"+ProjectFileName, bs, 0644)
}

func (c *Project) DataItems() []*DataItem {
	var items []*DataItem
	dir := utils.ProjectDataItemsDirPath(c.Id)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return items
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dataItem := NewDataItem(c.Id)
			err := dataItem.Load(entry.Name())
			if err == nil {
				items = append(items, dataItem)
			}
		}
	}
	return items
}
