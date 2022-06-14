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

/*
type testCaseUIStruct struct {
	//current *note
	//notes   *notelist

	content *widget.Entry
	//list    *widget.List

	tree         *widget.Label // *widget.Tree
	testcase     *widget.Label
	commandStack *widget.List
	fyneApp      fyne.App
	logger       *logrus.Logger
}


*/
//var myTestCase *testCaseUIStruct

var image *canvas.Image

// Main UI server module
func (uiServer *UIServerStruct) StartUIServer() {

	/*
		myUIServer = UIServerStruct{
			logger:  callersLoggerReference,
			grpcOut: grpc_out.GRPCOutStruct{Logger: callersLoggerReference},
		}

	*/
	uiServer.logger.WithFields(logrus.Fields{
		"id": "a4d2716f-ded1-4062-bffb-fd0c03d69ca3",
	}).Debug("Starting UI server")
	/*
		myTestCase = &testCaseUIStruct{
			logger: myUIServer.logger,
		}

	*/
	//myUIServer = UIServerStruct{}

	//var grpcOut grpc_out.GRPCOutStruct
	// myUIServer.grpcOut.SetLogger(myUIServer.logger)

	// Add/Forward variables to packages to be used later
	uiServer.grpcOut.SetLogger(uiServer.logger)
	uiServer.grpcOut.SetDialAddressString(uiServer.fenixGuiBuilderServerAddressToDial)

	uiServer.fyneApp = app.NewWithID("se.fenix.testcasebuilder")
	//fyneApp.Settings().SetTheme(&myTheme{})
	fyneMasterWindow := uiServer.fyneApp.NewWindow("Fenix TestCase Builder")
	fyneMasterWindow.SetMaster()

	// Get Available Building BLocks form GUI-server
	uiServer.loadAvailableBuildingBlocksFromServer()

	// Initate and create the tree structure for available building blocks, of TestInstructions and TestInstructionContainers
	uiServer.makeTreeUI()

	// Initiate the commandStack which describes how fyneApp TestCase is constructed
	uiServer.makeCommandStackUI()

	// Create fyneApp window for the Command Stack
	commandStackWindow := uiServer.fyneApp.NewWindow("Command Stack")
	commandStackWindow.SetContent(commandStackListUI)
	commandStackWindow.Show()

	//list := &notelist{pref: fyneApp.Preferences()}
	//list.load()
	//builderUI := &testCaseUIStruct{notes: list}
	/*
		builderUI := &testCaseUIStruct{
			content:      nil,
			tree:         nil,
			testcase:     nil,
			commandStack: nil,
		}

	*/
	fyneMasterWindow.SetContent(uiServer.loadUI())

	//fyneMasterWindow.SetContent(widget.NewLabel("Fenix TestCase Builder"))
	//builderUI.registerKeys(fyneMasterWindow)

	fyneMasterWindow.Resize(fyne.NewSize(400, 320))

	fyneMasterWindow.ShowAndRun()

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

func (uiServer *UIServerStruct) loadUI() fyne.CanvasObject {

	var _ desktop.Hoverable = (*CustomButton)(nil)

	uiServer.tree = widget.NewLabel("Available TestInstructions")
	//testCaseUI.testcase = widget.NewLabel("TestCase Builder Area")
	//testCaseUI.treeContainer = container.New(layout.NewHBoxLayout(), treeCanvasObject, layout.NewSpacer())
	uiServer.content = widget.NewMultiLineEntry()
	uiServer.content.SetText("Here you will build the TestCase")
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

	treeSide := uiServer.loadCompleteAvailableTestCaseBuildingBlocksUI()

	testCaseSide := uiServer.loadCompleteCurrentTestCaseUI()

	uiStructureContainer := newAdaptiveSplit(treeSide, testCaseSide)

	return uiStructureContainer
}

// Loads available TestInstructions and TestInstructionContainers and return the UI Bar and UI Tree-structure for them
func (uiServer *UIServerStruct) loadCompleteAvailableTestCaseBuildingBlocksUI() (completeAvailableTestCaseBuildingBlocksUI fyne.CanvasObject) {

	// Create toolbar for Available TestCase BuildingBlock area
	availableAvailableBuildingBlocksUIBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")
			uiServer.loadAvailableBuildingBlocksFromServer()
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			fmt.Println("Add to Pinned")
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			fmt.Println("Remove from Pinned")
		}),
	)

	// Load the Available TestCase BuildingBlocks TreeUI
	availableTestCaseBuildingBlocksTreeUI := uiServer.loadAvailableTestCaseBuildingBlocksTreeUI()

	// Create the complete TestCase BuildingBlocks UI area
	availableTestCaseBuildingBlocksBorderedLayout := layout.NewBorderLayout(availableAvailableBuildingBlocksUIBar, nil, nil, nil)
	tempcompleteAvailableTestCaseBuildingBlocksUI := container.New(availableTestCaseBuildingBlocksBorderedLayout, availableAvailableBuildingBlocksUIBar, container.NewVSplit(availableTestCaseBuildingBlocksTreeUI, uiServer.createTestCaseCommandsUI()))
	//tempcompleteAvailableTestCaseBuildingBlocksUI.MinSize(fyne.NewSize(float32(300), float32(400))

	completeAvailableTestCaseBuildingBlocksUI = tempcompleteAvailableTestCaseBuildingBlocksUI //container.New(layout.NewVBoxLayout(), tempcompleteAvailableTestCaseBuildingBlocksUI) //, layout.NewSpacer(), testCaseUI.createTestCaseCommandsUI())

	return completeAvailableTestCaseBuildingBlocksUI
}

// Loads current BuildingBlocksTree UI-structure
func (uiServer *UIServerStruct) loadAvailableTestCaseBuildingBlocksTreeUI() (availableTestCaseBuildingBlocksTreeUI fyne.CanvasObject) {

	//availableTestCaseBuildingBlocksTreeUI = widget.NewLabel("'currentTestCaseModelAreaUI'")
	availableTestCaseBuildingBlocksTreeUI = tree

	return availableTestCaseBuildingBlocksTreeUI
}

// Loads current TestCase return the UI-structure for it
func (uiServer *UIServerStruct) loadCompleteCurrentTestCaseUI() (completeCurrentTestCaseUIContainer fyne.CanvasObject) {

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
	currentTestCaseModelAreaUI := uiServer.loadCurrentTestCaseModelAreaUI()

	// Load the TestCase attributes UI area
	currentTestCaseAttributesAreaUI := uiServer.loadCurrentTestCaseAttributesAreaUI()

	// Create the UI area for both TestCase model UI and TestCase attributes UI
	testCaseAdaptiveSplitLayoutContainer := newAdaptiveSplit(currentTestCaseModelAreaUI, currentTestCaseAttributesAreaUI)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(testCaseToolUIBar, nil, nil, nil)
	completeCurrentTestCaseUIContainer = container.New(testCaseBorderedLayout, testCaseToolUIBar, testCaseAdaptiveSplitLayoutContainer)

	return completeCurrentTestCaseUIContainer
}

// Loads current TestCase model and return the UI-structure for it
func (uiServer *UIServerStruct) loadCurrentTestCaseModelAreaUI() (currentTestCaseModelAreaUI fyne.CanvasObject) {

	currentTestCaseModelAreaUI = widget.NewLabel("'currentTestCaseModelAreaUI'")

	return currentTestCaseModelAreaUI
}

// Loads current TestCase attributes and return the UI-structure for it
func (uiServer *UIServerStruct) loadCurrentTestCaseAttributesAreaUI() (currentTestCaseAttributesAreaUI fyne.CanvasObject) {

	currentTestCaseAttributesAreaUI = widget.NewLabel("'currentTestCaseAttributesAreaUI'")

	return currentTestCaseAttributesAreaUI
}
