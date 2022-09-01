package main

import (
	"FenixTesterGui/common_code"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func (fenixGuiBuilderProxyServerObject *fenixGuiBuilderProxyServerObjectStruct) InitLogger(filename string) {
	fenixGuiBuilderProxyServerObject.logger = logrus.StandardLogger()

	switch sharedCode.LoggingLevel {

	case logrus.DebugLevel:
		log.Println("'common_config.LoggingLevel': ", sharedCode.LoggingLevel)

	case logrus.InfoLevel:
		log.Println("'common_config.LoggingLevel': ", sharedCode.LoggingLevel)

	case logrus.WarnLevel:
		log.Println("'common_config.LoggingLevel': ", sharedCode.LoggingLevel)

	default:
		log.Println("Not correct value for debugging-level, this was used: ", sharedCode.LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(sharedCode.LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		fenixGuiBuilderProxyServerObject.logger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			fenixGuiBuilderProxyServerObject.logger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

}
