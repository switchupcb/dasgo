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
	FlagIntentGUILDS = 1 << 0

	// GUILD_MEMBER_ADD
	// GUILD_MEMBER_UPDATE
	// GUILD_MEMBER_REMOVE
	// THREAD_MEMBERS_UPDATE *
	FlagIntentGUILD_MEMBERS = 1 << 1

	// GUILD_BAN_ADD
	// GUILD_BAN_REMOVE
	FlagIntentGUILD_BANS = 1 << 2

	// GUILD_EMOJIS_UPDATE
	// GUILD_STICKERS_UPDATE
	FlagIntentGUILD_EMOJIS_AND_STICKERS = 1 << 3

	// GUILD_INTEGRATIONS_UPDATE
	// INTEGRATION_CREATE
	// INTEGRATION_UPDATE
	// INTEGRATION_DELETE
	FlagIntentGUILD_INTEGRATIONS = 1 << 4

	// WEBHOOKS_UPDATE
	FlagIntentGUILD_WEBHOOKS = 1 << 5

	// INVITE_CREATE
	// INVITE_DELETE
	FlagIntentGUILD_INVITES = 1 << 6

	// VOICE_STATE_UPDATE
	FlagIntentGUILD_VOICE_STATES = 1 << 7

	// PRESENCE_UPDATE
	FlagIntentGUILD_PRESENCES = 1 << 8

	// MESSAGE_CREATE
	// MESSAGE_UPDATE
	// MESSAGE_DELETE
	// MESSAGE_DELETE_BULK
	FlagIntentGUILD_MESSAGES = 1 << 9

	// MESSAGE_REACTION_ADD
	// MESSAGE_REACTION_REMOVE
	// MESSAGE_REACTION_REMOVE_ALL
	// MESSAGE_REACTION_REMOVE_EMOJI
	FlagIntentGUILD_MESSAGE_REACTIONS = 1 << 10

	// TYPING_START

	FlagIntentGUILD_MESSAGE_TYPING = 1 << 11

	// MESSAGE_CREATE
	// MESSAGE_UPDATE
	// MESSAGE_DELETE
	// CHANNEL_PINS_UPDATE
	FlagIntentDIRECT_MESSAGES = 1 << 12

	// MESSAGE_REACTION_ADD
	// MESSAGE_REACTION_REMOVE
	// MESSAGE_REACTION_REMOVE_ALL
	// MESSAGE_REACTION_REMOVE_EMOJI
	FlagIntentDIRECT_MESSAGE_REACTIONS = 1 << 13

	// TYPING_START
	FlagIntentDIRECT_MESSAGE_TYPING = 1 << 14

	// GUILD_SCHEDULED_EVENT_CREATE
	// GUILD_SCHEDULED_EVENT_UPDATE
	// GUILD_SCHEDULED_EVENT_DELETE
	// GUILD_SCHEDULED_EVENT_USER_ADD
	// GUILD_SCHEDULED_EVENT_USER_REMOVE
	FlagIntentGUILD_SCHEDULED_EVENTS = 1 << 16
)

// Gateway Commands
// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-commands
type Command interface{}

// Identify Structure
// https://discord.com/developers/docs/topics/gateway#identify-identify-structure
type Identify struct {
	Token          string                       `json:"token,omitempty"`
	Properties     IdentifyConnectionProperties `json:"properties,omitempty"`
	Compress       bool                         `json:"compress,omitempty"`
	LargeThreshold int                          `json:"large_threshold,omitempty"`
	Shard          *[2]int                      `json:"shard,omitempty"`
	Presence       GatewayPresenceUpdate        `json:"presence,omitempty"`
	Intents        BitFlag                      `json:"intents,omitempty"`
}

// Identify Connection Properties
// https://discord.com/developers/docs/topics/gateway#identify-identify-connection-properties
type IdentifyConnectionProperties struct {
	OS      string `json:"$os,omitempty"`
	Browser string `json:"$browser,omitempty"`
	Device  string `json:"$device,omitempty"`
}

// Resume Structure
// https://discord.com/developers/docs/topics/gateway#resume-resume-structure
type Resume struct {
	Token     string `json:"token,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	Seq       uint32 `json:"seq,omitempty"`
}

// Heartbeat
// https://discord.com/developers/docs/topics/gateway#heartbeat
type Heartbeat struct {
	Op   int   `json:"op,omitempty"`
	Data int64 `json:"d,omitempty"`
}

// Guild Request Members Structure
// https://discord.com/developers/docs/topics/gateway#request-guild-members-guild-request-members-structure
type GuildRequestMembers struct {
	GuildID   Snowflake   `json:"guild_id,omitempty"`
	Query     string      `json:"query,omitempty"`
	Limit     uint        `json:"limit,omitempty"`
	Presences bool        `json:"presences,omitempty"`
	UserIDs   []Snowflake `json:"user_ids,omitempty"`
	Nonce     string      `json:"nonce,omitempty"`
}

// Gateway Voice State Update Structure
// https://discord.com/developers/docs/topics/gateway#update-voice-state-gateway-voice-state-update-structure
type GatewayVoiceStateUpdate struct {
	GuildID   Snowflake `json:"guild_id,omitempty"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
	SelfMute  bool      `json:"self_mute,omitempty"`
	SelfDeaf  bool      `json:"self_deaf,omitempty"`
}

// Gateway Presence Update Structure
// https://discord.com/developers/docs/topics/gateway#update-presence-gateway-presence-update-structure
type GatewayPresenceUpdate struct {
	Since  int         `json:"since,omitempty"`
	Game   []*Activity `json:"game,omitempty"`
	Status string      `json:"status,omitempty"`
	AFK    bool        `json:"afk,omitempty"`
}

// Status Types
// https://discord.com/developers/docs/topics/gateway#update-presence-status-types
const (
	FlagStatusTypeOnline       = "online"
	FlagStatusTypeDoNotDisturb = "dnd"
	FlagStatusTypeAFK          = "idle"
	FlagStatusTypeInvisible    = "invisible"
	FlagStatusTypeOffline      = "offline"
)
