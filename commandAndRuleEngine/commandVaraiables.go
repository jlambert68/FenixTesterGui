package commandAndRuleEngine

import (
	"FenixTesterGui/testCase/testCaseModel"
	"github.com/sirupsen/logrus"
)

type commandsEngineStruct struct {
	logger    *logrus.Logger
	testcases *testCaseModel.TestCaseModelsStruct
}
