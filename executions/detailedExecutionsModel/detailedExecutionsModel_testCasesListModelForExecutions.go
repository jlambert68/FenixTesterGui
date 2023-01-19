package detailedExecutionsModel

import (
	"FenixTesterGui/executions/detailedTestCaseExecutionUI_summaryTableDefinition"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"image/color"
)

/*
func CreateSummaryTableForDetailedTestCaseExecutionsList() *fyne.Container {
	var tableForTestCaseExecutionsSummaryBindings []binding.DataMap

	// Create a binding for each TestCaseExecutionsSummary data
	for _, tempDetailedTestCaseExecutionReference := range detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap {

		tableForTestCaseExecutionsSummaryBindings = append(
			tableForTestCaseExecutionsSummaryBindings,
			binding.BindStruct(tempDetailedTestCaseExecutionReference.TestCaseExecutionsStatusForSummaryTable))
	}

	DetailedTestCaseExecutionsSummaryTableOptions.Bindings = tableForTestCaseExecutionsSummaryBindings

	ht := detailedTestCaseExecutionUI_summaryTableDefinition.NewTestCaseExecutionsSummaryTable(
		&DetailedTestCaseExecutionsSummaryTableOptions)

	detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsSummaryTable = ht

	mySortTable := container.NewMax(ht, layout.NewSpacer())
	mySortTable.Resize(ht.Size())
	//ht.Header.ScrollToTrailing()
	ht.Data.Refresh()
	//ht.Header.ScrollToLeading()

	return mySortTable

}

*/

/*
	#b6d7a8	INITIATED = 0; // All set up for execution, but has not been triggered to start execution
	#ffff00	EXECUTING = 1; // TestInstruction is execution
	#4a86e8	CONTROLLED_INTERRUPTION = 2; // Interrupted by in a controlled way
	#4a86e8	CONTROLLED_INTERRUPTION_CAN_BE_RERUN = 3; // Interrupted by in a controlled way, but can be rerun
	#00ff00	FINISHED_OK = 4; // Finish as expected to TestInstruction definition
	#00ff00	FINISHED_OK_CAN_BE_RERUN = 5; // Finish as expected to TestInstruction definition, but can be rerun
	#ff0000	FINISHED_NOT_OK = 6; // Finish with errors in validations
	#ff0000	FINISHED_NOT_OK_CAN_BE_RERUN = 7; // Finish with errors in validations, but can be rerun
	#9900ff	UNEXPECTED_INTERRUPTION = 8; // The TestInstruction stopped executed in an unexpected way
	#9900ff	UNEXPECTED_INTERRUPTION_CAN_BE_RERUN = 9; // The TestInstruction stopped executed in an unexpected way, but can be rerun
	#fbbc04	TIMEOUT_INTERRUPTION = 10; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}
	#fbbc04	TIMEOUT_INTERRUPTION_CAN_BE_RERUN = 11; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}, but can be rerun
*/

// ExecutionStatusColorMapStruct
// Holds the structure for one ExecutionStatus-definition
type ExecutionStatusColorMapStruct struct {
	TextColor       color.RGBA
	BackgroundColor color.RGBA
	StrokeColor     color.RGBA
	UseStroke       bool
}

const backgroundStrokeWidth = 4

// ExecutionStatusColorMap
// map[int32]ExecutionStatusColorMapStruct
// Holds the definitions for which colors should be used in UI for executions
var ExecutionStatusColorMap = map[int32]ExecutionStatusColorMapStruct{
	// 'INITIATED = 0'
	0: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0x00,
			A: 0xFF},
		UseStroke: false,
	},

	//  'EXECUTING = 1'
	1: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0x00,
			A: 0xFF},
		UseStroke: false,
	},

	// 'CONTROLLED_INTERRUPTION = 2'
	2: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x4a,
			G: 0x86,
			B: 0xe8,
			A: 0xFF},
		UseStroke: false,
	},

	// 'CONTROLLED_INTERRUPTION_CAN_BE_RERUN = 3'
	3: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x4a,
			G: 0x86,
			B: 0xe8,
			A: 0xFF},
		StrokeColor: color.RGBA{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF},
		UseStroke: true,
	},

	// 'FINISHED_OK = 4'
	4: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x00,
			G: 0xff,
			B: 0x00,
			A: 0xFF},
		UseStroke: false,
	},

	// 'FINISHED_OK_CAN_BE_RERUN = 5'
	5: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x00,
			G: 0xff,
			B: 0x00,
			A: 0xFF},
		StrokeColor: color.RGBA{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF},
		UseStroke: true,
	},

	// 'FINISHED_NOT_OK = 6'
	6: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xff,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		UseStroke: false,
	},

	// 'FINISHED_NOT_OK_CAN_BE_RERUN = 7'
	7: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xff,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		StrokeColor: color.RGBA{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF},
		UseStroke: true,
	},

	// 'UNEXPECTED_INTERRUPTION = 8'
	8: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x99,
			G: 0x00,
			B: 0xff,
			A: 0xFF},
		UseStroke: false,
	},

	// 'UNEXPECTED_INTERRUPTION_CAN_BE_RERUN = 9'
	9: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0x99,
			G: 0x00,
			B: 0xff,
			A: 0xFF},
		StrokeColor: color.RGBA{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF},
		UseStroke: true,
	},

	// 'TIMEOUT_INTERRUPTION = 10'
	10: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xfb,
			G: 0xbc,
			B: 0x04,
			A: 0xFF},
		UseStroke: false,
	},

	// 'TIMEOUT_INTERRUPTION_CAN_BE_RERUN = 11'
	11: ExecutionStatusColorMapStruct{
		TextColor: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF},
		BackgroundColor: color.RGBA{
			R: 0xfb,
			G: 0xbc,
			B: 0x04,
			A: 0xFF},
		StrokeColor: color.RGBA{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF},
		UseStroke: true,
	},
}

var testCaseRowBackgroundColorOddRow = color.RGBA{
	R: 0x11,
	G: 0x11,
	B: 0x11,
	A: 0x33}

var testCaseRowBackgroundColorEvenRow = color.RGBA{
	R: 0x22,
	G: 0x22,
	B: 0x22,
	A: 0x33}

func createExecutionSummary(
	testInstructionName string,
	testInstructionExecutionStatus int32) (
	tempItemExecutionSummary *fyne.Container,
	err error) {

	var testInstructionNameToUse string
	testInstructionNameToUse = " " + testInstructionName + ""

	var tempExecutionStatusColors ExecutionStatusColorMapStruct
	var existInMap bool

	tempExecutionStatusColors, existInMap = ExecutionStatusColorMap[testInstructionExecutionStatus]

	if existInMap == false {
		ErrorID := "22200080-6e9d-4654-aa1e-ddf89918e98b"
		err = errors.New(fmt.Sprintf("couldn't find 'TestInstructionExecutionStatus' within 'ExecutionStatusColorMap'. Got message: '%s'. [ErrorID:'%s']", testInstructionExecutionStatus, ErrorID))

		fmt.Println(err) // TODO Send on Error-channel

		return nil, err
	}

	// Create Text with correct color and Text
	var testInstructionText *canvas.Text
	testInstructionText = &canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color:     tempExecutionStatusColors.TextColor,
		Text:      testInstructionNameToUse,
		TextSize:  15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}

	// Create background with correct color
	var tempExecutionBackground *canvas.Rectangle
	tempExecutionBackground = canvas.NewRectangle(tempExecutionStatusColors.BackgroundColor)

	// Check if Stroke should be added
	if tempExecutionStatusColors.UseStroke == true {
		tempExecutionBackground.StrokeColor = tempExecutionStatusColors.StrokeColor
		tempExecutionBackground.StrokeWidth = 4
	}

	tempItemExecutionSummary = container.New(layout.NewMaxLayout(), tempExecutionBackground, testInstructionText)

	//return container.New(layout.NewMaxLayout(), widget.NewLabel("Dummy Text")), err
	return tempItemExecutionSummary, err
}

var backgroundRectangleBaseColorForOddRows = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0x33,
}

var headerBackgroundRectangleBaseColorForEvenRows = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0xff,
}

type summaryTableForDetailedTestCaseExecutionsStruct struct {
	testCaseNameLabel *widget.Label

	backgroundColorRectangle         *canvas.Rectangle
	showDetailedTestCaseExecution    *canvas.Image
	subscriptionForTestCaseExecution *canvas.Image

	testInstructionExecutionNames []fyne.CanvasObject
}

var summaryTableForDetailedTestCaseExecutions []*summaryTableForDetailedTestCaseExecutionsStruct

func CreateSummaryTableForDetailedTestCaseExecutionsList() (testcaseExecutionsSummaryReturnTable *fyne.Container) {

	var err error

	// Create the Header for the Summary table
	var summaryHeaderLabel *widget.Label
	summaryHeaderLabel = &widget.Label{
		BaseWidget: widget.BaseWidget{},
		Text:       "TestCaseExecutions Summary Table: ",
		Alignment:  0,
		Wrapping:   0,
		TextStyle:  fyne.TextStyle{Bold: true},
	}

	// Define the Summary Table
	var testcaseExecutionsSummaryTable *fyne.Container
	testcaseExecutionsSummaryTable = &fyne.Container{
		Hidden:  false,
		Layout:  layout.NewVBoxLayout(),
		Objects: nil,
	}

	// Add the Header to the Summary TableContainer
	testcaseExecutionsSummaryTable.Add(summaryHeaderLabel)

	// testCaseCounter
	var testCaseCounter int

	// Loop all TestCaseExecutions
	for _, tempDetailedTestCaseExecutionReference := range detailedTestCaseExecutionUI_summaryTableDefinition.TestCaseExecutionsDetailsMap {
		tempDetailedTestCaseExecution := *tempDetailedTestCaseExecutionReference

		testCaseCounter = testCaseCounter + 1

		// Create the Item for one Row in the Summary table
		var testCaseRow *fyne.Container
		testCaseRow = &fyne.Container{
			Hidden:  false,
			Layout:  layout.NewVBoxLayout(),
			Objects: nil,
		}

		// Extract TestCaseName
		var testCaseName string
		testCaseName = tempDetailedTestCaseExecution.TestCaseExecutionDatabaseResponseMessage.TestCaseExecutionBasicInformation.TestCaseName

		// Extract the status for the TestCase
		var testCaseStatus fenixExecutionServerGuiGrpcApi.TestCaseExecutionStatusEnum
		var detailedTestCaseExecutions []*fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage
		detailedTestCaseExecutions = tempDetailedTestCaseExecution.TestCaseExecutionDatabaseResponseMessage.TestCaseExecutionDetails

		// Get last row for TestCaseExecutions
		var lastRow int
		lastRow = len(tempDetailedTestCaseExecution.TestCaseExecutionDatabaseResponseMessage.TestCaseExecutionDetails) - 1

		// Extract the status
		testCaseStatus = detailedTestCaseExecutions[lastRow].TestCaseExecutionStatus

		// Create the TestCaseExecution-container
		var testCaseNameContainer *fyne.Container
		testCaseNameContainer, err = createExecutionSummary(testCaseName, int32(testCaseStatus))
		if err != nil {
			return nil
		}

		// Encapsulate TestCaseStatus-field into HBOX-container
		var testCaseNameContainerToBeAdded *fyne.Container
		testCaseNameContainerToBeAdded = &fyne.Container{
			Hidden:  false,
			Layout:  layout.NewHBoxLayout(),
			Objects: []fyne.CanvasObject{testCaseNameContainer},
		}

		// Add TestCaseName-container to Summary-row
		testCaseRow.Add(testCaseNameContainerToBeAdded)

		// Create the Item for all TestInstruction for one TestCase
		var testInstructionsForTestCase *fyne.Container
		testInstructionsForTestCase = &fyne.Container{
			Hidden:  false,
			Layout:  layout.NewHBoxLayout(),
			Objects: nil,
		}

		// Loop alla TestInstructionExecutions in TestCaseExecution
		for _, tempTestInstructionExecutionReference := range tempDetailedTestCaseExecution.TestCaseExecutionDatabaseResponseMessage.TestInstructionExecutions {
			tempTestInstructionExecution := *tempTestInstructionExecutionReference

			// Extract TestCaseName
			var testInstructionName string
			testInstructionName = tempTestInstructionExecution.TestInstructionExecutionBasicInformation.TestInstructionName

			// Extract the status for the TestCase
			var testInstructionStatus fenixExecutionServerGuiGrpcApi.TestInstructionExecutionStatusEnum
			var testInstructionExecutionsInformationMessages []*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
			testInstructionExecutionsInformationMessages = tempTestInstructionExecution.TestInstructionExecutionsInformation

			// Get last row for TestInstructionExecutions
			var lastTestInstructionRow int
			lastTestInstructionRow = len(testInstructionExecutionsInformationMessages) - 1

			// Extract TestInstructionStatus when there are something to extract
			if lastTestInstructionRow >= 0 {

				// Extract the status
				testInstructionStatus = testInstructionExecutionsInformationMessages[lastTestInstructionRow].TestInstructionExecutionStatus

				// Create the TestInstructionExecution-container
				var testInstructionNameContainer *fyne.Container
				testInstructionNameContainer, err = createExecutionSummary(testInstructionName, int32(testInstructionStatus))
				if err != nil {
					return nil
				}

				// Add TestInstructionName-container to containers for all TestInstructions for current Summary-row
				testInstructionsForTestCase.Add(testInstructionNameContainer)

			}
		}

		// Add the TestInstruction-container to the TestCase-container for the Row
		testCaseRow.Add(testInstructionsForTestCase)

		// Create background for TestCaseRow, with correct color
		var testCaseRowBackground *canvas.Rectangle

		// Check if even or odd row number
		if testCaseCounter%2 == 0 {
			// Even number
			testCaseRowBackground = canvas.NewRectangle(testCaseRowBackgroundColorEvenRow)
		} else {
			// Odd number
			testCaseRowBackground = canvas.NewRectangle(testCaseRowBackgroundColorOddRow)
		}

		// Encapsulate Row together with background color rectangle
		var testCaseRowContainer *fyne.Container
		testCaseRowContainer = container.New(layout.NewMaxLayout(), testCaseRowBackground, testCaseRow)

		// Add the Row to Summary-container
		testcaseExecutionsSummaryTable.Add(testCaseRowContainer)

	}

	// testcaseExecutionsSummaryTable.Refresh()

	// Define the Summary Return Table
	testcaseExecutionsSummaryReturnTable = &fyne.Container{
		Hidden:  false,
		Layout:  layout.NewHBoxLayout(),
		Objects: []fyne.CanvasObject{testcaseExecutionsSummaryTable},
	}

	return testcaseExecutionsSummaryReturnTable
}
