package listTestSuiteExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/testSuiteExecutions/testSuiteExecutionsModel"
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"log"
	"strconv"
)

func (testSuiteInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) GenerateTestSuiteExecutionPreviewContainer(
	testSuiteExecutionUuid string,
	testSuiteExecutionVersion uint32,
	testSuiteExecutionsModelRef *testSuiteExecutionsModel.TestSuiteExecutionsModelStruct,
	openedTestSuiteExecutionFrom openedTestSuiteExecutionFromType,
	currentWindowPtr *fyne.Window) {

	var currentWindow fyne.Window
	currentWindow = *currentWindowPtr

	var testSuiteExecutionPreviewTopContainer *fyne.Container
	var testSuiteExecutionPreviewBottomContainer *fyne.Container
	//var testCasePreviewScrollContainer *container.Scroll
	var testSuiteExecutionMainAreaForPreviewContainer *fyne.Container

	var err error
	var existInMap bool
	var foundValue bool
	var runTimeValueExists bool

	// Lock Map
	testSuiteExecutionAttributesForPreviewMapMutex.Lock()

	// Unlock map
	defer testSuiteExecutionAttributesForPreviewMapMutex.Unlock()

	// Clear out 'testSuiteExecutionAttributesForPreviewMapPtr'
	var testSuiteExecutionAttributesForPreviewMap map[testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct
	testSuiteExecutionAttributesForPreviewMap = make(map[testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct)

	testSuiteExecutionAttributesForPreviewMapPtr = &testSuiteExecutionAttributesForPreviewMap

	// Verify that number Headers match number of columns, constant 'numberColumnsInTestSuiteExecutionsListUI'
	if len(testSuiteExecutionsListTableHeader) != numberColumnsInTestSuiteExecutionsListUI {
		log.Fatalln(fmt.Sprintf("Number of elements in 'tempRowslice' missmatch contant 'numberColumnsInTestSuiteExecutionsListUI'. %d vs %d. ID: %s",
			testSuiteExecutionsListTableHeader,
			numberColumnsInTestSuiteExecutionsListUI,
			"c2b8a13c-ec20-46c2-adf9-965247732e07"))
	}

	// Get Data for the Preview
	var tempTestSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage

	// Can preview be found in Map for "One TestSuiteExecution per TestCase" or "All TestSuiteExecutions per TestCase"
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

	case AllExecutionsForOneTestSuite:
		tempTestSuiteExecutionsListMessage, _ = testSuiteExecutionsModelRef.GetSpecificTestSuiteExecutionForOneTestSuiteUuid(
			testSuiteExecutionsModel.TestSuiteUuidType(selectedTestSuiteExecutionObjected.
				allExecutionsFoOneTestSuiteListObject.testSuiteUuidForTestSuiteExecutionThatIsShownInPreview),
			testSuiteExecutionsModel.TestSuiteExecutionUuidType(selectedTestSuiteExecutionObjected.
				allExecutionsFoOneTestSuiteListObject.testSuiteExecutionUuidThatIsShownInPreview))

	case OneExecutionPerTestSuite:
		tempTestSuiteExecutionsListMessage, _ = testSuiteExecutionsModelRef.ReadFromTestSuiteExecutionsMap(
			testSuiteExecutionsModel.TestSuiteExecutionUuidType(testSuiteExecutionUuid))

	case NotDefined:

		tempTestSuiteExecutionsListMessage = &fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage{}

	}

	// Read from the TestSuiteExecutions-Map to Get object holding Logs, RunTimeAtributes and ...

	//var detailedTestSuiteExecutionsMapObject testSuiteExecutionsModel.DetailedTestSuiteExecutionsMapObjectStruct

	var detailedTestSuiteExecutionsObjectsMapPtr *map[testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType]*testSuiteExecutionsModel.DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionsObjectsMap map[testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType]*testSuiteExecutionsModel.DetailedTestSuiteExecutionsMapObjectStruct
	detailedTestSuiteExecutionsObjectsMapPtr = testSuiteExecutionsModelRef.DetailedTestSuiteExecutionsObjectsMapPtr
	detailedTestSuiteExecutionsObjectsMap = *detailedTestSuiteExecutionsObjectsMapPtr

	var detailedTestSuiteExecutionsObjectPtr *testSuiteExecutionsModel.DetailedTestSuiteExecutionsMapObjectStruct
	var detailedTestSuiteExecutionMapKey testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType
	detailedTestSuiteExecutionMapKey = testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType(
		testSuiteExecutionUuid + strconv.Itoa(int(testSuiteExecutionVersion)))

	detailedTestSuiteExecutionsObjectPtr, existInMap = detailedTestSuiteExecutionsObjectsMap[detailedTestSuiteExecutionMapKey]

	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                               "b245a7c7-ac7c-4958-a4ab-70d8a672bcd9",
			"detailedTestSuiteExecutionMapKey": detailedTestSuiteExecutionMapKey,
		}).Fatalln("Couldn't find TestSuiteExecution in 'detailedTestSuiteExecutionMap', should never happen!")
	}

	// Extract the run-time variables map
	var runTimeUpdatedAttributesMapPtr *map[testSuiteExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[testSuiteExecutionsModel.AttributeNameMapKeyType]testSuiteExecutionsModel.
		RunTimeUpdatedAttributeValueType
	var testInstructionExecutionsRunTimeUpdatedAttributesMap map[testSuiteExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[testSuiteExecutionsModel.AttributeNameMapKeyType]testSuiteExecutionsModel.
		RunTimeUpdatedAttributeValueType

	runTimeUpdatedAttributesMapPtr = detailedTestSuiteExecutionsObjectPtr.RunTimeUpdatedAttributesMapPtr
	testInstructionExecutionsRunTimeUpdatedAttributesMap = *runTimeUpdatedAttributesMapPtr

	// Create the Top container
	testSuiteExecutionPreviewTopContainer = container.New(layout.NewFormLayout())

	// Create the ExecutionStatus Rectangle for TestSuiteExecution-status
	var tempTestSuiteExecutionStatusRectangle *canvas.Rectangle
	tempTestSuiteExecutionStatusRectangle = canvas.NewRectangle(color.Transparent)

	// Resize the ExecutionStatus rectangle
	tempTestSuiteExecutionStatusRectangle.SetMinSize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHeight,
	})
	tempTestSuiteExecutionStatusRectangle.Resize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHeight,
	})

	// Set correct color on ExecutionStatus Rectangle
	var statusId uint8

	// Extract TestSuiteExecution-status
	var tempTestSuiteExecutionStatusEnum string
	tempTestSuiteExecutionStatusEnum = tempTestSuiteExecutionsListMessage.GetTestSuiteExecutionStatus().String()[4:]

	statusId = detailedExecutionsModel.
		ExecutionStatusColorNameToNumberMap[tempTestSuiteExecutionStatusEnum].ExecutionStatusNumber

	var executionStatusColorMapObjectForTestSuiteExecution detailedExecutionsModel.ExecutionStatusColorMapStruct
	executionStatusColorMapObjectForTestSuiteExecution, existInMap = detailedExecutionsModel.ExecutionStatusColorMap[int32(statusId)]
	if existInMap == false {
		// No matching Status color exist due to that TestInstruction exists in ExecutionQueue
		// 'INITIATED = 1'
		executionStatusColorMapObjectForTestSuiteExecution, _ = detailedExecutionsModel.ExecutionStatusColorMap[1]
		/*
			tempTestSuiteExecutionStatusRectangle.StrokeWidth = 2
			tempTestSuiteExecutionStatusRectangle.StrokeColor = color.NRGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xFF,
			}
		*/

	} else {
		// Status color found
		tempTestSuiteExecutionStatusRectangle.FillColor = executionStatusColorMapObjectForTestSuiteExecution.BackgroundColor

		if executionStatusColorMapObjectForTestSuiteExecution.UseStroke == true {
			tempTestSuiteExecutionStatusRectangle.StrokeWidth = 2
			tempTestSuiteExecutionStatusRectangle.StrokeColor = executionStatusColorMapObjectForTestSuiteExecution.StrokeColor
		}
	}

	// Add TestSuiteName to Top container
	tempTestCaseNameLabel := widget.NewLabel("TestSuiteName:")
	tempTestCaseNameLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewTopContainer.Add(tempTestCaseNameLabel)
	copyableTestCaseNameLabel := newCopyableLabel(tempTestSuiteExecutionsListMessage.GetTestSuiteName(), true)
	testSuiteExecutionPreviewTopContainer.Add(copyableTestCaseNameLabel)

	// Add TestSuiteExecutionStatus
	tempTestSuiteExecutionStatusLabel := widget.NewLabel("TestSuiteExecution-status:")
	tempTestSuiteExecutionStatusLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewTopContainer.Add(tempTestSuiteExecutionStatusLabel)

	var testSuiteExecutionStatusStackContainer *fyne.Container
	var testSuiteExecutionStatusHBoxContainer *fyne.Container
	var testSuiteExecutionStatusOuterHBoxContainer *fyne.Container
	testSuiteExecutionStatusStackContainer = container.NewStack()
	testSuiteExecutionStatusHBoxContainer = container.NewHBox(
		widget.NewLabel(" "),
		canvas.NewText(executionStatusColorMapObjectForTestSuiteExecution.ExecutionStatusName, color.Black),
		widget.NewLabel(" "))
	testSuiteExecutionStatusStackContainer.Add(tempTestSuiteExecutionStatusRectangle)
	testSuiteExecutionStatusStackContainer.Add(testSuiteExecutionStatusHBoxContainer)
	testSuiteExecutionStatusOuterHBoxContainer = container.NewHBox(
		testSuiteExecutionStatusStackContainer,
		layout.NewSpacer())

	testSuiteExecutionPreviewTopContainer.Add(testSuiteExecutionStatusOuterHBoxContainer)

	// Add TestSuiteOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewTopContainer.Add(tempOwnerDomainLabel)
	copyableDomainThatOwnTheTestCaseLabel := newCopyableLabel(
		tempTestSuiteExecutionsListMessage.TestSuitePreview.GetTestSuitePreview().GetDomainNameThatOwnTheTestSuite(), true)
	testSuiteExecutionPreviewTopContainer.Add(copyableDomainThatOwnTheTestCaseLabel)

	// Add empty row
	testSuiteExecutionPreviewTopContainer.Add(widget.NewLabel(""))
	testSuiteExecutionPreviewTopContainer.Add(widget.NewLabel(""))

	// Add TestCase
	tempTestCaseLabel := widget.NewLabel("TestCase with execution status:")
	tempTestCaseLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewTopContainer.Add(tempTestCaseLabel)
	testSuiteExecutionPreviewTopContainer.Add(widget.NewLabel(""))

	/*
		// Add Description Top container
		tempTestCaseDescription := widget.NewLabel("Description:")
		tempTestCaseDescription.TextStyle = fyne.TextStyle{Bold: true}
		testSuiteExecutionPreviewTopContainer.Add(tempTestCaseDescription)
		testSuiteExecutionPreviewTopContainer.Add(widget.NewRichTextWithText(tempTestSuiteExecutionsListMessage.GetTestCaseDescription()))

		// Create the Bottom container
		testSuiteExecutionPreviewBottomContainer = container.New(layout.NewFormLayout())

		// Add ComplexTextualDescription to Bottom container
		tempComplexTextualDescriptionLabel := widget.NewLabel("ComplexTextualDescription:")
		tempComplexTextualDescriptionLabel.TextStyle = fyne.TextStyle{Bold: true}
		testSuiteExecutionPreviewBottomContainer.Add(tempComplexTextualDescriptionLabel)
		testSuiteExecutionPreviewBottomContainer.Add(widget.NewLabel(tempTestSuiteExecutionsListMessage.GetComplexTextualDescription()))


	*/

	// Create the Bottom container
	testSuiteExecutionPreviewBottomContainer = container.New(layout.NewFormLayout())

	// Add TestSuiteExecutionVersion to Bottom container
	tempTestSuiteExecutionVersionLabel := widget.NewLabel("TestSuiteExecutionVersion:")
	tempTestSuiteExecutionVersionLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewBottomContainer.Add(tempTestSuiteExecutionVersionLabel)
	testSuiteExecutionPreviewBottomContainer.Add(widget.NewLabel(strconv.Itoa(int(tempTestSuiteExecutionsListMessage.
		GetTestSuiteExecutionVersion()))))

	// Add ExecutionStartTimeStamp to Bottom container
	tempExecutionStartTimeStampLabel := widget.NewLabel("TestCase Execution Start TimeStamp:")
	tempExecutionStartTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewBottomContainer.Add(tempExecutionStartTimeStampLabel)
	copyableExecutionStartTimeStampLabel := newCopyableLabel(
		tempTestSuiteExecutionsListMessage.
			GetExecutionStartTimeStamp().AsTime().String(), true)
	testSuiteExecutionPreviewBottomContainer.Add(copyableExecutionStartTimeStampLabel)

	// Add ExecutionStopTimeStamp to Bottom container
	tempExecutionStopTimeStampLabel := widget.NewLabel("TestCase Execution Stop TimeStamp:")
	tempExecutionStopTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewBottomContainer.Add(tempExecutionStopTimeStampLabel)
	copyableExecutionStopTimeStampLabel := newCopyableLabel(
		tempTestSuiteExecutionsListMessage.
			GetExecutionStopTimeStamp().AsTime().String(), true)
	testSuiteExecutionPreviewBottomContainer.Add(copyableExecutionStopTimeStampLabel)

	// Add ExecutionStatusUpdateTimeStamp to Bottom container
	tempExecutionStatusUpdateTimeStampLabel := widget.NewLabel("TestCase Execution Status Update TimeStamp:")
	tempExecutionStatusUpdateTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testSuiteExecutionPreviewBottomContainer.Add(tempExecutionStatusUpdateTimeStampLabel)
	copyableExecutionStatusUpdateTimeStampLabel := newCopyableLabel(
		tempTestSuiteExecutionsListMessage.
			GetExecutionStatusUpdateTimeStamp().AsTime().String(), true)
	testSuiteExecutionPreviewBottomContainer.Add(copyableExecutionStatusUpdateTimeStampLabel)

	// Add LastSavedByUserGCPAuthorization to Bottom container
	/*
		tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("Last saved by this GCP-user:")
		tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
		testSuiteExecutionPreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
		copyableExecutionStatusUpdateTimeStampLabel := newCopyableLabel(
			tempTestSuiteExecutionsListMessage..
				Get().AsTime().String(),true)
		testSuiteExecutionPreviewBottomContainer.Add(copyableExecutionStatusUpdateTimeStampLabel)

	*/
	// Create the area used for TIC, TI and the attributes
	testSuiteExecutionMainAreaForPreviewContainer = container.NewVBox()

	// Check if there is TestCase Preview information
	if tempTestSuiteExecutionsListMessage.GetTestCasePreview() == nil {

		testSuiteExecutionMainAreaForPreviewContainer.Add(widget.NewLabel("No Preview information found in database!"))

	} else {

		// Loop the preview objects and to container
		for previewObjectIndex, previewObject := range tempTestSuiteExecutionsListMessage.TestCasesPreviews.TestCasePreviews {

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
				Height: testCaseExecutionStatusRectangleHeight,
			})
			tempExecutionStatusRectangle.Resize(fyne.Size{
				Width:  testCaseExecutionStatusRectangleWidth,
				Height: testCaseExecutionStatusRectangleHeight,
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
					serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testSuiteNodeRectangleSize-4),
						float32(testSuiteNodeRectangleSize-4)))
					serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testSuiteNodeRectangleSize-4),
						float32(testSuiteNodeRectangleSize-4)))

				} else {
					// Convert the byte slice into an image.Image object
					if imageData_tic_parallellImage == nil {
						imageData_tic_parallellImage, err = png.Decode(bytes.NewReader(tic_parallellImage))
						if err != nil {
							log.Fatalf("Failed to decode image: %v", err)
						}
					}
					serialOrParallelRectangleImage = canvas.NewImageFromImage(imageData_tic_parallellImage)
					serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testSuiteNodeRectangleSize-4),
						float32(testSuiteNodeRectangleSize-4)))
					serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testSuiteNodeRectangleSize-4),
						float32(testSuiteNodeRectangleSize-4)))

				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var tempTIEAttributesContainerMapKey testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				tempTIEAttributesContainerMapKey = testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionContainerUuid())

				// Create the Name for the TestInstructionContainer
				// var tempTestInstructionContainerNameWidget *widget.Label
				//tempTestInstructionContainerNameWidget = widget.NewLabel(previewObject.GetTestInstructionContainerName())
				var tempTestInstructionContainerNameWidget *clickableTInTICNameLabelInPreviewStruct
				tempTestInstructionContainerNameWidget = newClickableTestInstructionNameLabelInPreview(
					previewObject.GetTestInstructionContainerName(),
					testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionUuid+
						strconv.Itoa(int(testSuiteExecutionVersion))),
					tempTIEAttributesContainerMapKey,
					nil,
					nil,
					labelIsTestInstructionContainer,
					testSuiteInstructionPreViewObjectRef)

				// Create the container containing the TestInstructionContainer
				var tempTestInstructionContainerContainer *fyne.Container
				tempTestInstructionContainerContainer = container.NewHBox(
					tempExecutionStatusRectangle,
					tempIndentationLevelRectangle,
					serialOrParallelRectangleImage,
					tempTestInstructionContainerNameWidget)

				// Add the TestInstructionContainerContainer to the main Area
				testSuiteExecutionMainAreaForPreviewContainer.Add(tempTestInstructionContainerContainer)

				// Create testSuiteExecutionAttributesForPreview-object to be placed in the map
				var tempTestSuiteExecutionAttributesForPreview testCaseExecutionAttributesForPreviewStruct

				// Create testSuiteExecutionAttributesForPreview-object to be placed in the map
				tempTestSuiteExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
					LabelType:                          labelIsTestInstructionContainer,
					LabelText:                          previewObject.GetTestInstructionContainerName(),
					attributesContainerShouldBeVisible: false,
					testInstructionExecutionAttributesContainer: nil,
					childObjectsWithAttributes:                  nil,
				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var testInstructionContainerExecutionAttributesContainerMapKey testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				testInstructionContainerExecutionAttributesContainerMapKey = testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionContainerUuid())

				// Loop rest of PreView-Objects up until we get back to same 'previewObject.IndentationLevel'
				// and add all references to all TestInstructionAttributes-map
				// Do this if we are not at the end of the slice
				if previewObjectIndex+1 < len(tempTestSuiteExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects) {

					var tempChildObjectsWithAttributes []testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType

					for counter := previewObjectIndex + 1; counter < len(tempTestSuiteExecutionsListMessage.
						GetTestCasePreview().TestCaseStructureObjects); counter++ {

						// Check if IndentationLevel for next object is same ur higher than current AttributeObjects IndentationLevel
						if tempTestSuiteExecutionsListMessage.GetTestCasePreview().
							TestCaseStructureObjects[counter].IndentationLevel > previewObject.IndentationLevel {

							// Add Object to slice of object within this TestInstructionExecution-container-object
							switch tempTestSuiteExecutionsListMessage.GetTestCasePreview().
								TestCaseStructureObjects[counter].GetTestCaseStructureObjectType() {

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestSuiteExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionContainerUuid()))

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstruction:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestSuiteExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionUuid()))

							default:
								sharedCode.Logger.WithFields(logrus.Fields{
									"id": "81a17228-8355-47a7-9446-8a0cd6b1755e",
									"GetTestCaseStructureObjectType": tempTestSuiteExecutionsListMessage.GetTestCasePreview().
										TestCaseStructureObjects[counter].GetTestCaseStructureObjectType(),
								}).Error("Unknown 'GetTestCaseStructureObjectType'")

							}
						} else {
							// Indentationlevel back on start IndentationLevel, so break loop
							break
						}
					}

					// Add objects with higher, or equal, Indentation-level to map-object
					tempTestSuiteExecutionAttributesForPreview.childObjectsWithAttributes = tempChildObjectsWithAttributes

					// Add attributes-container struct to attributes-container-map
					testSuiteExecutionAttributesForPreviewMap[testInstructionContainerExecutionAttributesContainerMapKey] = &tempTestSuiteExecutionAttributesForPreview
				}

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
				testInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testSuiteNodeRectangleSize-14),
					float32(testSuiteNodeRectangleSize-14)))
				testInstructionColorRectangle.Resize(fyne.NewSize(float32(testSuiteNodeRectangleSize-14),
					float32(testSuiteNodeRectangleSize-14)))

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var tempTIEAttributesContainerMapKey testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				tempTIEAttributesContainerMapKey = testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionUuid())

				// Create the Name for the TestInstruction
				//var tempTestInstructionNameWidget *widget.Label
				//tempTestInstructionNameWidget = widget.NewLabel(previewObject.GetTestInstructionName())
				var tempTestInstructionNameWidget *clickableTInTICNameLabelInPreviewStruct
				tempTestInstructionNameWidget = newClickableTestInstructionNameLabelInPreview(
					previewObject.GetTestInstructionName(),
					testSuiteExecutionsModel.DetailedTestSuiteExecutionMapKeyType(testSuiteExecutionUuid+
						strconv.Itoa(int(testSuiteExecutionVersion))),
					tempTIEAttributesContainerMapKey,
					nil,
					nil,
					labelIsTestInstruction,
					testSuiteInstructionPreViewObjectRef)

				// Set correct color on ExecutionStatus Rectangle
				var statusId uint8
				var statusBackgroundColor color.RGBA
				var statusStrokeColor color.RGBA
				var useStroke bool

				// Extract TestInstructionExecution from TestInstruction
				var temp2TestSuiteExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage
				// Can preview be found in Map for "One TestSuiteExecution per TestCase" or "All TestSuiteExecutions per TestCase"
				switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

				case AllExecutionsForOneTestSuite:
					temp2TestSuiteExecutionsListMessage, _ = testSuiteExecutionsModelRef.GetSpecificTestSuiteExecutionForOneTestCaseUuid(
						testSuiteExecutionsModel.TestSuiteUuidType(selectedTestSuiteExecutionObjected.
							allExecutionsFoOneTestSuiteListObject.testSuiteUuidForTestSuiteExecutionThatIsShownInPreview),
						testSuiteExecutionsModel.TestSuiteExecutionUuidType(selectedTestSuiteExecutionObjected.
							allExecutionsFoOneTestSuiteListObject.testSuiteExecutionUuidThatIsShownInPreview))

				case OneExecutionPerTestSuite:
					temp2TestSuiteExecutionsListMessage, _ = testSuiteExecutionsModelRef.ReadFromTestSuiteExecutionsMap(
						testSuiteExecutionsModel.TestSuiteExecutionUuidType(testSuiteExecutionUuid))

				case NotDefined:

					temp2TestSuiteExecutionsListMessage = &fenixExecutionServerGuiGrpcApi.TestSuiteExecutionsListMessage{}

				}

				/*
					if existInMap == false {
						var id string
						id = "7945551e-5e4d-41d3-8faf-54f1501daac9"
						log.Fatalf(fmt.Sprintf("Couldn't find testSuiteExecutionUuidThatIsShownInPreview '%s' in "+
							"TestSuiteExecutionsThatCanBeViewedByUserMap. ID='%s'",
							selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.
								testSuiteExecutionUuidThatIsShownInPreview,
							id))
					}


				*/
				var tempTestInstructionExecutionsStatusPreviewValues []*fenixExecutionServerGuiGrpcApi.
					TestInstructionExecutionStatusPreviewValueMessage
				tempTestInstructionExecutionsStatusPreviewValues = temp2TestSuiteExecutionsListMessage.
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
							testSuiteExecutionUuidThatIsShownInPreview,
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
				testSuiteExecutionMainAreaForPreviewContainer.Add(tempTestInstructionContainer)

				// Create testSuiteExecutionAttributesForPreview-object to be placed in the map
				var tempTestSuiteExecutionAttributesForPreview testCaseExecutionAttributesForPreviewStruct

				// Create testSuiteExecutionAttributesForPreview-object to be placed in the map
				tempTestSuiteExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
					LabelType:                          notDefined,
					LabelText:                          "",
					attributesContainerShouldBeVisible: false,
					testInstructionExecutionAttributesContainer: nil,
					childObjectsWithAttributes:                  nil,
				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var testInstructionExecutionAttributesContainerMapKey testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				testInstructionExecutionAttributesContainerMapKey = testSuiteExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionUuid())

				// Add Attributes if there are any
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
						var attributeLabel *copyableLabelStruct
						attributeLabel = newCopyableLabel(attribute.AttributeName, true)

						// Create value for original attribute value
						var originalTestCaseAttributeValue *clickableAttributeInPreviewStruct
						originalTestCaseAttributeValue = newClickableAttributeInPreview(
							attribute.AttributeValue,
							attribute.AttributeName,
							previewObject.TestInstructionName,
							nil,
							nil,
							attributeIsOriginal,
							testSuiteInstructionPreViewObjectRef)

						// Create the RunTime-changed value for the attribute, if is changed
						var runtTimeChangedAttributeValue *clickableAttributeInPreviewStruct

						// Extract Attributes for TestInstructionExecution

						var runTimeUpdatedAttributesNameMapPtr *map[testSuiteExecutionsModel.AttributeNameMapKeyType]testSuiteExecutionsModel.
							RunTimeUpdatedAttributeValueType
						var runTimeUpdatedAttributesMap map[testSuiteExecutionsModel.AttributeNameMapKeyType]testSuiteExecutionsModel.
							RunTimeUpdatedAttributeValueType

						// Convert TestInstructionUuid into TestInstructionExecutionMapKey
						var testInstructionExecutionAttributeRunTimeUpdatedMapKey testSuiteExecutionsModel.TestInstructionExecutionUuidType
						testInstructionExecutionAttributeRunTimeUpdatedMapKey, existInMap = testSuiteExecutionsModel.TestSuiteExecutionsModel.
							GetTestInstructionExecutionUuidFromTestInstructionUuid(
								testSuiteExecutionsModel.TestSuiteExecutionUuidType(detailedTestSuiteExecutionMapKey),
								testSuiteExecutionsModel.RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(previewObject.GetTestInstructionUuid()))

						runTimeUpdatedAttributesNameMapPtr, existInMap = testInstructionExecutionsRunTimeUpdatedAttributesMap[testSuiteExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType(
							testInstructionExecutionAttributeRunTimeUpdatedMapKey)]

						if existInMap == false {

							/*
								sharedCode.Logger.WithFields(logrus.Fields{
									"id": "31ae4664-1107-4928-a95b-aed7b618e1f7",
									"testInstructionExecutionAttributeRunTimeUpdatedMapKey": testInstructionExecutionAttributeRunTimeUpdatedMapKey,
								}).Fatalln("Couldn't find testInstructionExecutionAttributeRunTimeUpdatedMapKey in 'testInstructionExecutionsRunTimeUpdatedAttributesMap', should never happen")

							*/
							//runtTimeChangedAttributeValue = widget.NewLabel("No RunTime-value")
							runTimeValueExists = false

						} else {

							runTimeUpdatedAttributesMap = *runTimeUpdatedAttributesNameMapPtr

							var attributeNameMapKey testSuiteExecutionsModel.AttributeNameMapKeyType
							attributeNameMapKey = testSuiteExecutionsModel.AttributeNameMapKeyType(attribute.AttributeName)

							var attributeRunTimeValue, existInMap = runTimeUpdatedAttributesMap[attributeNameMapKey]

							if existInMap == true {

								// Attribute has RunTime-value
								runtTimeChangedAttributeValue = newClickableAttributeInPreview(
									string(attributeRunTimeValue),
									attribute.AttributeName,
									previewObject.TestInstructionName,
									nil,
									nil,
									attributeIsRunTimeChanged,
									testSuiteInstructionPreViewObjectRef)

								runTimeValueExists = true

							} else {

								//runtTimeChangedAttributeValue = widget.NewLabel("No RunTime-value")
								runTimeValueExists = false

							}
						}

						var testCaseAttributeValuesContainer *fyne.Container
						testCaseAttributeValuesContainer = container.New(layout.NewHBoxLayout())

						// Add Attributes in order depending on if there is any RunTime-value for the Attribute
						if runTimeValueExists == true {

							// RunTime-value exits
							testCaseAttributeValuesContainer.Add(runtTimeChangedAttributeValue)
							testCaseAttributeValuesContainer.Add(layout.NewSpacer())
							testCaseAttributeValuesContainer.Add(originalTestCaseAttributeValue)

						} else {

							// No RunTime-value exits
							testCaseAttributeValuesContainer.Add(originalTestCaseAttributeValue)
							testCaseAttributeValuesContainer.Add(layout.NewSpacer())

						}

						// Add label and value to the attribute container
						testInstructionAttributesContainer.Add(attributeLabel)
						testInstructionAttributesContainer.Add(testCaseAttributeValuesContainer)

					}

					// Create the container containing the Attributes
					var tempTestInstructionAttributesContainer *fyne.Container
					tempTestInstructionAttributesContainer = container.NewHBox(
						tempIndentationLevelRectangleForAttributes, testInstructionAttributesContainer)

					// Make attributes container invisible
					tempTestInstructionAttributesContainer.Hide()

					// Create testSuiteExecutionAttributesForPreview-object to be placed in the map
					tempTestSuiteExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
						LabelType:                          labelIsTestInstruction,
						LabelText:                          previewObject.GetTestInstructionName(),
						attributesContainerShouldBeVisible: false,
						testInstructionExecutionAttributesContainer: tempTestInstructionAttributesContainer,
						childObjectsWithAttributes:                  nil,
					}

					// Add the TestInstructionContainerContainer to the main Area
					testSuiteExecutionMainAreaForPreviewContainer.Add(tempTestInstructionAttributesContainer)
				}

				// Loop rest of PreView-Objects up until we get back to same 'previewObject.IndentationLevel' and add all references to all TestInstructionAttributes-map
				if previewObjectIndex < len(tempTestSuiteExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects) {

					var tempChildObjectsWithAttributes []testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType

					for counter := previewObjectIndex + 1; counter < len(tempTestSuiteExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects); counter++ {

						// Check if IndentationLevel for next object is same ur higher than current AttributeObjects IndentationLevel
						if tempTestSuiteExecutionsListMessage.GetTestCasePreview().
							TestCaseStructureObjects[counter].IndentationLevel > previewObject.IndentationLevel {

							// Add Object to slice of object within this TestInstructionExecution-container-object
							switch tempTestSuiteExecutionsListMessage.GetTestCasePreview().
								TestCaseStructureObjects[counter].GetTestCaseStructureObjectType() {

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestSuiteExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionContainerUuid()))

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstruction:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testSuiteExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestSuiteExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionUuid()))

							default:
								sharedCode.Logger.WithFields(logrus.Fields{
									"id": "d8b796bd-6fdd-4f58-bb57-9848c565ac2f",
									"GetTestCaseStructureObjectType": tempTestSuiteExecutionsListMessage.GetTestCasePreview().
										TestCaseStructureObjects[counter].GetTestCaseStructureObjectType(),
								}).Error("Unknown 'GetTestCaseStructureObjectType'")

							}
						}
					}

					// Add objects with higher, or equal, Indentation-level to map-object
					tempTestSuiteExecutionAttributesForPreview.childObjectsWithAttributes = tempChildObjectsWithAttributes

					// Add attributes-container struct to attributes-container-map
					testSuiteExecutionAttributesForPreviewMap[testInstructionExecutionAttributesContainerMapKey] = &tempTestSuiteExecutionAttributesForPreview
				}

			default:
				log.Fatalf("Unknown 'previewObject.TestCaseStructureObjectType' which never should happen; %s", previewObject.TestCaseStructureObjectType.String())
			}
		}
	}

	testCaseMainAreaForPreviewBorderContainer := container.NewBorder(nil, nil, nil, nil, testSuiteExecutionMainAreaForPreviewContainer)
	testCaseMainAreaForPreviewScrollContainer := container.NewScroll(testCaseMainAreaForPreviewBorderContainer)

	// Create the container used for the TestCase, with TIC, TI and Attributes
	//testCasePreviewScrollContainer = container.NewScroll(tempContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestSuiteExecution Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}
	tempTopHeaderContainer := container.NewHBox(tempTopHeaderLabel)

	// Extract if row is selected or not
	var tempRowIsSelected bool
	switch selectedTestSuiteExecutionObjected.ExecutionsInGuiIsOfType {

	case AllExecutionsForOneTestSuite:
		tempRowIsSelected = selectedTestSuiteExecutionObjected.allExecutionsFoOneTestSuiteListObject.isAnyRowSelected

	case OneExecutionPerTestSuite:
		tempRowIsSelected = selectedTestSuiteExecutionObjected.oneExecutionPerTestSuiteListObject.isAnyRowSelected

	case NotDefined:

	}

	if tempRowIsSelected == true {

		// A row is selected and Preview should be shown

		/*container.NewHSplit(container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testSuiteExecutionPreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testSuiteExecutionPreviewBottomContainer), nil, nil,
		testCaseMainAreaForPreviewScrollContainer))


		*/

		var openedDetailedTestSuiteExecutionsMap map[openedDetailedTestSuiteExecutionsMapKeyType]*openedDetailedTestSuiteExecutionStruct
		// Check if Map with Open TestSuiteExecutions, map or Window has been initialized
		if openedDetailedTestSuiteExecutionsMapPtr == nil {
			openedDetailedTestSuiteExecutionsMap = make(map[openedDetailedTestSuiteExecutionsMapKeyType]*openedDetailedTestSuiteExecutionStruct)
			openedDetailedTestSuiteExecutionsMapPtr = &openedDetailedTestSuiteExecutionsMap
		} else {
			openedDetailedTestSuiteExecutionsMap = *openedDetailedTestSuiteExecutionsMapPtr
		}

		// Generate map-key
		var openedDetailedTestSuiteExecutionsMapKey openedDetailedTestSuiteExecutionsMapKeyType
		openedDetailedTestSuiteExecutionsMapKey = openedDetailedTestSuiteExecutionsMapKeyType(testSuiteExecutionUuid + strconv.Itoa(int(testSuiteExecutionVersion)))

		// Get Object if is open
		//var openedDetailedTestSuiteExecution openedDetailedTestSuiteExecutionStruct
		//var openedDetailedTestSuiteExecutionPtr *openedDetailedTestSuiteExecutionStruct

		// Extract the Object holding all information for Tab/Window-container for the TestSuiteExecution
		var openedDetailedTestSuiteExecutionPtr *openedDetailedTestSuiteExecutionStruct
		openedDetailedTestSuiteExecutionPtr, existInMap = openedDetailedTestSuiteExecutionsMap[openedDetailedTestSuiteExecutionsMapKey]
		if existInMap == false {

			// Object doesn't exist som create a new one
			openedDetailedTestSuiteExecutionPtr = &openedDetailedTestSuiteExecutionStruct{
				isTestSuiteExecutionOpenInTab: false,
				TestSuiteInstructionPreViewObjectInTab: &TestSuiteInstructionPreViewStruct{
					testSuitePreviewTestInstructionExecutionLogSplitContainer: nil,
					testSuiteExecutionPreviewContainerScroll:                  nil,
					testSuiteExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:           nil,
					testInstructionsExecutionLogContainer:                     nil,
					testInstructionsExecutionAttributesContainerScroll:        nil,
					testInstructionsExecutionAttributesContainer:              nil,
					testInstructionsExecutionDetailsContainer:                 nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				},
				isTestSuiteExecutionOpenInExternalWindow:          false,
				TestSuiteInstructionPreViewObjectInExternalWindow: nil,

				externalWindow: nil,
				tabItem:        nil,
			}
		}

		// Reference to current window

		// Define function to open a TestSuiteExecution in a external window
		var openTestSuiteExecutionInExternalWindowFunction func()
		openTestSuiteExecutionInExternalWindowFunction = func() {

			var fenixApp fyne.App
			fenixApp = *sharedCode.FenixAppPtr

			// Object exist, but check if the window is open
			if openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInExternalWindow == true {
				// Window is open, so get the existing window so make it the top window
				openedDetailedTestSuiteExecutionPtr.externalWindow.RequestFocus()

			} else {

				// Initialize Window-object
				openedDetailedTestSuiteExecutionPtr.
					TestSuiteInstructionPreViewObjectInExternalWindow = &TestSuiteInstructionPreViewStruct{
					testSuitePreviewTestInstructionExecutionLogSplitContainer: nil,
					testSuiteExecutionPreviewContainerScroll:                  nil,
					testSuiteExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:           nil,
					testInstructionsExecutionLogContainer:                     nil,
					testInstructionsExecutionAttributesContainerScroll:        nil,
					testInstructionsExecutionAttributesContainer:              nil,
					testInstructionsExecutionDetailsContainer:                 nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				}

				// Create a new window for the TestSuiteExecution
				openedDetailedTestSuiteExecutionPtr.externalWindow = fenixApp.NewWindow(fmt.Sprintf(
					"TestSuiteExecution '%s'", testSuiteExecutionUuid))

				// Set boolean to indicate that window is open
				openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInExternalWindow = true

				// Catch the close action
				openedDetailedTestSuiteExecutionPtr.externalWindow.SetOnClosed(func() {

					// Remove Content container object from map and mark not there
					defer func() {
						openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInExternalWindow = false
						openedDetailedTestSuiteExecutionPtr.
							TestSuiteInstructionPreViewObjectInExternalWindow = &TestSuiteInstructionPreViewStruct{
							testSuitePreviewTestInstructionExecutionLogSplitContainer: nil,
							testSuiteExecutionPreviewContainerScroll:                  nil,
							testSuiteExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
							testInstructionsExecutionDetailsContainerScroll:           nil,
							testInstructionsExecutionLogContainer:                     nil,
							testInstructionsExecutionAttributesContainerScroll:        nil,
							testInstructionsExecutionAttributesContainer:              nil,
							testInstructionsExecutionDetailsContainer:                 nil,
							preViewTabs:                       nil,
							attributeExplorerTab:              nil,
							logsExplorerTab:                   nil,
							testInstructionDetailsExplorerTab: nil,
						}
						openedDetailedTestSuiteExecutionPtr.externalWindow = nil
					}()

					//openedDetailedTestSuiteExecutionPtr.externalWindow.Close()

				})

				// Generate base objects for the PreView
				openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInExternalWindow.
					testSuitePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(
					openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInExternalWindow)

				// Generate a copy of the TestSuiteExecutionPreview
				openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInExternalWindow.
					GenerateTestSuiteExecutionPreviewContainer(
						testSuiteExecutionUuid,
						testSuiteExecutionVersion,
						testSuiteExecutionsModelRef,
						fromExternalWindow,
						&openedDetailedTestSuiteExecutionPtr.externalWindow)

				// Save the Object back to the Map
				openedDetailedTestSuiteExecutionsMap[openedDetailedTestSuiteExecutionsMapKey] = openedDetailedTestSuiteExecutionPtr

				// Set TestSuiteExecution as content
				var externalWindowCanvas fyne.Canvas
				externalWindowCanvas = openedDetailedTestSuiteExecutionPtr.externalWindow.Canvas()

				externalWindowCanvas.SetContent(container.NewBorder(
					nil,
					widget.NewButton("Close TestSuiteExecution", func() {
						openedDetailedTestSuiteExecutionPtr.externalWindow.Close()
					}),
					nil, nil,
					openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInExternalWindow.
						testSuitePreviewTestInstructionExecutionLogSplitContainer,
				))

				// Set size and show
				openedDetailedTestSuiteExecutionPtr.externalWindow.Resize(fyne.NewSize(1000, 800))
				openedDetailedTestSuiteExecutionPtr.externalWindow.Show()
			}

		}

		// Define function to open a TestSuiteExecution in a Tab
		var openTestSuiteExecutionInTabFunction func()
		openTestSuiteExecutionInTabFunction = func() {

			// Object exist, but check if the Tab already exist
			if openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInTab == true {
				// Tab exist, so get the tab so make it the visible one
				detailedTestSuiteExecutionsUITabObjectRef.Select(openedDetailedTestSuiteExecutionPtr.tabItem)

			} else {

				// Initialize Tab-object
				openedDetailedTestSuiteExecutionPtr.
					TestSuiteInstructionPreViewObjectInTab = &TestSuiteInstructionPreViewStruct{
					testSuitePreviewTestInstructionExecutionLogSplitContainer: nil,
					testSuiteExecutionPreviewContainerScroll:                  nil,
					testSuiteExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:           nil,
					testInstructionsExecutionLogContainer:                     nil,
					testInstructionsExecutionAttributesContainerScroll:        nil,
					testInstructionsExecutionAttributesContainer:              nil,
					testInstructionsExecutionDetailsContainer:                 nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				}

				// Generate base objects for the PreView
				openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInTab.
					testSuitePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(
					openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInTab)

				// Generate a copy of the TestSuiteExecutionPreview
				openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInTab.
					GenerateTestSuiteExecutionPreviewContainer(
						testSuiteExecutionUuid,
						testSuiteExecutionVersion,
						testSuiteExecutionsModelRef,
						fromTab,
						sharedCode.FenixMasterWindowPtr)

				// Save the Object back to the Map
				openedDetailedTestSuiteExecutionsMap[openedDetailedTestSuiteExecutionsMapKey] = openedDetailedTestSuiteExecutionPtr

				// Create a new Tab used for a specific TestSuiteExecution
				var newDetailedTestSuiteExecutionsTab *container.TabItem
				newDetailedTestSuiteExecutionsTab = container.NewTabItem(
					testSuiteExecutionUuid,
					container.NewBorder(
						nil,
						nil,
						nil,
						nil,
						openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInTab.
							testSuitePreviewTestInstructionExecutionLogSplitContainer))

				// Add Tab to TabObjects TestSuiteExecutions
				detailedTestSuiteExecutionsUITabObjectRef.Append(newDetailedTestSuiteExecutionsTab)

				// Extract Map with exit-functions for Tabs
				var exitFunctionsForDetailedTestSuiteExecutionsUITabObject map[*container.TabItem]func()
				exitFunctionsForDetailedTestSuiteExecutionsUITabObject = *exitFunctionsForDetailedTestSuiteExecutionsUITabObjectPtr

				// Add Close function to TabItem
				exitFunctionsForDetailedTestSuiteExecutionsUITabObject[newDetailedTestSuiteExecutionsTab] = func() {

					// Remove and indicate that no Tab is onpen
					openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInTab = false
					openedDetailedTestSuiteExecutionPtr.TestSuiteInstructionPreViewObjectInTab = nil

					// Delete the Exit-function for the Tab
					defer func() {
						delete(exitFunctionsForDetailedTestSuiteExecutionsUITabObject, newDetailedTestSuiteExecutionsTab)
					}()

				}

				// Select the newly created TabItem
				detailedTestSuiteExecutionsUITabObjectRef.Select(newDetailedTestSuiteExecutionsTab)

				// Refresh the Tab
				detailedTestSuiteExecutionsUITabObjectRef.Refresh()

				// Store the pointer to the TabItem
				openedDetailedTestSuiteExecutionPtr.tabItem = newDetailedTestSuiteExecutionsTab

				// Set boolean to indicate that window is open
				openedDetailedTestSuiteExecutionPtr.isTestSuiteExecutionOpenInTab = true

			}

		}

		// Define the menu items to open TestSuiteExecution in Tab/External Window
		var items []*fyne.MenuItem

		// From where is the opening of the TestSuiteExecution initiated; FromExecutionList, FromExternalWindow, FromTab
		switch openedTestSuiteExecutionFrom {

		case fromExecutionList:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestSuiteExecution in Tab", openTestSuiteExecutionInTabFunction),
				fyne.NewMenuItem("Open TestSuiteExecution in separate window", openTestSuiteExecutionInExternalWindowFunction),
			}

		case fromExternalWindow:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestSuiteExecution in Tab", openTestSuiteExecutionInTabFunction),
			}

		case fromTab:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestSuiteExecution in separate window", openTestSuiteExecutionInExternalWindowFunction),
			}

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":                           "c4cb778e-6b4f-4a94-8cc6-88fde16de276",
				"openedTestSuiteExecutionFrom": openedTestSuiteExecutionFrom,
			}).Fatalln("Unhandled 'openedTestSuiteExecutionFrom', should never happen")

		}

		// Create the button
		var btn *widget.Button
		btn = widget.NewButton("Actions", func() {
			// 1) Get the absolute position of the button on the canvas
			pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(btn)

			// 2) Build the popup menu
			popup := widget.NewPopUpMenu(fyne.NewMenu("", items...), currentWindow.Canvas())

			// 3) Show it just below the button
			popup.ShowAtPosition(pos.Add(fyne.NewPos(0, btn.Size().Height)))
		})

		tempTopHeaderContainer.Add(btn)

		// Create the new Execution-Preview
		testSuiteInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Objects[0] = container.NewBorder(
			container.NewVBox(container.NewCenter(tempTopHeaderContainer), testSuiteExecutionPreviewTopContainer, widget.NewSeparator()),
			container.NewVBox(widget.NewSeparator(), testSuiteExecutionPreviewBottomContainer), nil, nil,
			testCaseMainAreaForPreviewScrollContainer)

		// Create a new temporary container for the logs
		testSuiteInstructionPreViewObjectRef.testInstructionsExecutionLogContainer.Objects[0] = container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))

	} else {

		// No row is selected and only an info text should be shown
		testSuiteInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))

		// Create a new temporary container for the logs
		testSuiteInstructionPreViewObjectRef.testInstructionsExecutionLogContainer.Objects[0] = container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))
	}

	// Add attributes container to attributes-container-map-ptr
	testSuiteExecutionAttributesForPreviewMapPtr = &testSuiteExecutionAttributesForPreviewMap

	// Refresh the 'testSuiteExecutionPreviewContainer'
	testSuiteInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Refresh()

}

// ClearTestSuiteExecutionPreviewContainer
// Clears the preview area for the TestSuiteExecution
func (testCaseInstructionPreViewObjectRef *TestSuiteInstructionPreViewStruct) ClearTestSuiteExecutionPreviewContainer() {
	testCaseInstructionPreViewObjectRef.testSuiteExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestSuiteExecution to get the Preview"))
}
