package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/testCaseExecutions/listTestCaseExecutionsModel"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
	"bytes"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"image"
	"image/color"
	"image/png"
	"log"
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

//go:embed resources/sort-up-ascending_cropped_32x51.png
var sortImageAscendingAsByteArray []byte
var sortImageAscendingAsImage image.Image

//go:embed resources/sort-down-descending_cropped_32x51.png
var sortImageDescendingAsByteArray []byte
var sortImageDescendingAsImage image.Image

// Create the UI used for list all TestCases that the User can edit
func GenerateListTestCaseExecutionsUI(
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) (
	listTestCasesUI fyne.CanvasObject) {

	//var testCaseTable *widget.Table

	var tempTestCaseListAndTestCasePreviewSplitContainer *container.Split

	var testCaseExecutionsListContainer *fyne.Container
	var testCaseExecutionsListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	var loadTestCaseExecutionsFromDataBaseButton *widget.Button
	var loadTestCaseExcutionsFromDataBaseFunction func()
	var filterTestCaseExcutionsButton *widget.Button
	var filterTestCaseExcutionsButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestCaseExecutionsAfterLocalFilterLabel *widget.Label
	var numberOfTestCaseExcutionsRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to load TestCaseExecutions from that Database that the user can view
	loadTestCaseExcutionsFromDataBaseFunction = func() {
		fmt.Println("'loadTestCaseExecutionsFromDataBaseButton' was pressed")
		listTestCaseExecutionsModel.LoadTestCaseExecutionsThatCanBeViewedByUser(
			testCaseExecutionsModel.LatestUniqueTestCaseExecutionDatabaseRowId,
			true,
			testCaseExecutionsModel.StandardTestCaseExecutionsBatchSize,
			testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
			testCaseExecutionsModel.NullTimeStampForTestCaseExecutionsSearch,
			true)

		filterTestCaseExcutionsButtonFunction()
	}

	// Define the 'loadTestCaseExecutionsFromDataBaseButton'
	loadTestCaseExecutionsFromDataBaseButton = widget.NewButton("Load TestCaseExecutions from Database",
		loadTestCaseExcutionsFromDataBaseFunction)

	// Define the function to be executed to filter TestCases that the user can edit
	filterTestCaseExcutionsButtonFunction = func() {
		fmt.Println("'filterTestCaseExecutionsButton' was pressed")
		loadTestCaseExecutionListTableTable(testCaseExecutionsModel)
		calculateAndSetCorrectColumnWidths()
		updateTestCaseExecutionsListTable(testCaseExecutionsModel)

		// Update the number TestCaseExcutionss in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testCaseExecutionsListTableTable))
		numberOfTestCaseExecutionsAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCases after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCases retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testCaseExecutionsListTableTable))
		numberOfTestCaseExecutionsInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCases retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

	}

	// Define the 'filterTestCaseExcutionsButton'
	filterTestCaseExcutionsButton = widget.NewButton("Filter TestCases", filterTestCaseExcutionsButtonFunction)

	// Define the function to be executed to list TestCases that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'filterTestCaseExcutionsButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(loadTestCaseExecutionsFromDataBaseButton, filterTestCaseExcutionsButton, clearFiltersButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestCaseExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	generateTestCaseExecutionsListTable(testCaseExecutionsModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testCaseExecutionsListTable)

	// Create the Scroll container for the List
	testCaseExecutionsListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCases in the local filter
	numberOfTestCaseExecutionsAfterLocalFilters = binding.NewString()
	_ = numberOfTestCaseExecutionsAfterLocalFilters.Set("No TestCases in the List")
	numberOfTestCaseExecutionsAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestCaseExecutionsAfterLocalFilters)

	// Create the label used for showing number of TestCases retrieved from the Database
	numberOfTestCaseExecutionsInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestCaseExecutionsInTheDatabaseSearch.Set("No TestCases retrieved from the Database")
	numberOfTestCaseExcutionsRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestCaseExecutionsInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCaseExecutionsAfterLocalFilterLabel, numberOfTestCaseExcutionsRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCaseExecutionsListScrollContainer' to 'testCaseExecutionsListContainer'
	testCaseExecutionsListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCaseExecutionsListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testCaseExecutionsListContainer)

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCase to get the Preview"))

	testCaseExecutionPreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testCaseExecutionPreviewContainer'
	testCaseExecutionPreviewContainerScroll = container.NewScroll(testCaseExecutionPreviewContainer)

	tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(testCasesListScrollContainer2, testCaseExecutionPreviewContainerScroll)
	tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.75

	TestCaseExecutionListAndTestCaseExecutionPreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

	return tempTestCaseListAndTestCasePreviewSplitContainer
}

func GenerateTestCaseExecutionPreviewContainer(
	testCaseExecutionUuid string,
	testCaseExecutionsModel *testCaseExecutionsModel.TestCaseExecutionsModelStruct) {

	var testCaseExecutionPreviewTopContainer *fyne.Container
	var testCaseExecutionPreviewBottomContainer *fyne.Container
	//var testCasePreviewScrollContainer *container.Scroll
	var testCaseExecutionMainAreaForPreviewContainer *fyne.Container

	var err error
	var existInMap bool
	var foundValue bool

	var tempTestCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Verify that number Headers match number of columns, constant 'numberColumnsInTestCaseExecutionsListUI'
	if len(testCaseExecutionsListTableHeader) != numberColumnsInTestCaseExecutionsListUI {
		log.Fatalln(fmt.Sprintf("Number of elements in 'tempRowslice' missmatch contant 'numberColumnsInTestCaseExecutionsListUI'. %d vs %d. ID: %s",
			testCaseExecutionsListTableHeader,
			numberColumnsInTestCaseExecutionsListUI,
			"c2b8a13c-ec20-46c2-adf9-965247732e07"))
	}

	// Get Data for the Preview
	tempTestCaseExecutionsListMessage = testCaseExecutionsModel.TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutionUuid]

	// Create the Top container
	testCaseExecutionPreviewTopContainer = container.New(layout.NewFormLayout())

	// Create the ExecutionStatus Rectangle for TestCaseExecution-status
	var tempTestCaseExecutionStatusRectangle *canvas.Rectangle
	tempTestCaseExecutionStatusRectangle = canvas.NewRectangle(color.Transparent)

	// Resize the ExecutionStatus rectangle
	tempTestCaseExecutionStatusRectangle.SetMinSize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHight,
	})
	tempTestCaseExecutionStatusRectangle.Resize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHight,
	})

	// Set correct color on ExecutionStatus Rectangle
	var statusId uint8

	// Extract TestCaseExecution-status
	var tempTestCaseExecutionStatusEnum string
	tempTestCaseExecutionStatusEnum = tempTestCaseExecutionsListMessage.GetTestCaseExecutionStatus().String()[4:]

	statusId = detailedExecutionsModel.
		ExecutionStatusColorNameToNumberMap[tempTestCaseExecutionStatusEnum].ExecutionStatusNumber

	var executionStatusColorMapObjectForTestCaseExecution detailedExecutionsModel.ExecutionStatusColorMapStruct
	executionStatusColorMapObjectForTestCaseExecution, existInMap = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)]
	if existInMap == false {
		// No matching Status color exist due to that TestInstruction exists in ExecutionQueue
		// 'INITIATED = 1'
		executionStatusColorMapObjectForTestCaseExecution, _ = detailedExecutionsModel.ExecutionStatusColorMap[1]
		/*
			tempTestCaseExecutionStatusRectangle.StrokeWidth = 2
			tempTestCaseExecutionStatusRectangle.StrokeColor = color.NRGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xFF,
			}
		*/

	} else {
		// Status color found
		tempTestCaseExecutionStatusRectangle.FillColor = executionStatusColorMapObjectForTestCaseExecution.BackgroundColor

		if executionStatusColorMapObjectForTestCaseExecution.UseStroke == true {
			tempTestCaseExecutionStatusRectangle.StrokeWidth = 2
			tempTestCaseExecutionStatusRectangle.StrokeColor = executionStatusColorMapObjectForTestCaseExecution.StrokeColor
		}
	}

	// Add TestCaseName to Top container
	tempTestCaseNameLabel := widget.NewLabel("TestCaseName:")
	tempTestCaseNameLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempTestCaseNameLabel)
	testCaseExecutionPreviewTopContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetTestCaseName()))

	// Add TestCaseExecutionStatus
	tempTestCaseExecutionStatusLabel := widget.NewLabel("TestCaseExecution-status:")
	tempTestCaseExecutionStatusLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempTestCaseExecutionStatusLabel)

	var testCaseNameHBoxContainer *fyne.Container
	testCaseNameHBoxContainer = container.NewHBox()
	testCaseNameHBoxContainer.Add(tempTestCaseExecutionStatusRectangle)
	testCaseExecutionPreviewTopContainer.Add(testCaseNameHBoxContainer)

	// Add TestCaseOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempOwnerDomainLabel)
	testCaseExecutionPreviewTopContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetTestCasePreview().GetDomainThatOwnTheTestCase()))

	// Add emtpy row
	testCaseExecutionPreviewTopContainer.Add(widget.NewLabel(""))
	testCaseExecutionPreviewTopContainer.Add(widget.NewLabel(""))

	// Add TestCase
	tempTestCaseLabel := widget.NewLabel("TestCase with execution status:")
	tempTestCaseLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempTestCaseLabel)
	testCaseExecutionPreviewTopContainer.Add(widget.NewLabel(""))

	/*
		// Add Description Top container
		tempTestCaseDescription := widget.NewLabel("Description:")
		tempTestCaseDescription.TextStyle = fyne.TextStyle{Bold: true}
		testCaseExecutionPreviewTopContainer.Add(tempTestCaseDescription)
		testCaseExecutionPreviewTopContainer.Add(widget.NewRichTextWithText(tempTestCaseExecutionsListMessage.GetTestCaseDescription()))

		// Create the Bottom container
		testCaseExecutionPreviewBottomContainer = container.New(layout.NewFormLayout())

		// Add ComplexTextualDescription to Bottom container
		tempComplexTextualDescriptionLabel := widget.NewLabel("ComplexTextualDescription:")
		tempComplexTextualDescriptionLabel.TextStyle = fyne.TextStyle{Bold: true}
		testCaseExecutionPreviewBottomContainer.Add(tempComplexTextualDescriptionLabel)
		testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetComplexTextualDescription()))


	*/

	// Create the Bottom container
	testCaseExecutionPreviewBottomContainer = container.New(layout.NewFormLayout())

	// Add TestCaseExecutionVersion to Bottom container
	tempTestCaseExecutionVersionLabel := widget.NewLabel("TestCaseExecutionVersion:")
	tempTestCaseExecutionVersionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempTestCaseExecutionVersionLabel)
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(strconv.Itoa(int(tempTestCaseExecutionsListMessage.GetTestCaseExecutionVersion()))))

	// Add ExecutionStartTimeStamp to Bottom container
	tempExecutionStartTimeStampLabel := widget.NewLabel("TestCase Execution Start TimeStamp:")
	tempExecutionStartTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStartTimeStampLabel)
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetExecutionStartTimeStamp().AsTime().String()))

	// Add ExecutionStopTimeStamp to Bottom container
	tempExecutionStopTimeStampLabel := widget.NewLabel("TestCase Execution Stop TimeStamp:")
	tempExecutionStopTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStopTimeStampLabel)
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetExecutionStopTimeStamp().AsTime().String()))

	// Add ExecutionStatusUpdateTimeStamp to Bottom container
	tempExecutionStatusUpdateTimeStampLabel := widget.NewLabel("TestCase Execution Status Update TimeStamp:")
	tempExecutionStatusUpdateTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStatusUpdateTimeStampLabel)
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetExecutionStatusUpdateTimeStamp().AsTime().String()))

	// Add LastSavedByUserGCPAuthorization to Bottom container
	tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("TestCase Execution Stop TimeStamp:")
	tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestCaseExecutionsListMessage.GetExecutionStopTimeStamp().AsTime().String()))

	// Create the area used for TIC, TI and the attributes
	testCaseExecutionMainAreaForPreviewContainer = container.NewVBox()

	// Loop the preview objects and to container
	for _, previewObject := range tempTestCaseExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects {

		// Create the Indentation rectangle
		var tempIndentationLevelRectangle *canvas.Rectangle
		tempIndentationLevelRectangle = canvas.NewRectangle(color.Transparent)

		// Resize the Indentation rectangle
		tempIndentationLevelRectangle.SetMinSize(fyne.Size{
			Width:  float32(10 * previewObject.GetIndentationLevel()),
			Height: 2,
		})
		tempIndentationLevelRectangle.Resize(fyne.Size{
			Width:  float32(10 * previewObject.GetIndentationLevel()),
			Height: 2,
		})

		// Create the ExecutionStatus Rectangle
		var tempExecutionStatusRectangle *canvas.Rectangle
		tempExecutionStatusRectangle = canvas.NewRectangle(color.Transparent)

		// Resize the ExecutionStatus rectangle
		tempExecutionStatusRectangle.SetMinSize(fyne.Size{
			Width:  testCaseExecutionStatusRectangleWidth,
			Height: testCaseExecutionStatusRectangleHight,
		})
		tempExecutionStatusRectangle.Resize(fyne.Size{
			Width:  testCaseExecutionStatusRectangleWidth,
			Height: testCaseExecutionStatusRectangleHight,
		})

		// Decide what type of object
		switch previewObject.TestCaseStructureObjectType {

		case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:

			var serialOrParallelRectangleImage *canvas.Image

			// Create graphics that show if TestInstruction is serial or parallel processed
			if previewObject.TestInstructionIsSerialProcessed == true {

				// Convert the byte slice into an image.Image object
				if imageData_tic_serialImage == nil {
					imageData_tic_serialImage, err = png.Decode(bytes.NewReader(tic_serialImage))
					if err != nil {
						log.Fatalf("Failed to decode image: %v", err)
					}
				}

				serialOrParallelRectangleImage = canvas.NewImageFromImage(imageData_tic_serialImage)
				serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))
				serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))

			} else {
				// Convert the byte slice into an image.Image object
				if imageData_tic_parallellImage == nil {
					imageData_tic_parallellImage, err = png.Decode(bytes.NewReader(tic_parallellImage))
					if err != nil {
						log.Fatalf("Failed to decode image: %v", err)
					}
				}
				serialOrParallelRectangleImage = canvas.NewImageFromImage(imageData_tic_parallellImage)
				serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))
				serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4), float32(testCaseNodeRectangleSize-4)))

			}

			// Create the Name for the TestInstructionContainer
			var tempTestInstructionContainerNameWidget *widget.Label
			tempTestInstructionContainerNameWidget = widget.NewLabel(previewObject.GetTestInstructionContainerName())

			// Create the container containing the TestInstructionContainer
			var tempTestInstructionContainerContainer *fyne.Container
			tempTestInstructionContainerContainer = container.NewHBox(
				tempExecutionStatusRectangle,
				tempIndentationLevelRectangle,
				serialOrParallelRectangleImage,
				tempTestInstructionContainerNameWidget)

			// Add the TestInstructionContainerContainer to the main Area
			testCaseExecutionMainAreaForPreviewContainer.Add(tempTestInstructionContainerContainer)

		case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstruction:

			// Create the Color for the color rectangle
			var rectangleColor color.Color
			rectangleColor, err = sharedCode.ConvertRGBAHexStringIntoRGBAColor(previewObject.TestInstructionColor)
			if err != nil {
				log.Fatalf("Failed to convert hex-color-string '%s' into 'color'. err='%s'", previewObject.TestInstructionColor, err.Error())
			}

			// Create color rectangle for TestInstruction
			var testInstructionColorRectangle *canvas.Rectangle
			testInstructionColorRectangle = canvas.NewRectangle(rectangleColor)

			// Set the size of the color rectangle
			testInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-14), float32(testCaseNodeRectangleSize-14)))
			testInstructionColorRectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-14), float32(testCaseNodeRectangleSize-14)))

			// Create the Name for the TestInstruction
			var tempTestInstructionNameWidget *widget.Label
			tempTestInstructionNameWidget = widget.NewLabel(previewObject.GetTestInstructionName())

			// Set correct color on ExecutionStatus Rectangle
			var statusId uint8
			var statusBackgroundColor color.RGBA
			var statusStrokeColor color.RGBA
			var useStroke bool

			// Extract TestInstructionExecution from TestInstruction
			previewObject.GetTestInstructionUuid()
			var tempTestCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
			tempTestCaseExecutionsListMessage, existInMap = testCaseExecutionsModel.
				TestCaseExecutionsThatCanBeViewedByUserMap[testCaseExecutionThatIsShownInPreview]
			if existInMap == false {
				var id string
				id = "7945551e-5e4d-41d3-8faf-54f1501daac9"
				log.Fatalf(fmt.Sprintf("Couldn't find testCaseExecutionThatIsShownInPreview '%s' in TestCaseExecutionsThatCanBeViewedByUserMap. ID='%s'",
					testCaseExecutionThatIsShownInPreview,
					id))
			}

			var tempTestInstructionExecutionsStatusPreviewValues []*fenixExecutionServerGuiGrpcApi.
				TestInstructionExecutionStatusPreviewValueMessage
			tempTestInstructionExecutionsStatusPreviewValues = tempTestCaseExecutionsListMessage.
				TestInstructionsExecutionStatusPreviewValues.GetTestInstructionExecutionStatusPreviewValues()

			// Loop to find correct TestInstructionExecution
			foundValue = false
			var foundTestInstructionExecutionStatusPreviewValues *fenixExecutionServerGuiGrpcApi.
				TestInstructionExecutionStatusPreviewValueMessage
			for _, tempTestInstructionExecutionStatusPreviewValues := range tempTestInstructionExecutionsStatusPreviewValues {
				if tempTestInstructionExecutionStatusPreviewValues.
					GetMatureTestInstructionUuid() == previewObject.GetTestInstructionUuid() {
					foundValue = true
					foundTestInstructionExecutionStatusPreviewValues = tempTestInstructionExecutionStatusPreviewValues
					break
				}
			}

			if foundValue == false {

				// No matching Status color exist due to that TestInstruction exists in ExecutionQueue
				// 'INITIATED = 1'

				statusBackgroundColor = detailedExecutionsModel.ExecutionStatusColorMap[1].BackgroundColor
				tempExecutionStatusRectangle.FillColor = statusBackgroundColor

				useStroke = detailedExecutionsModel.ExecutionStatusColorMap[1].UseStroke
				if useStroke == true {
					statusStrokeColor = detailedExecutionsModel.ExecutionStatusColorMap[1].StrokeColor
					tempExecutionStatusRectangle.StrokeColor = statusStrokeColor
					tempExecutionStatusRectangle.StrokeWidth = 2
				}
				/*


					// No TestInstructionExecution could be found, set Empty red box to indicate
					tempExecutionStatusRectangle.StrokeColor = color.NRGBA{
						R: 0xFF,
						G: 0x00,
						B: 0x00,
						A: 0xFF,
					}
					tempExecutionStatusRectangle.StrokeWidth = 2

				*/

				/*
					var id string
					id = "e12c6be8-614c-4379-b482-165ff18dd68d"
					log.Fatalf(fmt.Sprintf("Couldn't find TestInstruction '%s' in TestInstructionsExecution-data for TIE '%s'. ID='%s'",
						previewObject.GetTestInstructionUuid,
						testCaseExecutionThatIsShownInPreview,
						id))
				*/
			} else {

				statusId = detailedExecutionsModel.
					ExecutionStatusColorNameToNumberMap[foundTestInstructionExecutionStatusPreviewValues.
					TestInstructionExecutionStatus.String()[4:]].ExecutionStatusNumber
				statusBackgroundColor = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].BackgroundColor
				tempExecutionStatusRectangle.FillColor = statusBackgroundColor

				useStroke = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].UseStroke
				if useStroke == true {
					statusStrokeColor = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)].StrokeColor
					tempExecutionStatusRectangle.StrokeColor = statusStrokeColor
					tempExecutionStatusRectangle.StrokeWidth = 2
				}
			}

			// When no background color
			//if statusBackgroundColor.R+statusBackgroundColor.G+statusBackgroundColor.B+statusBackgroundColor.A == 0 {

			//					clickable.Alignment = fyne.TextAlignCenter
			//					clickable.textInsteadOfLabel.Hide()
			//					clickable.Show()

			// Create the container containing the TestInstruction
			var tempTestInstructionContainer *fyne.Container
			tempTestInstructionContainer = container.NewHBox(
				tempExecutionStatusRectangle,
				tempIndentationLevelRectangle,
				testInstructionColorRectangle,
				tempTestInstructionNameWidget)

			// Add the TestInstructionContainer to the main Area
			testCaseExecutionMainAreaForPreviewContainer.Add(tempTestInstructionContainer)

			/*

				// Attributes if there are any
				if len(previewObject.GetTestInstructionAttributes()) > 0 {

					// Create the Indentation rectangle for the attributes
					var tempIndentationLevelRectangleForAttributes *canvas.Rectangle
					tempIndentationLevelRectangleForAttributes = canvas.NewRectangle(color.Transparent)

					// Resize the Indentation rectangle
					tempIndentationLevelRectangleForAttributes.SetMinSize(fyne.Size{
						Width:  float32(10 * (previewObject.GetIndentationLevel() + 10)),
						Height: 2,
					})
					tempIndentationLevelRectangleForAttributes.Resize(fyne.Size{
						Width:  float32(10 * (previewObject.GetIndentationLevel() + 10)),
						Height: 2,
					})

					// Create the container for the attributes
					var testInstructionAttributesContainer *fyne.Container
					testInstructionAttributesContainer = container.New(layout.NewFormLayout())

					// Loop attributes and add to container
					for _, attribute := range previewObject.TestInstructionAttributes {

						// Create label for attribute
						var attributeLabel *widget.Label
						attributeLabel = widget.NewLabel(attribute.AttributeName)

						// Create value for attribute
						var attributeValue *widget.Label
						attributeValue = widget.NewLabel(attribute.AttributeValue)

						// Add label and value ro the attribute container
						testInstructionAttributesContainer.Add(attributeLabel)
						testInstructionAttributesContainer.Add(attributeValue)

					}

					// Create the container containing the Attributes
					var tempTestInstructionAttributesContainer *fyne.Container
					tempTestInstructionAttributesContainer = container.NewHBox(
						tempIndentationLevelRectangleForAttributes, testInstructionAttributesContainer)

					// Add the TestInstructionContainerContainer to the main Area
					testCaseExecutionMainAreaForPreviewContainer.Add(tempTestInstructionAttributesContainer)
				}

			*/

		default:
			log.Fatalf("Unknown 'previewObject.TestCaseStructureObjectType' which never should happen; %s", previewObject.TestCaseStructureObjectType.String())
		}
	}

	testCaseMainAreaForPreviewBorderContainer := container.NewBorder(nil, nil, nil, nil, testCaseExecutionMainAreaForPreviewContainer)
	testCaseMainAreaForPreviewScrollContainer := container.NewScroll(testCaseMainAreaForPreviewBorderContainer)

	// Create the container used for the TestCase, with TIC, TI and Attributes
	//testCasePreviewScrollContainer = container.NewScroll(tempContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestCase Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}

	testCaseExecutionPreviewContainer.Objects[0] = container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testCaseExecutionPreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testCaseExecutionPreviewBottomContainer), nil, nil,
		testCaseMainAreaForPreviewScrollContainer)

	// Refresh the 'testCaseExecutionPreviewContainer'
	testCaseExecutionPreviewContainer.Refresh()

}
