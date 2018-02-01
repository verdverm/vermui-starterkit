package home

import "github.com/hofstadter-io/connector-go"

func New() connector.Connector {
	items := []interface{}{
		NewHome(),
	}
	m := connector.New("home")
	m.Add(items)

	return m
}
