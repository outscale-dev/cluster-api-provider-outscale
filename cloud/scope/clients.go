package scope

import (
	"context"
	"os"

	"errors"
	osc "github.com/outscale/osc-sdk-go/v2"
)

// OscClient contains input client to use outscale api
type OscClient struct {
	auth context.Context
	api  *osc.APIClient
}

// newOscClient return OscLient using secret credentials
func newOscClient() (*OscClient, error) {
	accessKey := os.Getenv("OSC_ACCESS_KEY")
	if accessKey == "" {
		return nil, errors.New("environment variable OSC_ACCESS_KEY is required")
	}
	secretKey := os.Getenv("OSC_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("environment variable OSC_SECRET_KEY is required")
	}
	oscClient := &OscClient{
		api: osc.NewAPIClient(osc.NewConfiguration()),
		auth: context.WithValue(context.Background(), osc.ContextAWSv4, osc.AWSv4{
			AccessKey: os.Getenv("OSC_ACCESS_KEY"),
			SecretKey: os.Getenv("OSC_SECRET_KEY"),
		}),
	}
	return oscClient, nil
}
