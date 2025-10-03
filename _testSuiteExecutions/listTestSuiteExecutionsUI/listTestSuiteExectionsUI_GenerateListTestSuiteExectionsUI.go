package listTestSuiteExecutionsUI

import (
	detailedTestSuiteExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testSuiteExecutions/listTestSuiteExecutionsModel"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"strconv"
)

//go:embed resources/TIC-Horizontal_32x32.png
var tic_parallellImage []byte
var imageData_tic_parallellImage image.Image

//go:embed resources/TIC-Vertical_32x32.png
var tic_serialImage []byte
var imageData_tic_serialImage image.Image

//go:embed resources/sort_cropped_32x51.png
var sortUnspecifiedImageAsByteArray []byte
var sortImageUnspecifiedAsImage image.Image

//go:embed resources/sort-down-descending_cropped_32x51.png
var sortImageAscendingAsByteArray []byte
var sortImageAscendingAsImage image.Image

//go:embed resources/sort-up-ascending_cropped_32x51.png
var sortImageDescendingAsByteArray []byte
var sortImageDescendingAsImage image.Image

// Define the function to be executed to load TestSuiteExecutions from that Database that the user can view
// Only loads one TestSuiteExecution per TestCase
func LoadOneTestSuiteExecutionPerTestCaseFromDataBaseFunction(
	testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	updateGui bool,
	testSuiteInstructionPreViewObject *TestSuiteInstructionPreViewStruct) { // ***

	// If previously data set is of other kind then 'OneExecutionPerTestSuite' then no selection should be made
	if selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType != OneExecutionPerTestSuite {
		//selectedTestSuiteExecutionObjected.isAnyRowSelected = false
		testSuiteInstructionPreViewObject.ClearTestSuiteExecutionPreviewContainer() // ***
		loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Disable()

	}

	// Get number of TestCase retrieved from Database before and after loading more data
	//var beforeNumberOfRowsRetrievedFromDatabase int
	//var afterNumberOfRowsRetrievedFromDatabase int
	//beforeNumberOfRowsRetrievedFromDatabase = testSuiteExecutionsModel.
	//	GetNumberOfTestSuiteExecutionsRetrievedFromDatabase()

	// Specify from were Executions in the GUI comes from
	selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType = OneExecutionPerTestSuite

	// Channel used to trigger update of Gui
	var updateGuiChannelStep1 chan bool
	var updateGuiChannelStep2 chan bool
	updateGuiChannelStep1 = make(chan bool)
	updateGuiChannelStep2 = make(chan bool)

	listTestSuiteExecutionsModel.LoadTestSuiteExecutionsThatCanBeViewedByUser(
		testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.LatestUniqueTestSuiteExecutionDatabaseRowId,
		true,
		testSuiteExecutionsModel.StandardTestSuiteExecutionsBatchSize,
		false,
		"",
		testSuiteExecutionsModel.NullTimeStampForTestSuiteExecutionsSearch,
		testSuiteExecutionsModel.NullTimeStampForTestSuiteExecutionsSearch,
		true,
		&updateGuiChannelStep1)

	// Update the UI with the new data
	go func() {

		// Wait for trigger to update GUI
		<-updateGuiChannelStep1
		<-updateGuiChannelStep2

		// Update the GUI
		if updateGui == true {
			SortGuiTableOnCurrentColumnAndSorting()
		}

		/*
			// Get number of TestCase retrieved from Database after loading more data
			afterNumberOfRowsRetrievedFromDatabase = testSuiteExecutionsModel.
				GetNumberOfTestSuiteExecutionsRetrievedFromDatabase()

			// If they are different then sort the table
			if beforeNumberOfRowsRetrievedFromDatabase != afterNumberOfRowsRetrievedFromDatabase {
				//sortableHeaderReference.sortImage.onTapped()
				testSuiteExecutionsListTableHeadersMapRef[int(latestTestSuiteExecutionTimeStampColumnNumber)].sortImage.onTapped()

			}

		*/

	}()

	updateGuiChannelStep2 <- true
	//filterTestSuiteExcutionsButtonFunction()

}

// Create the UI used for list all TestCasesMapPtr that the User can edit
func GenerateListTestSuiteExecutionsUI(
	testSuiteExecutionsModel *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	detailedTestSuiteExecutionsUITabObject *container.DocTabs,
	exitFunctionsForDetailedTestSuiteExecutionsUITabObject *map[*container.TabItem]func(),
	testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct) ( // ***
	listTestCasesUI fyne.CanvasObject) {

	// save reference to 'detailedTestSuiteExecutionsUITabObject'
	detailedTestSuiteExecutionsUITabObjectRef = detailedTestSuiteExecutionsUITabObject

	// save Reference to Map over Exit-functions for the Tabs in 'detailedTestSuiteExecutionsUITabObject'
	exitFunctionsForDetailedTestSuiteExecutionsUITabObjectPtr = exitFunctionsForDetailedTestSuiteExecutionsUITabObject

	//var testCaseTable *widget.Table

	var tempTestCaseListAndTestCasePreviewSplitContainer *container.Split
	//var tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split

	var testSuiteExecutionsListContainer *fyne.Container
	var testSuiteExecutionsListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	// Button used for Retrieving all TestSuiteExecutions
	var loadTestSuiteExecutionsFromDataBaseButton *widget.Button
	var loadOneTestSuiteExecutionPerTestCaseFromDataBaseFunction func()

	var filterTestSuiteExcutionsButton *widget.Button
	var filterTestSuiteExcutionsButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var loadAllTestSuiteExecutionsForOneTestCaseButton *widget.Button
	var loadAllTestSuiteExecutionsForOneTestCaseButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestSuiteExecutionsAfterLocalFilterLabel *widget.Label
	var numberOfTestSuiteExecutionsRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Initiate map for Header-references
	testSuiteExecutionsListTableHeadersMapRef = make(map[int]*sortableHeaderLabelStruct)

	// Initiate 'selectedTestSuiteExecutionObjected' used for keep track of source of executions and what row is selected
	selectedTestSuiteExecutionObjected = selectedTestSuiteExecutionStruct{
		oneExecutionPerTestSuiteListObject: oneExecutionPerTestSuiteListObjectStruct{
			lastSelectedTestSuiteExecutionForOneExecutionPerTestSuite: "",
			testSuiteUuidForTestSuiteExecutionThatIsShownInPreview:    "",
			testSuiteExecutionUuidThatIsShownInPreview:                "",
			testSuiteExecutionVersionThatIsShownInPreview:             0,
			isAnyRowSelected:                false,
			currentSortColumn:               0,
			previousSortColumn:              0,
			currentHeader:                   nil,
			previousHeader:                  nil,
			currentSortColumnsSortDirection: 0,
		},
		allExecutionsFoOneTestSuiteListObject: allExecutionsFoOneTestSuiteListObjectStruct{
			lastSelectedTestSuiteExecutionForAllExecutionsForOneTestSuite: "",
			testSuiteUuidForTestSuiteExecutionThatIsShownInPreview:        "",
			testSuiteExecutionUuidThatIsShownInPreview:                    "",
			testSuiteExecutionVersionThatIsShownInPreview:                 0,
			isAnyRowSelected:                false,
			currentSortColumn:               0,
			previousSortColumn:              0,
			currentHeader:                   nil,
			previousHeader:                  nil,
			currentSortColumnsSortDirection: 0,
		},
		ExecutionsInGuiIsOfType: 0,
	}

	// Local function - Define the function to be executed to load TestSuiteExecutions from that Database that the user can view
	// Only loads one TestSuiteExecution per TestCase
	loadOneTestSuiteExecutionPerTestCaseFromDataBaseFunction = func() {

		// Call public load-function
		LoadOneTestSuiteExecutionPerTestCaseFromDataBaseFunction(testSuiteExecutionsModel,
			true,
			testCaseInstructionPreViewObject) // ***

	}

	// Define the 'loadTestSuiteExecutionsFromDataBaseButton'
	loadTestSuiteExecutionsFromDataBaseButton = widget.NewButton("Load one TestSuiteExecution per TestCase",
		loadOneTestSuiteExecutionPerTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCasesMapPtr that the user can edit
	filterTestSuiteExcutionsButtonFunction = func() {
		fmt.Println("'filterTestSuiteExcutionButton' was pressed")

		loadTestSuiteExecutionListTableTable(
			testSuiteExecutionsModel,
			false,
			"")
		calculateAndSetCorrectColumnWidths()
		updateTestSuiteExecutionsListTable(testSuiteExecutionsModel)

		// Update the number TestSuiteExcutionss in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testSuiteExecutionsListTableTable))
		numberOfTestSuiteExecutionsAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCasesMapPtr after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCasesMapPtr retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testSuiteExecutionsListTableTable))
		numberOfTestSuiteExecutionsInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCasesMapPtr retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

	}

	// Define the 'filterTestSuiteExcutionsButton'
	filterTestSuiteExcutionsButton = widget.NewButton("Filter TestCasesMapPtr", filterTestSuiteExcutionsButtonFunction)

	// Define the function to be executed to list TestCasesMapPtr that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'clearFiltersButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Set initial button text for 'loadAllTestSuiteExecutionsForOneTestCaseButton'
	loadAllTestSuiteExecutionsForOneTestSuiteButtonText = loadAllTestSuiteExecutionsForOneTestSuiteButtonTextPart1 + "_"

	// Define the function to be executed to load all TestSuiteExecutions for a specific TestCase. Executions are shown in list
	loadAllTestSuiteExecutionsForOneTestCaseButtonFunction = func() {

		// When no TestSuiteExecution is selected then inform the user
		if len(selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
			testSuiteExecutionUuidThatIsShownInPreview) == 0 ||
			selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
				isAnyRowSelected == false {

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Load TestSuiteExecutions",
				Content: "Select a TestSuiteExecution before loading alla executions for the TestCase",
			})

			return

		}

		fmt.Println("'loadAllTestSuiteExecutionsForOneTestCaseButton' was pressed")

		// If previously data set is of other kind then 'AllExecutionsForOneTestSuite' then no selection should be made
		if selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType != AllExecutionsForOneTestSuite {
			selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.
				isAnyRowSelected = false
			testCaseInstructionPreViewObject.ClearTestSuiteExecutionPreviewContainer()
			loadAllTestSuiteExecutionsForOneTestSuiteButtonReference.Disable()
		}

		// Specify from were Executions in the GUI comes from
		selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType = AllExecutionsForOneTestSuite

		// Channel used to trigger update of Gui
		var updateGuiChannelStep1 chan bool
		var updateGuiChannelStep2 chan bool
		updateGuiChannelStep1 = make(chan bool)
		updateGuiChannelStep2 = make(chan bool)

		listTestSuiteExecutionsModel.LoadTestSuiteExecutionsThatCanBeViewedByUser(
			testSuiteExecutionsModel.LatestTestSuiteExecutionForEachTestSuiteUuid.LatestUniqueTestSuiteExecutionDatabaseRowId,
			true,
			testSuiteExecutionsModel.StandardTestSuiteExecutionsBatchSize,
			true,
			selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
				testSuiteUuidForTestSuiteExecutionThatIsShownInPreview,
			testSuiteExecutionsModel.NullTimeStampForTestSuiteExecutionsSearch,
			testSuiteExecutionsModel.NullTimeStampForTestSuiteExecutionsSearch,
			true,
			&updateGuiChannelStep1)

		// Update the UI with the new data
		go func() {

			// Wait for trigger to update GUI
			<-updateGuiChannelStep1
			<-updateGuiChannelStep2

			// Update the GUI
			sortGuiTableAscendingOnTestSuiteExecutionTimeStamp()

		}()

		// Update the GUI
		sortGuiTableAscendingOnTestSuiteExecutionTimeStamp()
		updateGuiChannelStep2 <- true

	}

	// Define the 'loadAllTestSuiteExecutionsForOneTestCaseButton'
	loadAllTestSuiteExecutionsForOneTestCaseButton = widget.NewButton(loadAllTestSuiteExecutionsForOneTestSuiteButtonText,
		loadAllTestSuiteExecutionsForOneTestCaseButtonFunction)

	// Disable button from the beginning
	loadAllTestSuiteExecutionsForOneTestCaseButton.Disable()

	// Store reference to button
	loadAllTestSuiteExecutionsForOneTestSuiteButtonReference = loadAllTestSuiteExecutionsForOneTestCaseButton

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(
		loadTestSuiteExecutionsFromDataBaseButton,
		filterTestSuiteExcutionsButton,
		clearFiltersButton,
		loadAllTestSuiteExecutionsForOneTestCaseButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestSuiteExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	generateTestSuiteExecutionsListTable(testSuiteExecutionsModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testSuiteExecutionsListTable)

	// Create the Scroll container for the List
	testSuiteExecutionsListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCasesMapPtr in the local filter
	numberOfTestSuiteExecutionsAfterLocalFilters = binding.NewString()
	_ = numberOfTestSuiteExecutionsAfterLocalFilters.Set("No TestCasesMapPtr in the List")
	numberOfTestSuiteExecutionsAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestSuiteExecutionsAfterLocalFilters)

	// Create the label used for showing number of TestCasesMapPtr retrieved from the Database
	numberOfTestSuiteExecutionsInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestSuiteExecutionsInTheDatabaseSearch.Set("No TestCasesMapPtr retrieved from the Database")
	numberOfTestSuiteExecutionsRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestSuiteExecutionsInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestSuiteExecutionsAfterLocalFilterLabel, numberOfTestSuiteExecutionsRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testSuiteExecutionsListScrollContainer' to 'testSuiteExecutionsListContainer'
	testSuiteExecutionsListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testSuiteExecutionsListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testSuiteExecutionsListContainer)

	// ***********
	/*


		// Create the Temporary container that should be shown
		temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))

		testCaseInstructionPreViewObject.
			testSuiteExecutionPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

		// Generate the container for the Preview, 'testSuiteExecutionPreviewContainer'
		testCaseInstructionPreViewObject.
			testSuiteExecutionPreviewContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
			testSuiteExecutionPreviewContainer)

		// Create the temporary container for the logs
		tempTestInstructionsExecutionLogContainer := container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))

		// Add the temporary container to the Border container for the logs
		testCaseInstructionPreViewObject.
			testInstructionsExecutionLogContainer = container.
			NewBorder(nil, nil, nil, nil, tempTestInstructionsExecutionLogContainer)

		// Generate the Attribute-container for the Tab-object
		testCaseInstructionPreViewObject.
			testInstructionsExecutionAttributesContainer = container.
			NewBorder(nil, nil, nil, nil,
				widget.NewLabel("Select an attribute to get the full attribute-value"))

		// Generate the scroll-container used for Attributes-explorer
		testCaseInstructionPreViewObject.
			testInstructionsExecutionAttributesContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
			testInstructionsExecutionAttributesContainer)

		// Generate the TestInstructionExecution-container for the Tab-object
		testCaseInstructionPreViewObject.
			testInstructionsExecutionDetailsContainer = container.
			NewBorder(nil, nil, nil, nil,
				widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the details"))

		// Generate the scroll-container used for TestInstructionExecutionsDetails-explorer
		testCaseInstructionPreViewObject.
			testInstructionsExecutionDetailsContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
			testInstructionsExecutionDetailsContainer)

		// Generate PreViewTab-object
		testCaseInstructionPreViewObject.
			logsExplorerTab = container.NewTabItem("Logs Explorer", testCaseInstructionPreViewObject.
			testInstructionsExecutionLogContainer)

		// Generate AttributeExplorerTab
		testCaseInstructionPreViewObject.
			attributeExplorerTab = container.NewTabItem("Attribute Explorer", testCaseInstructionPreViewObject.
			testInstructionsExecutionAttributesContainerScroll)

		// Generate DestInstructionDetailsExplorerTab
		testCaseInstructionPreViewObject.
			testInstructionDetailsExplorerTab = container.NewTabItem("TestInstructionDetails Explorer",
			testCaseInstructionPreViewObject.testInstructionsExecutionDetailsContainerScroll)

		// Generate the 'PreView-tabs'-object
		testCaseInstructionPreViewObject.
			preViewTabs = container.NewAppTabs(testCaseInstructionPreViewObject.testInstructionDetailsExplorerTab,
			testCaseInstructionPreViewObject.logsExplorerTab,
			testCaseInstructionPreViewObject.attributeExplorerTab)
		testCaseInstructionPreViewObject.
			preViewTabs.OnSelected = func(item *container.TabItem) {
			item.Content.Refresh()
		}

		// Set the Tabs to be positioned in upper part of the object
		testCaseInstructionPreViewObject.
			preViewTabs.SetTabLocation(container.TabLocationTop)

		// make a hoverable transparent Execution-tree-overlay, to stop TestSuiteExecution-nodes to be clickable
		testCaseTreePreViewOverlay := NewHoverableRect(color.Transparent, nil)
		testCaseTreePreViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTestSuiteExecutionPreviewTree = false
			testCaseTreePreViewOverlay.Hide()
			testCaseTreePreViewOverlay.OtherHoverableRect.Show()
		}
		testCaseTreePreViewOverlay.OnMouseOut = func() {

		}

		// make a hoverable transparent preViewTab-overlay, to stop TestSuiteExecution-nodes to be clickable
		explorerTabOverlay := NewHoverableRect(color.Transparent, nil)
		explorerTabOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTestSuiteExecutionPreviewTree = true
			explorerTabOverlay.Hide()
			explorerTabOverlay.OtherHoverableRect.Show()
		}
		explorerTabOverlay.OnMouseOut = func() {

		}

		// Cross connect the two overlays
		testCaseTreePreViewOverlay.OtherHoverableRect = explorerTabOverlay
		explorerTabOverlay.OtherHoverableRect = testCaseTreePreViewOverlay

		testSuiteExecutionPreviewAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
			testSuiteExecutionPreviewContainerScroll, testCaseTreePreViewOverlay)
		preViewTabsAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
			preViewTabs, explorerTabOverlay)

		// Create the split-container that has the Execution-preview to the left and Logs/Attribute to the right
		tempTestCasePreviewTestInstructionExecutionLogSplitContainer = container.NewHSplit(
			testSuiteExecutionPreviewAndOverlayContainer, preViewTabsAndOverlayContainer)
		tempTestCasePreviewTestInstructionExecutionLogSplitContainer.Offset = 0.60

	*/
	// **************

	// Generates the Container structure for the PreView-container
	testCaseInstructionPreViewObject.
		testSuitePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(testCaseInstructionPreViewObject)

	// make a hoverable transparent PreView-overlay, to stop table-row hovering in left Table
	preViewOverlay := NewHoverableRect(color.Transparent, nil)
	preViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTable = true
		preViewOverlay.Hide()
		preViewOverlay.OtherHoverableRect.Show()
	}
	preViewOverlay.OnMouseOut = func() {

	}

	// make a hoverable transparent ExecutionList-overlay, to stop table-row hovering in left Table
	executionListViewOverlay := NewHoverableRect(color.Transparent, nil)
	executionListViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTable = false
		executionListViewOverlay.Hide()
		executionListViewOverlay.OtherHoverableRect.Show()
	}
	executionListViewOverlay.OnMouseOut = func() {

	}

	// Cross connect the two overlays
	preViewOverlay.OtherHoverableRect = executionListViewOverlay
	executionListViewOverlay.OtherHoverableRect = preViewOverlay

	preViewAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
		testSuitePreviewTestInstructionExecutionLogSplitContainer, preViewOverlay)
	executionListAndOverlayContainer := container.New(layout.NewStackLayout(), testCasesListScrollContainer2, executionListViewOverlay)

	// Generate the split-container holding the TestSuiteExecution-list and the Execution-Preview
	tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(executionListAndOverlayContainer, preViewAndOverlayContainer)
	tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.60

	testSuiteExecutionListAndTestSuiteExecutionPreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

	return tempTestCaseListAndTestCasePreviewSplitContainer
}

// Generates the Container structure for the PreView-container
func generatePreViewObject(
	testCaseInstructionPreViewObject *TestSuiteInstructionPreViewStruct) (
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split) {

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))

	testCaseInstructionPreViewObject.
		testSuiteExecutionPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testSuiteExecutionPreviewContainer'
	testCaseInstructionPreViewObject.
		testSuiteExecutionPreviewContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
		testSuiteExecutionPreviewContainer)

	// Create the temporary container for the logs
	tempTestInstructionsExecutionLogContainer := container.NewCenter(
		widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))

	// Add the temporary container to the Border container for the logs
	testCaseInstructionPreViewObject.
		testInstructionsExecutionLogContainer = container.
		NewBorder(nil, nil, nil, nil, tempTestInstructionsExecutionLogContainer)

	// Generate the Attribute-container for the Tab-object
	testCaseInstructionPreViewObject.
		testInstructionsExecutionAttributesContainer = container.
		NewBorder(nil, nil, nil, nil,
			widget.NewLabel("Select an attribute to get the full attribute-value"))

	// Generate the scroll-container used for Attributes-explorer
	testCaseInstructionPreViewObject.
		testInstructionsExecutionAttributesContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
		testInstructionsExecutionAttributesContainer)

	// Generate the TestInstructionExecution-container for the Tab-object
	testCaseInstructionPreViewObject.
		testInstructionsExecutionDetailsContainer = container.
		NewBorder(nil, nil, nil, nil,
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the details"))

	// Generate the scroll-container used for TestInstructionExecutionsDetails-explorer
	testCaseInstructionPreViewObject.
		testInstructionsExecutionDetailsContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
		testInstructionsExecutionDetailsContainer)

	// Generate PreViewTab-object
	testCaseInstructionPreViewObject.
		logsExplorerTab = container.NewTabItem("Logs Explorer", testCaseInstructionPreViewObject.
		testInstructionsExecutionLogContainer)

	// Generate AttributeExplorerTab
	testCaseInstructionPreViewObject.
		attributeExplorerTab = container.NewTabItem("Attribute Explorer", testCaseInstructionPreViewObject.
		testInstructionsExecutionAttributesContainerScroll)

	// Generate DestInstructionDetailsExplorerTab
	testCaseInstructionPreViewObject.
		testInstructionDetailsExplorerTab = container.NewTabItem("TestInstructionDetails Explorer",
		testCaseInstructionPreViewObject.testInstructionsExecutionDetailsContainerScroll)

	// Generate the 'PreView-tabs'-object
	testCaseInstructionPreViewObject.
		preViewTabs = container.NewAppTabs(testCaseInstructionPreViewObject.testInstructionDetailsExplorerTab,
		testCaseInstructionPreViewObject.logsExplorerTab,
		testCaseInstructionPreViewObject.attributeExplorerTab)
	testCaseInstructionPreViewObject.
		preViewTabs.OnSelected = func(item *container.TabItem) {
		item.Content.Refresh()
	}

	// Set the Tabs to be positioned in upper part of the object
	testCaseInstructionPreViewObject.
		preViewTabs.SetTabLocation(container.TabLocationTop)

	// make a hoverable transparent Execution-tree-overlay, to stop TestSuiteExecution-nodes to be clickable
	testCaseTreePreViewOverlay := NewHoverableRect(color.Transparent, nil)
	testCaseTreePreViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTestSuiteExecutionPreviewTree = false
		testCaseTreePreViewOverlay.Hide()
		testCaseTreePreViewOverlay.OtherHoverableRect.Show()
	}
	testCaseTreePreViewOverlay.OnMouseOut = func() {

	}

	// make a hoverable transparent preViewTab-overlay, to stop TestSuiteExecution-nodes to be clickable
	explorerTabOverlay := NewHoverableRect(color.Transparent, nil)
	explorerTabOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTestSuiteExecutionPreviewTree = true
		explorerTabOverlay.Hide()
		explorerTabOverlay.OtherHoverableRect.Show()
	}
	explorerTabOverlay.OnMouseOut = func() {

	}

	// Cross connect the two overlays
	testCaseTreePreViewOverlay.OtherHoverableRect = explorerTabOverlay
	explorerTabOverlay.OtherHoverableRect = testCaseTreePreViewOverlay

	testSuiteExecutionPreviewAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
		testSuiteExecutionPreviewContainerScroll, testCaseTreePreViewOverlay)
	preViewTabsAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
		preViewTabs, explorerTabOverlay)

	// Create the split-container that has the Execution-preview to the left and Logs/Attribute to the right
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer = container.NewHSplit(
		testSuiteExecutionPreviewAndOverlayContainer, preViewTabsAndOverlayContainer)
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer.Offset = 0.60

	return tempTestCasePreviewTestInstructionExecutionLogSplitContainer

}
