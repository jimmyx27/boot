package client_test

import (
	"boot/pkg/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var testT *testing.T

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	testT = t
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Client Suite", func() {
	var mockHttpClient *mocks.HttpClient
	var mockLogger *mocks.Logger

	BeforeEach(func() {
		mockLogger = mocks.NewLogger(TestT)
		mockHttpClient = mocks.NewHttpClient(TestT)
	})

	AfterEach(func() {
		url := "boot.io"
		amount := 5
		mockLogger = nil
		mockHttpClient = nil
	})

	var _ = Describe("Getting data", func() {
		It("Should give data", func() {
			mockLogger.On("Debug", mock.Anything)

			mockHttpClient.On("GetData", mock.Anything, mock.Anything).Return(`got data`, nil)

			cl, err := client.NewClient(mockLogger, mockHttpClient)
			Expect(err).NotTo(HaveOccured())

			res, err := cl.GetData(url, amount)
			Expect(err).NotTo(HaveOccured())
			Expect(res).To(Equal("Got data"))
		})
	})
})
