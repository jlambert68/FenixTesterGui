package testCaseModel

import (
	"errors"
	"fmt"
)

// SaveTestCase - Save the TestCase to the Database
func (testCaseModel *TestCasesModelsStruct) SaveTestCase(testCaseUuid string) (err error) {

	// Save changed Attributes, if there are any, into the TestCase-model under correct TestInstruction
	err = testCaseModel.SaveChangedTestCaseAttributeInTestCase(testCaseUuid)
	if err != nil {
		return err
	}

	// Get current TestCase
	_, existsInMap := testCaseModel.TestCases[testCaseUuid] //currentTestCase
	if existsInMap == false {

		errorId := "1ba54aac-6c59-40e5-8366-5b56ee2a7eaf"
		err = errors.New(fmt.Sprintf("testcase '%s' is missing in map with all TestCases [ErrorID: %s]", testCaseUuid, errorId))

		return err
	}

	// Save Attributes to model
	//testCaseModel.SaveTestCase()

	return err

}

// Save changed Attributes into the TestCase-model under correct TestInstruction
func (testCaseModel *TestCasesModelsStruct) SaveChangedTestCaseAttributeInTestCase(testCaseUuid string) (err error) {

	// Extract current TestCase
	testCase, existInMap := testCaseModel.TestCases[testCaseUuid]
	if existInMap == false {

		errorId := "40fc730f-87d4-4c44-96ff-ab1003e40751"
		err := errors.New(fmt.Sprintf("testCase %s is missing in TestCases-map [ErrorID: %s]", testCaseUuid, errorId))

		fmt.Println(err) //TODO Send error over error-channel
		return err
	}

	// Extract testInstructionElementMatureUuidUuid
	attributesList := *testCase.AttributesList
	testInstructionElementMatureUuidUuid := attributesList[0].TestInstructionElementMatureUuidUuid

	// Check if any attribute is changed
	if len(attributesList) > 0 {
		for _, attribute := range attributesList {
			if attribute.AttributeIsChanged == true {
				// Attribute is changed so save it,

				// Extract TestInstruction
				tempMatureTestInstruction, existInMap := testCase.MatureTestInstructionMap[testInstructionElementMatureUuidUuid]
				if existInMap == false {
					errorId := "83b64181-3a02-4b98-8eba-d1fbad61dcd5"
					err := errors.New(fmt.Sprintf("mature testInstruction %s is missing in MatureTestInstructionMap [ErrorID: %s]", testInstructionElementMatureUuidUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Extract  Attribute
				tempTestInstructionAttribute, existInMap := tempMatureTestInstruction.TestInstructionAttributesList[attribute.AttributeUuid]
				if existInMap == false {
					errorId := "77e03442-7ccc-46c7-891e-0c5e0dd5bd1c"
					err := errors.New(fmt.Sprintf("testInstruction attribute %s is missing in MatureTestInstructionMap [ErrorID: %s]", attribute.AttributeUuid, errorId))

					fmt.Println(err) //TODO Send error over error-channel
					return err
				}

				// Save changed value for Attribute
				tempTestInstructionAttribute.AttributeInformation.InputTextBoxProperty.TextBoxAttributeValue = attribute.AttributeChangedValue

			}
		}
	}

	return err

}
