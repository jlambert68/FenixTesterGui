package listTestSuitesUI

import (
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testSuites/listTestSuitesModel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

// GenerateTestSuitePreviewContainer
// Generates the PreViewContainer for the TestSuite
func (listTestSuiteUIObject *ListTestSuiteUIStruct) GenerateTestSuitePreviewContainer(
	testSuiteUuid string,
	testCasesModel *testCaseModel.TestCasesModelsStruct) {

	var testSuitePreviewTopContainer *fyne.Container
	var testSuitePreviewBottomContainer *fyne.Container
	//var testSuitePreviewScrollContainer *container.Scroll
	var testSuiteMainAreaForPreviewContainer *fyne.Container

	//var err error

	var tempTestSuitePreviewStructureMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestSuitePreviewStructureMessage

	// Get Data for the Preview
	tempTestSuitePreviewStructureMessage = listTestSuitesModel.TestSuitesThatCanBeEditedByUserMap[testSuiteUuid].TestSuitePreview.TestSuitePreview

	// Create the Top container
	testSuitePreviewTopContainer = container.New(layout.NewFormLayout())

	// Add TestSuiteName to Top container
	tempTestSuiteNameLabel := widget.NewLabel("TestSuiteName:")
	tempTestSuiteNameLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempTestSuiteNameLabel)
	testSuitePreviewTopContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetTestSuiteName()))

	// Add TestSuiteOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempOwnerDomainLabel)
	testSuitePreviewTopContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetDomainNameThatOwnTheTestSuite()))

	// Add Description Top container
	tempTestSuiteDescription := widget.NewLabel("Description:")
	tempTestSuiteDescription.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewTopContainer.Add(tempTestSuiteDescription)
	testSuitePreviewTopContainer.Add(widget.NewRichTextWithText(tempTestSuitePreviewStructureMessage.GetTestSuiteDescription()))

	// Create the Bottom container
	testSuitePreviewBottomContainer = container.New(layout.NewFormLayout())

	// Add TestSuiteVersion to Bottom container
	tempTestSuiteVersionLabel := widget.NewLabel("TestSuiteVersion:")
	tempTestSuiteVersionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempTestSuiteVersionLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetTestSuiteVersion()))

	// Add LastSavedByUserOnComputer to Bottom container
	tempLastSavedByUserOnComputerLabel := widget.NewLabel("Last saved by user (on computer)::")
	tempLastSavedByUserOnComputerLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedByUserOnComputerLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedByUserOnComputer()))

	// Add LastSavedByUserGCPAuthorization to Bottom container
	tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("Last saved by GCP authenticated user:")
	tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedByUserGCPAuthorization()))

	// Add LastSavedTimeStamp to Bottom container
	tempLastSavedTimeStamp := widget.NewLabel("Last saved TimeStamp:")
	tempLastSavedTimeStamp.TextStyle = fyne.TextStyle{Bold: true}
	testSuitePreviewBottomContainer.Add(tempLastSavedTimeStamp)
	testSuitePreviewBottomContainer.Add(widget.NewLabel(tempTestSuitePreviewStructureMessage.GetLastSavedTimeStamp()))

	// Create the area used for TIC, TI and the attributes
	testSuiteMainAreaForPreviewContainer = container.NewVBox()

	// Loop the preview objects and add to container
	/*
		for _, testCasePreview := range tempTestSuitePreviewStructureMessage.GetTestSuiteStructureObjects().TestCasePreViews {

			for _, previewObject := range testCasePreview.TestCasePreview.GetTestCaseStructureObjects() {

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
					testSuiteMainAreaForPreviewContainer.Add(tempTestInstructionContainerContainer)

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
					testSuiteMainAreaForPreviewContainer.Add(tempTestInstructionContainer)

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
						testSuiteMainAreaForPreviewContainer.Add(tempTestInstructionAttributesContainer)
					}

				default:
					log.Fatalf("Unknown 'previewObject.TestCaseStructureObjectType' which never should happen; %s", previewObject.TestCaseStructureObjectType.String())
				}
			}
		}

	*/

	testSuiteMainAreaForPreviewBorderContainer := container.NewBorder(nil, nil, nil, nil, testSuiteMainAreaForPreviewContainer)
	testSuiteMainAreaForPreviewScrollContainer := container.NewScroll(testSuiteMainAreaForPreviewBorderContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestSuite Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}

	listTestSuiteUIObject.testSuitePreviewContainer.Objects[0] = container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testSuitePreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testSuitePreviewBottomContainer), nil, nil,
		testSuiteMainAreaForPreviewScrollContainer)

	// Refresh the 'testSuitePreviewContainer'
	listTestSuiteUIObject.testSuitePreviewContainer.Refresh()

}
