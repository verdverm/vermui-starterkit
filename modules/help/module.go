package help

import "github.com/hofstadter-io/connector-go"

func New() connector.Connector {
	items := []interface{}{
		NewHelp(),
	}
	m := connector.New("help")
	m.Add(items)

	return m
}
