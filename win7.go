package main

import (
	"github.com/lxn/walk"
	."github.com/lxn/walk/declarative"
	"strings"
	"os/exec"
	"log"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title: "PHP Interpreter",
		MinSize: Size{600, 400},
		Layout: VBox{},
		Children: []Widget{
			Label{
				Text: "Escriba codigo PHP sin etiquetas de <?php apertura y cierre ?>",
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Let the magic happen!",
				OnClicked: func() {

					input := strings.Replace(inTE.Text(), "\n", " ", -1)
					out, err := exec.Command("php", "-r", input).Output()

					if err != nil {
						log.Println(err)
					}

					outTE.SetText(string(out))
				},
			},
		},
	}.Run()
}