/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// An occurrence of a particular package installation found within a system's filesystem. E.g., glibc was found in `/var/lib/dpkg/status`.
type V1beta1packageLocation struct {
	// Required. The CPE URI in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
	CpeUri string `json:"cpe_uri,omitempty"`
	// The version installed at this location.
	Version *PackageVersion `json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path string `json:"path,omitempty"`
}
