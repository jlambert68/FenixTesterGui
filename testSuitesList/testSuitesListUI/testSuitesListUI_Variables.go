package testSuitesListUI

import (
	"FenixTesterGui/testSuites/testSuitesModel"
)

var testSuitesModelPtr *TestSuitesModelStruct

// Holding all information about TestSuites
type TestSuitesModelStruct struct {
	// Map keeping track of all separate TestSuiteUiModels
	//x§TestSuitesUiMaplPtr *map[string]*testSuiteUI.TestSuiteUiModelStruct // Map-key = 'TestSuiteUuid'

	// Map keeping track of all separate TestSuiteModels
	TestSuitesMapPtr *map[string]*testSuitesModel.TestSuiteModelStruct // Map-key = 'TestSuiteUuid'

}
