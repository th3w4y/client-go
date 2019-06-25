/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Artifact describes a build product.
type ProvenanceArtifact struct {
	// Hash or checksum value of a binary, or Docker Registry 2.0 digest of a container.
	Checksum string `json:"checksum,omitempty"`
	// Artifact ID, if any; for container images, this will be a URL by digest like `gcr.io/projectID/imagename@sha256:123456`.
	Id string `json:"id,omitempty"`
	// Related artifact names. This may be the path to a binary or jar file, or in the case of a container build, the name used to push the container image to Google Container Registry, as presented to `docker push`. Note that a single Artifact ID can have multiple names, for example if two tags are applied to one image.
	Names []string `json:"names,omitempty"`
}
