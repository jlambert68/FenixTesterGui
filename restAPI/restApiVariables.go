package restAPI

import "github.com/sirupsen/logrus"

type RestApiStruct struct {
	logger *logrus.Logger
}

var RestAPI RestApiStruct
