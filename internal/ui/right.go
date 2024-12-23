package ui

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/calmlow/go-places/internal/io"
	"github.com/calmlow/go-places/internal/types"
)

const GITHUB_ORG_URL = "https://github.com/PricerAB"

func NewRepoHomeList(selectedRepo types.Place, exitApp func(exitCode int), onBackClick func()) *tview.List {
	title := fmt.Sprintf(" Selected Repository: %s ", selectedRepo.Name)

	rightList := tview.NewList()
	rightList.ShowSecondaryText(true)
	rightList.SetMainTextColor(tcell.ColorYellowGreen)

	rightList.SetBorder(true).SetTitle(title)

	// escape key is pressed - go back
	rightList.SetDoneFunc(onBackClick)

	rightList.AddItem("Goto "+selectedRepo.Name, "Browse to currrently selected repo in console", 'v', func() {
		log.Printf("Going to: %s\n", selectedRepo.Path)
		io.WriteTmpFile(io.CleanRepoPath(selectedRepo.Path))
		exitApp(0)
	})
	rightList.AddItem("Visit GitHub: "+selectedRepo.Name, "Visit GitHub for this repository", 'g', func() {
		fmt.Println("gh browse: ", selectedRepo.Path)
		githubUrl := fmt.Sprintf("%s/%s", GITHUB_ORG_URL, selectedRepo.Name)
		io.OpenBrowser(githubUrl)
		exitApp(0)
	})
	rightList.AddItem("Actions: "+selectedRepo.Name, "Visit GitHub Actions for this repository", 'a', func() {
		fmt.Println("gh browse: ", selectedRepo.Path)
		githubUrl := fmt.Sprintf("%s/%s/actions", GITHUB_ORG_URL, selectedRepo.Name)
		io.OpenBrowser(githubUrl)
		exitApp(0)
	})
	rightList.AddItem("Docs URL for "+selectedRepo.Name, "Visit GitHub Docs for this repository", 'd', func() {
		_, err := url.ParseRequestURI(selectedRepo.DocsUrl)
		if err != nil {
			panic(err)
		} else {
			io.OpenBrowser(selectedRepo.DocsUrl)
		}
		exitApp(0)
	})

	rightList.AddItem("Back", "Back to selector - also ESC key", 'b', onBackClick)

	rightList.AddItem("Quit", "Press to exit", 'q', func() {
		fmt.Println("Quitting...")
		exitApp(1)
	})

	return rightList
}
