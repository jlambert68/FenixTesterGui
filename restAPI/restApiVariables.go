package restAPI

import (
	"FenixTesterGui/fenixTestGuiObject"
	"github.com/sirupsen/logrus"
)

type RestApiStruct struct {
	Logger                        *logrus.Logger
	FenixTesterGuiObjectReference *fenixTestGuiObject.FenixGuiBuilderProxyServerObjectStruct
}

//var RestAPI RestApiStruct
