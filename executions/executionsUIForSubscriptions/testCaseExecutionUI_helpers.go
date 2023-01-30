package executionsUIForSubscriptions

import (
	sharedCode "FenixTesterGui/common_code"
	"FenixTesterGui/executions/executionsModelForSubscriptions"
	"FenixTesterGui/headertable"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
)

// Remove item from the DataItem-slice and keep order
func remove(slice []binding.DataMap, s int) []binding.DataMap {
	return append(slice[:s], slice[s+1:]...)
}

const headerColumnExtraWidth float32 = 75

func ResizeTableColumns(t *headertable.SortingHeaderTable) {

	// Set Column widths
	bindings := t.TableOpts.Bindings
	numberOfRows, _ := t.Data.Length()
	var columnWidthToBeUsed float32
	var totalTableWidth float32

	// Loop columns
	for i, colAttr := range t.TableOpts.ColAttrs {

		// Loop Rows to get MaxTestDataWidth
		var currentColumnsMaxWidth float32
		var tempColumnWidth float32
		for rowCounter := 0; rowCounter < numberOfRows; rowCounter++ {
			b1 := bindings[rowCounter]
			d1, err := b1.GetItem(colAttr.Name)
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id":      "584ebb7e-9e35-4c60-b1a7-0713f973d838",
					"colAttr": colAttr,
					"b1":      b1,
				}).Fatalln("Couldn't get TestCaseExecution data due to no match")
			}

			cellData, err := d1.(binding.String).Get()
			if err != nil {
				sharedCode.Logger.WithFields(logrus.Fields{
					"id": "14e392a6-390e-4321-96c1-1da2bcbd33e1",
					"d1": d1,
				}).Fatalln("Couldn't get TestCaseExecution data due to no match")
			}

			// Check if width for row data is greater than previous max width for column
			tempColumnWidth = widget.NewLabel(cellData).MinSize().Width
			if tempColumnWidth > currentColumnsMaxWidth {
				currentColumnsMaxWidth = tempColumnWidth
			}
		}

		// Get HeaderWidth
		headerWidth := widget.NewLabel(colAttr.Header).MinSize().Width + headerColumnExtraWidth

		// Decide to used HeaderWidth or DataWidth
		if headerWidth > currentColumnsMaxWidth {
			columnWidthToBeUsed = headerWidth
		} else {
			columnWidthToBeUsed = currentColumnsMaxWidth
		}

		// Add to total Table Width
		totalTableWidth = totalTableWidth + columnWidthToBeUsed

		// Set Width for Header and data column
		t.Header.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*columnWidthToBeUsed)
		t.Data.SetColumnWidth(i, float32(colAttr.WidthPercent)/100.0*columnWidthToBeUsed)

	}

	//t.Resize(fyne.NewSize(totalTableWidth, 200))
	t.Header.Resize(fyne.NewSize(totalTableWidth, t.Header.MinSize().Height))
	t.Data.Resize(fyne.NewSize(totalTableWidth, t.Data.MinSize().Height*float32(len(t.TableOpts.Bindings))))
}

// Check that TestCaseExecution doesn't already exist in any of OnQueue, UnderExecution or FinishedExecutions

func verifyThatTestCaseExecutionIsNotInUse(subscriptionsForTestCaseExecutionMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapOverallType) (err error) {

	// Variables needed
	var existInMap bool
	var errorId string
	var newErrorMessage string
	var fullErrorMessage string
	var testCaseExecutionMapKey executionsModelForSubscriptions.TestCaseExecutionMapKeyType
	var subscriptionsForTestCaseExecutionDetailsMap executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsType
	var shouldExistInTable bool
	var xorValidation bool

	// Verify that there was exact one testCaseExecutionMapKey
	var numberOfKeys int
	numberOfKeys = len(subscriptionsForTestCaseExecutionMap)
	if numberOfKeys != 1 {

		errorId = "f5b6464a-f98a-404e-99dc-a7f0eb1d9c6a"
		err = errors.New(fmt.Sprintf("expected  exact One 'testCaseExecutionMapKey', but got '%s' in  subscriptionsForTestCaseExecutionMap: '%s' [ErrorID: %s]. ", numberOfKeys, subscriptionsForTestCaseExecutionMap, errorId))

		fmt.Println(err) // TODO Send on Error Channel

		return err
	}

	// Loop over the "only element in map
	for testCaseExecutionMapKey, subscriptionsForTestCaseExecutionDetailsMap = range subscriptionsForTestCaseExecutionMap {

		// Loop over the Executions table to get if 'testCaseExecutionMapKey' should be in the table or not
		var subscriptionTableType executionsModelForSubscriptions.SubscriptionTableType
		var testCaseExecutionsDetails executionsModelForSubscriptions.SubscriptionsForTestCaseExecutionMapDetailsStruct
		for subscriptionTableType, testCaseExecutionsDetails = range subscriptionsForTestCaseExecutionDetailsMap {

			shouldExistInTable = testCaseExecutionsDetails.ShouldExistInTable

			switch subscriptionTableType {

			/*
				XOR =  (ShouldExistInTable AND NOT(existInMap)) OR (NOT(ShouldExistInTable) AND existInMap)
				ShouldExistInTable	existInMap	Error
				0					0			0
				0					1			1
				1					0			1
				1					1			0
			*/

			case executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionOnQueueTable:

				// Check if TestCaseExecution already exists in 'OnQueue'
				_, existInMap = executionsModelForSubscriptions.TestCaseExecutionsOnQueueMapAdaptedForUiTable[testCaseExecutionMapKey]

				xorValidation = shouldExistInTable && !existInMap || !shouldExistInTable && existInMap
				if xorValidation == true {

					errorId = "11865739-f07f-4042-83ac-6b4c4537e94f"
					newErrorMessage = fmt.Sprintf("the rule for validating if 'testCaseExecutionMapKey', '%s' "+
						"should or should not exisit in TestCaseExecutionsOnQueueMapAdaptedForUiTable failed. "+
						"shouldExistInTable='%t' and existInMap='%t' [ErrorID: %s]. ",
						testCaseExecutionMapKey, shouldExistInTable, existInMap, errorId)

					fullErrorMessage = fullErrorMessage + newErrorMessage

					err = errors.New("error found")

				}

			case executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionUnderExecutionTable:

				// Check if TestCaseExecution exists in 'UnderExecution'
				_, existInMap = executionsModelForSubscriptions.TestCaseExecutionsUnderExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]

				xorValidation = shouldExistInTable && !existInMap || !shouldExistInTable && existInMap
				if xorValidation == true {

					errorId = "f7ee589b-8075-44d8-94a3-2a3556358d00"
					newErrorMessage = fmt.Sprintf("the rule for validating if 'testCaseExecutionMapKey', '%s' "+
						"should or should not exisit in TestCaseExecutionsUnderExecutionMapAdaptedForUiTable failed. "+
						"shouldExistInTable='%t' and existInMap='%t' [ErrorID: %s]. ",
						testCaseExecutionMapKey, shouldExistInTable, existInMap, errorId)

					fullErrorMessage = fullErrorMessage + newErrorMessage

					err = errors.New("error found")

				}

			case executionsModelForSubscriptions.SubscriptionTableForTestCaseExecutionFinishedExecutionsTable:

				// Check if TestCaseExecution already exists in 'FinishedExecutions'
				_, existInMap = executionsModelForSubscriptions.TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable[testCaseExecutionMapKey]

				xorValidation = shouldExistInTable && !existInMap || !shouldExistInTable && existInMap
				if xorValidation == true {

					errorId = "aa576012-3aa0-4388-91a7-9d67b6258dd7"
					newErrorMessage = fmt.Sprintf("the rule for validating if 'testCaseExecutionMapKey', '%s' "+
						"should or should not exisit in TestCaseExecutionsFinishedExecutionMapAdaptedForUiTable failed. "+
						"shouldExistInTable='%t' and existInMap='%t' [ErrorID: %s]. ",
						testCaseExecutionMapKey, shouldExistInTable, existInMap, errorId)

					fullErrorMessage = fullErrorMessage + newErrorMessage

					err = errors.New("error found")

				}

			default:
				errorId = "53a5da5f-0f5e-4017-ae8d-09752f4f7925"
				err = errors.New(fmt.Sprintf("unhanlded subscriptionTableType '%s'  in subscriptionsForTestCaseExecutionDetailsMap '%s' [ErrorID: %s]. ", subscriptionTableType, subscriptionsForTestCaseExecutionDetailsMap, errorId))

				fmt.Println(err) // TODO Send on Error Channel
			}
		}

		break
	}

	// There were at least on match in any of OnQueue, UnderExecution or FinishedExecutions. So create new full error message
	if err != nil {

		err = errors.New(fullErrorMessage)
		// fmt.Println(err) // TODO Send on Error Channel
	}

	return err

}
