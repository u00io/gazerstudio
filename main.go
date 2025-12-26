package main

import (
	"github.com/u00io/gazerstudio/forms/mainform"
	"github.com/u00io/gazerstudio/localstorage"
	"github.com/u00io/gazerstudio/system"
)

func main() {
	localstorage.Init("gazer_studio")

	system.Instance = system.NewSystem()
	system.Instance.Start()
	system.Instance.InitDefaultProject()

	mainform.Run()
}
