package azdo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type AZDOResponse struct {
	Count int64    `json:"count"`
	Value []VarLib `json:"value"`
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

// CompareAndPrintDifference compares two variable groups and prints variables that are in the first group but not in the second.
func CompareAndPrintDifference(group1Name string, group1Vars map[string]struct{}, group2Name string, group2Vars map[string]struct{}) {
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

func GetVariableLibraries(pat string, org string, project string, ids ...int) (AZDOResponse, error) {
	var sbUrl strings.Builder
	sbUrl.WriteString(fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/distributedTask/variableGroups", org, project))
	if len(ids) != 0 {
		sbUrl.WriteString("?groupids=")
		for i, id := range ids {
			sbUrl.WriteString(strconv.Itoa(id))
			if i < len(ids)-1 {
				sbUrl.WriteString(",")
			}
		}
	}
	url := sbUrl.String()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AZDOResponse{}, err
	}

	req.SetBasicAuth("", pat)

	resp, err := client.Do(req)
	if err != nil {
		return AZDOResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AZDOResponse{}, err
	}

	var response AZDOResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return AZDOResponse{}, err
	}
	return response, nil
}
