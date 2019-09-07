package main

import (
	"github.com/whtiehack/wingui"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	log.Printf("resource %v %#[1]v  \n", IDD_DIALOG)
	var err error
	dlg, err = wingui.NewDialog(IDD_DIALOG, 0)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	dlg.SetIcon(IDI_ICON1)
	log.Println("dlg create end", dlg)
	var btn *wingui.Button
	btn, _ = dlg.NewButton(IDB_OK)
	btn.OnClicked = modalBtnClicked
	closeBtn, _ := dlg.NewButton(IDB_CANCEL)
	closeBtn.OnClicked = func() {
		dlg.Close()
	}
	dlg.Show()
	wingui.MessageLoop()
	log.Println("stoped")
}

func modalBtnClicked() {
	log.Println("btn clicked")
	wingui.NewModalDialog(IDD_DIALOG_OK, dlg.Handle(), func(okdlg *wingui.Dialog) {
		okbtn, _ := okdlg.NewButton(IDB_OK)
		okbtn.OnClicked = func() {
			log.Println("modal btn click")
			okdlg.Close()
		}
	})
}
