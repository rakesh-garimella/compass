package logging

import (
	"github.com/kyma-incubator/compass/tests/e2e/provisioning/pkg/client/v1_client"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type Client struct {
	logConfig Config
	expectedConfigMap string
}

func NewLoggingClient(config Config) *Client{
	return &Client{
		logConfig: config,
	}
}

func (s *Client)  setKubeConfig(k8sConfig *string, log logrus.FieldLogger) {
	//k8sConfig, err := config.GetConfig()
	//if err != nil {
	//	panic(err)
	//}

	cli, err := client.New(k8sConfig, client.Options{})
	if err != nil {
		panic(err)
	}
	cmClient := v1_client.NewConfigMapClient(log logrus.FieldLogger)
}

