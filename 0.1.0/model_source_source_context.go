/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A SourceContext is a reference to a tree of files. A SourceContext together with a path point to a unique revision of a single file or directory.
type SourceSourceContext struct {
	// A SourceContext referring to a revision in a Google Cloud Source Repo.
	CloudRepo *SourceCloudRepoSourceContext `json:"cloud_repo,omitempty"`
	// A SourceContext referring to a Gerrit project.
	Gerrit *SourceGerritSourceContext `json:"gerrit,omitempty"`
	// A SourceContext referring to any third party Git repo (e.g., GitHub).
	Git *SourceGitSourceContext `json:"git,omitempty"`
	// Labels with user defined metadata.
	Labels map[string]string `json:"labels,omitempty"`
}
