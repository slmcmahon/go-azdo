package azdo

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

type CommitListResponse struct {
	Count int64    `json:"count"`
	Value []Commit `json:"value"`
}

type ChangeListResponse struct {
	Counts ChangeCounts `json:"changecounts"`
	Value  []Change     `json:"changes"`
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
