package isolated

import (
	"fmt"
	"regexp"

	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = FDescribe("create-shared-domain command", func() {
	Context("Help", func() {
		It("displays the help information", func() {
			session := helpers.CF("create-shared-domain", "--help")
			Eventually(session).Should(Say("NAME:\n"))
			Eventually(session).Should(Say(regexp.QuoteMeta("create-shared-domain - Create a domain that can be used by all orgs (admin-only)")))
			Eventually(session).Should(Say("USAGE:\n"))
			Eventually(session).Should(Say(regexp.QuoteMeta("cf create-shared-domain DOMAIN [--router-group ROUTER_GROUP]")))
			Eventually(session).Should(Say("OPTIONS:\n"))
			Eventually(session).Should(Say("--router-group\\s+Routes for this domain will be configured only on the specified router group"))
			Eventually(session).Should(Say("SEE ALSO:\n"))
			Eventually(session).Should(Say("create-domain, domains, router-groups"))
		})
	})

	When("user is logged in as admin", func() {

	})

	When("user is not logged in as admin", func() {
		var (
			username string
			password string
		)

		BeforeEach(func() {
			username, password = helpers.CreateUser()
			helpers.LoginAs(username, password)
		})

		It("should not be able to create shared domain", func() {
			session := helpers.CF("create-shared-domain", "some-domain-name")
			Eventually(session).Should(Say(fmt.Sprintf("Creating shared domain some-domain-name as %s...", username)))
			Eventually(session).Should(Say("Server error, status code: 403, error code: 10003, message: You are not authorized to perform the requested action"))
			Eventually(session).Should(Say("FAILED"))
			Eventually(session).Should(Exit(1))
		})
	})
})
