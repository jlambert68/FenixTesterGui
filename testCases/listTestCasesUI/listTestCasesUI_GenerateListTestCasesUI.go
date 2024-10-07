package listTestCasesUI

import (
	sharedCode "FenixTesterGui/common_code"
	detailedTestCaseExecutionsUI "FenixTesterGui/executions/detailedExecutionsUI"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesModel"
	"bytes"
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image"
	"image/color"
	"image/png"
	"log"
	"strconv"
	"time"
)

//go:embed resources/TIC-Horizontal_32x32.png
var tic_parallellImage []byte
var imageData_tic_parallellImage image.Image

//go:embed resources/TIC-Vertical_32x32.png
var tic_serialImage []byte
var imageData_tic_serialImage image.Image

// Create the UI used for list all TestCases that the User can edit
func GenerateListTestCasesUI(testCasesModel *testCaseModel.TestCasesModelsStruct) (listTestCasesUI fyne.CanvasObject) {

	//var testCaseTable *widget.Table

	var tempTestCaseListAndTestCasePreviewSplitContainer *container.Split

	var testCasesListContainer *fyne.Container
	var testCasesListScrollContainer *container.Scroll
	var statisticsContainer *fyne.Container
	var executionColorPaletteContainer *fyne.Container
	var statisticsAndColorPaletteContainer *fyne.Container

	var loadTestCaseFromDataBaseButton *widget.Button
	var loadTestCaseFromDataBaseFunction func()
	var filterTestCasesButton *widget.Button
	var filterTestCasesButtonFunction func()
	var clearFiltersButton *widget.Button
	var clearFiltersButtonFunction func()
	var buttonsContainer *fyne.Container

	var numberOfTestCasesAfterLocalFilterLabel *widget.Label
	var numberOfTestCasesRetrievedFromDatabaseLabel *widget.Label

	var filterAndButtonsContainer *fyne.Container

	// Define the function to be executed to load TestCases from that Database that the user can edit
	loadTestCaseFromDataBaseFunction = func() {
		fmt.Println("'loadTestCaseFromDataBaseButton' was pressed")
		listTestCasesModel.LoadTestCaseThatCanBeEditedByUser(testCasesModel, time.Now().Add(-time.Hour*1000), time.Now().Add(-time.Hour*1000))
		filterTestCasesButtonFunction()
	}

	// Define the 'loadTestCaseFromDataBaseButton'
	loadTestCaseFromDataBaseButton = widget.NewButton("Load TestCases from Database", loadTestCaseFromDataBaseFunction)

	// Define the function to be executed to filter TestCases that the user can edit
	filterTestCasesButtonFunction = func() {
		fmt.Println("'filterTestCasesButton' was pressed")
		loadTestCaseListTableTable(testCasesModel)
		calculateAndSetCorrectColumnWidths()
		updateTestCasesListTable(testCasesModel)

		// Update the number TestCases in the list
		var numberOfRowsAsString string
		numberOfRowsAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesAfterLocalFilters.Set(
			fmt.Sprintf("Number of TestCases after local filters was applied: %s",
				numberOfRowsAsString))

		// Update the number TestCases retrieved from Database
		var numberOfRowsFromDatabaseAsString string
		numberOfRowsFromDatabaseAsString = strconv.Itoa(len(testCaseListTableTable))
		numberOfTestCasesInTheDatabaseSearch.Set(
			fmt.Sprintf("Number of TestCases retrieved from the Database: %s",
				numberOfRowsFromDatabaseAsString))

	}

	// Define the 'filterTestCasesButton'
	filterTestCasesButton = widget.NewButton("Filter TestCases", filterTestCasesButtonFunction)

	// Define the function to be executed to list TestCases that the user can edit
	clearFiltersButtonFunction = func() {
		fmt.Println("'clearFiltersButtonFunction' was pressed")
	}

	// Define the 'filterTestCasesButton'
	clearFiltersButton = widget.NewButton("Clear all search filters", clearFiltersButtonFunction)

	// Add the buttons to the buttonsContainer
	buttonsContainer = container.NewHBox(loadTestCaseFromDataBaseButton, filterTestCasesButton, clearFiltersButton)

	// Add objects to the 'filterAndButtonsContainer'
	filterAndButtonsContainer = container.NewVBox(buttonsContainer)

	// Generate the ExecutionColorPaletteContainer
	executionColorPaletteContainer = detailedTestCaseExecutionsUI.GenerateExecutionColorPalette()

	// Initiate the Table
	generateTestCasesListTable(testCasesModel)
	testCaseTableContainer := container.NewBorder(nil, nil, nil, nil, testCaseListTable)

	// Create the Scroll container for the List
	testCasesListScrollContainer = container.NewScroll(testCaseTableContainer)

	// Create the label used for showing number of TestCases in the local filter
	numberOfTestCasesAfterLocalFilters = binding.NewString()
	_ = numberOfTestCasesAfterLocalFilters.Set("No TestCases in the List")
	numberOfTestCasesAfterLocalFilterLabel = widget.NewLabelWithData(numberOfTestCasesAfterLocalFilters)

	// Create the label used for showing number of TestCases retrieved from the Database
	numberOfTestCasesInTheDatabaseSearch = binding.NewString()
	_ = numberOfTestCasesInTheDatabaseSearch.Set("No TestCases retrieved from the Database")
	numberOfTestCasesRetrievedFromDatabaseLabel = widget.NewLabelWithData(numberOfTestCasesInTheDatabaseSearch)

	// Initiate 'statisticsContainer'
	statisticsContainer = container.NewHBox(numberOfTestCasesAfterLocalFilterLabel, numberOfTestCasesRetrievedFromDatabaseLabel)

	statisticsAndColorPaletteContainer = container.NewVBox(executionColorPaletteContainer, statisticsContainer)

	// Add 'testCasesListScrollContainer' to 'testCasesListContainer'
	testCasesListContainer = container.NewBorder(filterAndButtonsContainer, statisticsAndColorPaletteContainer, nil, nil, testCasesListScrollContainer)
	testCasesListScrollContainer2 := container.NewScroll(testCasesListContainer)

	// Create the Temporary container that should be shown
	temporaryContainer := container.NewCenter(widget.NewLabel("Select a TestCase to get the Preview"))

	testCasePreviewContainer = container.NewBorder(nil, nil, nil, nil, temporaryContainer)

	// Generate the container for the Preview, 'testCasePreviewContainer'
	testCasePreviewContainerScroll = container.NewScroll(testCasePreviewContainer)

	tempTestCaseListAndTestCasePreviewSplitContainer = container.NewHSplit(testCasesListScrollContainer2, testCasePreviewContainerScroll)
	tempTestCaseListAndTestCasePreviewSplitContainer.Offset = 0.75

	TestCaseListAndTestCasePreviewSplitContainer = tempTestCaseListAndTestCasePreviewSplitContainer

	return tempTestCaseListAndTestCasePreviewSplitContainer
}

func GenerateTestCasePreviewContainer(
	testCaseUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var testCasePreviewTopContainer *fyne.Container
	var testCasePreviewBottomContainer *fyne.Container
	//var testCasePreviewScrollContainer *container.Scroll
	var testCaseMainAreaForPreviewContainer *fyne.Container

	var err error

	var tempTestCasePreviewStructureMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage

	// Get Data for the Preview
	tempTestCasePreviewStructureMessage = testCasesModel.TestCasesThatCanBeEditedByUserMap[testCaseUuid].TestCasePreview.TestCasePreview

	// Create the Top container
	testCasePreviewTopContainer = container.New(layout.NewFormLayout())

	// Add TestCaseName to Top container
	tempTestCaseNameLabel := widget.NewLabel("TestCaseName:")
	tempTestCaseNameLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewTopContainer.Add(tempTestCaseNameLabel)
	testCasePreviewTopContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetTestCaseName()))

	// Add TestCaseOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewTopContainer.Add(tempOwnerDomainLabel)
	testCasePreviewTopContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetDomainThatOwnTheTestCase()))

	// Add Description Top container
	tempTestCaseDescription := widget.NewLabel("Description:")
	tempTestCaseDescription.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewTopContainer.Add(tempTestCaseDescription)
	testCasePreviewTopContainer.Add(widget.NewRichTextWithText(tempTestCasePreviewStructureMessage.GetTestCaseDescription()))

	// Create the Bottom container
	testCasePreviewBottomContainer = container.New(layout.NewFormLayout())

	// Add ComplexTextualDescription to Bottom container
	tempComplexTextualDescriptionLabel := widget.NewLabel("ComplexTextualDescription:")
	tempComplexTextualDescriptionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewBottomContainer.Add(tempComplexTextualDescriptionLabel)
	testCasePreviewBottomContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetComplexTextualDescription()))

	// Add TestCaseVersion to Bottom container
	tempTestCaseVersionLabel := widget.NewLabel("TestCaseVersion:")
	tempTestCaseVersionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewBottomContainer.Add(tempTestCaseVersionLabel)
	testCasePreviewBottomContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetTestCaseVersion()))

	// Add LastSavedByUserOnComputer to Bottom container
	tempLastSavedByUserOnComputerLabel := widget.NewLabel("Last saved by user (on computer)::")
	tempLastSavedByUserOnComputerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewBottomContainer.Add(tempLastSavedByUserOnComputerLabel)
	testCasePreviewBottomContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetLastSavedByUserOnComputer()))

	// Add LastSavedByUserGCPAuthorization to Bottom container
	tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("Last saved by GCP authenticated user:")
	tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
	testCasePreviewBottomContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetLastSavedByUserGCPAuthorization()))

	// Add LastSavedTimeStamp to Bottom container
	tempLastSavedTimeStamp := widget.NewLabel("Last saved TimeStamp:")
	tempLastSavedTimeStamp.TextStyle = fyne.TextStyle{Bold: true}
	testCasePreviewBottomContainer.Add(tempLastSavedTimeStamp)
	testCasePreviewBottomContainer.Add(widget.NewLabel(tempTestCasePreviewStructureMessage.GetLastSavedTimeStamp()))

	// Create the area used for TIC, TI and the attributes
	testCaseMainAreaForPreviewContainer = container.NewVBox()

	// Loop the preview objects and to container
	for _, previewObject := range tempTestCasePreviewStructureMessage.TestCaseStructureObjects {

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

		// Decide what type of object
		switch previewObject.TestCaseStructureObjectType {

		case fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:

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
				tempIndentationLevelRectangle, serialOrParallelRectangleImage, tempTestInstructionContainerNameWidget)

			// Add the TestInstructionContainerContainer to the main Area
			testCaseMainAreaForPreviewContainer.Add(tempTestInstructionContainerContainer)

		case fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage_TestInstruction:

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

			// Create the container containing the TestInstruction
			var tempTestInstructionContainer *fyne.Container
			tempTestInstructionContainer = container.NewHBox(
				tempIndentationLevelRectangle, testInstructionColorRectangle, tempTestInstructionNameWidget)

			// Add the TestInstructionContainer to the main Area
			testCaseMainAreaForPreviewContainer.Add(tempTestInstructionContainer)

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
				testCaseMainAreaForPreviewContainer.Add(tempTestInstructionAttributesContainer)
			}

		default:
			log.Fatalf("Unknown 'previewObject.TestCaseStructureObjectType' which never should happen; %s", previewObject.TestCaseStructureObjectType.String())
		}
	}

	testCaseMainAreaForPreviewBorderContainer := container.NewBorder(nil, nil, nil, nil, testCaseMainAreaForPreviewContainer)
	testCaseMainAreaForPreviewScrollContainer := container.NewScroll(testCaseMainAreaForPreviewBorderContainer)

	// Create the container used for the TestCase, with TIC, TI and Attributes
	//testCasePreviewScrollContainer = container.NewScroll(tempContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestCase Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}

	testCasePreviewContainer.Objects[0] = container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testCasePreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testCasePreviewBottomContainer), nil, nil,
		testCaseMainAreaForPreviewScrollContainer)

	// Refresh the 'testCasePreviewContainer'
	testCasePreviewContainer.Refresh()

}
