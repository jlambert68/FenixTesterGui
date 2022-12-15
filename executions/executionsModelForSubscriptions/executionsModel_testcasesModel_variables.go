package executionsModelForSubscriptions

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

// TestCaseExecutionMapKeyType
// Type used to define that this is TestCaseExecutionKey for model-maps
type TestCaseExecutionMapKeyType string // Should consist of 'TestCaseExecutionUuid' + 'TestCaseExecutionVersion'

// BLOCK START
// The block below  is used when checking if  a TestCaseExecution exists in any of the tables (OnQueue, UnderExecution, FinishedExecutions)

// SubscriptionTableType
// The type for subscription tables for TestCaseExecutions
type SubscriptionTableType uint8

// Subscription tables for TestCaseExecutions
const (
	SubscriptionTableForTestCaseExecutionOnQueueTable SubscriptionTableType = iota
	SubscriptionTableForTestCaseExecutionUnderExecutionTable
	SubscriptionTableForTestCaseExecutionFinishedExecutionsTable
)

// SubscriptionsForTestCaseExecutionMapOverallType
// Map holding all information about all 'TestCaseExecutionMapKey'
type SubscriptionsForTestCaseExecutionMapOverallType map[TestCaseExecutionMapKeyType]SubscriptionsForTestCaseExecutionMapDetailsType

// SubscriptionsForTestCaseExecutionMapDetailsType
// Map holding all information about one 'TestCaseExecutionMapKey' if the TestCaseExecution should exist in specific table (OnQueue, UnderExecution, FinishedExecutions)
type SubscriptionsForTestCaseExecutionMapDetailsType map[SubscriptionTableType]SubscriptionsForTestCaseExecutionMapDetailsStruct

// SubscriptionsForTestCaseExecutionMapDetailsStruct
// Hold information about if a TestCaseExecution should exist in specific table (OnQueue, UnderExecution, FinishedExecutions)
type SubscriptionsForTestCaseExecutionMapDetailsStruct struct {
	ShouldExistInTable bool
}

// BLOCK END
