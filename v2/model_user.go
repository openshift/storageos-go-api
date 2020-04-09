/*
 * StorageOS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 * Contact: info@storageos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api
import (
	"time"
)
// User struct for User
type User struct {
	// A unique identifier for a user. The format of this type is undefined and may change but the defined properties will not change.. 
	Id string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	// If true, this user is an administrator of the cluster. Administrators bypass the usual authentication checks and are granted access to all resources. Some actions (such as adding a new user) can only be performed by an administrator. 
	IsAdmin bool `json:"isAdmin,omitempty"`
	// Defines a set of policy group IDs this user is a member of. Policy groups can be used to logically group users and  apply authorisation policies to all members. 
	Groups *[]string `json:"groups,omitempty"`
	// The time the entity was created. This timestamp is set by the node that created the entity, and may not be correct if the node's local clock was skewed. This value is for the user's informative purposes only, and correctness is not required. String format is RFC3339. 
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// The time the entity was last updated. This timestamp is set by the node that last updated the entity, and may not be correct if the node's local clock was skewed. This value is for the user's informative purposes only, and correctness is not required. String format is RFC3339. 
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// An opaque representation of an entity version at the time it was obtained from the API. All operations that mutate the entity must include this version field in the request unchanged. The format of this type is undefined and may change but the defined properties will not change. 
	Version string `json:"version,omitempty"`
}