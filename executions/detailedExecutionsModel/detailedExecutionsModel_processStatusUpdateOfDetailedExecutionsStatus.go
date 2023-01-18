package detailedExecutionsModel

import fenixExecutionServerGuiGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"

// Updates specific status information based on subscriptions updates from GuiExecutionServer
func (detailedExecutionsModelObject *DetailedExecutionsModelObjectStruct) processStatusUpdateOfDetailedExecutionsStatus(
	testCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage *fenixExecutionServerGuiGrpcApi.TestCaseExecutionsStatusAndTestInstructionExecutionsStatusMessage) {

	/*

	   // TestCaseStatus
	   // Add to existing structure for the TestCase-Status on summary page
	   var testCaseExecutionsStatusForSummaryTableData TestCaseExecutionsStatusForSummaryTableStruct

	   // First get the latest data from DB-content
	   for testCaseExecutionDetailsCounter, testCaseExecutionDetailsMessage := range testCaseExecutionResponse.TestCaseExecutionDetails {

	   // When it's the first instance of status then use that as the base
	   if testCaseExecutionDetailsCounter == 0 {
	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   } else {
	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testCaseExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }
	   }

	   // Second get the latest data from status updates
	   for _, testCaseExecutionDetailsMessage := range testCaseExecutionsDetails.TestCaseExecutionsStatusUpdates {

	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testCaseExecutionDetailsMessage.TestCaseExecutionDetails.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testCaseExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testCaseExecutionsStatusForSummaryTableData = TestCaseExecutionsStatusForSummaryTableStruct{
	   TestCaseUIName:                 testCaseExecutionResponse.TestCaseExecutionBasicInformation.
	   TestCaseName,
	   TestCaseStatusValue:            uint32(testCaseExecutionDetailsMessage.TestCaseExecutionDetails.
	   TestCaseExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testCaseExecutionDetailsMessage.TestCaseExecutionDetails.
	   ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }

	   // Add the TestStatus for Summary page
	   testCaseExecutionsDetails.TestCaseExecutionsStatusForSummaryTable = testCaseExecutionsStatusForSummaryTableData

	   // TestInstructionsStatus

	   // First get the latest data from DB-content
	   for _, testInstructionExecutionDetailsMessage := range testCaseExecutionResponse.TestInstructionExecutions {
	   var testInstructionExecutionsStatusForSummaryTableData TestInstructionExecutionsStatusForSummaryTable
	   // Loop all status messages
	   for testInstructionExecutionInformationCounter, testInstructionExecutionInformation := range testInstructionExecutionDetailsMessage.
	   TestInstructionExecutionsInformation {

	   // When it's the first instance of status then use that as the base
	   if testInstructionExecutionInformationCounter == 0 {
	   testInstructionExecutionsStatusForSummaryTableData = TestInstructionExecutionsStatusForSummaryTable{
	   TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
	   TestInstructionStatusValue:     uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   } else {
	   // Check if the new timestamp > existing timestamp, if so then use new instance
	   if testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime().After(
	   testInstructionExecutionsStatusForSummaryTableData.ExecutionStatusUpdateTimeStamp) {

	   testInstructionExecutionsStatusForSummaryTableData = TestInstructionExecutionsStatusForSummaryTable{
	   TestInstructionExecutionUIName: testInstructionExecutionDetailsMessage.TestInstructionExecutionBasicInformation.TestInstructionName,
	   TestInstructionStatusValue: uint32(testInstructionExecutionInformation.TestInstructionExecutionStatus),
	   ExecutionStatusUpdateTimeStamp: testInstructionExecutionInformation.ExecutionStatusUpdateTimeStamp.AsTime(),
	   }
	   }
	   }
	   }

	   // Append the TestInstructionsStatus for Summary page
	   testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable = append(
	   testCaseExecutionsDetails.TestInstructionExecutionsStatusForSummaryTable,
	   testInstructionExecutionsStatusForSummaryTableData)
	   }

	   hmm hur matchar man dessa två slice:ar?

	   }

	   }
	*/

}
