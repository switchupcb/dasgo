package dasgo

// Get Global Application Commands
// GET /applications/{application.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#get-global-application-commands
type GetGlobalApplicationCommands struct {
	ApplicationID     Snowflake `url:"-"`
	WithLocalizations *bool     `url:"with_localizations,omitempty"`
}

// Create Global Application Command
// POST /applications/{application.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#create-global-application-command
type CreateGlobalApplicationCommand struct {
	ApplicationID            Snowflake                   `json:"-"`
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`
	Description              *string                     `json:"description,omitempty"`
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions **string                    `json:"default_member_permissions,omitempty"`
	DMPermission             **bool                      `json:"dm_permission,omitempty"`
	Type                     *Flag                       `json:"type,omitempty"`
	NSFW                     *bool                       `json:"nsfw,omitempty"`
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
	ApplicationID            Snowflake                   `json:"-"`
	CommandID                Snowflake                   `json:"-"`
	Name                     *string                     `json:"name,omitempty"`
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`
	Description              *string                     `json:"description,omitempty"`
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions **string                    `json:"default_member_permissions,omitempty"`
	DMPermission             **bool                      `json:"dm_permission,omitempty"`
	NSFW                     *bool                       `json:"nsfw,omitempty"`
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
	ApplicationID       Snowflake             `json:"-"`
	ApplicationCommands []*ApplicationCommand `json:"commands"`
}

// Get Guild Application Commands
// GET /applications/{application.id}/guilds/{guild.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#get-guild-application-commands
type GetGuildApplicationCommands struct {
	ApplicationID     Snowflake `url:"-"`
	GuildID           Snowflake `url:"-"`
	WithLocalizations *bool     `url:"with_localizations,omitempty"`
}

// Create Guild Application Command
// POST /applications/{application.id}/guilds/{guild.id}/commands
// https://discord.com/developers/docs/interactions/application-commands#create-guild-application-command
type CreateGuildApplicationCommand struct {
	ApplicationID            Snowflake                   `json:"-"`
	GuildID                  Snowflake                   `json:"-"`
	Name                     string                      `json:"name"`
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`
	Description              *string                     `json:"description,omitempty"`
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions **string                    `json:"default_member_permissions,omitempty"`
	Type                     *Flag                       `json:"type,omitempty"`
	NSFW                     *bool                       `json:"nsfw,omitempty"`
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
	ApplicationID            Snowflake                   `json:"-"`
	GuildID                  Snowflake                   `json:"-"`
	CommandID                Snowflake                   `json:"-"`
	Name                     *string                     `json:"name,omitempty"`
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`
	Description              *string                     `json:"description,omitempty"`
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions **string                    `json:"default_member_permissions,omitempty"`
	NSFW                     *bool                       `json:"nsfw,omitempty"`
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
	ApplicationID       Snowflake             `json:"-"`
	GuildID             Snowflake             `json:"-"`
	ApplicationCommands []*ApplicationCommand `json:"commands"`
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
	ApplicationID Snowflake                        `json:"-"`
	GuildID       Snowflake                        `json:"-"`
	CommandID     Snowflake                        `json:"-"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions"`
}

// Batch Edit Application Command Permissions
// PUT /applications/{application.id}/guilds/{guild.id}/commands/permissions
// https://discord.com/developers/docs/interactions/application-commands#batch-edit-application-command-permissions
type BatchEditApplicationCommandPermissions struct {
	ApplicationID Snowflake
	GuildID       Snowflake
}

// Create Interaction Response
// POST /interactions/{interaction.id}/{interaction.token}/callback
// https://discord.com/developers/docs/interactions/receiving-and-responding#create-interaction-response
type CreateInteractionResponse struct {
	InteractionID    Snowflake
	InteractionToken string
	*InteractionResponse
}

// Get Original Interaction Response
// GET /webhooks/{application.id}/{interaction.token}/messages/@original
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-original-interaction-response
type GetOriginalInteractionResponse struct {
	ApplicationID    Snowflake  `url:"-"`
	InteractionToken string     `url:"-"`
	ThreadID         *Snowflake `url:"thread_id,omitempty"`
}

// Edit Original Interaction Response
// PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-original-interaction-response
type EditOriginalInteractionResponse struct {
	ApplicationID    Snowflake         `json:"-" url:"-"`
	InteractionToken string            `json:"-" url:"-"`
	ThreadID         *Snowflake        `json:"-" url:"thread_id,omitempty"`
	Content          **string          `json:"content,omitempty" url:"-"`
	Embeds           *[]*Embed         `json:"embeds,omitempty" url:"-"`
	Components       *[]Component      `json:"components,omitempty" url:"-"`
	Files            []*File           `json:"-" url:"-" dasgo:"files"`
	AllowedMentions  **AllowedMentions `json:"allowed_mentions,omitempty" url:"-"`
	PayloadJSON      string            `json:"-" url:"-" multipart:"payload_json"`
	Attachments      *[]*Attachment    `json:"attachments,omitempty" url:"-"`
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
	ApplicationID    Snowflake        `json:"-" url:"-"`
	InteractionToken string           `json:"-" url:"-"`
	ThreadID         *Snowflake       `json:"-" url:"thread_id,omitempty"`
	Content          *string          `json:"content,omitempty" url:"-"`
	Username         *string          `json:"username,omitempty" url:"-"`
	AvatarURL        *string          `json:"avatar_url,omitempty" url:"-"`
	TTS              *bool            `json:"tts,omitempty" url:"-"`
	Embeds           []*Embed         `json:"embeds,omitempty" url:"-"`
	AllowedMentions  *AllowedMentions `json:"allowed_mentions,omitempty" url:"-"`
	Components       []Component      `json:"components,omitempty" url:"-"`
	Files            []*File          `json:"-" url:"-" dasgo:"files"`
	PayloadJSON      string           `json:"-" url:"-" multipart:"payload_json"`
	Attachments      []*Attachment    `json:"attachments,omitempty" url:"-"`
	Flags            *BitFlag         `json:"flags,omitempty" url:"-"`
	ThreadName       *string          `json:"thread_name,omitempty" url:"-"`
}

// Get Followup Message
// GET /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#get-followup-message
type GetFollowupMessage struct {
	ApplicationID    Snowflake  `url:"-"`
	InteractionToken string     `url:"-"`
	MessageID        Snowflake  `url:"-"`
	ThreadID         *Snowflake `url:"thread_id,omitempty"`
}

// Edit Followup Message
// PATCH /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#edit-followup-message
type EditFollowupMessage struct {
	ApplicationID    Snowflake         `json:"-" url:"-"`
	InteractionToken string            `json:"-" url:"-"`
	MessageID        Snowflake         `json:"-" url:"-"`
	ThreadID         *Snowflake        `json:"-" url:"thread_id,omitempty"`
	Content          **string          `json:"content,omitempty" url:"-"`
	Embeds           *[]*Embed         `json:"embeds,omitempty" url:"-"`
	Components       *[]Component      `json:"components,omitempty" url:"-"`
	Files            []*File           `json:"-" url:"-" dasgo:"files"`
	AllowedMentions  **AllowedMentions `json:"allowed_mentions,omitempty" url:"-"`
	PayloadJSON      string            `json:"-" url:"-" multipart:"payload_json"`
	Attachments      *[]*Attachment    `json:"attachments,omitempty" url:"-"`
}

// Delete Followup Message
// DELETE /webhooks/{application.id}/{interaction.token}/messages/{message.id}
// https://discord.com/developers/docs/interactions/receiving-and-responding#delete-followup-message
type DeleteFollowupMessage struct {
	ApplicationID    Snowflake
	InteractionToken string
	MessageID        Snowflake
}

// Get Current Application
// GET /applications/@me
// https://discord.com/developers/docs/resources/application#get-current-application
type GetCurrentApplication struct{}

// Get Application Role Connection Metadata Records
// GET /applications/{application.id}/role-connections/metadata
// https://discord.com/developers/docs/resources/application-role-connection-metadata#get-application-role-connection-metadata-records
type GetApplicationRoleConnectionMetadataRecords struct {
	ApplicationID Snowflake
}

// Update Application Role Connection Metadata Records
// PUT /applications/{application.id}/role-connections/metadata
// https://discord.com/developers/docs/resources/application-role-connection-metadata#update-application-role-connection-metadata-records
type UpdateApplicationRoleConnectionMetadataRecords struct {
	ApplicationID Snowflake
}

// Get Guild Audit Log
// GET /guilds/{guild.id}/audit-logs
// https://discord.com/developers/docs/resources/audit-log#get-guild-audit-log
type GetGuildAuditLog struct {
	GuildID Snowflake `url:"-"`

	// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
	UserID     Snowflake `url:"user_id,omitempty"`
	ActionType Flag      `url:"action_type,omitempty"`
	Before     Snowflake `url:"before,omitempty"`
	After      Snowflake `url:"after,omitempty"`
	Limit      int       `url:"limit,omitempty"`
}

// List Auto Moderation Rules for Guild
// GET /guilds/{guild.id}/auto-moderation/rules
// https://discord.com/developers/docs/resources/auto-moderation#list-auto-moderation-rules-for-guild
type ListAutoModerationRulesForGuild struct {
	GuildID Snowflake
}

// Get Auto Moderation Rule
// GET /guilds/{guild.id}/auto-moderation/rules/{auto_moderation_rule.id}
// https://discord.com/developers/docs/resources/auto-moderation#get-auto-moderation-rule
type GetAutoModerationRule struct {
	GuildID              Snowflake
	AutoModerationRuleID Snowflake
}

// Create Auto Moderation Rule
// POST /guilds/{guild.id}/auto-moderation/rules
// https://discord.com/developers/docs/resources/auto-moderation#create-auto-moderation-rule
type CreateAutoModerationRule struct {
	GuildID         Snowflake               `json:"-"`
	Name            string                  `json:"name"`
	EventType       Flag                    `json:"event_type"`
	TriggerType     Flag                    `json:"trigger_type"`
	TriggerMetadata *TriggerMetadata        `json:"trigger_metadata,omitempty"`
	Actions         []*AutoModerationAction `json:"actions"`
	Enabled         *bool                   `json:"enabled,omitempty"`
	ExemptRoles     []Snowflake             `json:"exempt_roles,omitempty"`
	ExemptChannels  []Snowflake             `json:"exempt_channels,omitempty"`
}

// Modify Auto Moderation Rule
// PATCH /guilds/{guild.id}/auto-moderation/rules/{auto_moderation_rule.id}
// https://discord.com/developers/docs/resources/auto-moderation#modify-auto-moderation-rule
type ModifyAutoModerationRule struct {
	GuildID              Snowflake               `json:"-"`
	AutoModerationRuleID Snowflake               `json:"-"`
	Name                 *string                 `json:"name,omitempty"`
	EventType            *Flag                   `json:"event_type,omitempty"`
	TriggerType          *Flag                   `json:"trigger_type,omitempty"`
	TriggerMetadata      *TriggerMetadata        `json:"trigger_metadata,omitempty"`
	Actions              []*AutoModerationAction `json:"actions,omitempty"`
	Enabled              *bool                   `json:"enabled,omitempty"`
	ExemptRoles          []Snowflake             `json:"exempt_roles,omitempty"`
	ExemptChannels       []Snowflake             `json:"exempt_channels,omitempty"`
}

// Delete Auto Moderation Rule
// DELETE /guilds/{guild.id}/auto-moderation/rules/{auto_moderation_rule.id}
// https://discord.com/developers/docs/resources/auto-moderation#delete-auto-moderation-rule
type DeleteAutoModerationRule struct {
	GuildID              Snowflake
	AutoModerationRuleID Snowflake
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

// Modify Channel Group DM
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-group-dm
type ModifyChannelGroupDM struct {
	ChannelID Snowflake `json:"-"`
	Name      *string   `json:"name,omitempty"`
	Icon      *string   `json:"icon,omitempty"`
}

// Modify Channel Guild
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-guild-channel
type ModifyChannelGuild struct {
	ChannelID                     Snowflake               `json:"-"`
	Name                          *string                 `json:"name,omitempty"`
	Type                          *Flag                   `json:"type,omitempty"`
	Position                      **int                   `json:"position,omitempty"`
	Topic                         **string                `json:"topic,omitempty"`
	NSFW                          **bool                  `json:"nsfw,omitempty"`
	RateLimitPerUser              **int                   `json:"rate_limit_per_user,omitempty"`
	Bitrate                       **int                   `json:"bitrate,omitempty"`
	UserLimit                     **int                   `json:"user_limit,omitempty"`
	PermissionOverwrites          *[]*PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                      **Snowflake             `json:"parent_id,omitempty"`
	RTCRegion                     **string                `json:"rtc_region,omitempty"`
	VideoQualityMode              **Flag                  `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    **int                   `json:"default_auto_archive_duration,omitempty"`
	Flags                         *BitFlag                `json:"flags,omitempty"`
	AvailableTags                 []*ForumTag             `json:"available_tags,omitempty"`
	DefaultReactionEmoji          **DefaultReaction       `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int                    `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              **Flag                  `json:"default_sort_order,omitempty"`
}

// Modify Channel
// PATCH /channels/{channel.id}
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-thread
type ModifyChannelThread struct {
	ChannelID           Snowflake   `json:"-"`
	Name                *string     `json:"name,omitempty"`
	Archived            *bool       `json:"archived,omitempty"`
	AutoArchiveDuration *int        `json:"auto_archive_duration,omitempty"`
	Locked              *bool       `json:"locked,omitempty"`
	Invitable           *bool       `json:"invitable,omitempty"`
	RateLimitPerUser    **int       `json:"rate_limit_per_user,omitempty"`
	Flags               *BitFlag    `json:"flags,omitempty"`
	AppliedTags         []Snowflake `json:"applied_tags,omitempty"`
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
	ChannelID Snowflake  `url:"-"`
	Around    *Snowflake `url:"around,omitempty"`
	Before    *Snowflake `url:"before,omitempty"`
	After     *Snowflake `url:"after,omitempty"`
	Limit     *Flag      `url:"limit,omitempty"`
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
	ChannelID        Snowflake         `json:"-"`
	Content          *string           `json:"content,omitempty"`
	Nonce            *Nonce            `json:"nonce,omitempty"`
	TTS              *bool             `json:"tts,omitempty"`
	Embeds           []*Embed          `json:"embeds,omitempty"`
	AllowedMentions  *AllowedMentions  `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReference `json:"message_reference,omitempty"`
	Components       []Component       `json:"components,omitempty"`
	StickerIDS       []*Snowflake      `json:"sticker_ids,omitempty"`
	Files            []*File           `json:"-" dasgo:"files,omitempty"`
	PayloadJSON      *string           `json:"-" multipart:"payload_json,omitempty"`
	Attachments      []*Attachment     `json:"attachments,omitempty"`
	Flags            *BitFlag          `json:"flags,omitempty"`
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
	ChannelID Snowflake  `url:"-"`
	MessageID Snowflake  `url:"-"`
	Emoji     string     `url:"-"`
	After     *Snowflake `url:"after,omitempty"`
	Limit     *int       `url:"limit,omitempty"`
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
	ChannelID       Snowflake         `json:"-"`
	MessageID       Snowflake         `json:"-"`
	Content         **string          `json:"content,omitempty"`
	Embeds          *[]*Embed         `json:"embeds,omitempty"`
	Flags           **BitFlag         `json:"flags,omitempty"`
	AllowedMentions **AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      *[]Component      `json:"components,omitempty"`
	Files           []*File           `json:"-" dasgo:"files"`
	PayloadJSON     *string           `json:"-" multipart:"payload_json"`
	Attachments     *[]*Attachment    `json:"attachments,omitempty"`
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
	ChannelID Snowflake    `json:"-"`
	Messages  []*Snowflake `json:"messages"`
}

// Edit Channel Permissions
// PUT /channels/{channel.id}/permissions/{overwrite.id}
// https://discord.com/developers/docs/resources/channel#edit-channel-permissions
type EditChannelPermissions struct {
	ChannelID   Snowflake `json:"-"`
	OverwriteID Snowflake `json:"-"`
	Allow       **string  `json:"allow,omitempty"`
	Deny        **string  `json:"deny,omitempty"`
	Type        Flag      `json:"type"`
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
	ChannelID           Snowflake `json:"-"`
	MaxAge              *int      `json:"max_age"`
	MaxUses             *int      `json:"max_uses"`
	Temporary           bool      `json:"temporary"`
	Unique              bool      `json:"unique"`
	TargetType          Flag      `json:"target_type"`
	TargetUserID        Snowflake `json:"target_user_id"`
	TargetApplicationID Snowflake `json:"target_application_id"`
}

// Delete Channel Permission
// DELETE /channels/{channel.id}/permissions/{overwrite.id}
// https://discord.com/developers/docs/resources/channel#delete-channel-permission
type DeleteChannelPermission struct {
	ChannelID   Snowflake
	OverwriteID Snowflake
}

// Follow Announcement Channel
// POST /channels/{channel.id}/followers
// https://discord.com/developers/docs/resources/channel#follow-announcement-channel
type FollowAnnouncementChannel struct {
	ChannelID        Snowflake `json:"-"`
	WebhookChannelID Snowflake `json:"webhook_channel_id"`
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
	ChannelID   Snowflake `json:"-"`
	UserID      Snowflake `json:"-"`
	AccessToken string    `json:"access_token"`
	Nickname    *string   `json:"nick"`
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
	ChannelID           Snowflake `json:"-"`
	MessageID           Snowflake `json:"-"`
	Name                string    `json:"name"`
	AutoArchiveDuration *int      `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    **int     `json:"rate_limit_per_user,omitempty"`
}

// Start Thread without Message
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-without-message
type StartThreadwithoutMessage struct {
	ChannelID           Snowflake `json:"-"`
	Name                string    `json:"name"`
	AutoArchiveDuration *int      `json:"auto_archive_duration,omitempty"`
	Type                *Flag     `json:"type,omitempty"`
	Invitable           *bool     `json:"invitable,omitempty"`
	RateLimitPerUser    **int     `json:"rate_limit_per_user,omitempty"`
}

// Start Thread in Forum Channel
// POST /channels/{channel.id}/threads
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel
type StartThreadinForumChannel struct {
	ChannelID           Snowflake                 `json:"-"`
	Name                string                    `json:"name"`
	AutoArchiveDuration *int                      `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    **int                     `json:"rate_limit_per_user,omitempty"`
	Message             *ForumThreadMessageParams `json:"message"`
	AppliedTags         []Snowflake               `json:"applied_tags,omitempty"`
}

// Forum Thread Message Params Object
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-forum-thread-message-params-object
type ForumThreadMessageParams struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	StickerIDS      []*Snowflake     `json:"sticker_ids,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
	Files           []*File          `json:"-" dasgo:"files"`
	PayloadJSON     string           `json:"-" multipart:"payload_json,omitempty"`
	Flags           *BitFlag         `json:"flags,omitempty"`
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
	ChannelID  Snowflake `url:"-"`
	UserID     Snowflake `url:"-"`
	WithMember *bool     `url:"with_member,omitempty"`
}

// List Thread Members
// GET /channels/{channel.id}/thread-members
// https://discord.com/developers/docs/resources/channel#list-thread-members
type ListThreadMembers struct {
	ChannelID  Snowflake  `url:"-"`
	WithMember *bool      `url:"with_member,omitempty"`
	After      *Snowflake `url:"after,omitempty"`
	Limit      *int       `url:"limit,omitempty"`
}

// List Public Archived Threads
// GET /channels/{channel.id}/threads/archived/public
// https://discord.com/developers/docs/resources/channel#list-public-archived-threads
type ListPublicArchivedThreads struct {
	ChannelID Snowflake  `url:"-"`
	Before    *Timestamp `url:"before,omitempty"`
	Limit     *int       `url:"limit,omitempty"`
}

// List Private Archived Threads
// GET /channels/{channel.id}/threads/archived/private
// https://discord.com/developers/docs/resources/channel#list-private-archived-threads
type ListPrivateArchivedThreads struct {
	ChannelID Snowflake  `url:"-"`
	Before    *Timestamp `url:"before,omitempty"`
	Limit     *int       `url:"limit,omitempty"`
}

// List Joined Private Archived Threads
// GET /channels/{channel.id}/users/@me/threads/archived/private
// https://discord.com/developers/docs/resources/channel#list-joined-private-archived-threads
type ListJoinedPrivateArchivedThreads struct {
	ChannelID Snowflake  `url:"-"`
	Before    *Timestamp `url:"before,omitempty"`
	Limit     *int       `url:"limit,omitempty"`
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
	GuildID Snowflake    `json:"-"`
	Name    string       `json:"name"`
	Image   string       `json:"image"`
	Roles   []*Snowflake `json:"roles"`
}

// Modify Guild Emoji
// PATCH /guilds/{guild.id}/emojis/{emoji.id}
// https://discord.com/developers/docs/resources/emoji#modify-guild-emoji
type ModifyGuildEmoji struct {
	GuildID Snowflake     `json:"-"`
	EmojiID Snowflake     `json:"-"`
	Name    *string       `json:"name,omitempty"`
	Roles   *[]*Snowflake `json:"roles,omitempty"`
}

// Delete Guild Emoji
// DELETE /guilds/{guild.id}/emojis/{emoji.id}
// https://discord.com/developers/docs/resources/emoji#delete-guild-emoji
type DeleteGuildEmoji struct {
	GuildID Snowflake
	EmojiID Snowflake
}

// Create Guild
// POST /guilds
// https://discord.com/developers/docs/resources/guild#create-guild
type CreateGuild struct {
	Name                        string     `json:"name"`
	Region                      **string   `json:"region,omitempty"`
	Icon                        *string    `json:"icon,omitempty"`
	VerificationLevel           *Flag      `json:"verification_level,omitempty"`
	DefaultMessageNotifications *Flag      `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *Flag      `json:"explicit_content_filter,omitempty"`
	Roles                       []*Role    `json:"roles,omitempty"`
	Channels                    []*Channel `json:"channels,omitempty"`
	AfkChannelID                *Snowflake `json:"afk_channel_id,omitempty"`
	AfkTimeout                  *int       `json:"afk_timeout,omitempty"`
	SystemChannelID             *Snowflake `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *BitFlag   `json:"system_channel_flags,omitempty"`
}

// Get Guild
// GET /guilds/{guild.id}
// https://discord.com/developers/docs/resources/guild#get-guild
type GetGuild struct {
	GuildID    Snowflake `url:"-"`
	WithCounts *bool     `url:"with_counts,omitempty"`
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
	GuildID                     Snowflake   `json:"-"`
	Name                        *string     `json:"name,omitempty"`
	VerificationLevel           **Flag      `json:"verification_level,omitempty"`
	DefaultMessageNotifications **Flag      `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       **Flag      `json:"explicit_content_filter,omitempty"`
	AFKChannelID                **Snowflake `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int         `json:"afk_timeout,omitempty"`
	Icon                        **string    `json:"icon,omitempty"`
	OwnerID                     Snowflake   `json:"owner_id,omitempty"`
	Splash                      **string    `json:"splash,omitempty"`
	DiscoverySplash             **string    `json:"discovery_splash,omitempty"`
	Banner                      **string    `json:"banner,omitempty"`
	SystemChannelID             **Snowflake `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *BitFlag    `json:"system_channel_flags,omitempty"`
	RulesChannelID              **Snowflake `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      **Snowflake `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             **string    `json:"preferred_locale,omitempty"`
	Features                    []*string   `json:"features,omitempty"`
	Description                 **string    `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool       `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       **Snowflake `json:"safety_alerts_channel_id,omitempty"`
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
	GuildID                    Snowflake               `json:"-"`
	Name                       string                  `json:"name"`
	Type                       **Flag                  `json:"type,omitempty"`
	Topic                      **string                `json:"topic,omitempty"`
	Bitrate                    **int                   `json:"bitrate,omitempty"`
	UserLimit                  **int                   `json:"user_limit,omitempty"`
	RateLimitPerUser           **int                   `json:"rate_limit_per_user,omitempty"`
	Position                   **int                   `json:"position,omitempty"`
	PermissionOverwrites       *[]*PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                   **Snowflake             `json:"parent_id,omitempty"`
	NSFW                       **bool                  `json:"nsfw,omitempty"`
	RTCRegion                  **string                `json:"rtc_region,omitempty"`
	VideoQualityMode           **Flag                  `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration **int                   `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji       **DefaultReaction       `json:"default_reaction_emoji,omitempty"`
	AvailableTags              *[]*ForumTag            `json:"available_tags,omitempty"`
	DefaultSortOrder           **Flag                  `json:"default_sort_order,omitempty"`
	DefaultForumLayout         **Flag                  `json:"default_forum_layout,omitempty"`
}

// Modify Guild Channel Positions
// PATCH /guilds/{guild.id}/channels
// https://discord.com/developers/docs/resources/guild#modify-guild-channel-positions
type ModifyGuildChannelPositions struct {
	GuildID    Snowflake                               `json:"-"`
	Parameters []*ModifyGuildChannelPositionParameters `json:"parameters"`
}

// Modify Guild Channel Position Parameters
// https://discord.com/developers/docs/resources/guild#modify-guild-channel-positions-json-params
type ModifyGuildChannelPositionParameters struct {
	ID              Snowflake   `json:"id"`
	Position        **int       `json:"position,omitempty"`
	LockPermissions **bool      `json:"lock_permissions,omitempty"`
	ParentID        **Snowflake `json:"parent_id,omitempty"`
}

// List Active Guild Threads
// GET /guilds/{guild.id}/threads/active
// https://discord.com/developers/docs/resources/guild#list-active-guild-threads
type ListActiveGuildThreads struct {
	GuildID Snowflake `json:"-"`
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
	GuildID Snowflake  `url:"-"`
	Limit   *int       `url:"limit,omitempty"`
	After   *Snowflake `url:"after,omitempty"`
}

// Search Guild Members
// GET /guilds/{guild.id}/members/search
// https://discord.com/developers/docs/resources/guild#search-guild-members
type SearchGuildMembers struct {
	GuildID Snowflake `url:"-"`
	Query   string    `url:"query"`
	Limit   *int      `url:"limit,omitempty"`
}

// Add Guild Member
// PUT /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#add-guild-member
type AddGuildMember struct {
	GuildID     Snowflake   `json:"-"`
	UserID      Snowflake   `json:"-"`
	AccessToken string      `json:"access_token"`
	Nick        *string     `json:"nick,omitempty"`
	Roles       []Snowflake `json:"roles,omitempty"`
	Mute        *bool       `json:"mute,omitempty"`
	Deaf        *bool       `json:"deaf,omitempty"`
}

// Modify Guild Member
// PATCH /guilds/{guild.id}/members/{user.id}
// https://discord.com/developers/docs/resources/guild#modify-guild-member
type ModifyGuildMember struct {
	GuildID                    Snowflake    `json:"-"`
	UserID                     Snowflake    `json:"-"`
	Nick                       **string     `json:"nick,omitempty"`
	Roles                      *[]Snowflake `json:"roles,omitempty"`
	Mute                       **bool       `json:"mute,omitempty"`
	Deaf                       **bool       `json:"deaf,omitempty"`
	ChannelID                  **Snowflake  `json:"channel_id,omitempty"`
	CommunicationDisabledUntil **Timestamp  `json:"communication_disabled_until,omitempty"`
	Flags                      **BitFlag    `json:"flags,omitempty"`
}

// Modify Current Member
// PATCH /guilds/{guild.id}/members/@me
// https://discord.com/developers/docs/resources/guild#modify-current-member
type ModifyCurrentMember struct {
	GuildID Snowflake `json:"-"`
	Nick    **string  `json:"nick,omitempty"`
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
	GuildID Snowflake  `url:"-"`
	Limit   *int       `url:"limit,omitempty"`
	Before  *Snowflake `url:"before,omitempty"`
	After   *Snowflake `url:"after,omitempty"`
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
	GuildID              Snowflake `json:"-"`
	UserID               Snowflake `json:"-"`
	DeleteMessageSeconds *int      `json:"delete_message_seconds,omitempty"`
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
	GuildID      Snowflake `json:"-"`
	Name         *string   `json:"name,omitempty"`
	Permissions  *string   `json:"permissions,omitempty"`
	Color        *int      `json:"color,omitempty"`
	Hoist        *bool     `json:"hoist,omitempty"`
	Icon         **string  `json:"icon,omitempty"`
	UnicodeEmoji **string  `json:"unicode_emoji,omitempty"`
	Mentionable  *bool     `json:"mentionable,omitempty"`
}

// Modify Guild Role Positions
// PATCH /guilds/{guild.id}/roles
// https://discord.com/developers/docs/resources/guild#modify-guild-role-positions
type ModifyGuildRolePositions struct {
	GuildID    Snowflake                            `json:"-"`
	Parameters []*ModifyGuildRolePositionParameters `json:"parameters"`
}

// Modify Guild Role Position Parameters
// https://discord.com/developers/docs/resources/guild#create-guild-role-json-params
type ModifyGuildRolePositionParameters struct {
	ID       Snowflake `json:"id"`
	Position **int     `json:"position,omitempty"`
}

// Modify Guild Role
// PATCH /guilds/{guild.id}/roles/{role.id}
// https://discord.com/developers/docs/resources/guild#modify-guild-role
type ModifyGuildRole struct {
	GuildID      Snowflake `json:"-"`
	RoleID       Snowflake `json:"-"`
	Name         **string  `json:"name,omitempty"`
	Permissions  **string  `json:"permissions,omitempty"`
	Color        **int     `json:"color,omitempty"`
	Hoist        **bool    `json:"hoist,omitempty"`
	Icon         **string  `json:"icon,omitempty"`
	UnicodeEmoji **string  `json:"unicode_emoji,omitempty"`
	Mentionable  **bool    `json:"mentionable,omitempty"`
}

// Modify Guild MFA Level
// POST /guilds/{guild.id}/mfa
// https://discord.com/developers/docs/resources/guild#modify-guild-mfa-level
type ModifyGuildMFALevel struct {
	GuildID Snowflake `json:"-"`
	Level   Flag      `json:"level"`
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
	GuildID      Snowflake   `url:"-"`
	Days         int         `url:"days,omitempty"`
	IncludeRoles []Snowflake `url:"include_roles,omitempty"`
}

// Begin Guild Prune
// POST /guilds/{guild.id}/prune
// https://discord.com/developers/docs/resources/guild#begin-guild-prune
type BeginGuildPrune struct {
	GuildID           Snowflake   `json:"-"`
	Days              int         `json:"days"`
	ComputePruneCount bool        `json:"compute_prune_count"`
	IncludeRoles      []Snowflake `json:"include_roles"`
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
	GuildID Snowflake `json:"-"`
	Code    *string   `json:"code"`
	Uses    int       `json:"uses,omitempty"`
}

// Get Guild Widget Image
// GET /guilds/{guild.id}/widget.png
// https://discord.com/developers/docs/resources/guild#get-guild-widget-image
type GetGuildWidgetImage struct {
	GuildID Snowflake `url:"-"`

	// Widget Style Options
	// https://discord.com/developers/docs/resources/guild#get-guild-widget-image-widget-style-options
	Style *string `url:"style,omitempty"`
}

// Widget Style Options
// https://discord.com/developers/docs/resources/guild#get-guild-widget-image-widget-style-options
const (
	FlagWidgetStyleOptionShield  = "shield"
	FlagWidgetStyleOptionBanner1 = "banner1"
	FlagWidgetStyleOptionBanner2 = "banner2"
	FlagWidgetStyleOptionBanner3 = "banner3"
	FlagWidgetStyleOptionBanner4 = "banner4"
)

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
	GuildID         Snowflake                `json:"-"`
	Enabled         **bool                   `json:"enabled,omitempty"`
	WelcomeChannels *[]*WelcomeScreenChannel `json:"welcome_channels,omitempty"`
	Description     **string                 `json:"description,omitempty"`
}

// Get Guild Onboarding
// GET /guilds/{guild.id}/onboarding
// https://discord.com/developers/docs/resources/guild#get-guild-onboarding
type GetGuildOnboarding struct {
	GuildID Snowflake
}

// Modify Guild Onboarding
// PUT /guilds/{guild.id}/onboarding
// https://discord.com/developers/docs/resources/guild#modify-guild-onboarding
type ModifyGuildOnboarding struct {
	GuildID           Snowflake           `json:"-"`
	Prompts           []*OnboardingPrompt `json:"prompts"`
	DefaultChannelIDs []Snowflake         `json:"default_channel_ids"`
	Enabled           bool                `json:"enabled"`
	Mode              Flag                `json:"mode"`
}

// Modify Current User Voice State
// PATCH /guilds/{guild.id}/voice-states/@me
// https://discord.com/developers/docs/resources/guild#modify-current-user-voice-state
type ModifyCurrentUserVoiceState struct {
	GuildID                 Snowflake   `json:"-"`
	ChannelID               *Snowflake  `json:"channel_id,omitempty"`
	Suppress                *bool       `json:"suppress,omitempty"`
	RequestToSpeakTimestamp **Timestamp `json:"request_to_speak_timestamp,omitempty"`
}

// Modify User Voice State
// PATCH /guilds/{guild.id}/voice-states/{user.id}
// https://discord.com/developers/docs/resources/guild#modify-user-voice-state
type ModifyUserVoiceState struct {
	GuildID   Snowflake `json:"-"`
	UserID    Snowflake `json:"-"`
	ChannelID Snowflake `json:"channel_id"`
	Suppress  *bool     `json:"suppress,omitempty"`
}

// List Scheduled Events for Guild
// GET /guilds/{guild.id}/scheduled-events
// https://discord.com/developers/docs/resources/guild-scheduled-event#list-scheduled-events-for-guild
type ListScheduledEventsforGuild struct {
	GuildID       Snowflake `url:"-"`
	WithUserCount *bool     `url:"with_user_count,omitempty"`
}

// Create Guild Scheduled Event
// POST /guilds/{guild.id}/scheduled-events
// https://discord.com/developers/docs/resources/guild-scheduled-event#create-guild-scheduled-event
type CreateGuildScheduledEvent struct {
	GuildID            Snowflake                          `json:"-"`
	ChannelID          *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name               string                             `json:"name"`
	PrivacyLevel       Flag                               `json:"privacy_level"`
	ScheduledStartTime Timestamp                          `json:"scheduled_start_time"`
	ScheduledEndTime   *Timestamp                         `json:"scheduled_end_time,omitempty"`
	Description        *string                            `json:"description,omitempty"`
	EntityType         *Flag                              `json:"entity_type,omitempty"`
	Image              *string                            `json:"image,omitempty"`
}

// Get Guild Scheduled Event
// GET /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}
// https://discord.com/developers/docs/resources/guild-scheduled-event#get-guild-scheduled-event
type GetGuildScheduledEvent struct {
	GuildID               Snowflake `url:"-"`
	GuildScheduledEventID Snowflake `url:"-"`
	WithUserCount         *bool     `url:"with_user_count,omitempty"`
}

// Modify Guild Scheduled Event
// PATCH /guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}
// https://discord.com/developers/docs/resources/guild-scheduled-event#modify-guild-scheduled-event
type ModifyGuildScheduledEvent struct {
	GuildID               Snowflake                           `json:"-"`
	GuildScheduledEventID Snowflake                           `json:"-"`
	ChannelID             *Snowflake                          `json:"channel_id,omitempty"`
	EntityMetadata        **GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	Name                  *string                             `json:"name,omitempty"`
	PrivacyLevel          *Flag                               `json:"privacy_level,omitempty"`
	ScheduledStartTime    *Timestamp                          `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime      *Timestamp                          `json:"scheduled_end_time,omitempty"`
	Description           **string                            `json:"description,omitempty"`
	EntityType            *Flag                               `json:"entity_type,omitempty"`
	Status                *Flag                               `json:"status,omitempty"`
	Image                 *string                             `json:"image,omitempty"`
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
	GuildID               Snowflake  `url:"-"`
	GuildScheduledEventID Snowflake  `url:"-"`
	Limit                 *int       `url:"limit,omitempty"`
	WithMember            *bool      `url:"with_member,omitempty"`
	Before                *Snowflake `url:"before,omitempty"`
	After                 *Snowflake `url:"after,omitempty"`
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
	TemplateCode string  `json:"-"`
	Name         string  `json:"name"`
	Icon         *string `json:"icon,omitempty"`
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
	GuildID     Snowflake `json:"-"`
	Name        string    `json:"name"`
	Description **string  `json:"description,omitempty"`
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
	TemplateCode string   `json:"-"`
	Name         *string  `json:"name,omitempty"`
	Description  **string `json:"description,omitempty"`
}

// Delete Guild Template
// DELETE /guilds/{guild.id}/templates/{template.code}
// https://discord.com/developers/docs/resources/guild-template#delete-guild-template
type DeleteGuildTemplate struct {
	GuildID      Snowflake
	TemplateCode string
}

// Get Invite
// GET /invites/{invite.code}
// https://discord.com/developers/docs/resources/invite#get-invite
type GetInvite struct {
	InviteCode            string     `url:"-"`
	WithCounts            *bool      `url:"with_counts,omitempty"`
	WithExpiration        *bool      `url:"with_expiration,omitempty"`
	GuildScheduledEventID *Snowflake `url:"guild_scheduled_event_id,omitempty"`
}

// Delete Invite
// DELETE /invites/{invite.code}
// https://discord.com/developers/docs/resources/invite#delete-invite
type DeleteInvite struct {
	InviteCode string
}

// Create Stage Instance
// POST /stage-instances
// https://discord.com/developers/docs/resources/stage-instance#create-stage-instance
type CreateStageInstance struct {
	ChannelID             Snowflake `json:"channel_id"`
	Topic                 string    `json:"topic"`
	PrivacyLevel          *Flag     `json:"privacy_level,omitempty"`
	SendStartNotification *bool     `json:"send_start_notification,omitempty"`
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
	ChannelID    Snowflake `json:"-"`
	Topic        *string   `json:"topic,omitempty"`
	PrivacyLevel *Flag     `json:"privacy_level,omitempty"`
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
type ListNitroStickerPacks struct{}

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
	GuildID     Snowflake `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tags        *string   `json:"tags"`
	File        File      `json:"-" dasgo:"file"`
}

// Modify Guild Sticker
// PATCH /guilds/{guild.id}/stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#modify-guild-sticker
type ModifyGuildSticker struct {
	GuildID     Snowflake `json:"-"`
	StickerID   Snowflake `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Description **string  `json:"description,omitempty"`
	Tags        *string   `json:"tags,omitempty"`
}

// Delete Guild Sticker
// DELETE /guilds/{guild.id}/stickers/{sticker.id}
// https://discord.com/developers/docs/resources/sticker#delete-guild-sticker
type DeleteGuildSticker struct {
	GuildID   Snowflake
	StickerID Snowflake
}

// Get Current User
// GET/users/@me
// https://discord.com/developers/docs/resources/user#get-current-user
type GetCurrentUser struct{}

// Get User
// GET/users/{user.id}
// https://discord.com/developers/docs/resources/user#get-user
type GetUser struct {
	UserID Snowflake
}

// Modify Current User
// PATCH /users/@me
// https://discord.com/developers/docs/resources/user#modify-current-user
type ModifyCurrentUser struct {
	Username *string `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

// Get Current User Guilds
// GET /users/@me/guilds
// https://discord.com/developers/docs/resources/user#get-current-user-guilds
type GetCurrentUserGuilds struct {
	Before     *Snowflake `json:"before,omitempty"`
	After      *Snowflake `json:"after,omitempty"`
	Limit      *int       `json:"limit,omitempty"`
	WithCounts *bool      `json:"with_counts,omitempty"`
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
	RecipientID Snowflake `json:"recipient_id"`
}

// Create Group DM
// POST /users/@me/channels
// https://discord.com/developers/docs/resources/user#create-group-dm
type CreateGroupDM struct {
	AccessTokens []*string            `json:"access_tokens"`
	Nicks        map[Snowflake]string `json:"nicks"`
}

// Get User Connections
// GET /users/@me/connections
// https://discord.com/developers/docs/resources/user#get-user-connections
type GetUserConnections struct{}

// Get User Application Role Connection
// GET /users/@me/applications/{application.id}/role-connection
// https://discord.com/developers/docs/resources/user#get-user-application-role-connection
type GetUserApplicationRoleConnection struct {
	ApplicationID Snowflake
}

// Update User Application Role Connection
// PUT /users/@me/applications/{application.id}/role-connection
// https://discord.com/developers/docs/resources/user#update-user-application-role-connection
type UpdateUserApplicationRoleConnection struct {
	ApplicationID    Snowflake         `json:"-"`
	PlatformName     *string           `json:"platform_name,omitempty"`
	PlatformUsername *string           `json:"platform_user,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}

// List Voice Regions
// GET /voice/regions
// https://discord.com/developers/docs/resources/voice#list-voice-regions
type ListVoiceRegions struct{}

// Create Webhook
// POST /channels/{channel.id}/webhooks
// https://discord.com/developers/docs/resources/webhook#create-webhook
type CreateWebhook struct {
	ChannelID Snowflake `json:"-"`
	Name      string    `json:"name"`
	Avatar    **string  `json:"avatar,omitempty"`
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
	WebhookID Snowflake  `json:"-"`
	Name      *string    `json:"name,omitempty"`
	Avatar    **string   `json:"avatar,omitempty"`
	ChannelID *Snowflake `json:"channel_id,omitempty"`
}

// Modify Webhook with Token
// PATCH /webhooks/{webhook.id}/{webhook.token}
// https://discord.com/developers/docs/resources/webhook#modify-webhook-with-token
type ModifyWebhookwithToken struct {
	WebhookID    Snowflake
	WebhookToken string
	Name         *string  `json:"name,omitempty"`
	Avatar       **string `json:"avatar,omitempty"`
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
	WebhookID       Snowflake        `json:"-" url:"-"`
	WebhookToken    string           `json:"-" url:"-"`
	Wait            *bool            `json:"-" url:"wait,omitempty"`
	ThreadID        *Snowflake       `json:"-" url:"thread_id,omitempty"`
	Content         *string          `json:"content,omitempty" url:"-"`
	Username        *string          `json:"username,omitempty" url:"-"`
	AvatarURL       *string          `json:"avatar_url,omitempty" url:"-"`
	TTS             *bool            `json:"tts,omitempty" url:"-"`
	Embeds          []*Embed         `json:"embeds,omitempty" url:"-"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty" url:"-"`
	Components      []Component      `json:"components,omitempty" url:"-"`
	Files           []*File          `json:"-" url:"-" dasgo:"files"`
	PayloadJSON     string           `json:"-" url:"-" multipart:"payload_json"`
	Attachments     []*Attachment    `json:"attachments,omitempty" url:"-"`
	Flags           *BitFlag         `json:"flags,omitempty" url:"-"`
	ThreadName      *string          `json:"thread_name,omitempty" url:"-"`
}

// Execute Slack-Compatible Webhook
// POST /webhooks/{webhook.id}/{webhook.token}/slack
// https://discord.com/developers/docs/resources/webhook#execute-slackcompatible-webhook
type ExecuteSlackCompatibleWebhook struct {
	WebhookID    Snowflake  `url:"-"`
	WebhookToken string     `url:"-"`
	ThreadID     *Snowflake `url:"thread_id,omitempty"`
	Wait         *bool      `url:"wait,omitempty"`
}

// Execute GitHub-Compatible Webhook
// POST /webhooks/{webhook.id}/{webhook.token}/github
// https://discord.com/developers/docs/resources/webhook#execute-githubcompatible-webhook
type ExecuteGitHubCompatibleWebhook struct {
	WebhookID    Snowflake  `url:"-"`
	WebhookToken string     `url:"-"`
	ThreadID     *Snowflake `url:"thread_id,omitempty"`
	Wait         *bool      `url:"wait,omitempty"`
}

// Get Webhook Message
// GET /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#get-webhook-message
type GetWebhookMessage struct {
	WebhookID    Snowflake  `url:"-"`
	WebhookToken string     `url:"-"`
	MessageID    Snowflake  `url:"-"`
	ThreadID     *Snowflake `url:"thread_id,omitempty"`
}

// Edit Webhook Message
// PATCH /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message
type EditWebhookMessage struct {
	WebhookID       Snowflake         `json:"-" url:"-"`
	WebhookToken    string            `json:"-" url:"-"`
	MessageID       Snowflake         `json:"-" url:"-"`
	ThreadID        *Snowflake        `url:"thread_id,omitempty"`
	Content         **string          `json:"content,omitempty" url:"-"`
	Embeds          *[]*Embed         `json:"embeds,omitempty" url:"-"`
	Components      *[]Component      `json:"components,omitempty" url:"-"`
	Files           []*File           `json:"-" url:"-" dasgo:"files"`
	AllowedMentions **AllowedMentions `json:"allowed_mentions,omitempty" url:"-"`
	PayloadJSON     string            `json:"-" url:"-" multipart:"payload_json"`
	Attachments     *[]*Attachment    `json:"attachments,omitempty" url:"-"`
}

// Delete Webhook Message
// DELETE /webhooks/{webhook.id}/{webhook.token}/messages/{message.id}
// https://discord.com/developers/docs/resources/webhook#delete-webhook-message
type DeleteWebhookMessage struct {
	WebhookID    Snowflake  `url:"-"`
	WebhookToken string     `url:"-"`
	MessageID    Snowflake  `url:"-"`
	ThreadID     *Snowflake `url:"thread_id,omitempty"`
}

// Get Current Bot Application Information
// GET /oauth2/applications/@me
// https://discord.com/developers/docs/topics/oauth2#get-current-bot-application-information
type GetCurrentBotApplicationInformation struct{}

// Get Current Authorization Information
// GET /oauth2/@me
// https://discord.com/developers/docs/topics/oauth2#get-current-authorization-information
type GetCurrentAuthorizationInformation struct{}

// Get Gateway
// GET /gateway
// https://discord.com/developers/docs/topics/gateway#get-gateway
type GetGateway struct{}

// Get Gateway Bot
// GET /gateway/bot
// https://discord.com/developers/docs/topics/gateway#get-gateway-bot
type GetGatewayBot struct{}

// Authorization URL
// GET /oauth2/authorize
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-authorization-url-example
type AuthorizationURL struct {
	ResponseType string    `url:"response_type,omitempty"`
	ClientID     Snowflake `url:"client_id,omitempty"`
	Scope        string    `url:"scope,omitempty"`
	State        string    `url:"state,omitempty"`
	RedirectURI  string    `url:"redirect_uri,omitempty"`
	Prompt       string    `url:"prompt,omitempty"`
}

// Access Token Exchange
// POST /oauth2/token
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-access-token-exchange-example
type AccessTokenExchange struct {
	ClientID     Snowflake `url:"client_id,omitempty"`
	ClientSecret string    `url:"client_secret,omitempty"`
	GrantType    string    `url:"grant_type,omitempty"`
	Code         string    `url:"code,omitempty"`
	RedirectURI  string    `url:"redirect_uri,omitempty"`
}

// Refresh Token Exchange
// POST /oauth2/token
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-refresh-token-exchange-example
type RefreshTokenExchange struct {
	ClientID     Snowflake `url:"client_id,omitempty"`
	ClientSecret string    `url:"client_secret,omitempty"`
	GrantType    string    `url:"grant_type,omitempty"`
	RefreshToken string    `url:"refresh_token,omitempty"`
}

// Client Credentials Token Request
// POST /oauth2/token
// https://discord.com/developers/docs/topics/oauth2#client-credentials-grant-client-credentials-token-request-example
type ClientCredentialsTokenRequest struct {
	GrantType string `url:"grant_type,omitempty"`
	Scope     string `url:"scope,omitempty"`
}

// Bot Auth Parameters
// GET /oauth2/authorize
// https://discord.com/developers/docs/topics/oauth2#bot-authorization-flow-bot-auth-parameters
type BotAuth struct {
	ClientID           Snowflake `url:"client_id"`
	Scope              string    `url:"scope"`
	Permissions        BitFlag   `url:"permissions"`
	GuildID            Snowflake `url:"guild_id"`
	DisableGuildSelect bool      `url:"disable_guild_select"`
}
