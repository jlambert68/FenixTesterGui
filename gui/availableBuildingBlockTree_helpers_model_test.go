package gui

import (
	common_config "FenixTesterGui/common_code"
	"FenixTesterGui/grpc_out"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	fenixGuiTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"testing"
	"time"

	"FenixTesterGui/gui/UnitTestTestData"
	"github.com/stretchr/testify/assert"
)

/*
func testlist() *notelist {
	a := test.NewApp()
	n := &notelist{pref: a.Preferences()}

	return n
}

*/

func initLoggerForTest(filename string) (myTestLogger *logrus.Logger) {
	myTestLogger = logrus.StandardLogger()

	switch common_config.LoggingLevel {

	case logrus.DebugLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.InfoLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.WarnLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	default:
		log.Println("Not correct value for debugging-level, this was used: ", common_config.LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(common_config.LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		myTestLogger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			myTestLogger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

	return myTestLogger
}

var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

func TestLoadModelWithAvailableBuildingBlocksRegardingTestInstructions(t *testing.T) {

	var testInstructionsAndTestContainersMessage fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage
	//testInstructionsAndTestContainersMessage = &fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage

	//res := fenixGuiTestCaseBuilderServerGrpcApi.TestInstructionsAndTestContainersMessage{}
	//json.Unmarshal(byt, &res)
	mystring := UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001

	if err := jsonpb.UnmarshalString(mystring, &testInstructionsAndTestContainersMessage); err != nil {
		panic(err)
	}

	myLogger := initLoggerForTest("")

	var availableBuildingBlocksModel *availableBuildingBlocksModelStruct

	availableBuildingBlocksModel = &availableBuildingBlocksModelStruct{
		logger:                             myLogger,
		fenixGuiBuilderServerAddressToDial: "",
		fullDomainTestInstructionTypeTestInstructionRelationsMap:                   nil,
		fullDomainTestInstructionContainerTypeTestInstructionContainerRelationsMap: nil,
		availableBuildingBlocksForUITreeNodes:                                      nil,
		grpcOut:                                                                    grpc_out.GRPCOutStruct{},
		availableBuildingBlockModelSuitedForFyneTreeView:                           nil,
	}

	availableBuildingBlocksModel.loadModelWithAvailableBuildingBlocksRegardingTestInstructions(&testInstructionsAndTestContainersMessage)
	fmt.Println(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)

	b, err := json.Marshal(&availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	assert.Equal(t, fmt.Sprint(availableBuildingBlocksModel.availableBuildingBlocksForUITreeNodes), UnitTestTestData.TestInstructionsAndTestInstructionsRespons_ABB001_ExpectedResult)
}
