package commandAndRuleEngine

import (
	"errors"
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
)

// VerifyIfComponentCanBeDeleted - Verify if a component can be deleted or not
//	Element 	Can be deleted 		Rule
//	B0			False				TCRuleDeletion001
//	B1			False				TCRuleDeletion002
//	B10			False				TCRuleDeletion003
//	B11f		False				TCRuleDeletion004a
//	B11l		False				TCRuleDeletion004b
//	B12			False				TCRuleDeletion005
//	B10*x* 		False				TCRuleDeletion006
//	B10*x 		False				TCRuleDeletion007
//	B10x*		False				TCRuleDeletion008
//	B11fx		False				TCRuleDeletion009a
//	B11lx		False				TCRuleDeletion009b
//	B12x		False				TCRuleDeletion010
//	TI			True				TCRuleDeletion011
//	Tix			False				TCRuleDeletion012
//	TIC(X)		True				TCRuleDeletion013
//	TICx(X)		False				TCRuleDeletion014
func VerifyIfComponentCanBeDeleted(uuid string, testCaseModelMap *map[string]fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage) (err error) {

	return err
}

// Verify the complex rules if a component can be deleted or not
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedSimpleRules(componentType fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum) (canBeDeleted bool, err error) {

	// Check simple rules of component can be deleted or not
	switch componentType {

	//	B0 - False - TCRuleDeletion001
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND:
		canBeDeleted = false

		//	B1 - False - TCRuleDeletion002
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	B10 - False - TCRuleDeletion003
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND:
		canBeDeleted = false

		//	B11			False				TCRuleDeletion004
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND:
		canBeDeleted = false

		//	B12			False				TCRuleDeletion005
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND:
		canBeDeleted = false

		//	B10*x* 		False				TCRuleDeletion006
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND:
		canBeDeleted = false

		//	B10*x 		False				TCRuleDeletion007
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND:
		canBeDeleted = false

		//	B10x*		False				TCRuleDeletion008
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND:

		//	B11x		False				TCRuleDeletion009
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
		fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	B12x		False				TCRuleDeletion010
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE:
		canBeDeleted = false

		//	TI			True				TCRuleDeletion011
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION:
		canBeDeleted = true

		//	Tix			False				TCRuleDeletion012
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIx_TESTINSTRUCTION_NONE_REMOVABLE:
		canBeDeleted = false

		//	TIC(X)		True				TCRuleDeletion013
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TIC_TESTINSTRUCTIONCONTAINER:
		canBeDeleted = true

		//	TICx(X)		False				TCRuleDeletion014
	case fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TICx_TESTINSTRUCTIONCONTAINER_NONE_REMOVABLE:
		canBeDeleted = false

	default:
		canBeDeleted = false
		err = errors.New(componentType.String() + " is an unknown componentType")

		commandAndRuleEngineObject.logger.WithFields(logrus.Fields{
			"id":            "3be29c9d-1db6-47dc-8ee1-8dcdfecda074",
			"componentType": componentType,
		}).Error(componentType.String() + " is an unknown componentType")

	}

	return canBeDeleted, err
}

// Verify the complex rules if a component can be deleted or not
// Rules how deletion of an element is done
// X = any allowed structure
// What to remove			Remove in structure				Result after deletion		Rule
// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
// n=TI or TIC(X)			B11f-n-B11l						B10							TCRuleDeletion102
// n=TI or TIC(X)			B11fx-n-B11lx					B10*x*						TCRuleDeletion103
// n=TI or TIC(X)			B11f-n-B11lx					B10x*						TCRuleDeletion104
// n=TI or TIC(X)			B11fx-n-B11l					B10*x						TCRuleDeletion105
// n=TI or TIC(X)			B11f-n-B12-X					B11f-X						TCRuleDeletion106
// n=TI or TIC(X)			B11fx-n-B12x-X					B11fx-X						TCRuleDeletion107
// n=TI or TIC(X)			B11f-n-B12x-X					B11fx-X						TCRuleDeletion108
// n=TI or TIC(X)			B11fx-n-B12-X					B11fx-X						TCRuleDeletion109
// n=TI or TIC(X)			X-B12-n-B11l					X-B11l						TCRuleDeletion110
// n=TI or TIC(X)			X-B12x-n-B11lx					X-B11lx						TCRuleDeletion111
// n=TI or TIC(X)			X-B12-n-B11lx					X-B11lx						TCRuleDeletion112
// n=TI or TIC(X)			X-B12x-n-B11l					X-B11lx						TCRuleDeletion113
// n=TI or TIC(X)			X-B12-n-B12-X					X-B12-X						TCRuleDeletion114
// n=TI or TIC(X)			X-B12x-n-B12x-X					X-B12x-X					TCRuleDeletion115
// n=TI or TIC(X)			X-B12-n-B12x-X					X-B12x-X					TCRuleDeletion116
// n=TI or TIC(X)			X-B12x-n-B12-X					X-B12x-X					TCRuleDeletion117
func (commandAndRuleEngineObject *commandAndRuleEngineObjectStruct) verifyIfComponentCanBeDeletedComplexRules(uuid string, testCaseModelMap map[string]fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage) (err error) {

	var ruleName string
	var ruleProcessed bool

	// Extract data for Previous Element
	currentElementUuid := uuid
	currentElementType := testCaseModelMap[currentElementUuid].TestCaseModelElementType

	// Extract data for Previous Element
	previousElementUuid := testCaseModelMap[uuid].PreviousElementUuid
	previousElementType := testCaseModelMap[previousElementUuid].TestCaseModelElementType

	// Extract data for Next Element
	nextElementUuid := testCaseModelMap[uuid].NextElementUuid
	nextlementType := testCaseModelMap[nextElementUuid].TestCaseModelElementType

	// Extract dta for Parent Element
	parrentElementUuid := testCaseModelMap[uuid].ParentElementUuid
	//parrentElemenType := testCaseModelMap[parrentElementUuid].TestCaseModelElementType

	// What to remove			Remove in structure				Result after deletion		Rule
	// n= TIC(X)				B1-n-B1							B0							TCRuleDeletion101
	if previousElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE &&
		currentElementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_TI_TESTINSTRUCTION &&
		nextlementType == fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1_BOND_NONE_SWAPPABLE {

		// Rule OK
		ruleName = "TCRuleDeletion101"
		ruleProcessed = true

		/*
			// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
			message TestCaseModelElementMessage {
			  string OriginalElementUuid = 1; // The original elements UUID, e.g. a TestInstruction unique UUID set by client system
			  string OriginalElementName = 2; // The original elements Name, e.g. a TestInstruction Name set by client system
			  string MatureElementUuid = 3; // The UUID that is created in the TestCase to give it a unique id
			  string PreviousElementUuid = 4;  // The UUID of the previous element. When there are no previous element then this field is populated with current element UUID
			  string NextElementUuid = 5;  // The UUID of the previous element. When there are no next element then this field is populated with current element UUID
			  string FirstChildElementUuid = 6; // The UUID of the first child element. Only applicable when this is a TestInstructionContainer. When there are no child element then this field is populated with current element UUID
			  string ParentElementUuid = 7; // The UUID of the parent, TestInstructionContainer. Only applicable when this is the last element inside a TestInstructionContainer. When there are no parent element then this field is populated with current element UUID
			  TestCaseModelElementTypeEnum TestCaseModelElementType = 8; // The specific type of this TestCase-element
			  string CurrentElementModelElement = 9; // The UUID of the element that this data act on, e.g. For TI & TIC the it is the same as 'OriginalElementUuid' but for BONDs then it is the BONDs UUID
			}
		*/

		newB0Element := fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementMessage{
			OriginalElementUuid:        commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
			OriginalElementName:        commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
			MatureElementUuid:          uuidGenerator.New().String(),
			PreviousElementUuid:        commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
			NextElementUuid:            commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
			FirstChildElementUuid:      commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
			ParentElementUuid:          parrentElementUuid,
			TestCaseModelElementType:   fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
			CurrentElementModelElement: commandAndRuleEngineObject.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		}

		// Add New Elements to Map
		testCaseModelMap[newB0Element.MatureElementUuid] = newB0Element

		// Remove Old Elements from Map
		delete(testCaseModelMap, previousElementUuid)
		delete(testCaseModelMap, currentElementUuid)
		delete(testCaseModelMap, nextElementUuid)

		// Update Reference in parent TIC
		tempParrentElement := testCaseModelMap[parrentElementUuid]
		tempParrentElement.FirstChildElementUuid = newB0Element.MatureElementUuid
		testCaseModelMap[parrentElementUuid] = tempParrentElement

	}

	return err

}
