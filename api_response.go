/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucketv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// APIResponse contains generic data from API Response
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
	Values  map[string]interface{}
}

type SelfLink struct {
	Href string `json:"href"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type Links struct {
	Self []SelfLink `json:"self,omitempty"`
}

type Project struct {
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Type        string `json:"type"`
	Links       Links  `json:"links"`
}

// Repository contains data from a BitBucket Repository
type Repository struct {
	Slug          string  `json:"slug"`
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	ScmID         string  `json:"scmId"`
	State         string  `json:"state"`
	StatusMessage string  `json:"statusMessage"`
	Forkable      bool    `json:"forkable"`
	Project       Project `json:"project"`
	Public        bool    `json:"public"`
	Links         struct {
		Clone []CloneLink `json:"clone"`
		Self  []SelfLink  `json:"self"`
	} `json:"links"`
}

type UserWithNameEmail struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
}

type UserWithLinks struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ID           int    `json:"id,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Type         string `json:"type,omitempty"`
	Links        Links  `json:"links,omitempty"`
}

type User struct {
	Name                        string `json:"name"`
	EmailAddress                string `json:"emailAddress"`
	ID                          int    `json:"id"`
	DisplayName                 string `json:"displayName"`
	Active                      bool   `json:"active"`
	Slug                        string `json:"slug"`
	Type                        string `json:"type"`
	DirectoryName               string `json:"directoryName"`
	Deletable                   bool   `json:"deletable"`
	LastAuthenticationTimestamp int64  `json:"lastAuthenticationTimestamp"`
	MutableDetails              bool   `json:"mutableDetails"`
	MutableGroups               bool   `json:"mutableGroups"`
}

type UserWithMetadata struct {
	User               UserWithLinks `json:"user,omitempty"`
	Role               string        `json:"role,omitempty"`
	Approved           bool          `json:"approved,omitempty"`
	Status             string        `json:"status,omitempty"`
	LastReviewedCommit string        `json:"lastReviewedCommit,omitempty"`
}

// PermissionGlobal are global permissions
type PermissionGlobal string

const (
	// PermissionGlobalLicensedUser represents the ability to log into the system
	PermissionGlobalLicensedUser PermissionGlobal = "LICENSED_USER"
	// PermissionGlobalProjectCreate allows project creation
	PermissionGlobalProjectCreate PermissionGlobal = "PROJECT_CREATE"
	// PermissionGlobalAdmin represents an administrator
	PermissionGlobalAdmin PermissionGlobal = "ADMIN"
	// PermissionGlobalSysAdmin represents a system administrator
	PermissionGlobalSysAdmin PermissionGlobal = "SYS_ADMIN"
)

// PermissionProject are project level permissions
type PermissionProject string

const (
	// PermissionProjectAdmin grants admin priviledges
	PermissionProjectAdmin PermissionProject = "PROJECT_ADMIN"
	// PermissionProjectRead grants read priviledges
	PermissionProjectRead PermissionProject = "PROJECT_READ"
	// PermissionProjectWrite grants write priviledges
	PermissionProjectWrite PermissionProject = "PROJECT_WRITE"
)

// PermissionRepository are repository level permissions
type PermissionRepository string

const (
	// PermissionRepositoryAdmin grants admin priviledges
	PermissionRepositoryAdmin PermissionRepository = "REPO_ADMIN"
	// PermissionRepositoryRead grants read priviledges
	PermissionRepositoryRead PermissionRepository = "REPO_READ"
	// PermissionRepositoryWrite grants write priviledges
	PermissionRepositoryWrite PermissionRepository = "REPO_WRITE"
)

// UserPermission contains a user with its permission
type UserPermission struct {
	User       User   `json:"user"`
	Permission string `json:"permission"`
}

// Group represents a user group
type Group struct {
	Name string `json:"name"`
}

// GroupPermission contains a group with its permission
type GroupPermission struct {
	Group      Group  `json:"group"`
	Permission string `json:"permission"`
}

type MergeResult struct {
	Outcome string `json:"outcome"`
	Current bool   `json:"current"`
}

type MergeGetResponse struct {
	CanMerge   bool                     `json:"canMerge"`
	Conflicted bool                     `json:"conflicted"`
	Outcome    string                   `json:"outcome"`
	Vetoes     []map[string]interface{} `json:"vetoes"`
}

type PullRequestRef struct {
	ID           string     `json:"id"`
	DisplayID    string     `json:"displayId"`
	LatestCommit string     `json:"latestCommit"`
	Repository   Repository `json:"repository"`
}

type PullRequest struct {
	ID           int                `json:"id"`
	Version      int32              `json:"version"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	State        string             `json:"state"`
	Open         bool               `json:"open"`
	Closed       bool               `json:"closed"`
	CreatedDate  int64              `json:"createdDate"`
	UpdatedDate  int64              `json:"updatedDate"`
	FromRef      PullRequestRef     `json:"fromRef"`
	ToRef        PullRequestRef     `json:"toRef"`
	Locked       bool               `json:"locked"`
	Author       *UserWithMetadata  `json:"author,omitempty"`
	Reviewers    []UserWithMetadata `json:"reviewers"`
	Participants []UserWithMetadata `json:"participants,omitempty"`
	Properties   struct {
		MergeResult       MergeResult `json:"mergeResult"`
		ResolvedTaskCount int         `json:"resolvedTaskCount"`
		OpenTaskCount     int         `json:"openTaskCount"`
	} `json:"properties"`
	Links Links `json:"links"`
}

// SSHKey contains data from a SSHKey in the BitBucket Server
type SSHKey struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Label string `json:"label"`
}

// Commit contains data from a commit in BitBucket
type Commit struct {
	ID                 string `json:"id"`
	DisplayID          string `json:"displayId"`
	Author             User   `json:"author"`
	AuthorTimestamp    int64  `json:"authorTimestamp"`
	Committer          User   `json:"committer"`
	CommitterTimestamp int64  `json:"committerTimestamp"`
	Message            string `json:"message"`
	Parents            []struct {
		ID        string `json:"id"`
		DisplayID string `json:"displayId"`
	} `json:"parents"`
}

type BuildStatus struct {
	State       string `json:"state"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	DateAdded   int64  `json:"dateAdded"`
}

type Diff struct {
	Diffs []struct {
		Source struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"source"`
		Destination struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"destination"`
		Hunks []struct {
			SourceLine      int `json:"sourceLine"`
			SourceSpan      int `json:"sourceSpan"`
			DestinationLine int `json:"destinationLine"`
			DestinationSpan int `json:"destinationSpan"`
			Segments        []struct {
				Type  string `json:"type"`
				Lines []struct {
					Destination int    `json:"destination"`
					Source      int    `json:"source"`
					Line        string `json:"line"`
					Truncated   bool   `json:"truncated"`
				} `json:"lines"`
				Truncated bool `json:"truncated"`
			} `json:"segments"`
			Truncated bool `json:"truncated"`
		} `json:"hunks"`
		Truncated    bool    `json:"truncated"`
		ContextLines float64 `json:"contextLines"`
		FromHash     string  `json:"fromHash"`
		ToHash       string  `json:"toHash"`
		WhiteSpace   string  `json:"whiteSpace"`
	} `json:"diffs"`
	Truncated    bool    `json:"truncated"`
	ContextLines float64 `json:"contextLines"`
	FromHash     string  `json:"fromHash"`
	ToHash       string  `json:"toHash"`
	WhiteSpace   string  `json:"whiteSpace"`
}

// Tag contaings git Tag information
type Tag struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	Hash            string `json:"hash"`
}

// Branch contains git Branch information
type Branch struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	IsDefault       bool   `json:"isDefault"`
}

// Content contains repository content information (and files content)
type Content struct {
	Children struct {
		IsLastPage bool `json:"isLastPage"`
		Limit      int  `json:"limit"`
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Values     []struct {
			ContentID string `json:"contentId,omitempty"`
			Path      struct {
				Components []string `json:"components"`
				Extension  string   `json:"extension"`
				Name       string   `json:"name"`
				Parent     string   `json:"parent"`
				ToString   string   `json:"toString"`
			} `json:"path"`
			Size int    `json:"size,omitempty"`
			Type string `json:"type"`
			Node string `json:"node,omitempty"`
		} `json:"values"`
	} `json:"children"`
	Path struct {
		Components []string `json:"components"`
		Name       string   `json:"name"`
		ToString   string   `json:"toString"`
	} `json:"path"`
	Revision string `json:"revision"`
}

type WebhookConfiguration struct {
	Secret string `json:"secret"`
}

type Webhook struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Events        []string             `json:"events"`
	Url           string               `json:"url"`
	Active        bool                 `json:"active"`
	Configuration WebhookConfiguration `json:"configuration"`
}

// WebHookCallback contains payload to use while reading handling webhooks from bitbucket
type WebHookCallback struct {
	Actor      Actor      `json:"actor"`
	Repository Repository `json:"repository"`
	Push       struct {
		Changes []Change `json:"changes"`
	} `json:"push"`
}

// Actor contains the actor of reported changes from a webhook
type Actor struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

// Change contains changes reported by webhooks
type Change struct {
	Created bool        `json:"created"`
	Closed  bool        `json:"closed"`
	Old     interface{} `json:"old"`
	New     struct {
		Type   string `json:"type"`
		Name   string `json:"name"`
		Target struct {
			Type string `json:"type"`
			Hash string `json:"hash"`
		} `json:"target"`
	} `json:"new"`
}

// String converts global permission to its string representation
func (p PermissionGlobal) String() string {
	return string(p)
}

// String converts project permission to its string representation
func (p PermissionProject) String() string {
	return string(p)
}

// String converts repository permission to its string representation
func (p PermissionRepository) String() string {
	return string(p)
}

func (k *SSHKey) String() string {
	parts := make([]string, 1, 2)
	parts[0] = strings.TrimSpace(k.Text)
	return strings.Join(parts, " ")
}

// GetCommitsResponse cast Commits into structure
func GetCommitsResponse(r *APIResponse) ([]Commit, error) {
	var m []Commit
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetTagsResponse cast Tags into structure
func GetTagsResponse(r *APIResponse) ([]Tag, error) {
	var m []Tag
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetBranchesResponse cast Tags into structure
func GetBranchesResponse(r *APIResponse) ([]Branch, error) {
	var m []Branch
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoriesResponse cast Repositories into structure
func GetRepositoriesResponse(r *APIResponse) ([]Repository, error) {
	var m []Repository
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoryResponse cast Repositories into structure
func GetRepositoryResponse(r *APIResponse) (Repository, error) {
	var m Repository
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetDiffResponse cast Diff into structure
func GetDiffResponse(r *APIResponse) (Diff, error) {
	var m Diff
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetSSHKeysResponse cast SSHKeys into structure
func GetSSHKeysResponse(r *APIResponse) ([]SSHKey, error) {
	var m []SSHKey
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetBuildStatusResponse cast BuildStatus into structure
func GetBuildStatusesResponse(r *APIResponse) ([]BuildStatus, error) {
	var m []BuildStatus
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetPullRequestResponse cast PullRequest into structure
func GetPullRequestResponse(r *APIResponse) (PullRequest, error) {
	var m PullRequest
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetPullRequestResponse PullRequests into structure
func GetPullRequestsResponse(r *APIResponse) ([]PullRequest, error) {
	var m []PullRequest
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetContentResponse cast Content into structure
func GetContentResponse(r *APIResponse) (Content, error) {
	var c Content
	err := mapstructure.Decode(r.Values, &c)
	return c, err
}

// GetUserWithMetadataResponse casts users into structure
func GetUserWithMetadataResponse(r *APIResponse) (UserWithMetadata, error) {
	var c UserWithMetadata
	err := mapstructure.Decode(r.Values, &c)
	return c, err
}

// GetUsersResponse casts users into structure
func GetUsersResponse(r *APIResponse) ([]User, error) {
	var c []User
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetUsersPermissionResponse casts user permissions into structure
func GetUsersPermissionResponse(r *APIResponse) ([]UserPermission, error) {
	var c []UserPermission
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetGroupsPermissionResponse casts group permissions into structure
func GetGroupsPermissionResponse(r *APIResponse) ([]GroupPermission, error) {
	var c []GroupPermission
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetWebhooksResponse cast Webhooks into structure
func GetWebhooksResponse(r *APIResponse) ([]Webhook, error) {
	var h []Webhook
	err := mapstructure.Decode(r.Values["values"], &h)
	return h, err
}

// NewAPIResponse create new APIResponse from http.Response
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError create new erroneous API response from http.response and error
func NewAPIResponseWithError(r *http.Response, bodyBytes []byte, err error) (*APIResponse, error) {
	response := &APIResponse{Response: r, Message: strings.Replace(err.Error(), "\"", "", -1)}
	if bodyBytes != nil {
		parseErr := json.Unmarshal(bodyBytes, &response.Values)
		if parseErr != nil {
			err = fmt.Errorf("%s ParseError: %v", err.Error(), parseErr)
		}
	}
	return response, err
}

// NewBitbucketAPIResponse create new API response from http.response
func NewBitbucketAPIResponse(r *http.Response) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	err := json.NewDecoder(r.Body).Decode(&response.Values)
	if err != nil {
		return nil, err
	}
	return response, err
}

// NewRawAPIResponse create new API response from http.response with raw data
func NewRawAPIResponse(r *http.Response) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	response.Payload = raw
	return response, nil
}

// HasNextPage returns if response is paged and has next page and where it does start
func HasNextPage(response *APIResponse) (isLastPage bool, nextPageStart int) {
	isLastPage = response.Values["isLastPage"].(bool)
	if !isLastPage {
		nextPageStart = int(response.Values["nextPageStart"].(float64))
	}
	return !isLastPage, nextPageStart
}
