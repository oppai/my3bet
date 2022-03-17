package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Counter struct {
	openCount     int64
	threeBetCount int64
	threeBetRatio float32
}

func (c *Counter) IncrementOpen() {
	c.openCount += 1
}

func (c *Counter) IncrementThreebet() {
	c.threeBetCount += 1
	c.IncrementOpen()
}

func (c *Counter) Ratio3bet() {
	if c.openCount == 0 {
		c.threeBetRatio = 0.0
	} else {
		c.threeBetRatio = float32(c.threeBetCount) / float32(c.openCount)
	}
}

type CounterHandler struct {
	Counter                *Counter
	OpenLabelBind          binding.String
	ThreeBetLabelBind      binding.String
	ThreeBetRatioLabelBind binding.String
}

func (c *CounterHandler) Init() {
	c.update()
}

func (c *CounterHandler) IncrementOpen() {
	c.Counter.IncrementOpen()
	c.Counter.Ratio3bet()
	c.update()
}

func (c *CounterHandler) IncrementThreebet() {
	c.Counter.IncrementThreebet()
	c.Counter.Ratio3bet()
	c.update()
}

func (c *CounterHandler) update() {
	c.OpenLabelBind.Set(fmt.Sprintf("%d", c.Counter.openCount))
	c.ThreeBetLabelBind.Set(fmt.Sprintf("%d", c.Counter.threeBetCount))
	c.ThreeBetRatioLabelBind.Set(fmt.Sprintf("%.1f%%", 100.0*c.Counter.threeBetRatio))
}

func main() {
	a := app.New()
	w := a.NewWindow("3bet checker")
	w.SetFixedSize(true)

	c := &CounterHandler{
		Counter:                &Counter{},
		OpenLabelBind:          binding.NewString(),
		ThreeBetLabelBind:      binding.NewString(),
		ThreeBetRatioLabelBind: binding.NewString(),
	}
	c.Init()

	openLabel := widget.NewLabelWithData(c.OpenLabelBind)
	openButton := widget.NewButton("Open", func() {
		c.IncrementOpen()
	})

	threeBetLabel := widget.NewLabelWithData(c.ThreeBetLabelBind)
	threeBetButton := widget.NewButton("3bet+", func() {
		c.IncrementThreebet()
	})

	threeBetRatioLabel := widget.NewLabelWithData(c.ThreeBetRatioLabelBind)

	rowItem1 := container.New(layout.NewVBoxLayout(), openLabel, openButton)
	rowItem2 := container.New(layout.NewVBoxLayout(), threeBetLabel, threeBetButton)
	controller := container.New(layout.NewHBoxLayout(), rowItem1, rowItem2)

	content := container.New(layout.NewHBoxLayout(), controller, threeBetRatioLabel)

	w.SetContent(container.New(layout.NewMaxLayout(), content))
	w.ShowAndRun()
}
