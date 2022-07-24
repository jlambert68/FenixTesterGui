package commandAndRuleEngine

import (
	uuidGenerator "github.com/google/uuid"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
)

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

/*
	// One TestCaseElement, that is not a TestInstructionContainer, that hold type, reference to previous, next and parent elements
	message MatureTestCaseModelElementMessage {
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

// Create a new B0-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB0Element() (newBondB0Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// Create new Bond element
	newBondB0Element = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        matureElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B0_BOND,
	}

	return newBondB0Element

}

// Create a new B1f-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB1fElement(parentElementUuid string) (newBondB1fElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB1fElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          "",
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1f_BOND_NONE_SWAPPABLE,
	}

	return newBondB1fElement

}

// Create a new B1l-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB1lElement(parentElementUuid string) (newBondB1lElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB1lElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      "",
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B1l_BOND_NONE_SWAPPABLE,
	}

	return newBondB1lElement

}

// Create a new B10-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB10Element(parentElementUuid string) (newBondB10Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB10Element = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10_BOND,
	}

	return newBondB10Element

}

// Create a new B10*x*-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB10oxoElement(parentElementUuid string) (newBondB10oxoElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB10oxoElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10oxo_BOND,
	}

	return newBondB10oxoElement

}

// Create a new B10x*-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB10xoElement(parentElementUuid string) (newBondB10xoElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB10xoElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10xo_BOND,
	}

	return newBondB10xoElement

}

// Create a new B10*x-bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB10oxElement(parentElementUuid string) (newBondB10oxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB10oxElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B10ox_BOND,
	}

	return newBondB10oxElement

}

// Create a new B11f-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB11fElement(parentElementUuid string) (newBondB11fElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB11fElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          "",
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11f_BOND,
	}

	return newBondB11fElement

}

// Create a new B11l-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB11lElement(parentElementUuid string) (newBondB11lElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB11lElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      "",
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11l_BOND,
	}

	return newBondB11lElement

}

// Create a new B11fx-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB11fxElement(parentElementUuid string) (newBondB11fxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB11fxElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      matureElementUuid,
		NextElementUuid:          "",
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11fx_BOND_NONE_SWAPPABLE,
	}

	return newBondB11fxElement

}

// Create a new B11lx-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB11lxElement(parentElementUuid string) (newBondB11lxElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB11lxElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      "",
		NextElementUuid:          matureElementUuid,
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B11lx_BOND_NONE_SWAPPABLE,
	}

	return newBondB11lxElement

}

// Create a new B12-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB12Element(parentElementUuid string) (newBondB12Element fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB12Element = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      "",
		NextElementUuid:          "",
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12_BOND,
	}

	return newBondB12Element

}

// Create a new B12x-Bond to be used in the TestCase-model
func (commandAndRuleEngine *commandAndRuleEngineObjectStruct) createNewBondB12xElement(parentElementUuid string) (newBondB12xElement fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage) {

	// Generate new UUID
	matureElementUuid := uuidGenerator.New().String()

	// If there is no parent element then use 'matureElementUuid'
	if parentElementUuid == "" {
		parentElementUuid = matureElementUuid
	}

	// Create new Bond element
	newBondB12xElement = fenixGuiTestCaseBuilderServerGrpcApi.MatureTestCaseModelElementMessage{
		OriginalElementUuid:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondUuid,
		OriginalElementName:      commandAndRuleEngine.availableBondsMap[fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE].BasicBondInformation.VisibleBondAttributes.BondName,
		MatureElementUuid:        matureElementUuid,
		PreviousElementUuid:      "",
		NextElementUuid:          "",
		FirstChildElementUuid:    matureElementUuid,
		ParentElementUuid:        parentElementUuid,
		TestCaseModelElementType: fenixGuiTestCaseBuilderServerGrpcApi.TestCaseModelElementTypeEnum_B12x_BOND_NONE_SWAPPABLE,
	}

	return newBondB12xElement

}
