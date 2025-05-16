package listTestCaseExecutionsUI

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/detailedExecutionsModel"
	"FenixTesterGui/testCaseExecutions/testCaseExecutionsModel"
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

func (testCaseInstructionPreViewObjectRef *TestCaseInstructionPreViewStruct) GenerateTestCaseExecutionPreviewContainer(
	testCaseExecutionUuid string,
	testCaseExecutionVersion uint32,
	testCaseExecutionsModelRef *testCaseExecutionsModel.TestCaseExecutionsModelStruct,
	openedTestCaseExecutionFrom openedTestCaseExecutionFromType,
	currentWindowPtr *fyne.Window) {

	var currentWindow fyne.Window
	currentWindow = *currentWindowPtr

	var testCaseExecutionPreviewTopContainer *fyne.Container
	var testCaseExecutionPreviewBottomContainer *fyne.Container
	//var testCasePreviewScrollContainer *container.Scroll
	var testCaseExecutionMainAreaForPreviewContainer *fyne.Container

	var err error
	var existInMap bool
	var foundValue bool
	var runTimeValueExists bool

	// Lock Map
	testCaseExecutionAttributesForPreviewMapMutex.Lock()

	// Unlock map
	defer testCaseExecutionAttributesForPreviewMapMutex.Unlock()

	// Clear out 'testCaseExecutionAttributesForPreviewMapPtr'
	var testCaseExecutionAttributesForPreviewMap map[testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct
	testCaseExecutionAttributesForPreviewMap = make(map[testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType]*testCaseExecutionAttributesForPreviewStruct)

	testCaseExecutionAttributesForPreviewMapPtr = &testCaseExecutionAttributesForPreviewMap

	// Verify that number Headers match number of columns, constant 'numberColumnsInTestCaseExecutionsListUI'
	if len(testCaseExecutionsListTableHeader) != numberColumnsInTestCaseExecutionsListUI {
		log.Fatalln(fmt.Sprintf("Number of elements in 'tempRowslice' missmatch contant 'numberColumnsInTestCaseExecutionsListUI'. %d vs %d. ID: %s",
			testCaseExecutionsListTableHeader,
			numberColumnsInTestCaseExecutionsListUI,
			"c2b8a13c-ec20-46c2-adf9-965247732e07"))
	}

	// Get Data for the Preview
	var tempTestCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage

	// Can preview be found in Map for "One TestCaseExecution per TestCase" or "All TestCaseExecutions per TestCase"
	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

	case AllExecutionsForOneTestCase:
		tempTestCaseExecutionsListMessage, _ = testCaseExecutionsModelRef.GetSpecificTestCaseExecutionForOneTestCaseUuid(
			testCaseExecutionsModel.TestCaseUuidType(selectedTestCaseExecutionObjected.
				allExecutionsFoOneTestCaseListObject.testCaseUuidForTestCaseExecutionThatIsShownInPreview),
			testCaseExecutionsModel.TestCaseExecutionUuidType(selectedTestCaseExecutionObjected.
				allExecutionsFoOneTestCaseListObject.testCaseExecutionUuidThatIsShownInPreview))

	case OneExecutionPerTestCase:
		tempTestCaseExecutionsListMessage, _ = testCaseExecutionsModelRef.ReadFromTestCaseExecutionsMap(
			testCaseExecutionsModel.TestCaseExecutionUuidType(testCaseExecutionUuid))

	case NotDefined:

		tempTestCaseExecutionsListMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage{}

	}

	// Read from the TestCaseExecutions-Map to Get object holding Logs, RunTimeAtributes and ...

	//var detailedTestCaseExecutionsMapObject testCaseExecutionsModel.DetailedTestCaseExecutionsMapObjectStruct

	var detailedTestCaseExecutionsObjectsMapPtr *map[testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType]*testCaseExecutionsModel.DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionsObjectsMap map[testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType]*testCaseExecutionsModel.DetailedTestCaseExecutionsMapObjectStruct
	detailedTestCaseExecutionsObjectsMapPtr = testCaseExecutionsModelRef.DetailedTestCaseExecutionsObjectsMapPtr
	detailedTestCaseExecutionsObjectsMap = *detailedTestCaseExecutionsObjectsMapPtr

	var detailedTestCaseExecutionsObjectPtr *testCaseExecutionsModel.DetailedTestCaseExecutionsMapObjectStruct
	var detailedTestCaseExecutionMapKey testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType
	detailedTestCaseExecutionMapKey = testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType(
		testCaseExecutionUuid + strconv.Itoa(int(testCaseExecutionVersion)))

	detailedTestCaseExecutionsObjectPtr, existInMap = detailedTestCaseExecutionsObjectsMap[detailedTestCaseExecutionMapKey]

	if existInMap == false {

		sharedCode.Logger.WithFields(logrus.Fields{
			"id":                              "98ac73f4-2b99-4fc8-826d-0f3282ca9870",
			"detailedTestCaseExecutionMapKey": detailedTestCaseExecutionMapKey,
		}).Fatalln("Couldn't find TestCaseExecution in 'detailedTestCaseExecutionMap', should never happen!")
	}

	// Extract the run-time variables map
	var runTimeUpdatedAttributesMapPtr *map[testCaseExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[testCaseExecutionsModel.AttributeNameMapKeyType]testCaseExecutionsModel.
		RunTimeUpdatedAttributeValueType
	var testInstructionExecutionsRunTimeUpdatedAttributesMap map[testCaseExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType]*map[testCaseExecutionsModel.AttributeNameMapKeyType]testCaseExecutionsModel.
		RunTimeUpdatedAttributeValueType

	runTimeUpdatedAttributesMapPtr = detailedTestCaseExecutionsObjectPtr.RunTimeUpdatedAttributesMapPtr
	testInstructionExecutionsRunTimeUpdatedAttributesMap = *runTimeUpdatedAttributesMapPtr

	// Create the Top container
	testCaseExecutionPreviewTopContainer = container.New(layout.NewFormLayout())

	// Create the ExecutionStatus Rectangle for TestCaseExecution-status
	var tempTestCaseExecutionStatusRectangle *canvas.Rectangle
	tempTestCaseExecutionStatusRectangle = canvas.NewRectangle(color.Transparent)

	// Resize the ExecutionStatus rectangle
	tempTestCaseExecutionStatusRectangle.SetMinSize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHeight,
	})
	tempTestCaseExecutionStatusRectangle.Resize(fyne.Size{
		Width:  testCaseExecutionStatusRectangleWidth,
		Height: testCaseExecutionStatusRectangleHeight,
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
	copyableTestCaseNameLabel := newCopyableLabel(tempTestCaseExecutionsListMessage.GetTestCaseName(), true)
	testCaseExecutionPreviewTopContainer.Add(copyableTestCaseNameLabel)

	// Add TestCaseExecutionStatus
	tempTestCaseExecutionStatusLabel := widget.NewLabel("TestCaseExecution-status:")
	tempTestCaseExecutionStatusLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempTestCaseExecutionStatusLabel)

	var testCaseExecutionStatusStackContainer *fyne.Container
	var testCaseExecutionStatusHBoxContainer *fyne.Container
	var testCaseExecutionStatusOuterHBoxContainer *fyne.Container
	testCaseExecutionStatusStackContainer = container.NewStack()
	testCaseExecutionStatusHBoxContainer = container.NewHBox(
		widget.NewLabel(" "),
		canvas.NewText(executionStatusColorMapObjectForTestCaseExecution.ExecutionStatusName, color.Black),
		widget.NewLabel(" "))
	testCaseExecutionStatusStackContainer.Add(tempTestCaseExecutionStatusRectangle)
	testCaseExecutionStatusStackContainer.Add(testCaseExecutionStatusHBoxContainer)
	testCaseExecutionStatusOuterHBoxContainer = container.NewHBox(
		testCaseExecutionStatusStackContainer,
		layout.NewSpacer())

	testCaseExecutionPreviewTopContainer.Add(testCaseExecutionStatusOuterHBoxContainer)

	// Add TestCaseOwner Domain Top container
	tempOwnerDomainLabel := widget.NewLabel("OwnerDomain:")
	tempOwnerDomainLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewTopContainer.Add(tempOwnerDomainLabel)
	copyableDomainThatOwnTheTestCaseLabel := newCopyableLabel(
		tempTestCaseExecutionsListMessage.GetTestCasePreview().GetDomainThatOwnTheTestCase(), true)
	testCaseExecutionPreviewTopContainer.Add(copyableDomainThatOwnTheTestCaseLabel)

	// Add empty row
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
	testCaseExecutionPreviewBottomContainer.Add(widget.NewLabel(strconv.Itoa(int(tempTestCaseExecutionsListMessage.
		GetTestCaseExecutionVersion()))))

	// Add ExecutionStartTimeStamp to Bottom container
	tempExecutionStartTimeStampLabel := widget.NewLabel("TestCase Execution Start TimeStamp:")
	tempExecutionStartTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStartTimeStampLabel)
	copyableExecutionStartTimeStampLabel := newCopyableLabel(
		tempTestCaseExecutionsListMessage.
			GetExecutionStartTimeStamp().AsTime().String(), true)
	testCaseExecutionPreviewBottomContainer.Add(copyableExecutionStartTimeStampLabel)

	// Add ExecutionStopTimeStamp to Bottom container
	tempExecutionStopTimeStampLabel := widget.NewLabel("TestCase Execution Stop TimeStamp:")
	tempExecutionStopTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStopTimeStampLabel)
	copyableExecutionStopTimeStampLabel := newCopyableLabel(
		tempTestCaseExecutionsListMessage.
			GetExecutionStopTimeStamp().AsTime().String(), true)
	testCaseExecutionPreviewBottomContainer.Add(copyableExecutionStopTimeStampLabel)

	// Add ExecutionStatusUpdateTimeStamp to Bottom container
	tempExecutionStatusUpdateTimeStampLabel := widget.NewLabel("TestCase Execution Status Update TimeStamp:")
	tempExecutionStatusUpdateTimeStampLabel.TextStyle = fyne.TextStyle{Bold: true}
	testCaseExecutionPreviewBottomContainer.Add(tempExecutionStatusUpdateTimeStampLabel)
	copyableExecutionStatusUpdateTimeStampLabel := newCopyableLabel(
		tempTestCaseExecutionsListMessage.
			GetExecutionStatusUpdateTimeStamp().AsTime().String(), true)
	testCaseExecutionPreviewBottomContainer.Add(copyableExecutionStatusUpdateTimeStampLabel)

	// Add LastSavedByUserGCPAuthorization to Bottom container
	/*
		tempLastSavedByUserGCPAuthorizationLabel := widget.NewLabel("Last saved by this GCP-user:")
		tempLastSavedByUserGCPAuthorizationLabel.TextStyle = fyne.TextStyle{Bold: true}
		testCaseExecutionPreviewBottomContainer.Add(tempLastSavedByUserGCPAuthorizationLabel)
		copyableExecutionStatusUpdateTimeStampLabel := newCopyableLabel(
			tempTestCaseExecutionsListMessage..
				Get().AsTime().String(),true)
		testCaseExecutionPreviewBottomContainer.Add(copyableExecutionStatusUpdateTimeStampLabel)

	*/
	// Create the area used for TIC, TI and the attributes
	testCaseExecutionMainAreaForPreviewContainer = container.NewVBox()

	// Check if there is TestCase Preview information
	if tempTestCaseExecutionsListMessage.GetTestCasePreview() == nil {

		testCaseExecutionMainAreaForPreviewContainer.Add(widget.NewLabel("No Preview information found in database!"))

	} else {

		// Loop the preview objects and to container
		for previewObjectIndex, previewObject := range tempTestCaseExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects {

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
					serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4),
						float32(testCaseNodeRectangleSize-4)))
					serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4),
						float32(testCaseNodeRectangleSize-4)))

				} else {
					// Convert the byte slice into an image.Image object
					if imageData_tic_parallellImage == nil {
						imageData_tic_parallellImage, err = png.Decode(bytes.NewReader(tic_parallellImage))
						if err != nil {
							log.Fatalf("Failed to decode image: %v", err)
						}
					}
					serialOrParallelRectangleImage = canvas.NewImageFromImage(imageData_tic_parallellImage)
					serialOrParallelRectangleImage.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-4),
						float32(testCaseNodeRectangleSize-4)))
					serialOrParallelRectangleImage.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-4),
						float32(testCaseNodeRectangleSize-4)))

				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var tempTIEAttributesContainerMapKey testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				tempTIEAttributesContainerMapKey = testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionContainerUuid())

				// Create the Name for the TestInstructionContainer
				// var tempTestInstructionContainerNameWidget *widget.Label
				//tempTestInstructionContainerNameWidget = widget.NewLabel(previewObject.GetTestInstructionContainerName())
				var tempTestInstructionContainerNameWidget *clickableTInTICNameLabelInPreviewStruct
				tempTestInstructionContainerNameWidget = newClickableTestInstructionNameLabelInPreview(
					previewObject.GetTestInstructionContainerName(),
					testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType(testCaseExecutionUuid+
						strconv.Itoa(int(testCaseExecutionVersion))),
					tempTIEAttributesContainerMapKey,
					nil,
					nil,
					labelIsTestInstructionContainer,
					testCaseInstructionPreViewObjectRef)

				// Create the container containing the TestInstructionContainer
				var tempTestInstructionContainerContainer *fyne.Container
				tempTestInstructionContainerContainer = container.NewHBox(
					tempExecutionStatusRectangle,
					tempIndentationLevelRectangle,
					serialOrParallelRectangleImage,
					tempTestInstructionContainerNameWidget)

				// Add the TestInstructionContainerContainer to the main Area
				testCaseExecutionMainAreaForPreviewContainer.Add(tempTestInstructionContainerContainer)

				// Create testCaseExecutionAttributesForPreview-object to be placed in the map
				var tempTestCaseExecutionAttributesForPreview testCaseExecutionAttributesForPreviewStruct

				// Create testCaseExecutionAttributesForPreview-object to be placed in the map
				tempTestCaseExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
					LabelType:                          labelIsTestInstructionContainer,
					LabelText:                          previewObject.GetTestInstructionContainerName(),
					attributesContainerShouldBeVisible: false,
					testInstructionExecutionAttributesContainer: nil,
					childObjectsWithAttributes:                  nil,
				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var testInstructionContainerExecutionAttributesContainerMapKey testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				testInstructionContainerExecutionAttributesContainerMapKey = testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionContainerUuid())

				// Loop rest of PreView-Objects up until we get back to same 'previewObject.IndentationLevel'
				// and add all references to all TestInstructionAttributes-map
				// Do this if we are not at the end of the slice
				if previewObjectIndex+1 < len(tempTestCaseExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects) {

					var tempChildObjectsWithAttributes []testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType

					for counter := previewObjectIndex + 1; counter < len(tempTestCaseExecutionsListMessage.
						GetTestCasePreview().TestCaseStructureObjects); counter++ {

						// Check if IndentationLevel for next object is same ur higher than current AttributeObjects IndentationLevel
						if tempTestCaseExecutionsListMessage.GetTestCasePreview().
							TestCaseStructureObjects[counter].IndentationLevel > previewObject.IndentationLevel {

							// Add Object to slice of object within this TestInstructionExecution-container-object
							switch tempTestCaseExecutionsListMessage.GetTestCasePreview().
								TestCaseStructureObjects[counter].GetTestCaseStructureObjectType() {

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestCaseExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionContainerUuid()))

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstruction:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestCaseExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionUuid()))

							default:
								sharedCode.Logger.WithFields(logrus.Fields{
									"id": "5607884d-8884-48b0-8636-2ad3ca7046f9",
									"GetTestCaseStructureObjectType": tempTestCaseExecutionsListMessage.GetTestCasePreview().
										TestCaseStructureObjects[counter].GetTestCaseStructureObjectType(),
								}).Error("Unknown 'GetTestCaseStructureObjectType'")

							}
						} else {
							// Indentationlevel back on start IndentationLevel, so break loop
							break
						}
					}

					// Add objects with higher, or equal, Indentation-level to map-object
					tempTestCaseExecutionAttributesForPreview.childObjectsWithAttributes = tempChildObjectsWithAttributes

					// Add attributes-container struct to attributes-container-map
					testCaseExecutionAttributesForPreviewMap[testInstructionContainerExecutionAttributesContainerMapKey] = &tempTestCaseExecutionAttributesForPreview
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
				testInstructionColorRectangle.SetMinSize(fyne.NewSize(float32(testCaseNodeRectangleSize-14),
					float32(testCaseNodeRectangleSize-14)))
				testInstructionColorRectangle.Resize(fyne.NewSize(float32(testCaseNodeRectangleSize-14),
					float32(testCaseNodeRectangleSize-14)))

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var tempTIEAttributesContainerMapKey testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				tempTIEAttributesContainerMapKey = testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType(previewObject.GetTestInstructionUuid())

				// Create the Name for the TestInstruction
				//var tempTestInstructionNameWidget *widget.Label
				//tempTestInstructionNameWidget = widget.NewLabel(previewObject.GetTestInstructionName())
				var tempTestInstructionNameWidget *clickableTInTICNameLabelInPreviewStruct
				tempTestInstructionNameWidget = newClickableTestInstructionNameLabelInPreview(
					previewObject.GetTestInstructionName(),
					testCaseExecutionsModel.DetailedTestCaseExecutionMapKeyType(testCaseExecutionUuid+
						strconv.Itoa(int(testCaseExecutionVersion))),
					tempTIEAttributesContainerMapKey,
					nil,
					nil,
					labelIsTestInstruction,
					testCaseInstructionPreViewObjectRef)

				// Set correct color on ExecutionStatus Rectangle
				var statusId uint8
				var statusBackgroundColor color.RGBA
				var statusStrokeColor color.RGBA
				var useStroke bool

				// Extract TestInstructionExecution from TestInstruction
				var temp2TestCaseExecutionsListMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage
				// Can preview be found in Map for "One TestCaseExecution per TestCase" or "All TestCaseExecutions per TestCase"
				switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

				case AllExecutionsForOneTestCase:
					temp2TestCaseExecutionsListMessage, _ = testCaseExecutionsModelRef.GetSpecificTestCaseExecutionForOneTestCaseUuid(
						testCaseExecutionsModel.TestCaseUuidType(selectedTestCaseExecutionObjected.
							allExecutionsFoOneTestCaseListObject.testCaseUuidForTestCaseExecutionThatIsShownInPreview),
						testCaseExecutionsModel.TestCaseExecutionUuidType(selectedTestCaseExecutionObjected.
							allExecutionsFoOneTestCaseListObject.testCaseExecutionUuidThatIsShownInPreview))

				case OneExecutionPerTestCase:
					temp2TestCaseExecutionsListMessage, _ = testCaseExecutionsModelRef.ReadFromTestCaseExecutionsMap(
						testCaseExecutionsModel.TestCaseExecutionUuidType(testCaseExecutionUuid))

				case NotDefined:

					temp2TestCaseExecutionsListMessage = &fenixExecutionServerGuiGrpcApi.TestCaseExecutionsListMessage{}

				}

				/*
					if existInMap == false {
						var id string
						id = "7945551e-5e4d-41d3-8faf-54f1501daac9"
						log.Fatalf(fmt.Sprintf("Couldn't find testCaseExecutionUuidThatIsShownInPreview '%s' in "+
							"TestCaseExecutionsThatCanBeViewedByUserMap. ID='%s'",
							selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.
								testCaseExecutionUuidThatIsShownInPreview,
							id))
					}


				*/
				var tempTestInstructionExecutionsStatusPreviewValues []*fenixExecutionServerGuiGrpcApi.
					TestInstructionExecutionStatusPreviewValueMessage
				tempTestInstructionExecutionsStatusPreviewValues = temp2TestCaseExecutionsListMessage.
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
							testCaseExecutionUuidThatIsShownInPreview,
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

				// Create testCaseExecutionAttributesForPreview-object to be placed in the map
				var tempTestCaseExecutionAttributesForPreview testCaseExecutionAttributesForPreviewStruct

				// Create testCaseExecutionAttributesForPreview-object to be placed in the map
				tempTestCaseExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
					LabelType:                          notDefined,
					LabelText:                          "",
					attributesContainerShouldBeVisible: false,
					testInstructionExecutionAttributesContainer: nil,
					childObjectsWithAttributes:                  nil,
				}

				// Create Map-key; TCEoTICoTIEAttributesContainerMapKeyType
				var testInstructionExecutionAttributesContainerMapKey testCaseExecutionsModel.
					TCEoTICoTIEAttributesContainerMapKeyType
				testInstructionExecutionAttributesContainerMapKey = testCaseExecutionsModel.
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
							testCaseInstructionPreViewObjectRef)

						// Create the RunTime-changed value for the attribute, if is changed
						var runtTimeChangedAttributeValue *clickableAttributeInPreviewStruct

						// Extract Attributes for TestInstructionExecution

						var runTimeUpdatedAttributesNameMapPtr *map[testCaseExecutionsModel.AttributeNameMapKeyType]testCaseExecutionsModel.
							RunTimeUpdatedAttributeValueType
						var runTimeUpdatedAttributesMap map[testCaseExecutionsModel.AttributeNameMapKeyType]testCaseExecutionsModel.
							RunTimeUpdatedAttributeValueType

						// Convert TestInstructionUuid into TestInstructionExecutionMapKey
						var testInstructionExecutionAttributeRunTimeUpdatedMapKey testCaseExecutionsModel.TestInstructionExecutionUuidType
						testInstructionExecutionAttributeRunTimeUpdatedMapKey, existInMap = testCaseExecutionsModel.TestCaseExecutionsModel.
							GetTestInstructionExecutionUuidFromTestInstructionUuid(
								testCaseExecutionsModel.TestCaseExecutionUuidType(detailedTestCaseExecutionMapKey),
								testCaseExecutionsModel.RelationBetweenTestInstructionUuidAndTestInstructionExectuionMapKeyType(previewObject.GetTestInstructionUuid()))

						runTimeUpdatedAttributesNameMapPtr, existInMap = testInstructionExecutionsRunTimeUpdatedAttributesMap[testCaseExecutionsModel.TestInstructionExecutionAttributeRunTimeUpdatedMapKeyType(
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

							var attributeNameMapKey testCaseExecutionsModel.AttributeNameMapKeyType
							attributeNameMapKey = testCaseExecutionsModel.AttributeNameMapKeyType(attribute.AttributeName)

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
									testCaseInstructionPreViewObjectRef)

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

					// Create testCaseExecutionAttributesForPreview-object to be placed in the map
					tempTestCaseExecutionAttributesForPreview = testCaseExecutionAttributesForPreviewStruct{
						LabelType:                          labelIsTestInstruction,
						LabelText:                          previewObject.GetTestInstructionName(),
						attributesContainerShouldBeVisible: false,
						testInstructionExecutionAttributesContainer: tempTestInstructionAttributesContainer,
						childObjectsWithAttributes:                  nil,
					}

					// Add the TestInstructionContainerContainer to the main Area
					testCaseExecutionMainAreaForPreviewContainer.Add(tempTestInstructionAttributesContainer)
				}

				// Loop rest of PreView-Objects up until we get back to same 'previewObject.IndentationLevel' and add all references to all TestInstructionAttributes-map
				if previewObjectIndex < len(tempTestCaseExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects) {

					var tempChildObjectsWithAttributes []testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType

					for counter := previewObjectIndex + 1; counter < len(tempTestCaseExecutionsListMessage.GetTestCasePreview().TestCaseStructureObjects); counter++ {

						// Check if IndentationLevel for next object is same ur higher than current AttributeObjects IndentationLevel
						if tempTestCaseExecutionsListMessage.GetTestCasePreview().
							TestCaseStructureObjects[counter].IndentationLevel > previewObject.IndentationLevel {

							// Add Object to slice of object within this TestInstructionExecution-container-object
							switch tempTestCaseExecutionsListMessage.GetTestCasePreview().
								TestCaseStructureObjects[counter].GetTestCaseStructureObjectType() {

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstructionContainer:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestCaseExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionContainerUuid()))

							case fenixExecutionServerGuiGrpcApi.TestCasePreviewStructureMessage_TestInstruction:
								tempChildObjectsWithAttributes = append(tempChildObjectsWithAttributes,
									testCaseExecutionsModel.TCEoTICoTIEAttributesContainerMapKeyType(
										tempTestCaseExecutionsListMessage.GetTestCasePreview().
											TestCaseStructureObjects[counter].GetTestInstructionUuid()))

							default:
								sharedCode.Logger.WithFields(logrus.Fields{
									"id": "5607884d-8884-48b0-8636-2ad3ca7046f9",
									"GetTestCaseStructureObjectType": tempTestCaseExecutionsListMessage.GetTestCasePreview().
										TestCaseStructureObjects[counter].GetTestCaseStructureObjectType(),
								}).Error("Unknown 'GetTestCaseStructureObjectType'")

							}
						}
					}

					// Add objects with higher, or equal, Indentation-level to map-object
					tempTestCaseExecutionAttributesForPreview.childObjectsWithAttributes = tempChildObjectsWithAttributes

					// Add attributes-container struct to attributes-container-map
					testCaseExecutionAttributesForPreviewMap[testInstructionExecutionAttributesContainerMapKey] = &tempTestCaseExecutionAttributesForPreview
				}

			default:
				log.Fatalf("Unknown 'previewObject.TestCaseStructureObjectType' which never should happen; %s", previewObject.TestCaseStructureObjectType.String())
			}
		}
	}

	testCaseMainAreaForPreviewBorderContainer := container.NewBorder(nil, nil, nil, nil, testCaseExecutionMainAreaForPreviewContainer)
	testCaseMainAreaForPreviewScrollContainer := container.NewScroll(testCaseMainAreaForPreviewBorderContainer)

	// Create the container used for the TestCase, with TIC, TI and Attributes
	//testCasePreviewScrollContainer = container.NewScroll(tempContainer)

	// Create Top header for Preview
	tempTopHeaderLabel := widget.NewLabel("TestCaseExecution Preview")
	tempTopHeaderLabel.TextStyle = fyne.TextStyle{Bold: true}
	tempTopHeaderContainer := container.NewHBox(tempTopHeaderLabel)

	// Extract if row is selected or not
	var tempRowIsSelected bool
	switch selectedTestCaseExecutionObjected.ExecutionsInGuiIsOfType {

	case AllExecutionsForOneTestCase:
		tempRowIsSelected = selectedTestCaseExecutionObjected.allExecutionsFoOneTestCaseListObject.isAnyRowSelected

	case OneExecutionPerTestCase:
		tempRowIsSelected = selectedTestCaseExecutionObjected.oneExecutionPerTestCaseListObject.isAnyRowSelected

	case NotDefined:

	}

	if tempRowIsSelected == true {

		// A row is selected and Preview should be shown

		/*container.NewHSplit(container.NewBorder(
		container.NewVBox(container.NewCenter(tempTopHeaderLabel), testCaseExecutionPreviewTopContainer, widget.NewSeparator()),
		container.NewVBox(widget.NewSeparator(), testCaseExecutionPreviewBottomContainer), nil, nil,
		testCaseMainAreaForPreviewScrollContainer))


		*/

		var openedDetailedTestCaseExecutionsMap map[openedDetailedTestCaseExecutionsMapKeyType]*openedDetailedTestCaseExecutionStruct
		// Check if Map with Open TestCaseExecutions, map or Window has been initialized
		if openedDetailedTestCaseExecutionsMapPtr == nil {
			openedDetailedTestCaseExecutionsMap = make(map[openedDetailedTestCaseExecutionsMapKeyType]*openedDetailedTestCaseExecutionStruct)
			openedDetailedTestCaseExecutionsMapPtr = &openedDetailedTestCaseExecutionsMap
		} else {
			openedDetailedTestCaseExecutionsMap = *openedDetailedTestCaseExecutionsMapPtr
		}

		// Generate map-key
		var openedDetailedTestCaseExecutionsMapKey openedDetailedTestCaseExecutionsMapKeyType
		openedDetailedTestCaseExecutionsMapKey = openedDetailedTestCaseExecutionsMapKeyType(testCaseExecutionUuid + strconv.Itoa(int(testCaseExecutionVersion)))

		// Get Object if is open
		//var openedDetailedTestCaseExecution openedDetailedTestCaseExecutionStruct
		//var openedDetailedTestCaseExecutionPtr *openedDetailedTestCaseExecutionStruct

		// Extract the Object holding all information for Tab/Window-container for the TestCaseExecution
		var openedDetailedTestCaseExecutionPtr *openedDetailedTestCaseExecutionStruct
		openedDetailedTestCaseExecutionPtr, existInMap = openedDetailedTestCaseExecutionsMap[openedDetailedTestCaseExecutionsMapKey]
		if existInMap == false {

			// Object doesn't exist som create a new one
			openedDetailedTestCaseExecutionPtr = &openedDetailedTestCaseExecutionStruct{
				isTestCaseExecutionOpenInTab: false,
				TestCaseInstructionPreViewObjectInTab: &TestCaseInstructionPreViewStruct{
					testCasePreviewTestInstructionExecutionLogSplitContainer: nil,
					testCaseExecutionPreviewContainerScroll:                  nil,
					testCaseExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:          nil,
					testInstructionsExecutionLogContainer:                    nil,
					testInstructionsExecutionAttributesContainerScroll:       nil,
					testInstructionsExecutionAttributesContainer:             nil,
					testInstructionsExecutionDetailsContainer:                nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				},
				isTestCaseExecutionOpenInExternalWindow:          false,
				TestCaseInstructionPreViewObjectInExternalWindow: nil,

				externalWindow: nil,
				tabItem:        nil,
			}
		}

		// Reference to current window

		// Define function to open a TestCaseExecution in a external window
		var openTestCaseExecutionInExternalWindowFunction func()
		openTestCaseExecutionInExternalWindowFunction = func() {

			var fenixApp fyne.App
			fenixApp = *sharedCode.FenixAppPtr

			// Object exist, but check if the window is open
			if openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInExternalWindow == true {
				// Window is open, so get the existing window so make it the top window
				openedDetailedTestCaseExecutionPtr.externalWindow.RequestFocus()

			} else {

				// Initialize Window-object
				openedDetailedTestCaseExecutionPtr.
					TestCaseInstructionPreViewObjectInExternalWindow = &TestCaseInstructionPreViewStruct{
					testCasePreviewTestInstructionExecutionLogSplitContainer: nil,
					testCaseExecutionPreviewContainerScroll:                  nil,
					testCaseExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:          nil,
					testInstructionsExecutionLogContainer:                    nil,
					testInstructionsExecutionAttributesContainerScroll:       nil,
					testInstructionsExecutionAttributesContainer:             nil,
					testInstructionsExecutionDetailsContainer:                nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				}

				// Create a new window for the TestCaseExecution
				openedDetailedTestCaseExecutionPtr.externalWindow = fenixApp.NewWindow(fmt.Sprintf(
					"TestCaseExecution '%s'", testCaseExecutionUuid))

				// Set boolean to indicate that window is open
				openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInExternalWindow = true

				// Catch the close action
				openedDetailedTestCaseExecutionPtr.externalWindow.SetOnClosed(func() {

					// Remove Content container object from map and mark not there
					defer func() {
						openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInExternalWindow = false
						openedDetailedTestCaseExecutionPtr.
							TestCaseInstructionPreViewObjectInExternalWindow = &TestCaseInstructionPreViewStruct{
							testCasePreviewTestInstructionExecutionLogSplitContainer: nil,
							testCaseExecutionPreviewContainerScroll:                  nil,
							testCaseExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
							testInstructionsExecutionDetailsContainerScroll:          nil,
							testInstructionsExecutionLogContainer:                    nil,
							testInstructionsExecutionAttributesContainerScroll:       nil,
							testInstructionsExecutionAttributesContainer:             nil,
							testInstructionsExecutionDetailsContainer:                nil,
							preViewTabs:                       nil,
							attributeExplorerTab:              nil,
							logsExplorerTab:                   nil,
							testInstructionDetailsExplorerTab: nil,
						}
						openedDetailedTestCaseExecutionPtr.externalWindow = nil
					}()

					//openedDetailedTestCaseExecutionPtr.externalWindow.Close()

				})

				// Generate base objects for the PreView
				openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInExternalWindow.
					testCasePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(
					openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInExternalWindow)

				// Generate a copy of the TestCaseExecutionPreview
				openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInExternalWindow.
					GenerateTestCaseExecutionPreviewContainer(
						testCaseExecutionUuid,
						testCaseExecutionVersion,
						testCaseExecutionsModelRef,
						fromExternalWindow,
						&openedDetailedTestCaseExecutionPtr.externalWindow)

				// Save the Object back to the Map
				openedDetailedTestCaseExecutionsMap[openedDetailedTestCaseExecutionsMapKey] = openedDetailedTestCaseExecutionPtr

				// Set TestCaseExecution as content
				var externalWindowCanvas fyne.Canvas
				externalWindowCanvas = openedDetailedTestCaseExecutionPtr.externalWindow.Canvas()

				externalWindowCanvas.SetContent(container.NewBorder(
					widget.NewLabel("This is a secondary window."),
					widget.NewButton("Close TestCaseExecution", func() {
						openedDetailedTestCaseExecutionPtr.externalWindow.Close()
					}),
					nil, nil,
					openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInExternalWindow.
						testCasePreviewTestInstructionExecutionLogSplitContainer,
				))

				// Set size and show
				openedDetailedTestCaseExecutionPtr.externalWindow.Resize(fyne.NewSize(1000, 800))
				openedDetailedTestCaseExecutionPtr.externalWindow.Show()
			}

		}

		// Define function to open a TestCaseExecution in a Tab
		var openTestCaseExecutionInTabFunction func()
		openTestCaseExecutionInTabFunction = func() {

			// Object exist, but check if the Tab already exist
			if openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInTab == true {
				// Tab exist, so get the tab so make it the visible one
				detailedTestCaseExecutionsUITabObjectRef.Select(openedDetailedTestCaseExecutionPtr.tabItem)

			} else {

				// Initialize Tab-object
				openedDetailedTestCaseExecutionPtr.
					TestCaseInstructionPreViewObjectInTab = &TestCaseInstructionPreViewStruct{
					testCasePreviewTestInstructionExecutionLogSplitContainer: nil,
					testCaseExecutionPreviewContainerScroll:                  nil,
					testCaseExecutionPreviewContainer:                        container.New(layout.NewVBoxLayout()),
					testInstructionsExecutionDetailsContainerScroll:          nil,
					testInstructionsExecutionLogContainer:                    nil,
					testInstructionsExecutionAttributesContainerScroll:       nil,
					testInstructionsExecutionAttributesContainer:             nil,
					testInstructionsExecutionDetailsContainer:                nil,
					preViewTabs:                       nil,
					attributeExplorerTab:              nil,
					logsExplorerTab:                   nil,
					testInstructionDetailsExplorerTab: nil,
				}

				// Generate base objects for the PreView
				openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInTab.
					testCasePreviewTestInstructionExecutionLogSplitContainer = generatePreViewObject(
					openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInTab)

				// Generate a copy of the TestCaseExecutionPreview
				openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInTab.
					GenerateTestCaseExecutionPreviewContainer(
						testCaseExecutionUuid,
						testCaseExecutionVersion,
						testCaseExecutionsModelRef,
						fromTab,
						sharedCode.FenixMasterWindowPtr)

				// Save the Object back to the Map
				openedDetailedTestCaseExecutionsMap[openedDetailedTestCaseExecutionsMapKey] = openedDetailedTestCaseExecutionPtr

				// Create a new Tab used for a specific TestCaseExecution
				var newDetailedTestCaseExecutionsTab *container.TabItem
				newDetailedTestCaseExecutionsTab = container.NewTabItem(
					testCaseExecutionUuid,
					container.NewBorder(
						nil,
						nil,
						nil,
						nil,
						openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInTab.
							testCasePreviewTestInstructionExecutionLogSplitContainer))

				// Add Tab to TabObjects TestCaseExecutions
				detailedTestCaseExecutionsUITabObjectRef.Append(newDetailedTestCaseExecutionsTab)

				// Extract Map with exit-functions for Tabs
				var exitFunctionsForDetailedTestCaseExecutionsUITabObject map[*container.TabItem]func()
				exitFunctionsForDetailedTestCaseExecutionsUITabObject = *exitFunctionsForDetailedTestCaseExecutionsUITabObjectPtr

				// Add Close function to TabItem
				exitFunctionsForDetailedTestCaseExecutionsUITabObject[newDetailedTestCaseExecutionsTab] = func() {

					// Remove and indicate that no Tab is onpen
					openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInTab = false
					openedDetailedTestCaseExecutionPtr.TestCaseInstructionPreViewObjectInTab = nil

					// Delete the Exit-function for the Tab
					defer func() {
						delete(exitFunctionsForDetailedTestCaseExecutionsUITabObject, newDetailedTestCaseExecutionsTab)
					}()

				}

				// Select the newly created TabItem
				detailedTestCaseExecutionsUITabObjectRef.Select(newDetailedTestCaseExecutionsTab)

				// Refresh the Tab
				detailedTestCaseExecutionsUITabObjectRef.Refresh()

				// Store the pointer to the TabItem
				openedDetailedTestCaseExecutionPtr.tabItem = newDetailedTestCaseExecutionsTab

				// Set boolean to indicate that window is open
				openedDetailedTestCaseExecutionPtr.isTestCaseExecutionOpenInTab = true

			}

		}

		// Define the menu items to open TestCaseExecution in Tab/External Window
		var items []*fyne.MenuItem

		// From where is the opening of the TestCaseExecution initiated; FromExecutionList, FromExternalWindow, FromTab
		switch openedTestCaseExecutionFrom {

		case fromExecutionList:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestCaseExecution in Tab", openTestCaseExecutionInTabFunction),
				fyne.NewMenuItem("Open TestCaseExecution in separate window", openTestCaseExecutionInExternalWindowFunction),
			}

		case fromExternalWindow:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestCaseExecution in Tab", openTestCaseExecutionInTabFunction),
			}

		case fromTab:
			// Define the popup menu
			items = []*fyne.MenuItem{
				fyne.NewMenuItem("Open TestCaseExecution in separate window", openTestCaseExecutionInExternalWindowFunction),
			}

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"Id":                          "f6c3f4ec-91c3-4b2a-dab-0aef96453a2a",
				"openedTestCaseExecutionFrom": openedTestCaseExecutionFrom,
			}).Fatalln("Unhandled 'openedTestCaseExecutionFrom', should never happen")

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
		testCaseInstructionPreViewObjectRef.testCaseExecutionPreviewContainer.Objects[0] = container.NewBorder(
			container.NewVBox(container.NewCenter(tempTopHeaderContainer), testCaseExecutionPreviewTopContainer, widget.NewSeparator()),
			container.NewVBox(widget.NewSeparator(), testCaseExecutionPreviewBottomContainer), nil, nil,
			testCaseMainAreaForPreviewScrollContainer)

		// Create a new temporary container for the logs
		testCaseInstructionPreViewObjectRef.testInstructionsExecutionLogContainer.Objects[0] = container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))

	} else {

		// No row is selected and only an info text should be shown
		testCaseInstructionPreViewObjectRef.testCaseExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestCaseExecution to get the Preview"))

		// Create a new temporary container for the logs
		testCaseInstructionPreViewObjectRef.testInstructionsExecutionLogContainer.Objects[0] = container.NewCenter(
			widget.NewLabel("Select a TestInstructionExecution or a TesInstructionContainer to get the Logs"))
	}

	// Add attributes container to attributes-container-map-ptr
	testCaseExecutionAttributesForPreviewMapPtr = &testCaseExecutionAttributesForPreviewMap

	// Refresh the 'testCaseExecutionPreviewContainer'
	testCaseInstructionPreViewObjectRef.testCaseExecutionPreviewContainer.Refresh()

}

// ClearTestCaseExecutionPreviewContainer
// Clears the preview area for the TestCaseExecution
func (testCaseInstructionPreViewObjectRef *TestCaseInstructionPreViewStruct) ClearTestCaseExecutionPreviewContainer() {
	testCaseInstructionPreViewObjectRef.testCaseExecutionPreviewContainer.Objects[0] = container.NewCenter(widget.NewLabel("Select a TestCaseExecution to get the Preview"))
}
