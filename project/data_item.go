package project

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/u00io/gazerstudio/utils"
)

type DataItemValue struct {
	DT    int64  `json:"dt"`
	Value string `json:"value"`
}

func (c *DataItemValue) DateTimeString() string {
	return utils.FormatDateTime(c.DT)
}

func (c *DataItemValue) FloatValue() float64 {
	fv, err := strconv.ParseFloat(c.Value, 64)
	if err != nil {
		return 0
	}
	return fv
}

type DataItemHistory struct {
	Values []DataItemValue `json:"values"`
}

type DataItem struct {
	projectId string

	Id   string
	Name string
}

func NewDataItem(projectId string) *DataItem {
	var c DataItem
	c.projectId = projectId
	c.Id = utils.GenerateId()
	c.Name = "New Data"
	return &c
}

func (c *DataItem) Load(id string) error {
	dir := utils.ProjectDataItemDirPath(c.projectId, id)
	bs, err := os.ReadFile(dir + "/" + DataItemFileName)
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

func (c *DataItem) Save() {
	dir := utils.ProjectDataItemDirPath(c.projectId, c.Id)

	bs, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return
	}

	_ = os.MkdirAll(dir, 0755)
	_ = os.WriteFile(dir+"/"+DataItemFileName, bs, 0644)
}

func (c *DataItem) SetValues(values []DataItemValue) {
	history := DataItemHistory{
		Values: values,
	}
	dir := utils.ProjectDataItemDirPath(c.projectId, c.Id)
	bs, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return
	}
	_ = os.MkdirAll(dir, 0755)
	_ = os.WriteFile(dir+"/history.json", bs, 0644)
}

func (c *DataItem) GetValues() []DataItemValue {
	var history DataItemHistory
	dir := utils.ProjectDataItemDirPath(c.projectId, c.Id)
	bs, err := os.ReadFile(dir + "/history.json")
	if err != nil {
		return history.Values
	}
	err = json.Unmarshal(bs, &history)
	if err != nil {
		return history.Values
	}
	return history.Values
}

func (c *DataItem) FFT() []float64 {
	values := c.GetValues()
	var data []float64
	for _, v := range values {
		fv, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			continue
		}
		data = append(data, fv)
	}

	// add zero padding to the next power of two
	n := len(data)
	power := 1
	for power < n {
		power <<= 1
	}
	for len(data) < power {
		data = append(data, 0)
	}

	fft := utils.FFTDouble(data)
	return fft
}
