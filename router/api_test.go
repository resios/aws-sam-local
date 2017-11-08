package router_test

import (
	"github.com/awslabs/aws-sam-local/router"
	"github.com/awslabs/goformation/cloudformation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getApiResourceFromTemplate(path string) (*router.AWSServerlessApi) {
	templateUri := &path
	apiResource := &router.AWSServerlessApi{
		AWSServerlessApi: &cloudformation.AWSServerlessApi{
			DefinitionUri: &cloudformation.AWSServerlessApi_StringOrS3Location{
				String: templateUri,
			},
		},
	}
	return apiResource
}

var _ = Describe("Api", func() {

	Context("Load local Swagger definitions", func() {
		/*apiResource := router.AWSServerlessApi{}
		templateUri := new(string)
		*templateUri = "test/templates/open-api/pet-store-simple.json"
		definitionUri := new(cloudformation.AWSServerlessApi_StringOrS3Location)
		definitionUri.String = templateUri

		apiResource.DefinitionUri = definitionUri*/
		It("Succesfully loads the basic template", func() {
			apiResource := getApiResourceFromTemplate("../test/templates/open-api/pet-store-simple.json")

			mounts, err := apiResource.Mounts()

			Expect(err).Should(BeNil())
			Expect(mounts).ShouldNot(BeNil())

			Expect(mounts).ShouldNot(BeEmpty())
			Expect(len(mounts)).Should(BeIdenticalTo(4))
		})

		It("Succesfully reads integration definition", func() {
			apiResource := getApiResourceFromTemplate("../test/templates/open-api/pet-store-simple.json")

			mounts, err := apiResource.Mounts()

			Expect(err).Should(BeNil())
			Expect(mounts).ShouldNot(BeNil())

			for _, mount := range mounts {
				if mount.Method == "get" && mount.Path == "/pets" {
					Expect(mount.IntegrationArn.Arn).Should(BeIdenticalTo("arn:aws:lambda:us-west-2:123456789012:function:Calc"))
				}
			}
		})

	})
})
