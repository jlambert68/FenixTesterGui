package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

type testCaseUIStruct struct {
	//current *note
	//notes   *notelist

	content *widget.Entry
	//list    *widget.List

	tree     *widget.Label // *widget.Tree
	testcase *widget.Label
}

var image *canvas.Image

// Main UI server module
func (uiServerStruct *UIServerStruct) StartUIServer() {

	uiServerStruct.logger.WithFields(logrus.Fields{
		"id": "a4d2716f-ded1-4062-bffb-fd0c03d69ca3",
	}).Debug("Starting UI server")

	a := app.NewWithID("se.fenix.testcasebuilder")
	//a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Fenix TestCase Builder")

	makeTree()

	//list := &notelist{pref: a.Preferences()}
	//list.load()
	//builderUI := &testCaseUIStruct{notes: list}
	builderUI := &testCaseUIStruct{
		content:  nil,
		tree:     nil,
		testcase: nil,
	}
	w.SetContent(builderUI.loadUI())

	//w.SetContent(widget.NewLabel("Fenix TestCase Builder"))
	//builderUI.registerKeys(w)

	w.Resize(fyne.NewSize(400, 320))
	w.ShowAndRun()

}

type MouseHandler struct {
	widget.BaseWidget
}

/*
func (m *MouseHandler) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(canvas.NewCircle(color.RGBA{255, 0, 0, 255}))
}

*/

type CustomButton struct {
	widget.Button
}

func (m *CustomButton) MouseIn(e *desktop.MouseEvent) {
	fmt.Println("Mouse In")
	fmt.Println(m.Position())
	mypos := fyne.NewPos(float32(200), float32(200))
	image.Move(mypos)
	image.Show()
	image.Refresh()

}

/*
func (m *CustomButton) MouseMoved(e *desktop.MouseEvent) {
	fmt.Println("Mouse Moved")
}
*/
func (m *CustomButton) MouseOut() {
	fmt.Println("Mouse Out")
	image.Hide()
	image.Refresh()
}

func (testCaseUI *testCaseUIStruct) loadUI() fyne.CanvasObject {

	var _ desktop.Hoverable = (*CustomButton)(nil)

	testCaseUI.tree = widget.NewLabel("Available TestInstructions")
	//testCaseUI.testcase = widget.NewLabel("TestCase Builder Area")
	//testCaseUI.treeContainer = container.New(layout.NewHBoxLayout(), treeCanvasObject, layout.NewSpacer())
	testCaseUI.content = widget.NewMultiLineEntry()
	testCaseUI.content.SetText("Here you will build the TestCase")
	/*
		myButton := widget.Button{
			DisableableWidget: widget.DisableableWidget{},
			Text:              "My new Button",
			OnTapped: func() {
				fmt.Println("MyButton was clicked...")
			},
		}

		//myExtededButton := &CustomButton{myButton}

	*/

	//text1 := canvas.NewText("Hello", color.White)

	treeSide := testCaseUI.loadCompleteAvailableTestCaseBuildingBlocksUI()

	testCaseSide := testCaseUI.loadCompleteCurrentTestCaseUI()

	uiStructureContainer := newAdaptiveSplit(treeSide, testCaseSide)

	return uiStructureContainer
}

// Loads available TestInstructions and TestInstructionContainers and return the UI Bar and UI Tree-structure for them
func (testCaseUI *testCaseUIStruct) loadCompleteAvailableTestCaseBuildingBlocksUI() (completeAvailableTestCaseBuildingBlocksUI fyne.CanvasObject) {

	// Create toolbar for Available TestCase BuildingBlock area
	availableAvailableBuildingBlocksUIBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			fmt.Println("Add to Pinned")
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			fmt.Println("Remove from Pinned")
		}),
	)

	// Load the Available TestCase BuildingBlocks TreeUI
	availableTestCaseBuildingBlocksTreeUI := testCaseUI.loadAvailableTestCaseBuildingBlocksTreeUI()

	// Create the complete TestCase BuildingBlocks UI area
	availableTestCaseBuildingBlocksBorderedLayout := layout.NewBorderLayout(availableAvailableBuildingBlocksUIBar, nil, nil, nil)
	completeAvailableTestCaseBuildingBlocksUI = container.New(availableTestCaseBuildingBlocksBorderedLayout, availableAvailableBuildingBlocksUIBar, availableTestCaseBuildingBlocksTreeUI)

	return completeAvailableTestCaseBuildingBlocksUI
}

// Loads current BuildingBlocksTree UI-structure
func (testCaseUI *testCaseUIStruct) loadAvailableTestCaseBuildingBlocksTreeUI() (availableTestCaseBuildingBlocksTreeUI fyne.CanvasObject) {

	//availableTestCaseBuildingBlocksTreeUI = widget.NewLabel("'currentTestCaseModelAreaUI'")
	availableTestCaseBuildingBlocksTreeUI = tree

	return availableTestCaseBuildingBlocksTreeUI
}

// Loads current TestCase return the UI-structure for it
func (testCaseUI *testCaseUIStruct) loadCompleteCurrentTestCaseUI() (completeCurrentTestCaseUIContainer fyne.CanvasObject) {

	// Create toolbar for TestCase area
	testCaseToolUIBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload GUI TestCase from model")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy Node")
		}),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut Node")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Past Node")
		}),
	)

	// Load the TestCase model UI area
	currentTestCaseModelAreaUI := testCaseUI.loadCurrentTestCaseModelAreaUI()

	// Load the TestCase attributes UI area
	currentTestCaseAttributesAreaUI := testCaseUI.loadCurrentTestCaseAttributesAreaUI()

	// Create the UI area for both TestCase model UI and TestCase attributes UI
	testCaseAdaptiveSplitLayoutContainer := newAdaptiveSplit(currentTestCaseModelAreaUI, currentTestCaseAttributesAreaUI)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(testCaseToolUIBar, nil, nil, nil)
	completeCurrentTestCaseUIContainer = container.New(testCaseBorderedLayout, testCaseToolUIBar, testCaseAdaptiveSplitLayoutContainer)

	return completeCurrentTestCaseUIContainer
}

// Loads current TestCase model and return the UI-structure for it
func (testCaseUI *testCaseUIStruct) loadCurrentTestCaseModelAreaUI() (currentTestCaseModelAreaUI fyne.CanvasObject) {

	currentTestCaseModelAreaUI = widget.NewLabel("'currentTestCaseModelAreaUI'")

	return currentTestCaseModelAreaUI
}

// Loads current TestCase attributes and return the UI-structure for it
func (testCaseUI *testCaseUIStruct) loadCurrentTestCaseAttributesAreaUI() (currentTestCaseAttributesAreaUI fyne.CanvasObject) {

	currentTestCaseAttributesAreaUI = widget.NewLabel("'currentTestCaseAttributesAreaUI'")

	return currentTestCaseAttributesAreaUI
}
