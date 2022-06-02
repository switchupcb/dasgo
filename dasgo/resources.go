// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import (
	"time"
)

//I assumed type snowflake is nonzero

// Application Command Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	ID                       Snowflake                   `json:"id"`
	Type                     Flag                        `json:"type,omitempty"`
	ApplicationID            Snowflake                   `json:"application_id"`
	GuildID                  Snowflake                   `json:"guild_id,omitempty"`
	Name                     string                      `json:"name"`
	NameLocalizations        map[Flag]string             `json:"name_localizations"`
	Description              string                      `json:"description"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                      `json:"default_member_permissions"`
	DMPermission             *bool                        `json:"dm_permission,omitempty"`
	Version                  Snowflake                   `json:"version,omitempty"`
}

// Application Command Types
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
const (
	FlagApplicationCommandTypeCHAT_INPUT = 1
	FlagApplicationCommandTypeUSER       = 2
	FlagApplicationCommandTypeMESSAGE    = 3
)

// Application Command Option Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     Flag                              `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        map[Flag]string                   `json:"name_localizations"`
	Description              string                            `json:"description"`
	DescriptionLocalizations map[Flag]string                   `json:"description_localizations"`
	Required                 *bool                              `json:"required,omitempty"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []*ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []*Flag                           `json:"channel_types,omitempty"`
	MinValue                 *float64                           `json:"min_value,omitempty"`
	MaxValue                 *float64                           `json:"max_value,omitempty"`
	Autocomplete             *bool                              `json:"autocomplete,omitempty"`
}

// Application Command Option Type
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
const (
	FlagApplicationCommandOptionTypeSUB_COMMAND       = 1
	FlagApplicationCommandOptionTypeSUB_COMMAND_GROUP = 2
	FlagApplicationCommandOptionTypeSTRING            = 3
	FlagApplicationCommandOptionTypeINTEGER           = 4
	FlagApplicationCommandOptionTypeBOOLEAN           = 5
	FlagApplicationCommandOptionTypeUSER              = 6
	FlagApplicationCommandOptionTypeCHANNEL           = 7
	FlagApplicationCommandOptionTypeROLE              = 8
	FlagApplicationCommandOptionTypeMENTIONABLE       = 9
	FlagApplicationCommandOptionTypeNUMBER            = 10
	FlagApplicationCommandOptionTypeATTACHMENT        = 11
)

// Application Command Option Choice
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure
type ApplicationCommandOptionChoice struct {
	Name              string          `json:"name"`
	NameLocalizations map[Flag]string `json:"name_localizations"`
	Value             interface{}     `json:"value"`
}

// Application Command Interaction Data Option Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    Flag                                       `json:"type"`
	Value   *interface{}                                `json:"value,omitempty"`
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused *bool                                       `json:"focused,omitempty"`
}

// Guild Application Command Permissions Object
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                        `json:"id"`
	ApplicationID Snowflake                        `json:"application_id"`
	GuildID       Snowflake                        `json:"guild_id"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions"`
}

// Application Command Permissions Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure
type ApplicationCommandPermissions struct {
	ID         Snowflake `json:"id"`
	Type       Flag      `json:"type"`
	Permission bool      `json:"permission"`
}

// Application Command Permission Type
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permission-type
const (
	FlagApplicationCommandPermissionTypeROLE = 1
	FlagApplicationCommandPermissionTypeUSER = 2
)

// Component Object
type Component interface {
	Type()
}

// Component Types
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
const (
	FlagComponentTypeActionRow  = 1
	FlagComponentTypeButton     = 2
	FlagComponentTypeSelectMenu = 3
	FlagComponentTypeTextInput  = 4
)

func (c ActionsRow) Type() Flag {
	return FlagComponentTypeActionRow
}

func (c Button) Type() Flag {
	return FlagComponentTypeButton
}

func (c SelectMenu) Type() Flag {
	return FlagComponentTypeSelectMenu
}

func (c TextInput) Type() Flag {
	return FlagComponentTypeTextInput
}

//was unclear on optionality and nullablility so I dropped omitempty.

// https://discord.com/developers/docs/interactions/message-components#component-object
type ActionsRow struct {
	Components []Component `json:"components"`
}



// Button Object

// *The type field was not included

// https://discord.com/developers/docs/interactions/message-components#button-object
type Button struct {
	Style    Flag    `json:"style"`
	Label    *string `json:"label,omitempty"`
	Emoji    *Emoji  `json:"emoji,omitempty"`
	CustomID *string  `json:"custom_id,omitempty"`
	URL      *string  `json:"url,omitempty"`
	Disabled *bool    `json:"disabled,omitempty"`
}

// Button Styles
// https://discord.com/developers/docs/interactions/message-components#button-object-button-styles
const (
	FlagButtonStylePRIMARY   = 1
	FlagButtonStyleBLURPLE   = 1
	FlagButtonStyleSecondary = 2
	FlagButtonStyleGREY      = 2
	FlagButtonStyleSuccess   = 3
	FlagButtonStyleGREEN     = 3
	FlagButtonStyleDanger    = 4
	FlagButtonStyleRED       = 4
	FlagButtonStyleLINK      = 5
)

// Select Menu Structure

// *

// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure
type SelectMenu struct {
	CustomID    string             `json:"custom_id"`
	Options     []SelectMenuOption `json:"options"`
	Placeholder *string             `json:"placeholder,omitempty"`
	MinValues   *Flag              `json:"min_values,omitempty"`
	MaxValues   *Flag               `json:"max_values,omitempty"`
	Disabled    *bool               `json:"disabled,omitempty"`
}

// Select Menu Option Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectMenuOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description *string `json:"description,omitempty"`
	Emoji       *Emoji   `json:"emoji,omitempty"`
	Default     *bool    `json:"default,omitempty"`
}

// Text Input Structure

// *

// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInput struct {
	CustomID    string    `json:"custom_id"`
	Style       Flag      `json:"style"`
	Label       *string   `json:"label"`
	MinLength   *CodeFlag `json:"min_length,omitempty"`
	MaxLength   *CodeFlag  `json:"max_length,omitempty"`
	Required    *bool      `json:"required,omitempty"`
	Value       *string    `json:"value,omitempty"`
	Placeholder *string   `json:"placeholder,omitempty"`
}

// Text Input Styles
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-styles
const (
	FlagTextInputStyleShort     = 1
	FlagTextInputStyleParagraph = 2
)

// Interaction Object
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	ID            Snowflake       `json:"id"`
	ApplicationID Snowflake       `json:"application_id"`
	Type          Flag            `json:"type"`
	Data          InteractionData `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id,omitempty"`
	ChannelID     Snowflake       `json:"channel_id,omitempty"`
	Member        *GuildMember    `json:"member,omitempty"`
	User          *User           `json:"user,omitempty"`
	Token         string          `json:"token"`
	Version       Flag            `json:"version"`
	Message       *Message        `json:"message,omitempty"`
	Locale        *string          `json:"locale,omitempty"`
	GuildLocale   *string          `json:"guild_locale,omitempty"`
}

// Interaction Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
const (
	FlagInteractionTypePING                             = 1
	FlagInteractionTypeAPPLICATION_COMMAND              = 2
	FlagInteractionTypeMESSAGE_COMPONENT                = 3
	FlagInteractionTypeAPPLICATION_COMMAND_AUTOCOMPLETE = 4
	FlagInteractionTypeMODAL_SUBMIT                     = 5
)

// Interaction Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data-structure
type InteractionData struct {
	ID            Snowflake                                  `json:"id"`
	Name          string                                     `json:"name"`
	Type          Flag                                       `json:"type"`
	Resolved      *ResolvedData                              `json:"resolved,omitempty"`
	Options       []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildID       Snowflake                                  `json:"guild_id,omitempty"`
	CustomID      *string                                     `json:"custom_id,omitempty"`
	ComponentType Flag                                       `json:"component_type,omitempty"`
	Values        []*string                                  `json:"values,omitempty"`
	TargetID      Snowflake                                  `json:"target_id,omitempty"`
	Components    []*Component                               `json:"components,omitempty"`
}

// Resolved Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type ResolvedData struct {
	Users       map[Snowflake]*User        `json:"users"`
	Members     map[Snowflake]*GuildMember `json:"members"`
	Roles       map[Snowflake]*Role        `json:"roles"`
	Channels    map[Snowflake]*Channel     `json:"channels"`
	Messages    map[Snowflake]*Message     `json:"messages"`
	Attachments map[Snowflake]*Attachment  `json:"attachments"`
}

// Message Interaction Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	ID     Snowflake    `json:"id"`
	Type   Flag         `json:"type"`
	Name   string       `json:"name"`
	User   *User        `json:"user"`
	Member *GuildMember `json:"member,omitempty"`
}

// Interaction Response Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure
type InteractionResponse struct {
	Type Flag                     `json:"type"`
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// Interaction Callback Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
const (
	FlagInteractionCallbackTypePONG                                    = 1
	FlagInteractionCallbackTypeCHANNEL_MESSAGE_WITH_SOURCE             = 4
	FlagInteractionCallbackTypeDEFERRED_CHANNEL_MESSAGE_WITH_SOURCE    = 5
	FlagInteractionCallbackTypeDEFERRED_UPDATE_MESSAGE                 = 6
	FlagInteractionCallbackTypeUPDATE_MESSAGE                          = 7
	FlagInteractionCallbackTypeAPPLICATION_COMMAND_AUTOCOMPLETE_RESULT = 8
	FlagInteractionCallbackTypeMODAL                                   = 9
)

// Interaction Callback Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionCallbackData interface{}

// Messages
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-messages
type Messages struct {
	TTS             *bool             `json:"tts,omitempty"`
	Content         *string           `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           *BitFlag          `json:"flags,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// Autocomplete
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-autocomplete
type Autocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices"`
}

// Modal
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-modal
type ModalSubmitInteractionData struct {
	CustomID   *string     `json:"custom_id"`
	Title      string      `json:"title"`
	Components []Component `json:"components"`
}

// Application Object
// https://discord.com/developers/docs/resources/application
type Application struct {
	ID                  Snowflake      `json:"id"`
	Name                string         `json:"name"`
	Icon                string         `json:"icon"`
	Description         string         `json:"description"`
	RPCOrigins          []string       `json:"rpc_origins,omitempty"`
	BotPublic           bool           `json:"bot_public"`
	BotRequireCodeGrant bool           `json:"bot_require_code_grant"`
	TermsOfServiceURL   string         `json:"terms_of_service_url,omitempty"`
	PrivacyProxyURL     string         `json:"privacy_policy_url,omitempty"`
	Owner               *User          `json:"owner,omitempty"`
	VerifyKey           string         `json:"verify_key"`
	Team                *Team          `json:"team"`
	GuildID             Snowflake      `json:"guild_id,omitempty"`
	PrimarySKUID        Snowflake      `json:"primary_sku_id,omitempty"`
	Slug                *string        `json:"slug,omitempty"`
	CoverImage          *string         `json:"cover_image,omitempty"`
	Flags               *Flag           `json:"flags,omitempty"`
	Tags                []string       `json:"tags,omitempty"`
	InstallParams       *InstallParams `json:"install_params,omitempty"`
	CustomInstallURL    string         `json:"custom_install_url,omitempty"`
}

// Application Flags
// https://discord.com/developers/docs/resources/application#application-object-application-flags
const (
	FlagApplicationGATEWAY_PRESENCE                 = 1 << 12
	FlagApplicationGATEWAY_PRESENCE_LIMITED         = 1 << 13
	FlagApplicationGATEWAY_GUILD_MEMBERS            = 1 << 14
	FlagApplicationGATEWAY_GUILD_MEMBERS_LIMITED    = 1 << 15
	FlagApplicationVERIFICATION_PENDING_GUILD_LIMIT = 1 << 16
	FlagApplicationEMBEDDED                         = 1 << 17
	FlagApplicationGATEWAY_MESSAGE_CONTENT          = 1 << 18
	FlagApplicationGATEWAY_MESSAGE_CONTENT_LIMITED  = 1 << 19
)

// Install Params Object
// https://discord.com/developers/docs/resources/application#install-params-object
type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}

// Audit Log Object
// https://discord.com/developers/docs/resources/audit-log
type AuditLog struct {
	AuditLogEntries      []*AuditLogEntry       `json:"audit_log_entries"`
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events"`
	Integration          []*Integration         `json:"integrations"`
	Threads              []*Channel             `json:"threads"`
	Users                []*User                `json:"users"`
	Webhooks             []*Webhook             `json:"webhooks"`
}

// Audit Log Entry Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type AuditLogEntry struct {
	TargetID   string            `json:"target_id"`
	Changes    []*AuditLogChange `json:"changes,omitempty"`
	UserID     Snowflake         `json:"user_id"`
	ID         Snowflake         `json:"id"`
	ActionType Flag              `json:"action_type"`
	Options    *AuditLogOptions  `json:"options,omitempty"`
	Reason     *string           `json:"reason,omitempty"`
}

// Audit Log Events
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
const (
	FlagAuditLogEventGUILD_UPDATE                          = 1
	FlagAuditLogEventCHANNEL_CREATE                        = 10
	FlagAuditLogEventCHANNEL_UPDATE                        = 11
	FlagAuditLogEventCHANNEL_DELETE                        = 12
	FlagAuditLogEventCHANNEL_OVERWRITE_CREATE              = 13
	FlagAuditLogEventCHANNEL_OVERWRITE_UPDATE              = 14
	FlagAuditLogEventCHANNEL_OVERWRITE_DELETE              = 15
	FlagAuditLogEventMEMBER_KICK                           = 20
	FlagAuditLogEventMEMBER_PRUNE                          = 21
	FlagAuditLogEventMEMBER_BAN_ADD                        = 22
	FlagAuditLogEventMEMBER_BAN_REMOVE                     = 23
	FlagAuditLogEventMEMBER_UPDATE                         = 24
	FlagAuditLogEventMEMBER_ROLE_UPDATE                    = 25
	FlagAuditLogEventMEMBER_MOVE                           = 26
	FlagAuditLogEventMEMBER_DISCONNECT                     = 27
	FlagAuditLogEventBOT_ADD                               = 28
	FlagAuditLogEventROLE_CREATE                           = 30
	FlagAuditLogEventROLE_UPDATE                           = 31
	FlagAuditLogEventROLE_DELETE                           = 32
	FlagAuditLogEventINVITE_CREATE                         = 40
	FlagAuditLogEventINVITE_UPDATE                         = 41
	FlagAuditLogEventINVITE_DELETE                         = 42
	FlagAuditLogEventWEBHOOK_CREATE                        = 50
	FlagAuditLogEventWEBHOOK_UPDATE                        = 51
	FlagAuditLogEventWEBHOOK_DELETE                        = 52
	FlagAuditLogEventEMOJI_CREATE                          = 60
	FlagAuditLogEventEMOJI_UPDATE                          = 61
	FlagAuditLogEventEMOJI_DELETE                          = 62
	FlagAuditLogEventMESSAGE_DELETE                        = 72
	FlagAuditLogEventMESSAGE_BULK_DELETE                   = 73
	FlagAuditLogEventMESSAGE_PIN                           = 74
	FlagAuditLogEventMESSAGE_UNPIN                         = 75
	FlagAuditLogEventINTEGRATION_CREATE                    = 80
	FlagAuditLogEventINTEGRATION_UPDATE                    = 81
	FlagAuditLogEventINTEGRATION_DELETE                    = 82
	FlagAuditLogEventSTAGE_INSTANCE_CREATE                 = 83
	FlagAuditLogEventSTAGE_INSTANCE_UPDATE                 = 84
	FlagAuditLogEventSTAGE_INSTANCE_DELETE                 = 85
	FlagAuditLogEventSTICKER_CREATE                        = 90
	FlagAuditLogEventSTICKER_UPDATE                        = 91
	FlagAuditLogEventSTICKER_DELETE                        = 92
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_CREATE          = 100
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_UPDATE          = 101
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_DELETE          = 102
	FlagAuditLogEventTHREAD_CREATE                         = 110
	FlagAuditLogEventTHREAD_UPDATE                         = 111
	FlagAuditLogEventTHREAD_DELETE                         = 112
	FlagAuditLogEventAPPLICATION_COMMAND_PERMISSION_UPDATE = 121
)

// Optional Audit Entry Info
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogOptions struct {
	ApplicationID    Snowflake `json:"application_id"`
	ChannelID        Snowflake `json:"channel_id"`
	Count            string    `json:"count"`
	DeleteMemberDays string    `json:"delete_member_days"`
	ID               Snowflake `json:"id"`
	MembersRemoved   string    `json:"members_removed"`
	MessageID        Snowflake `json:"message_id"`
	RoleName         string    `json:"role_name"`
	Type             string    `json:"type"`
}

// Audit Log Change Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object
type AuditLogChange struct {
	NewValue interface{} `json:"new_value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
	Key      string      `json:"key"`
}

// Audit Log Change Exceptions
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-exceptions

// not sure for last pin timestamp

// Channel Object
// https://discord.com/developers/docs/resources/channel
type Channel struct {
	ID                         Snowflake             `json:"id"`
	Type                       *Flag                 `json:"type"`
	GuildID                    Snowflake             `json:"guild_id,omitempty"`
	Position                   *int                   `json:"position,omitempty"`
	PermissionOverwrites       []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                       string                `json:"name,omitempty"`
	Topic                      *string               `json:"topic,omitempty"`
	NSFW                       *bool                  `json:"nsfw,omitempty"`
	LastMessageID              *Snowflake             `json:"last_message_id"`
	Bitrate                    *Flag                  `json:"bitrate,omitempty"`
	UserLimit                  *Flag                  `json:"user_limit,omitempty"`
	RateLimitPerUser           *CodeFlag             `json:"rate_limit_per_user,omitempty"`
	Recipients                 []*User               `json:"recipients,omitempty"`
	Icon                       *string                `json:"icon"`
	OwnerID                    Snowflake             `json:"owner_id,omitempty"`
	ApplicationID              Snowflake             `json:"application_id,omitempty"`
	ParentID                   *Snowflake             `json:"parent_id"`
	LastPinTimestamp           *time.Time             `json:"last_pin_timestamp"`
	RTCRegion                  *string                `json:"rtc_region"`
	VideoQualityMode           Flag                  `json:"video_quality_mode,omitempty"`
	MessageCount               *Flag                  `json:"message_count,omitempty"`
	MemberCount                *Flag                  `json:"member_count,omitempty"`
	ThreadMetadata             *ThreadMetadata       `json:"thread_metadata,omitempty"`
	Member                     *ThreadMember         `json:"member,omitempty"`
	DefaultAutoArchiveDuration int                   `json:"default_auto_archive_duration,omitempty"`
	Permissions                *string               `json:"permissions,omitempty"`
	Flags                      BitFlag               `json:"flags,omitempty"`
}

// Channel Types
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
const (
	FlagChannelTypeGUILD_TEXT           = 0
	FlagChannelTypeDM                   = 1
	FlagChannelTypeGUILD_VOICE          = 2
	FlagChannelTypeGROUP_DM             = 3
	FlagChannelTypeGUILD_CATEGORY       = 4
	FlagChannelTypeGUILD_NEWS           = 5
	FlagChannelTypeGUILD_NEWS_THREAD    = 10
	FlagChannelTypeGUILD_PUBLIC_THREAD  = 11
	FlagChannelTypeGUILD_PRIVATE_THREAD = 12
	FlagChannelTypeGUILD_STAGE_VOICE    = 13
	FlagChannelTypeGUILD_DIRECTORY      = 14
	FlagChannelTypeGUILD_FORUM          = 15
)

// Video Quality Modes
// https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
const (
	FlagVideoQualityModeAUTO = 1
	FlagVideoQualityModeFULL = 2
)

// Channel Flags
// https://discord.com/developers/docs/resources/channel#channel-object-channel-flags
const (
	FlagChannelPINNED = 1 << 1
)

// Message Object
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                Snowflake         `json:"id"`
	ChannelID         *Snowflake        `json:"channel_id"`
	GuildID           *Snowflake        `json:"guild_id,omitempty"`
	Author            *User             `json:"author"`
	Member            *GuildMember      `json:"member,omitempty"`
	Content           string            `json:"content"`
	Timestamp         time.Time         `json:"timestamp"`
	EditedTimestamp   time.Time         `json:"edited_timestamp"`
	TTS               bool              `json:"tts"`
	MentionEveryone   bool              `json:"mention_everyone"`
	Mentions          []*User           `json:"mentions"`
	MentionRoles      []*Snowflake      `json:"mention_roles"`
	MentionChannels   []*ChannelMention `json:"mention_channels,omitempty"`
	Attachments       []*Attachment     `json:"attachments"`
	Embeds            []*Embed          `json:"embeds"`
	Reactions         []*Reaction       `json:"reactions,omitempty"`
	Nonce             interface{}       `json:"nonce,omitempty"`
	Pinned            bool              `json:"pinned"`
	WebhookID         *Snowflake        `json:"webhook_id,omitempty"`
	Type              *Flag             `json:"type"`
	Activity          MessageActivity   `json:"activity,omitempty"`
	Application       *Application      `json:"application,omitempty"`
	ApplicationID     Snowflake         `json:"application_id,omitempty"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	Flags             *CodeFlag          `json:"flags,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	Interaction       *Interaction      `json:"interaction,omitempty"`
	Thread            *Channel          `json:"thread,omitempty"`
	Components        []*Component      `json:"components,omitempty"`
	StickerItems      []*StickerItem    `json:"sticker_items,omitempty"`
}

// Message Types
// https://discord.com/developers/docs/resources/channel#message-object-message-types
const (
	FlagMessageTypeDEFAULT                                      = 0
	FlagMessageTypeRECIPIENT_ADD                                = 1
	FlagMessageTypeRECIPIENT_REMOVE                             = 2
	FlagMessageTypeCALL                                         = 3
	FlagMessageTypeCHANNEL_NAME_CHANGE                          = 4
	FlagMessageTypeCHANNEL_ICON_CHANGE                          = 5
	FlagMessageTypeCHANNEL_PINNED_MESSAGE                       = 6
	FlagMessageTypeGUILD_MEMBER_JOIN                            = 7
	FlagMessageTypeUSER_PREMIUM_GUILD_SUBSCRIPTION              = 8
	FlagMessageTypeUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_ONE     = 9
	FlagMessageTypeUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_TWO     = 10
	FlagMessageTypeUSER_PREMIUM_GUILD_SUBSCRIPTION_TIER_THREE   = 11
	FlagMessageTypeCHANNEL_FOLLOW_ADD                           = 12
	FlagMessageTypeGUILD_DISCOVERY_DISQUALIFIED                 = 14
	FlagMessageTypeGUILD_DISCOVERY_REQUALIFIED                  = 15
	FlagMessageTypeGUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING = 16
	FlagMessageTypeGUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING   = 17
	FlagMessageTypeTHREAD_CREATED                               = 18
	FlagMessageTypeREPLY                                        = 19
	FlagMessageTypeCHAT_INPUT_COMMAND                           = 20
	FlagMessageTypeTHREAD_STARTER_MESSAGE                       = 21
	FlagMessageTypeGUILD_INVITE_REMINDER                        = 22
	FlagMessageTypeCONTEXT_MENU_COMMAND                         = 23
)

// Message Activity Structure
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	Type    int     `json:"type"`
	PartyID *string `json:"party_id,omitempty"`
}

// Message Activity Types
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
const (
	FlagMessageActivityTypeJOIN         = 1
	FlagMessageActivityTypeSPECTATE     = 2
	FlagMessageActivityTypeLISTEN       = 3
	FlagMessageActivityTypeJOIN_REQUEST = 5
)

// Message Flags
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
const (
	FlagMessageCROSSPOSTED                            = 1 << 0
	FlagMessageIS_CROSSPOST                           = 1 << 1
	FlagMessageSUPPRESS_EMBEDS                        = 1 << 2
	FlagMessageSOURCE_MESSAGE_DELETED                 = 1 << 3
	FlagMessageURGENT                                 = 1 << 4
	FlagMessageHAS_THREAD                             = 1 << 5
	FlagMessageEPHEMERAL                              = 1 << 6
	FlagMessageLOADING                                = 1 << 7
	FlagMessageFAILED_TO_MENTION_SOME_ROLES_IN_THREAD = 1 << 8
)

// Message Reference Object
// https://discord.com/developers/docs/resources/channel#message-reference-object
type MessageReference struct {
	MessageID       Snowflake  `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists *bool       `json:"fail_if_not_exists,omitempty"`
}

// Followed Channel Structure
// https://discord.com/developers/docs/resources/channel#followed-channel-object-followed-channel-structure
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// Reaction Object
// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count CodeFlag `json:"count"`
	Me    bool     `json:"me"`
	Emoji *Emoji   `json:"emoji"`
}

// Overwrite Object
// https://discord.com/developers/docs/resources/channel#overwrite-object
type PermissionOverwrite struct {
	ID    Snowflake `json:"id"`
	Type  *Flag     `json:"type"`
	Deny  string    `json:"deny"`
	Allow string    `json:"allow"`
}

// Thread Metadata Object
// https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool      `json:"archived"`
	AutoArchiveDuration int       `json:"auto_archive_duration"`
	Locked              bool      `json:"locked"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp"`
}

// Thread Member Object
// https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	ThreadID      Snowflake `json:"id,omitempty"`
	UserID        Snowflake `json:"user_id,omitempty"`
	JoinTimestamp time.Time `json:"join_timestamp"`
	Flags         CodeFlag  `json:"flags"`
}

// Embed Object
// https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       *string          `json:"title,omitempty"`
	Type        *string          `json:"type,omitempty"`
	Description *string         `json:"description,omitempty"`
	URL         *string          `json:"url,omitempty"`
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
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int     `json:"height,omitempty"`
	Width    *int     `json:"width,omitempty"`
}

// Embed Video Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	URL      *string  `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int     `json:"height,omitempty"`
	Width    *int     `json:"width,omitempty"`
}

// Embed Image Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int     `json:"height,omitempty"`
	Width    *int     `json:"width,omitempty"`
}

// Embed Provider Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// Embed Author Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

// Embed Footer Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      *string  `json:"icon_url,omitempty"`
	ProxyIconURL *string  `json:"proxy_icon_url,omitempty"`
}

// Embed Field Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool   `json:"inline,omitempty"`
}

// Embed Limits
// https://discord.com/developers/docs/resources/channel#embed-object-embed-limits
const (
	FlagEmbedLimitTitle       = 256
	FlagEmbedLimitDescription = 4096
	FlagEmbedLimitFields      = 25
	FlagEmbedLimitFieldName   = 256
	FlagEmbedLimitFieldValue  = 1024
	FlagEmbedLimitFooterText  = 2048
	FlagEmbedLimitAuthorName  = 256
)

// Message Attachment Object
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID          Snowflake `json:"id"`
	Filename    string    `json:"filename"`
	Description *string    `json:"description,omitempty"`
	ContentType *string    `json:"content_type,omitempty"`
	Size        int       `json:"size"`
	URL         string    `json:"url"`
	ProxyURL    *string   `json:"proxy_url"`
	Height      *int       `json:"height"`
	Width       *int       `json:"width"`
	Emphemeral  *bool      `json:"ephemeral,omitempty"`
}

// Channel Mention Object
// https://discord.com/developers/docs/resources/channel#channel-mention-object
type ChannelMention struct {
	ID      Snowflake `json:"id"`
	GuildID Snowflake `json:"guild_id"`
	Type    *Flag     `json:"type"`
	Name    string    `json:"name"`
}

// Allowed Mentions Structure
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMentions struct {
	Parse       []*string    `json:"parse"`
	Roles       []*Snowflake `json:"roles"`
	Users       []*Snowflake `json:"users"`
	RepliedUser bool         `json:"replied_user"`
}

// Allowed Mention Types
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
const (
	FlagAllowedMentionTypeRoles    = "roles"
	FlagAllowedMentionTypeUsers    = "users"
	FlagAllowedMentionTypeEveryone = "everyone"
)

// Emoji Object
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	ID            Snowflake   `json:"id"`
	Name          *string     `json:"name"`
	Roles         []Snowflake `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons *bool        `json:"require_colons,omitempty"`
	Managed       *bool        `json:"managed,omitempty"`
	Animated      *bool        `json:"animated,omitempty"`
	Available     *bool        `json:"available,omitempty"`
}

// Guild Object
// https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	ID                          Snowflake      `json:"id"`
	Name                        string         `json:"name"`
	Icon                        string         `json:"icon"`
	IconHash                    *string         `json:"icon_hash"`
	Splash                      string         `json:"splash"`
	DiscoverySplash             string         `json:"discovery_splash"`
	Owner                       *bool           `json:"owner,omitempty"`
	OwnerID                     Snowflake      `json:"owner_id"`
	Permissions                 *string        `json:"permissions,omitempty"`
	Region                      *string         `json:"region"`
	AfkChannelID                Snowflake      `json:"afk_channel_id"`
	AfkTimeout                  int            `json:"afk_timeout"`
	WidgetEnabled               *bool           `json:"widget_enabled,omitempty"`
	WidgetChannelID             *Snowflake      `json:"widget_channel_id"`
	VerificationLevel           *Flag          `json:"verification_level"`
	DefaultMessageNotifications *Flag          `json:"default_message_notifications"`
	ExplicitContentFilter       *Flag          `json:"explicit_content_filter"`
	Roles                       []*Role        `json:"roles"`
	Emojis                      []*Emoji       `json:"emojis"`
	Features                    []*string      `json:"features"`
	MFALevel                    *Flag          `json:"mfa_level"`
	ApplicationID               Snowflake      `json:"application_id"`
	SystemChannelID             Snowflake      `json:"system_channel_id"`
	SystemChannelFlags          BitFlag        `json:"system_channel_flags"`
	RulesChannelID              Snowflake      `json:"rules_channel_id"`
	MaxPresences                *CodeFlag       `json:"max_presences"`
	MaxMembers                  int            `json:"max_members,omitempty"`
	VanityUrl                   *string        `json:"vanity_url_code"`
	Description                 *string        `json:"description"`
	Banner                      *string         `json:"banner"`
	PremiumTier                 *Flag          `json:"premium_tier"`
	PremiumSubscriptionCount    *CodeFlag      `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale"`
	PublicUpdatesChannelID      Snowflake      `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        *int            `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      *int            `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    *int            `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NSFWLevel                   *Flag          `json:"nsfw_level"`
	Stickers                    []*Sticker     `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled"`

	// Unavailable Guild Object
	// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Unavailable bool `json:"unavailable,omitempty"`
}

// Default Message Notification Level
// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
const (
	FlagDefaultMessageNotificationLevelALL_MESSAGES  = 0
	FlagDefaultMessageNotificationLevelONLY_MENTIONS = 1
)

// Explicit Content Filter Level
// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
const (
	FlagExplicitContentFilterLevelDISABLED              = 0
	FlagExplicitContentFilterLevelMEMBERS_WITHOUT_ROLES = 1
	FlagExplicitContentFilterLevelALL_MEMBERS           = 2
)

// MFA Level
// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
const (
	FlagMFALevelNONE     = 0
	FlagMFALevelELEVATED = 1
)

// Verification Level
// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
const (
	FlagVerificationLevelNONE      = 0
	FlagVerificationLevelLOW       = 1
	FlagVerificationLevelMEDIUM    = 2
	FlagVerificationLevelHIGH      = 3
	FlagVerificationLevelVERY_HIGH = 4
)

// Guild NSFW Level
// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
const (
	FlagGuildNSFWLevelDEFAULT        = 0
	FlagGuildNSFWLevelEXPLICIT       = 1
	FlagGuildNSFWLevelSAFE           = 2
	FlagGuildNSFWLevelAGE_RESTRICTED = 3
)

// Premium Tier
// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
const (
	FlagPremiumTierNONE  = 0
	FlagPremiumTierONE   = 1
	FlagPremiumTierTWO   = 2
	FlagPremiumTierTHREE = 3
)

// System Channel Flags
// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
const (
	FlagSystemChannelSUPPRESS_JOIN_NOTIFICATIONS           = 1 << 0
	FlagSystemChannelSUPPRESS_PREMIUM_SUBSCRIPTIONS        = 1 << 1
	FlagSystemChannelSUPPRESS_GUILD_REMINDER_NOTIFICATIONS = 1 << 2
	FlagSystemChannelSUPPRESS_JOIN_NOTIFICATION_REPLIES    = 1 << 3
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
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Icon                     string     `json:"icon"`
	Splash                   string     `json:"splash"`
	DiscoverySplash          string     `json:"discovery_splash"`
	Emojis                   []*Emoji   `json:"emojis"`
	Features                 []*string  `json:"features"`
	ApproximateMemberCount   int        `json:"approximate_member_count"`
	ApproximatePresenceCount int        `json:"approximate_presence_count"`
	Description              *string    `json:"description"`
	Stickers                 []*Sticker `json:"stickers"`
}

// Guild Widget Settings Object
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object
type GuildWidgetSettings struct {
	Enabled   bool      `json:"enabled"`
	ChannelID Snowflake `json:"channel_id"`
}

// Guild Widget Object
// https://discord.com/developers/docs/resources/guild#et-gguild-widget-object-get-guild-widget-structure*
type GuildWidget struct {
	ID            Snowflake  `json:"id"`
	Name          string     `json:"name"`
	InstantInvite *string     `json:"instant_invite"`
	Channels      []*Channel `json:"channels"`
	Members       []*User    `json:"members"`
	PresenceCount int        `json:"presence_count"`
}

// Guild Member Object

//GuildID not in api documentation

// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User                       *User        `json:"user,omitempty"`
	Nick                       *string      `json:"nick"`
	Avatar                     *string       `json:"avatar"`
	Roles                      []*Snowflake `json:"roles"`
	GuildID                    Snowflake    `json:"guild_id,omitempty"`
	JoinedAt                   time.Time    `json:"joined_at"`
	PremiumSince               *time.Time    `json:"premium_since"`
	Deaf                       bool         `json:"deaf"`
	Mute                       bool         `json:"mute"`
	Pending                    *bool         `json:"pending,omitempty"`
	Permissions                *string      `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time    `json:"communication_disabled_until,omitempty"`
}

// Integration Object
// https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	ID                Snowflake          `json:"id"`
	Name              string             `json:"name"`
	Type              string             `json:"type"`
	Enabled           *bool               `json:"enabled,omitempty"`
	Syncing           *bool               `json:"syncing,omitempty"`
	RoleID            Snowflake          `json:"role_id,omitempty"`
	EnableEmoticons   *bool               `json:"enable_emoticons,omitempty"`
	ExpireBehavior    *Flag              `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *int               `json:"expire_grace_period,omitempty"`
	User              *User              `json:"user,omitempty"`
	Account           IntegrationAccount `json:"account"`
	SyncedAt          *time.Time          `json:"synced_at,omitempty"`
	SubscriberCount   *int               `json:"subscriber_count,omitempty"`
	Revoked           *bool               `json:"revoked,omitempty"`
	Application       *Application       `json:"application,omitempty"`
}

// Integration Expire Behaviors
// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
const (
	FlagIntegrationExpireBehaviorREMOVEROLE = 0
	FlagIntegrationExpireBehaviorKICK       = 1
)

// Integration Account Object
// https://discord.com/developers/docs/resources/guild#integration-account-object
type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Integration Application Object
// https://discord.com/developers/docs/resources/guild#integration-application-object-integration-application-structure
type IntegrationApplication struct {
	ID          Snowflake `json:"id"`
	Name        string    `json:"name"`
	Icon        *string    `json:"icon"`
	Description *string   `json:"description"`
	Bot         *User     `json:"bot,omitempty"`
}

// Guild Ban Object
// https://discord.com/developers/docs/resources/guild#ban-object
type Ban struct {
	Reason *string `json:"reason"`
	User   *User   `json:"user"`
}

// Welcome Screen Object
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-structure
type WelcomeScreen struct {
	Description           *string                 `json:"description"`
	WelcomeScreenChannels []*WelcomeScreenChannel `json:"welcome_channels"`
}

// Welcome Screen Channel Structure
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure
type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id"`
	Description *string    `json:"description"`
	EmojiID     *Snowflake `json:"emoji_id"`
	EmojiName   *string    `json:"emoji_name"`
}

// Guild Scheduled Event Object
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	ID                 Snowflake                         `json:"id"`
	GuildID            Snowflake                         `json:"guild_id"`
	ChannelID          *Snowflake                         `json:"channel_id"`
	CreatorID          *Snowflake                         `json:"creator_id"`
	Name               string                            `json:"name"`
	Description        *string                            `json:"description"`
	ScheduledStartTime time.Time                         `json:"scheduled_start_time"`
	ScheduledEndTime   time.Time                         `json:"scheduled_end_time"`
	PrivacyLevel       Flag                              `json:"privacy_level"`
	Status             Flag                              `json:"status"`
	EntityType         Flag                              `json:"entity_type"`
	EntityID           *Snowflake                         `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
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
	FlagGuildScheduledEventEntityTypeSTAGE_INSTANCE = 1
	FlagGuildScheduledEventEntityTypeVOICE          = 2
	FlagGuildScheduledEventEntityTypeEXTERNAL       = 3
)

// Guild Scheduled Event Status
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-status
const (
	FlagGuildScheduledEventStatusSCHEDULED = 1
	FlagGuildScheduledEventStatusACTIVE    = 2
	FlagGuildScheduledEventStatusCOMPLETED = 3
	FlagGuildScheduledEventStatusCANCELED  = 4
)

// Guild Scheduled Event Entity Metadata
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata
type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}

// Guild Scheduled Event User Object
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object-guild-scheduled-event-user-structure
type GuildScheduledEventUser struct {
	GuildScheduledEventID Snowflake    `json:"guild_scheduled_event_id"`
	User                  *User        `json:"user"`
	Member                *GuildMember `json:"member,omitempty"`
}

// Guild Template Object
// https://discord.com/developers/docs/resources/guild-template#guild-template-object
type GuildTemplate struct {
	Code                  string    `json:"code"`
	Name                  string    `json:"name"`
	Description           *string   `json:"description"`
	UsageCount            *int      `json:"usage_count"`
	CreatorID             Snowflake `json:"creator_id"`
	Creator               *User     `json:"creator"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	SourceGuildID         Snowflake `json:"source_guild_id"`
	SerializedSourceGuild *Guild    `json:"serialized_source_guild"`
	IsDirty               *bool      `json:"is_dirty"`
}

// Invite Object
// https://discord.com/developers/docs/resources/invite#invite-object
type Invite struct {
	Code                     string               `json:"code"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               Flag                 `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount int                  `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int                  `json:"approximate_member_count,omitempty"`
	ExpiresAt                *time.Time            `json:"expires_at"`
	StageInstance            StageInstance        `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// Invite Target Types
// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
const (
	FlagInviteTargetTypeSTREAM               = 1
	FlagInviteTargetTypeEMBEDDED_APPLICATION = 2
)

// Invite Metadata Object
// https://discord.com/developers/docs/resources/invite#invite-metadata-object-invite-metadata-structure
type InviteMetadata struct {
	Uses      *int      `json:"uses"`
	MaxUses   *int      `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
}

// Stage Instance Object
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
	ID                    Snowflake  `json:"id"`
	GuildID               *Snowflake `json:"guild_id"`
	ChannelID             *Snowflake `json:"channel_id"`
	Topic                 string     `json:"topic"`
	PrivacyLevel          Flag       `json:"privacy_level"`
	DiscoverableDisabled  bool       `json:"discoverable_disabled"`
	GuildScheduledEventID *Snowflake  `json:"guild_scheduled_event_id"`
}

// Privacy Level
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
const (
	FlagPrivacyLevelPUBLIC     = 1
	FlagPrivacyLevelGUILD_ONLY = 2
)

// Sticker Structure
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	ID          Snowflake  `json:"id"`
	PackID      Snowflake  `json:"pack_id,omitempty"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Tags        *string    `json:"tags"`
	Asset       *string    `json:"asset"`
	Type        Flag       `json:"type"`
	FormatType  Flag       `json:"format_type"`
	Available   *bool       `json:"available,omitempty"`
	GuildID     *Snowflake `json:"guild_id,omitempty"`
	User        *User      `json:"user,omitempty"`
	SortValue   *int        `json:"sort_value,omitempty"`
}

// Sticker Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
const (
	FlagStickerTypeSTANDARD = 1
	FlagStickerTypeGUILD    = 2
)

// Sticker Format Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
const (
	FlagStickerFormatTypePNG    = 1
	FlagStickerFormatTypeAPNG   = 2
	FlagStickerFormatTypeLOTTIE = 3
)

// Sticker Item Object
// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type StickerItem struct {
	ID         Snowflake `json:"id"`
	Name       string    `json:"name"`
	FormatType Flag      `json:"format_type"`
}

// Sticker Pack Object
// https://discord.com/developers/docs/resources/sticker#sticker-pack-object-sticker-pack-structure
type StickerPack struct {
	ID             Snowflake  `json:"id"`
	Stickers       []*Sticker `json:"stickers"`
	Name           string     `json:"name"`
	SKU_ID         Snowflake  `json:"sku_id"`
	CoverStickerID Snowflake  `json:"cover_sticker_id,omitempty"`
	Description    string     `json:"description"`
	BannerAssetID  Snowflake  `json:"banner_asset_id,omitempty"`
}

// User Object
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID            Snowflake `json:"id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Avatar        *string    `json:"avatar"`
	Bot           *bool      `json:"bot,omitempty"`
	System        *bool      `json:"system,omitempty"`
	MFAEnabled    *bool      `json:"mfa_enabled,omitempty"`
	Banner        *string    `json:"banner"`
	AccentColor   *int       `json:"accent_color"`
	Locale        string    `json:"locale,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         *string   `json:"email"`
	Flags         *BitFlag  `json:"flag,omitempty"`
	PremiumType   *Flag     `json:"premium_type,omitempty"`
	PublicFlags   BitFlag   `json:"public_flag,omitempty"`
}

// User Flags
// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	FlagUserNONE                         = 0
	FlagUserSTAFF                        = 1 << 0
	FlagUserPARTNER                      = 1 << 1
	FlagUserHYPESQUAD                    = 1 << 2
	FlagUserBUG_HUNTER_LEVEL_1           = 1 << 3
	FlagUserHYPESQUAD_ONLINE_HOUSE_ONE   = 1 << 6
	FlagUserHYPESQUAD_ONLINE_HOUSE_TWO   = 1 << 7
	FlagUserHYPESQUAD_ONLINE_HOUSE_THREE = 1 << 8
	FlagUserPREMIUM_EARLY_SUPPORTER      = 1 << 9
	FlagUserTEAM_PSEUDO_USER             = 1 << 10
	FlagUserBUG_HUNTER_LEVEL_2           = 1 << 14
	FlagUserVERIFIED_BOT                 = 1 << 16
	FlagUserVERIFIED_DEVELOPER           = 1 << 17
	FlagUserCERTIFIED_MODERATOR          = 1 << 18
	FlagUserBOT_HTTP_INTERACTIONS        = 1 << 19
)

// Premium Types
// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	FlagPremiumTypeNONE         = 0
	FlagPremiumTypeNITROCLASSIC = 1
	FlagPremiumTypeNITRO        = 2
)

// User Connection Object
// https://discord.com/developers/docs/resources/user#connection-object-connection-structure
type Connection struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Revoked      *bool           `json:"revoked,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
	Verified     bool           `json:"verified"`
	FriendSync   bool           `json:"friend_sync"`
	ShowActivity bool           `json:"show_activity"`
	Visibility   Flag           `json:"visibility"`
}

// Visibility Types
// https://discord.com/developers/docs/resources/user#connection-object-visibility-types
const (
	FlagVisibilityTypeNONE     = 0
	FlagVisibilityTypeEVERYONE = 1
)

// Voice State Object
// https://discord.com/developers/docs/resources/voice#voice-state-object-voice-state-structure
type VoiceState struct {
	GuildID                 Snowflake    `json:"guild_id,omitempty"`
	ChannelID               *Snowflake    `json:"channel_id"`
	UserID                  Snowflake    `json:"user_id"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              *bool         `json:"self_stream,omitempty"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp *time.Time    `json:"request_to_speak_timestamp"`
}

// Voice Region Object
// https://discord.com/developers/docs/resources/voice#voice-region-object-voice-region-structure
type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}

// Webhook Object
// https://discord.com/developers/docs/resources/webhook#webhook-object
type Webhook struct {
	ID            Snowflake  `json:"id"`
	Type          Flag       `json:"type"`
	GuildID       *Snowflake `json:"guild_id"`
	ChannelID     *Snowflake `json:"channel_id"`
	User          *User      `json:"user,omitempty"`
	Name          *string     `json:"name"`
	Avatar        *string     `json:"avatar"`
	Token         string     `json:"token,omitempty"`
	ApplicationID *Snowflake `json:"application_id"`
	SourceGuild   *Guild     `json:"source_guild,omitempty"`
	SourceChannel *Channel   `json:"source_channel,omitempty"`
	URL           *string     `json:"url,omitempty"`
}

// Webhook Types
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
const (
	FlagWebhookTypeINCOMING        = 1
	FlagWebhookTypeCHANNELFOLLOWER = 2
	FlagWebhookTypeAPPLICATION     = 3
)

// Bitwise Permission Flags
// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
const (
	FlagBitwisePermissionCREATE_INSTANT_INVITE      = 1 << 0
	FlagBitwisePermissionKICK_MEMBERS               = 1 << 1
	FlagBitwisePermissionBAN_MEMBERS                = 1 << 2
	FlagBitwisePermissionADMINISTRATOR              = 1 << 3
	FlagBitwisePermissionMANAGE_CHANNELS            = 1 << 4
	FlagBitwisePermissionMANAGE_GUILD               = 1 << 5
	FlagBitwisePermissionADD_REACTIONS              = 1 << 6
	FlagBitwisePermissionVIEW_AUDIT_LOG             = 1 << 7
	FlagBitwisePermissionPRIORITY_SPEAKER           = 1 << 8
	FlagBitwisePermissionSTREAM                     = 1 << 9
	FlagBitwisePermissionVIEW_CHANNEL               = 1 << 10
	FlagBitwisePermissionSEND_MESSAGES              = 1 << 11
	FlagBitwisePermissionSEND_TTS_MESSAGES          = 1 << 12
	FlagBitwisePermissionMANAGE_MESSAGES            = 1 << 13
	FlagBitwisePermissionEMBED_LINKS                = 1 << 14
	FlagBitwisePermissionATTACH_FILES               = 1 << 15
	FlagBitwisePermissionREAD_MESSAGE_HISTORY       = 1 << 16
	FlagBitwisePermissionMENTION_EVERYONE           = 1 << 17
	FlagBitwisePermissionUSE_EXTERNAL_EMOJIS        = 1 << 18
	FlagBitwisePermissionVIEW_GUILD_INSIGHTS        = 1 << 19
	FlagBitwisePermissionCONNECT                    = 1 << 20
	FlagBitwisePermissionSPEAK                      = 1 << 21
	FlagBitwisePermissionMUTE_MEMBERS               = 1 << 22
	FlagBitwisePermissionDEAFEN_MEMBERS             = 1 << 23
	FlagBitwisePermissionMOVE_MEMBERS               = 1 << 24
	FlagBitwisePermissionUSE_VAD                    = 1 << 25
	FlagBitwisePermissionCHANGE_NICKNAME            = 1 << 26
	FlagBitwisePermissionMANAGE_NICKNAMES           = 1 << 27
	FlagBitwisePermissionMANAGE_ROLES               = 1 << 28
	FlagBitwisePermissionMANAGE_WEBHOOKS            = 1 << 29
	FlagBitwisePermissionMANAGE_EMOJIS_AND_STICKERS = 1 << 30
	FlagBitwisePermissionUSE_APPLICATION_COMMANDS   = 1 << 31
	FlagBitwisePermissionREQUEST_TO_SPEAK           = 1 << 32
	FlagBitwisePermissionMANAGE_EVENTS              = 1 << 33
	FlagBitwisePermissionMANAGE_THREADS             = 1 << 34
	FlagBitwisePermissionCREATE_PUBLIC_THREADS      = 1 << 35
	FlagBitwisePermissionCREATE_PRIVATE_THREADS     = 1 << 36
	FlagBitwisePermissionUSE_EXTERNAL_STICKERS      = 1 << 37
	FlagBitwisePermissionSEND_MESSAGES_IN_THREADS   = 1 << 38
	FlagBitwisePermissionUSE_EMBEDDED_ACTIVITIES    = 1 << 39
	FlagBitwisePermissionMODERATE_MEMBERS           = 1 << 40
)

// Permission Overwrite Types
const (
	FlagPermissionOverwriteTypeRole   = 0
	FlagPermissionOverwriteTypeMember = 1
)

// Role Object
// https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID           Snowflake `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         *string    `json:"icon,omitempty"`
	UnicodeEmoji *string    `json:"unicode_emoji,omitempty"`
	Position     int       `json:"position"`
	Permissions  string    `json:"permissions"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         *RoleTags `json:"tags,omitempty"`
}

// Role Tags Structure
// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID             Snowflake `json:"bot_id,omitempty"`
	IntegrationID     Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber bool      `json:"premium_subscriber,omitempty"`
}

// Team Object
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        *string        `json:"icon"`
	ID          Snowflake     `json:"id"`
	Members     []*TeamMember `json:"members"`
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	OwnerUserID Snowflake     `json:"owner_user_id"`
}

// Team Member Object
// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState Flag      `json:"membership_state"`
	Permissions     []string  `json:"permissions"`
	TeamID          Snowflake `json:"team_id"`
	User            *User     `json:"user"`
}

// Membership State Enum
// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
const (
	FlagMembershipStateEnumINVITED  = 1
	FlagMembershipStateEnumACCEPTED = 2
)

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
	Name          string              `json:"name"`
	Type          *Flag               `json:"type"`
	URL           *string              `json:"url"`
	CreatedAt     int                 `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID Snowflake           `json:"application_id,omitempty"`
	Details       *string              `json:"details"`
	State         *string              `json:"state"`
	Emoji         *Emoji              `json:"emoji"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      *bool                `json:"instance,omitempty"`
	Flags         BitFlag             `json:"flags,omitempty"`
	Buttons       []Button            `json:"buttons,omitempty"`
}

// Activity Types
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
const (
	FlagActivityTypePlaying   = 0
	FlagActivityTypeStreaming = 1
	FlagActivityTypeListening = 2
	FlagActivityTypeWatching  = 3
	FlagActivityTypeCustom    = 4
	FlagActivityTypeCompeting = 5
)

// Activity Timestamps Struct
// htthttps://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type ActivityTimestamps struct {
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// Activity Emoji
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
type ActivityEmoji struct {
	Name     string    `json:"name"`
	ID       Snowflake `json:"id,omitempty"`
	Animated *bool      `json:"animated,omitempty"`
}

// Activity Party Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type ActivityParty struct {
	ID   *string  `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

// Activity Assets Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
type ActivityAssets struct {
	LargeImage *string `json:"large_image,omitempty"`
	LargeText  *string `json:"large_text,omitempty"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

// Activity Asset Image
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-asset-image
type ActivityAssetImage struct {
	ApplicationAsset string `json:"application_asset_id"`
	MediaProxyImage  string `json:"image_id"`
}

// Activity Secrets Struct
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
type ActivitySecrets struct {
	Join     *string `json:"join,omitempty"`
	Spectate *string `json:"spectate,omitempty"`
	Match    *string `json:"match,omitempty"`
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

// OAuth2 Scopes
// https://discord.com/developers/docs/topics/oauth2#shared-resources-oauth2-scopes
const (
	FlagOAuth2ScopeActivitiesRead                        = "activities.read"
	FlagOAuth2ScopeActivitiesWrite                       = "activities.write"
	FlagOAuth2ScopeApplicationsBuildsRead                = "applications.builds.read"
	FlagOAuth2ScopeApplicationsBuildsUpload              = "applications.builds.upload"
	FlagOAuth2ScopeApplicationsCommands                  = "applications.commands"
	FlagOAuth2ScopeApplicationsCommandsUpdate            = "applications.commands.update"
	FlagOAuth2ScopeApplicationsCommandsPermissionsUpdate = "applications.commands.permissions.update"
	FlagOAuth2ScopeApplicationsEntitlements              = "applications.entitlements"
	FlagOAuth2ScopeApplicationsStoreUpdate               = "applications.store.update"
	FlagOAuth2ScopeBot                                   = "bot"
	FlagOAuth2ScopeConnections                           = "connections"
	FlagOAuth2ScopeDM_channelsRead                       = "dm_channels.read"
	FlagOAuth2ScopeEmail                                 = "email"
	FlagOAuth2ScopeGDMJoin                               = "gdm.join"
	FlagOAuth2ScopeGuilds                                = "guilds"
	FlagOAuth2ScopeGuildsJoin                            = "guilds.join"
	FlagOAuth2ScopeGuildsMembersRead                     = "guilds.members.read"
	FlagOAuth2ScopeIdentify                              = "identify"
	FlagOAuth2ScopeMessagesRead                          = "messages.read"
	FlagOAuth2ScopeRelationshipsRead                     = "relationships.read"
	FlagOAuth2ScopeRPC                                   = "rpc"
	FlagOAuth2ScopeRPCActivitiesWrite                    = "rpc.activities.write"
	FlagOAuth2ScopeRPCNotificationsRead                  = "rpc.notifications.read"
	FlagOAuth2ScopeRPCVoiceRead                          = "rpc.voice.read"
	FlagOAuth2ScopeRPCVoiceWrite                         = "rpc.voice.write"
	FlagOAuth2ScopeVoice                                 = "voice"
	FlagOAuth2ScopeWebhookIncoming                       = "webhook.incoming"
)
