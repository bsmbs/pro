package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	haxory = map[string]string{
		"a": "4",
		"b": "8",
		"c": "[",
		"d": ")",
		"e": "3",
		"f": "|\"",
		"g": "6",
		"h": "#",
		"i": "1",
		"j": "_|",
		"k": "x",
		"l": "1",
		"m": "|v|",
		"n": "^/",
		"o": "0",
		"p": "|*",
		"q": "(_,)",
		"r": "|`",
		"s": "5",
		"t": "7",
		"u": "(_)",
		"v": "\\/",
		"w": "\\/\\/",
		"x": "><",
		"y": "\\/",
		"z": "2",
	}
	samologski = []string{"a", "e", "i", "o", "u", "y"}
)

func main() {
	app, err := gtk.ApplicationNew("io.github.pizza61.pro", glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	app.Connect("activate", func() {
		bd, err := gtk.BuilderNew()
		if err != nil {
			log.Fatal("Could not", err)
		}
		glad, err := Asset("lib/bs.glade")
		if err != nil {
			log.Fatal("Could not", err)
		}

		icon, err := Asset("lib/icon.png")
		if err != nil {
			log.Fatal("Could not", err)
		}

		bd.AddFromString(string(glad))
		win, err := bd.GetObject("Window")
		if err != nil {
			log.Fatal("Could not", err)
		}
		window := win.(*gtk.Window)
		window.SetTitle("pr0")

		err = ioutil.WriteFile("/tmp/haxor.png", icon, 0644)
		window.SetIconFromFile("/tmp/haxor.png")
		window.Show()

		resultB, err := bd.GetObject("word")
		if err != nil {
			log.Fatal("Could not", err)
		}
		result := resultB.(*gtk.Entry)
		result.SetPlaceholderText("Wpisz tu swój tekst")

		/********************
			h4x0r
		*********************/
		btn1B, err := bd.GetObject("b1")
		if err != nil {
			log.Fatal("Could not", err)
		}

		btn1 := btn1B.(*gtk.Button)
		btn1.Connect("clicked", func() {
			tt, _ := result.GetText()
			result.SetText(hackowanie(tt, false))
		})
		/********************
			#4><0r
		*********************/
		btn2B, err := bd.GetObject("b2")
		if err != nil {
			log.Fatal("Could not", err)
		}

		btn2 := btn2B.(*gtk.Button)
		btn2.Connect("clicked", func() {
			tt, _ := result.GetText()
			result.SetText(hackowanie(tt, true))
		})

		/********************
			hxr
		*********************/
		btn3B, err := bd.GetObject("b3")
		if err != nil {
			log.Fatal("Could not", err)
		}

		btn3 := btn3B.(*gtk.Button)
		btn3.Connect("clicked", func() {
			tt, _ := result.GetText()
			result.SetText(usuwanieSamoglosek(tt))
		})

		app.AddWindow(window)
	})
	app.Run(os.Args)
}

func hackowanie(wejscie string, pro bool) string {
	wrra := strings.Split(wejscie, "")

	if !pro {
		for i, l := range wrra {
			ten := strings.ToLower(l)
			if ten == "a" || ten == "e" || ten == "i" || ten == "l" || ten == "o" || ten == "s" || ten == "t" || ten == "z" {
				wrra[i] = haxory[ten]
			}
		}
	} else {
		for i, l := range wrra {
			ten := strings.ToLower(l)
			wrra[i] = haxory[ten]
		}
	}
	wejscie = strings.Join(wrra, "")
	return wejscie
}

func usuwanieSamoglosek(wejscie string) string {
	if len(wejscie) > 0 {
		ostatniaLiterka := wejscie[len(wejscie)-1:]
		wejscie = wejscie[:len(wejscie)-1]

		replacer := strings.NewReplacer("a", "", "e", "", "i", "", "o", "", "u", "", "y", "")
		wejscie = replacer.Replace(wejscie) + ostatniaLiterka
	}

	return wejscie
}
