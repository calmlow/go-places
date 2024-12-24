package ui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/calmlow/go-places/internal/io"
	"github.com/calmlow/go-places/internal/places"
	"github.com/calmlow/go-places/internal/types"
)

const (
	PAGE1_MAIN      = "main"
	PAGE2_REPO_HOME = "repoHome"
	COLS_MAX        = 435
	N_FOLDER        = "\ue613"
)

func RunGui() {

	app := tview.NewApplication()
	pages := tview.NewPages()
	flex := tview.NewFlex()

	list := tview.NewList()
	list.ShowSecondaryText(false)
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetBorderPadding(0, 0, 2, 0)
	list.SetMainTextColor(tcell.ColorBisque)
	list.SetHighlightFullLine(false)
	list.SetSelectedBackgroundColor(tcell.Color122)

	flex.AddItem(list, COLS_MAX, 0, true)

	placesArr, err := places.GetPlaces()
	if err != nil {
		// Error is "caught" here to show the error in the UI
		eName := fmt.Sprintf("Problem getting the places list from config. %v", err)
		errorPlace := types.Place{}
		errorPlace.Name = eName
		placesArr = []types.Place{}
		placesArr = append(placesArr, errorPlace)

	}
	hiddenPlacesArr, err2 := places.GetHiddenPlaces()
	if err2 != nil {
		log.Printf("Problem getting the hidden places list from config. %v", err)
	}

	exitApp := func(exitCode int) {
		app.Stop()
		os.Exit(exitCode)
	}
	onBackClick := func() {
		pages.RemovePage(PAGE2_REPO_HOME)
	}

	contextMenuOffset := 0

	twoLevelMenu := func() {
		selectedIndex := list.GetCurrentItem() - contextMenuOffset
		selectedRepo := placesArr[selectedIndex]

		repoHomeFlex := tview.NewFlex()

		textView := GetHelpTextView(selectedRepo)

		rightList := NewRepoHomeList(selectedRepo, exitApp, onBackClick)

		repoHomeFlex.AddItem(rightList, 65, 0, true)
		repoHomeFlex.AddItem(textView, 100, 0, true)

		pages.AddPage(PAGE2_REPO_HOME, repoHomeFlex, true, false)
		pages.SwitchToPage(PAGE2_REPO_HOME)
	}

	noSubMenu := func() {
		selectedIndex := list.GetCurrentItem() - contextMenuOffset
		selectedPlace := placesArr[selectedIndex]
		log.Printf("Going to file: %s\n", selectedPlace.Path)
		io.WriteTmpFile(io.CleanRepoPath(selectedPlace.Path))
		exitApp(0)
	}

	addListItems(list, noSubMenu, twoLevelMenu, placesArr)

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
		os.Exit(1)
	})

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyInsert:
			contextMenuOffset = 1
			placesArr = append(placesArr, hiddenPlacesArr...)
			addListItems(list, noSubMenu, twoLevelMenu, hiddenPlacesArr)
			return nil
		}
		return event
	})

	list.SetDoneFunc(func() {
		app.Stop()
		os.Exit(1)
	})

	pages.AddPage(PAGE1_MAIN, flex, true, true)

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}

}

func addListItems(list *tview.List, noSubMenu func(), selectedMenu func(), repoList []types.Place) {
	for _, repo := range repoList {

		if fi, err := os.Stat(io.CleanRepoPath(repo.Path)); err == nil {
			if fi.Mode().IsDir() {
				list.AddItem(N_FOLDER+" "+repo.Name, repo.Description, repo.ShortcutAsRune(), selectedMenu)
			} else {
				list.AddItem(getIcon(repo.Path)+" "+repo.Name, repo.Description, repo.ShortcutAsRune(), noSubMenu)
			}
		} else {
			panic(fmt.Errorf("one of the paths in your config doesn't resolve to a place in the file system: %s. %v", repo.Path, err))
		}
	}
}

func getIconOld(p string) string {

	if strings.Contains(p, "linux") {
		return "\u033d"
	}

	switch filepath.Ext(p) {
	case ".jpg", ".jpeg", ".png", ".gif":
		return "\ue70f"
	case ".yaml", ".yml":
		return "\uf301"
	case ".pdf":
		return "\ue737"
	case ".go", ".java", ".py":
		return "\ue751"
	case ".js":
		return "\ue781"
	case ".md":
		return "\uf48a"
	case ".zip", ".tar", ".gz":
		return "\ue79b"
	default:
		return "\ue709"
	}

}

func getIcon(p string) string {

	if strings.Contains(p, "linux") {
		return "\u033d" // Or a more appropriate Linux icon
	}

	switch filepath.Ext(p) {
	case ".jpg", ".jpeg", ".png", ".gif", ".svg", ".bmp", ".tiff":
		return "\ue70f" // Image file
	case ".yaml", ".yml":
		return "\uf89c" // YAML file (Nerd Font icon)
	case ".pdf":
		return "\ue737" // PDF file
	case ".go":
		return "\ue626" // Go file (Nerd Font icon)
	case ".java":
		return "\ue738" // Java file (Nerd Font icon)
	case ".py":
		return "\ue606" // Python file (Nerd Font icon)
	case ".js", ".jsx", ".ts", ".tsx":
		return "\ue74e" // Javascript/Typescript file (Nerd Font icon)
	case ".html", ".htm":
		return "\uf13b" // HTML file (Nerd Font icon)
	case ".css":
		return "\ue749" // CSS file (Nerd Font icon)
	case ".json":
		return "\ue60b" // JSON file (Nerd Font icon)
	case ".md":
		return "\uf48a" // Markdown file
	case ".zip", ".tar", ".gz", ".rar", ".7z":
		return "\ue79b" // Archive file
	case ".doc", ".docx":
		return "\uf1c2" // Word document (Nerd Font icon)
	case ".xls", ".xlsx":
		return "\uf1c3" // Excel spreadsheet (Nerd Font icon)
	case ".ppt", ".pptx":
		return "\uf1c4" // PowerPoint presentation (Nerd Font icon)
	case ".txt":
		return "\uf0f6" // Text file (Nerd Font icon)
	default:
		return "\ue709" // Generic file
	}
}
