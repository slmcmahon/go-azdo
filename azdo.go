package azdo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func NewAZDOOperations(pat string, org string, project string) *AZDOOperations {
	return &AZDOOperations{
		PAT:          pat,
		Organization: org,
		Project:      project,
	}
}

// CompareAndPrintDifference compares two variable groups and prints variables that are in the first group but not in the second.
func (ops *AZDOOperations) CompareAndPrintDifference(group1Name string, group1Vars map[string]struct{}, group2Name string, group2Vars map[string]struct{}) {
	var hasDifference bool

	fmt.Printf("\nVariables in %s but not in %s:\n", group1Name, group2Name)
	for name := range group1Vars {
		if _, exists := group2Vars[name]; !exists {
			fmt.Println(" - " + name)
			hasDifference = true
		}
	}

	if !hasDifference {
		fmt.Println(" - No differences found")
	}
}

// azureDevOpsGetRequest makes a GET request to the Azure DevOps API and unmarshals the response.
func azureDevOpsGetRequest(pat, url string, response AzureDevOpsAPIResponse) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth("", pat)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}

func (ops *AZDOOperations) GetRepositories() ([]Repository, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories", ops.Organization, ops.Project)
	var response AZDOResponse[Repository]
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response.Value, err
}

func (ops *AZDOOperations) GetRepository(repository string) (Repository, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s", ops.Organization, ops.Project, repository)
	var response Repository
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response, err
}

func (ops *AZDOOperations) GetCommits(repository string) ([]Commit, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s/commits", ops.Organization, ops.Project, repository)
	var response AZDOResponse[Commit]
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response.Value, err
}

func (ops *AZDOOperations) GetCommit(repository string, commitId string) (Commit, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s/commits/%s", ops.Organization, ops.Project, repository, commitId)
	var response Commit
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response, err
}

func (ops *AZDOOperations) GetChanges(repository string, commit string) ([]Change, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s/commits/%s/changes", ops.Organization, ops.Project, repository, commit)
	var response ChangeListResponse
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response.Value, err
}

func (ops *AZDOOperations) GetPullRequests(repository string) ([]PullRequest, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s/pullrequests", ops.Organization, ops.Project, repository)
	var response AZDOResponse[PullRequest]
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response.Value, err
}

func (ops *AZDOOperations) GetVariableLibraries(ids ...int) ([]VarLib, error) {
	var sbUrl strings.Builder
	sbUrl.WriteString(fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/distributedTask/variableGroups", ops.Organization, ops.Project))
	if len(ids) != 0 {
		idStrings := make([]string, len(ids))
		for i, id := range ids {
			idStrings[i] = strconv.Itoa(id)
		}
		sbUrl.WriteString("?groupids=" + strings.Join(idStrings, ","))
	}
	url := sbUrl.String()

	var response AZDOResponse[VarLib]
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response.Value, err
}
