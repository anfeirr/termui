// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	sbc := widgets.NewStackedBarChart()
	sbc.Title = "Student's Marks: X-Axis=Name, Y-Axis=Grade% (Math, English, Science, Computer Science)"
	sbc.Labels = []string{"Ken", "Rob", "Dennis", "Linus"}

	sbc.Data = make([][]int, 4)
	sbc.Data[0] = []int{90, 85, 90, 80}
	sbc.Data[1] = []int{70, 85, 75, 60}
	sbc.Data[2] = []int{75, 60, 80, 85}
	sbc.Data[3] = []int{100, 100, 100, 100}
	sbc.SetRect(5, 5, 100, 30)
	sbc.BarWidth = 5

	ui.Render(sbc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
