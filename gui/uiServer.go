package gui

import (
	"FenixTesterGui/commandAndRuleEngine"
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	executionsModelForExecutions "FenixTesterGui/executions/executionsModelForTestCaseExecutions"
	"FenixTesterGui/executions/executionsUIForExecutions"
	"FenixTesterGui/executions/executionsUIForSubscriptions"
	"FenixTesterGui/grpc_out_GuiTestCaseBuilderServer"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCase/testCaseUI"
	"FenixTesterGui/testCaseSubscriptionHandler"
	"FenixTesterGui/testCases/listTestCasesModel"
	"FenixTesterGui/testCases/listTestCasesUI"
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
	"github.com/sirupsen/logrus"
	"image/color"
	"log"
	"os"
	"strconv"
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
	fenixApp      fyne.App
	logger       *logrus.Logger
}


*/
//var myTestCase *testCaseUIStruct

var image *canvas.Image

func (globalUISServer *GlobalUIServerStruct) StartUIServer() {

	uiServer := &UIServerStruct{
		logger:                             nil,
		fenixApp:                           nil,
		fenixMasterWindow:                  nil,
		tree:                               nil,
		content:                            nil,
		fenixGuiBuilderServerAddressToDial: "",
		AvailableBuildingBlocksModel: AvailableBuildingBlocksModelStruct{
			logger:                             nil,
			fenixGuiBuilderServerAddressToDial: "",
			fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
			fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
			AvailableBuildingBlocksForUITreeNodes:                                      nil,
			grpcOut:                                                                    grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		},
		testCasesModel: testCaseModel.TestCasesModelsStruct{
			TestCases:        nil,
			CurrentUser:      sharedCode.CurrentUserIdLogedInOnComputer,
			GrpcOutReference: nil,
		},
		commandAndRuleEngine: commandAndRuleEngine.CommandAndRuleEngineObjectStruct{},
		grpcOut:              grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{},
		testCasesUiModel: testCaseUI.TestCasesUiModelStruct{
			TestCaseToolUIBar:       nil,
			TestCasesTabs:           nil,
			TestCasesUiModelMap:     nil,
			TestCasesModelReference: nil,
		},
	}
	// Add gRPC-out Reference
	uiServer.commandAndRuleEngine.GrpcOutReference = &uiServer.AvailableBuildingBlocksModel.grpcOut
	uiServer.testCasesModel.GrpcOutReference = &uiServer.AvailableBuildingBlocksModel.grpcOut

	// Add TestCasesReference to CommandEngine
	uiServer.commandAndRuleEngine.Testcases = &uiServer.testCasesModel

	// Add TestCasesReference to TestUI-engine
	uiServer.testCasesUiModel.TestCasesModelReference = &uiServer.testCasesModel

	// Add CommandEngineReference to TestUI-engine
	uiServer.testCasesUiModel.CommandAndRuleEngineReference = &uiServer.commandAndRuleEngine

	// Forward logger and Dail string
	uiServer.SetLogger(globalUISServer.logger)
	uiServer.SetDialAddressString(globalUISServer.fenixGuiBuilderServerAddressToDial)

	// Create Channel used for sending Commands to CommandsEngine
	sharedCode.CommandChannel = make(chan sharedCode.ChannelCommandStruct)
	myCommandChannelRef := &sharedCode.CommandChannel
	uiServer.testCasesUiModel.CommandAndRuleEngineReference.CommandChannelReference = myCommandChannelRef
	uiServer.testCasesUiModel.CommandChannelReference = myCommandChannelRef

	// Start Receiver channel for Commands
	uiServer.commandAndRuleEngine.InitiateCommandChannelReader()

	// Start Channel used for updating status on TestCaseExecutions
	detailedExecutionsModel.InitiateCommandChannelReaderForDetailedStatusUpdates()

	// Create Channel used for triggering TestCase Graphics update
	sharedCode.CommandChannelGraphicsUpdate = make(chan sharedCode.ChannelCommandGraphicsUpdatedStruct)
	myGraphicsUpdateChannelRef := &sharedCode.CommandChannelGraphicsUpdate
	uiServer.testCasesUiModel.CommandAndRuleEngineReference.GraphicsUpdateChannelReference = myGraphicsUpdateChannelRef
	uiServer.testCasesUiModel.GraphicsUpdateChannelReference = myGraphicsUpdateChannelRef

	// Start Receiver channel for Graphics Update
	uiServer.testCasesUiModel.InitiateGraphicsUpdateChannelReader()

	uiServer.startTestCaseUIServer()

}

// Main UI server module
func (uiServer *UIServerStruct) startTestCaseUIServer() {
	/*
		myUIServer = UIServerStruct{
			logger:  callersLoggerReference,
			grpcOut: grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct{Logger: callersLoggerReference},
		}

	*/
	uiServer.logger.WithFields(logrus.Fields{
		"id": "a4d2716f-ded1-4062-bffb-fd0c03d69ca3",
	}).Debug("Starting UI server")

	// var err error
	/*
		myTestCase = &testCaseUIStruct{
			logger: myUIServer.logger,
		}

	*/
	//myUIServer = UIServerStruct{}

	//var grpcOut grpc_out_GuiTestCaseBuilderServer.GRPCOutGuiTestCaseBuilderServerStruct
	// myUIServer.grpcOut.SetLogger(myUIServer.logger)

	// Add/Forward variables to packages to be used later
	uiServer.AvailableBuildingBlocksModel.SetLogger(uiServer.logger)
	uiServer.commandAndRuleEngine.SetLogger(uiServer.logger)
	uiServer.AvailableBuildingBlocksModel.grpcOut.SetLogger(uiServer.logger)
	uiServer.AvailableBuildingBlocksModel.grpcOut.SetDialAddressString(uiServer.fenixGuiBuilderServerAddressToDial)

	uiServer.fenixApp = app.NewWithID("se.fenix.testcasebuilder")
	//fenixApp.Settings().SetTheme(&myTheme{})
	uiServer.fenixMasterWindow = uiServer.fenixApp.NewWindow("Fenix TestCase Builder")
	uiServer.fenixMasterWindow.SetMaster()
	uiServer.fenixMasterWindow.CenterOnScreen()

	/*
		var w fyne.Window
		if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
			w = drv.CreateSplashWindow()
			w.SetContent(widget.NewLabel("\"If you want to change the world, don't protest. Write code!\" - Hal Finney (1994)"))
			w.Show()

			go func() {
				time.Sleep(time.Second * 5)
				w.Close()

			}()
		}

	*/
	// Create Fenix Splash screen
	var splashWindow fyne.Window
	var splashWindowProlongedVisibleChannel chan time.Duration
	splashWindowProlongedVisibleChannel = make(chan time.Duration)
	createSplashWindow(&splashWindow, &splashWindowProlongedVisibleChannel)

	uiServer.commandAndRuleEngine.MasterFenixWindow = &uiServer.fenixMasterWindow

	// Get Available Building BLocks form GUI-server
	uiServer.AvailableBuildingBlocksModel.loadAvailableBuildingBlocksFromServer(&uiServer.testCasesModel)

	// Get Available Building Blocks form GUI-server
	uiServer.AvailableBuildingBlocksModel.loadPinnedBuildingBlocksFromServer()

	// Load list with TemplateRepositoryApiUrls form GUI-server
	uiServer.AvailableBuildingBlocksModel.loadTemplateRepositoryApiUrls(&uiServer.testCasesModel)

	// Load available TestData for the User
	uiServer.AvailableBuildingBlocksModel.loadTestData(&uiServer.testCasesModel)

	// Load list with TestCases that the user can edit
	listTestCasesModel.LoadTestCaseThatCanBeEditedByUser(&uiServer.testCasesModel)

	// Load Available Bonds
	uiServer.commandAndRuleEngine.LoadAvailableBondsFromServer()

	// Load Immature TestInstruction Attributes into TestCase-model
	// Not used directly and is instead moved into when getting TI and TIC from BuilderServer
	//uiServer.testCasesModel.LoadModelWithImmatureTestInstructionAttributes()

	// Create the Available Building Blocks adapted to Fyne tree-view
	uiServer.AvailableBuildingBlocksModel.makeTreeUIModel()

	// Initiate all variables needed by the TestCaseExecution-SubscriptionHandler
	testCaseSubscriptionHandler.TestCaseExecutionStatusSubscriptionHandlerObject.InitiateTestCaseExecutionStatusSubscriptionHandler()

	// Initiate the channels used when Adding or Removing items to/from OnQueue-table, UnderExecution-table or FinishedExecutions-table
	executionsModelForSubscriptions.InitiateAndStartChannelsUsedByListModel()

	// Start Channel readers for testCases OnQueue, UnderExecutions or Finished Executions
	executionsUIForSubscriptions.StartTableAddAndRemoveChannelReaders()

	// Initiate Models for Subscription regarding TestCaseExecutionsOnExecutionQueue
	executionsModelForSubscriptions.InitiateSubscriptionModelForTestCaseOnExecutionQueue()

	// Load TestCaseExecutionsOnExecutionQueue
	var domainsList []string
	domainsList = nil
	_ = executionsModelForExecutions.ExecutionsModelObject.LoadAndCreateModelForTestCasesOnExecutionQueue(domainsList)

	// Initiate Models for Subscription regarding UnderExecution
	executionsModelForSubscriptions.InitiateSubscriptionModelForTestCaseUnderExecution()

	// Load TestCaseExecutionsUnderExecution
	domainsList = nil
	_ = executionsModelForExecutions.ExecutionsModelObject.LoadAndCreateModelForTestCaseUnderExecutions(domainsList)

	// Initiate Models for Subscription regarding FinishedExecutions
	executionsModelForSubscriptions.InitiateSubscriptionModelForTestCaseWithFinishedExecutions()

	// Load TestCaseExecutionsWithFinishedExecutions
	domainsList = nil
	_ = executionsModelForExecutions.ExecutionsModelObject.LoadAndCreateModelForTestCaseWithFinishedExecutions(domainsList)

	// Initiate and create the tree structure for available building blocks, of TestInstructions and TestInstructionContainers
	uiServer.makeTreeUI()
	tree.OpenAllBranches()

	// Initiate the commandStack which describes how fenixApp TestCase is constructed
	uiServer.makeCommandStackUI()

	// Create fenixApp window for the Command Stack
	//TODO Remove StackWindow
	//commandStackWindow := uiServer.fenixApp.NewWindow("Command Stack")
	//commandStackWindow.SetContent(commandStackListUI)
	//commandStackWindow.Show()

	//list := &notelist{pref: fenixApp.Preferences()}
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

	myCanvas := uiServer.fenixMasterWindow.Canvas()

	//blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	//rect := canvas.NewRectangle(blue)

	//var rect *customRectangle
	//rect = rect.newCustomRect()
	//myCanvasLabel := widget.NewLabel("My Canvas Overlay Label")

	// Create the UI for Build TestCase-UI
	var buildTestCasesUI fyne.CanvasObject
	buildTestCasesUI = uiServer.loadUI()

	// Create the UI for List TestCases-UI
	var tempListTestCasesUI fyne.CanvasObject
	tempListTestCasesUI = listTestCasesUI.GenerateListTestCasesUI(&uiServer.testCasesModel)

	mySliderDataAsString := binding.NewString()

	uiServer.fenixApp.Settings().Scale()
	scaleEnvKey := "FYNE_SCALE"
	envVal := os.Getenv(scaleEnvKey)
	fmt.Println(envVal)
	//defer os.Setenv(scaleEnvKey, envVal)

	//	_ = os.Setenv(scaleEnvKey, "auto")

	sizeSlider := widget.NewSlider(40, 200)
	sizeSliderSizeLabel := widget.NewLabelWithData(mySliderDataAsString)
	sizeContainer := container.NewVBox(sizeSliderSizeLabel, sizeSlider)
	sizeSlider.Resize(fyne.NewSize(300, 0))

	configContainerGrid := container.New(layout.NewAdaptiveGridLayout(2), sizeContainer, widget.NewLabel("Test"))

	// Generate 'left' Execution Tab for Subscriptions, that holds listings for Executions and individual detailed Executions
	subscriptionExecutionsUITab := executionsUIForSubscriptions.ExecutionsUIObject.GenerateBaseUITabForExecutions() //MySortTable() //CreateTableObject()

	// Generate 'left' Execution Tab, that holds listings for Executions and individual detailed Executions
	executionsUITab := executionsUIForExecutions.ExecutionsUIObject.GenerateBaseUITabForExecutions() //MySortTable() //CreateTableObject()

	// Generate a test tab for Detailed TestCaseExecutions
	detailedTestCaseExecutionTab := detailedTestCaseExecutionsUI.DetailedTestCaseExecutionsUIObject.GenerateBaseUITabForDetailedTestCaseExecutions()

	tabs := container.NewAppTabs(
		container.NewTabItem("Build TestCase", buildTestCasesUI),
		container.NewTabItem("List TestCases", tempListTestCasesUI),
		container.NewTabItem("Executions (Subscriptions)", subscriptionExecutionsUITab),
		container.NewTabItem("Executions", executionsUITab),
		container.NewTabItem("Detailed TestCaseExecutions", detailedTestCaseExecutionTab),
		container.NewTabItem("Config", configContainerGrid),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.OnSelected = func(tabItem *container.TabItem) {
		// TODO UGLY Update of page
		executionsUIForSubscriptions.ExecutionsUIObject.OnQueueTable.Header.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.OnQueueTable.Header.Refresh()
		executionsUIForSubscriptions.ExecutionsUIObject.UnderExecutionTable.Header.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.UnderExecutionTable.Header.Refresh()
		executionsUIForSubscriptions.ExecutionsUIObject.FinishedExecutionTable.Header.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.FinishedExecutionTable.Header.Refresh()

		executionsUIForSubscriptions.ExecutionsUIObject.OnQueueTable.Data.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.OnQueueTable.Data.Refresh()
		executionsUIForSubscriptions.ExecutionsUIObject.UnderExecutionTable.Data.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.UnderExecutionTable.Data.Refresh()
		executionsUIForSubscriptions.ExecutionsUIObject.FinishedExecutionTable.Data.ScrollToLeading()
		executionsUIForSubscriptions.ExecutionsUIObject.FinishedExecutionTable.Data.Refresh()
	}

	tabs.SetTabLocation(container.TabLocationLeading)

	sizeSlider.OnChanged = func(f float64) {

		err := os.Setenv(scaleEnvKey, "1.5")
		uiServer.fenixMasterWindow.Hide()
		uiServer.fenixMasterWindow.Show()
		fmt.Println("err: ", err)
		//		_ = os.Setenv(scaleEnvKey, s)
		set := uiServer.fenixApp.Settings().Scale()
		fmt.Println(set)
		set = fyne.CurrentApp().Settings().Scale()
		fmt.Println(set)

		myInt := int(f)
		myString := strconv.Itoa(myInt)
		s := "Screen zoomm: " + myString + "%"

		mySliderDataAsString.Set(s)

		_ = fyne.CurrentApp().Settings().Scale()

	}

	myCanvas.SetContent(tabs)
	_ = os.Setenv(scaleEnvKey, sharedCode.FYNE_SCALE)

	//myCanvas.Overlays().Add(myCanvasLabel)

	//uiServer.fenixMasterWindow.SetContent(myCanvas) //(uiServer.loadUI())

	//uiServer.fenixMasterWindow.SetContent(widget.NewLabel("Fenix TestCase Builder"))
	//builderUI.registerKeys(uiServer.fenixMasterWindow)

	uiServer.fenixMasterWindow.Resize(fyne.NewSize(3000, 1500))

	//w.Hide()

	// Store pointers to FenixApp and FenixMainWindow
	sharedCode.FenixAppPtr = &uiServer.fenixApp
	sharedCode.FenixMasterWindowPtr = &uiServer.fenixMasterWindow

	splashWindow.RequestFocus()
	splashWindow.Show()
	go func() {
		time.Sleep(time.Millisecond * 500)
		splashWindow.RequestFocus()
		splashWindowProlongedVisibleChannel <- time.Second * 6
	}()
	uiServer.fenixMasterWindow.RequestFocus()
	uiServer.fenixMasterWindow.ShowAndRun()

}

type customRectangle struct {
	widget.Label
	myrect canvas.Rectangle
}

func (c *customRectangle) newCustomRect() *customRectangle {
	myRectangle := &customRectangle{}
	c.ExtendBaseWidget(myRectangle)
	//blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	//myRectangle.FillColor = blue
	//myRectangle.StrokeColor = color.Black
	//myRectangle.StrokeWidth = 10

	return myRectangle
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *customRectangle) MouseMoved(a *desktop.MouseEvent) {
	log.Println("I have been 'MouseMoved'")
	fmt.Println(a.Position, a.AbsolutePosition)
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

	// Old original solution for creating TestCase-UI:::: testCaseSide := uiServer.loadCompleteCurrentTestCaseUI()
	testCaseSide := uiServer.testCasesUiModel.GenerateBaseCanvasObjectForTestCaseUI()

	uiStructureContainer := newAdaptiveSplit(treeSide, testCaseSide)

	// Create Object used when Dragging TI and TIC from Available Building Blocks
	// TODO REALLY UGLY CODE AND SHOULD BE BROKEN OUT INTO FUNCTION
	// Add Text to be used for Drag n Drop, for now it's for testing only
	colorBlack := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	draggingText := canvas.NewText("Drag n Drop Object", colorBlack)
	draggingBackgroundRectangle := canvas.NewRectangle(color.RGBA{
		R: 0x55,
		G: 0x55,
		B: 0x55,
		A: 0x22,
	})
	draggingBackgroundRectangle.StrokeWidth = 2
	draggingBackgroundRectangle.StrokeColor = color.RGBA{
		R: 0x88,
		G: 0x88,
		B: 0x88,
		A: 0x99,
	}

	draggingTextBackgroundRectangle := canvas.NewRectangle(color.RGBA{
		R: 0x55,
		G: 0x55,
		B: 0x55,
		A: 0xaa,
	})

	draggingBackgroundRectangle.SetMinSize(draggingText.Size().Add(fyne.NewSize(50, 50)))
	contentGroupDragginObject := container.NewCenter(draggingBackgroundRectangle, draggingTextBackgroundRectangle, draggingText)

	contentGroupDragginObject.Move(fyne.NewPos(320, 320))
	contentDraggingObject := container.NewWithoutLayout(contentGroupDragginObject)

	uiServer.testCasesUiModel.DragNDropObject.DragNDropText = draggingText
	uiServer.testCasesUiModel.DragNDropObject.DragNDropRectangle = draggingBackgroundRectangle
	uiServer.testCasesUiModel.DragNDropObject.DragNDropRectangleTextBackground = draggingTextBackgroundRectangle
	uiServer.testCasesUiModel.DragNDropObject.DragNDropContainer = contentGroupDragginObject
	uiServer.testCasesUiModel.DragNDropObject.DragNDropContainer.Hide()

	uiServer.testCasesUiModel.DragNDropStateMachine.InitiateStateStateMachine(
		uiServer.testCasesUiModel.DragNDropObject.DragNDropText,
		uiServer.testCasesUiModel.DragNDropObject.DragNDropRectangle,
		uiServer.testCasesUiModel.DragNDropObject.DragNDropRectangleTextBackground,
		uiServer.testCasesUiModel.DragNDropObject.DragNDropContainer,
		uiServer.commandAndRuleEngine.CommandChannelReference,
		uiServer.commandAndRuleEngine.Testcases)

	// ****************************
	// TODO Used for Testing only and can be removed
	// Add Text to be used for Drag n Drop, for now it's for testing only
	black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	text2 := canvas.NewText("There...", black)
	backgroundRect := canvas.NewRectangle(color.RGBA{
		R: 0x55,
		G: 0x55,
		B: 0x55,
		A: 0x22,
	})
	backgroundRect.StrokeWidth = 2
	backgroundRect.StrokeColor = color.RGBA{
		R: 0x88,
		G: 0x88,
		B: 0x88,
		A: 0x99,
	}

	middlebackgroundRect := canvas.NewRectangle(color.RGBA{
		R: 0x55,
		G: 0x55,
		B: 0x55,
		A: 0xaa,
	})

	backgroundRect.SetMinSize(text2.Size().Add(fyne.NewSize(50, 50)))
	contentGroup := container.NewCenter(backgroundRect, middlebackgroundRect, text2)

	contentGroup.Move(fyne.NewPos(120, 120))
	content := container.NewWithoutLayout(contentGroup)

	uiServer.testCasesUiModel.DragNDropText = text2
	uiServer.testCasesUiModel.DragNDropRectangle = backgroundRect
	uiServer.testCasesUiModel.DragNDropRectangleTextBackground = middlebackgroundRect
	uiServer.testCasesUiModel.DragNDropContainer = contentGroup
	uiServer.testCasesUiModel.DragNDropContainer.Hide()

	// ****************************

	myLoayout := container.NewMax(uiStructureContainer, content, contentDraggingObject)

	return myLoayout
}

// Loads available TestInstructions and TestInstructionContainers and return the UI Bar and UI Tree-structure for them
func (uiServer *UIServerStruct) loadCompleteAvailableTestCaseBuildingBlocksUI() (completeAvailableTestCaseBuildingBlocksUI fyne.CanvasObject) {

	// Create toolbar for Available TestCase BuildingBlockType area
	availableAvailableBuildingBlocksUIBar := widget.NewToolbar(

		// Icon for reloading Building Blocks from Server
		widget.NewToolbarAction(theme.ContentRedoIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")

			// Load Available Building Blocks and Pinned Building Blocks from Server
			uiServer.AvailableBuildingBlocksModel.loadAvailableBuildingBlocksFromServer((&uiServer.testCasesModel))
			uiServer.AvailableBuildingBlocksModel.loadPinnedBuildingBlocksFromServer()

			// Load list with TemplateRepositoryApiUrls form GUI-server
			uiServer.AvailableBuildingBlocksModel.loadTemplateRepositoryApiUrls(&uiServer.testCasesModel)

			// Load available TestData for the User
			uiServer.AvailableBuildingBlocksModel.loadTestData(&uiServer.testCasesModel)

			// Recreate the TreeUIModel
			uiServer.AvailableBuildingBlocksModel.makeTreeUIModel()

			// Recreate the UI-tree-component
			uiServer.makeTreeUI()
		}),

		// Icon for saving pinned Building Blocks to Server
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			fmt.Println("Reload Available Components from GuiServer")

			// Load Available Building Blocks and Pinned Building Blocks from Server
			_ = uiServer.AvailableBuildingBlocksModel.savePinnedBuildingBlocksFromServer()
			/*
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
			*/
		}),

		// Icon for Adding Building Block to Pinned Building Blocks
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			err := uiServer.AvailableBuildingBlocksModel.verifyBeforePinTestInstructionOrTestInstructionContainer(uiServer.AvailableBuildingBlocksModel.clickedNodeName, true)
			if err == nil {
				fmt.Println("Add to Pinned")
				err := uiServer.AvailableBuildingBlocksModel.pinTestInstructionOrTestInstructionContainer(uiServer.AvailableBuildingBlocksModel.clickedNodeName)
				if err == nil {
					// Update the testCaseModel, which will refresh UI
					uiServer.AvailableBuildingBlocksModel.makeTreeUIModel()
				}
			}

		}),

		// Icon for Removing Pinned Building Block
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			err := uiServer.AvailableBuildingBlocksModel.verifyBeforeUnPinTestInstructionOrTestInstructionContainer(uiServer.AvailableBuildingBlocksModel.clickedNodeName, true)
			if err == nil {
				fmt.Println("Remove from Pinned")
				err := uiServer.AvailableBuildingBlocksModel.unPinTestInstructionOrTestInstructionContainer(uiServer.AvailableBuildingBlocksModel.clickedNodeName)
				if err == nil {
					// Update the testCaseModel, which will refrsh UI
					uiServer.AvailableBuildingBlocksModel.makeTreeUIModel()
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

	buildingBlockSplit := container.NewVSplit(availableTestCaseBuildingBlocksTreeUI, uiServer.createTestCaseCommandsUI())
	buildingBlockSplit.Offset = 0.8

	tempcompleteAvailableTestCaseBuildingBlocksUI := container.New(availableTestCaseBuildingBlocksBorderedLayout, availableAvailableBuildingBlocksUIBar, buildingBlockSplit)
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
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

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
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureSimple = binding.NewString()
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureSimple.Set("'currentTestCaseTextualStructureSimple'")

	// Set initial values for TestCase Textual Structure - Complex
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureComplex = binding.NewString()
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureComplex.Set("'currentTestCaseTextualStructureComplex'")

	// Set initial values for TestCase Textual Structure - Simple
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureExtended = binding.NewString()
	uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureExtended.Set("'currentTestCaseTextualStructureExtended'")

	// Create the Labels to be used for showing the TestCase Textual Structures
	testCaseTextualStructureSimpleWidget := widget.NewLabelWithData(uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureSimple)
	testCaseTextualStructureComplexWidget := widget.NewLabelWithData(uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureComplex)
	testCaseTextualStructureExtendedWidget := widget.NewLabelWithData(uiServer.AvailableBuildingBlocksModel.currentTestCaseTextualStructureExtended)

	// Create GUI Canvas object to be used
	currentTestCaseModelAreaUI = container.NewVBox(testCaseTextualStructureSimpleWidget, testCaseTextualStructureComplexWidget, testCaseTextualStructureExtendedWidget)

	return currentTestCaseModelAreaUI
}

// Loads current TestCase attributes and return the UI-structure for it
func (uiServer *UIServerStruct) loadCurrentTestCaseAttributesAreaUI() (currentTestCaseAttributesAreaUI fyne.CanvasObject) {

	currentTestCaseAttributesAreaUI = widget.NewLabel("'currentTestCaseAttributesAreaUI'")

	return currentTestCaseAttributesAreaUI
}
