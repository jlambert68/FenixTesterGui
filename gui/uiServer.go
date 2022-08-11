package gui

import (
	"FenixTesterGui/commandAndRuleEngine"
	"FenixTesterGui/grpc_out"
	"FenixTesterGui/testCase/testCaseModel"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
	//"FenixTesterGui/resources"
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

func (globalUISServer *GlobalUIServerStruct) StartUIServer() {

	uiServer := &UIServerStruct{
		logger:                             nil,
		fyneApp:                            nil,
		tree:                               nil,
		content:                            nil,
		fenixGuiBuilderServerAddressToDial: "",
		availableBuildingBlocksModel: AvailableBuildingBlocksModelStruct{
			logger:                             nil,
			fenixGuiBuilderServerAddressToDial: "",
			fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
			fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
			availableBuildingBlocksForUITreeNodes:                                      nil,
			grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		},
		testCasesModel: testCaseModel.TestCasesModelsStruct{
			TestCases:        nil,
			CurrentUser:      "s41797",
			GrpcOutReference: nil,
		},
		commandAndRuleEngine: commandAndRuleEngine.CommandAndRuleEngineObjectStruct{},
		grpcOut:              grpc_out.GRPCOutStruct{},
	}
	// Add gRPC-out Reference
	uiServer.commandAndRuleEngine.GrpcOutReference = &uiServer.availableBuildingBlocksModel.grpcOut
	uiServer.testCasesModel.GrpcOutReference = &uiServer.availableBuildingBlocksModel.grpcOut

	// Add TestCasesReference to CommandEngine
	uiServer.commandAndRuleEngine.Testcases = &uiServer.testCasesModel

	// Forward logger and Dail string
	uiServer.SetLogger(globalUISServer.logger)
	uiServer.SetDialAddressString(globalUISServer.fenixGuiBuilderServerAddressToDial)

	uiServer.startTestCaseUIServer()

}

// Main UI server module
func (uiServer *UIServerStruct) startTestCaseUIServer() {
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
	uiServer.availableBuildingBlocksModel.SetLogger(uiServer.logger)
	uiServer.commandAndRuleEngine.SetLogger(uiServer.logger)
	uiServer.availableBuildingBlocksModel.grpcOut.SetLogger(uiServer.logger)
	uiServer.availableBuildingBlocksModel.grpcOut.SetDialAddressString(uiServer.fenixGuiBuilderServerAddressToDial)

	uiServer.fyneApp = app.NewWithID("se.fenix.testcasebuilder")
	//fyneApp.Settings().SetTheme(&myTheme{})
	fyneMasterWindow := uiServer.fyneApp.NewWindow("Fenix TestCase Builder")
	fyneMasterWindow.SetMaster()

	// Get Available Building BLocks form GUI-server
	uiServer.availableBuildingBlocksModel.loadAvailableBuildingBlocksFromServer()

	// Get Available Building BLocks form GUI-server
	uiServer.availableBuildingBlocksModel.loadPinnedBuildingBlocksFromServer()

	// Load Available Bonds
	uiServer.commandAndRuleEngine.LoadAvailableBondsFromServer()

	// Create the Available Building Blocks adapted to Fyne tree-view
	uiServer.availableBuildingBlocksModel.makeTreeUIModel()

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

	fyneMasterWindow.Resize(fyne.NewSize(3000, 1500))

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

		// Icon for reloading Building Blocks from Server
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")

			// Load Available Building Blocks and Pinned Building Blocks from Server
			uiServer.availableBuildingBlocksModel.loadAvailableBuildingBlocksFromServer()
			uiServer.availableBuildingBlocksModel.loadPinnedBuildingBlocksFromServer()

			// Recreate the TreeUIModel
			uiServer.availableBuildingBlocksModel.makeTreeUIModel()

			// Recreate the UI-tree-component
			uiServer.makeTreeUI()
		}),

		// Icon for saving pinned Building Blocks to Server
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")

			// Load Available Building Blocks and Pinned Building Blocks from Server
			err := uiServer.availableBuildingBlocksModel.savePinnedBuildingBlocksFromServer()

			f, err := os.Open("resources/s_ui_error_stereo_04-35938.mp3")
			if err != nil {
				log.Println(err)
			}

			streamer, format, err := mp3.Decode(f)
			if err != nil {
				log.Println(err)
			}
			defer streamer.Close()

			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

			done := make(chan bool)
			speaker.Play(beep.Seq(streamer, beep.Callback(func() {
				done <- true
			})))

			<-done

		}),

		// Icon for Adding Building Block to Pinned Building Blocks
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			err := uiServer.availableBuildingBlocksModel.verifyBeforePinTestInstructionOrTestInstructionContainer(uiServer.availableBuildingBlocksModel.clickedNodeName, true)
			if err == nil {
				fmt.Println("Add to Pinned")
				err := uiServer.availableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer(uiServer.availableBuildingBlocksModel.clickedNodeName)
				if err == nil {
					// Update the testCaseModel, which will refrsh UI
					uiServer.availableBuildingBlocksModel.makeTreeUIModel()
				}
			}

		}),

		// Icon for Removing Pinned Building Block
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			err := uiServer.availableBuildingBlocksModel.verifyBeforeUnPinTestInstructionOrTestInstructionContainer(uiServer.availableBuildingBlocksModel.clickedNodeName, true)
			if err == nil {
				fmt.Println("Remove from Pinned")
				err := uiServer.availableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer(uiServer.availableBuildingBlocksModel.clickedNodeName)
				if err == nil {
					// Update the testCaseModel, which will refrsh UI
					uiServer.availableBuildingBlocksModel.makeTreeUIModel()
					//uiServer.tree.Refresh()
				}
			}
		}),
	)

	// Load the Available TestCase BuildingBlocks TreeUI
	availableTestCaseBuildingBlocksTreeUI := uiServer.loadAvailableTestCaseBuildingBlocksTreeUI()

	//	commandParametersAndCommandLayout := container.New(layout.NewVBoxLayout(), uiServer.createTestCaseCommandParametersUI(), uiServer.createTestCaseCommandsUI())
	//commandParametersAndCommandLayout := container.New(layout.NewVBoxLayout(), uiServer.createTestCaseCommandParametersUI(), uiServer.createTestCaseCommandsUI())
	//commandParametersAndCommandLayout := container.New(layout.NewVBoxLayout(), uiServer.createTestCaseCommandsUI())

	// Create the complete TestCase BuildingBlocks UI area
	availableTestCaseBuildingBlocksBorderedLayout := layout.NewBorderLayout(availableAvailableBuildingBlocksUIBar, nil, nil, nil)
	tempcompleteAvailableTestCaseBuildingBlocksUI := container.New(availableTestCaseBuildingBlocksBorderedLayout, availableAvailableBuildingBlocksUIBar, container.NewVSplit(availableTestCaseBuildingBlocksTreeUI, uiServer.createTestCaseCommandsUI()))
	//tempcompleteAvailableTestCaseBuildingBlocksUI.MinSize(fyne.NewSize(float32(300), float32(400))

	//templabel := widget.NewLabel("MyLabel")

	//newAll := container.NewWithoutLayout(tempcompleteAvailableTestCaseBuildingBlocksUI, templabel)

	//templabel.Move(fyne.NewPos(200, 200))

	//completeAvailableTestCaseBuildingBlocksUI = newAll
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
			fmt.Println("Reload GUI TestCase from testCaseModel")
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

	// Load the TestCase testCaseModel UI area
	currentTestCaseModelAreaUI := uiServer.loadCurrentTestCaseModelAreaUI()

	// Load the TestCase attributes UI area
	currentTestCaseAttributesAreaUI := uiServer.loadCurrentTestCaseAttributesAreaUI()

	// Create the UI area for both TestCase testCaseModel UI and TestCase attributes UI
	testCaseAdaptiveSplitLayoutContainer := newAdaptiveSplit(currentTestCaseModelAreaUI, currentTestCaseAttributesAreaUI)

	// Create the complete TestCase UI area
	testCaseBorderedLayout := layout.NewBorderLayout(testCaseToolUIBar, nil, nil, nil)
	completeCurrentTestCaseUIContainer = container.New(testCaseBorderedLayout, testCaseToolUIBar, testCaseAdaptiveSplitLayoutContainer)

	return completeCurrentTestCaseUIContainer
}

// Loads current TestCase testCaseModel and return the UI-structure for it
func (uiServer *UIServerStruct) loadCurrentTestCaseModelAreaUI() (currentTestCaseModelAreaUI fyne.CanvasObject) {

	// Set initial values for TestCase Textual Structure - Simple
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureSimple = binding.NewString()
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureSimple.Set("'currentTestCaseTextualStructureSimple'")

	// Set initial values for TestCase Textual Structure - Complex
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureComplex = binding.NewString()
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureComplex.Set("'currentTestCaseTextualStructureComplex'")

	// Set initial values for TestCase Textual Structure - Simple
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureExtended = binding.NewString()
	uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureExtended.Set("'currentTestCaseTextualStructureExtended'")

	// Create the Labels to be used for showing the TestCase Textual Structures
	testCaseTextualStructureSimpleWidget := widget.NewLabelWithData(uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureSimple)
	testCaseTextualStructureComplexWidget := widget.NewLabelWithData(uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureComplex)
	testCaseTextualStructureExtendedWidget := widget.NewLabelWithData(uiServer.availableBuildingBlocksModel.currentTestCaseTextualStructureExtended)

	// Create GUI Canvas object to be used
	currentTestCaseModelAreaUI = container.NewVBox(testCaseTextualStructureSimpleWidget, testCaseTextualStructureComplexWidget, testCaseTextualStructureExtendedWidget)

	return currentTestCaseModelAreaUI
}

// Loads current TestCase attributes and return the UI-structure for it
func (uiServer *UIServerStruct) loadCurrentTestCaseAttributesAreaUI() (currentTestCaseAttributesAreaUI fyne.CanvasObject) {

	currentTestCaseAttributesAreaUI = widget.NewLabel("'currentTestCaseAttributesAreaUI'")

	return currentTestCaseAttributesAreaUI
}
