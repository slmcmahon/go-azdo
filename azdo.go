package azdo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type AZDOOperations struct {
	PAT          string
	Organization string
	Project      string
}

type VarLibResponse struct {
	Count int64    `json:"count"`
	Value []VarLib `json:"value"`
}

type RepoListResponse struct {
	Count int64        `json:"count"`
	Value []Repository `json:"value"`
}

type VarLib struct {
	Variables                      map[string]Variable `json:"variables"`
	ID                             int64               `json:"id"`
	Type                           string              `json:"type"`
	Name                           string              `json:"name"`
	Description                    string              `json:"description"`
	CreatedBy                      Person              `json:"createdBy"`
	CreatedOn                      string              `json:"createdOn"`
	ModifiedBy                     Person              `json:"modifiedBy"`
	ModifiedOn                     string              `json:"modifiedOn"`
	IsShared                       bool                `json:"isShared"`
	VariableGroupProjectReferences interface{}         `json:"variableGroupProjectReferences"`
}

type Person struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
	UniqueName  string `json:"uniqueName"`
}

type Variable struct {
	Value string `json:"value"`
}

type Repository struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	DefaultBranch   string `json:"defaultBranch"`
	Size            int64  `json:"size"`
	RemoteURL       string `json:"remoteUrl"`
	SSHURL          string `json:"sshUrl"`
	WebURL          string `json:"webUrl"`
	IsDisabled      bool   `json:"isDisabled"`
	IsInMaintenance bool   `json:"isInMaintenance"`
}

type ChangeCounts struct {
	Add    int64 `json:"Add"`
	Edit   int64 `json:"Edit"`
	Delete int64 `json:"Delete"`
}

type Commit struct {
	CommitID     string       `json:"commitId"`
	Author       Person       `json:"author"`
	Committer    Person       `json:"committer"`
	Comment      string       `json:"comment"`
	ChangeCounts ChangeCounts `json:"changeCounts"`
	URL          string       `json:"url"`
	RemoteURL    string       `json:"remoteUrl"`
}

type Change struct {
	Item       ChangeItem `json:"item"`
	ChangeType string     `json:"changeType"`
}

type ChangeItem struct {
	ObjectID         string `json:"objectId"`
	OriginalObjectID string `json:"originalObjectId"`
	GitObjectType    string `json:"gitObjectType"`
	CommitID         string `json:"commitId"`
	Path             string `json:"path"`
	URL              string `json:"url"`
}

// AzureDevOpsAPIResponse is an interface that different Azure DevOps response types can implement.
type AzureDevOpsAPIResponse interface{}

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

func (ops *AZDOOperations) GetRepositories() (RepoListResponse, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories", ops.Organization, ops.Project)
	var response RepoListResponse
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response, err
}

func (ops *AZDOOperations) GetRepository(repsitory string) (Repository, error) {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories/%s", ops.Organization, ops.Project, repsitory)
	var response Repository
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response, err
}

func (ops *AZDOOperations) GetVariableLibraries(ids ...int) (VarLibResponse, error) {
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

	var response VarLibResponse
	err := azureDevOpsGetRequest(ops.PAT, url, &response)
	return response, err
}
