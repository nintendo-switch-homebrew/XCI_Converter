/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"container/list"
	"fmt"
	"log"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

var labelList = list.New()
var win *gtk.Window

func gui() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("xci_Converter")
	win.SetDefaultSize(280, 120)
	win.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)
	//if err != nil {
	//log.Fatal(err)
	//}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget(win))
	win.ShowAll()

	gtk.Main()
}

func hello() {
	fmt.Println("Hello")
}

func windowWidget(win *gtk.Window) *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	titleIDStr, err := gtk.LabelNew("Original title ID ")
	if err != nil {
		log.Fatal("Unable to create new label", err)
	}
	titleID, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create scrolled window:", err)
	}

	xciStr, err := gtk.LabelNew("Path to your XCI ")
	if err != nil {
		log.Fatal("Unable to create new label", err)
	}
	xciPath, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create scrolled window:", err)
	}

	convertButton, err := gtk.ButtonNewWithLabel("Convert")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	convertButton.Connect("clicked", func() {
		var popup *gtk.MessageDialog

		titleIDBuffer, _ := titleID.GetBuffer()
		titleIDName, _ := titleIDBuffer.GetText()
		titleIDName = strings.TrimSpace(titleIDName)

		xciPathBuffer, _ := xciPath.GetBuffer()
		path, _ := xciPathBuffer.GetText()
		path = strings.TrimSpace(path)

		if isHex(titleIDName) == false {
			popup = gtk.MessageDialogNew(win,
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
				"Please set a valid titleID")
		} else if isValidFile(path) == false {
			popup = gtk.MessageDialogNew(win,
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
				"Path invalid or insufficient permission")
		} else {
			convert(titleIDName, path)
			popup = gtk.MessageDialogNew(win,
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
				"Finished !")
		}
		popup.Run()
		popup.Destroy()

	})

	grid.Attach(titleIDStr, 0, 0, 1, 1)
	grid.AttachNextTo(titleID, titleIDStr, gtk.POS_RIGHT, 1, 1)

	grid.Attach(xciStr, 0, 1, 1, 1)
	grid.AttachNextTo(xciPath, xciStr, 1, 1, 1)

	grid.Attach(convertButton, 1, 2, 1, 1)

	return &grid.Container.Widget
}
