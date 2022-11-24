package testCaseSubscriptionHandler

import (
	"errors"
	"fmt"
)

// AddTestCaseExecutionStatusSubscription - Add a TestCaseExecutionStatusSubscription
func (testCaseExecutionStatusSubscriptionHandlerObject *TestCaseExecutionStatusSubscriptionHandlerStruct) AddTestCaseExecutionStatusSubscription(testCaseExecutionStatusSubscription *TestCaseExecutionStatusSubscriptionStruct) (err error) {

	var existsInMap bool

	// Create MapKey
	var testCaseExecutionStatusSubscriptionMapKey TestCaseExecutionStatusSubscriptionMapKeyType
	testCaseExecutionStatusSubscriptionMapKey =
		TestCaseExecutionStatusSubscriptionMapKeyType(
			testCaseExecutionStatusSubscription.TestCaseExecutionUuid +
				testCaseExecutionStatusSubscription.TestCaseExecutionVersion)

	// Check if this Subscription already exist in Map
	_, existsInMap = TestCaseExecutionExecutionStatusSubscriptionMap[testCaseExecutionStatusSubscriptionMapKey]
	if existsInMap == true {

		errorId := "9eb0c19e-fe5e-4d4b-921b-a37d1f3408a3"
		err = errors.New(fmt.Sprintf("testCaseExecutionSubscription, '%s', already exits in 'TestCaseExecutionExecutionStatusSubscriptionMap'. [ErrorID: %s]", testCaseExecutionStatusSubscriptionMapKey, errorId))

		return err
	}

	// Add Subscription to Map
	TestCaseExecutionExecutionStatusSubscriptionMap[testCaseExecutionStatusSubscriptionMapKey] = testCaseExecutionStatusSubscription

	return err

}

// RemoveTestCaseExecutionStatusSubscription - Remove a TestCaseExecutionStatusSubscription
func (testCaseExecutionStatusSubscriptionHandlerObject *TestCaseExecutionStatusSubscriptionHandlerStruct) RemoveTestCaseExecutionStatusSubscription(testCaseExecutionStatusSubscriptionMapKey TestCaseExecutionStatusSubscriptionMapKeyType) (err error) {

	var existsInMap bool

	// Check if this subscription exist in Map
	_, existsInMap = TestCaseExecutionExecutionStatusSubscriptionMap[testCaseExecutionStatusSubscriptionMapKey]
	if existsInMap == false {

		errorId := "9d48e50d-ca76-4e9e-9646-39738660697f"
		err = errors.New(fmt.Sprintf("testCaseExecutionSubscription, '%s', doesn't exit in 'TestCaseExecutionExecutionStatusSubscriptionMap'. [ErrorID: %s]", testCaseExecutionStatusSubscriptionMapKey, errorId))

		return err
	}

	// Remove Subscription from Map
	delete(TestCaseExecutionExecutionStatusSubscriptionMap, testCaseExecutionStatusSubscriptionMapKey)

	return err

}
