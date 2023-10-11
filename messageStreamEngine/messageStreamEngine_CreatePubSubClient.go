package messageStreamEngine

import (
	"FenixTesterGui/common_code"
	"cloud.google.com/go/pubsub"
	"context"
	"crypto/tls"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func creatNewPubSubClient(ctx context.Context) (pubSubClient *pubsub.Client, err error) {

	// Check that some type of initialization has been done
	if len(gcpProject) == 0 {
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":         "4353c13c-1132-4300-b316-338a815ecb50",
			"gcpProject": gcpProject,
		}).Error("The variable 'gcpProject' is not initialized")

		return nil, errors.New("the variable 'gcpProject' is not initialized")
	}

	projectID := gcpProject

	var opts []grpc.DialOption

	// PubSub is handled within GCP so add TLS
	var creds credentials.TransportCredentials
	creds = credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	opts = []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	pubSubClient, err = pubsub.NewClient(ctx, projectID, option.WithGRPCDialOption(opts[0]))

	if err != nil {

		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":  "4488a860-a059-44ec-b7cf-870ce8f6b8a2",
			"err": err,
		}).Error("Got some problem when creating 'pubsub.NewClient'")

		return nil, err
	}

	return pubSubClient, err
}
