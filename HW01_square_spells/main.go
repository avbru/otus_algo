package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	spells := newSpells()
	go func() {
		for _, spell := range spells {
			g.Update(spell.layout)
			time.Sleep(time.Second * 2)
		}
		os.Exit(0)
	}()

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func (sp Spell) layout(g *gocui.Gui) error {
	tname := time.Now().String()
	if v, err := g.SetView(tname, 0, 0, 52, 26); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return fmt.Errorf("unable to set view: %w", err)
		}
		for x := 0; x < 25; x++ {
			for y := 0; y < 25; y++ {
				if sp(x, y) {
					fmt.Fprint(v, "# ")
				} else {
					fmt.Fprint(v, ". ")
				}
			}
			fmt.Fprint(v, "\n")
		}

		if _, err := g.SetCurrentView(tname); err != nil {
			return fmt.Errorf("unable to set current view: %w", err)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
