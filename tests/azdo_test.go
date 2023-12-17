package tests

import (
	"os"
	"testing"

	"github.com/slmcmahon/go-azdo"
)

func getOps() *azdo.AZDOOperations {
	return azdo.NewAZDOOperations(os.Getenv("AZDO_PAT"), os.Getenv("AZDO_ORG"), os.Getenv("AZDO_PROJECT"))
}

func TestGetRepoList(t *testing.T) {
	ops := getOps()
	repoList, _ := ops.GetRepositories()
	if len(repoList) == 0 {
		t.Errorf("No repos were returned")
	}
}

func TestGetVariables(t *testing.T) {
	ops := getOps()
	vars, _ := ops.GetVariableLibraries()
	if len(vars) == 0 {
		t.Errorf("No variable libraries were found.")
	}
}

func TestGetRepository(t *testing.T) {
	ops := getOps()
	repo, _ := ops.GetRepository("competency-core")
	if repo.Name != "competency-core" {
		t.Errorf("Did not find the repository")
	}
}

func TestGetCommits(t *testing.T) {
	ops := getOps()
	commits, _ := ops.GetCommits("competency-core")
	if len(commits) == 0 {
		t.Errorf("Did not find any commits")
	}
}

func TestGetCommit(t *testing.T) {
	ops := getOps()
	commit, _ := ops.GetCommit("competency-core", "7bac654c332acd0faf6fc03170656ae19370d556")
	if commit.Comment != "2023-12-15: Remove MI PCP rules" {
		t.Errorf("Did not find the commit")
	}
}

func TestGetChanges(t *testing.T) {
	ops := getOps()
	changes, _ := ops.GetChanges("competency-core", "7bac654c332acd0faf6fc03170656ae19370d556")
	if len(changes) == 0 {
		t.Errorf("Did not find the changes")
	}
}
