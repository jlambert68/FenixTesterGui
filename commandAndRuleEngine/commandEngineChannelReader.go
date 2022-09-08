package commandAndRuleEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/testCase/testCaseModel"
	"errors"
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"log"
	"sync"
)

// Channel reader which is used for reading out commands to CommandEngine
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) startCommandChannelReader() {

	var incomingChannelCommand sharedCode.ChannelCommandStruct

	for {
		// Wait for incoming command over channel
		incomingChannelCommand = <-*commandAndRuleEngine.CommandChannelReference

		switch incomingChannelCommand.ChannelCommand {

		case sharedCode.ChannelCommandNewTestCase:
			commandAndRuleEngine.channelCommandCreateNewTestCase(incomingChannelCommand)

		case sharedCode.ChannelCommandSwapElement:
			commandAndRuleEngine.channelCommandSwap(incomingChannelCommand)

		case sharedCode.ChannelCommandRemoveElement:
			commandAndRuleEngine.channelCommandRemove(incomingChannelCommand)

		// No other command is supported
		default:
			//TODO Send Error over ERROR-channel
		}
	}

}

// Execute command 'New TestCase' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandCreateNewTestCase(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	// Create New TestCase
	testCaseUuid, err := commandAndRuleEngine.NewTestCaseModel()

	if err != nil {
		fmt.Println(err)

		return

	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(testCaseUuid)
	if err != nil {
		fmt.Println(err)
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		CreateNewTestCaseUI:     true,
		ActiveTestCase:          testCaseUuid,
		TextualTestCaseSimple:   textualTestCaseSimple,
		TextualTestCaseComplex:  textualTestCaseComplex,
		TextualTestCaseExtended: textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Execute command 'Remove Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandRemove(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	currentTestCaseUuid := incomingChannelCommand.ActiveTestCase
	elementToRemove := incomingChannelCommand.FirstParameter

	// Delete Element from TestCase
	err := commandAndRuleEngine.DeleteElementFromTestCaseModel(currentTestCaseUuid, elementToRemove)
	if err != nil {
		//TODO Send ERRORS over error-channel
		fmt.Println(err)

		return
	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(currentTestCaseUuid)
	if err != nil {
		fmt.Println(err)
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		CreateNewTestCaseUI:     false,
		ActiveTestCase:          currentTestCaseUuid,
		TextualTestCaseSimple:   textualTestCaseSimple,
		TextualTestCaseComplex:  textualTestCaseComplex,
		TextualTestCaseExtended: textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

func printDropZone(index int) {
	log.Println("Chosen DropZone: is ", index)
}

// Execute command 'Swap Element' received from Channel reader
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) channelCommandSwap(incomingChannelCommand sharedCode.ChannelCommandStruct) {

	currentTestCaseUuid := incomingChannelCommand.ActiveTestCase
	elementUuidTobeSwappedOut := incomingChannelCommand.FirstParameter
	elementUuidToBeSwappedIn := incomingChannelCommand.SecondParameter
	elementType := incomingChannelCommand.ElementType

	// Get the ImmatureElement To Swap In
	var immatureElementToSwapInTestCaseFormat testCaseModel.ImmatureElementStruct

	switch elementType {

	case sharedCode.TestInstruction:
		tempMap := commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionsMap
		immatureElementToSwapInOriginal := tempMap[elementUuidToBeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = commandAndRuleEngine.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

		// Handle DropZones
		availableDropZones := tempMap[elementUuidToBeSwappedIn].ImmatureTestInstructionInformation.AvailableDropZones
		numberOfDropZones := len(availableDropZones)

		// If there are more than one DropZone then pop up window so user can choose
		if numberOfDropZones > 1 {
			numberOfDropZones = 2
		}

		// Create DropZone window
		dropZoneContainer := container.NewVBox()
		var dropZoneWindow dialog.Dialog
		var dropZoneWaitGroup sync.WaitGroup
		dropZoneWaitGroup.Add(1)
		chosenDropZoneName := "NO_DROPZONE_WAS_CHOSEN"

		switch numberOfDropZones {
		case 0:

			// Set the TestInstructionColor to transparent
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = "#00000000"
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = "No DropZone exists"
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = "No DropZone exists"

		case 1:
			// Move the uuid and color for the only DropZone
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = availableDropZones[0].DropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = availableDropZones[0].DropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = availableDropZones[0].DropZoneColor

		case 2:
			for _, dropZoneItem := range availableDropZones {

				dropZoneButton := &widget.Button{
					DisableableWidget: widget.DisableableWidget{},
					Text:              dropZoneItem.DropZoneName,
					Icon:              nil,
					Importance:        0,
					Alignment:         0,
					IconPlacement:     0,
					OnTapped:          nil,
				}

				// Create the DropZone Button
				dropZoneButton.OnTapped = func() {
					fmt.Println(dropZoneButton.Text)
					chosenDropZoneName = dropZoneButton.Text
					defer dropZoneWaitGroup.Done()
				}

				// Create the Background colored rectangle
				dropZoneColor, err := sharedCode.ConvertRGBAHexStringIntoRGBAColor(dropZoneItem.DropZoneColor)
				if err != nil {
					fmt.Println(err.Error()) // TODO Send on Error-channel
				}
				backgroundColorRectangle := canvas.NewRectangle(dropZoneColor)

				// Create the DropZoneContainer
				dropZoneButtonContainer := container.NewMax(backgroundColorRectangle, dropZoneButton)

				// Add the DropZoneButton-container to the object to be put up for user to chose DropZone from
				dropZoneContainer.Add(dropZoneButtonContainer)

			}

			dropZoneContainer.Add(widget.NewSeparator())

			// Open up the DropZone-choser to the user
			dropZoneWindow = dialog.NewCustom("Choose DropZone", "Cancel", dropZoneContainer, *commandAndRuleEngine.MasterFenixWindow)
			dropZoneWindow.Show()

			// Wait for user to choose a DropZone
			dropZoneWaitGroup.Wait()
			dropZoneWindow.Hide()

			// If no DropZone was chosen then just exit
			if chosenDropZoneName == "NO_DROPZONE_WAS_CHOSEN" {
				return
			}

			// Find correct DropZone
			var dropZoneColorAsHexString string
			var dropZoneUuid string
			var dropZoneName string
			for _, dropZoneItem := range availableDropZones {

				if dropZoneItem.DropZoneName == chosenDropZoneName {
					dropZoneColorAsHexString = dropZoneItem.DropZoneColor
					dropZoneUuid = dropZoneItem.DropZoneUuid
					dropZoneName = dropZoneItem.DropZoneName
					break
				}

			}

			// Set the DropZoneUuid and TestInstructionColor from Chosen DropZone
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneUuid = dropZoneUuid
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneName = dropZoneName
			immatureElementToSwapInTestCaseFormat.ChosenDropZoneColor = dropZoneColorAsHexString

		}

		/*

			// If number of DropZones for Droppen object is greater than one the Open DropZone-picker Window
			dropZone1 := widget.NewButton("DropZone 1", func() {
				log.Println("Chosen DropZone: is One")
			})

			stateMachineDragAndDrop. stateMachineDragAndDrop.SourceUuid

			dropZoneColor, err :=
			dropZone1Container := container.NewMax(canvas.NewRectangle(color.Gray{}), dropZone1)

			dropZone2 := widget.NewButton("DropZone 1", func() {
				log.Println("Chosen DropZone: is Two")
			})
			dropZoneContainer := container.NewVBox(dropZone1, dropZone2)

			dropZoneWindow := dialog.NewCustom("Choose DropZone", "Cancel", dropZoneContainer, *masterFenixWindow)
			dropZoneWindow.Show()
		*/

	case sharedCode.TestInstructionContainer:
		tempMap := commandAndRuleEngine.Testcases.AvailableImmatureTestInstructionContainersMap
		immatureElementToSwapInOriginal := tempMap[elementUuidToBeSwappedIn].ImmatureSubTestCaseModel
		immatureElementToSwapInTestCaseFormat = commandAndRuleEngine.convertGrpcElementModelIntoTestCaseElementModel(immatureElementToSwapInOriginal)

	default:

		errorId := "6e8ed2ec-84df-42eb-a95d-41ba6920a9cb"
		err := errors.New(fmt.Sprintf("unknown Building BLock Type: '%s' [ErrorID: %s]", elementType, errorId))

		fmt.Println(err) //TODO Send error over error-channel

		// Exit function
		return

	}

	// Execute Swap of Elements
	err := commandAndRuleEngine.SwapElementsInTestCaseModel(currentTestCaseUuid, elementUuidTobeSwappedOut, &immatureElementToSwapInTestCaseFormat)
	if err != nil {
		fmt.Println(err)

		return //TODO Send error over error-channel
	}

	// Update UI with TestCase Textual Representation
	textualTestCaseSimple, textualTestCaseComplex, textualTestCaseExtended, err := commandAndRuleEngine.Testcases.CreateTextualTestCase(currentTestCaseUuid)
	if err != nil {
		fmt.Println(err) //TODO Send error over error-channel
	}

	// Send 'update TestCase graphics' command over channel
	outgoingChannelCommandGraphicsUpdatedData := sharedCode.ChannelCommandGraphicsUpdatedStruct{
		CreateNewTestCaseUI:     false,
		ActiveTestCase:          currentTestCaseUuid,
		TextualTestCaseSimple:   textualTestCaseSimple,
		TextualTestCaseComplex:  textualTestCaseComplex,
		TextualTestCaseExtended: textualTestCaseExtended,
	}

	*commandAndRuleEngine.GraphicsUpdateChannelReference <- outgoingChannelCommandGraphicsUpdatedData

}

// Convert gRPC-message for TI or TIC into model used within the TestCase-model
func (commandAndRuleEngine *CommandAndRuleEngineObjectStruct) convertGrpcElementModelIntoTestCaseElementModel(immatureGrpcElementModelMessage *fenixGuiTestCaseBuilderServerGrpcApi.ImmatureElementModelMessage) (immatureTestCaseElementModel testCaseModel.ImmatureElementStruct) {

	// Initiate map used in TestCaseModel
	immatureTestCaseElementModel.ImmatureElementMap = make(map[string]fenixGuiTestCaseBuilderServerGrpcApi.ImmatureTestCaseModelElementMessage)

	// Loop over gRPC-element-model-structure
	for _, gRpcElementModel := range immatureGrpcElementModelMessage.TestCaseModelElements {
		immatureTestCaseElementModel.ImmatureElementMap[gRpcElementModel.ImmatureElementUuid] = *gRpcElementModel
	}

	// Set the first Element
	immatureTestCaseElementModel.FirstElementUuid = immatureGrpcElementModelMessage.FirstImmatureElementUuid

	return immatureTestCaseElementModel
}
