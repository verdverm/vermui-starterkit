package root

import (
	"github.com/gdamore/tcell"
	"github.com/verdverm/tview"
	"github.com/verdverm/vermui/hoc/cmdbox"
	"github.com/verdverm/vermui/hoc/console"
	"github.com/verdverm/vermui/hoc/layouts/panels"
	"github.com/verdverm/vermui/hoc/statusbar"

	"github.com/hofstadter-io/connector-go"
)

type Commandables interface {
	Commands() []Command
}
type Command interface {
	CommandName() string
	CommandUsage() string
	CommandHelp() string
	CommandCallback(args []string, context map[string]interface{})
}

type RootView struct {
	*panels.Layout

	//
	// Top Panel elements
	//
	// Always Visible
	cbox *cmdbox.CmdBoxWidget
	sbar *statusbar.StatusBar
	// Hidden
	errConsole *console.ErrConsoleWidget

	//
	// Main Panel element
	//
	mainPanel *MainPanel

	//
	// Bottom Panel elements
	//
	devConsole *console.DevConsoleWidget
}

func New() *RootView {

	V := &RootView{
		Layout: panels.New(),
	}

	V.SetDirection(tview.FlexRow)

	V.buildTopPanel()

	V.buildMainPanel()

	V.buildBotPanel()

	return V
}

func (V *RootView) Connect(C connector.Connector) {
	cmds := C.Get((*Command)(nil))
	for _, Cmd := range cmds {
		cmd := Cmd.(Command)
		V.cbox.AddCommand(cmd)
	}

	V.mainPanel.Connect(C)
}

func (V *RootView) buildTopPanel() {
	V.cbox = cmdbox.New()
	V.cbox.
		SetTitle("  VermUI  ").
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorLime).
		SetBorder(true).
		SetBorderColor(tcell.Color27)

	V.sbar = statusbar.New()

	// topBar is a Flex with 2 columns
	topBar := tview.NewFlex().SetDirection(tview.FlexColumn)
	topBar.AddItem(V.cbox, 0, 1, false)
	topBar.AddItem(V.sbar, 0, 1, true)

	// error console
	V.errConsole = console.NewErrConsoleWidget()

	// Top Panels
	V.AddFirstPanel("top-bar", topBar, 3, 0, 0, "", false, "")
	V.AddFirstPanel("err-console", V.errConsole, 0, 1, 0, "", true, "C-e")

}

func (V *RootView) buildMainPanel() {
	// A Horizontal Layout with a Router as the main element
	V.mainPanel = NewMainPanel()
	V.SetMainPanel("main-panel", V.mainPanel, 0, 1, 0, "C-a")
}

func (V *RootView) buildBotPanel() {
	// dev console
	V.devConsole = console.NewDevConsoleWidget()

	// Bottom Panels
	V.AddLastPanel("dev-console", V.devConsole, 0, 1, 1, "", true, "C-l")
}
