package modules

import (
	"github.com/verdverm/tview"

	"github.com/hofstadter-io/connector-go"

	"github.com/verdverm/vermui-starterkit/modules/demos"
	"github.com/verdverm/vermui-starterkit/modules/help"
	"github.com/verdverm/vermui-starterkit/modules/home"
	"github.com/verdverm/vermui-starterkit/modules/root"
)

var (
	Module   connector.Connector
	rootView tview.Primitive
)

func Init() {
	rootView = root.New()

	items := []interface{}{
		// primary layout components
		rootView,

		// routable pages
		home.New(),
		help.New(),
		demos.New(),
	}

	conn := connector.New("root")
	conn.Add(items)
	Module = conn

	Module.Connect(Module)
}

func RootView() tview.Primitive {
	return rootView
}
