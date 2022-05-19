// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import "time"

// List Public Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-active-threads-response-body
type ListPublicArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// List Private Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-active-threads-response-body
type ListPrivateArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// List Joined Private Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-active-threads-response-body
type ListJoinedPrivateArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// Get Gateway Response
// https://discord.com/developers/docs/topics/gateway#get-gateway-example-response
type GetGatewayResponse struct {
	URL string `json:"url,omitempty"`
}

// Get Gateway Bot Response
// https://discord.com/developers/docs/topics/gateway#get-gateway-example-response
type GetGatewayBotResponse struct {
	URL               string            `json:"url,omitempty"`
	Shards            *int              `json:"shards,omitempty"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

// Current Authorization Information Response Structure
// https://discord.com/developers/docs/topics/oauth2#get-current-authorization-information
type CurrentAuthorizationInformationResponse struct {
	Application *Application `json:"application"`
	Scopes      []*int       `json:"scopes"`
	Expires     *time.Time   `json:"expires"`
	User        *User        `json:"user"`
}

// Modify Current User Nick Response
// https://discord.com/developers/docs/topics/gateway#get-gateway-example-response
type ModifyCurrentUserNickResponse struct {
	Nick *string `json:"nick,omitempty"`
}
