package tool

import (
	"flag"
	"github.com/caothu159/grunt"
	"github.com/caothu159/hosts"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/themes/light"
)

var _defaultScaleFactor float32
var _flagTheme string

func init() {
	flagTheme := flag.String("theme", "dark", "Theme to use {dark|light}.")
	defaultScaleFactor := flag.Float64("scaling", 1.0, "Adjusts the scaling of UI rendering")
	flag.Parse()

	_defaultScaleFactor = float32(*defaultScaleFactor)
	_flagTheme = *flagTheme
}

func CreateTheme(driver gxui.Driver) gxui.Theme {
	if _flagTheme == "light" {
		return light.CreateTheme(driver)
	}
	return dark.CreateTheme(driver)
}

func panelHolder(name string, panel gxui.LinearLayout, theme gxui.Theme) gxui.PanelHolder {
	holder := theme.CreatePanelHolder()
	holder.AddPanel(panel, name)
	return holder
}

func appMain(driver gxui.Driver) {
	theme := CreateTheme(driver)

	label := func(text string) gxui.LinearLayout {
		label := theme.CreateLabel()
		label.SetText(text)
		layout := theme.CreateLinearLayout()
		layout.SetDirection(gxui.LeftToRight)
		layout.AddChild(label)
		return layout
	}

	// ┌───────┐║┌───────┐
	// │       │║│       │
	// │   A   │║│   B   │
	// │       │║│       │
	// │       │║└───────┘
	// │       │║═════════
	// │       │║┌───────┐
	// │       │║│       │
	// │       │║│   D   │
	// │       │║│       │
	// └───────┘║└───────┘

	// Left
	leftPanelHolder := theme.CreatePanelHolder()
	leftPanelHolder.AddPanel(label("A content"), "A")

	splitterLeft := theme.CreateSplitterLayout()
	splitterLeft.SetOrientation(gxui.Vertical)
	splitterLeft.AddChild(leftPanelHolder)

	// Right
	rightBPanelHolder := theme.CreatePanelHolder()
	rightBPanelHolder.AddPanel(label("B 1 content"), "B 1")
	rightBPanelHolder.AddPanel(label("B 2 content"), "B 2")

	rightDPanelHolder := theme.CreatePanelHolder()
	rightDPanelHolder.AddPanel(label("D 1 content"), "D 1")
	rightDPanelHolder.AddPanel(label("D 2 content"), "D 2")
	rightDPanelHolder.AddPanel(label("D 3 content"), "D 3")

	splitterRight := theme.CreateSplitterLayout()
	splitterRight.SetOrientation(gxui.Vertical)
	splitterRight.AddChild(rightBPanelHolder)
	splitterRight.AddChild(rightDPanelHolder)

	splitter := theme.CreateSplitterLayout()
	splitter.SetOrientation(gxui.Horizontal)
	splitter.AddChild(splitterLeft)
	splitter.AddChild(splitterRight)

	// Main Panel
	mainPanelHolder := theme.CreatePanelHolder()
	mainPanelHolder.AddPanel(grunt.CreateGrunt(theme), "Grunt")
	mainPanelHolder.AddPanel(hosts.CreateHosts(theme), "Hosts")
	mainPanelHolder.AddPanel(splitter, "In Feature Developing")

	window := theme.CreateWindow(800, 600, "Code tools")
	window.SetScale(_defaultScaleFactor)
	window.SetPadding(math.Spacing{
		//left
		L: 10,
		//right
		R: 10,
		//bottom
		B: 10,
		//top
		T: 10,
	})
	window.OnClose(driver.Terminate)
	window.AddChild(mainPanelHolder)
}

func CreateWindow() {
	gl.StartDriver(appMain)
}

func Main() {
	CreateWindow()
}

func main() {
	Main()
}
