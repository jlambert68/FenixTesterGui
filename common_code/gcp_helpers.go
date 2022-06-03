package common_config

import (
	"FenixTesterGui/notToBeSentToGithub"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/idtoken"
	grpcMetadata "google.golang.org/grpc/metadata"
	"time"
)

// Generate Google access token. Used when running in GCP
func (fenixGuiBuilderProxyServerObject *fenixGuiBuilderProxyServerObjectStruct) generateGCPAccessToken(ctx context.Context) (appendedCtx context.Context, returnAckNack bool, returnMessage string) {

	// Only create the token if there is none, or it has expired
	if fenixGuiBuilderProxyServerObject.gcpAccessToken == nil || fenixGuiBuilderProxyServerObject.gcpAccessToken.Expiry.Before(time.Now()) {

		// Create an identity token.
		// With a global TokenSource tokens would be reused and auto-refreshed at need.
		// A given TokenSource is specific to the audience.
		/*
			tokenSource, err := idtoken.NewTokenSource(ctx, "https://"+common_config.FenixGuiBuilderServerAddress)
			if err != nil {
				fenixGuiBuilderProxyServerObject.logger.WithFields(logrus.Fields{
					"ID":  "8ba622d8-b4cd-46c7-9f81-d9ade2568eca",
					"err": err,
				}).Error("Couldn't generate access token")

				return nil, false, "Couldn't generate access token"
			}

			token, err := tokenSource.Token()
		*/
		/*
			var eMailAndPrivateKey = struct {
				Email      string `json:"client_email"`
				PrivateKey string `json:"private_key"`
			}{}
			json.Unmarshal(serviceAccountKeyJson, &eMailAndPrivateKey)
			config := &jwt.Config{
				Email:      eMailAndPrivateKey.Email,
				PrivateKey: []byte(eMailAndPrivateKey.PrivateKey),
				Scopes: []string{
					gcp_scope,
				},
				TokenURL:   google.JWTTokenURL,
				UseIDToken: false,
			}

		*/

		tokenSource, err := idtoken.NewTokenSource(ctx, notToBeSentToGithub.gcp_scope, idtoken.WithCredentialsJSON(notToBeSentToGithub.serviceAccountKeyJson))

		if err != nil {
			fenixGuiBuilderProxyServerObject.logger.WithFields(logrus.Fields{
				"ID":  "8ba622d8-b4cd-46c7-9f81-d9ade2568eca",
				"err": err,
			}).Error("Couldn't generate access token")

			return nil, false, "Couldn't generate access token"
		}

		token, err := tokenSource.Token()
		//token, err := config.TokenSource(oauth2.NoContext).Token()

		if err != nil {
			fenixGuiBuilderProxyServerObject.logger.WithFields(logrus.Fields{
				"ID":  "0cf31da5-9e6b-41bc-96f1-6b78fb446194",
				"err": err,
			}).Error("Problem getting the token")

			return nil, false, "Problem getting the token"
		} else {
			fenixGuiBuilderProxyServerObject.logger.WithFields(logrus.Fields{
				"ID":    "8b1ca089-0797-4ee6-bf9d-f9b06f606ae9",
				"token": "Nothing to see", //token,
			}).Debug("Got Bearer Token")
		}

		fenixGuiBuilderProxyServerObject.gcpAccessToken = token

	}

	fenixGuiBuilderProxyServerObject.logger.WithFields(logrus.Fields{
		"ID": "cd124ca3-87bb-431b-9e7f-e044c52b4960",
		"fenixGuiBuilderProxyServerObject.gcpAccessToken": "Nothing to see", //fenixGuiBuilderProxyServerObject.gcpAccessToken,
	}).Debug("Will use Bearer Token")

	// Add token to gRPC Request.
	appendedCtx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+fenixGuiBuilderProxyServerObject.gcpAccessToken.AccessToken)

	return appendedCtx, true, ""

}
