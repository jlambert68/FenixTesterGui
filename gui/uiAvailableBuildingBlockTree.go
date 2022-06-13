package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"time"
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
			return list[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := list[uid]

			return ok && len(children) > 0
		},

		CreateNode: func(branch bool) fyne.CanvasObject {
			fmt.Println("CreateNode: ")
			return widget.NewLabel("Collection Widgets: ")
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
			obj.(*widget.Label).SetText(uid + time.Now().String())
			fmt.Println(tree.Size())
		},

		OnSelected: func(uid string) {
			fmt.Println(uid, list[uid])
			//if t, ok := list[uid]; ok {
			//	fmt.Println(tree.Root)
			//	fmt.Println(t)

			//}
		},
	}

}

// Load all Available Building Blocks from Gui-server
func (uiServer *UIServerStruct) loadAvailableBuildingBlocksFromServer() {

	var fenixGuiTestCaseBuilderServerGrpcApi *fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//grpcOut := grpc_out.GRPCOutStruct{}
	fenixGuiTestCaseBuilderServerGrpcApi = uiServer.grpcOut.SendGetTestInstructionsAndTestContainers("s41797")

	fmt.Println(fenixGuiTestCaseBuilderServerGrpcApi)

}
