package demos

import "github.com/hofstadter-io/connector-go"

func New() connector.Connector {
	items := []interface{}{
		NewDemos(),
	}

	m := connector.New("demos")
	m.Add(items)

	return m
}
