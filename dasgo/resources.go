// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import (
	"time"
)

// Application Command Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	ID                       Snowflake                   `json:"id,omitempty"`
	Type                     Flag                        `json:"type,omitempty"`
	ApplicationID            Snowflake                   `json:"application_id,omitempty"`
	GuildID                  Snowflake                   `json:"guild_id,omitempty"`
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions string                      `json:"default_member_permissions,omitempty"`
	DMPermission             bool                        `json:"dm_permission,omitempty"`
	Version                  Snowflake                   `json:"version,omitempty"`
}

// Application Command Types
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
const (
	FlagTypesCommandApplicationCHAT_INPUT = 1
	FlagTypesCommandApplicationUSER       = 2
	FlagTypesCommandApplicationMESSAGE    = 3
)

// Application Command Option Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     Flag                              `json:"type,omitempty"`
	Name                     string                            `json:"name,omitempty"`
	NameLocalizations        map[Flag]string                   `json:"name_localizations,omitempty"`
	Description              string                            `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string                   `json:"description_localizations,omitempty"`
	Required                 bool                              `json:"required,omitempty"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []*ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []*Flag                           `json:"channel_types,omitempty"`
	MinValue                 float64                           `json:"min_value,omitempty"`
	MaxValue                 float64                           `json:"max_value,omitempty"`
	Autocomplete             bool                              `json:"autocomplete,omitempty"`
}

// Application Command Option Type
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
const (
	FlagTypeOptionCommandApplicationSUB_COMMAND       = 1
	FlagTypeOptionCommandApplicationSUB_COMMAND_GROUP = 2
	FlagTypeOptionCommandApplicationSTRING            = 3
	FlagTypeOptionCommandApplicationINTEGER           = 4
	FlagTypeOptionCommandApplicationBOOLEAN           = 5
	FlagTypeOptionCommandApplicationUSER              = 6
	FlagTypeOptionCommandApplicationCHANNEL           = 7
	FlagTypeOptionCommandApplicationROLE              = 8
	FlagTypeOptionCommandApplicationMENTIONABLE       = 9
	FlagTypeOptionCommandApplicationNUMBER            = 10
	FlagTypeOptionCommandApplicationATTACHMENT        = 11
)

// Application Command Option Choice
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure
type ApplicationCommandOptionChoice struct {
	Name              string          `json:"name,omitempty"`
	NameLocalizations map[Flag]string `json:"name_localizations,omitempty"`
	Value             interface{}     `json:"value,omitempty"`
}

// Application Command Interaction Data Option Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name,omitempty"`
	Type    Flag                                       `json:"type,omitempty"`
	Value   interface{}                                `json:"value,omitempty"`
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused bool                                       `json:"focused,omitempty"`
}

// Guild Application Command Permissions Object
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                        `json:"id,omitempty"`
	ApplicationID Snowflake                        `json:"application_id,omitempty"`
	GuildID       Snowflake                        `json:"guild_id,omitempty"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions,omitempty"`
}

// Application Command Permissions Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure
type ApplicationCommandPermissions struct {
	ID         Snowflake `json:"id,omitempty"`
	Type       Flag      `json:"type,omitempty"`
	Permission bool      `json:"permission,omitempty"`
}

// Application Command Permission Type
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permission-type
const (
	FlagTypePermissionCommandApplicationROLE = 1
	FlagTypePermissionCommandApplicationUSER = 2
)

// Component Object
type Component interface {
	Type()
}

// Component Types
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
const (
	FlagTypesComponentActionRow  = 1
	FlagTypesComponentButton     = 2
	FlagTypesComponentSelectMenu = 3
	FlagTypesComponentTextInput  = 4
)

func (c ActionsRow) Type() Flag {
	return FlagTypesComponentActionRow
}

func (c Button) Type() Flag {
	return FlagTypesComponentButton
}

func (c SelectMenu) Type() Flag {
	return FlagTypesComponentSelectMenu
}

func (c TextInput) Type() Flag {
	return FlagTypesComponentTextInput
}

// https://discord.com/developers/docs/interactions/message-components#component-object
type ActionsRow struct {
	Components []Component `json:"components,omitempty"`
}

// Button Object
// https://discord.com/developers/docs/interactions/message-components#button-object
type Button struct {
	Style    Flag    `json:"style,omitempty"`
	Label    *string `json:"label,omitempty"`
	Emoji    *Emoji  `json:"emoji,omitempty"`
	CustomID string  `json:"custom_id,omitempty"`
	URL      string  `json:"url,omitempty"`
	Disabled bool    `json:"disabled,omitempty"`
}

// Button Styles
// https://discord.com/developers/docs/interactions/message-components#button-object-button-styles
const (
	FlagStylesbuttonPRIMARY   = 1
	FlagStylesbuttonBLURPLE   = 1
	FlagStylesbuttonSecondary = 2
	FlagStylesbuttonGREY      = 2
	FlagStylesbuttonSuccess   = 3
	FlagStylesbuttonGREEN     = 3
	FlagStylesbuttonDanger    = 4
	FlagStylesbuttonRED       = 4
	FlagStylesbuttonLINK      = 5
)

// Select Menu Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure
type SelectMenu struct {
	CustomID    string             `json:"custom_id,omitempty"`
	Options     []SelectMenuOption `json:"options,omitempty"`
	Placeholder string             `json:"placeholder,omitempty"`
	MinValues   *Flag              `json:"min_values,omitempty"`
	MaxValues   Flag               `json:"max_values,omitempty"`
	Disabled    bool               `json:"disabled,omitempty"`
}

// Select Menu Option Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectMenuOption struct {
	Label       *string `json:"label,omitempty"`
	Value       *string `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
	Emoji       Emoji   `json:"emoji,omitempty"`
	Default     bool    `json:"default,omitempty"`
}

// Text Input Structure
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInput struct {
	CustomID    string    `json:"custom_id,omitempty"`
	Style       Flag      `json:"style,omitempty"`
	Label       *string   `json:"label,omitempty"`
	MinLength   *CodeFlag `json:"min_length,omitempty"`
	MaxLength   CodeFlag  `json:"max_length,omitempty"`
	Required    bool      `json:"required,omitempty"`
	Value       string    `json:"value,omitempty"`
	Placeholder *string   `json:"placeholder,omitempty"`
}

// TextInputStyle
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-styles
const (
	FlagStyleInputTextShort     = 1
	FlagStyleInputTextParagraph = 2
)

// Interaction Object
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	ID            Snowflake       `json:"id,omitempty"`
	ApplicationID Snowflake       `json:"application_id,omitempty"`
	Type          Flag            `json:"type,omitempty"`
	Data          InteractionData `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id,omitempty"`
	ChannelID     Snowflake       `json:"channel_id,omitempty"`
	Member        *GuildMember    `json:"member,omitempty"`
	User          *User           `json:"user,omitempty"`
	Token         string          `json:"token,omitempty"`
	Version       Flag            `json:"version,omitempty"`
	Message       *Message        `json:"message,omitempty"`
	Locale        string          `json:"locale,omitempty"`
	GuildLocale   string          `json:"guild_locale,omitempty"`
}

// Interaction Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
const (
	FlagTypeInteractionPING                             = 1
	FlagTypeInteractionAPPLICATION_COMMAND              = 2
	FlagTypeInteractionMESSAGE_COMPONENT                = 3
	FlagTypeInteractionAPPLICATION_COMMAND_AUTOCOMPLETE = 4
	FlagTypeInteractionMODAL_SUBMIT                     = 5
)

// Interaction Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data-structure
type InteractionData struct {
	ID            Snowflake                                  `json:"id,omitempty"`
	Name          string                                     `json:"name,omitempty"`
	Type          Flag                                       `json:"type,omitempty"`
	Resolved      *ResolvedData                              `json:"resolved,omitempty"`
	Options       []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildID       Snowflake                                  `json:"guild_id,omitempty"`
	CustomID      string                                     `json:"custom_id,omitempty"`
	ComponentType Flag                                       `json:"component_type,omitempty"`
	Values        []*string                                  `json:"values,omitempty"`
	TargetID      Snowflake                                  `json:"target_id,omitempty"`
	Components    []*Component                               `json:"components,omitempty"`
}

// Resolved Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type ResolvedData struct {
	Users       map[Snowflake]*User        `json:"users,omitempty"`
	Members     map[Snowflake]*GuildMember `json:"members,omitempty"`
	Roles       map[Snowflake]*Role        `json:"roles,omitempty"`
	Channels    map[Snowflake]*Channel     `json:"channels,omitempty"`
	Messages    map[Snowflake]*Message     `json:"messages,omitempty"`
	Attachments map[Snowflake]*Attachment  `json:"attachments,omitempty"`
}

// Message Interaction Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	ID     Snowflake    `json:"id,omitempty"`
	Type   Flag         `json:"type,omitempty"`
	Name   string       `json:"name,omitempty"`
	User   *User        `json:"user,omitempty"`
	Member *GuildMember `json:"member,omitempty"`
}

// Interaction Response Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure
type InteractionResponse struct {
	Type Flag                     `json:"type,omitempty"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// Interaction Callback Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
const (
	FlagTypeCallbackInteractionPONG                                    = 1
	FlagTypeCallbackInteractionCHANNEL_MESSAGE_WITH_SOURCE             = 4
	FlagTypeCallbackInteractionDEFERRED_CHANNEL_MESSAGE_WITH_SOURCE    = 5
	FlagTypeCallbackInteractionDEFERRED_UPDATE_MESSAGE                 = 6
	FlagTypeCallbackInteractionUPDATE_MESSAGE                          = 7
	FlagTypeCallbackInteractionAPPLICATION_COMMAND_AUTOCOMPLETE_RESULT = 8
	FlagTypeCallbackInteractionMODAL                                   = 9
)

// Interaction Callback Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionCallbackData interface{}

// Messages
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-messages
type Messages struct {
	TTS             bool             `json:"tts,omitempty"`
	Content         string           `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           BitFlag          `json:"flags,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// Autocomplete
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-autocomplete
type Autocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
}

// Modal
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-modal
type ModalSubmitInteractionData struct {
	CustomID   *string     `json:"custom_id,omitempty"`
	Title      string      `json:"title,omitempty"`
	Components []Component `json:"components,omitempty"`
}

// Application Object
// https://discord.com/developers/docs/resources/application
type Application struct {
	ID                  Snowflake      `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	Icon                string         `json:"icon,omitempty"`
	Description         string         `json:"description,omitempty"`
	RPCOrigins          []string       `json:"rpc_origins,omitempty"`
	BotPublic           bool           `json:"bot_public,omitempty"`
	BotRequireCodeGrant bool           `json:"bot_require_code_grant,omitempty"`
	TermsOfServiceURL   string         `json:"terms_of_service_url,omitempty"`
	PrivacyProxyURL     string         `json:"privacy_policy_url,omitempty"`
	Owner               *User          `json:"owner,omitempty"`
	VerifyKey           string         `json:"verify_key,omitempty"`
	Team                *Team          `json:"team,omitempty"`
	GuildID             Snowflake      `json:"guild_id,omitempty"`
	PrimarySKUID        Snowflake      `json:"primary_sku_id,omitempty"`
	Slug                *string        `json:"slug,omitempty"`
	CoverImage          string         `json:"cover_image,omitempty"`
	Flags               Flag           `json:"flags,omitempty"`
	Tags                []string       `json:"tags,omitempty"`
	InstallParams       *InstallParams `json:"install_params,omitempty"`
	CustomInstallURL    string         `json:"custom_install_url,omitempty"`
}

// Application Flags
// https://discord.com/developers/docs/resources/application#application-object-application-flags
const (
	FlagFlagsApplicationGATEWAY_PRESENCE                 = 1 << 12
	FlagFlagsApplicationGATEWAY_PRESENCE_LIMITED         = 1 << 13
	FlagFlagsApplicationGATEWAY_GUILD_MEMBERS            = 1 << 14
	FlagFlagsApplicationGATEWAY_GUILD_MEMBERS_LIMITED    = 1 << 15
	FlagFlagsApplicationVERIFICATION_PENDING_GUILD_LIMIT = 1 << 16
	FlagFlagsApplicationEMBEDDED                         = 1 << 17
	FlagFlagsApplicationGATEWAY_MESSAGE_CONTENT          = 1 << 18
	FlagFlagsApplicationGATEWAY_MESSAGE_CONTENT_LIMITED  = 1 << 19
)

// Install Params Object
// https://discord.com/developers/docs/resources/application#install-params-object
type InstallParams struct {
	Scopes      []string `json:"scopes,omitempty"`
	Permissions string   `json:"permissions,omitempty"`
}

// Audit Log Object
// https://discord.com/developers/docs/resources/audit-log
type AuditLog struct {
	AuditLogEntries      []*AuditLogEntry       `json:"audit_log_entries,omitempty"`
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events,omitempty"`
	Integration          []*Integration         `json:"integrations,omitempty"`
	Threads              []*Channel             `json:"threads,omitempty"`
	Users                []*User                `json:"users,omitempty"`
	Webhooks             []*Webhook             `json:"webhooks,omitempty"`
}

// Audit Log Entry Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type AuditLogEntry struct {
	TargetID   string            `json:"target_id,omitempty"`
	Changes    []*AuditLogChange `json:"changes,omitempty"`
	UserID     Snowflake         `json:"user_id,omitempty"`
	ID         Snowflake         `json:"id,omitempty"`
	ActionType Flag              `json:"action_type,omitempty"`
	Options    *AuditLogOptions  `json:"options,omitempty"`
	Reason     *string           `json:"reason,omitempty"`
}

// Audit Log Events
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
const (
	FlagEventsLogAuditGUILD_UPDATE                          = 1
	FlagEventsLogAuditCHANNEL_CREATE                        = 10
	FlagEventsLogAuditCHANNEL_UPDATE                        = 11
	FlagEventsLogAuditCHANNEL_DELETE                        = 12
	FlagEventsLogAuditCHANNEL_OVERWRITE_CREATE              = 13
	FlagEventsLogAuditCHANNEL_OVERWRITE_UPDATE              = 14
	FlagEventsLogAuditCHANNEL_OVERWRITE_DELETE              = 15
	FlagEventsLogAuditMEMBER_KICK                           = 20
	FlagEventsLogAuditMEMBER_PRUNE                          = 21
	FlagEventsLogAuditMEMBER_BAN_ADD                        = 22
	FlagEventsLogAuditMEMBER_BAN_REMOVE                     = 23
	FlagEventsLogAuditMEMBER_UPDATE                         = 24
	FlagEventsLogAuditMEMBER_ROLE_UPDATE                    = 25
	FlagEventsLogAuditMEMBER_MOVE                           = 26
	FlagEventsLogAuditMEMBER_DISCONNECT                     = 27
	FlagEventsLogAuditBOT_ADD                               = 28
	FlagEventsLogAuditROLE_CREATE                           = 30
	FlagEventsLogAuditROLE_UPDATE                           = 31
	FlagEventsLogAuditROLE_DELETE                           = 32
	FlagEventsLogAuditINVITE_CREATE                         = 40
	FlagEventsLogAuditINVITE_UPDATE                         = 41
	FlagEventsLogAuditINVITE_DELETE                         = 42
	FlagEventsLogAuditWEBHOOK_CREATE                        = 50
	FlagEventsLogAuditWEBHOOK_UPDATE                        = 51
	FlagEventsLogAuditWEBHOOK_DELETE                        = 52
	FlagEventsLogAuditEMOJI_CREATE                          = 60
	FlagEventsLogAuditEMOJI_UPDATE                          = 61
	FlagEventsLogAuditEMOJI_DELETE                          = 62
	FlagEventsLogAuditMESSAGE_DELETE                        = 72
	FlagEventsLogAuditMESSAGE_BULK_DELETE                   = 73
	FlagEventsLogAuditMESSAGE_PIN                           = 74
	FlagEventsLogAuditMESSAGE_UNPIN                         = 75
	FlagEventsLogAuditINTEGRATION_CREATE                    = 80
	FlagEventsLogAuditINTEGRATION_UPDATE                    = 81
	FlagEventsLogAuditINTEGRATION_DELETE                    = 82
	FlagEventsLogAuditSTAGE_INSTANCE_CREATE                 = 83
	FlagEventsLogAuditSTAGE_INSTANCE_UPDATE                 = 84
	FlagEventsLogAuditSTAGE_INSTANCE_DELETE                 = 85
	FlagEventsLogAuditSTICKER_CREATE                        = 90
	FlagEventsLogAuditSTICKER_UPDATE                        = 91
	FlagEventsLogAuditSTICKER_DELETE                        = 92
	FlagEventsLogAuditGUILD_SCHEDULED_EVENT_CREATE          = 100
	FlagEventsLogAuditGUILD_SCHEDULED_EVENT_UPDATE          = 101
	FlagEventsLogAuditGUILD_SCHEDULED_EVENT_DELETE          = 102
	FlagEventsLogAuditTHREAD_CREATE                         = 110
	FlagEventsLogAuditTHREAD_UPDATE                         = 111
	FlagEventsLogAuditTHREAD_DELETE                         = 112
	FlagEventsLogAuditAPPLICATION_COMMAND_PERMISSION_UPDATE = 121
)

// Optional Audit Entry Info
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogOptions struct {
	ApplicationID    Snowflake `json:"application_id,omitempty"`
	ChannelID        Snowflake `json:"channel_id,omitempty"`
	Count            string    `json:"count,omitempty"`
	DeleteMemberDays string    `json:"delete_member_days,omitempty"`
	ID               Snowflake `json:"id,omitempty"`
	MembersRemoved   string    `json:"members_removed,omitempty"`
	MessageID        Snowflake `json:"message_id,omitempty"`
	RoleName         string    `json:"role_name,omitempty"`
	Type             string    `json:"type,omitempty"`
}

// Audit Log Change Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object
type AuditLogChange struct {
	NewValue interface{} `json:"new_value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
	Key      string      `json:"key,omitempty"`
}

// Audit Log Change Exceptions
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-exceptions

// Channel Object
// https://discord.com/developers/docs/resources/channel
type Channel struct {
	ID                         Snowflake             `json:"id,omitempty"`
	Type                       *Flag                 `json:"type,omitempty"`
	GuildID                    Snowflake             `json:"guild_id,omitempty"`
	Position                   int                   `json:"position,omitempty"`
	PermissionOverwrites       []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                       string                `json:"name,omitempty"`
	Topic                      *string               `json:"topic,omitempty"`
	NSFW                       bool                  `json:"nsfw,omitempty"`
	LastMessageID              Snowflake             `json:"last_message_id,omitempty"`
	Bitrate                    Flag                  `json:"bitrate,omitempty"`
	UserLimit                  Flag                  `json:"user_limit,omitempty"`
	RateLimitPerUser           *CodeFlag             `json:"rate_limit_per_user,omitempty"`
	Recipients                 []*User               `json:"recipients,omitempty"`
	Icon                       string                `json:"icon,omitempty"`
	OwnerID                    Snowflake             `json:"owner_id,omitempty"`
	ApplicationID              Snowflake             `json:"application_id,omitempty"`
	ParentID                   Snowflake             `json:"parent_id,omitempty"`
	LastPinTimestamp           time.Time             `json:"last_pin_timestamp,omitempty"`
	RTCRegion                  string                `json:"rtc_region,omitempty"`
	VideoQualityMode           Flag                  `json:"video_quality_mode,omitempty"`
	MessageCount               Flag                  `json:"message_count,omitempty"`
	MemberCount                Flag                  `json:"member_count,omitempty"`
	ThreadMetadata             *ThreadMetadata       `json:"thread_metadata,omitempty"`
	Member                     *ThreadMember         `json:"member,omitempty"`
	DefaultAutoArchiveDuration int                   `json:"default_auto_archive_duration,omitempty"`
	Permissions                *string               `json:"permissions,omitempty"`
	Flags                      BitFlag               `json:"flags,omitempty"`
}

// Channel Types
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
const (
	FlagTypesChannelGUILD_TEXT           = 0
	FlagTypesChannelDM                   = 1
	FlagTypesChannelGUILD_VOICE          = 2
	FlagTypesChannelGROUP_DM             = 3
	FlagTypesChannelGUILD_CATEGORY       = 4
	FlagTypesChannelGUILD_NEWS           = 5
	FlagTypesChannelGUILD_NEWS_THREAD    = 10
	FlagTypesChannelGUILD_PUBLIC_THREAD  = 11
	FlagTypesChannelGUILD_PRIVATE_THREAD = 12
	FlagTypesChannelGUILD_STAGE_VOICE    = 13
	FlagTypesChannelGUILD_DIRECTORY      = 14
	FlagTypesChannelGUILD_FORUM          = 15
)

// Video Quality Modes
// https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
const (
	FlagModesQualityVideoAUTO = 1
	FlagModesQualityVideoFULL = 2
)

// Channel Flags
// https://discord.com/developers/docs/resources/channel#channel-object-channel-flags
const (
	FlagChannelPINNED = 1 << 1
)

// Message Object
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                Snowflake         `json:"id,omitempty"`
	ChannelID         *Snowflake        `json:"channel_id,omitempty"`
	GuildID           *Snowflake        `json:"guild_id,omitempty"`
	Author            *User             `json:"author,omitempty"`
	Member            *GuildMember      `json:"member,omitempty"`
	Content           string            `json:"content,omitempty"`
	Timestamp         time.Time         `json:"timestamp,omitempty"`
	EditedTimestamp   time.Time         `json:"edited_timestamp,omitempty"`
	TTS               bool              `json:"tts,omitempty"`
	MentionEveryone   bool              `json:"mention_everyone,omitempty"`
	Mentions          []*User           `json:"mentions,omitempty"`
	MentionRoles      []*Snowflake      `json:"mention_roles,omitempty"`
	MentionChannels   []*ChannelMention `json:"mention_channels,omitempty"`
	Attachments       []*Attachment     `json:"attachments,omitempty"`
	Embeds            []*Embed          `json:"embeds,omitempty"`
	Reactions         []*Reaction       `json:"reactions,omitempty"`
	Nonce             interface{}       `json:"nonce,omitempty"`
	Pinned            bool              `json:"pinned,omitempty"`
	WebhookID         *Snowflake        `json:"webhook_id,omitempty"`
	Type              *Flag             `json:"type,omitempty"`
	Activity          MessageActivity   `json:"activity,omitempty"`
	Application       *Application      `json:"application,omitempty"`
	ApplicationID     Snowflake         `json:"application_id,omitempty"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	Flags             CodeFlag          `json:"flags,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	Interaction       *Interaction      `json:"interaction,omitempty"`
	Thread            *Channel          `json:"thread,omitempty"`
	Components        []*Component      `json:"components,omitempty"`
	StickerItems      []*StickerItem    `json:"sticker_items,omitempty"`
}

// Message Types
// https://discord.com/developers/docs/resources/channel#message-object-message-types
const (
	FlagTypesMessageDEFAULT                                      = 0
	FlagTypesMessageRECIPIENT_ADD                                = 1
	FlagTypesMessageRECIPIENT_REMOVE                             = 2
	FlagTypesMessageCALL                                         = 3
	FlagTypesMessageCHANNEL_NAME_CHANGE                          = 4
	FlagTypesMessageCHANNEL_ICON_CHANGE                          = 5
	FlagTypesMessageCHANNEL_PINNED_MESSAGE                       = 6
	FlagTypesMessageGUILD_MEMBER_JOIN                            = 7
	FlagTypesMessageUSER_PREMIUM_GUILD_SUBSCRIPTION              = 8
	FlagTypesMessageUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_ONE     = 9
	FlagTypesMessageUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_TWO     = 10
	FlagTypesMessageUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_THREE   = 11
	FlagTypesMessageCHANNEL_FOLLOW_ADD                           = 12
	FlagTypesMessageGUILD_DISCOVERY_DISQUALIFIED                 = 14
	FlagTypesMessageGUILD_DISCOVERY_REQUALIFIED                  = 15
	FlagTypesMessageGUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING = 16
	FlagTypesMessageGUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING   = 17
	FlagTypesMessageTHREAD_CREATED                               = 18
	FlagTypesMessageREPLY                                        = 19
	FlagTypesMessageCHAT_INPUT_COMMAND                           = 20
	FlagTypesMessageTHREAD_STARTER_MESSAGE                       = 21
	FlagTypesMessageGUILD_INVITE_REMINDER                        = 22
	FlagTypesMessageCONTEXT_MENU_COMMAND                         = 23
)

// Message Activity Structure
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	Type    int     `json:"type,omitempty"`
	PartyID *string `json:"party_id,omitempty"`
}

// Message Activity Types
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
const (
	FlagTypesActivityMessageJOIN         = 1
	FlagTypesActivityMessageSPECTATE     = 2
	FlagTypesActivityMessageLISTEN       = 3
	FlagTypesActivityMessageJOIN_REQUEST = 5
)

// Message Flags
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
const (
	FlagFlagsMessageCROSSPOSTED                            = 1 << 0
	FlagFlagsMessageIS_CROSSPOST                           = 1 << 1
	FlagFlagsMessageSUPPRESS_EMBEDS                        = 1 << 2
	FlagFlagsMessageSOURCE_MESSAGE_DELETED                 = 1 << 3
	FlagFlagsMessageURGENT                                 = 1 << 4
	FlagFlagsMessageHAS_THREAD                             = 1 << 5
	FlagFlagsMessageEPHEMERAL                              = 1 << 6
	FlagFlagsMessageLOADING                                = 1 << 7
	FlagFlagsMessageFAILED_TO_MENTION_SOME_ROLES_IN_THREAD = 1 << 8
)

// Message Reference Object
// https://discord.com/developers/docs/resources/channel#message-reference-object
type MessageReference struct {
	MessageID       Snowflake  `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool       `json:"fail_if_not_exists,omitempty"`
}

// Followed Channel Structure
// https://discord.com/developers/docs/resources/channel#followed-channel-object-followed-channel-structure
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id,omitempty"`
	WebhookID Snowflake `json:"webhook_id,omitempty"`
}

// Reaction Object
// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count CodeFlag `json:"count,omitempty"`
	Me    bool     `json:"me,omitempty"`
	Emoji *Emoji   `json:"emoji,omitempty"`
}

// Overwrite Object
// https://discord.com/developers/docs/resources/channel#overwrite-object
type PermissionOverwrite struct {
	ID    Snowflake `json:"id,omitempty"`
	Type  *Flag     `json:"type,omitempty"`
	Deny  string    `json:"deny,omitempty"`
	Allow string    `json:"allow,omitempty"`
}

// Thread Metadata Object
// https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool      `json:"archived,omitempty"`
	AutoArchiveDuration int       `json:"auto_archive_duration,omitempty"`
	Locked              bool      `json:"locked,omitempty"`
	Invitable           bool      `json:"invitable,omitempty"`
	CreateTimestamp     time.Time `json:"create_timestamp,omitempty"`
}

// Thread Member Object
// https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	ThreadID      Snowflake `json:"id,omitempty"`
	UserID        Snowflake `json:"user_id,omitempty"`
	JoinTimestamp time.Time `json:"join_timestamp,omitempty"`
	Flags         CodeFlag  `json:"flags,omitempty"`
}

// Embed Object
// https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description *string         `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   time.Time       `json:"timestamp,omitempty"`
	Color       CodeFlag        `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

// Embed Thumbnail Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	URL      string  `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   int     `json:"height,omitempty"`
	Width    int     `json:"width,omitempty"`
}

// Embed Video Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	URL      string  `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   int     `json:"height,omitempty"`
	Width    int     `json:"width,omitempty"`
}

// Embed Image Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	URL      string  `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   int     `json:"height,omitempty"`
	Width    int     `json:"width,omitempty"`
}

// Embed Provider Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Embed Author Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// Embed Footer Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         *string `json:"text,omitempty"`
	IconURL      string  `json:"icon_url,omitempty"`
	ProxyIconURL string  `json:"proxy_icon_url,omitempty"`
}

// Embed Field Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

// Embed Limits
// https://discord.com/developers/docs/resources/channel#embed-object-embed-limits
const (
	FlagLimitsEmbedTitle            = 256
	FlagLimitsEmbedDescription      = 4096
	FlagLimitsEmbedEmbedLimitFields = 25
	FlagLimitsEmbedFieldName        = 256
	FlagLimitsEmbedFieldValue       = 1024
	FlagLimitsEmbedFooterText       = 2048
	FlagLimitsEmbedAuthorName       = 256
)

// Message Attachment Object
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID          Snowflake `json:"id,omitempty"`
	Filename    string    `json:"filename,omitempty"`
	Description string    `json:"description,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Size        int       `json:"size,omitempty"`
	URL         string    `json:"url,omitempty"`
	ProxyURL    *string   `json:"proxy_url,omitempty"`
	Height      int       `json:"height,omitempty"`
	Width       int       `json:"width,omitempty"`
	Emphemeral  bool      `json:"ephemeral,omitempty"`
}

// Channel Mention Object
// https://discord.com/developers/docs/resources/channel#channel-mention-object
type ChannelMention struct {
	ID      Snowflake `json:"id,omitempty"`
	GuildID Snowflake `json:"guild_id,omitempty"`
	Type    *Flag     `json:"type,omitempty"`
	Name    string    `json:"name,omitempty"`
}

// Allowed Mentions Structure
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMentions struct {
	Parse       []*string    `json:"parse,omitempty"`
	Roles       []*Snowflake `json:"roles,omitempty"`
	Users       []*Snowflake `json:"users,omitempty"`
	RepliedUser bool         `json:"replied_user,omitempty"`
}

// Allowed Mention Types
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
const (
	FlagTypesMentionAllowedRoles     = "roles"
	FlagTypesMentionAllowedsUsers    = "users"
	FlagTypesMentionAllowedsEveryone = "everyone"
)

// Emoji Object
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	ID            Snowflake   `json:"id,omitempty"`
	Name          *string     `json:"name,omitempty"`
	Roles         []Snowflake `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons bool        `json:"require_colons,omitempty"`
	Managed       bool        `json:"managed,omitempty"`
	Animated      bool        `json:"animated,omitempty"`
	Available     bool        `json:"available,omitempty"`
}

// Guild Object
// https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	ID                          Snowflake      `json:"id,omitempty"`
	Name                        string         `json:"name,omitempty"`
	Icon                        string         `json:"icon,omitempty"`
	IconHash                    string         `json:"icon_hash,omitempty"`
	Splash                      string         `json:"splash,omitempty"`
	DiscoverySplash             string         `json:"discovery_splash,omitempty"`
	Owner                       bool           `json:"owner,omitempty"`
	OwnerID                     Snowflake      `json:"owner_id,omitempty"`
	Permissions                 *string        `json:"permissions,omitempty"`
	Region                      string         `json:"region,omitempty"`
	AfkChannelID                Snowflake      `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int            `json:"afk_timeout,omitempty"`
	WidgetEnabled               bool           `json:"widget_enabled,omitempty"`
	WidgetChannelID             Snowflake      `json:"widget_channel_id,omitempty"`
	VerificationLevel           *Flag          `json:"verification_level,omitempty"`
	DefaultMessageNotifications *Flag          `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *Flag          `json:"explicit_content_filter,omitempty"`
	Roles                       []*Role        `json:"roles,omitempty"`
	Emojis                      []*Emoji       `json:"emojis,omitempty"`
	Features                    []*string      `json:"features,omitempty"`
	MFALevel                    *Flag          `json:"mfa_level,omitempty"`
	ApplicationID               Snowflake      `json:"application_id,omitempty"`
	SystemChannelID             Snowflake      `json:"system_channel_id,omitempty"`
	SystemChannelFlags          BitFlag        `json:"system_channel_flags,omitempty"`
	RulesChannelID              Snowflake      `json:"rules_channel_id,omitempty"`
	MaxPresences                CodeFlag       `json:"max_presences,omitempty"`
	MaxMembers                  int            `json:"max_members,omitempty"`
	VanityUrl                   *string        `json:"vanity_url_code,omitempty"`
	Description                 *string        `json:"description,omitempty"`
	Banner                      string         `json:"banner,omitempty"`
	PremiumTier                 *Flag          `json:"premium_tier,omitempty"`
	PremiumSubscriptionCount    *CodeFlag      `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale,omitempty"`
	PublicUpdatesChannelID      Snowflake      `json:"public_updates_channel_id,omitempty"`
	MaxVideoChannelUsers        int            `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      int            `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int            `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NSFWLevel                   *Flag          `json:"nsfw_level,omitempty"`
	Stickers                    []*Sticker     `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled,omitempty"`

	// Unavailable Guild Object
	// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Unavailable bool `json:"unavailable,omitempty"`
}

// Default Message Notification Level
// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
const (
	FlagLevelNotificationMessageDefaultALL_MESSAGES  = 0
	FlagLevelNotificationMessageDefaultONLY_MENTIONS = 1
)

// Explicit Content Filter Level
// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
const (
	FlagLevelFilterContentExplicitDISABLED              = 0
	FlagLevelFilterContentExplicitMEMBERS_WITHOUT_ROLES = 1
	FlagLevelFilterContentExplicitALL_MEMBERS           = 2
)

// MFA Level
// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
const (
	FlagLevelMFANONE     = 0
	FlagLevelMFAELEVATED = 1
)

// Verification Level
// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
const (
	FlagLevelVerificationNONE      = 0
	FlagLevelVerificationLOW       = 1
	FlagLevelVerificationMEDIUM    = 2
	FlagLevelVerificationHIGH      = 3
	FlagLevelVerificationVERY_HIGH = 4
)

// Guild NSFW Level
// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
const (
	FlagLevelNSFWGuildDEFAULT        = 0
	FlagLevelNSFWGuildEXPLICIT       = 1
	FlagLevelNSFWGuildSAFE           = 2
	FlagLevelNSFWGuildAGE_RESTRICTED = 3
)

// Premium Tier
// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
const (
	FlagTierPremiumNONE  = 0
	FlagTierPremiumONE   = 1
	FlagTierPremiumTWO   = 2
	FlagTierPremiumTHREE = 3
)

// System Channel Flags
// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
const (
	FlagFlagsChannelSystemSUPPRESS_JOIN_NOTIFICATIONS           = 1 << 0
	FlagFlagsChannelSystemSUPPRESS_PREMIUM_SUBSCRIPTIONS        = 1 << 1
	FlagFlagsChannelSystemSUPPRESS_GUILD_REMINDER_NOTIFICATIONS = 1 << 2
	FlagFlagsChannelSystemSUPPRESS_JOIN_NOTIFICATION_REPLIES    = 1 << 3
)

// Guild Features
// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
var (
	GuildFeatures = map[string]string{
		"ANIMATED_BANNER":                  "guild has access to set an animated guild banner image",
		"ANIMATED_ICON":                    "guild has access to set an animated guild icon",
		"BANNER":                           "guild has access to set a guild banner image",
		"COMMERCE":                         "guild has access to use commerce features (i.e. create store channels)",
		"COMMUNITY":                        "guild can enable welcome screen, Membership Screening, stage channels and discovery, and receives community updates",
		"DISCOVERABLE":                     "guild is able to be discovered in the directory",
		"FEATURABLE":                       "guild is able to be featured in the directory",
		"INVITE_SPLASH":                    "guild has access to set an invite splash background",
		"MEMBER_VERIFICATION_GATE_ENABLED": "guild has enabled Membership Screening",
		"MONETIZATION_ENABLED":             "guild has enabled monetization",
		"MORE_STICKERS":                    "guild has increased custom sticker slots",
		"NEWS":                             "guild has access to create news channels",
		"PARTNERED":                        "guild is partnered",
		"PREVIEW_ENABLED":                  "guild can be previewed before joining via Membership Screening or the directory",
		"PRIVATE_THREADS":                  "guild has access to create private threads",
		"ROLE_ICONS":                       "guild is able to set role icons",
		"SEVEN_DAY_THREAD_ARCHIVE":         "guild has access to the seven day archive time for threads",
		"THREE_DAY_THREAD_ARCHIVE":         "guild has access to the three day archive time for threads",
		"TICKETED_EVENTS_ENABLED":          "guild has enabled ticketed events",
		"VANITY_URL":                       "guild has access to set a vanity URL",
		"VERIFIED":                         "guild is verified",
		"VIP_REGIONS":                      "guild has access to set 384kbps bitrate in voice (previously VIP voice servers)",
		"WELCOME_SCREEN_ENABLED":           "guild has enabled the welcome screen",
	}
)

// Guild Preview Object
// https://discord.com/developers/docs/resources/guild#guild-preview-object-guild-preview-structure
type GuildPreview struct {
	ID                       string     `json:"id,omitempty"`
	Name                     string     `json:"name,omitempty"`
	Icon                     string     `json:"icon,omitempty"`
	Splash                   string     `json:"splash,omitempty"`
	DiscoverySplash          string     `json:"discovery_splash,omitempty"`
	Emojis                   []*Emoji   `json:"emojis,omitempty"`
	Features                 []*string  `json:"features,omitempty"`
	ApproximateMemberCount   int        `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount int        `json:"approximate_presence_count,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Stickers                 []*Sticker `json:"stickers,omitempty"`
}

// Guild Widget Settings Object
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object
type GuildWidgetSettings struct {
	Enabled   bool      `json:"enabled,omitempty"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
}

// Guild Widget Object
// https://discord.com/developers/docs/resources/guild#et-gguild-widget-object-get-guild-widget-structure*
type GuildWidget struct {
	ID            Snowflake  `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	InstantInvite string     `json:"instant_invite,omitempty"`
	Channels      []*Channel `json:"channels,omitempty"`
	Members       []*User    `json:"members,omitempty"`
	PresenceCount int        `json:"presence_count,omitempty"`
}

// Guild Member Object
// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User                       *User        `json:"user,omitempty"`
	Nick                       *string      `json:"nick,omitempty"`
	Avatar                     string       `json:"avatar,omitempty"`
	Roles                      []*Snowflake `json:"roles,omitempty"`
	GuildID                    Snowflake    `json:"guild_id,omitempty"`
	JoinedAt                   time.Time    `json:"joined_at,omitempty"`
	PremiumSince               time.Time    `json:"premium_since,omitempty"`
	Deaf                       bool         `json:"deaf,omitempty"`
	Mute                       bool         `json:"mute,omitempty"`
	Pending                    bool         `json:"pending,omitempty"`
	Permissions                *string      `json:"permissions,omitempty"`
	CommunicationDisabledUntil time.Time    `json:"communication_disabled_until,omitempty"`
}

// Integration Object
// https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	ID                Snowflake          `json:"id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Type              string             `json:"type,omitempty"`
	Enabled           bool               `json:"enabled,omitempty"`
	Syncing           bool               `json:"syncing,omitempty"`
	RoleID            Snowflake          `json:"role_id,omitempty"`
	EnableEmoticons   bool               `json:"enable_emoticons,omitempty"`
	ExpireBehavior    *Flag              `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *int               `json:"expire_grace_period,omitempty"`
	User              *User              `json:"user,omitempty"`
	Account           IntegrationAccount `json:"account,omitempty"`
	SyncedAt          time.Time          `json:"synced_at,omitempty"`
	SubscriberCount   *int               `json:"subscriber_count,omitempty"`
	Revoked           bool               `json:"revoked,omitempty"`
	Application       *Application       `json:"application,omitempty"`
}

// Integration Expire Behaviors
// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
const (
	FlagBehaviorsExpireIntegrationREMOVEROLE = 0
	FlagBehaviorsExpireIntegrationKICK       = 1
)

// Integration Account Object
// https://discord.com/developers/docs/resources/guild#integration-account-object
type IntegrationAccount struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Integration Application Object
// https://discord.com/developers/docs/resources/guild#integration-application-object-integration-application-structure
type IntegrationApplication struct {
	ID          Snowflake `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	Description *string   `json:"description,omitempty"`
	Bot         *User     `json:"bot,omitempty"`
}

// Guild Ban Object
// https://discord.com/developers/docs/resources/guild#ban-object
type Ban struct {
	Reason *string `json:"reason,omitempty"`
	User   *User   `json:"user,omitempty"`
}

// Welcome Screen Object
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-structure
type WelcomeScreen struct {
	Description           *string                 `json:"description,omitempty"`
	WelcomeScreenChannels []*WelcomeScreenChannel `json:"welcome_channels,omitempty"`
}

// Welcome Screen Channel Structure
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure
type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id,omitempty"`
	Description *string    `json:"description,omitempty"`
	EmojiID     *Snowflake `json:"emoji_id,omitempty"`
	EmojiName   *string    `json:"emoji_name,omitempty"`
}

// Guild Scheduled Event Object
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	ID                 Snowflake                         `json:"id,omitempty"`
	GuildID            Snowflake                         `json:"guild_id,omitempty"`
	ChannelID          Snowflake                         `json:"channel_id,omitempty"`
	CreatorID          Snowflake                         `json:"creator_id,omitempty"`
	Name               string                            `json:"name,omitempty"`
	Description        string                            `json:"description,omitempty"`
	ScheduledStartTime time.Time                         `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime   time.Time                         `json:"scheduled_end_time,omitempty"`
	PrivacyLevel       Flag                              `json:"privacy_level,omitempty"`
	Status             Flag                              `json:"status,omitempty"`
	EntityType         Flag                              `json:"entity_type,omitempty"`
	EntityID           Snowflake                         `json:"entity_id,omitempty"`
	EntityMetadata     GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Creator            *User                             `json:"creator,omitempty"`
	UserCount          int                               `json:"user_count,omitempty"`
	Image              string                            `json:"image,omitempty"`
}

// Guild Scheduled Event Privacy Level
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-privacy-level
const (
	FlagGuildScheduledEventPrivacyLevelGUILD_ONLY = 2
)

// Guild Scheduled Event Entity Types
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-types
const (
	FlagTypesEntityEventScheduledGuildSTAGE_INSTANCE = 1
	FlagTypesEntityEventScheduledGuildVOICE          = 2
	FlagTypesEntityEventScheduledGuildEXTERNAL       = 3
)

// Guild Scheduled Event Status
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-status
const (
	FlagStatusEventScheduledGuildSCHEDULED = 1
	FlagStatusEventScheduledGuildACTIVE    = 2
	FlagStatusEventScheduledGuildCOMPLETED = 3
	FlagStatusEventScheduledGuildCANCELED  = 4
)

// Guild Scheduled Event Entity Metadata
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata
type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}

// Guild Scheduled Event User Object
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object-guild-scheduled-event-user-structure
type GuildScheduledEventUser struct {
	GuildScheduledEventID Snowflake    `json:"guild_scheduled_event_id,omitempty"`
	User                  *User        `json:"user,omitempty"`
	Member                *GuildMember `json:"member,omitempty"`
}

// Guild Template Object
// https://discord.com/developers/docs/resources/guild-template#guild-template-object
type GuildTemplate struct {
	Code                  string    `json:"code,omitempty"`
	Name                  string    `json:"name,omitempty"`
	Description           *string   `json:"description,omitempty"`
	UsageCount            *int      `json:"usage_count,omitempty"`
	CreatorID             Snowflake `json:"creator_id,omitempty"`
	Creator               *User     `json:"creator,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	SourceGuildID         Snowflake `json:"source_guild_id,omitempty"`
	SerializedSourceGuild *Guild    `json:"serialized_source_guild,omitempty"`
	IsDirty               bool      `json:"is_dirty,omitempty"`
}

// Client Status Object
// https://discord.com/developers/docs/topics/gateway#client-status-object
type ClientStatus struct {
	Desktop *string `json:"desktop,omitempty"`
	Mobile  *string `json:"mobile,omitempty"`
	Web     *string `json:"web,omitempty"`
}

// Activity Object
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-structure
type Activity struct {
	Name          string              `json:"name,omitempty"`
	Type          *Flag               `json:"type,omitempty"`
	URL           string              `json:"url,omitempty"`
	CreatedAt     int                 `json:"created_at,omitempty"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID Snowflake           `json:"application_id,omitempty"`
	Details       string              `json:"details,omitempty"`
	State         string              `json:"state,omitempty"`
	Emoji         *Emoji              `json:"emoji,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      bool                `json:"instance,omitempty"`
	Flags         BitFlag             `json:"flags,omitempty"`
	Buttons       []Button            `json:"buttons,omitempty"`
}

// Activity Types
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
const (
	FlagEnumTypeActivityPlaying   = 0
	FlagEnumTypeActivityStreaming = 1
	FlagEnumTypeActivityListening = 2
	FlagEnumTypeActivityWatching  = 3
	FlagEnumTypeActivityCustom    = 4
	FlagEnumTypeActivityCompeting = 5
)

// Activity Timestamps Struct
// htthttps://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type ActivityTimestamps struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// Activity Emoji
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
type ActivityEmoji struct {
	Name     string    `json:"name,omitempty"`
	ID       Snowflake `json:"id,omitempty"`
	Animated bool      `json:"animated,omitempty"`
}

// Activity Party Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type ActivityParty struct {
	ID   string  `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

// Activity Assets Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

// Activity Asset Image
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
type ActivityAssetImage struct {
	ApplicationAsset string `json:"application_asset_id,omitempty"`
	MediaProxyImage  string `json:"image_id,omitempty"`
}

// Activity Secrets Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

// Activity Flags
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
const (
	FlagActivityINSTANCE                    = 1 << 0
	FlagActivityJOIN                        = 1 << 1
	FlagActivitySPECTATE                    = 1 << 2
	FlagActivityJOIN_REQUEST                = 1 << 3
	FlagActivitySYNC                        = 1 << 4
	FlagActivityPLAY                        = 1 << 5
	FlagActivityPARTY_PRIVACY_FRIENDS       = 1 << 6
	FlagActivityPARTY_PRIVACY_VOICE_CHANNEL = 1 << 7
	FlagActivityEMBEDDED                    = 1 << 8
)

// Invite Object
// https://discord.com/developers/docs/resources/invite#invite-object
type Invite struct {
	Code                     string               `json:"code,omitempty"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel,omitempty"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetType               Flag                 `json:"target_type,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount int                  `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int                  `json:"approximate_member_count,omitempty"`
	ExpiresAt                time.Time            `json:"expires_at,omitempty"`
	StageInstance            StageInstance        `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// Invite Target Types
// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
const (
	FlagTypesTargetInviteSTREAM               = 1
	FlagTypesTargetInviteEMBEDDED_APPLICATION = 2
)

// Invite Metadata Object
// https://discord.com/developers/docs/resources/invite#invite-metadata-object-invite-metadata-structure
type InviteMetadata struct {
	Uses      *int      `json:"uses,omitempty"`
	MaxUses   *int      `json:"max_uses,omitempty"`
	MaxAge    int       `json:"max_age,omitempty"`
	Temporary bool      `json:"temporary,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Sticker Structure
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	ID          Snowflake  `json:"id,omitempty"`
	PackID      Snowflake  `json:"pack_id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Tags        *string    `json:"tags,omitempty"`
	Asset       *string    `json:"asset,omitempty"`
	Type        Flag       `json:"type,omitempty"`
	FormatType  Flag       `json:"format_type,omitempty"`
	Available   bool       `json:"available,omitempty"`
	GuildID     *Snowflake `json:"guild_id,omitempty"`
	User        *User      `json:"user,omitempty"`
	SortValue   int        `json:"sort_value,omitempty"`
}

// Sticker Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
const (
	FlagTypesStickerSTANDARD = 1
	FlagTypesStickerGUILD    = 2
)

// Sticker Format Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
const (
	FlagTypesFormatStickerPNG    = 1
	FlagTypesFormatStickerAPNG   = 2
	FlagTypesFormatStickerLOTTIE = 3
)

// Sticker Item Object
// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type StickerItem struct {
	ID         Snowflake `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	FormatType Flag      `json:"format_type,omitempty"`
}

// Sticker Pack Object
// StickerPack represents a pack of standard stickers.
type StickerPack struct {
	ID            Snowflake `json:"id,omitempty"`
	Type          Flag      `json:"type,omitempty"`
	GuildID       Snowflake `json:"guild_id,omitempty"`
	ChannelID     Snowflake `json:"channel_id,omitempty"`
	User          *User     `json:"user,omitempty"`
	Name          string    `json:"name,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Token         string    `json:"token,omitempty"`
	ApplicationID Snowflake `json:"application_id,omitempty"`
	SourceGuild   *Guild    `json:"source_guild,omitempty"`
	SourceChannel *Channel  `json:"source_channel,omitempty"`
	URL           string    `json:"url,omitempty"`
}

// Webhook Object
// https://discord.com/developers/docs/resources/webhook#webhook-object
type Webhook struct {
	ID            Snowflake  `json:"id,omitempty"`
	Type          Flag       `json:"type,omitempty"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	ChannelID     *Snowflake `json:"channel_id,omitempty"`
	User          *User      `json:"user,omitempty"`
	Name          string     `json:"name,omitempty"`
	Avatar        string     `json:"avatar,omitempty"`
	Token         string     `json:"token,omitempty"`
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	SourceGuild   *Guild     `json:"source_guild,omitempty"`
	SourceChannel *Channel   `json:"source_channel,omitempty"`
	URL           string     `json:"url,omitempty"`
}

// Webhook Types
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
const (
	FlagTypesWebhookINCOMING        = 1
	FlagTypesWebhookCHANNELFOLLOWER = 2
	FlagTypesWebhookAPPLICATION     = 3
)

// Bitwise Permission Flags
// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
const (
	FlagFlagsPermissionBitwiseCREATE_INSTANT_INVITE      = 1 << 0
	FlagFlagsPermissionBitwiseKICK_MEMBERS               = 1 << 1
	FlagFlagsPermissionBitwiseBAN_MEMBERS                = 1 << 2
	FlagFlagsPermissionBitwiseADMINISTRATOR              = 1 << 3
	FlagFlagsPermissionBitwiseMANAGE_CHANNELS            = 1 << 4
	FlagFlagsPermissionBitwiseMANAGE_GUILD               = 1 << 5
	FlagFlagsPermissionBitwiseADD_REACTIONS              = 1 << 6
	FlagFlagsPermissionBitwiseVIEW_AUDIT_LOG             = 1 << 7
	FlagFlagsPermissionBitwisePRIORITY_SPEAKER           = 1 << 8
	FlagFlagsPermissionBitwiseSTREAM                     = 1 << 9
	FlagFlagsPermissionBitwiseVIEW_CHANNEL               = 1 << 10
	FlagFlagsPermissionBitwiseSEND_MESSAGES              = 1 << 11
	FlagFlagsPermissionBitwiseSEND_TTS_MESSAGES          = 1 << 12
	FlagFlagsPermissionBitwiseMANAGE_MESSAGES            = 1 << 13
	FlagFlagsPermissionBitwiseEMBED_LINKS                = 1 << 14
	FlagFlagsPermissionBitwiseATTACH_FILES               = 1 << 15
	FlagFlagsPermissionBitwiseREAD_MESSAGE_HISTORY       = 1 << 16
	FlagFlagsPermissionBitwiseMENTION_EVERYONE           = 1 << 17
	FlagFlagsPermissionBitwiseUSE_EXTERNAL_EMOJIS        = 1 << 18
	FlagFlagsPermissionBitwiseVIEW_GUILD_INSIGHTS        = 1 << 19
	FlagFlagsPermissionBitwiseCONNECT                    = 1 << 20
	FlagFlagsPermissionBitwiseSPEAK                      = 1 << 21
	FlagFlagsPermissionBitwiseMUTE_MEMBERS               = 1 << 22
	FlagFlagsPermissionBitwiseDEAFEN_MEMBERS             = 1 << 23
	FlagFlagsPermissionBitwiseMOVE_MEMBERS               = 1 << 24
	FlagFlagsPermissionBitwiseUSE_VAD                    = 1 << 25
	FlagFlagsPermissionBitwiseCHANGE_NICKNAME            = 1 << 26
	FlagFlagsPermissionBitwiseMANAGE_NICKNAMES           = 1 << 27
	FlagFlagsPermissionBitwiseMANAGE_ROLES               = 1 << 28
	FlagFlagsPermissionBitwiseMANAGE_WEBHOOKS            = 1 << 29
	FlagFlagsPermissionBitwiseMANAGE_EMOJIS_AND_STICKERS = 1 << 30
	FlagFlagsPermissionBitwiseUSE_APPLICATION_COMMANDS   = 1 << 31
	FlagFlagsPermissionBitwiseREQUEST_TO_SPEAK           = 1 << 32
	FlagFlagsPermissionBitwiseMANAGE_EVENTS              = 1 << 33
	FlagFlagsPermissionBitwiseMANAGE_THREADS             = 1 << 34
	FlagFlagsPermissionBitwiseCREATE_PUBLIC_THREADS      = 1 << 35
	FlagFlagsPermissionBitwiseCREATE_PRIVATE_THREADS     = 1 << 36
	FlagFlagsPermissionBitwiseUSE_EXTERNAL_STICKERS      = 1 << 37
	FlagFlagsPermissionBitwiseSEND_MESSAGES_IN_THREADS   = 1 << 38
	FlagFlagsPermissionBitwiseUSE_EMBEDDED_ACTIVITIES    = 1 << 39
	FlagFlagsPermissionBitwiseMODERATE_MEMBERS           = 1 << 40
)

const (
	FlagPermissionOverwriteTypeRole   = 0
	FlagPermissionOverwriteTypeMember = 1
)

// Role Object
// https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID           Snowflake `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Color        int       `json:"color,omitempty"`
	Hoist        bool      `json:"hoist,omitempty"`
	Icon         string    `json:"icon,omitempty"`
	UnicodeEmoji string    `json:"unicode_emoji,omitempty"`
	Position     int       `json:"position,omitempty"`
	Permissions  string    `json:"permissions,omitempty"`
	Managed      bool      `json:"managed,omitempty"`
	Mentionable  bool      `json:"mentionable,omitempty"`
	Tags         *RoleTags `json:"tags,omitempty"`
}

// Role Tags Structure
// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID             Snowflake `json:"bot_id,omitempty"`
	IntegrationID     Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber bool      `json:"premium_subscriber,omitempty"`
}

// Stage Instance Object
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
	ID                    Snowflake  `json:"id,omitempty"`
	GuildID               *Snowflake `json:"guild_id,omitempty"`
	ChannelID             *Snowflake `json:"channel_id,omitempty"`
	Topic                 string     `json:"topic,omitempty"`
	PrivacyLevel          Flag       `json:"privacy_level,omitempty"`
	DiscoverableDisabled  bool       `json:"discoverable_disabled,omitempty"`
	GuildScheduledEventID Snowflake  `json:"guild_scheduled_event_id,omitempty"`
}

// Privacy Level
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
const (
	FlagLevelPrivacyPUBLIC     = 1
	FlagLevelPrivacyGUILD_ONLY = 2
)

// Team Object
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        string        `json:"icon,omitempty"`
	ID          Snowflake     `json:"id,omitempty"`
	Members     []*TeamMember `json:"members,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	OwnerUserID Snowflake     `json:"owner_user_id,omitempty"`
}

// Team Member Object
// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState Flag      `json:"membership_state,omitempty"`
	Permissions     []string  `json:"permissions,omitempty"`
	TeamID          Snowflake `json:"team_id,omitempty"`
	User            *User     `json:"user,omitempty"`
}

// Membership State Enum
// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
const (
	FlagEnumStateMembershipINVITED  = 1
	FlagEnumStateMembershipACCEPTED = 2
)

// User Object
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID            Snowflake `json:"id,omitempty"`
	Username      string    `json:"username,omitempty"`
	Discriminator string    `json:"discriminator,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Bot           bool      `json:"bot,omitempty"`
	System        bool      `json:"system,omitempty"`
	MFAEnabled    bool      `json:"mfa_enabled,omitempty"`
	Banner        string    `json:"banner,omitempty"`
	AccentColor   int       `json:"accent_color,omitempty"`
	Locale        string    `json:"locale,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         *string   `json:"email,omitempty"`
	Flags         *BitFlag  `json:"flag,omitempty"`
	PremiumType   *Flag     `json:"premium_type,omitempty"`
	PublicFlags   BitFlag   `json:"public_flag,omitempty"`
}

// User Flags
// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	FlagFlagsUserNONE                         = 0
	FlagFlagsUserSTAFF                        = 1 << 0
	FlagFlagsUserPARTNER                      = 1 << 1
	FlagFlagsUserHYPESQUAD                    = 1 << 2
	FlagFlagsUserBUG_HUNTER_LEVEL_1           = 1 << 3
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_ONE   = 1 << 6
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_TWO   = 1 << 7
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_THREE = 1 << 8
	FlagFlagsUserPREMIUM_EARLY_SUPPORTER      = 1 << 9
	FlagFlagsUserTEAM_PSEUDO_USER             = 1 << 10
	FlagFlagsUserBUG_HUNTER_LEVEL_2           = 1 << 14
	FlagFlagsUserVERIFIED_BOT                 = 1 << 16
	FlagFlagsUserVERIFIED_DEVELOPER           = 1 << 17
	FlagFlagsUserCERTIFIED_MODERATOR          = 1 << 18
	FlagFlagsUserBOT_HTTP_INTERACTIONS        = 1 << 19
)

// Premium Types
// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	FlagTypesPremiumNONE         = 0
	FlagTypesPremiumNITROCLASSIC = 1
	FlagTypesPremiumNITRO        = 2
)

// User Connection Object
// https://discord.com/developers/docs/resources/user#connection-object-connection-structure
type Connection struct {
	ID           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Type         string         `json:"type,omitempty"`
	Revoked      bool           `json:"revoked,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
	Verified     bool           `json:"verified,omitempty"`
	FriendSync   bool           `json:"friend_sync,omitempty"`
	ShowActivity bool           `json:"show_activity,omitempty"`
	Visibility   Flag           `json:"visibility,omitempty"`
}

// Visibility Types
// https://discord.com/developers/docs/resources/user#connection-object-visibility-types
const (
	FlagTypesVisibilityNONE     = 0
	FlagTypesVisibilityEVERYONE = 1
)

// Voice State Object
// https://discord.com/developers/docs/resources/voice#voice-state-object-voice-state-structure
type VoiceState struct {
	GuildID                 Snowflake    `json:"guild_id,omitempty"`
	ChannelID               Snowflake    `json:"channel_id,omitempty"`
	UserID                  Snowflake    `json:"user_id,omitempty"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id,omitempty"`
	Deaf                    bool         `json:"deaf,omitempty"`
	Mute                    bool         `json:"mute,omitempty"`
	SelfDeaf                bool         `json:"self_deaf,omitempty"`
	SelfMute                bool         `json:"self_mute,omitempty"`
	SelfStream              bool         `json:"self_stream,omitempty"`
	SelfVideo               bool         `json:"self_video,omitempty"`
	Suppress                bool         `json:"suppress,omitempty"`
	RequestToSpeakTimestamp time.Time    `json:"request_to_speak_timestamp,omitempty"`
}

// Voice Region Object
// https://discord.com/developers/docs/resources/voice#voice-region-object-voice-region-structure
type VoiceRegion struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Optimal    bool   `json:"optimal,omitempty"`
	Deprecated bool   `json:"deprecated,omitempty"`
	Custom     bool   `json:"custom,omitempty"`
}
