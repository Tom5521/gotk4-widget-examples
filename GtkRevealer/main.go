package main

import (
	"os"
	"time"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.test.window", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		activate(app)
	})
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

const loremIpsum string = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	tv := gtk.NewTextView()
	tv.SetVExpand(true)
	tb := tv.Buffer()

	tb.SetText(loremIpsum)
	tv.SetWrapMode(gtk.WrapWordChar)

	revealer := gtk.NewRevealer()
	revealer.SetChild(tv)
	go func() {
		time.Sleep(4 * time.Second)
		revealer.SetRevealChild(true)
	}()

	w.SetChild(revealer)
	w.Show()
}
