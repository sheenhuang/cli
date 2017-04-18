package plugin_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/command/commandfakes"
	. "code.cloudfoundry.org/cli/command/plugin"
	"code.cloudfoundry.org/cli/command/plugin/pluginfakes"
	"code.cloudfoundry.org/cli/command/plugin/shared"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("uninstall-plugin command", func() {
	var (
		cmd        UninstallPluginCommand
		testUI     *ui.UI
		fakeConfig *commandfakes.FakeConfig
		fakeActor  *pluginfakes.FakeUninstallPluginActor
		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		plugins := map[string]configv3.Plugin{
			"some-plugin": configv3.Plugin{
				Version: configv3.PluginVersion{
					Major: 1,
					Minor: 2,
					Build: 3,
				},
			},
		}
		fakeConfig.PluginsReturns(plugins)
		fakeActor = new(pluginfakes.FakeUninstallPluginActor)

		cmd = UninstallPluginCommand{
			UI:     testUI,
			Config: fakeConfig,
			Actor:  fakeActor,
		}

		cmd.RequiredArgs.PluginName = "some-plugin"
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when the plugin is installed", func() {
		BeforeEach(func() {
			fakeActor.UninstallPluginReturns(nil)
		})

		Context("when no errors are encountered", func() {
			It("uninstalls the plugin and outputs the success message", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(testUI.Out).To(Say("Uninstalling plugin some-plugin..."))
				Expect(testUI.Out).To(Say("OK"))
				Expect(testUI.Out).To(Say("Plugin some-plugin 1.2.3 successfully uninstalled."))

				Expect(fakeActor.UninstallPluginCallCount()).To(Equal(1))
				_, pluginName := fakeActor.UninstallPluginArgsForCall(0)
				Expect(pluginName).To(Equal("some-plugin"))
			})
		})

		Context("when uninstalling the plugin encounters an error", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some error")
				fakeActor.UninstallPluginReturns(expectedErr)
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError(expectedErr))
			})
		})
	})

	Context("when the plugin is not installed", func() {
		BeforeEach(func() {
			fakeActor.UninstallPluginReturns(
				pluginaction.PluginNotFoundError{
					Name: "some-plugin",
				},
			)
		})

		It("returns a PluginNotFoundError", func() {
			Expect(testUI.Out).To(Say("Uninstalling plugin some-plugin..."))
			Expect(executeErr).To(MatchError(shared.PluginNotFoundError{
				Name: "some-plugin",
			}))
		})
	})
})
