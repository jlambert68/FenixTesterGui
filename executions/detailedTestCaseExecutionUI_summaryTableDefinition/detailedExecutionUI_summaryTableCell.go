package detailedTestCaseExecutionUI_summaryTableDefinition

import (
	"FenixTesterGui/resources"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"image/color"
	"time"
)

var _ fyne.Widget = (*TestCaseExecutionSummaryTableCellStruct)(nil)

type TestCaseExecutionsBaseInformationStruct struct {
	// BaseInformation for TestCase
	TestCaseExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestCaseExecutionBasicInformationMessage

	// map[ExecutionStatusUpdateTimeStamp]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage
	AllTestCaseExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestCaseExecutionDetailsMessage
}

// TestCaseExecutionsStatusForSummaryTableStruct
// The definition used in SummaryTable to represent one TestCaseExecution and its current execution status
type TestCaseExecutionsStatusForSummaryTableStruct struct {
	TestCaseUIName                                          string
	TestCaseStatusValue                                     uint32
	ExecutionStatusUpdateTimeStamp                          time.Time
	TestCaseExecutionUuid                                   string
	TestCaseExecutionVersion                                uint32
	SortOrder                                               string
	TestInstructionExecutionsStatusForSummaryTableReference *[]*TestInstructionExecutionsStatusForSummaryTableStruct
}

type TestInstructionExecutionsBaseInformationStruct struct {
	// BaseInformation for TestInstruction
	TestInstructionExecutionBasicInformation *fenixExecutionServerGuiGrpcApi.TestInstructionExecutionBasicInformationMessage

	// map[ExecutionStatusUpdateTimeStamp]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage
	AllTestInstructionsExecutionsStatusUpdatesInformationMap map[string]*fenixExecutionServerGuiGrpcApi.TestInstructionExecutionsInformationMessage

	LatestStatusUpdateTimeStamp time.Time
}

// TestInstructionExecutionsStatusForSummaryTableStruct
// The definition used in SummaryTable to represent one TestInstructionExecution and its current execution status
type TestInstructionExecutionsStatusForSummaryTableStruct struct {
	TestInstructionExecutionUIName  string
	TestInstructionName             string
	TestInstructionExecutionUuid    string
	TestInstructionExecutionVersion uint32
	TestInstructionStatusValue      uint32
	ExecutionStatusUpdateTimeStamp  time.Time
	SortOrder                       string
}

type TestCaseExecutionSummaryTableCellStruct struct {
	widget.BaseWidget
	Label                                   *widget.Label
	testInstructionExecutionNames           []fyne.CanvasObject
	backgroundColorRectangle                *canvas.Rectangle
	showDetailedTestCaseExecution           *canvas.Image
	subscriptionForTestCaseExecution        *canvas.Image
	rowNumber                               int
	TestCaseExecutionMapKey                 string
	TestCaseExecutionMapKeyLabel            *widget.Label
	MyLabels                                []fyne.CanvasObject
	MyLabelsTextValues                      []string
	testCaseExecutionsDetails               *TestCaseExecutionsDetailsStruct
	testCaseExecutionsStatusForSummaryTable *TestCaseExecutionsStatusForSummaryTableStruct
	//testInstructionExecutionsStatusForSummaryTable *[]*TestInstructionExecutionsStatusForSummaryTable

}

func (t *TestCaseExecutionSummaryTableCellStruct) Tapped(_ *fyne.PointEvent) {
	fmt.Println("I was clicked!!! ", t.Label)

}

func (t *TestCaseExecutionSummaryTableCellStruct) TappedSecondary(_ *fyne.PointEvent) {
	fmt.Println("I was Right clicked!!!")

}

func (testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) DoubleTapped(_ *fyne.PointEvent) {

	fmt.Println("I was Double clicked!!!")

}

var backgroundRectangleBaseColor = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0x33,
}

var headerBackgroundRectangleBaseColor = color.RGBA{
	R: 0x33,
	G: 0x33,
	B: 0x33,
	A: 0x88,
}

func FlashAddedRow(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

	go func(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0x00,
				G: 0xFF,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				testcaseExecutionSummaryTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(testcaseExecutionSummaryTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(testcaseExecutionSummaryTableCell)

}

func FlashRowToBeRemoved(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {
	go func(testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) {

		// Define how the Color-flash should look like
		rectangleColorAnimation := canvas.NewColorRGBAAnimation(backgroundRectangleBaseColor,
			color.RGBA{
				R: 0xFF,
				G: 0x00,
				B: 0x00,
				A: 0xAA,
			}, time.Millisecond*200, func(animationColorValue color.Color) {
				testcaseExecutionSummaryTableCell.backgroundColorRectangle.FillColor = animationColorValue
				canvas.Refresh(testcaseExecutionSummaryTableCell.backgroundColorRectangle)
			})

		// Initiate Color-flash
		rectangleColorAnimation.AutoReverse = true
		rectangleColorAnimation.Start()

	}(testcaseExecutionSummaryTableCell)

}

func NewTestcaseExecutionSummaryTableCell(text string) *TestCaseExecutionSummaryTableCellStruct {
	newtestcaseExecutionSummaryTableCell := &TestCaseExecutionSummaryTableCellStruct{
		Label:                            widget.NewLabel(text),
		backgroundColorRectangle:         canvas.NewRectangle(backgroundRectangleBaseColor),
		showDetailedTestCaseExecution:    canvas.NewImageFromResource(resources.ResourceIcons8CheckMarkButton48Png),
		subscriptionForTestCaseExecution: canvas.NewImageFromResource(resources.ResourceIcons8CheckMarkButton48Png),
		TestCaseExecutionMapKeyLabel:     widget.NewLabel("text"),
		MyLabels:                         []fyne.CanvasObject{},
	}
	newtestcaseExecutionSummaryTableCell.MyLabels = make([]fyne.CanvasObject, 10)
	newtestcaseExecutionSummaryTableCell.MyLabels[0] = widget.NewLabel("SKapad när..1..")
	newtestcaseExecutionSummaryTableCell.MyLabels[1] = widget.NewLabel("SKapad när..2..")

	// Hide the showDetailedTestCaseExecution-image and set it to fill its parent
	newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution.FillMode = canvas.ImageFillContain
	newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution.Hide()

	// Hide the subscriptionForTestCaseExecution-image and set it to fill its parent
	newtestcaseExecutionSummaryTableCell.subscriptionForTestCaseExecution.FillMode = canvas.ImageFillContain
	newtestcaseExecutionSummaryTableCell.subscriptionForTestCaseExecution.Hide()

	newtestcaseExecutionSummaryTableCell.ExtendBaseWidget(newtestcaseExecutionSummaryTableCell)
	return newtestcaseExecutionSummaryTableCell

}

func (newtestcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct) CreateRenderer() fyne.WidgetRenderer {

	// Generate slice of TestInstructionExecutions
	/*
		var testInstructionExecutionNames []fyne.CanvasObject

		if newtestcaseExecutionSummaryTableCell.testCaseExecutionsDetails != nil {
			var testInstructionExecutionsStatusForSummaryTableSReference *[]*TestInstructionExecutionsStatusForSummaryTableStruct
			testInstructionExecutionsStatusForSummaryTableSReference = newtestcaseExecutionSummaryTableCell.testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable.TestInstructionExecutionsStatusForSummaryTable

			var testInstructionExecutionsStatusForSummaryTable []*TestInstructionExecutionsStatusForSummaryTableStruct
			testInstructionExecutionsStatusForSummaryTable = *testInstructionExecutionsStatusForSummaryTableSReference

			for _, testInstructionNameRef := range testInstructionExecutionsStatusForSummaryTable {
				testInstructionExecutionNames = append(testInstructionExecutionNames,
					container.NewMax(widget.NewLabel(testInstructionNameRef.TestInstructionExecutionUIName)))
			}
		}

	*/

	myNewVBox := container.NewVBox()
	myNewVBox.Add(newtestcaseExecutionSummaryTableCell.Label)
	myNewVBox.Add(newtestcaseExecutionSummaryTableCell.TestCaseExecutionMapKeyLabel)
	if len(newtestcaseExecutionSummaryTableCell.MyLabelsTextValues) > 0 {
		for i, _ := range newtestcaseExecutionSummaryTableCell.MyLabelsTextValues {
			//newtestcaseExecutionSummaryTableCell.MyLabels[i] = widget.NewLabel(textValue)
			myNewVBox.Add(newtestcaseExecutionSummaryTableCell.MyLabels[i])
		}

	}
	myNewVBox.Add(newtestcaseExecutionSummaryTableCell.MyLabels[0])

	// Use standard cell
	return &testcaseExecutionSummaryTableCellRenderer{
		testcaseExecutionSummaryTableCell: newtestcaseExecutionSummaryTableCell,
		container: container.NewMax(
			myNewVBox,
			newtestcaseExecutionSummaryTableCell.backgroundColorRectangle), //newtestcaseExecutionSummaryTableCell.showDetailedTestCaseExecution,
	}
}

var _ fyne.WidgetRenderer = (*testcaseExecutionSummaryTableCellRenderer)(nil)

type testcaseExecutionSummaryTableCellRenderer struct {
	testcaseExecutionSummaryTableCell *TestCaseExecutionSummaryTableCellStruct
	container                         *fyne.Container
}

func (r *testcaseExecutionSummaryTableCellRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *testcaseExecutionSummaryTableCellRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *testcaseExecutionSummaryTableCellRenderer) Refresh() {
	r.container.Refresh()
}

func (r *testcaseExecutionSummaryTableCellRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *testcaseExecutionSummaryTableCellRenderer) Destroy() {}
