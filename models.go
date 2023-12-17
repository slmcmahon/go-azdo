package azdo

type AZDOOperations struct {
	PAT          string
	Organization string
	Project      string
}

type AZDOResponse[T any] struct {
	Count int64 `json:"count"`
	Value []T   `json:"value"`
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

type CreatedBy struct {
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
	ID          string `json:"id"`
	UniqueName  string `json:"uniqueName"`
	ImageURL    string `json:"imageUrl"`
	Descriptor  string `json:"descriptor"`
}

type CompletionOptions struct {
	MergeCommitMessage          string        `json:"mergeCommitMessage"`
	DeleteSourceBranch          bool          `json:"deleteSourceBranch"`
	MergeStrategy               string        `json:"mergeStrategy"`
	AutoCompleteIgnoreConfigIDS []interface{} `json:"autoCompleteIgnoreConfigIds"`
	TransitionWorkItems         *bool         `json:"transitionWorkItems,omitempty"`
}

type LastMergeCommit struct {
	CommitID string `json:"commitId"`
	URL      string `json:"url"`
}

type VotedFor struct {
	ReviewerURL string `json:"reviewerUrl"`
	Vote        int64  `json:"vote"`
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
	ID          string `json:"id"`
	UniqueName  string `json:"uniqueName"`
	ImageURL    string `json:"imageUrl"`
	IsContainer bool   `json:"isContainer"`
}

type Reviewer struct {
	ReviewerURL string     `json:"reviewerUrl"`
	Vote        int64      `json:"vote"`
	HasDeclined bool       `json:"hasDeclined"`
	IsRequired  *bool      `json:"isRequired,omitempty"`
	IsFlagged   bool       `json:"isFlagged"`
	DisplayName string     `json:"displayName"`
	URL         string     `json:"url"`
	ID          string     `json:"id"`
	UniqueName  string     `json:"uniqueName"`
	ImageURL    string     `json:"imageUrl"`
	IsContainer *bool      `json:"isContainer,omitempty"`
	VotedFor    []VotedFor `json:"votedFor"`
}

type Label struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type PullRequest struct {
	Repository            Repository         `json:"repository"`
	PullRequestID         int64              `json:"pullRequestId"`
	CodeReviewID          int64              `json:"codeReviewId"`
	Status                string             `json:"status"`
	CreatedBy             CreatedBy          `json:"createdBy"`
	CreationDate          string             `json:"creationDate"`
	Title                 string             `json:"title"`
	Description           *string            `json:"description,omitempty"`
	SourceRefName         string             `json:"sourceRefName"`
	TargetRefName         string             `json:"targetRefName"`
	MergeStatus           string             `json:"mergeStatus"`
	IsDraft               bool               `json:"isDraft"`
	MergeID               string             `json:"mergeId"`
	LastMergeSourceCommit LastMergeCommit    `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit LastMergeCommit    `json:"lastMergeTargetCommit"`
	LastMergeCommit       *LastMergeCommit   `json:"lastMergeCommit,omitempty"`
	Reviewers             []Reviewer         `json:"reviewers"`
	URL                   string             `json:"url"`
	CompletionOptions     *CompletionOptions `json:"completionOptions,omitempty"`
	SupportsIterations    bool               `json:"supportsIterations"`
	AutoCompleteSetBy     *CreatedBy         `json:"autoCompleteSetBy,omitempty"`
	Labels                []Label            `json:"labels"`
}

// AzureDevOpsAPIResponse is an interface that different Azure DevOps response types can implement.
type AzureDevOpsAPIResponse interface{}
