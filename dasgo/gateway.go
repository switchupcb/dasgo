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

// Gateway URL Query String Params
// https://discord.com/developers/docs/topics/gateway#connecting-gateway-url-query-string-params
type GatewayURLQueryString struct {
	V        int    `json:"v,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Compress string `json:"compress,omitempty"`
}

// Session Start Limit Structure
// https://discord.com/developers/docs/topics/gateway#session-start-limit-object-session-start-limit-structure
type SessionStartLimit struct {
	Total          int `json:"total,omitempty"`
	Remaining      int `json:"remaining,omitempty"`
	ResetAfter     int `json:"reset_after,omitempty"`
	MaxConcurrency int `json:"max_concurrency,omitempty"`
}

// List of Intents
// https://discord.com/developers/docs/topics/gateway#list-of-intents
const (
	// GUILD_CREATE
	// GUILD_UPDATE
	// GUILD_DELETE
	// GUILD_ROLE_CREATE
	// GUILD_ROLE_UPDATE
	// GUILD_ROLE_DELETE
	// CHANNEL_CREATE
	// CHANNEL_UPDATE
	// CHANNEL_DELETE
	// CHANNEL_PINS_UPDATE
	// THREAD_CREATE
	// THREAD_UPDATE
	// THREAD_DELETE
	// THREAD_LIST_SYNC
	// THREAD_MEMBER_UPDATE
	// THREAD_MEMBERS_UPDATE *
	// STAGE_INSTANCE_CREATE
	// STAGE_INSTANCE_UPDATE
	// STAGE_INSTANCE_DELETE
	FlagIntentsofListGUILDS = 1 << 0

	// GUILD_MEMBER_ADD
	// GUILD_MEMBER_UPDATE
	// GUILD_MEMBER_REMOVE
	// THREAD_MEMBERS_UPDATE *
	FlagIntentsofListGUILD_MEMBERS = 1 << 1

	// GUILD_BAN_ADD
	// GUILD_BAN_REMOVE
	FlagIntentsofListGUILD_BANS = 1 << 2

	// GUILD_EMOJIS_UPDATE
	// GUILD_STICKERS_UPDATE
	FlagIntentsofListGUILD_EMOJIS_AND_STICKERS = 1 << 3

	// GUILD_INTEGRATIONS_UPDATE
	// INTEGRATION_CREATE
	// INTEGRATION_UPDATE
	// INTEGRATION_DELETE
	FlagIntentsofListGUILD_INTEGRATIONS = 1 << 4

	// WEBHOOKS_UPDATE
	FlagIntentsofListGUILD_WEBHOOKS = 1 << 5

	// INVITE_CREATE
	// INVITE_DELETE
	FlagIntentsofListGUILD_INVITES = 1 << 6

	// VOICE_STATE_UPDATE
	FlagIntentsofListGUILD_VOICE_STATES = 1 << 7

	// PRESENCE_UPDATE
	FlagIntentsofListGUILD_PRESENCES = 1 << 8

	// MESSAGE_CREATE
	// MESSAGE_UPDATE
	// MESSAGE_DELETE
	// MESSAGE_DELETE_BULK
	FlagIntentsofListGUILD_MESSAGES = 1 << 9

	// MESSAGE_REACTION_ADD
	// MESSAGE_REACTION_REMOVE
	// MESSAGE_REACTION_REMOVE_ALL
	// MESSAGE_REACTION_REMOVE_EMOJI
	FlagIntentsofListGUILD_MESSAGE_REACTIONS = 1 << 10

	// TYPING_START

	FlagIntentsofListGUILD_MESSAGE_TYPING = 1 << 11

	// MESSAGE_CREATE
	// MESSAGE_UPDATE
	// MESSAGE_DELETE
	// CHANNEL_PINS_UPDATE
	FlagIntentsofListDIRECT_MESSAGES = 1 << 12

	// MESSAGE_REACTION_ADD
	// MESSAGE_REACTION_REMOVE
	// MESSAGE_REACTION_REMOVE_ALL
	// MESSAGE_REACTION_REMOVE_EMOJI
	FlagIntentsofListDIRECT_MESSAGE_REACTIONS = 1 << 13

	// TYPING_START
	FlagIntentsofListDIRECT_MESSAGE_TYPING = 1 << 14

	// GUILD_SCHEDULED_EVENT_CREATE
	// GUILD_SCHEDULED_EVENT_UPDATE
	// GUILD_SCHEDULED_EVENT_DELETE
	// GUILD_SCHEDULED_EVENT_USER_ADD
	// GUILD_SCHEDULED_EVENT_USER_REMOVE
	FlagIntentsofListGUILD_SCHEDULED_EVENTS = 1 << 16
)
