package cnc_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/celestiaorg/go-cnc"
)

type IntegrationTestSuite struct {
	suite.Suite

	dockerCompose *testcontainers.LocalDockerCompose
}

func (i *IntegrationTestSuite) SetupSuite() {
	composeFilePaths := []string{"docker/test-docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	i.dockerCompose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	i.dockerCompose.WaitForService("bridge0",
		wait.ForHTTP("/balance").WithPort("26659").
			WithStartupTimeout(60*time.Second).
			WithPollInterval(3*time.Second))
	execError := i.dockerCompose.WithCommand([]string{"up", "-d"}).Invoke()
	err := execError.Error
	if err != nil {
		i.Fail("failed to execute docker compose up:", "error: %v\nstdout: %v\nstderr: %v", err, execError.Stdout, execError.Stderr)
	}
}

func (i *IntegrationTestSuite) TearDownSuite() {
	execError := i.dockerCompose.Down()
	if err := execError.Error; err != nil {
		i.Fail("failed to execute docker compose down", "error: %v\nstdout: %v\nstderr: %v", err, execError.Stdout, execError.Stderr)
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (i *IntegrationTestSuite) TestNewClient() {
	cases := []struct {
		name          string
		options       []cnc.Option
		expectedError error
	}{
		{"without options", nil, nil},
		{"with timeout", []cnc.Option{cnc.WithTimeout(1 * time.Second)}, nil},
	}

	for _, c := range cases {
		i.Run(c.name, func() {
			client, err := cnc.NewClient("", c.options...)
			i.ErrorIs(err, c.expectedError)
			if c.expectedError != nil {
				i.Nil(client)
			} else {
				i.NotNil(client)
			}
		})
	}
}

func (i *IntegrationTestSuite) TestDataRoundTrip() {
	client, err := cnc.NewClient("http://localhost:26659", cnc.WithTimeout(30*time.Second))
	i.NoError(err)
	i.NotNil(client)

	randomData := []byte("random data")
	txRes, err := client.SubmitPFD(context.TODO(), [8]byte{1, 2, 3, 4, 5, 6, 7, 8}, randomData, 10000, 100000)
	i.Require().NoError(err)
	i.Require().NotNil(txRes)
	i.Assert().Zero(txRes.Code)
	expectedHeight := txRes.Height

	data, err := client.NamespacedData(context.TODO(), [8]byte{1, 2, 3, 4, 5, 6, 7, 8}, uint64(expectedHeight))
	i.Require().NoError(err)
	i.Require().NotNil(data)
	i.Len(data, 1)
	i.Contains(data, randomData)
}
