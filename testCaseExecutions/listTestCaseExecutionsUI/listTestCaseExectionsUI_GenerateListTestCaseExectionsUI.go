package listTestCaseExecutionsUI

import (
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/soundEngine"
	"FenixTesterGui/testCaseExecutions/listTestCaseExecutionsModel"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
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
	"time"
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

// Define the function to be executed to load TestCaseExecutions from that Database that the user can view
// Only loads one TestCaseExecution per TestCase
func LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction(
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct,
	updateGui bool,
	testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) {

	// If previously data set is of other kind then 'OneExecutionPerTestCase' then no selection should be made
	if selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType != OneExecutionPerTestCase {
		//selectedTestCaseExecutionObjected.isAnyRowSelected = false
		testCaseInstructionPreViewObject.ClearTestCaseExecutionPreviewContainer()
		loadAllTestCaseExecutionsForOneTestCaseButtonReference.Disable()

	}

	// Get number of TestCase retrieved from Database before and after loading more data
	//var beforeNumberOfRowsRetrievedFromDatabase int
	//var afterNumberOfRowsRetrievedFromDatabase int
	//beforeNumberOfRowsRetrievedFromDatabase = testCaseExecutionsModel.
	//	GetNumberOfTestCaseExecutionsRetrievedFromDatabase()

	// Specify from were Executions in the GUI comes from
	selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType = OneExecutionPerTestCase

	// Channel used to trigger update of Gui
	var updateGuiChannelStep1 chan bool
	var updateGuiChannelStep2 chan bool
	updateGuiChannelStep1 = make(chan bool)
	updateGuiChannelStep2 = make(chan bool)

	listTestCaseExecutionsModel.LoadTestCaseExecutionsThatCanBeViewedByUser(
		testCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.LatestUniqueTestCaseExecutionDatabaseRowId,
		true,
		testCaseExecutionsModel.StandardTestCaseExecutionsBatchSize,
		false,
		"",
		testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
		testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
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
			afterNumberOfRowsRetrievedFromDatabase = testCaseExecutionsModel.
				GetNumberOfTestCaseExecutionsRetrievedFromDatabase()

			// If they are different then sort the table
			if beforeNumberOfRowsRetrievedFromDatabase != afterNumberOfRowsRetrievedFromDatabase {
				//sortableHeaderReference.sortImage.onTapped()
				testCaseExecutionsListTableHeadersMapRef[int(latestTestCaseExecutionTimeStampColumnNumber)].sortImage.onTapped()

			}

		*/

	}()

	updateGuiChannelStep2 <- true
	//filterTestCaseExcutionsButtonFunction()

}

// Create the UI used for list all TestCasesMapPtr that the User can edit
func GenerateListTestCaseExecutionsUI(
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct,
	detailedTestCaseExecutionsUITabObject *container.DocTabs,
	exitFunctionsForDetailedTestCaseExecutionsUITabObject *map[*container.TabItem]func(),
	testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) (
	listTestCasesUI fyne.CanvasObject) {

	// save reference to 'detailedTestCaseExecutionsUITabObject'
	detailedTestCaseExecutionsUITabObjectRef = detailedTestCaseExecutionsUITabObject

	// save Reference to Map over Exit-functions for the Tabs in 'detailedTestCaseExecutionsUITabObject'
	exitFunctionsForDetailedTestCaseExecutionsUITabObjectPtr = exitFunctionsForDetailedTestCaseExecutionsUITabObject

	//var testCaseTable *widget.Table

	var tempTestCaseListAndTestCasePreviewSplitContainer *container.Split
	//var tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split

	var testCaseExecutionsListContainer *fyne.Container
	var testCaseExecutionsListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	// Button used for Retrieving all TestCaseExecutions
	var loadTestCaseExecutionsFromDataBaseButton *widget.Button
	var loadOneTestCaseExecutionPerTestCaseFromDataBaseFunction func()

	var filterTestCaseExcutionsButton *widget.Button
	var filterTestCaseExcutionsButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var loadAllTestCaseExecutionsForOneTestCaseButton *widget.Button
	var loadAllTestCaseExecutionsForOneTestCaseButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestCaseExecutionsAfterLocalFilterLabel *widget.Label
	var numberOfTestCaseExecutionsRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Initiate map for Header-references
	testCaseExecutionsListTableHeadersMapRef = make(map[int]*sortableHeaderLabelStruct)

	// Initiate 'selectedTestCaseExecutionObjected' used for keep track of source of executions and what row is selected
	selectedTestCaseExecutionObjected = selectedTestCaseExecutionStruct{
		oneExecutionPerTestCaseListObject: oneExecutionPerTestCaseListObjectStruct{
			lastSelectedTestCaseExecutionForOneExecutionPerTestCase: "",
			testCaseUuidForTestCaseExecutionThatIsShownInPreview:    "",
			testCaseExecutionUuidThatIsShownInPreview:               "",
			testCaseExecutionVersionThatIsShownInPreview:            0,
			isAnyRowSelected:                false,
			currentSortColumn:               0,
			previousSortColumn:              0,
			currentHeader:                   nil,
			previousHeader:                  nil,
			currentSortColumnsSortDirection: 0,
		},
		allExecutionsFoOneTestCaseListObject: allExecutionsFoOneTestCaseListObjectStruct{
			lastSelectedTestCaseExecutionForAllExecutionsForOneTestCase: "",
			testCaseUuidForTestCaseExecutionThatIsShownInPreview:        "",
			testCaseExecutionUuidThatIsShownInPreview:                   "",
			testCaseExecutionVersionThatIsShownInPreview:                0,
			isAnyRowSelected:                false,
			currentSortColumn:               0,
			previousSortColumn:              0,
			currentHeader:                   nil,
			previousHeader:                  nil,
			currentSortColumnsSortDirection: 0,
		},
		ExecutionsInGuiIsOfType: 0,
	}

	// Local function - Define the function to be executed to load TestCaseExecutions from that Database that the user can view
	// Only loads one TestCaseExecution per TestCase
	loadOneTestCaseExecutionPerTestCaseFromDataBaseFunction = func() {

		// Call public load-function
		LoadOneTestCaseExecutionPerTestCaseFromDataBaseFunction(testCaseExecutionsModel,
			true,
			testCaseInstructionPreViewObject)

	}

	// Define the 'loadTestCaseExecutionsFromDataBaseButton'
	loadTestCaseExecutionsFromDataBaseButton = widget.NewButton("Load one TestCaseExecution per TestCase",
		loadOneTestCaseExecutionPerTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCasesMapPtr that the user can edit
	filterTestCaseExcutionsButtonFunction = func() {
		fmt.Println("'filterTestCaseExcutionButton' was pressed")

		loadTestCaseExecutionListTableTable(
			testCaseExecutionsModel,
			false,
			"")
		calculateAndSetCorrectColumnWidths()
		updateTestCaseExecutionsListTable(testCaseExecutionsModel)

		// Update the number TestCaseExcutionss in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testCaseExecutionsListTableTable))
		numberOfTestCaseExecutionsAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCasesMapPtr after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCasesMapPtr retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testCaseExecutionsListTableTable))
		numberOfTestCaseExecutionsInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCasesMapPtr retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

	}

	// Define the 'filterTestCaseExcutionsButton'
	filterTestCaseExcutionsButton = widget.NewButton("Filter TestCasesMapPtr", filterTestCaseExcutionsButtonFunction)

	// Define the function to be executed to list TestCasesMapPtr that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'clearFiltersButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Set initial button text for 'loadAllTestCaseExecutionsForOneTestCaseButton'
	loadAllTestCaseExecutionsForOneTestCaseButtonText = loadAllTestCaseExecutionsForOneTestCaseButtonTextPart1 + "_"

	// Define the function to be executed to load all TestCaseExecutions for a specific TestCase. Executions are shown in list
	loadAllTestCaseExecutionsForOneTestCaseButtonFunction = func() {

		// When no TestCaseExecution is selected then inform the user
		if len(selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
			testCaseExecutionUuidThatIsShownInPreview) == 0 ||
			selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
				isAnyRowSelected == false {

			// Trigger System Notification sound
			soundEngine.PlaySoundChannel <- soundEngine.SystemNotificationSound

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Load TestCaseExecutions",
				Content: "Select a TestCaseExecution before loading alla executions for the TestCase",
			})

			return

		}

		fmt.Println("'loadAllTestCaseExecutionsForOneTestCaseButton' was pressed")

		// If previously data set is of other kind then 'AllExecutionsForOneTestCase' then no selection should be made
		if selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType != AllExecutionsForOneTestCase {
			selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.
				isAnyRowSelected = false
			testCaseInstructionPreViewObject.ClearTestCaseExecutionPreviewContainer()
			loadAllTestCaseExecutionsForOneTestCaseButtonReference.Disable()
		}

		// Specify from were Executions in the GUI comes from
		selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType = AllExecutionsForOneTestCase

		// Channel used to trigger update of Gui
		/*
			var updateGuiChannelStep1 chan bool
			var updateGuiChannelStep2 chan bool
			updateGuiChannelStep1 = make(chan bool)
			updateGuiChannelStep2 = make(chan bool)


		*/
		// Use one buffered channel to handle both batch updates
		updateGuiChannel := make(chan bool, 2)

		fmt.Println("Finished loading all TestCaseExecutions for the TestCase 1a", time.Now().String())

		listTestCaseExecutionsModel.LoadTestCaseExecutionsThatCanBeViewedByUser(
			testCaseExecutionsModel.LatestTestCaseExecutionForEachTestCaseUuid.LatestUniqueTestCaseExecutionDatabaseRowId,
			true,
			testCaseExecutionsModel.StandardTestCaseExecutionsBatchSize,
			true,
			selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
				testCaseUuidForTestCaseExecutionThatIsShownInPreview,
			testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
			testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
			true,
			&updateGuiChannel)

		// ✅ Listen for refresh triggers
		go func() {
			for i := 0; i < 2; i++ {
				<-updateGuiChannel
				fmt.Printf("Refreshing GUI after batch %d at %s\n", i+1, time.Now().String())

				fyne.Do(func() {
					sortGuiTableAscendingOnTestCaseExecutionTimeStamp()
				})
			}

			fmt.Println("Finished loading all TestCaseExecutions for the TestCase 2", time.Now().String())
		}()

		/*
				// Update the UI with the new data
				go func() {

					// Wait for trigger to update GUI
					<-updateGuiChannelStep1
					<-updateGuiChannelStep2

					fmt.Println("Finished loading all TestCaseExecutions for the TestCase 2", time.Now().String())

					// Update the GUI
					sortGuiTableAscendingOnTestCaseExecutionTimeStamp()

				}()



			fmt.Println("Finished loading all TestCaseExecutions for the TestCase 1b", time.Now().String())

			// Update the GUI
			sortGuiTableAscendingOnTestCaseExecutionTimeStamp()
			updateGuiChannelStep2 <- true

		*/

		// Immediate first refresh
		fmt.Println("Finished loading all TestCaseExecutions for the TestCase 1b", time.Now().String())
		fyne.Do(func() {
			sortGuiTableAscendingOnTestCaseExecutionTimeStamp()
		})

	}

	// Define the 'loadAllTestCaseExecutionsForOneTestCaseButton'
	loadAllTestCaseExecutionsForOneTestCaseButton = widget.NewButton(loadAllTestCaseExecutionsForOneTestCaseButtonText,
		loadAllTestCaseExecutionsForOneTestCaseButtonFunction)

	// Disable button from the beginning
	loadAllTestCaseExecutionsForOneTestCaseButton.Disable()

	// Store reference to button
	loadAllTestCaseExecutionsForOneTestCaseButtonReference = loadAllTestCaseExecutionsForOneTestCaseButton

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(
		loadTestCaseExecutionsFromDataBaseButton,
		filterTestCaseExcutionsButton,
		clearFiltersButton,
		loadAllTestCaseExecutionsForOneTestCaseButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestCaseExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	generateTestCaseExecutionsListTable(testCaseExecutionsModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testCaseExecutionsListTable)

	// Create the Scroll container for the List
	testCaseExecutionsListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCasesMapPtr in the local filter
	numberOfTestCaseExecutionsAfterLocalFilters = binding.NewString()
	_ = numberOfTestCaseExecutionsAfterLocalFilters.Set("No TestCasesMapPtr in the List")
	numberOfTestCaseExecutionsAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestCaseExecutionsAfterLocalFilters)

	// Create the label used for showing number of TestCasesMapPtr retrieved from the Database
	numberOfTestCaseExecutionsInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestCaseExecutionsInTheDatabaseSearch.Set("No TestCasesMapPtr retrieved from the Database")
	numberOfTestCaseExecutionsRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestCaseExecutionsInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCaseExecutionsAfterLocalFilterLabel, numberOfTestCaseExecutionsRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCaseExecutionsListScrollContainer' to 'testCaseExecutionsListContainer'
	testCaseExecutionsListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCaseExecutionsListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testCaseExecutionsListContainer)

	// ***********
	/*


		// Create the Temporary container that should be shown
		temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCaseExecution to get the Preview"))

		testCaseInstructionPreViewObject.
			testCaseExecutionPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

		// Generate the container for the Preview, 'testCaseExecutionPreviewContainer'
		testCaseInstructionPreViewObject.
			testCaseExecutionPreviewContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
			testCaseExecutionPreviewContainer)

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

		// make a hoverable transparent Execution-tree-overlay, to stop TestCaseExecution-nodes to be clickable
		testCaseTreePreViewOverlay := NewHoverableRect(color.Transparent, nil)
		testCaseTreePreViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTestCaseExecutionPreviewTree = false
			testCaseTreePreViewOverlay.Hide()
			testCaseTreePreViewOverlay.OtherHoverableRect.Show()
		}
		testCaseTreePreViewOverlay.OnMouseOut = func() {

		}

		// make a hoverable transparent preViewTab-overlay, to stop TestCaseExecution-nodes to be clickable
		explorerTabOverlay := NewHoverableRect(color.Transparent, nil)
		explorerTabOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

			mouseHasLeftTestCaseExecutionPreviewTree = true
			explorerTabOverlay.Hide()
			explorerTabOverlay.OtherHoverableRect.Show()
		}
		explorerTabOverlay.OnMouseOut = func() {

		}

		// Cross connect the two overlays
		testCaseTreePreViewOverlay.OtherHoverableRect = explorerTabOverlay
		explorerTabOverlay.OtherHoverableRect = testCaseTreePreViewOverlay

		testCaseExecutionPreviewAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
			testCaseExecutionPreviewContainerScroll, testCaseTreePreViewOverlay)
		preViewTabsAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
			preViewTabs, explorerTabOverlay)

		// Create the split-container that has the Execution-preview to the left and Logs/Attribute to the right
		tempTestCasePreviewTestInstructionExecutionLogSplitContainer = container.NewHSplit(
			testCaseExecutionPreviewAndOverlayContainer, preViewTabsAndOverlayContainer)
		tempTestCasePreviewTestInstructionExecutionLogSplitContainer.Offset = 0.60

	*/
	// **************

	// Generates the Container structure for the PreView-container
	testCaseInstructionPreViewObject.
		testCasePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(testCaseInstructionPreViewObject)

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
		testCasePreviewTestInstructionExecutionLogSplitContainer, preViewOverlay)
	executionListAndOverlayContainer := container.New(layout.NewStackLayout(),
		testCasesListScrollContainer2,
		executionListViewOverlay)

	// Generate the split-container holding the TestCaseExecution-list and the Execution-Preview
	tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(executionListAndOverlayContainer, preViewAndOverlayContainer)
	tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.60

	TestCaseExecutionListAndTestCaseExecutionPreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

	return tempTestCaseListAndTestCasePreviewSplitContainer
}

// Generates the Container structure for the PreView-container
func generatePreViewObject(
	testCaseInstructionPreViewObject *TestCaseInstructionPreViewStruct) (
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer *container.Split) {

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCaseExecution to get the Preview"))

	testCaseInstructionPreViewObject.
		testCaseExecutionPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testCaseExecutionPreviewContainer'
	testCaseInstructionPreViewObject.
		testCaseExecutionPreviewContainerScroll = container.NewScroll(testCaseInstructionPreViewObject.
		testCaseExecutionPreviewContainer)

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

	// make a hoverable transparent Execution-tree-overlay, to stop TestCaseExecution-nodes to be clickable
	testCaseTreePreViewOverlay := NewHoverableRect(color.Transparent, nil)
	testCaseTreePreViewOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTestCaseExecutionPreviewTree = false
		testCaseTreePreViewOverlay.Hide()
		testCaseTreePreViewOverlay.OtherHoverableRect.Show()
	}
	testCaseTreePreViewOverlay.OnMouseOut = func() {

	}

	// make a hoverable transparent preViewTab-overlay, to stop TestCaseExecution-nodes to be clickable
	explorerTabOverlay := NewHoverableRect(color.Transparent, nil)
	explorerTabOverlay.OnMouseIn = func(ev *desktop.MouseEvent) {

		mouseHasLeftTestCaseExecutionPreviewTree = true
		explorerTabOverlay.Hide()
		explorerTabOverlay.OtherHoverableRect.Show()
	}
	explorerTabOverlay.OnMouseOut = func() {

	}

	// Cross connect the two overlays
	testCaseTreePreViewOverlay.OtherHoverableRect = explorerTabOverlay
	explorerTabOverlay.OtherHoverableRect = testCaseTreePreViewOverlay

	testCaseExecutionPreviewAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
		testCaseExecutionPreviewContainerScroll, testCaseTreePreViewOverlay)
	preViewTabsAndOverlayContainer := container.New(layout.NewStackLayout(), testCaseInstructionPreViewObject.
		preViewTabs, explorerTabOverlay)

	// Create the split-container that has the Execution-preview to the left and Logs/Attribute to the right
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer = container.NewHSplit(
		testCaseExecutionPreviewAndOverlayContainer, preViewTabsAndOverlayContainer)
	tempTestCasePreviewTestInstructionExecutionLogSplitContainer.Offset = 0.60

	return tempTestCasePreviewTestInstructionExecutionLogSplitContainer

}
