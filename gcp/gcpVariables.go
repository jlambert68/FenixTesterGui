package gcp

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type gcpObjectStruct struct {
	logger         *logrus.Logger
	gcpAccessToken *oauth2.Token
}

var Gcp gcpObjectStruct
