package executionsModelForExecutions

// Initiate the channels used when Adding or Removing items to/from OnQueue-table, UnderExecution-table or FinishedExecutions-table
func InitiateAndStartChannelsUsedByListModel() {

	// Initiate Channel used for Adding and Deleting Execution items in OnQueue-table
	OnQueueTableAddRemoveChannel = make(chan OnQueueTableAddRemoveChannelStruct, MaximumNumberOfItemsForOnQueueTableAddRemoveChannel)

	// Initiate Channel used for Adding and Deleting Execution items in UnderExecution-table
	UnderExecutionTableAddRemoveChannel = make(chan UnderExecutionTableAddRemoveChannelStruct, MaximumNumberOfItemsForUnderExecutionTableAddRemoveChannel)

	// Initiate Channel used for Adding and Deleting Execution items in FinishedExecutions-table
	FinishedExecutionsTableAddRemoveChannel = make(chan FinishedExecutionsTableAddRemoveChannelStruct, MaximumNumberOfItemsForFinishedExecutionsTableAddRemoveChannel)

}
