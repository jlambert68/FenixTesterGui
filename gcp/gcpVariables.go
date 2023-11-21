package gcp

import (
	"github.com/markbates/goth"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"sync"
)

type GcpObjectStruct struct {
	logger                                                 *logrus.Logger
	gcpAccessTokenForServiceAccounts                       *oauth2.Token
	gcpAccessTokenForAuthorizedAccounts                    goth.User
	mutexWhenGeneratingGcpAccessTokenForAuthorizedAccounts *sync.Mutex

	gcpAccessTokenForAuthorizedAccountsPubSub goth.User
	refreshTokenResponse                      *RefreshTokenResponse
}

var GcpObject GcpObjectStruct

const (
	TargetServerGuiTestCaseBuilderServer TargetServerType = iota
	TargetServerGuiExecutionServer
	TargetServerGuiExecutionServerNew
)

type TargetServerType int
