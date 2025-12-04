package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	
	_ "git.ucloudadmin.com/taishan/ustack-api-test/test/e2e/vm" // Register VM tests
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UStack API E2E Suite")
}

