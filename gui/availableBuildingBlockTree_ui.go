package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
	//"golang.org/x/exp/maps"
)

//______________________________________________________________________________

var list map[string][]string
var tree *widget.Tree

func (uiServer *UIServerStruct) makeTreeUI() {
	list = map[string][]string{
		"":  {"A"},
		"A": {"B", "D"},
		"B": {"C"},
		"C": {"abc"},
		"D": {"E"},
		"E": {"F", "G"},
	}

	/*

		tree = widget.NewTreeWithStrings(list)
		tree.OnSelected = func(id string) {
			dbg.Green("Tree node selected: %s", id)

		}
		tree.OnUnselected = func(id string) {
			dbg.Red("Tree node unselected: %s", id)
		}

		tree.OpenAllBranches()

	*/

	tree = &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return uiServer.availableBuildingBlocksModel.getAvailableBuildingBlocksModel()[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := uiServer.availableBuildingBlocksModel.getAvailableBuildingBlocksModel()[uid]

			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {
			fmt.Println("CreateNode: ")
			return newTappableLabel() //widget.NewLabel("Collection Widgets: ")
		},

		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			fmt.Println("UpdateNode: ", uid)
			/*
				_, ok := list[uid]
				if !ok {
					fyne.LogError("Missing tutorial panel: "+uid, nil)
					return
				}
			*/
			obj.(*tappableLabel).SetText(uid) //obj.(*widget.Label).SetText(uid) // + time.Now().String())
			fmt.Println(tree.Size())
		},

		OnSelected: func(uid string) {
			fmt.Println(uid, uiServer.availableBuildingBlocksModel.getAvailableBuildingBlocksModel()[uid])
			//if t, ok := list[uid]; ok {
			//	fmt.Println(tree.Root)
			//	fmt.Println(t)

			//}
		},
	}

}

type tappableLabel struct {
	widget.Label
	movableLable *widget.Label
}

func newTappableLabel() *tappableLabel {
	label := &tappableLabel{}
	label.ExtendBaseWidget(label)
	//icon.SetResource(res)

	return label
}

func (t *tappableLabel) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")

	fmt.Println(t.Position())
}

func (t *tappableLabel) TappedSecondary(_ *fyne.PointEvent) {
	log.Println("I have been Secondary tapped")
}

func (t *tappableLabel) Dragged(ev *fyne.DragEvent) {
	log.Println("I have been Dragged: ", t.Position())
	t.movableLable = widget.NewLabel("DRAGGED")
	t.movableLable.Move(ev.Position)
	t.TextStyle.Bold = true
	fmt.Println(fmt.Println(t.Text))

}

func (t *tappableLabel) DragEnd() {
	log.Println("I have been DragEnd")
	t.TextStyle.Bold = false
}
