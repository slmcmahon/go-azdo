package tests

import (
	"os"
	"testing"

	"github.com/slmcmahon/go-azdo"
)

func TestGetRepoList(t *testing.T) {
	ops := azdo.NewAZDOOperations(os.Getenv("AZDO_PAT"), os.Getenv("AZDO_ORG"), os.Getenv("AZDO_PROJECT"))
	repoList, _ := ops.GetRepositories()
	if repoList.Count == 0 {
		t.Errorf("No repos were returned")
	}
}

func TestGetVariables(t *testing.T) {
	ops := azdo.NewAZDOOperations(os.Getenv("AZDO_PAT"), os.Getenv("AZDO_ORG"), os.Getenv("AZDO_PROJECT"))
	vars, _ := ops.GetVariableLibraries()
	if vars.Count == 0 {
		t.Errorf("No variable libraries were found.")
	}
}

func TestGetRepository(t *testing.T) {
	ops := azdo.NewAZDOOperations(os.Getenv("AZDO_PAT"), os.Getenv("AZDO_ORG"), os.Getenv("AZDO_PROJECT"))
	repo, _ := ops.GetRepository("competency-core")
	if repo.Name != "competency-core" {
		t.Errorf("Did not find the repository")
	}
}
