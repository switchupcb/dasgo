// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import "encoding/json"

// Gateway Payload Structure
// https://discord.com/developers/docs/topics/gateway#payloads-gateway-payload-structure
type GatewayPayload struct {
	Op             *Flag           `json:"op,omitempty"`
	Data           json.RawMessage `json:"d,omitempty"`
	SequenceNumber uint32          `json:"s,omitempty"`
	EventName      string          `json:"t,omitempty"`
}

// Presence Update Event Fields
// https://discord.com/developers/docs/topics/gateway#presence-update-presence-update-event-fields
type PresenceUpdate struct {
	User         *User         `json:"user,omitempty"`
	GuildID      Snowflake     `json:"guild_id,omitempty"`
	Status       string        `json:"status,omitempty"`
	Activities   []*Activity   `json:"activities,omitempty"`
	ClientStatus *ClientStatus `json:"client_status,omitempty"`
}

// Client Status Object
// https://discord.com/developers/docs/topics/gateway#client-status-object
type ClientStatus struct {
	Desktop *string `json:"desktop,omitempty"`
	Mobile  *string `json:"mobile,omitempty"`
	Web     *string `json:"web,omitempty"`
}
