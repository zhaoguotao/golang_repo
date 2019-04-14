package main

import (
	//	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window
var MultiEntry *ui.MultilineEntry

//var filename string

func FirstPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(false)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	MultiEntry = ui.NewNonWrappingMultilineEntry()
	buttonOpenFolder := ui.NewButton(" 打开目录   ")
	buttonOpenFolder.OnClicked(func(*ui.Button) {
		cmd := exec.Command("cmd", "/c", "start .")
		if err := cmd.Run(); err != nil {
			MultiEntry.Append("Error: " + err.Error() + "\n")
			return
		}
	})
	buttonRun := ui.NewButton("    Run    ")
	buttonRun.OnClicked(func(*ui.Button) {
		fin := entry.Text()
		if fin != "" { //First, make sure file/folder exist
			fi, err := os.Lstat(fin)
			if err == nil {
				if fi.Mode().IsDir() {
					//MultiEntry.Append(fin + " : is a folder\n")
					BatchRun(fin)
				} else if fi.Mode().IsRegular() {
					SingleRun(fin)
				} else {
					MultiEntry.Append(fin + ": not exist.\n")
					ui.MsgBox(mainwin, "Warning", fin+": not exist.\n")
				}
			} else {
				ui.MsgBoxError(mainwin, "Error", fin+" not exist.")
			}
		} else {
			ui.MsgBox(mainwin, "Warning", "Import a file first.")
		}

	})
	buttonClear := ui.NewButton("    Clear    ")
	buttonClear.OnClicked(func(*ui.Button) {
		MultiEntry.SetText("")
	})
	grid.Append(button, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	grid.Append(buttonOpenFolder, 2, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(buttonRun, 3, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(buttonClear, 4, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	entryForm.Append("", MultiEntry, true)
	vbox.Append(entryForm, true)

	return vbox
}

func setupUI() {
	mainwin = ui.NewWindow("batch rename raw v0.01", 640, 280, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	hbox.Append(FirstPage(), true)
	mainwin.SetChild(hbox)

	mainwin.Show()
}

func RenameFile(fname, new_name string) {
	// Rename or move file from one location to another.
	ret := os.Rename(fname, new_name)
	if ret != nil {
		MultiEntry.Append(ret.Error())
		return
	} else {
		MultiEntry.Append(fname + " >>> " + new_name + "\n")
		return
	}
}

func SingleRun(fname string) {
	r, _ := regexp.Compile("20\\d{6}_\\d{6}|20\\d{12}")
	// Only handle RAWPLAIN16 file
	ok := strings.HasSuffix(fname, ".RAWPLAIN16")
	if ok {
		if r.FindString(fname) != "" {
			fpre := filepath.Dir(fname) //去除最后一个元素的路径
			new_name := fpre + "/IMG_" + r.FindString(fname) + ".raw"
			RenameFile(fname, new_name)
		} else {
			MultiEntry.Append(fname + ": NO match...\n")
		}
	} else {
		MultiEntry.Append(fname + ": not a RAWPLAIN16 file\n")
	}
}

func BatchRun(path string) {
	r, _ := regexp.Compile("20\\d{6}_\\d{6}|20\\d{12}")
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// Only handle RAWPLAIN16 file
		ok := strings.HasSuffix(path, ".RAWPLAIN16")
		if ok {
			if r.FindString(path) != "" {
				fpre := filepath.Dir(path) //去除最后一个元素的路径
				new_name := fpre + "/IMG_" + r.FindString(path) + ".raw"
				RenameFile(path, new_name)
			} else {
				MultiEntry.Append(path + ": not a RAWPLAIN16 file\n")
			}
		}
		return nil
	})
	if err != nil {
		MultiEntry.Append(err.Error())
	}
}

func main() {
	ui.Main(setupUI)
}
