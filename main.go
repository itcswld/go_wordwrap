package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"wordwrap/xls"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("")
	a.Settings().SetTheme(theme.DarkTheme())
	w.Resize(fyne.NewSize(400, 300))
	img, _ := filepath.Abs("logo_png.png")
	r, _ := fyne.LoadResourceFromPath(img)
	w.SetIcon(r)

	e_sheet := widget.NewEntry()
	e_sheet.SetPlaceHolder("sheet")
	e_sheet.Resize(fyne.NewSize(250, 30))

	e_copy := widget.NewEntry()
	e_copy.Text = "1"
	e_copy.Validator = func(s string) error {
		var err error
		is_numeric := regexp.MustCompile(`\d`).MatchString(s)
		if !is_numeric {
			err = errors.New("need number")
			dialog.NewInformation("err:", err.Error(), w)
			w.Canvas().Focus(e_copy)
		}
		return err
	}
	e_copy.Resize(fyne.NewSize(250, 30))
	form := &widget.Form{
		Items: []*widget.FormItem{{
			Text:   "sheet name",
			Widget: e_sheet,
		}},
		SubmitText: "print",
		OnSubmit: func() {
			data := xls.GetXlsData("./xls/test.xlsx", e_sheet.Text)
			for i, v := range data {
				fmt.Println(i, ",", v)
			}
			e_sheet.SetText("")
			w.Canvas().Focus(e_sheet)
		},
	}

	w.SetContent(container.NewVBox(form))

	w.ShowAndRun()
}
