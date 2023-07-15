package login_test

import (
	"testing"

	"jennifer/dealls-tech-test/internal/server/config"

	"github.com/orlangure/gnomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLogin(t *testing.T) {
	var container *gnomock.Container

	config.SetupSuiteTest(container)

	RegisterFailHandler(Fail)
	RunSpecs(t, "login Suite")
}
