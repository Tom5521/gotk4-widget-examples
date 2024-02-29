package main

import (
	"math/rand/v2"
	"os"

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

func activate(app *gtk.Application) {
	w := gtk.NewApplicationWindow(app)

	// GtkFixed places its child widgets at fixed positions and with fixed sizes.
	// GtkFixed performs no automatic layout management.
	//
	// For most applications, you should not use this container!
	// It keeps you from having to learn about the other GTK containers,
	// but it results in broken applications. With GtkFixed,
	// the following things will result in truncated text, overlapping widgets,
	// and other display bugs.
	fixed := gtk.NewFixed()

	setController := func(widget gtk.Widgetter) *gtk.EventControllerMotion {
		// We create a new controller.
		controller := gtk.NewEventControllerMotion()
		controller.ConnectEnter(func(_, _ float64) {
			// Create the new coordinates for the widget based on the window size.
			newX := float64(rand.N(w.Width()))
			newY := float64(rand.N(w.Height()))
			// Apply the new coordinates.
			fixed.Move(widget, newX, newY)
		})
		return controller
	}
	newLabel := func(text string) *gtk.Label {
		l := gtk.NewLabel(text)
		l.AddController(setController(l))
		return l
	}

	button := gtk.NewButtonWithLabel("Click Me!")
	button.ConnectClicked(func() {
		button.SetLabel("How do you clicked me?")
	})
	button.AddController(setController(button))

	fixed.Put(button, 200, 200)
	fixed.Put(newLabel("Hi world!"), 200, 300)
	fixed.Put(newLabel("Hi world! (2)"), 500, 400)

	w.SetChild(fixed)
	w.Show()
}
