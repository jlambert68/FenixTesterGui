package testCaseSubscriptionHandler

// InitiateTestCaseExecutionStatusSubscriptionHandler - Initiate all variables needed by the TestCaseExecution-SubscriptionHandler
func (testCaseExecutionStatusSubscriptionHandlerObject *TestCaseExecutionStatusSubscriptionHandlerStruct) InitiateTestCaseExecutionStatusSubscriptionHandler() {

	// Initiate Subscription-Map
	TestCaseExecutionExecutionStatusSubscriptionMap = make(map[TestCaseExecutionStatusSubscriptionMapKeyType]*TestCaseExecutionStatusSubscriptionStruct)

}
