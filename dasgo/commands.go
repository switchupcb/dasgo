package dasgo

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
	FlagTypesStatusOnline       = "online"
	FlagTypesStatusDoNotDisturb = "dnd"
	FlagTypesStatusAFK          = "idle"
	FlagTypesStatusInvisible    = "invisible"
	FlagTypesStatusOffline      = "offline"
)
