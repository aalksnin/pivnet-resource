package pivnet_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/pivotal-cf-experimental/pivnet-resource"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

const (
	apiPrefix = "/api/v2"
)

var _ = Describe("PivnetClient", func() {
	var (
		server *ghttp.Server
		client pivnet.Client
		token  string
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		apiAddress := server.URL() + apiPrefix
		token = "my-auth-token"
		client = pivnet.NewClient(apiAddress, token)
	})

	AfterEach(func() {
		server.Close()
	})

	It("has authenticated headers for each request", func() {
		response := fmt.Sprintf(`{"releases": [{"version": "1234"}]}`)

		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", apiPrefix+"/products/my-product-id/releases"),
				ghttp.VerifyHeaderKV("Authorization", fmt.Sprintf("Token: %s", token)),
				ghttp.RespondWith(http.StatusOK, response),
			),
		)

		_, err := client.ProductVersions("my-product-id")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Product Versions", func() {
		Context("when parsing the url fails with error", func() {
			It("forwards the error", func() {
				badURL := "%%%"
				client = pivnet.NewClient(badURL, token)

				_, err := client.ProductVersions("some product")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("%%%"))
			})
		})

		Context("when making the request fails with error", func() {
			It("forwards the error", func() {
				badURL := "https://not-a-real-url.com"
				client = pivnet.NewClient(badURL, token)

				_, err := client.ProductVersions("some product")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("no such host"))
			})
		})

		Context("when the json unmarshalling fails with error", func() {
			It("forwards the error", func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", apiPrefix+"/products/my-product-id/releases"),
						ghttp.RespondWith(http.StatusOK, "%%%"),
					),
				)

				_, err := client.ProductVersions("my-product-id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid character"))
			})
		})

		It("gets versions", func() {
			productVersion := "v" + strconv.Itoa(rand.Int())
			response := fmt.Sprintf(
				`{"releases": [{"version": %q}, {"version": %q}]}`,
				productVersion, productVersion,
			)

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", apiPrefix+"/products/my-product-id/releases"),
					ghttp.RespondWith(http.StatusOK, response),
				),
			)

			versions, err := client.ProductVersions("my-product-id")
			Expect(err).NotTo(HaveOccurred())
			Expect(server.ReceivedRequests()).Should(HaveLen(1))
			Expect(versions).To(HaveLen(2))
			Expect(versions[0]).Should(Equal(productVersion))
			Expect(versions[1]).Should(Equal(productVersion))
		})
	})
})
