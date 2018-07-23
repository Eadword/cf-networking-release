package acceptance_test

import (
	"cf-pusher/cf_cli_adapter"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = PDescribe("external connectivity", func() {
	var (
		appA      string
		orgName   string
		spaceName string
		appRoute  string
		cli       *cf_cli_adapter.Adapter
	)

	BeforeEach(func() {
		if testConfig.Internetless {
			Skip("skipping egress policy tests")
		}

		cli = &cf_cli_adapter.Adapter{CfCliPath: "cf"}
		appA = fmt.Sprintf("appA-%d", rand.Int31())

		orgName = testConfig.Prefix + "egress-policy-org"
		spaceName = testConfig.Prefix + "space"
		setupOrgAndSpace(orgName, spaceName)

		By("unbinding all running ASGs")
		for _, sg := range testConfig.DefaultSecurityGroups {
			Expect(cf.Cf("unbind-running-security-group", sg).Wait(Timeout_Short)).To(gexec.Exit(0))
		}

		By("pushing the test app")
		pushProxy(appA)
		appRoute = fmt.Sprintf("http://%s.%s/", appA, config.AppsDomain)
	})

	AfterEach(func() {
		By("adding back all the original running ASGs")
		for _, sg := range testConfig.DefaultSecurityGroups {
			Expect(cf.Cf("bind-running-security-group", sg).Wait(Timeout_Short)).To(gexec.Exit(0))
		}

		By("deleting the test org")
		Expect(cf.Cf("delete-org", orgName, "-f").Wait(Timeout_Push)).To(gexec.Exit(0))
	})

	checkRequest := func(route string, expectedStatusCode int, expectedResponseSubstring string) error {
		resp, err := http.Get(route)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		respBody := string(respBytes)

		if resp.StatusCode != expectedStatusCode {
			return fmt.Errorf("test http get to %s: expected response code %d but got %d.  response body:\n%s", route, expectedStatusCode, resp.StatusCode, respBody)
		}
		if !strings.Contains(respBody, expectedResponseSubstring) {
			return fmt.Errorf("test http get to %s: expected response to contain %q but instead saw:\n%s", route, expectedResponseSubstring, respBody)
		}
		return nil
	}

	canProxy := func() error {
		return checkRequest(appRoute+"proxy/example.com", 200, "Example Domain")
	}
	cannotProxy := func() error {
		return checkRequest(appRoute+"proxy/example.com", 500, "example.com")
	}

	Describe("egress policy connectivity", func() {
		It("the app can reach the internet when egress policy is present", func(done Done) {
			By("checking that the app cannot reach the internet using http and dns")
			Eventually(cannotProxy, "10s", "1s").Should(Succeed())
			Consistently(cannotProxy, "2s", "0.5s").Should(Succeed())

			By("creating egress policy")
			appAGuid, err := cli.AppGuid(appA)
			Expect(err).NotTo(HaveOccurred())
			createEgressPolicy(cli, fmt.Sprintf(testEgressPolicies, appAGuid))

			By("checking that the app can use dns and http to reach the internet")
			Eventually(canProxy, "10s", "1s").Should(Succeed())
			Consistently(canProxy, "2s", "0.5s").Should(Succeed())

			close(done)
		}, 180 /* <-- overall spec timeout in seconds */)
	})
})

func createEgressPolicy(cli *cf_cli_adapter.Adapter, payload string) {
	payloadFile, err := ioutil.TempFile("", "")
	Expect(err).NotTo(HaveOccurred())

	_, err = payloadFile.Write([]byte(payload))
	Expect(err).NotTo(HaveOccurred())

	err = payloadFile.Close()
	Expect(err).NotTo(HaveOccurred())

	_, err = cli.Curl("POST", "/networking/v1/external/policies", payloadFile.Name())
	Expect(err).NotTo(HaveOccurred())
}

var testEgressPolicies = `
	"egress_policies": [
		{
			"source": {"id": %q},
			"destination": {
				"protocol": "tcp",
				"ips": [
					"start": "0.0.0.0",
					"end": "255.255.255.255",
				]
			}
		}
	]
`
