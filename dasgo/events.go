// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import "time"

// Gateway Events
// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
type Event interface{}

// Gateway Event Names
// https://discord.com/developers/docs/topics/gateway#event-names
const (
	FlagGatewayEventNameHello                               = "HELLO"
	FlagGatewayEventNameReady                               = "READY"
	FlagGatewayEventNameResumed                             = "RESUMED"
	FlagGatewayEventNameReconnect                           = "RECONNECT"
	FlagGatewayEventNameInvalidSession                      = "INVALID_SESSION"
	FlagGatewayEventNameApplicationCommandPermissionsUpdate = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	FlagGatewayEventNameChannelCreate                       = "CHANNEL_CREATE"
	FlagGatewayEventNameChannelUpdate                       = "CHANNEL_UPDATE"
	FlagGatewayEventNameChannelDelete                       = "CHANNEL_DELETE"
	FlagGatewayEventNameChannelPinsUpdate                   = "CHANNEL_PINS_UPDATE"
	FlagGatewayEventNameThreadCreate                        = "THREAD_CREATE"
	FlagGatewayEventNameThreadUpdate                        = "THREAD_UPDATE"
	FlagGatewayEventNameThreadDelete                        = "THREAD_DELETE"
	FlagGatewayEventNameThreadListSync                      = "THREAD_LIST_SYNC"
	FlagGatewayEventNameThreadMemberUpdate                  = "THREAD_MEMBER_UPDATE"
	FlagGatewayEventNameThreadMembersUpdate                 = "THREAD_MEMBERS_UPDATE"
	FlagGatewayEventNameGuildCreate                         = "GUILD_CREATE"
	FlagGatewayEventNameGuildUpdate                         = "GUILD_UPDATE"
	FlagGatewayEventNameGuildDelete                         = "GUILD_DELETE"
	FlagGatewayEventNameGuildBanAdd                         = "GUILD_BAN_ADD"
	FlagGatewayEventNameGuildBanRemove                      = "GUILD_BAN_REMOVE"
	FlagGatewayEventNameGuildEmojisUpdate                   = "GUILD_EMOJIS_UPDATE"
	FlagGatewayEventNameGuildStickersUpdate                 = "GUILD_STICKERS_UPDATE"
	FlagGatewayEventNameGuildIntegrationsUpdate             = "GUILD_INTEGRATIONS_UPDATE"
	FlagGatewayEventNameGuildMemberAdd                      = "GUILD_MEMBER_ADD"
	FlagGatewayEventNameGuildMemberRemove                   = "GUILD_MEMBER_REMOVE"
	FlagGatewayEventNameGuildMemberUpdate                   = "GUILD_MEMBER_UPDATE"
	FlagGatewayEventNameGuildMembersChunk                   = "GUILD_MEMBERS_CHUNK"
	FlagGatewayEventNameGuildRoleCreate                     = "GUILD_ROLE_CREATE"
	FlagGatewayEventNameGuildRoleUpdate                     = "GUILD_ROLE_UPDATE"
	FlagGatewayEventNameGuildRoleDelete                     = "GUILD_ROLE_DELETE"
	FlagGatewayEventNameGuildScheduledEventCreate           = "GUILD_SCHEDULED_EVENT_CREATE"
	FlagGatewayEventNameGuildScheduledEventUpdate           = "GUILD_SCHEDULED_EVENT_UPDATE"
	FlagGatewayEventNameGuildScheduledEventDelete           = "GUILD_SCHEDULED_EVENT_DELETE"
	FlagGatewayEventNameGuildScheduledEventUserAdd          = "GUILD_SCHEDULED_EVENT_USER_ADD"
	FlagGatewayEventNameGuildScheduledEventUserRemove       = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	FlagGatewayEventNameIntegrationCreate                   = "INTEGRATION_CREATE"
	FlagGatewayEventNameIntegrationUpdate                   = "INTEGRATION_UPDATE"
	FlagGatewayEventNameIntegrationDelete                   = "INTEGRATION_DELETE"
	FlagGatewayEventNameInteractionCreate                   = "INTERACTION_CREATE"
	FlagGatewayEventNameInviteCreate                        = "INVITE_CREATE"
	FlagGatewayEventNameInviteDelete                        = "INVITE_DELETE"
	FlagGatewayEventNameMessageCreate                       = "MESSAGE_CREATE"
	FlagGatewayEventNameMessageUpdate                       = "MESSAGE_UPDATE"
	FlagGatewayEventNameMessageDelete                       = "MESSAGE_DELETE"
	FlagGatewayEventNameMessageDeleteBulk                   = "MESSAGE_DELETE_BULK"
	FlagGatewayEventNameMessageReactionAdd                  = "MESSAGE_REACTION_ADD"
	FlagGatewayEventNameMessageReactionRemove               = "MESSAGE_REACTION_REMOVE"
	FlagGatewayEventNameMessageReactionRemoveAll            = "MESSAGE_REACTION_REMOVE_ALL"
	FlagGatewayEventNameMessageReactionRemoveEmoji          = "MESSAGE_REACTION_REMOVE_EMOJI"
	FlagGatewayEventNamePresenceUpdate                      = "PRESENCE_UPDATE"
	FlagGatewayEventNameStageInstanceCreate                 = "STAGE_INSTANCE_CREATE"
	FlagGatewayEventNameStageInstanceDelete                 = "STAGE_INSTANCE_DELETE"
	FlagGatewayEventNameStageInstanceUpdate                 = "STAGE_INSTANCE_UPDATE"
	FlagGatewayEventNameTypingStart                         = "TYPING_START"
	FlagGatewayEventNameUserUpdate                          = "USER_UPDATE"
	FlagGatewayEventNameVoiceStateUpdate                    = "VOICE_STATE_UPDATE"
	FlagGatewayEventNameVoiceServerUpdate                   = "VOICE_SERVER_UPDATE"
	FlagGatewayEventNameWebhooksUpdate                      = "WEBHOOKS_UPDATE"
)

// Hello Structure
// https://discord.com/developers/docs/topics/gateway#hello-hello-structure
type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// Ready Event Fields
// https://discord.com/developers/docs/topics/gateway#ready-ready-event-fields
type Ready struct {
	Version     int          `json:"v"`
	User        *User        `json:"user"`
	Guilds      []*Guild     `json:"guilds"`
	SessionID   string       `json:"session_id"`
	Shard       *[2]int      `json:"shard,omitempty"`
	Application *Application `json:"application"`
}

// Resumed
// https://discord.com/developers/docs/topics/gateway#resumed
type Resumed struct{}

// Reconnect
// https://discord.com/developers/docs/topics/gateway#reconnect
type Reconnect struct{}

// Invalid Session
// https://discord.com/developers/docs/topics/gateway#invalid-session
type InvalidSession struct {
	Data bool `json:"d"`
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
	GuildID    Snowflake       `json:"guild_id"`
	ChannelIDs []Snowflake     `json:"channel_ids,omitempty"`
	Threads    []*Channel      `json:"threads"`
	Members    []*ThreadMember `json:"members"`
}

// Thread Member Update
// https://discord.com/developers/docs/topics/gateway#thread-member-update
type ThreadMemberUpdate struct {
	*ThreadMember
	GuildID Snowflake `json:"guild_id"`
}

// Thread Members Update
// https://discord.com/developers/docs/topics/gateway#thread-members-update
type ThreadMembersUpdate struct {
	ID             Snowflake       `json:"id"`
	GuildID        Snowflake       `json:"guild_id"`
	MemberCount    int             `json:"member_count"`
	AddedMembers   []*ThreadMember `json:"added_members,omitempty"`
	RemovedMembers []Snowflake     `json:"removed_member_ids,omitempty"`
}

// Channel Pins Update
// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type ChannelPinsUpdate struct {
	GuildID          Snowflake `json:"guild_id,omitempty"`
	ChannelID        Snowflake `json:"channel_id"`
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
	GuildID string `json:"guild_id"`
	User    *User  `json:"user"`
}

// Guild Ban Remove
// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GuildBanRemove struct {
	GuildID Snowflake `json:"guild_id"`
	User    *User     `json:"user"`
}

// Guild Emojis Update
// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GuildEmojisUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []*Emoji  `json:"emojis"`
}

// Guild Stickers Update
// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GuildStickersUpdate struct {
	GuildID  Snowflake  `json:"guild_id"`
	Stickers []*Sticker `json:"stickers"`
}

// Guild Integrations Update
// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GuildIntegrationsUpdate struct {
	GuildID Snowflake `json:"guild_id"`
}

// Guild Member Add
// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GuildMemberAdd struct {
	GuildID Snowflake `json:"guild_id"`
	*GuildMember
}

// Guild Member Remove
// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GuildMemberRemove struct {
	GuildID Snowflake `json:"guild_id"`
	User    *User     `json:"user"`
}

// Guild Member Update
// https://discord.com/developers/docs/topics/gateway#guild-member-update
type GuildMemberUpdate struct {
	*GuildMember
}

// Guild Members Chunk
// https://discord.com/developers/docs/topics/gateway#guild-members-chunk
type GuildMembersChunk struct {
	GuildID    Snowflake         `json:"guild_id"`
	Members    []*GuildMember    `json:"members"`
	ChunkIndex int               `json:"chunk_index"`
	ChunkCount int               `json:"chunk_count"`
	Presences  []*PresenceUpdate `json:"presences,omitempty"`
	NotFound   []Snowflake       `json:"not_found,omitempty"`
	Nonce      *string           `json:"nonce,omitempty"`
}

// Guild Role Create
// https://discord.com/developers/docs/topics/gateway#guild-role-create
type GuildRoleCreate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    *Role     `json:"role"`
}

// Guild Role Update
// https://discord.com/developers/docs/topics/gateway#guild-role-update
type GuildRoleUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    *Role     `json:"role"`
}

// Guild Role Delete
// https://discord.com/developers/docs/topics/gateway#guild-role-delete
type GuildRoleDelete struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
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
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

// Guild Scheduled Event User Remove
// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-remove
type GuildScheduledEventUserRemove struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

// Integration Create
// https://discord.com/developers/docs/topics/gateway#integration-create
type IntegrationCreate struct {
	*Integration
	GuildID Snowflake `json:"guild_id"`
}

// Integration Update
// https://discord.com/developers/docs/topics/gateway#integration-update
type IntegrationUpdate struct {
	*Integration
	GuildID Snowflake `json:"guild_id"`
}

// Integration Delete
// https://discord.com/developers/docs/topics/gateway#integration-delete
type IntegrationDelete struct {
	IntegrationID Snowflake `json:"id"`
	GuildID       Snowflake `json:"guild_id"`
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
	ChannelID         Snowflake    `json:"channel_id"`
	Code              string       `json:"code"`
	CreatedAt         time.Time    `json:"created_at"`
	GuildID           Snowflake    `json:"guild_id,omitempty"`
	Inviter           *User        `json:"inviter,omitempty"`
	MaxAge            int          `json:"max_age"`
	MaxUses           int          `json:"max_uses"`
	TargetType        *int         `json:"target_user_type,omitempty"`
	TargetUser        *User        `json:"target_user,omitempty"`
	TargetApplication *Application `json:"target_application,omitempty"`
	Temporary         bool         `json:"temporary"`
	Uses              int          `json:"uses"`
}

// Invite Delete
// https://discord.com/developers/docs/topics/gateway#invite-delete
type InviteDelete struct {
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
	Code      Snowflake `json:"code"`
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
	MessageID Snowflake `json:"id"`
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
}

// Message Delete Bulk
// https://discord.com/developers/docs/topics/gateway#message-delete-bulk
type MessageDeleteBulk struct {
	MessageIDs []Snowflake `json:"ids"`
	ChannelID  Snowflake   `json:"channel_id"`
	GuildID    Snowflake   `json:"guild_id,omitempty"`
}

// Message Reaction Add
// https://discord.com/developers/docs/topics/gateway#message-reaction-add
type MessageReactionAdd struct {
	UserID    Snowflake    `json:"user_id"`
	ChannelID Snowflake    `json:"channel_id"`
	MessageID Snowflake    `json:"message_id"`
	GuildID   Snowflake    `json:"guild_id,omitempty"`
	Member    *GuildMember `json:"member,omitempty"`
	Emoji     *Emoji       `json:"emoji"`
}

// Message Reaction Remove
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove
type MessageReactionRemove struct {
	UserID    Snowflake `json:"user_id"`
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
	Emoji     *Emoji    `json:"emoji"`
}

// Message Reaction Remove All
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-all
type MessageReactionRemoveAll struct {
	ChannelID Snowflake `json:"channel_id"`
	MessageID Snowflake `json:"message_id"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
}

// Message Reaction Remove Emoji
// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-emoji
type MessageReactionRemoveEmoji struct {
	ChannelID Snowflake `json:"channel_id"`
	GuildID   Snowflake `json:"guild_id,omitempty"`
	MessageID Snowflake `json:"message_id"`
	Emoji     *Emoji    `json:"emoji"`
}

// Presence Update Event Fields
// https://discord.com/developers/docs/topics/gateway#presence-update-presence-update-event-fields
type PresenceUpdate struct {
	User         *User         `json:"user"`
	GuildID      Snowflake     `json:"guild_id"`
	Status       string        `json:"status"`
	Activities   []*Activity   `json:"activities"`
	ClientStatus *ClientStatus `json:"client_status"`
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
	ChannelID Snowflake    `json:"channel_id"`
	GuildID   Snowflake    `json:"guild_id,omitempty"`
	UserID    Snowflake    `json:"user_id"`
	Member    *GuildMember `json:"member,omitempty"`
	Timestamp int          `json:"timestamp"`
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
	Token    string    `json:"token"`
	GuildID  Snowflake `json:"guild_id"`
	Endpoint string    `json:"endpoint"`
}

// Webhooks Update
// https://discord.com/developers/docs/topics/gateway#webhooks-update
type WebhooksUpdate struct {
	GuildID   Snowflake `json:"guild_id"`
	ChannelID Snowflake `json:"channel_id"`
}
