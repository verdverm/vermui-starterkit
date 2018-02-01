package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/verdverm/vermui"
	"github.com/verdverm/vermui/events"

	"github.com/verdverm/vermui-starterkit/modules"
)

func main() {
	// initialize VermUI
	err := vermui.Init()
	if err != nil {
		panic(err)
	}

	// initialize our modules
	modules.Init()

	// Set the root view
	root := modules.RootView()
	vermui.SetRootView(root)

	// Ctrl-c to quit program
	vermui.AddGlobalHandler("/sys/key/C-c", func(e events.Event) {
		vermui.Stop()
	})

	// Log Key presses (if you want to)
	logKeys()

	// Run PProf (useful for catching hangs)
	// go runPprofServer()

	// catch panics and exit, vermui will catch, clean up, format error, print, and repanic
	defer func() {
		err := recover()
		if err != nil {
			vermui.Stop()
			panic(err)
		}
	}()

	go func() {
		// some latent locksups occur randomly
		time.Sleep(time.Millisecond * 10)
		events.SendCustomEvent("/router/dispatch", "/")
		events.SendCustomEvent("/status/message", "Welcome to [lime]VermUI[white]!!")
	}()

	// Start the Main (Blocking) Loop
	vermui.Start()
}

func logKeys() {
	vermui.AddGlobalHandler("/sys/key", func(e events.Event) {
		if k, ok := e.Data.(events.EventKey); ok {
			go events.SendCustomEvent("/console/key", k.KeyStr)
		}
	})
}
func runPprofServer() {
	runtime.SetMutexProfileFraction(1)
	http.ListenAndServe(":8888", nil)
}
