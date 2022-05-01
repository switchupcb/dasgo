package dasgo

import "time"

// Gateway Events
// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
type Event interface{}

// Hello Structure
// https://discord.com/developers/docs/topics/gateway#hello-hello-structure
type Hello struct {
	HeartbeatInterval time.Duration `json:"heartbeat_interval,omitempty"`
}

// Ready Event Fields
// https://discord.com/developers/docs/topics/gateway#ready-ready-event-fields
type Ready struct {
	Version     int          `json:"v,omitempty"`
	User        *User        `json:"user,omitempty"`
	Guilds      []*Guild     `json:"guilds,omitempty"`
	SessionID   string       `json:"session_id,omitempty"`
	Shard       *[2]int      `json:"shard,omitempty"`
	Application *Application `json:"application,omitempty"`
}

// Resumed
// https://discord.com/developers/docs/topics/gateway#resumed
type Resumed struct {
	Op int `json:"op,omitempty"`
}

// Reconnect
// https://discord.com/developers/docs/topics/gateway#reconnect
type Reconnect struct {
	Op int `json:"op,omitempty"`
}

// Invalid Session
// https://discord.com/developers/docs/topics/gateway#invalid-session
type InvalidSession struct {
	Op   int  `json:"op,omitempty"`
	Data bool `json:"d,omitempty"`
}

// Application Command Permissions Update
// https://discord.com/developers/docs/topics/gateway#application-command-permissions-update
type ApplicationCommandPermissionsUpdate struct {
	*GuildApplicationCommandPermissions
}

// Channel Create
// https://discord.com/developers/docs/topics/gateway#channel-create
type ChannelCreate struct {
	*Channel
}

// Channel Update
// https://discord.com/developers/docs/topics/gateway#channel-update
type ChannelUpdate struct {
	*Channel
}

// Channel Delete
// https://discord.com/developers/docs/topics/gateway#channel-delete
type ChannelDelete struct {
	*Channel
}

// Thread Create
// https://discord.com/developers/docs/topics/gateway#thread-create
type ThreadCreate struct {
	*Channel
	NewlyCreated bool `json:"newly_created,omitempty"`
}

// Thread Update
// https://discord.com/developers/docs/topics/gateway#thread-update
type ThreadUpdate struct {
	*Channel
}

// Thread Delete
// https://discord.com/developers/docs/topics/gateway#thread-delete
type ThreadDelete struct {
	*Channel
}

// Thread List Sync Event Fields
// https://discord.com/developers/docs/topics/gateway#thread-list-sync
type ThreadListSync struct {
	GuildID    Snowflake       `json:"guild_id,omitempty"`
	ChannelIDs []Snowflake     `json:"channel_ids,omitempty"`
	Threads    []*Channel      `json:"threads,omitempty"`
	Members    []*ThreadMember `json:"members,omitempty"`
}

// Thread Member Update
// https://discord.com/developers/docs/topics/gateway#thread-member-update
type ThreadMemberUpdate struct {
	*ThreadMember
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Thread Members Update
// https://discord.com/developers/docs/topics/gateway#thread-members-update
type ThreadMembersUpdate struct {
	ID             Snowflake       `json:"id,omitempty"`
	GuildID        Snowflake       `json:"guild_id,omitempty"`
	MemberCount    int             `json:"member_count,omitempty"`
	AddedMembers   []*ThreadMember `json:"added_members,omitempty"`
	RemovedMembers []Snowflake     `json:"removed_member_ids,omitempty"`
}

// Channel Pins Update
// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type ChannelPinsUpdate struct {
	GuildID          Snowflake `json:"guild_id,omitempty"`
	ChannelID        Snowflake `json:"channel_id,omitempty"`
	LastPinTimestamp time.Time `json:"last_pin_timestamp,omitempty"`
}

// Guild Create
// https://discord.com/developers/docs/topics/gateway#guild-create
type GuildCreate struct {
	*Guild

	// https://discord.com/developers/docs/topics/threads#gateway-events
	Threads []*Channel `json:"threads,omitempty"`
}

// Guild Update
// https://discord.com/developers/docs/topics/gateway#guild-update
type GuildUpdate struct {
	*Guild
}

// Guild Delete
// https://discord.com/developers/docs/topics/gateway#guild-delete
type GuildDelete struct {
	*Guild
}

// Guild Ban Add
// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GuildBanAdd struct {
	GuildID string `json:"guild_id,omitempty"`
	User    *User  `json:"user,omitempty"`
}

// Guild Ban Remove
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GuildBanRemove struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	User    *User     `json:"user,omitempty"`
}

// Guild Emojis Update
// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GuildEmojisUpdate struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	Emojis  []*Emoji  `json:"emojis,omitempty"`
}

// Guild Stickers Update
// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GuildStickersUpdate struct {
	GuildID  Snowflake  `json:"guild_id,omitempty"`
	Stickers []*Sticker `json:"stickers,omitempty"`
}

// Guild Integrations Update
// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GuildIntegrationsUpdate struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Guild Member Add
// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GuildMemberAdd struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	*GuildMember
}

// Guild Member Remove
// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GuildMemberRemove struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	User    *User     `json:"user,omitempty"`
}

// Guild Member Update
// https://discord.com/developers/docs/topics/gateway#guild-member-update
type GuildMemberUpdate struct {
	*GuildMember
}

// Guild Members Chunk
// https://discord.com/developers/docs/topics/gateway#guild-members-chunk
type GuildMembersChunk struct {
	GuildID    Snowflake         `json:"guild_id,omitempty"`
	Members    []*GuildMember    `json:"members,omitempty"`
	ChunkIndex int               `json:"chunk_index,omitempty"`
	ChunkCount int               `json:"chunk_count,omitempty"`
	Presences  []*PresenceUpdate `json:"presences,omitempty"`
	NotFound   []Snowflake       `json:"not_found,omitempty"`
	Nonce      string            `json:"nonce,omitempty"`
}

// Guild Role Create
// https://discord.com/developers/docs/topics/gateway#guild-role-create
type GuildRoleCreate struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	Role    *Role     `json:"role,omitempty"`
}

// Guild Role Update
// https://discord.com/developers/docs/topics/gateway#guild-role-update
type GuildRoleUpdate struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	Role    *Role     `json:"role,omitempty"`
}

// Guild Role Delete
// https://discord.com/developers/docs/topics/gateway#guild-role-delete
type GuildRoleDelete struct {
	GuildID Snowflake `json:"guild_id,omitempty"`
	RoleID  Snowflake `json:"role_id,omitempty"`
}

// Guild Scheduled Event Create
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-create
type GuildScheduledEventCreate struct {
	*GuildScheduledEvent
}

// Guild Scheduled Event Update
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-update
type GuildScheduledEventUpdate struct {
	*GuildScheduledEvent
}

// Guild Scheduled Event Delete
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-delete
type GuildScheduledEventDelete struct {
	*GuildScheduledEvent
}

// Guild Scheduled Event User Add
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-add
type GuildScheduledEventUserAdd struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id,omitempty"`
	UserID                Snowflake `json:"user_id,omitempty"`
	GuildID               Snowflake `json:"guild_id,omitempty"`
}

// Guild Scheduled Event User Remove
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-remove
type GuildScheduledEventUserRemove struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id,omitempty"`
	UserID                Snowflake `json:"user_id,omitempty"`
	GuildID               Snowflake `json:"guild_id,omitempty"`
}

// Integration Create
// https://discord.com/developers/docs/topics/gateway#integration-create
type IntegrationCreate struct {
	*Integration
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Integration Update
// https://discord.com/developers/docs/topics/gateway#integration-update
type IntegrationUpdate struct {
	*Integration
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// Integration Delete
// https://discord.com/developers/docs/topics/gateway#integration-delete
type IntegrationDelete struct {
	IntegrationID Snowflake `json:"id,omitempty"`
	GuildID       Snowflake `json:"guild_id,omitempty"`
	ApplicationID Snowflake `json:"application_id,omitempty"`
}

// Interaction Create
// https://discord.com/developers/docs/topics/gateway#interaction-create
type InteractionCreate struct {
	*Interaction
}

// Invite Create
// https://discord.com/developers/docs/topics/gateway#invite-create
type InviteCreate struct {
	ChannelID         Snowflake    `json:"channel_id,omitempty"`
	Code              string       `json:"code,omitempty"`
	CreatedAt         time.Time    `json:"created_at,omitempty"`
	GuildID           Snowflake    `json:"guild_id,omitempty"`
	Inviter           *User        `json:"inviter,omitempty"`
	MaxAge            int          `json:"max_age,omitempty"`
	MaxUses           int          `json:"max_uses,omitempty"`
	TargetType        int          `json:"target_user_type,omitempty"`
	TargetUser        *User        `json:"target_user,omitempty"`
	TargetApplication *Application `json:"target_application,omitempty"`
	Temporary         bool         `json:"temporary,omitempty"`
	Uses              int          `json:"uses,omitempty"`
}

// Invite Delete
// https://discord.com/developers/docs/topics/gateway#invite-delete
type InviteDelete struct {
	ChannelID Snowflake `json:"channel_id,omitempty"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
	Code      Snowflake `json:"code,omitempty"`
}

// Message Create
// https://discord.com/developers/docs/topics/gateway#message-create
type MessageCreate struct {
	*Message
}

// Message Update
// https://discord.com/developers/docs/topics/gateway#message-update
type MessageUpdate struct {
	Message *Message
}

// Message Delete
// https://discord.com/developers/docs/topics/gateway#message-delete
type MessageDelete struct {
	MessageID Snowflake `json:"id,omitempty"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
}

// Message Delete Bulk
// https://discord.com/developers/docs/topics/gateway#message-delete-bulk
type MessageDeleteBulk struct {
	MessageIDs []Snowflake `json:"ids,omitempty"`
	ChannelID  Snowflake   `json:"channel_id,omitempty"`
	GuildID    Snowflake   `json:"guild_id,omitempty"`
}

// Message Reaction Add
// https://discord.com/developers/docs/topics/gateway#message-reaction-add
type MessageReactionAdd struct {
	UserID    Snowflake    `json:"user_id,omitempty"`
	ChannelID Snowflake    `json:"channel_id,omitempty"`
	MessageID Snowflake    `json:"message_id,omitempty"`
	GuildID   Snowflake    `json:"guild_id,omitempty"`
	Member    *GuildMember `json:"member,omitempty"`
	Emoji     *Emoji       `json:"emoji,omitempty"`
}

// Message Reaction Remove
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove
type MessageReactionRemove struct {
	UserID    Snowflake    `json:"user_id,omitempty"`
	ChannelID Snowflake    `json:"channel_id,omitempty"`
	MessageID Snowflake    `json:"message_id,omitempty"`
	GuildID   Snowflake    `json:"guild_id,omitempty"`
	Member    *GuildMember `json:"member,omitempty"`
	Emoji     *Emoji       `json:"emoji,omitempty"`
}

// Message Reaction Remove All
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-all
type MessageReactionRemoveAll struct {
	ChannelID Snowflake `json:"channel_id,omitempty"`
	MessageID Snowflake `json:"message_id,omitempty"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
}

// Message Reaction Remove Emoji
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-emoji
type MessageReactionRemoveEmoji struct {
	ChannelID Snowflake `json:"channel_id,omitempty"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
	MessageID Snowflake `json:"message_id,omitempty"`
	Emoji     *Emoji    `json:"emoji,omitempty"`
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

// Stage Instance Create
// https://discord.com/developers/docs/topics/gateway#stage-instance-create
type StageInstanceCreate struct {
	*StageInstance
}

// Stage Instance Update
// https://discord.com/developers/docs/topics/gateway#stage-instance-update
type StageInstanceUpdate struct {
	*StageInstance
}

// Stage Instance Delete
// https://discord.com/developers/docs/topics/gateway#stage-instance-delete
type StageInstanceDelete struct {
	*StageInstance
}

// Typing Start
// https://discord.com/developers/docs/topics/gateway#typing-start
type TypingStart struct {
	ChannelID Snowflake    `json:"channel_id,omitempty"`
	GuildID   Snowflake    `json:"guild_id,omitempty"`
	UserID    Snowflake    `json:"user_id,omitempty"`
	Member    *GuildMember `json:"member,omitempty"`
	Timestamp int          `json:"timestamp,omitempty"`
}

// User Update
// https://discord.com/developers/docs/topics/gateway#user-update
type UserUpdate struct {
	*User
}

// Voice State Update
// https://discord.com/developers/docs/topics/gateway#voice-state-update
type VoiceStateUpdate struct {
	*VoiceState
}

// Voice Server Update
// https://discord.com/developers/docs/topics/gateway#voice-server-update
type VoiceServerUpdate struct {
	Token    string    `json:"token,omitempty"`
	GuildID  Snowflake `json:"guild_id,omitempty"`
	Endpoint string    `json:"endpoint,omitempty"`
}

// Webhooks Update
// https://discord.com/developers/docs/topics/gateway#webhooks-update
type WebhooksUpdate struct {
	GuildID   Snowflake `json:"guild_id,omitempty"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
}
