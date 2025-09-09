package listTestCasesUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"FenixTesterGui/testCases/listTestCasesModel"
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"image/color"
	"image/png"
	"log"
)

// GenerateTestCasePreviewContainer
// Generates the PreViewContainer for the TestCase
func (listTestCaseUIObject *ListTestCaseUIStruct) GenerateTestCasePreviewContainer(
	testCaseUuid string,
	testCasePreviewContainer *fyne.Container,
	testCasesModel *testCaseModel.TestCasesModelsStruct,
	fromTestCaseList bool) {

	var testCasePreviewTopContainer *fyne.Container
	var testCasePreviewBottomContainer *fyne.Container
	//var testCasePreviewScrollContainer *container.Scroll
	var testCaseMainAreaForPreviewContainer *fyne.Container

	var err error

	var tempTestCasePreviewStructureMessage *fenixGuiTestCaseBuilderServerGrpcApi.TestCasePreviewStructureMessage

	// Get Data for the Preview
	tempTestCasePreviewStructureMessage = listTestCasesModel.TestCasesThatCanBeEditedByUserMap[testCaseUuid].TestCasePreview.TestCasePreview

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

	if fromTestCaseList == true {

		testCasePreviewContainer.Objects[0] = container.NewBorder(
			container.NewVBox(container.NewCenter(tempTopHeaderLabel), testCasePreviewTopContainer, widget.NewSeparator()),
			container.NewVBox(widget.NewSeparator(), testCasePreviewBottomContainer), nil, nil,
			testCaseMainAreaForPreviewScrollContainer)

	} else {

		testCasePreviewContainer.Objects[0] = container.NewBorder(
			container.NewVBox(container.NewHBox(tempTopHeaderLabel), testCasePreviewTopContainer, widget.NewSeparator()),
			container.NewVBox(widget.NewSeparator(), testCasePreviewBottomContainer), nil, nil,
			testCaseMainAreaForPreviewScrollContainer)
	}

	// Refresh the 'testCasePreviewContainer'
	testCasePreviewContainer.Refresh()

}
