package executionsModel

// ExecutionsModelObjectStruct - struct to object that hold all parts to of TestCaseExecution-model together
type ExecutionsModelObjectStruct struct {
}

// ExecutionsModelObject -  Object that hold all parts to of TestCaseExecution-model together
var ExecutionsModelObject ExecutionsModelObjectStruct

// Type used to hold one Domain with UUID and Name
type domainsStruct struct {
	domainUuid string
	domainName string
}

// Type used to define that this is TestCaseExecutionKey for model-maps
type TestCaseExecutionMapKeyType string // Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'
