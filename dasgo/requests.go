// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import (
	"time"
)

// Get Global Application Commands
// GET /applications/{application.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#get-global-application-commands
type GetGlobalApplicationCommands struct {
	ApplicationID     Snowflake
	WithLocalizations bool `json:"with_localizations,omitempty"`
}

// Create Global Application Command
// POST /applications/{application.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#create-global-application-command
type CreateGlobalApplicationCommand struct {
	ApplicationID            Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
	Type                     Flag                        `json:"type,omitempty"`
}

// Get Global Application Command
// GET /applications/{application.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#get-global-application-command
type GetGlobalApplicationCommand struct {
	ApplicationID Snowflake
	CommandID     Snowflake
}

// Edit Global Application Command
// PATCH /applications/{application.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#edit-global-application-command
type EditGlobalApplicationCommand struct {
	ApplicationID            Snowflake
	CommandID                Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
}

// Delete Global Application Command
// DELETE /applications/{application.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#delete-global-application-command
type DeleteGlobalApplicationCommand struct {
	ApplicationID Snowflake
	CommandID     Snowflake
}

// Bulk Overwrite Global Application Commands
// PUT /applications/{application.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#bulk-overwrite-global-application-commands
type BulkOverwriteGlobalApplicationCommands struct {
	ApplicationID       Snowflake
	ApplicationCommands []*ApplicationCommand
}

// Get Guild Application Commands
// GET /applications/{application.id}/guilds/{guild.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#get-guild-application-commands
type GetGuildApplicationCommands struct {
	ApplicationID     Snowflake
	GuildID           Snowflake
	WithLocalizations bool `json:"with_localizations,omitempty"`
}

// Create Guild Application Command
// POST /applications/{application.id}/guilds/{guild.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#create-guild-application-command
type CreateGuildApplicationCommand struct {
	ApplicationID            Snowflake
	GuildID                  Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
	Type                     Flag                        `json:"type,omitempty"`
}

// Get Guild Application Command
// GET /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#get-guild-application-command
type GetGuildApplicationCommand struct {
	ApplicationID Snowflake
	GuildID       Snowflake
	CommandID     Snowflake
}

// Edit Guild Application Command
// PATCH /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#edit-guild-application-command
type EditGuildApplicationCommand struct {
	ApplicationID            Snowflake
	GuildID                  Snowflake
	CommandID                Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
}

// Delete Guild Application Command
// DELETE /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#delete-guild-application-command
type DeleteGuildApplicationCommand struct {
	ApplicationID Snowflake
	GuildID       Snowflake
	CommandID     Snowflake
}

// Bulk Overwrite Guild Application Commands
// PUT /applications/{application.id}/guilds/{guild.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#bulk-overwrite-guild-application-commands
type BulkOverwriteGuildApplicationCommands struct {
	ApplicationID            Snowflake
	GuildID                  Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
	Type                     Flag                        `json:"type,omitempty"`
}

// Get Guild Application Command Permissions
// GET /applications/{application.id}/guilds/{guild.id}/commands/permissions
// https://discord.com/developers/docs/interactions/application-commands#get-guild-application-command-permissions
type GetGuildApplicationCommandPermissions struct {
	ApplicationID Snowflake
	GuildID       Snowflake
}

// Get Application Command Permissions
// GET /applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions
// https://discord.com/developers/docs/interactions/application-commands#get-application-command-permissions
type GetApplicationCommandPermissions struct {
	ApplicationID Snowflake
	GuildID       Snowflake
	CommandID     Snowflake
}

// Edit Application Command Permissions
// PUT /applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions
// https://discord.com/developers/docs/interactions/application-commands#edit-application-command-permissions
type EditApplicationCommandPermissions struct {
	ApplicationID Snowflake
	GuildID       Snowflake
	CommandID     Snowflake
	Permissions   []*ApplicationCommandPermissions `json:"permissions,omitempty"`
}

// Batch Edit Application Command Permissions
// PUT /applications/{application.id}/guilds/{guild.id}/commands/permissions
// https://discord.com/developers/docs/interactions/application-commands#batch-edit-application-command-permissions
type BatchEditApplicationCommandPermissions struct {
	ApplicationID Snowflake
	GuildID       Snowflake
}

// Get Guild Audit Log
// GET /guilds/{guild.id}/audit-logs
// https://discord.com/developers/docs/resources/audit-log#get-guild-audit-log
type GetGuildAuditLog struct {
	GuildID    Snowflake
	UserID     Snowflake `json:"user_id"`
	ActionType Flag      `json:"action_type"`
	Before     Snowflake `json:"before,omitempty"`
	Limit      Flag      `json:"limit,omitempty"`
}

// Get Channel
// GET /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#get-channel
type GetChannel struct {
	ChannelID Snowflake
}

// Modify Channel
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel
type ModifyChannel struct {
	ChannelID Snowflake
}

// Modify Channel
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-group-dm
type ModifyChannelGroupDM struct {
	ChannelID Snowflake
	Name      string `json:"name,omitempty"`
	Icon      int    `json:"icon,omitempty"`
}

// Modify Channel
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-guild-channel
type ModifyChannelGuild struct {
	ChannelID                  Snowflake
	Name                       *string                `json:"name,omitempty"`
	Type                       *Flag                  `json:"type,omitempty"`
	Position                   *uint                  `json:"position,omitempty"`
	Topic                      *string                `json:"topic,omitempty"`
	NSFW                       bool                   `json:"nsfw,omitempty"`
	RateLimitPerUser           *CodeFlag              `json:"rate_limit_per_user,omitempty"`
	Bitrate                    *uint                  `json:"bitrate,omitempty"`
	UserLimit                  *Flag                  `json:"user_limit,omitempty"`
	PermissionOverwrites       *[]PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                   *Snowflake             `json:"parent_id,omitempty"`
	RTCRegion                  *string                `json:"rtc_region,omitempty"`
	VideoQualityMode           Flag                   `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration *uint                  `json:"default_auto_archive_duration,omitempty"`
}

// Modify Channel
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-thread
type ModifyChannelThread struct {
	ChannelID           Snowflake
	Name                string    `json:"name,omitempty"`
	Archived            bool      `json:"archived,omitempty"`
	AutoArchiveDuration CodeFlag  `json:"auto_archive_duration,omitempty"`
	Locked              bool      `json:"locked,omitempty"`
	Invitable           bool      `json:"invitable,omitempty"`
	RateLimitPerUser    *CodeFlag `json:"rate_limit_per_user,omitempty"`
}

// Delete/Close Channel
// DELETE /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#deleteclose-channel
type DeleteCloseChannel struct {
	ChannelID Snowflake
}

// Get Channel Messages
// GET /channels/{channel.id}/messages
// https://discord.com/developers/docs/resources/channel#get-channel-messages
type GetChannelMessages struct {
	ChannelID Snowflake
	Around    *Snowflake `json:"around,omitempty"`
	Before    *Snowflake `json:"before,omitempty"`
	After     *Snowflake `json:"after,omitempty"`
	Limit     Flag       `json:"limit,omitempty"`
}

// Get Channel Message
// GET /channels/{channel.id}/messages/{message.id}
// https://discord.com/developers/docs/resources/channel#get-channel-message
type GetChannelMessage struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Create Message
// POST /channels/{channel.id}/messages
// https://discord.com/developers/docs/resources/channel#create-message
type CreateMessage struct {
	ChannelID       Snowflake
	Content         string            `json:"content,omitempty"`
	TTS             bool              `json:"tts,omitempty"`
	Embeds          []*Embed          `json:"embeds,omitempty"`
	Embed           *Embed            `json:"embed,omitempty"`
	AllowedMentions *AllowedMentions  `json:"allowed_mentions,omitempty"`
	Reference       *MessageReference `json:"message_reference,omitempty"`
	StickerID       []*Snowflake      `json:"sticker_ids,omitempty"`
	Components      []*Component      `json:"components,omitempty"`
	Files           []byte            `dasgo:"files"`
	PayloadJSON     *string           `json:"payload_json,omitempty"`
	Attachments     []*Attachment     `json:"attachments,omitempty"`
	Flags           BitFlag           `json:"flags,omitempty"`
}

// Crosspost Message
// POST /channels/{channel.id}/messages/{message.id}/crosspost
// https://discord.com/developers/docs/resources/channel#crosspost-message
type CrosspostMessage struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Create Reaction
// PUT /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me
// https://discord.com/developers/docs/resources/channel#create-reaction
type CreateReaction struct {
	ChannelID Snowflake
	MessageID Snowflake
	Emoji     string
}

// Delete Own Reaction
// DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me
// https://discord.com/developers/docs/resources/channel#delete-own-reaction
type DeleteOwnReaction struct {
	ChannelID Snowflake
	MessageID Snowflake
	Emoji     string
}

// Delete User Reaction
// DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/{user.id}
// https://discord.com/developers/docs/resources/channel#delete-user-reaction
type DeleteUserReaction struct {
	ChannelID Snowflake
	MessageID Snowflake
	Emoji     string
	UserID    Snowflake
}

// Get Reactions
// GET /channels/{channel.id}/messages/{message.id}/reactions/{emoji}
// https://discord.com/developers/docs/resources/channel#get-reactions
type GetReactions struct {
	ChannelID Snowflake
	MessageID Snowflake
	Emoji     string
	After     Snowflake `json:"after,omitempty"`
	Limit     Flag      `json:"limit,omitempty"`
}

// Delete All Reactions
// DELETE /channels/{channel.id}/messages/{message.id}/reactions
// https://discord.com/developers/docs/resources/channel#delete-all-reactions
type DeleteAllReactions struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Delete All Reactions for Emoji
// DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}
// https://discord.com/developers/docs/resources/channel#delete-all-reactions-for-emoji
type DeleteAllReactionsforEmoji struct {
	ChannelID Snowflake
	MessageID Snowflake
	Emoji     string
}

// Edit Message
// PATCH /channels/{channel.id}/messages/{message.id}
// https://discord.com/developers/docs/resources/channel#edit-message
type EditMessage struct {
	ChannelID       Snowflake
	MessageID       Snowflake
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	Embed           *Embed           `json:"embed,omitempty"`
	Flags           *BitFlag         `json:"flags,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []*Component     `json:"components,omitempty"`
	Files           []byte           `dasgo:"files"`
	PayloadJSON     *string          `json:"payload_json,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// Delete Message
// DELETE /channels/{channel.id}/messages/{message.id}
// https://discord.com/developers/docs/resources/channel#delete-message
type DeleteMessage struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Bulk Delete Messages
// POST /channels/{channel.id}/messages/bulk-delete
// https://discord.com/developers/docs/resources/channel#bulk-delete-messages
type BulkDeleteMessages struct {
	ChannelID Snowflake
	Messages  []*Snowflake `json:"messages,omitempty"`
}

// Edit Channel Permissions
// PUT /channels/{channel.id}/permissions/{overwrite.id}
// https://discord.com/developers/docs/resources/channel#edit-channel-permissions
type EditChannelPermissions struct {
	ChannelID   Snowflake
	OverwriteID Snowflake
	Allow       string `json:"allow,omitempty"`
	Deny        string `json:"deny,omitempty"`
	Type        *Flag  `json:"type,omitempty"`
}

// Get Channel Invites
// GET /channels/{channel.id}/invites
// https://discord.com/developers/docs/resources/channel#get-channel-invites
type GetChannelInvites struct {
	ChannelID Snowflake
}

// Create Channel Invite
// POST /channels/{channel.id}/invites
// https://discord.com/developers/docs/resources/channel#create-channel-invite
type CreateChannelInvite struct {
	ChannelID           Snowflake
	MaxAge              *int      `json:"max_age,omitempty"`
	MaxUses             *Flag     `json:"max_uses,omitempty"`
	Temporary           bool      `json:"temporary,omitempty"`
	Unique              bool      `json:"unique,omitempty"`
	TargetType          Flag      `json:"target_type,omitempty"`
	TargetUserID        Snowflake `json:"target_user_id,omitempty"`
	TargetApplicationID Snowflake `json:"target_application_id,omitempty"`
}

// Delete Channel Permission
// DELETE /channels/{channel.id}/permissions/{overwrite.id}
// https://discord.com/developers/docs/resources/channel#delete-channel-permission
type DeleteChannelPermission struct {
	ChannelID   Snowflake
	OverwriteID Snowflake
}

// Follow News Channel
// POST /channels/{channel.id}/followers
// https://discord.com/developers/docs/resources/channel#follow-news-channel
type FollowNewsChannel struct {
	ChannelID Snowflake
}

// Trigger Typing Indicator
// POST /channels/{channel.id}/typing
// https://discord.com/developers/docs/resources/channel#trigger-typing-indicator
type TriggerTypingIndicator struct {
	ChannelID Snowflake
}

// Get Pinned Messages
// GET /channels/{channel.id}/pins
// https://discord.com/developers/docs/resources/channel#get-pinned-messages
type GetPinnedMessages struct {
	ChannelID Snowflake
}

// Pin Message
// PUT /channels/{channel.id}/pins/{message.id}
// https://discord.com/developers/docs/resources/channel#pin-message
type PinMessage struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Unpin Message
// DELETE /channels/{channel.id}/pins/{message.id}
// https://discord.com/developers/docs/resources/channel#unpin-message
type UnpinMessage struct {
	ChannelID Snowflake
	MessageID Snowflake
}

// Group DM Add Recipient
// PUT /channels/{channel.id}/recipients/{user.id}
// https://discord.com/developers/docs/resources/channel#group-dm-add-recipient
type GroupDMAddRecipient struct {
	ChannelID   Snowflake
	UserID      Snowflake
	AccessToken string  `json:"access_token,omitempty"`
	Nickname    *string `json:"nick,omitempty"`
}

// Group DM Remove Recipient
// DELETE /channels/{channel.id}/recipients/{user.id}
// https://discord.com/developers/docs/resources/channel#group-dm-remove-recipient
type GroupDMRemoveRecipient struct {
	ChannelID Snowflake
	UserID    Snowflake
}

// Start Thread from Message
// POST /channels/{channel.id}/messages/{message.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-from-message
type StartThreadfromMessage struct {
	ChannelID           Snowflake
	MessageID           Snowflake
	Name                string `json:"name,omitempty"`
	RateLimitPerUser    uint   `json:"rate_limit_per_user,omitempty"`
	AutoArchiveDuration *int   `json:"auto_archive_duration,omitempty"`
}

// Start Thread without Message
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-without-message
type StartThreadwithoutMessage struct {
	ChannelID           Snowflake
	Name                string    `json:"name,omitempty"`
	AutoArchiveDuration CodeFlag  `json:"auto_archive_duration,omitempty"`
	Type                *Flag     `json:"type,omitempty"`
	Invitable           bool      `json:"invitable,omitempty"`
	RateLimitPerUser    *CodeFlag `json:"rate_limit_per_user,omitempty"`
}

// Start Thread in Forum Channel
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel
type StartThreadinForumChannel struct {
	ChannelID Snowflake
}

// Start Thread in Forum Channel
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-json-params-for-the-thread
type StartThreadinForumChannelThread struct {
	ChannelID           Snowflake
	Name                string    `json:"name,omitempty"`
	RateLimitPerUser    uint      `json:"rate_limit_per_user,omitempty"`
	AutoArchiveDuration *CodeFlag `json:"auto_archive_duration,omitempty"`
}

// Start Thread in Forum Channel
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-json-params-for-the-message
type StartThreadinForumChannelMessage struct {
	ChannelID       Snowflake
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []*Component     `json:"components,omitempty"`
	StickerIDS      []*Snowflake     `json:"sticker_ids,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
	Files           []byte           `dasgo:"files"`
	PayloadJSON     string           `json:"payload_json,omitempty"`
	Flags           BitFlag          `json:"flags,omitempty"`
}

// Join Thread
// PUT /channels/{channel.id}/thread-members/@me
// https://discord.com/developers/docs/resources/channel#join-thread
type JoinThread struct {
	ChannelID Snowflake
}

// Add Thread Member
// PUT /channels/{channel.id}/thread-members/{user.id}
// https://discord.com/developers/docs/resources/channel#add-thread-member
type AddThreadMember struct {
	ChannelID Snowflake
	UserID    Snowflake
}

// Leave Thread
// DELETE /channels/{channel.id}/thread-members/@me
// https://discord.com/developers/docs/resources/channel#leave-thread
type LeaveThread struct {
	ChannelID Snowflake
}

// Remove Thread Member
// DELETE /channels/{channel.id}/thread-members/{user.id}
// https://discord.com/developers/docs/resources/channel#remove-thread-member
type RemoveThreadMember struct {
	ChannelID Snowflake
	UserID    Snowflake
}

// Get Thread Member
// GET /channels/{channel.id}/thread-members/{user.id}
// https://discord.com/developers/docs/resources/channel#get-thread-member
type GetThreadMember struct {
	ChannelID Snowflake
	UserID    Snowflake
}

// List Thread Members
// GET /channels/{channel.id}/thread-members
// https://discord.com/developers/docs/resources/channel#list-thread-members
type ListThreadMembers struct {
	ChannelID Snowflake
}

// List Active Channel Threads
// GET /channels/{channel.id}/threads/active
// https://discord.com/developers/docs/resources/channel#list-active-threads
type ListActiveChannelThreads struct {
	ChannelID Snowflake
	Before    Snowflake `json:"before,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

// List Public Archived Threads
// GET /channels/{channel.id}/threads/archived/public
// https://discord.com/developers/docs/resources/channel#list-public-archived-threads
type ListPublicArchivedThreads struct {
	ChannelID Snowflake
	Before    Snowflake `json:"before,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

// List Private Archived Threads
// GET /channels/{channel.id}/threads/archived/private
// https://discord.com/developers/docs/resources/channel#list-private-archived-threads
type ListPrivateArchivedThreads struct {
	ChannelID Snowflake
	Before    Snowflake `json:"before,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

// List Joined Private Archived Threads
// GET /channels/{channel.id}/users/@me/threads/archived/private
// https://discord.com/developers/docs/resources/channel#list-joined-private-archived-threads
type ListJoinedPrivateArchivedThreads struct {
	ChannelID Snowflake
	Before    Snowflake `json:"before,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

// List Guild Emojis
// GET /guilds/{guild.id}/emojis
// https://discord.com/developers/docs/resources/emoji#list-guild-emojis
type ListGuildEmojis struct {
	GuildID Snowflake
}

// Get Guild Emoji
// GET /guilds/{guild.id}/emojis/{emoji.id}
// https://discord.com/developers/docs/resources/emoji#get-guild-emoji
type GetGuildEmoji struct {
	GuildID Snowflake
	EmojiID Snowflake
}

// Create Guild Emoji
// POST /guilds/{guild.id}/emojis
// https://discord.com/developers/docs/resources/emoji#create-guild-emoji
type CreateGuildEmoji struct {
	GuildID Snowflake
	Name    string       `json:"name,omitempty"`
	Image   string       `json:"image,omitempty"`
	Roles   []*Snowflake `json:"roles,omitempty"`
}

// Modify Guild Emoji
// PATCH /guilds/{guild.id}/emojis/{emoji.id}
// https://discord.com/developers/docs/resources/emoji#modify-guild-emoji
type ModifyGuildEmoji struct {
	GuildID Snowflake
	EmojiID Snowflake
	Name    string       `json:"name,omitempty"`
	Roles   []*Snowflake `json:"roles,omitempty"`
}

// Delete Guild Emoji
// DELETE /guilds/{guild.id}/emojis/{emoji.id}
// https://discord.com/developers/docs/resources/emoji#delete-guild-emoji
type DeleteGuildEmoji struct {
	GuildID Snowflake
	EmojiID Snowflake
}

// Get Gateway
// GET /gateway
// https://discord.com/developers/docs/topics/gateway#get-gateway
type GetGateway struct{}

// Get Gateway Bot
// GET /gateway/bot
// https://discord.com/developers/docs/topics/gateway#get-gateway-bot
type GetGatewayBot struct {
	URL               string `json:"url,omitempty"`
	Shards            int    `json:"shards,omitempty"`
	SessionStartLimit int    `json:"session_start_limit,omitempty"`
}

// Create Guild
// POST /guilds
// https://discord.com/developers/docs/resources/guild#create-guild
type CreateGuild struct {
	Name                        string     `json:"name,omitempty"`
	Region                      string     `json:"region,omitempty"`
	Icon                        string     `json:"icon,omitempty"`
	VerificationLevel           *Flag      `json:"verification_level,omitempty"`
	DefaultMessageNotifications *Flag      `json:"default_message_notifications,omitempty"`
	AfkChannelID                string     `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int        `json:"afk_timeout,omitempty"`
	OwnerID                     string     `json:"owner_id,omitempty"`
	Splash                      string     `json:"splash,omitempty"`
	Banner                      string     `json:"banner,omitempty"`
	Roles                       []*Role    `json:"roles,omitempty"`
	Channels                    []*Channel `json:"channels,omitempty"`
	SystemChannelID             Snowflake  `json:"system_channel_id,omitempty"`
	SystemChannelFlags          BitFlag    `json:"system_channel_flags,omitempty"`
	ExplicitContentFilter       *Flag      `json:"explicit_content_filter,omitempty"`
}

// Get Guild
// GET /guilds/{guild.id}
// https://discord.com/developers/docs/resources/guild#get-guild
type GetGuild struct {
	GuildID    Snowflake
	WithCounts bool `json:"with_counts,omitempty"`
}

// Get Guild Preview
// GET /guilds/{guild.id}/preview
// https://discord.com/developers/docs/resources/guild#get-guild-preview
type GetGuildPreview struct {
	GuildID Snowflake
}

// Modify Guild
// PATCH /guilds/{guild.id}
// https://discord.com/developers/docs/resources/guild#modify-guild
type ModifyGuild struct {
	GuildID                     Snowflake
	Name                        string     `json:"name,omitempty"`
	Region                      string     `json:"region,omitempty"`
	VerificationLevel           *Flag      `json:"verification_lvl,omitempty"`
	DefaultMessageNotifications *Flag      `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *Flag      `json:"explicit_content_filter,omitempty"`
	AFKChannelID                Snowflake  `json:"afk_channel_id,omitempty"`
	Icon                        *string    `json:"icon,omitempty"`
	OwnerID                     Snowflake  `json:"owner_id,omitempty"`
	Splash                      *string    `json:"splash,omitempty"`
	DiscoverySplash             *string    `json:"discovery_splash,omitempty"`
	Banner                      *string    `json:"banner,omitempty"`
	SystemChannelID             Snowflake  `json:"system_channel_id,omitempty"`
	SystemChannelFlags          BitFlag    `json:"system_channel_flags,omitempty"`
	RulesChannelID              Snowflake  `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *Snowflake `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *string    `json:"preferred_locale,omitempty"`
	Features                    []*string  `json:"features,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	PremiumProgressBarEnabled   bool       `json:"premium_progress_bar_enabled,omitempty"`
}

// Delete Guild
// DELETE /guilds/{guild.id}
// https://discord.com/developers/docs/resources/guild#delete-guild
type DeleteGuild struct {
	GuildID Snowflake
}

// Get Guild Channels
// GET /guilds/{guild.id}/channels
// https://discord.com/developers/docs/resources/guild#get-guild-channels
type GetGuildChannels struct {
	GuildID Snowflake
}

// Create Guild Channel
// POST /guilds/{guild.id}/channels
// https://discord.com/developers/docs/resources/guild#create-guild-channel
type CreateGuildChannel struct {
	GuildID                    Snowflake
	Name                       string                 `json:"name,omitempty"`
	Type                       *Flag                  `json:"type,omitempty"`
	Topic                      *string                `json:"topic,omitempty"`
	NSFW                       bool                   `json:"nsfw,omitempty"`
	Position                   int                    `json:"position,omitempty"`
	Bitrate                    int                    `json:"bitrate,omitempty"`
	UserLimit                  int                    `json:"user_limit,omitempty"`
	PermissionOverwrites       []*PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                   *Snowflake             `json:"parent_id,omitempty"`
	RateLimitPerUser           *CodeFlag              `json:"rate_limit_per_user,omitempty"`
	DefaultAutoArchiveDuration int                    `json:"default_auto_archive_duration,omitempty"`
}

// Modify Guild Channel Positions
// PATCH /guilds/{guild.id}/channels
// https://discord.com/developers/docs/resources/guild#modify-guild-channel-positions
type ModifyGuildChannelPositions struct {
	GuildID         Snowflake
	ID              Snowflake  `json:"id,omitempty"`
	Position        int        `json:"position,omitempty"`
	LockPermissions bool       `json:"lock_permissions,omitempty"`
	ParentID        *Snowflake `json:"parent_id,omitempty"`
}

// List Active Guild Threads
// GET /guilds/{guild.id}/threads/active
// https://discord.com/developers/docs/resources/guild#list-active-threads
type ListActiveGuildThreads struct {
	GuildID Snowflake
	Threads []*Channel      `json:"threads,omitempty"`
	Members []*ThreadMember `json:"members,omitempty"`
}

// Get Guild Member
// GET /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#get-guild-member
type GetGuildMember struct {
	GuildID Snowflake
	UserID  Snowflake
}

// List Guild Members
// GET /guilds/{guild.id}/members
// https://discord.com/developers/docs/resources/guild#list-guild-members
type ListGuildMembers struct {
	GuildID Snowflake
	After   *Snowflake `json:"after,omitempty"`
	Limit   *CodeFlag  `json:"limit,omitempty"`
}

// Search Guild Members
// GET /guilds/{guild.id}/members/search
// https://discord.com/developers/docs/resources/guild#search-guild-members
type SearchGuildMembers struct {
	GuildID Snowflake
	Query   string    `json:"query,omitempty"`
	Limit   *CodeFlag `json:"limit,omitempty"`
}

// Add Guild Member
// PUT /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#add-guild-member
type AddGuildMember struct {
	GuildID     Snowflake
	UserID      Snowflake
	AccessToken string       `json:"access_token,omitempty"`
	Nick        string       `json:"nick,omitempty"`
	Roles       []*Snowflake `json:"roles,omitempty"`
	Mute        bool         `json:"mute,omitempty"`
	Deaf        bool         `json:"deaf,omitempty"`
}

// Modify Guild Member
// PATCH /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#modify-guild-member
type ModifyGuildMember struct {
	GuildID                    Snowflake
	UserID                     Snowflake
	Nick                       string       `json:"nick,omitempty"`
	Roles                      []*Snowflake `json:"roles,omitempty"`
	Mute                       bool         `json:"mute,omitempty"`
	Deaf                       bool         `json:"deaf,omitempty"`
	ChannelID                  Snowflake    `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time   `json:"communication_disabled_until,omitempty"`
}

// Modify Current Member
// PATCH /guilds/{guild.id}/members/@me
// https://discord.com/developers/docs/resources/guild#modify-current-member
type ModifyCurrentMember struct {
	GuildID Snowflake
	Nick    string `json:"nick,omitempty"`
}

// Modify Current User Nick
// PATCH /guilds/{guild.id}/members/@me/nick
// https://discord.com/developers/docs/resources/guild#modify-current-user-nick
type ModifyCurrentUserNick struct {
	GuildID Snowflake
	Nick    string `json:"nick,omitempty"`
}

// Add Guild Member Role
// PUT /guilds/{guild.id}/members/{user.id}/roles/{role.id}
// https://discord.com/developers/docs/resources/guild#add-guild-member-role
type AddGuildMemberRole struct {
	GuildID Snowflake
	UserID  Snowflake
	RoleID  Snowflake
}

// Remove Guild Member Role
// DELETE /guilds/{guild.id}/members/{user.id}/roles/{role.id}
// https://discord.com/developers/docs/resources/guild#remove-guild-member-role
type RemoveGuildMemberRole struct {
	GuildID Snowflake
	UserID  Snowflake
	RoleID  Snowflake
}

// Remove Guild Member
// DELETE /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#remove-guild-member
type RemoveGuildMember struct {
	GuildID Snowflake
	UserID  Snowflake
}

// Get Guild Bans
// GET /guilds/{guild.id}/bans
// https://discord.com/developers/docs/resources/guild#get-guild-bans
type GetGuildBans struct {
	GuildID Snowflake
	Before  *Snowflake `json:"before,omitempty"`
	After   *Snowflake `json:"after,omitempty"`
	Limit   *CodeFlag  `json:"limit,omitempty"`
}

// Get Guild Ban
// GET /guilds/{guild.id}/bans/{user.id}
// https://discord.com/developers/docs/resources/guild#get-guild-ban
type GetGuildBan struct {
	GuildID Snowflake
	UserID  Snowflake
}

// Create Guild Ban
// PUT /guilds/{guild.id}/bans/{user.id}
// https://discord.com/developers/docs/resources/guild#create-guild-ban
type CreateGuildBan struct {
	GuildID           Snowflake
	UserID            Snowflake
	DeleteMessageDays *Flag   `json:"delete_message_days,omitempty"`
	Reason            *string `json:"reason,omitempty"`
}

// Remove Guild Ban
// DELETE /guilds/{guild.id}/bans/{user.id}
// https://discord.com/developers/docs/resources/guild#remove-guild-ban
type RemoveGuildBan struct {
	GuildID Snowflake
	UserID  Snowflake
}

// Get Guild Roles
// GET /guilds/{guild.id}/roles
// https://discord.com/developers/docs/resources/guild#get-guild-roles
type GetGuildRoles struct {
	GuildID Snowflake
}

// Create Guild Role
// POST /guilds/{guild.id}/roles
// https://discord.com/developers/docs/resources/guild#create-guild-role
type CreateGuildRole struct {
	GuildID      Snowflake
	Name         string  `json:"name,omitempty"`
	Permissions  string  `json:"permissions,omitempty"`
	Color        *int    `json:"color,omitempty"`
	Hoist        bool    `json:"hoist,omitempty"`
	Icon         *int    `json:"icon,omitempty"`
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	Mentionable  bool    `json:"mentionable,omitempty"`
}

// Modify Guild Role Positions
// PATCH /guilds/{guild.id}/roles
// https://discord.com/developers/docs/resources/guild#modify-guild-role-positions
type ModifyGuildRolePositions struct {
	GuildID  Snowflake
	ID       Snowflake `json:"id,omitempty"`
	Position int       `json:"position,omitempty"`
}

// Modify Guild Role
// PATCH /guilds/{guild.id}/roles/{role.id}
// https://discord.com/developers/docs/resources/guild#modify-guild-role
type ModifyGuildRole struct {
	GuildID      Snowflake
	RoleID       Snowflake
	Name         string  `json:"name,omitempty"`
	Permissions  int64   `json:"permissions,string,omitempty"`
	Color        *int    `json:"color,omitempty"`
	Hoist        bool    `json:"hoist,omitempty"`
	Icon         *int    `json:"icon,omitempty"`
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	Mentionable  bool    `json:"mentionable,omitempty"`
}

// Delete Guild Role
// DELETE /guilds/{guild.id}/roles/{role.id}
// https://discord.com/developers/docs/resources/guild#delete-guild-role
type DeleteGuildRole struct {
	GuildID Snowflake
	RoleID  Snowflake
}

// Get Guild Prune Count
// GET /guilds/{guild.id}/prune
// https://discord.com/developers/docs/resources/guild#get-guild-prune-count
type GetGuildPruneCount struct {
	GuildID      Snowflake
	Days         Flag         `json:"days,omitempty"`
	IncludeRoles []*Snowflake `json:"include_roles,omitempty"`
}

// Begin Guild Prune
// POST /guilds/{guild.id}/prune
// https://discord.com/developers/docs/resources/guild#begin-guild-prune
type BeginGuildPrune struct {
	GuildID           Snowflake
	Days              Flag         `json:"days,omitempty"`
	ComputePruneCount bool         `json:"compute_prune_count,omitempty"`
	IncludeRoles      []*Snowflake `json:"include_roles,omitempty"`
	Reason            *string      `json:"reason,omitempty"`
}

// Get Guild Voice Regions
// GET /guilds/{guild.id}/regions
// https://discord.com/developers/docs/resources/guild#get-guild-voice-regions
type GetGuildVoiceRegions struct {
	GuildID Snowflake
}

// Get Guild Invites
// GET /guilds/{guild.id}/invites
// https://discord.com/developers/docs/resources/guild#get-guild-invites
type GetGuildInvites struct {
	GuildID Snowflake
}

// Get Guild Integrations
// GET /guilds/{guild.id}/integrations
// https://discord.com/developers/docs/resources/guild#get-guild-integrations
type GetGuildIntegrations struct {
	GuildID Snowflake
}

// Delete Guild Integration
// DELETE /guilds/{guild.id}/integrations/{integration.id}
// https://discord.com/developers/docs/resources/guild#delete-guild-integration
type DeleteGuildIntegration struct {
	GuildID       Snowflake
	IntegrationID Snowflake
}

// Get Guild Widget Settings
// GET /guilds/{guild.id}/widget
// https://discord.com/developers/docs/resources/guild#get-guild-widget-settings
type GetGuildWidgetSettings struct {
	GuildID Snowflake
}

// Modify Guild Widget
// PATCH /guilds/{guild.id}/widget
// https://discord.com/developers/docs/resources/guild#modify-guild-widget
type ModifyGuildWidget struct {
	GuildID Snowflake
}

// Get Guild Widget
// GET /guilds/{guild.id}/widget.json
// https://discord.com/developers/docs/resources/guild#get-guild-widget
type GetGuildWidget struct {
	GuildID Snowflake
}

// Get Guild Vanity URL
// GET /guilds/{guild.id}/vanity-url
// https://discord.com/developers/docs/resources/guild#get-guild-vanity-url
type GetGuildVanityURL struct {
	GuildID Snowflake
}

// Get Guild Widget Image
// GET /guilds/{guild.id}/widget.png
// https://discord.com/developers/docs/resources/guild#get-guild-widget-image
type GetGuildWidgetImage struct {
	GuildID Snowflake
	Style   string `json:"style,omitempty"`
}

// Get Guild Welcome Screen
// GET /guilds/{guild.id}/welcome-screen
// https://discord.com/developers/docs/resources/guild#get-guild-welcome-screen
type GetGuildWelcomeScreen struct {
	GuildID Snowflake
}

// Modify Guild Welcome Screen
// PATCH /guilds/{guild.id}/welcome-screen
// https://discord.com/developers/docs/resources/guild#modify-guild-welcome-screen
type ModifyGuildWelcomeScreen struct {
	GuildID         Snowflake
	Enabled         bool                    `json:"enabled,omitempty"`
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels,omitempty"`
	Description     *string                 `json:"description,omitempty"`
}

// Modify Current User Voice State
// PATCH /guilds/{guild.id}/voice-states/@me
// https://discord.com/developers/docs/resources/guild#modify-current-user-voice-state
type ModifyCurrentUserVoiceState struct {
	GuildID                 Snowflake
	Suppress                bool       `json:"suppress,omitempty"`
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
}

// Modify User Voice State
// PATCH /guilds/{guild.id}/voice-states/{user.id}
// https://discord.com/developers/docs/resources/guild#modify-user-voice-state
type ModifyUserVoiceState struct {
	GuildID  Snowflake
	UserID   Snowflake
	Suppress bool `json:"suppress,omitempty"`
}

// List Scheduled Events for Guild
// GET /guilds/{guild.id}/scheduled-events
// https://discord.com/developers/docs/resources/guild-scheduled-event#list-scheduled-events-for-guild
type ListScheduledEventsforGuild struct {
	GuildID       Snowflake
	WithUserCount bool `json:"with_user_count,omitempty"`
}

// Create Guild Scheduled Event
// POST /guilds/{guild.id}/scheduled-events
// https://discord.com/developers/docs/resources/guild-scheduled-event#create-guild-scheduled-event
type CreateGuildScheduledEvent struct {
	GuildID            Snowflake
	ChannelID          *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name               *string                            `json:"name,omitempty"`
	PrivacyLevel       Flag                               `json:"privacy_level,omitempty"`
	ScheduledStartTime Snowflake                          `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime   Snowflake                          `json:"scheduled_end_time,omitempty"`
	Description        *string                            `json:"description,omitempty"`
	EntityType         *Flag                              `json:"entity_type,omitempty"`
	Image              *string                            `json:"image,omitempty"`
}

// Get Guild Scheduled Event
// GET /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}
// https://discord.com/developers/docs/resources/guild-scheduled-event#get-guild-scheduled-event
type GetGuildScheduledEvent struct {
	GuildID               Snowflake
	GuildScheduledEventID Snowflake
	WithUserCount         bool `json:"with_user_count,omitempty"`
}

// Modify Guild Scheduled Event
// PATCH /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}
// https://discord.com/developers/docs/resources/guild-scheduled-event#modify-guild-scheduled-event
type ModifyGuildScheduledEvent struct {
	GuildID               Snowflake
	GuildScheduledEventID Snowflake
	ChannelID             *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata        *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name                  *string                            `json:"name,omitempty"`
	PrivacyLevel          Flag                               `json:"privacy_level,omitempty"`
	ScheduledStartTime    Snowflake                          `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime      Snowflake                          `json:"scheduled_end_time,omitempty"`
	Description           *string                            `json:"description,omitempty"`
	EntityType            *Flag                              `json:"entity_type,omitempty"`
	Image                 *string                            `json:"image,omitempty"`
	Status                Flag                               `json:"status,omitempty"`
}

// Delete Guild Scheduled Event
// DELETE /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}
// https://discord.com/developers/docs/resources/guild-scheduled-event#delete-guild-scheduled-event
type DeleteGuildScheduledEvent struct {
	GuildID               Snowflake
	GuildScheduledEventID Snowflake
}

// Get Guild Scheduled Event Users
// GET /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}/users
// https://discord.com/developers/docs/resources/guild-scheduled-event#get-guild-scheduled-event-users
type GetGuildScheduledEventUsers struct {
	GuildID               Snowflake
	GuildScheduledEventID Snowflake
	Limit                 Flag       `json:"limit,omitempty"`
	WithMember            bool       `json:"with_member,omitempty"`
	Before                *Snowflake `json:"before,omitempty"`
	After                 *Snowflake `json:"after,omitempty"`
}

// Get Guild Template
// GET /guilds/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#get-guild-template
type GetGuildTemplate struct {
	TemplateCode string
}

// Create Guild from Guild Template
// POST /guilds/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#create-guild-from-guild-template
type CreateGuildfromGuildTemplate struct {
	TemplateCode string
	Name         string `json:"name,omitempty"`
	Icon         string `json:"icon,omitempty"`
}

// Get Guild Templates
// GET /guilds/{guild.id}/templates
// https://discord.com/developers/docs/resources/guild-template#get-guild-templates
type GetGuildTemplates struct {
	GuildID Snowflake
}

// Create Guild Template
// POST /guilds/{guild.id}/templates
// https://discord.com/developers/docs/resources/guild-template#create-guild-template
type CreateGuildTemplate struct {
	GuildID     Snowflake
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Sync Guild Template
// PUT /guilds/{guild.id}/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#sync-guild-template
type SyncGuildTemplate struct {
	GuildID      Snowflake
	TemplateCode string
}

// Modify Guild Template
// PATCH /guilds/{guild.id}/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#modify-guild-template
type ModifyGuildTemplate struct {
	GuildID      Snowflake
	TemplateCode string
	Name         string  `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
}

// Delete Guild Template
// DELETE /guilds/{guild.id}/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#delete-guild-template
type DeleteGuildTemplate struct {
	GuildID      Snowflake
	TemplateCode string
}

// Create Interaction Response
// POST /interactions/{interaction.id}/{interaction.token}/callback
// https://discord.com/developers/docs/interactions/receiving-and-responding#create-interaction-response
type CreateInteractionResponse struct {
	InteractionID    Snowflake
	InteractionToken string
}

// Get Original Interaction Response
// GET /webhooks/{application.id}/{interaction.token}/messages/@original
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-original-interaction-response
type GetOriginalInteractionResponse struct {
	ApplicationID    Snowflake
	InteractionToken string
}

// Edit Original Interaction Response
// PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-original-interaction-response
type EditOriginalInteractionResponse struct {
	ApplicationID    Snowflake
	InteractionToken string
}

// Delete Original Interaction Response
// DELETE /webhooks/{application.id}/{interaction.token}/messages/@original
// https://discord.com/developers/docs/interactions/receiving-and-responding#delete-original-interaction-response
type DeleteOriginalInteractionResponse struct {
	ApplicationID    Snowflake
	InteractionToken string
}

// Create Followup Message
// POST /webhooks/{application.id}/{interaction.token}
// https://discord.com/developers/docs/interactions/receiving-and-responding#create-followup-message
type CreateFollowupMessage struct {
	ApplicationID    Snowflake
	InteractionToken string
}

// Get Followup Message
// GET /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-followup-message
type GetFollowupMessage struct {
	ApplicationID    Snowflake
	InteractionToken string
	MessageID        Snowflake
}

// Edit Followup Message
// PATCH /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-followup-message
type EditFollowupMessage struct {
	ApplicationID    Snowflake
	InteractionToken string
	MessageID        Snowflake
}

// Delete Followup Message
// DELETE /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#delete-followup-message
type DeleteFollowupMessage struct {
	ApplicationID    Snowflake
	InteractionToken string
	MessageID        Snowflake
}

// Get Invite
// GET /invites/{invite.code}
// https://discord.com/developers/docs/resources/invite#get-invite
type GetInvite struct {
	InviteCode            string
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id,omitempty"`
	WithCounts            bool      `json:"with_counts,omitempty"`
	WithExpiration        bool      `json:"with_expiration,omitempty"`
}

// Delete Invite
// DELETE /invites/{invite.code}
// https://discord.com/developers/docs/resources/invite#delete-invite
type DeleteInvite struct {
	InviteCode string
}

// Get Current Bot Application Information
// GET /oauth2/applications/@me
// https://discord.com/developers/docs/topics/oauth2#get-current-bot-application-information
type GetCurrentBotApplicationInformation struct{}

// Get Current Authorization Information
// GET /oauth2/@me
// https://discord.com/developers/docs/topics/oauth2#get-current-authorization-information
type GetCurrentAuthorizationInformation struct{}

// Create Stage Instance
// POST /stage-instances
// https://discord.com/developers/docs/resources/stage-instance#create-stage-instance
type CreateStageInstance struct {
	ChannelID             Snowflake `json:"channel_id,omitempty"`
	Topic                 string    `json:"topic,omitempty"`
	PrivacyLevel          Flag      `json:"privacy_level,omitempty"`
	SendStartNotification bool      `json:"send_start_notification,omitempty"`
}

// Get Stage Instance
// GET /stage-instances/{channel.id}
// https://discord.com/developers/docs/resources/stage-instance#get-stage-instance
type GetStageInstance struct {
	ChannelID Snowflake
}

// Modify Stage Instance
// PATCH /stage-instances/{channel.id}
// https://discord.com/developers/docs/resources/stage-instance#modify-stage-instance
type ModifyStageInstance struct {
	ChannelID    Snowflake
	Topic        string `json:"topic,omitempty"`
	PrivacyLevel Flag   `json:"privacy_level,omitempty"`
}

// Delete Stage Instance
// DELETE /stage-instances/{channel.id}
// https://discord.com/developers/docs/resources/stage-instance#delete-stage-instance
type DeleteStageInstance struct {
	ChannelID Snowflake
}

// Get Sticker
// GET /stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#get-sticker
type GetSticker struct {
	StickerID Snowflake
}

// List Nitro Sticker Packs
// GET /sticker-packs
// https://discord.com/developers/docs/resources/sticker#list-nitro-sticker-packs
type ListNitroStickerPacks struct {
	StickerPacks []*StickerPack `json:"sticker_packs,omitempty"`
}

// List Guild Stickers
// GET /guilds/{guild.id}/stickers
// https://discord.com/developers/docs/resources/sticker#list-guild-stickers
type ListGuildStickers struct {
	GuildID Snowflake
}

// Get Guild Sticker
// GET /guilds/{guild.id}/stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#get-guild-sticker
type GetGuildSticker struct {
	GuildID   Snowflake
	StickerID Snowflake
}

// Create Guild Sticker
// POST /guilds/{guild.id}/stickers
// https://discord.com/developers/docs/resources/sticker#create-guild-sticker
type CreateGuildSticker struct {
	GuildID     Snowflake
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Tags        *string `json:"tags,omitempty"`
	Files       []byte  `dasgo:"files"`
}

// Modify Guild Sticker
// PATCH /guilds/{guild.id}/stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#modify-guild-sticker
type ModifyGuildSticker struct {
	GuildID     Snowflake
	StickerID   Snowflake
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Tags        *string `json:"tags,omitempty"`
}

// Delete Guild Sticker
// DELETE /guilds/{guild.id}/stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#delete-guild-sticker
type DeleteGuildSticker struct {
	GuildID   Snowflake
	StickerID Snowflake
}

// Modify Current User
// PATCH /users/@me
// https://discord.com/developers/docs/resources/user#modify-current-user
type ModifyCurrentUser struct {
	Username string  `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

// Get Current User Guilds
// GET /users/@me/guilds
// https://discord.com/developers/docs/resources/user#get-current-user-guilds
type GetCurrentUserGuilds struct {
	Before *Snowflake `json:"before,omitempty"`
	After  *Snowflake `json:"after,omitempty"`
	Limit  Flag       `json:"limit,omitempty"`
}

// Get Current User Guild Member
// GET /users/@me/guilds/{guild.id}/member
// https://discord.com/developers/docs/resources/user#get-current-user-guild-member
type GetCurrentUserGuildMember struct {
	GuildID Snowflake
}

// Leave Guild
// DELETE /users/@me/guilds/{guild.id}
// https://discord.com/developers/docs/resources/user#leave-guild
type LeaveGuild struct {
	GuildID Snowflake
}

// Create DM
// POST /users/@me/channels
// https://discord.com/developers/docs/resources/user#create-dm
type CreateDM struct {
	RecipientID Snowflake `json:"recipient_id,omitempty"`
}

// Create Group DM
// POST /users/@me/channels
// https://discord.com/developers/docs/resources/user#create-group-dm
type CreateGroupDM struct {
	AccessTokens []*string            `json:"access_tokens,omitempty"`
	Nicks        map[Snowflake]string `json:"nicks,omitempty"`
}

// Get User Connections
// GET /users/@me/connections
// https://discord.com/developers/docs/resources/user#get-user-connections
type GetUserConnections struct {
	RecipientID Snowflake `json:"recipient_id,omitempty"`
}

// List Voice Regions
// GET /voice/regions
// https://discord.com/developers/docs/resources/voice#list-voice-regions
type ListVoiceRegions struct{}

// Create Webhook
// POST /channels/{channel.id}/webhooks
// https://discord.com/developers/docs/resources/webhook#create-webhook
type CreateWebhook struct {
	ChannelID Snowflake
	Name      string `json:"name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

// Get Channel Webhooks
// GET /channels/{channel.id}/webhooks
// https://discord.com/developers/docs/resources/webhook#get-channel-webhooks
type GetChannelWebhooks struct {
	ChannelID Snowflake
}

// Get Guild Webhooks
// GET /guilds/{guild.id}/webhooks
// https://discord.com/developers/docs/resources/webhook#get-guild-webhooks
type GetGuildWebhooks struct {
	GuildID Snowflake
}

// Get Webhook
// GET /webhooks/{webhook.id}
// https://discord.com/developers/docs/resources/webhook#get-webhook
type GetWebhook struct {
	WebhookID Snowflake
}

// Get Webhook with Token
// GET /webhooks/{webhook.id}/{webhook.token}
// https://discord.com/developers/docs/resources/webhook#get-webhook-with-token
type GetWebhookwithToken struct {
	WebhookID    Snowflake
	WebhookToken string
}

// Modify Webhook
// PATCH /webhooks/{webhook.id}
// https://discord.com/developers/docs/resources/webhook#modify-webhook
type ModifyWebhook struct {
	WebhookID Snowflake
	Name      string    `json:"name,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
}

// Modify Webhook with Token
// PATCH /webhooks/{webhook.id}/{webhook.token}
// https://discord.com/developers/docs/resources/webhook#modify-webhook-with-token
type ModifyWebhookwithToken struct {
	WebhookID    Snowflake
	WebhookToken string
}

// Delete Webhook
// DELETE /webhooks/{webhook.id}
// https://discord.com/developers/docs/resources/webhook#delete-webhook
type DeleteWebhook struct {
	WebhookID Snowflake
}

// Delete Webhook with Token
// DELETE /webhooks/{webhook.id}/{webhook.token}
// https://discord.com/developers/docs/resources/webhook#delete-webhook-with-token
type DeleteWebhookwithToken struct {
	WebhookID    Snowflake
	WebhookToken string
}

// Execute Webhook
// POST /webhooks/{webhook.id}/{webhook.token}
// https://discord.com/developers/docs/resources/webhook#execute-webhook
type ExecuteWebhook struct {
	WebhookID       Snowflake
	WebhookToken    string
	Wait            bool             `json:"wait,omitempty"`
	ThreadID        Snowflake        `json:"thread_id,omitempty"`
	Content         string           `json:"content,omitempty"`
	Username        string           `json:"username,omitempty"`
	AvatarURL       string           `json:"avatar_url,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Files           []byte           `dasgo:"files"`
	Components      []Component      `json:"components,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	PayloadJSON     string           `json:"payload_json,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// Execute Slack-Compatible Webhook
// POST /webhooks/{webhook.id}/{webhook.token}/slack
// https://discord.com/developers/docs/resources/webhook#execute-slackcompatible-webhook
type ExecuteSlackCompatibleWebhook struct {
	WebhookID    Snowflake
	WebhookToken string
	ThreadID     Snowflake `json:"thread_id,omitempty"`
	Wait         bool      `json:"wait,omitempty"`
}

// Execute GitHub-Compatible Webhook
// POST /webhooks/{webhook.id}/{webhook.token}/github
// https://discord.com/developers/docs/resources/webhook#execute-githubcompatible-webhook
type ExecuteGitHubCompatibleWebhook struct {
	WebhookID    Snowflake
	WebhookToken string
	ThreadID     Snowflake `json:"thread_id,omitempty"`
	Wait         bool      `json:"wait,omitempty"`
}

// Get Webhook Message
// GET /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#get-webhook-message
type GetWebhookMessage struct {
	WebhookID    Snowflake
	WebhookToken string
	MessageID    Snowflake
	ThreadID     Snowflake `json:"thread_id,omitempty"`
}

// Edit Webhook Message
// PATCH /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message
type EditWebhookMessage struct {
	WebhookID       Snowflake
	WebhookToken    string
	MessageID       Snowflake
	ThreadID        Snowflake        `json:"thread_id,omitempty"`
	Content         *string          `json:"content,omitempty"`
	Components      []*Component     `json:"components,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	Files           []byte           `dasgo:"files"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	PayloadJSON     string           `json:"payload_json,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
}

// Delete Webhook Message
// DELETE /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#delete-webhook-message
type DeleteWebhookMessage struct {
	WebhookID    Snowflake
	WebhookToken string
	MessageID    Snowflake
}
