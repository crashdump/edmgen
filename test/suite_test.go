//go:build integration
// +build integration

package edm

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/crashdump/edmgen/pkg/edm"
)

var fixturesDir = "../test/fixtures"

type EdmTestSuite struct {
	suite.Suite
}

func (suite *EdmTestSuite) SetupSuite() {
	if _, err := os.Stat(fmt.Sprintf("%s/linux/.git", fixturesDir)); err == nil {
		// linux source code already present, we can continue
		return
	}

	suite.T().Log("downloading linux source, please be patient")
	cmd := exec.Command("git", "clone", "--depth=1", "https://github.com/torvalds/linux.git")
	cmd.Dir = fixturesDir
	_, err := cmd.Output()
	if err != nil {
		suite.FailNow("could not download the linux source code fixtures", err)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(EdmTestSuite))
}

func (suite *EdmTestSuite) TestWalk_Linux() {
	edmc, err := edm.New(edm.Opts{})
	suite.NoError(err)

	err = edmc.SelectFiles(fmt.Sprintf("%s/linux", fixturesDir))
	suite.NoError(err)
	suite.NotEmpty(edmc.Paths)

	err = edmc.ExamineFiles(paths)
	suite.NoError(err)
	suite.NotEmpty(edmc.Content)
}
