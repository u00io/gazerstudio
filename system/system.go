package system

import (
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/u00io/gazerstudio/project"
)

type System struct {
	mtx    sync.Mutex
	events []Event

	project *project.Project
}

type Event struct {
	Name      string
	Parameter string
}

var Instance *System

func NewSystem() *System {
	var c System
	return &c
}

func (c *System) Start() {
	go c.thWork()
}

func (c *System) Stop() {
}

func (c *System) thWork() {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func (c *System) EmitEvent(event string, parameter string) {
	c.mtx.Lock()
	c.events = append(c.events, Event{Name: event, Parameter: parameter})
	c.mtx.Unlock()
}

func (c *System) GetAndClearEvents() []Event {
	c.mtx.Lock()
	events := c.events
	c.events = make([]Event, 0)
	c.mtx.Unlock()
	return events
}

func (c *System) GetProject() *project.Project {
	return c.project
}

func (c *System) OpenProject(id string) {
	c.project = project.NewProject()
	c.project.Id = id
	c.project.Name = "Project " + id

	c.project.Load(id)
}

func (c *System) InitDefaultProject() {

	c.project = project.NewProject()
	c.project.Id = "cdda2e3e6400d1f0abae1f0aeb07d080"
	c.project.Name = "Default Project"
	c.project.Save()

	return

	projectId := c.project.Id

	dataItem := project.NewDataItem(projectId)
	dataItem.Name = "Default Data Item"
	dataItem.Save()

	values := make([]project.DataItemValue, 0)
	for i := 0; i < 4096; i++ {
		dt := time.Now().Add(time.Duration(i) * time.Second).UnixMilli()

		nowSec := time.Now().Add(-1 * time.Hour).Add(time.Duration(i) * time.Second).Unix()

		demoData := ""
		//rnd := rand.Int31() % 100
		sinValue := math.Sin(float64(nowSec%600)/600.0*2.0*math.Pi)*100 + 100
		// add slow sin wave
		//sinValue += math.Sin(float64(nowSec%300)/300.0*2.0*math.Pi)*1050 + 50
		// add fast sin wave
		//sinValue += math.Sin(float64(nowSec%10)/10.0*2.0*math.Pi)*20 + 20

		demoData = strconv.FormatFloat(sinValue, 'f', 1, 64)

		value := project.DataItemValue{
			DT:    dt,
			Value: demoData,
		}

		values = append(values, value)
	}

	dataItem.SetValues(values)
	c.project = project.NewProject()
	c.project.Load(projectId)
}
