package dasgo

import (
	"encoding/json"
)

// Application Command Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	ID                       Snowflake                   `json:"id"`
	Type                     *Flag                       `json:"type,omitempty"`
	ApplicationID            Snowflake                   `json:"application_id"`
	GuildID                  *Snowflake                  `json:"guild_id,omitempty"`
	Name                     string                      `json:"name"`
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`
	Description              string                      `json:"description"`
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                     `json:"default_member_permissions"`
	NSFW                     *bool                       `json:"nsfw,omitempty"`
	IntegrationTypes         []Flag                      `json:"integration_types,omitempty"`
	Contexts                 *[]Flag                     `json:"contexts,omitempty"`
	Version                  Snowflake                   `json:"version,omitempty"`
	Handler                  Flag                        `json:"handler,omitempty"`
}

// Application Command Types
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
const (
	FlagApplicationCommandTypeCHAT_INPUT          Flag = 1
	FlagApplicationCommandTypeUSER                Flag = 2
	FlagApplicationCommandTypeMESSAGE             Flag = 3
	FlagApplicationCommandTypePRIMARY_ENTRY_POINT Flag = 4
)

// Application Command Option Structure
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     Flag                              `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        *map[string]string                `json:"name_localizations,omitempty"`
	Description              string                            `json:"description"`
	DescriptionLocalizations *map[string]string                `json:"description_localizations,omitempty"`
	Required                 *bool                             `json:"required,omitempty"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []*ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             Flags                             `json:"channel_types,omitempty"`
	MinValue                 *float64                          `json:"min_value,omitempty"`
	MaxValue                 *float64                          `json:"max_value,omitempty"`
	MinLength                *int                              `json:"min_length,omitempty"`
	MaxLength                *int                              `json:"max_length,omitempty"`
	Autocomplete             *bool                             `json:"autocomplete,omitempty"`
}

// Application Command Option Type
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
const (
	FlagApplicationCommandOptionTypeSUB_COMMAND       Flag = 1
	FlagApplicationCommandOptionTypeSUB_COMMAND_GROUP Flag = 2
	FlagApplicationCommandOptionTypeSTRING            Flag = 3
	FlagApplicationCommandOptionTypeINTEGER           Flag = 4
	FlagApplicationCommandOptionTypeBOOLEAN           Flag = 5
	FlagApplicationCommandOptionTypeUSER              Flag = 6
	FlagApplicationCommandOptionTypeCHANNEL           Flag = 7
	FlagApplicationCommandOptionTypeROLE              Flag = 8
	FlagApplicationCommandOptionTypeMENTIONABLE       Flag = 9
	FlagApplicationCommandOptionTypeNUMBER            Flag = 10
	FlagApplicationCommandOptionTypeATTACHMENT        Flag = 11
)

// Application Command Option Choice
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure
type ApplicationCommandOptionChoice struct {
	Name              string             `json:"name"`
	NameLocalizations *map[string]string `json:"name_localizations,omitempty"`
	Value             Value              `json:"value"`
}

// Entry Point Command Handler Types
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-entry-point-command-handler-types
const (
	FlagEntryPointCommandHandlerTypesAPP_HANDLER             Flag = 1
	FlagEntryPointCommandHandlerTypesDISCORD_LAUNCH_ACTIVITY Flag = 2
)

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
	FlagApplicationCommandPermissionTypeROLE    Flag = 1
	FlagApplicationCommandPermissionTypeUSER    Flag = 2
	FlagApplicationCommandPermissionTypeCHANNEL Flag = 3
)

// Component Object
// https://discord.com/developers/docs/interactions/message-components#component-object
type Component interface {
	ComponentType() Flag
}

// Component Types
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
const (
	FlagComponentTypeInvalid           Flag = 0
	FlagComponentTypeActionRow         Flag = 1
	FlagComponentTypeButton            Flag = 2
	FlagComponentTypeStringSelect      Flag = 3
	FlagComponentTypeTextInput         Flag = 4
	FlagComponentTypeUserSelect        Flag = 5
	FlagComponentTypeRoleSelect        Flag = 6
	FlagComponentTypeMentionableSelect Flag = 7
	FlagComponentTypeChannelSelect     Flag = 8
)

func (c ActionRow) ComponentType() Flag {
	return FlagComponentTypeActionRow
}

func (c Button) ComponentType() Flag {
	return FlagComponentTypeButton
}

func (c SelectMenu) ComponentType() Flag {
	return FlagComponentTypeInvalid
}

func (c SelectMenu) FlagComponentTypeStringSelect() Flag {
	return FlagComponentTypeStringSelect
}

func (c SelectMenu) ComponentTypeUserSelect() Flag {
	return FlagComponentTypeUserSelect
}

func (c SelectMenu) ComponentTypeRoleSelect() Flag {
	return FlagComponentTypeRoleSelect
}

func (c SelectMenu) ComponentTypeMentionableSelect() Flag {
	return FlagComponentTypeMentionableSelect
}

func (c SelectMenu) ComponentTypeChannelSelect() Flag {
	return FlagComponentTypeChannelSelect
}

func (c TextInput) ComponentType() Flag {
	return FlagComponentTypeTextInput
}

// https://discord.com/developers/docs/interactions/message-components#action-rows
type ActionRow struct {
	Type       Flag        `json:"type"`
	Components []Component `json:"components"`
}

// Button Object
// https://discord.com/developers/docs/interactions/message-components#button-object
type Button struct {
	Type     Flag       `json:"type"`
	Style    Flag       `json:"style"`
	Label    *string    `json:"label,omitempty"`
	Emoji    *Emoji     `json:"emoji,omitempty"`
	CustomID *string    `json:"custom_id,omitempty"`
	SKUID    *Snowflake `json:"sku_id,omitempty"`
	URL      *string    `json:"url,omitempty"`
	Disabled *bool      `json:"disabled,omitempty"`
}

// Button Styles
// https://discord.com/developers/docs/interactions/message-components#button-object-button-styles
const (
	FlagButtonStylePRIMARY   Flag = 1
	FlagButtonStyleBLURPLE   Flag = 1
	FlagButtonStyleSecondary Flag = 2
	FlagButtonStyleGREY      Flag = 2
	FlagButtonStyleSuccess   Flag = 3
	FlagButtonStyleGREEN     Flag = 3
	FlagButtonStyleDanger    Flag = 4
	FlagButtonStyleRED       Flag = 4
	FlagButtonStyleLINK      Flag = 5
	FlagButtonStylePremium   Flag = 6
)

// Select Menu Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure
type SelectMenu struct {
	Type          Flag                  `json:"type"`
	CustomID      string                `json:"custom_id"`
	Options       []*SelectMenuOption   `json:"options"`
	ChannelTypes  Flags                 `json:"channel_types,omitempty"`
	Placeholder   *string               `json:"placeholder,omitempty"`
	DefaultValues []*SelectDefaultValue `json:"default_values,omitempty"`
	MinValues     *Flag                 `json:"min_values,omitempty"`
	MaxValues     *Flag                 `json:"max_values,omitempty"`
	Disabled      *bool                 `json:"disabled,omitempty"`
}

// Select Menu Option Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectMenuOption struct {
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Description *string `json:"description,omitempty"`
	Emoji       *Emoji  `json:"emoji,omitempty"`
	Default     *bool   `json:"default,omitempty"`
}

// Select Default Value Structure
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-default-value-structure
type SelectDefaultValue struct {
	ID   Snowflake `json:"id"`
	Type string    `json:"type"`
}

// Text Input Structure
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInput struct {
	Type        Flag    `json:"type"`
	Style       Flag    `json:"style"`
	CustomID    string  `json:"custom_id"`
	Label       *string `json:"label"`
	MinLength   *int    `json:"min_length,omitempty"`
	MaxLength   *int    `json:"max_length,omitempty"`
	Required    *bool   `json:"required,omitempty"`
	Value       *string `json:"value,omitempty"`
	Placeholder *string `json:"placeholder,omitempty"`
}

// Text Input Styles
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-styles
const (
	FlagTextInputStyleShort     Flag = 1
	FlagTextInputStyleParagraph Flag = 2
)

// Interaction Object
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	ID                           Snowflake          `json:"id"`
	ApplicationID                Snowflake          `json:"application_id"`
	Type                         Flag               `json:"type"`
	Data                         InteractionData    `json:"data,omitempty"`
	Guild                        *Guild             `json:"guild,omitempty"`
	GuildID                      *Snowflake         `json:"guild_id,omitempty"`
	Channel                      *Channel           `json:"channel,omitempty"`
	ChannelID                    *Snowflake         `json:"channel_id,omitempty"`
	Member                       *GuildMember       `json:"member,omitempty"`
	User                         *User              `json:"user,omitempty"`
	Token                        string             `json:"token"`
	Version                      int                `json:"version,omitempty"`
	Message                      *Message           `json:"message,omitempty"`
	AppPermissions               *BitFlag           `json:"app_permissions,omitempty,string"`
	Locale                       *string            `json:"locale,omitempty"`
	GuildLocale                  *string            `json:"guild_locale,omitempty"`
	Entitlements                 []*Entitlement     `json:"entitlement"`
	AuthorizingIntegrationOwners map[Flag]Snowflake `json:"authorizing_integration_owners"`
	Context                      *Flag              `json:"context,omitempty"`
}

// Interaction Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
const (
	FlagInteractionTypePING                             Flag = 1
	FlagInteractionTypeAPPLICATION_COMMAND              Flag = 2
	FlagInteractionTypeMESSAGE_COMPONENT                Flag = 3
	FlagInteractionTypeAPPLICATION_COMMAND_AUTOCOMPLETE Flag = 4
	FlagInteractionTypeMODAL_SUBMIT                     Flag = 5
)

// Interaction Context Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-context-types
const (
	FlagInteractionContextTypeGUILD           Flag = 0
	FlagInteractionContextTypeBOT_DM          Flag = 1
	FlagInteractionContextTypePRIVATE_CHANNEL      = 2
)

// Interaction Data
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data
type InteractionData interface {
	InteractionDataType() Flag
}

func (d ApplicationCommandData) InteractionDataType() Flag {
	return FlagInteractionTypeAPPLICATION_COMMAND
}

func (d MessageComponentData) InteractionDataType() Flag {
	return FlagInteractionTypeMESSAGE_COMPONENT
}

func (d ModalSubmitData) InteractionDataType() Flag {
	return FlagInteractionTypeMODAL_SUBMIT
}

// Application Command Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-application-command-data-structure
type ApplicationCommandData struct {
	ID       Snowflake                                  `json:"id"`
	Name     string                                     `json:"name"`
	Type     Flag                                       `json:"type"`
	Resolved *ResolvedData                              `json:"resolved,omitempty"`
	Options  []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildID  *Snowflake                                 `json:"guild_id,omitempty"`
	TargetID *Snowflake                                 `json:"target_id,omitempty"`
}

// Message Component Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-message-component-data-structure
type MessageComponentData struct {
	CustomID      string              `json:"custom_id"`
	ComponentType Flag                `json:"component_type"`
	Values        []*SelectMenuOption `json:"values,omitempty"`
	Resolved      *ResolvedData       `json:"resolved,omitempty"`
}

// Modal Submit Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-modal-submit-data-structure
type ModalSubmitData struct {
	CustomID   string      `json:"custom_id"`
	Components []Component `json:"components"`
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

// Application Command Interaction Data Option Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    Flag                                       `json:"type"`
	Value   *Value                                     `json:"value,omitempty"`
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused *bool                                      `json:"focused,omitempty"`
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
	Type Flag                    `json:"type"`
	Data InteractionCallbackData `json:"data,omitempty"`
}

// Interaction Callback Type
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
const (
	FlagInteractionCallbackTypePONG                                    Flag = 1
	FlagInteractionCallbackTypeCHANNEL_MESSAGE_WITH_SOURCE             Flag = 4
	FlagInteractionCallbackTypeDEFERRED_CHANNEL_MESSAGE_WITH_SOURCE    Flag = 5
	FlagInteractionCallbackTypeDEFERRED_UPDATE_MESSAGE                 Flag = 6
	FlagInteractionCallbackTypeUPDATE_MESSAGE                          Flag = 7
	FlagInteractionCallbackTypeAPPLICATION_COMMAND_AUTOCOMPLETE_RESULT Flag = 8
	FlagInteractionCallbackTypeMODAL                                   Flag = 9
	FlagInteractionCallbackTypeLAUNCH_ACTIVITY                         Flag = 12
)

// Interaction Callback Data Structure
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionCallbackData interface {
	InteractionCallbackDataType() Flag
}

func (d Messages) InteractionCallbackDataType() Flag {
	return FlagInteractionCallbackTypeCHANNEL_MESSAGE_WITH_SOURCE
}

func (d Autocomplete) InteractionCallbackDataType() Flag {
	return FlagInteractionCallbackTypeAPPLICATION_COMMAND_AUTOCOMPLETE_RESULT
}

func (d Modal) InteractionCallbackDataType() Flag {
	return FlagInteractionCallbackTypeMODAL
}

// Messages
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-messages
type Messages struct {
	TTS             *bool            `json:"tts,omitempty"`
	Content         *string          `json:"content,omitempty"`
	Embeds          []*Embed         `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           *BitFlag         `json:"flags,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []*Attachment    `json:"attachments,omitempty"`
	Poll            *Poll            `json:"poll,omitempty"`
}

// Autocomplete
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-autocomplete
type Autocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices"`
}

// Modal
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-modal
type Modal struct {
	CustomID   string      `json:"custom_id"`
	Title      string      `json:"title"`
	Components []Component `json:"components"`
}

// Application Object
// https://discord.com/developers/docs/resources/application
type Application struct {
	ID                             Snowflake                                         `json:"id"`
	Name                           string                                            `json:"name"`
	Icon                           *string                                           `json:"icon"`
	Description                    string                                            `json:"description"`
	RPCOrigins                     []string                                          `json:"rpc_origins,omitempty"`
	BotPublic                      bool                                              `json:"bot_public"`
	BotRequireCodeGrant            bool                                              `json:"bot_require_code_grant"`
	Bot                            *User                                             `json:"bot,omitempty"`
	TermsOfServiceURL              *string                                           `json:"terms_of_service_url,omitempty"`
	PrivacyProxyURL                *string                                           `json:"privacy_policy_url,omitempty"`
	Owner                          *User                                             `json:"owner,omitempty"`
	VerifyKey                      string                                            `json:"verify_key"`
	Team                           *Team                                             `json:"team"`
	GuildID                        *Snowflake                                        `json:"guild_id,omitempty"`
	Guild                          *Guild                                            `json:"guild,omitempty"`
	PrimarySKUID                   *Snowflake                                        `json:"primary_sku_id,omitempty"`
	Slug                           *string                                           `json:"slug,omitempty"`
	CoverImage                     *string                                           `json:"cover_image,omitempty"`
	Flags                          *BitFlag                                          `json:"flags,omitempty"`
	ApproximateGuildCount          *int                                              `json:"approximate_guild_count,omitempty"`
	ApproximateUserInstallCount    *int                                              `json:"approximate_user_install_count,omitempty"`
	RedirectURIs                   []string                                          `json:"redirect_uris,omitempty"`
	InteractionsEndpointURL        **string                                          `json:"interactions_endpoint_url,omitempty"`
	RoleConnectionsVerificationURL **string                                          `json:"role_connections_verification_url,omitempty"`
	EventWebhooksURL               **string                                          `json:"event_webhooks_url,omitempty"`
	EventWebhooksStatus            Flag                                              `json:"event_webhooks_status"`
	EventWebhooksTypes             []string                                          `json:"event_webhooks_types,omitempty"`
	Tags                           []string                                          `json:"tags,omitempty"`
	InstallParams                  *InstallParams                                    `json:"install_params,omitempty"`
	IntegrationTypesConfig         map[Flag]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	CustomInstallURL               *string                                           `json:"custom_install_url,omitempty"`
}

// Application Integration Type Configuration Object
// https://discord.com/developers/docs/resources/application#application-object-application-integration-type-configuration-object
type ApplicationIntegrationTypeConfiguration struct {
	OAuth2InstallParams *InstallParams `json:"oauth2_install_params,omitempty"`
}

// Application Integration Types
// https://discord.com/developers/docs/resources/application#application-object-application-integration-types
const (
	FlagApplicationIntegrationTypeGUILD_INSTALL Flag = 0
	FlagApplicationIntegrationTypeUSER_INSTALL  Flag = 1
)

// Application Event Webhook Status
// https://discord.com/developers/docs/resources/application#application-object-application-event-webhook-status
const (
	FlagApplicationEventWebhookStatusDISABLED            Flag = 1
	FlagApplicationEventWebhookStatusENABLED             Flag = 2
	FlagApplicationEventWebhookStatusDISABLED_BY_DISCORD Flag = 3
)

// Application Flags
// https://discord.com/developers/docs/resources/application#application-object-application-flags
const (
	FlagApplicationAPPLICATION_AUTO_MODERATION_RULE_CREATE_BADGE BitFlag = 1 << 6
	FlagApplicationGATEWAY_PRESENCE                              BitFlag = 1 << 12
	FlagApplicationGATEWAY_PRESENCE_LIMITED                      BitFlag = 1 << 13
	FlagApplicationGATEWAY_GUILD_MEMBERS                         BitFlag = 1 << 14
	FlagApplicationGATEWAY_GUILD_MEMBERS_LIMITED                 BitFlag = 1 << 15
	FlagApplicationVERIFICATION_PENDING_GUILD_LIMIT              BitFlag = 1 << 16
	FlagApplicationEMBEDDED                                      BitFlag = 1 << 17
	FlagApplicationGATEWAY_MESSAGE_CONTENT                       BitFlag = 1 << 18
	FlagApplicationGATEWAY_MESSAGE_CONTENT_LIMITED               BitFlag = 1 << 19
	FlagApplicationAPPLICATION_COMMAND_BADGE                     BitFlag = 1 << 23
)

// Install Params Object
// https://discord.com/developers/docs/resources/application#install-params-object
type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}

// Activity Instance Object
// https://discord.com/developers/docs/resources/application#get-application-activity-instance-activity-instance-object
type ActivityInstance struct {
	ApplicationID Snowflake         `json:"application_id"`
	InstanceID    string            `json:"instance_id"`
	LaunchID      Snowflake         `json:"launch_id"`
	Location      *ActivityLocation `json:"location"`
	Users         []Snowflake       `json:"users"`
}

// Activity Location Object
// https://discord.com/developers/docs/resources/application#get-application-activity-instance-activity-location-object
type ActivityLocation struct {
	ID        string      `json:"id"`
	Kind      string      `json:"kind"`
	ChannelID Snowflake   `json:"channel_id"`
	GuildID   **Snowflake `json:"guild_id,omitempty"`
}

// Activity Location Kind Enum
// https://discord.com/developers/docs/resources/application#get-application-activity-instance-activity-location-kind-enum
var (
	ActivityLocationKindEnum = map[string]string{
		"gc": "Location is a Guild Channel",
		"pc": "Location is a Private Channel, such as a DM or GDM",
	}
)

// Application Role Connection Metadata Object
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-structure
type ApplicationRoleConnectionMetadata struct {
	Type                     Flag               `json:"type"`
	Key                      string             `json:"key"`
	Name                     string             `json:"name"`
	NameLocalizations        *map[string]string `json:"name_localizations,omitempty"`
	Description              string             `json:"description"`
	DescriptionLocalizations *map[string]string `json:"description_localizations,omitempty"`
}

// Application Role Connection Metadata Types
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-type
const (
	FlagApplicationRoleConnectionMetadataTypeINTEGER_LESS_THAN_OR_EQUAL     Flag = 1
	FlagApplicationRoleConnectionMetadataTypeINTEGER_GREATER_THAN_OR_EQUAL  Flag = 2
	FlagApplicationRoleConnectionMetadataTypeINTEGER_EQUAL                  Flag = 3
	FlagApplicationRoleConnectionMetadataTypeINTEGER_NOT_EQUAL              Flag = 4
	FlagApplicationRoleConnectionMetadataTypeDATETIME_LESS_THAN_OR_EQUAL    Flag = 5
	FlagApplicationRoleConnectionMetadataTypeDATETIME_GREATER_THAN_OR_EQUAL Flag = 6
	FlagApplicationRoleConnectionMetadataTypeBOOLEAN_EQUAL                  Flag = 7
	FlagApplicationRoleConnectionMetadataTypeBOOLEAN_NOT_EQUAL              Flag = 8
)

// Audit Log Object
// https://discord.com/developers/docs/resources/audit-log
type AuditLog struct {
	ApplicationCommands  []*ApplicationCommand  `json:"application_commands"`
	AuditLogEntries      []*AuditLogEntry       `json:"audit_log_entries"`
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events"`
	Integration          []*Integration         `json:"integrations"`
	Threads              []*Channel             `json:"threads"`
	Users                []*User                `json:"users"`
	Webhooks             []*Webhook             `json:"webhooks"`
}

// Audit Log Entry Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-entry-structure
type AuditLogEntry struct {
	TargetID   *string           `json:"target_id"`
	Changes    []*AuditLogChange `json:"changes,omitempty"`
	UserID     *Snowflake        `json:"user_id"`
	ID         Snowflake         `json:"id"`
	ActionType Flag              `json:"action_type"`
	Options    *AuditLogOptions  `json:"options,omitempty"`
	Reason     *string           `json:"reason,omitempty"`
}

// Audit Log Events
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
const (
	FlagAuditLogEventGUILD_UPDATE                                Flag = 1
	FlagAuditLogEventCHANNEL_CREATE                              Flag = 10
	FlagAuditLogEventCHANNEL_UPDATE                              Flag = 11
	FlagAuditLogEventCHANNEL_DELETE                              Flag = 12
	FlagAuditLogEventCHANNEL_OVERWRITE_CREATE                    Flag = 13
	FlagAuditLogEventCHANNEL_OVERWRITE_UPDATE                    Flag = 14
	FlagAuditLogEventCHANNEL_OVERWRITE_DELETE                    Flag = 15
	FlagAuditLogEventMEMBER_KICK                                 Flag = 20
	FlagAuditLogEventMEMBER_PRUNE                                Flag = 21
	FlagAuditLogEventMEMBER_BAN_ADD                              Flag = 22
	FlagAuditLogEventMEMBER_BAN_REMOVE                           Flag = 23
	FlagAuditLogEventMEMBER_UPDATE                               Flag = 24
	FlagAuditLogEventMEMBER_ROLE_UPDATE                          Flag = 25
	FlagAuditLogEventMEMBER_MOVE                                 Flag = 26
	FlagAuditLogEventMEMBER_DISCONNECT                           Flag = 27
	FlagAuditLogEventBOT_ADD                                     Flag = 28
	FlagAuditLogEventROLE_CREATE                                 Flag = 30
	FlagAuditLogEventROLE_UPDATE                                 Flag = 31
	FlagAuditLogEventROLE_DELETE                                 Flag = 32
	FlagAuditLogEventINVITE_CREATE                               Flag = 40
	FlagAuditLogEventINVITE_UPDATE                               Flag = 41
	FlagAuditLogEventINVITE_DELETE                               Flag = 42
	FlagAuditLogEventWEBHOOK_CREATE                              Flag = 50
	FlagAuditLogEventWEBHOOK_UPDATE                              Flag = 51
	FlagAuditLogEventWEBHOOK_DELETE                              Flag = 52
	FlagAuditLogEventEMOJI_CREATE                                Flag = 60
	FlagAuditLogEventEMOJI_UPDATE                                Flag = 61
	FlagAuditLogEventEMOJI_DELETE                                Flag = 62
	FlagAuditLogEventMESSAGE_DELETE                              Flag = 72
	FlagAuditLogEventMESSAGE_BULK_DELETE                         Flag = 73
	FlagAuditLogEventMESSAGE_PIN                                 Flag = 74
	FlagAuditLogEventMESSAGE_UNPIN                               Flag = 75
	FlagAuditLogEventINTEGRATION_CREATE                          Flag = 80
	FlagAuditLogEventINTEGRATION_UPDATE                          Flag = 81
	FlagAuditLogEventINTEGRATION_DELETE                          Flag = 82
	FlagAuditLogEventSTAGE_INSTANCE_CREATE                       Flag = 83
	FlagAuditLogEventSTAGE_INSTANCE_UPDATE                       Flag = 84
	FlagAuditLogEventSTAGE_INSTANCE_DELETE                       Flag = 85
	FlagAuditLogEventSTICKER_CREATE                              Flag = 90
	FlagAuditLogEventSTICKER_UPDATE                              Flag = 91
	FlagAuditLogEventSTICKER_DELETE                              Flag = 92
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_CREATE                Flag = 100
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_UPDATE                Flag = 101
	FlagAuditLogEventGUILD_SCHEDULED_EVENT_DELETE                Flag = 102
	FlagAuditLogEventTHREAD_CREATE                               Flag = 110
	FlagAuditLogEventTHREAD_UPDATE                               Flag = 111
	FlagAuditLogEventTHREAD_DELETE                               Flag = 112
	FlagAuditLogEventAPPLICATION_COMMAND_PERMISSION_UPDATE       Flag = 121
	FlagAuditLogEventAUTO_MODERATION_RULE_CREATE                 Flag = 140
	FlagAuditLogEventAUTO_MODERATION_RULE_UPDATE                 Flag = 141
	FlagAuditLogEventAUTO_MODERATION_RULE_DELETE                 Flag = 142
	FlagAuditLogEventAUTO_MODERATION_BLOCK_MESSAGE               Flag = 143
	FlagAuditLogEventAUTO_MODERATION_FLAG_TO_CHANNEL             Flag = 144
	FlagAuditLogEventAUTO_MODERATION_USER_COMMUNICATION_DISABLED Flag = 145
	FlagAuditLogEventCREATOR_MONETIZATION_REQUEST_CREATED        Flag = 150
	FlagAuditLogEventCREATOR_MONETIZATION_TERMS_ACCEPTED         Flag = 151
	FlagAuditLogEventONBOARDING_PROMPT_CREATE                    Flag = 163
	FlagAuditLogEventONBOARDING_PROMPT_UPDATE                    Flag = 164
	FlagAuditLogEventONBOARDING_PROMPT_DELETE                    Flag = 165
	FlagAuditLogEventONBOARDING_CREATE                           Flag = 166
	FlagAuditLogEventONBOARDING_UPDATE                           Flag = 167
	FlagAuditLogEventHOME_SETTINGS_CREATE                        Flag = 190
	FlagAuditLogEventHOME_SETTINGS_UPDATE                        Flag = 191
)

// Optional Audit Entry Info
// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogOptions struct {
	ApplicationID                 Snowflake `json:"application_id"`
	AutoModerationRuleName        string    `json:"auto_moderation_rule_name"`
	AutoModerationRuleTriggerType string    `json:"auto_moderation_rule_trigger_type"`
	ChannelID                     Snowflake `json:"channel_id"`
	Count                         string    `json:"count"`
	DeleteMemberDays              string    `json:"delete_member_days"`
	ID                            Snowflake `json:"id"`
	MembersRemoved                string    `json:"members_removed"`
	MessageID                     Snowflake `json:"message_id"`
	RoleName                      string    `json:"role_name"`
	Type                          string    `json:"type"`
}

// Audit Log Change Object
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object
type AuditLogChange struct {
	NewValue json.RawMessage `json:"new_value,omitempty"`
	OldValue json.RawMessage `json:"old_value,omitempty"`
	Key      string          `json:"key"`
}

// Audit Log Change Exceptions
// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-exceptions

// Auto Moderation Rule Structure
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-auto-moderation-rule-structure
type AutoModerationRule struct {
	ID              Snowflake               `json:"id"`
	GuildID         Snowflake               `json:"guild_id"`
	Name            string                  `json:"name"`
	CreatorID       Snowflake               `json:"creator_id"`
	EventType       Flag                    `json:"event_type"`
	TriggerType     Flag                    `json:"trigger_type"`
	TriggerMetadata TriggerMetadata         `json:"trigger_metadata"`
	Actions         []*AutoModerationAction `json:"actions"`
	Enabled         bool                    `json:"enabled"`
	ExemptRoles     []Snowflake             `json:"exempt_roles"`
	ExemptChannels  []Snowflake             `json:"exempt_channels"`
}

// Trigger Types
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-types
const (
	FlagTriggerTypeKEYWORD        Flag = 1
	FlagTriggerTypeHARMFUL_LINK   Flag = 2
	FlagTriggerTypeSPAM           Flag = 3
	FlagTriggerTypeKEYWORD_PRESET Flag = 4
	FlagTriggerTypeMENTION_SPAM   Flag = 5
	FlagTriggerTypeMEMBER_PROFILE Flag = 6
)

// Trigger Metadata
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-metadata
type TriggerMetadata struct {
	// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-keyword-matching-strategies
	KeywordFilter                []string `json:"keyword_filter"`
	RegexPatterns                []string `json:"regex_patterns"`
	Presets                      Flags    `json:"presets"`
	AllowList                    []string `json:"allow_list"`
	MentionTotalLimit            int      `json:"mention_total_limit"`
	MentionRaidProtectionEnabled bool     `json:"mention_raid_protection_enabled"`
}

// Keyword Preset Types
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-keyword-preset-types
const (
	FlagKeywordPresetTypePROFANITY      Flag = 1
	FlagKeywordPresetTypeSEXUAL_CONTENT Flag = 2
	FlagKeywordPresetTypeSLURS          Flag = 3
)

// Event Types
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-event-types
const (
	FlagEventTypeMESSAGE_SEND  Flag = 1
	FlagEventTypeMEMBER_UPDATE Flag = 2
)

// Auto Moderation Action Structure
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object
type AutoModerationAction struct {
	Type     Flag            `json:"type"`
	Metadata *ActionMetadata `json:"metadata,omitempty"`
}

// Action Types
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-types
const (
	FlagActionTypeBLOCK_MESSAGE            Flag = 1
	FlagActionTypeSEND_ALERT_MESSAGE       Flag = 2
	FlagActionTypeTIMEOUT                  Flag = 3
	FlagActionTypeBLOCK_MEMBER_INTERACTION Flag = 4
)

// Action Metadata
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-metadata
type ActionMetadata struct {
	ChannelID       Snowflake `json:"channel_id"`
	DurationSeconds int       `json:"duration_seconds"`
	CustomMessage   *string   `json:"custom_message,omitempty"`
}

// Channel Object
// https://discord.com/developers/docs/resources/channel
type Channel struct {
	ID                            Snowflake              `json:"id"`
	Type                          *Flag                  `json:"type"`
	GuildID                       *Snowflake             `json:"guild_id,omitempty"`
	Position                      *int                   `json:"position,omitempty"`
	PermissionOverwrites          []*PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                          **string               `json:"name,omitempty"`
	Topic                         **string               `json:"topic,omitempty"`
	NSFW                          *bool                  `json:"nsfw,omitempty"`
	LastMessageID                 **Snowflake            `json:"last_message_id,omitempty"`
	Bitrate                       *int                   `json:"bitrate,omitempty"`
	UserLimit                     *int                   `json:"user_limit,omitempty"`
	RateLimitPerUser              *int                   `json:"rate_limit_per_user,omitempty"`
	Recipients                    []*User                `json:"recipients,omitempty"`
	Icon                          **string               `json:"icon,omitempty"`
	OwnerID                       *Snowflake             `json:"owner_id,omitempty"`
	ApplicationID                 *Snowflake             `json:"application_id,omitempty"`
	Managed                       *bool                  `json:"managed,omitempty"`
	ParentID                      **Snowflake            `json:"parent_id,omitempty"`
	LastPinTimestamp              **Timestamp            `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     **string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *Flag                  `json:"video_quality_mode,omitempty"`
	MessageCount                  *int                   `json:"message_count,omitempty"`
	MemberCount                   *int                   `json:"member_count,omitempty"`
	ThreadMetadata                *ThreadMetadata        `json:"thread_metadata,omitempty"`
	Member                        *ThreadMember          `json:"member,omitempty"`
	DefaultAutoArchiveDuration    *int                   `json:"default_auto_archive_duration,omitempty"`
	Permissions                   *string                `json:"permissions,omitempty"`
	Flags                         *BitFlag               `json:"flags,omitempty"`
	TotalMessageSent              *int                   `json:"total_message_sent,omitempty"`
	AvailableTags                 []*ForumTag            `json:"available_tags,omitempty"`
	AppliedTags                   []Snowflake            `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction       `json:"default_reaction_emoji"`
	DefaultThreadRateLimitPerUser *int                   `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              **Flag                 `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *Flag                  `json:"default_forum_layout,omitempty"`
}

// Channel Types
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
const (
	FlagChannelTypeGUILD_TEXT          Flag = 0
	FlagChannelTypeDM                  Flag = 1
	FlagChannelTypeGUILD_VOICE         Flag = 2
	FlagChannelTypeGROUP_DM            Flag = 3
	FlagChannelTypeGUILD_CATEGORY      Flag = 4
	FlagChannelTypeGUILD_ANNOUNCEMENT  Flag = 5
	FlagChannelTypeANNOUNCEMENT_THREAD Flag = 10
	FlagChannelTypePUBLIC_THREAD       Flag = 11
	FlagChannelTypePRIVATE_THREAD      Flag = 12
	FlagChannelTypeGUILD_STAGE_VOICE   Flag = 13
	FlagChannelTypeGUILD_DIRECTORY     Flag = 14
	FlagChannelTypeGUILD_FORUM         Flag = 15
	FlagChannelTypeGUILD_MEDIA         Flag = 16
)

// Video Quality Modes
// https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
const (
	FlagVideoQualityModeAUTO Flag = 1
	FlagVideoQualityModeFULL Flag = 2
)

// Channel Flags
// https://discord.com/developers/docs/resources/channel#channel-object-channel-flags
const (
	FlagChannelPINNED      BitFlag = 1 << 1
	FlagChannelREQUIRE_TAG BitFlag = 1 << 4
)

// Sort Order Types
// https://discord.com/developers/docs/resources/channel#channel-object-sort-order-types
const (
	FlagSortOrderTypeLATEST_ACTIVITY Flag = 0
	FlagSortOrderTypeCREATION_DATE   Flag = 1
)

// Forum Layout Types
// https://discord.com/developers/docs/resources/channel#channel-object-forum-layout-types
const (
	FlagForumLayoutTypeNOT_SET      Flag = 0
	FlagForumLayoutTypeLIST_VIEW    Flag = 1
	FlagForumLayoutTypeGALLERY_VIEW Flag = 2
)

// Followed Channel Structure
// https://discord.com/developers/docs/resources/channel#followed-channel-object-followed-channel-structure
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// Message Object
// https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	ID                   Snowflake                           `json:"id"`
	ChannelID            Snowflake                           `json:"channel_id"`
	Author               *User                               `json:"author"`
	Content              string                              `json:"content"`
	Timestamp            Timestamp                           `json:"timestamp"`
	EditedTimestamp      *Timestamp                          `json:"edited_timestamp"`
	TTS                  bool                                `json:"tts"`
	MentionEveryone      bool                                `json:"mention_everyone"`
	Mentions             []*User                             `json:"mentions"`
	MentionRoles         []*Snowflake                        `json:"mention_roles"`
	MentionChannels      []*ChannelMention                   `json:"mention_channels,omitempty"`
	Attachments          []*Attachment                       `json:"attachments"`
	Embeds               []*Embed                            `json:"embeds"`
	Reactions            []*Reaction                         `json:"reactions,omitempty"`
	Nonce                *Nonce                              `json:"nonce,omitempty"`
	Pinned               bool                                `json:"pinned"`
	WebhookID            *Snowflake                          `json:"webhook_id,omitempty"`
	Type                 Flag                                `json:"type"`
	Activity             *MessageActivity                    `json:"activity,omitempty"`
	Application          *Application                        `json:"application,omitempty"`
	ApplicationID        *Snowflake                          `json:"application_id,omitempty"`
	Flags                *BitFlag                            `json:"flags,omitempty"`
	MessageReference     *MessageReference                   `json:"message_reference,omitempty"`
	MessageSnapshots     []*MessageSnapshot                  `json:"message_snapshots,omitempty"`
	ReferencedMessage    **Message                           `json:"referenced_message,omitempty"`
	InteractionMetadata  MessageComponentInteractionMetadata `json:"interaction_metadata,omitempty"`
	Interaction          *MessageInteraction                 `json:"interaction"`
	Thread               *Channel                            `json:"thread"`
	Components           []Component                         `json:"components"`
	StickerItems         []*StickerItem                      `json:"sticker_items"`
	Position             *int                                `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData               `json:"role_subscription_data,omitempty"`
	Resolved             *ResolvedData                       `json:"resolved,omitempty"`
	Poll                 *Poll                               `json:"poll,omitempty"`
	Call                 *MessageCall                        `json:"call,omitempty"`

	// MessageCreate Event Extra Fields
	// https://discord.com/developers/docs/topics/gateway-events#message-create
	GuildID *Snowflake   `json:"guild_id,omitempty"`
	Member  *GuildMember `json:"member,omitempty"`
}

// Message Types
// https://discord.com/developers/docs/resources/channel#message-object-message-types
const (
	FlagMessageTypeDEFAULT                                      Flag = 0
	FlagMessageTypeRECIPIENT_ADD                                Flag = 1
	FlagMessageTypeRECIPIENT_REMOVE                             Flag = 2
	FlagMessageTypeCALL                                         Flag = 3
	FlagMessageTypeCHANNEL_NAME_CHANGE                          Flag = 4
	FlagMessageTypeCHANNEL_ICON_CHANGE                          Flag = 5
	FlagMessageTypeCHANNEL_PINNED_MESSAGE                       Flag = 6
	FlagMessageTypeUSER_JOIN                                    Flag = 7
	FlagMessageTypeGUILD_BOOST                                  Flag = 8
	FlagMessageTypeGUILD_BOOST_TIER_1                           Flag = 9
	FlagMessageTypeGUILD_BOOST_TIER_2                           Flag = 10
	FlagMessageTypeGUILD_BOOST_TIER_3                           Flag = 11
	FlagMessageTypeCHANNEL_FOLLOW_ADD                           Flag = 12
	FlagMessageTypeGUILD_DISCOVERY_DISQUALIFIED                 Flag = 14
	FlagMessageTypeGUILD_DISCOVERY_REQUALIFIED                  Flag = 15
	FlagMessageTypeGUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING Flag = 16
	FlagMessageTypeGUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING   Flag = 17
	FlagMessageTypeTHREAD_CREATED                               Flag = 18
	FlagMessageTypeREPLY                                        Flag = 19
	FlagMessageTypeCHAT_INPUT_COMMAND                           Flag = 20
	FlagMessageTypeTHREAD_STARTER_MESSAGE                       Flag = 21
	FlagMessageTypeGUILD_INVITE_REMINDER                        Flag = 22
	FlagMessageTypeCONTEXT_MENU_COMMAND                         Flag = 23
	FlagMessageTypeAUTO_MODERATION_ACTION                       Flag = 24
	FlagMessageTypeROLE_SUBSCRIPTION_PURCHASE                   Flag = 25
	FlagMessageTypeINTERACTION_PREMIUM_UPSELL                   Flag = 26
	FlagMessageTypeSTAGE_START                                  Flag = 27
	FlagMessageTypeSTAGE_END                                    Flag = 28
	FlagMessageTypeSTAGE_SPEAKER                                Flag = 29
	FlagMessageTypeSTAGE_RAISE_HAND                             Flag = 30
	FlagMessageTypeSTAGE_TOPIC                                  Flag = 31
	FlagMessageTypeGUILD_APPLICATION_PREMIUM_SUBSCRIPTION       Flag = 32
	FlagMessageTypeGUILD_INCIDENT_ALERT_MODE_ENABLED            Flag = 36
	FlagMessageTypeGUILD_INCIDENT_ALERT_MODE_DISABLED           Flag = 37
	FlagMessageTypeGUILD_INCIDENT_REPORT_RAID                   Flag = 38
	FlagMessageTypeGUILD_INCIDENT_REPORT_FALSE_ALARM            Flag = 39
	FlagMessageTypePURCHASE_NOTIFICATION                        Flag = 44
	FlagMessageTypePOLL_RESULT                                  Flag = 46
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
	FlagMessageActivityTypeJOIN         Flag = 1
	FlagMessageActivityTypeSPECTATE     Flag = 2
	FlagMessageActivityTypeLISTEN       Flag = 3
	FlagMessageActivityTypeJOIN_REQUEST Flag = 5
)

// Message Flags
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
const (
	FlagMessageCROSSPOSTED                            BitFlag = 1 << 0
	FlagMessageIS_CROSSPOST                           BitFlag = 1 << 1
	FlagMessageSUPPRESS_EMBEDS                        BitFlag = 1 << 2
	FlagMessageSOURCE_MESSAGE_DELETED                 BitFlag = 1 << 3
	FlagMessageURGENT                                 BitFlag = 1 << 4
	FlagMessageHAS_THREAD                             BitFlag = 1 << 5
	FlagMessageEPHEMERAL                              BitFlag = 1 << 6
	FlagMessageLOADING                                BitFlag = 1 << 7
	FlagMessageFAILED_TO_MENTION_SOME_ROLES_IN_THREAD BitFlag = 1 << 8
	FlagMessageSUPPRESS_NOTIFICATIONS                 BitFlag = 1 << 12
	FlagMessageIS_VOICE_MESSAGE                       BitFlag = 1 << 13
	FlagMessageHAS_SNAPSHOT                           BitFlag = 1 << 14
)

// Message Interaction Metadata Object
// https://discord.com/developers/docs/resources/message#message-interaction-metadata-object
type MessageInteractionMetadata interface {
	MessageInteractionMetadata() Flag
}

func (d ApplicationCommandInteractionMetadata) MessageInteractionMetadata() Flag {
	return FlagInteractionTypeAPPLICATION_COMMAND
}

func (d MessageComponentInteractionMetadata) MessageInteractionMetadata() Flag {
	return FlagInteractionTypeMESSAGE_COMPONENT
}

func (d ModalSubmitInteractionMetadata) MessageInteractionMetadata() Flag {
	return FlagInteractionTypeMODAL_SUBMIT
}

// Application Command Interaction Metadata Structure
// https://discord.com/developers/docs/resources/message#message-interaction-metadata-object-application-command-interaction-metadata-structure
type ApplicationCommandInteractionMetadata struct {
	ID                           Snowflake          `json:"id"`
	Type                         Flag               `json:"type"`
	User                         *User              `json:"user"`
	AuthorizingIntegrationOwners map[Flag]Snowflake `json:"authorizing_integration_owners"`
	OriginalResponseMessageID    *Snowflake         `json:"original_response_message_id,omitempty"`
	TargetUser                   *User              `json:"target_user,omitempty"`
	TargetMessageID              *Snowflake         `json:"target_message_id,omitempty"`
}

// Message Component Interaction Metadata Structure
// https://discord.com/developers/docs/resources/message#message-interaction-metadata-object-message-component-interaction-metadata-structure
type MessageComponentInteractionMetadata struct {
	ID                           Snowflake          `json:"id"`
	Type                         Flag               `json:"type"`
	User                         *User              `json:"user"`
	AuthorizingIntegrationOwners map[Flag]Snowflake `json:"authorizing_integration_owners"`
	OriginalResponseMessageID    *Snowflake         `json:"original_response_message_id,omitempty"`
	InteractedMessageID          Snowflake          `json:"interacted_message_id"`
}

// Modal Submit Interaction Metadata Structure
// https://discord.com/developers/docs/resources/message#message-interaction-metadata-object-modal-submit-interaction-metadata-structure
type ModalSubmitInteractionMetadata struct {
	ID                            Snowflake                  `json:"id"`
	Type                          Flag                       `json:"type"`
	User                          *User                      `json:"user"`
	AuthorizingIntegrationOwners  map[Flag]Snowflake         `json:"authorizing_integration_owners"`
	OriginalResponseMessageID     *Snowflake                 `json:"original_response_message_id,omitempty"`
	TriggeringInteractionMetadata MessageInteractionMetadata `json:"triggering_interaction_metadata"`
}

// Message Call Object Structure
// https://discord.com/developers/docs/resources/message#message-call-object
type MessageCall struct {
	Participants   []Snowflake `json:"participants"`
	EndedTimestamp **Timestamp `json:"ended_timestamp,omitempty"`
}

// Message Reference Object
// https://discord.com/developers/docs/resources/channel#message-reference-object
type MessageReference struct {
	Type            Flag       `json:"type"`
	MessageID       *Snowflake `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists *bool      `json:"fail_if_not_exists,omitempty"`
}

// Message Reference Types
// https://discord.com/developers/docs/resources/message#message-reference-types
const (
	FlagMessageReferenceTypeDEFAULT Flag = 0
	FlagMessageReferenceTypeFORWARD Flag = 1
)

// Message Snapshot Structure
// https://discord.com/developers/docs/resources/message#message-snapshot-structure
type MessageSnapshot struct {
	Message *Message `json:"message"`
}

// Reaction Object
// https://discord.com/developers/docs/resources/channel#reaction-object
type Reaction struct {
	Count        int                   `json:"count"`
	CountDetails *ReactionCountDetails `json:"count_details"`
	Me           bool                  `json:"me"`
	MeBurst      bool                  `json:"me_burst"`
	Emoji        *Emoji                `json:"emoji"`
	BurstColors  []string              `json:"burst_colors"`
}

// Reaction Count Details Structure
// https://discord.com/developers/docs/resources/message#reaction-count-details-object-reaction-count-details-structure
type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

// Overwrite Object
// https://discord.com/developers/docs/resources/channel#overwrite-object
type PermissionOverwrite struct {
	ID    Snowflake `json:"id"`
	Type  Flag      `json:"type"`
	Deny  string    `json:"deny"`
	Allow string    `json:"allow"`
}

// Thread Metadata Object
// https://discord.com/developers/docs/resources/channel#thread-metadata-object
type ThreadMetadata struct {
	Archived            bool        `json:"archived"`
	AutoArchiveDuration int         `json:"auto_archive_duration"`
	ArchiveTimestamp    Timestamp   `json:"archive_timestamp"`
	Locked              bool        `json:"locked"`
	Invitable           *bool       `json:"invitable,omitempty"`
	CreateTimestamp     **Timestamp `json:"create_timestamp"`
}

// Thread Member Object
// https://discord.com/developers/docs/resources/channel#thread-member-object
type ThreadMember struct {
	ID            *Snowflake   `json:"id,omitempty"`
	UserID        *Snowflake   `json:"user_id,omitempty"`
	JoinTimestamp Timestamp    `json:"join_timestamp"`
	Flags         Flag         `json:"flags"`
	Member        *GuildMember `json:"member,omitempty"`
}

// Default Reaction Structure
// https://discord.com/developers/docs/resources/channel#default-reaction-object-default-reaction-structure
type DefaultReaction struct {
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
}

// Forum Tag Structure
// https://discord.com/developers/docs/resources/channel#forum-tag-object-forum-tag-structure
type ForumTag struct {
	ID        Snowflake  `json:"id"`
	Name      string     `json:"name"`
	Moderated bool       `json:"moderated"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
}

// Embed Object
// https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	Title       *string         `json:"title,omitempty"`
	Type        *string         `json:"type,omitempty"`
	Description *string         `json:"description,omitempty"`
	URL         *string         `json:"url,omitempty"`
	Timestamp   *Timestamp      `json:"timestamp,omitempty"`
	Color       *int            `json:"color,omitempty"`
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
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

// Embed Video Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

// Embed Image Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
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
	Name         string  `json:"name"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

// Embed Footer Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string  `json:"text"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

// Embed Field Structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
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

// Attachment Object
// https://discord.com/developers/docs/resources/message#attachment-object-attachment-structure
type Attachment struct {
	ID              Snowflake `json:"id"`
	Filename        string    `json:"filename"`
	Title           *string   `json:"title,omitempty"`
	Description     *string   `json:"description,omitempty"`
	ContentType     *string   `json:"content_type,omitempty"`
	Size            int       `json:"size"`
	URL             string    `json:"url"`
	ProxyURL        string    `json:"proxy_url"`
	Height          **int     `json:"height,omitempty"`
	Width           **int     `json:"width,omitempty"`
	Emphemeral      *bool     `json:"ephemeral,omitempty"`
	DurationSeconds *float64  `json:"duration_secs,omitempty"`
	Waveform        *string   `json:"waveform,omitempty"`
	Flags           *BitFlag  `json:"flags,omitempty"`
}

// Attachment Flags
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-flags
const (
	IS_REMIX BitFlag = 1 << 2
)

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

// Role Subscription Data Object Structure
// https://discord.com/developers/docs/resources/channel#role-subscription-data-object-role-subscription-data-object-structure
type RoleSubscriptionData struct {
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"`
	TierName                  string    `json:"tier_name"`
	TotalMonthsSubscribed     int       `json:"total_months_subscribed"`
	IsRenewal                 bool      `json:"is_renewal"`
}

// Reaction Types
// https://discord.com/developers/docs/resources/message#get-reactions-reaction-types
const (
	FlagReactionTypeNORMAL = 0
	FlagReactionTypeBURST  = 1
)

// Poll
// https://discord.com/developers/docs/resources/poll
type Poll struct {
	Question         *PollMedia    `json:"question"`
	Answers          []*PollAnswer `json:"answers"`
	Expiry           *Timestamp    `json:"expiry"`
	AllowMultiselect bool          `json:"allow_multiselect"`
	LayoutType       int           `json:"layout_type"`
	Results          *PollResults  `json:"results,omitempty"`
}

// Poll Create Request Object Structure
// https://discord.com/developers/docs/resources/poll#poll-create-request-object-poll-create-request-object-structure
type PollCreateRequest struct {
	Question         *PollMedia    `json:"question"`
	Answers          []*PollAnswer `json:"answers"`
	Duration         *int          `json:"expiry,omitempty"`
	AllowMultiselect *bool         `json:"allow_multiselect,omitempty"`
	LayoutType       *int          `json:"layout_type,omitempty"`
}

// Poll Layout Types
// https://discord.com/developers/docs/resources/poll#layout-type
const (
	FlagPollLayoutTypeDEFAULT Flag = 1
)

// Poll Media Object Structure
// https://discord.com/developers/docs/resources/poll#layout-type
type PollMedia struct {
	Text  *string `json:"text,omitempty"`
	Emoji *Emoji  `json:"emoji,omitempty"`
}

// Poll Answer Object Structure
// https://discord.com/developers/docs/resources/poll#poll-answer-object-poll-answer-object-structure
type PollAnswer struct {
	AnswerID  int        `json:"answer_id"`
	PollMedia *PollMedia `json:"poll_media"`
}

// Poll Results Object Structure
// https://discord.com/developers/docs/resources/poll#poll-results-object-poll-results-object-structure
type PollResults struct {
	IsFinalized  bool               `json:"is_finalized"`
	AnswerCounts []*PollAnswerCount `json:"answer_counts"`
}

// Poll Answer Count Object Structure
// https://discord.com/developers/docs/resources/poll#poll-results-object-poll-answer-count-object-structure
type PollAnswerCount struct {
	ID      int  `json:"id"`
	Count   int  `json:"count"`
	MeVoted bool `json:"me_voted"`
}

// SKU Structure
// https://discord.com/developers/docs/resources/sku#sku-object-sku-structure
type SKU struct {
	ID            Snowflake `json:"id"`
	Type          Flag      `json:"type"`
	ApplicationID Snowflake `json:"application_id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Flags         BitFlag   `json:"flags"`
}

// SKU Types
// https://discord.com/developers/docs/resources/sku#sku-object-sku-types
const (
	FlagSKUTypeDURABLE            Flag = 2
	FlagSKUTypeCONSUMABLE         Flag = 3
	FlagSKUTypeSUBSCRIPTION       Flag = 5
	FlagSKUTypeSUBSCRIPTION_GROUP Flag = 6
)

// SKU Flags
// https://discord.com/developers/docs/resources/sku#sku-object-sku-flags
const (
	FlagSKUFlagAVAILABLE          BitFlag = 1 << 2
	FlagSKUFlagGUILD_SUBSCRIPTION BitFlag = 1 << 7
	FlagSKUFlagUSER_SUBSCRIPTION  BitFlag = 1 << 8
)

// Soundboard Sound Structure
// https://discord.com/developers/docs/resources/soundboard#soundboard-sound-object-soundboard-sound-structure
type SoundboardSound struct {
	Name      string     `json:"name"`
	SoundID   Snowflake  `json:"sound_id"`
	Volume    float32    `json:"volume"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Available bool       `json:"available"`
	User      *User      `json:"user,omitempty"`
}

// Emoji Object
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	ID            *Snowflake  `json:"id"`
	Name          *string     `json:"name,omitempty"`
	Roles         []Snowflake `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons *bool       `json:"require_colons,omitempty"`
	Managed       *bool       `json:"managed,omitempty"`
	Animated      *bool       `json:"animated,omitempty"`
	Available     *bool       `json:"available,omitempty"`
}

// Entitlement Structure
// https://discord.com/developers/docs/resources/entitlement#entitlement-object-entitlement-structure
type Entitlement struct {
	ID            Snowflake  `json:"id"`
	SKUID         Snowflake  `json:"sku_id"`
	ApplicationID Snowflake  `json:"application_id"`
	UserID        *Snowflake `json:"user_id,omitempty"`
	Type          Flag       `json:"type"`
	Deleted       bool       `json:"deleted"`
	StartsAt      *Timestamp `json:"starts_at"`
	EndsAt        *Timestamp `json:"ends_at"`
	GuildID       *Snowflake `json:"guild_id,omitempty"`
	Consumed      *bool      `json:"consumed,omitempty"`
}

// Entitlement Types
// https://discord.com/developers/docs/resources/entitlement#entitlement-object-entitlement-types
const (
	FlagEntitlementTypePURCHASE                 Flag = 1
	FlagEntitlementTypePREMIUM_SUBSCRIPTION     Flag = 2
	FlagEntitlementTypeDEVELOPER_GIFT           Flag = 3
	FlagEntitlementTypeTEST_MODE_PURCHASE       Flag = 4
	FlagEntitlementTypeFREE_PURCHASE            Flag = 5
	FlagEntitlementTypeUSER_GIFT                Flag = 6
	FlagEntitlementTypePREMIUM_PURCHASE         Flag = 7
	FlagEntitlementTypeAPPLICATION_SUBSCRIPTION Flag = 8
)

// Guild Object
// https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	ID                          Snowflake      `json:"id"`
	Name                        string         `json:"name"`
	Icon                        *string        `json:"icon"`
	IconHash                    **string       `json:"icon_hash,omitempty"`
	Splash                      *string        `json:"splash"`
	DiscoverySplash             *string        `json:"discovery_splash"`
	Owner                       *bool          `json:"owner,omitempty"`
	OwnerID                     Snowflake      `json:"owner_id"`
	Permissions                 *string        `json:"permissions,omitempty"`
	AfkChannelID                *Snowflake     `json:"afk_channel_id"`
	AfkTimeout                  int            `json:"afk_timeout"`
	WidgetEnabled               *bool          `json:"widget_enabled,omitempty"`
	WidgetChannelID             **Snowflake    `json:"widget_channel_id,omitempty"`
	VerificationLevel           Flag           `json:"verification_level"`
	DefaultMessageNotifications Flag           `json:"default_message_notifications"`
	ExplicitContentFilter       Flag           `json:"explicit_content_filter"`
	Roles                       []*Role        `json:"roles"`
	Emojis                      []*Emoji       `json:"emojis"`
	Features                    []*string      `json:"features"`
	MFALevel                    Flag           `json:"mfa_level"`
	ApplicationID               *Snowflake     `json:"application_id"`
	SystemChannelID             *Snowflake     `json:"system_channel_id"`
	SystemChannelFlags          BitFlag        `json:"system_channel_flags"`
	RulesChannelID              *Snowflake     `json:"rules_channel_id"`
	MaxPresences                **int          `json:"max_presences,omitempty"`
	MaxMembers                  *int           `json:"max_members,omitempty"`
	VanityUrl                   *string        `json:"vanity_url_code"`
	Description                 *string        `json:"description"`
	Banner                      *string        `json:"banner"`
	PremiumTier                 Flag           `json:"premium_tier"`
	PremiumSubscriptionCount    *int           `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale"`
	PublicUpdatesChannelID      *Snowflake     `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        *int           `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   *int           `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      *int           `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    *int           `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NSFWLevel                   Flag           `json:"nsfw_level"`
	Stickers                    []*Sticker     `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelID       *Snowflake     `json:"safety_alerts_channel_id"`

	// Unavailable Guild Object
	// https://discord.com/developers/docs/resources/guild#unavailable-guild-object
	Unavailable *bool `json:"unavailable,omitempty"`
}

// Default Message Notification Level
// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
const (
	FlagDefaultMessageNotificationLevelALL_MESSAGES  Flag = 0
	FlagDefaultMessageNotificationLevelONLY_MENTIONS Flag = 1
)

// Explicit Content Filter Level
// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
const (
	FlagExplicitContentFilterLevelDISABLED              Flag = 0
	FlagExplicitContentFilterLevelMEMBERS_WITHOUT_ROLES Flag = 1
	FlagExplicitContentFilterLevelALL_MEMBERS           Flag = 2
)

// MFA Level
// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
const (
	FlagMFALevelNONE     Flag = 0
	FlagMFALevelELEVATED Flag = 1
)

// Verification Level
// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
const (
	FlagVerificationLevelNONE      Flag = 0
	FlagVerificationLevelLOW       Flag = 1
	FlagVerificationLevelMEDIUM    Flag = 2
	FlagVerificationLevelHIGH      Flag = 3
	FlagVerificationLevelVERY_HIGH Flag = 4
)

// Guild NSFW Level
// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
const (
	FlagGuildNSFWLevelDEFAULT        Flag = 0
	FlagGuildNSFWLevelEXPLICIT       Flag = 1
	FlagGuildNSFWLevelSAFE           Flag = 2
	FlagGuildNSFWLevelAGE_RESTRICTED Flag = 3
)

// Premium Tier
// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
const (
	FlagPremiumTierNONE  Flag = 0
	FlagPremiumTierONE   Flag = 1
	FlagPremiumTierTWO   Flag = 2
	FlagPremiumTierTHREE Flag = 3
)

// System Channel Flags
// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
const (
	FlagSystemChannelSUPPRESS_JOIN_NOTIFICATIONS                              BitFlag = 1 << 0
	FlagSystemChannelSUPPRESS_PREMIUM_SUBSCRIPTIONS                           BitFlag = 1 << 1
	FlagSystemChannelSUPPRESS_GUILD_REMINDER_NOTIFICATIONS                    BitFlag = 1 << 2
	FlagSystemChannelSUPPRESS_JOIN_NOTIFICATION_REPLIES                       BitFlag = 1 << 3
	FlagSystemChannelSUPPRESS_ROLE_SUBSCRIPTION_PURCHASE_NOTIFICATIONS        BitFlag = 1 << 4
	FlagSystemChannelSUPPRESS_ROLE_SUBSCRIPTION_PURCHASE_NOTIFICATION_REPLIES BitFlag = 1 << 5
)

// Guild Features
// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
const (
	FlagGuildFeatureANIMATED_BANNER                           = "ANIMATED_BANNER"
	FlagGuildFeatureANIMATED_ICON                             = "ANIMATED_ICON"
	FlagGuildFeatureAPPLICATION_COMMAND_PERMISSIONS_V2        = "APPLICATION_COMMAND_PERMISSIONS_V2"
	FlagGuildFeatureAUTO_MODERATION                           = "AUTO_MODERATION"
	FlagGuildFeatureBANNER                                    = "BANNER"
	FlagGuildFeatureCOMMUNITY                                 = "COMMUNITY"
	FlagGuildFeatureCREATOR_MONETIZABLE_PROVISIONAL           = "CREATOR_MONETIZABLE_PROVISIONAL"
	FlagGuildFeatureCREATOR_STORE_PAGE                        = "CREATOR_STORE_PAGE"
	FlagGuildFeatureDEVELOPER_SUPPORT_SERVER                  = "DEVELOPER_SUPPORT_SERVER"
	FlagGuildFeatureDISCOVERABLE                              = "DISCOVERABLE"
	FlagGuildFeatureFEATURABLE                                = "FEATURABLE"
	FlagGuildFeatureINVITES_DISABLED                          = "INVITES_DISABLED"
	FlagGuildFeatureINVITE_SPLASH                             = "INVITE_SPLASH"
	FlagGuildFeatureMEMBER_VERIFICATION_GATE_ENABLED          = "MEMBER_VERIFICATION_GATE_ENABLED"
	FlagGuildFeatureMORE_STICKERS                             = "MORE_STICKERS"
	FlagGuildFeatureNEWS                                      = "NEWS"
	FlagGuildFeaturePARTNERED                                 = "PARTNERED"
	FlagGuildFeaturePREVIEW_ENABLED                           = "PREVIEW_ENABLED"
	FlagGuildFeatureRAID_ALERTS_DISABLED                      = "RAID_ALERTS_DISABLED"
	FlagGuildFeatureROLE_ICONS                                = "ROLE_ICONS"
	FlagGuildFeatureROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	FlagGuildFeatureROLE_SUBSCRIPTIONS_ENABLED                = "ROLE_SUBSCRIPTIONS_ENABLED"
	FlagGuildFeatureTICKETED_EVENTS_ENABLED                   = "TICKETED_EVENTS_ENABLED"
	FlagGuildFeatureVANITY_URL                                = "VANITY_URL"
	FlagGuildFeatureVERIFIED                                  = "VERIFIED"
	FlagGuildFeatureVIP_REGIONS                               = "VIP_REGIONS"
	FlagGuildFeatureWELCOME_SCREEN_ENABLED                    = "WELCOME_SCREEN_ENABLED"
)

// Mutable Guild Features
// https://discord.com/developers/docs/resources/guild#guild-object-mutable-guild-features
var (
	MutableGuildFeatures = map[string]bool{
		FlagGuildFeatureCOMMUNITY:            true,
		FlagGuildFeatureINVITES_DISABLED:     true,
		FlagGuildFeatureDISCOVERABLE:         true,
		FlagGuildFeatureRAID_ALERTS_DISABLED: true,
	}
)

// Guild Preview Object
// https://discord.com/developers/docs/resources/guild#guild-preview-object-guild-preview-structure
type GuildPreview struct {
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Icon                     *string    `json:"icon"`
	Splash                   *string    `json:"splash"`
	DiscoverySplash          *string    `json:"discovery_splash"`
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
	Enabled   bool       `json:"enabled"`
	ChannelID *Snowflake `json:"channel_id"`
}

// Guild Widget Object
// https://discord.com/developers/docs/resources/guild#et-gguild-widget-object-get-guild-widget-structure*
type GuildWidget struct {
	ID            Snowflake  `json:"id"`
	Name          string     `json:"name"`
	InstantInvite *string    `json:"instant_invite"`
	Channels      []*Channel `json:"channels"`
	Members       []*User    `json:"members"`
	PresenceCount int        `json:"presence_count"`
}

// Guild Member Object
// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User                       *User                  `json:"user,omitempty"`
	Nick                       **string               `json:"nick,omitempty"`
	Avatar                     **string               `json:"avatar,omitempty"`
	Banner                     **string               `json:"banner,omitempty"`
	Roles                      []*Snowflake           `json:"roles"`
	JoinedAt                   Timestamp              `json:"joined_at"`
	PremiumSince               **Timestamp            `json:"premium_since,omitempty"`
	Deaf                       bool                   `json:"deaf"`
	Mute                       bool                   `json:"mute"`
	Flags                      BitFlag                `json:"flags"`
	Pending                    *bool                  `json:"pending,omitempty"`
	Permissions                *string                `json:"permissions,omitempty"`
	CommunicationDisabledUntil **Timestamp            `json:"communication_disabled_until,omitempty"`
	AvatarDecorationData       **AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
}

// Guild Member Flags
// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-flags
const (
	FlagGuildMemberDID_REJOIN                      BitFlag = 1 << 0
	FlagGuildMemberCOMPLETED_ONBOARDING            BitFlag = 1 << 1
	FlagGuildMemberBYPASSES_VERIFICATION           BitFlag = 1 << 2
	FlagGuildMemberSTARTED_ONBOARDING              BitFlag = 1 << 3
	FlagGuildMemberIS_GUEST                        BitFlag = 1 << 4
	FlagGuildMemberSTARTED_HOME_ACTIONS            BitFlag = 1 << 5
	FlagGuildMemberCOMPLETED_HOME_ACTIONS          BitFlag = 1 << 6
	FlagGuildMemberAUTOMOD_QUARANTINED_USERNAME    BitFlag = 1 << 7
	FlagGuildMemberDM_SETTINGS_UPSELL_ACKNOWLEDGED BitFlag = 1 << 9
)

// Integration Object
// https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	ID                Snowflake          `json:"id"`
	Name              string             `json:"name"`
	Type              string             `json:"type"`
	Enabled           bool               `json:"enabled"`
	Syncing           *bool              `json:"syncing,omitempty"`
	RoleID            *Snowflake         `json:"role_id,omitempty"`
	EnableEmoticons   *bool              `json:"enable_emoticons,omitempty"`
	ExpireBehavior    *Flag              `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *int               `json:"expire_grace_period,omitempty"`
	User              *User              `json:"user,omitempty"`
	Account           IntegrationAccount `json:"account"`
	SyncedAt          *Timestamp         `json:"synced_at,omitempty"`
	SubscriberCount   *int               `json:"subscriber_count,omitempty"`
	Revoked           *bool              `json:"revoked,omitempty"`
	Application       *Application       `json:"application,omitempty"`
	Scopes            []string           `json:"scopes,omitempty"`
}

// Integration Expire Behaviors
// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
const (
	FlagIntegrationExpireBehaviorREMOVEROLE Flag = 0
	FlagIntegrationExpireBehaviorKICK       Flag = 1
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
	Icon        *string   `json:"icon"`
	Description string    `json:"description"`
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

// Guild Onboarding Structure
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object
type GuildOnboarding struct {
	GuildID           Snowflake           `json:"guild_id"`
	Prompts           []*OnboardingPrompt `json:"prompt"`
	DefaultChannelIDs []Snowflake         `json:"default_channel_ids"`
	Enabled           bool                `json:"enabled"`
	Mode              Flag                `json:"mode"`
}

// Onboarding Prompt Structure
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-onboarding-prompt-structure
type OnboardingPrompt struct {
	ID           Snowflake       `json:"id"`
	Type         Flag            `json:"type"`
	Options      []*PromptOption `json:"options"`
	Title        string          `json:"title"`
	SingleSelect bool            `json:"single_select"`
	Required     bool            `json:"required"`
	InOnboarding bool            `json:"in_onboarding"`
}

// Prompt Option Structure
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-prompt-option-structure
type PromptOption struct {
	ID            Snowflake   `json:"id"`
	ChannelIDs    []Snowflake `json:"channel_ids"`
	RoleIDs       []Snowflake `json:"role_ids"`
	Emoji         []*Emoji    `json:"emoji"`
	EmojiID       *Snowflake  `json:"emoji_id,omitempty"`
	EmojiName     *string     `json:"emoji_name,omitempty"`
	EmojiAnimated *bool       `json:"emoji_animated,omitempty"`
	Title         string      `json:"title"`
	Description   *string     `json:"description"`
}

// Onboarding Mode
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-onboarding-mode
const (
	ONBOARDING_DEFAULT  Flag = 0
	ONBOARDING_ADVANCED Flag = 1
)

// Prompt Types
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-prompt-types
const (
	FlagPromptTypeMULTIPLE_CHOICE = 0
	FlagPromptTypeDROPDOWN        = 1
)

// Guild Scheduled Event Object
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	ID                 Snowflake                          `json:"id"`
	GuildID            Snowflake                          `json:"guild_id"`
	ChannelID          *Snowflake                         `json:"channel_id"`
	CreatorID          **Snowflake                        `json:"creator_id,omitempty"`
	Name               string                             `json:"name"`
	Description        **string                           `json:"description,omitempty"`
	ScheduledStartTime Timestamp                          `json:"scheduled_start_time"`
	ScheduledEndTime   *Timestamp                         `json:"scheduled_end_time"`
	PrivacyLevel       Flag                               `json:"privacy_level"`
	Status             Flag                               `json:"status"`
	EntityType         Flag                               `json:"entity_type"`
	EntityID           *Snowflake                         `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator,omitempty"`
	UserCount          *int                               `json:"user_count,omitempty"`
	Image              **string                           `json:"image,omitempty"`
	RecurrenceRule     *GuildScheduledEventRecurrenceRule `json:"recurrence_rule"`
}

// Guild Scheduled Event Privacy Level
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-privacy-level
const (
	FlagGuildScheduledEventPrivacyLevelGUILD_ONLY Flag = 2
)

// Guild Scheduled Event Entity Types
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-types
const (
	FlagGuildScheduledEventEntityTypeSTAGE_INSTANCE Flag = 1
	FlagGuildScheduledEventEntityTypeVOICE          Flag = 2
	FlagGuildScheduledEventEntityTypeEXTERNAL       Flag = 3
)

// Guild Scheduled Event Status
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-status
const (
	FlagGuildScheduledEventStatusSCHEDULED Flag = 1
	FlagGuildScheduledEventStatusACTIVE    Flag = 2
	FlagGuildScheduledEventStatusCOMPLETED Flag = 3
	FlagGuildScheduledEventStatusCANCELED  Flag = 4
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

// Guild Scheduled Event Recurrence Rule Structure
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-structure
type GuildScheduledEventRecurrenceRule struct {
	Start      Timestamp                                    `json:"start"`
	End        *Timestamp                                   `json:"end"`
	Frequency  Flag                                         `json:"frequency"`
	Interval   int                                          `json:"interval"`
	ByWeekday  []Flag                                       `json:"by_weekday"`
	ByNWeekday []*GuildScheduledEventRecurrenceRuleNWeekday `json:"by_n_weekday"`
	ByMonth    []Flag                                       `json:"by_month"`
	ByMonthDay []int                                        `json:"by_month_day"`
	ByYearDay  []int                                        `json:"by_year_day"`
	Count      *int                                         `json:"count"`
}

// Guild Scheduled Event Recurrence Rule - Frequency
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-frequency
const (
	FlagGuildScheduledEventRecurrenceRuleFrequencyYEARLY  Flag = 0
	FlagGuildScheduledEventRecurrenceRuleFrequencyMONTHLY Flag = 1
	FlagGuildScheduledEventRecurrenceRuleFrequencyWEEKLY  Flag = 2
	FlagGuildScheduledEventRecurrenceRuleFrequencyDAILY   Flag = 3
)

// Guild Scheduled Event Recurrence Rule - Weekday
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-weekday
const (
	FlagGuildScheduledEventRecurrenceRuleWeekdayMONDAY    Flag = 0
	FlagGuildScheduledEventRecurrenceRuleWeekdayTUESDAY   Flag = 1
	FlagGuildScheduledEventRecurrenceRuleWeekdayWEDNESDAY Flag = 2
	FlagGuildScheduledEventRecurrenceRuleWeekdayTHURSDAY  Flag = 3
	FlagGuildScheduledEventRecurrenceRuleWeekdayFRIDAY    Flag = 4
	FlagGuildScheduledEventRecurrenceRuleWeekdaySATURDAY  Flag = 5
	FlagGuildScheduledEventRecurrenceRuleWeekdaySUNDAY    Flag = 6
)

// Guild Scheduled Event Recurrence Rule - N_Weekday Structure
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-weekday
type GuildScheduledEventRecurrenceRuleNWeekday struct {
	N   int  `json:"n"`
	Day Flag `json:"day"`
}

// Guild Scheduled Event Recurrence Rule - Month
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-month
const (
	FlagGuildScheduledEventRecurrenceRuleJANUARY   Flag = 1
	FlagGuildScheduledEventRecurrenceRuleFEBRUARY  Flag = 2
	FlagGuildScheduledEventRecurrenceRuleMARCH     Flag = 3
	FlagGuildScheduledEventRecurrenceRuleAPRIL     Flag = 4
	FlagGuildScheduledEventRecurrenceRuleMAY       Flag = 5
	FlagGuildScheduledEventRecurrenceRuleJUNE      Flag = 6
	FlagGuildScheduledEventRecurrenceRuleJULY      Flag = 7
	FlagGuildScheduledEventRecurrenceRuleAUGUST    Flag = 8
	FlagGuildScheduledEventRecurrenceRuleSEPTEMBER Flag = 9
	FlagGuildScheduledEventRecurrenceRuleOCTOBER   Flag = 10
	FlagGuildScheduledEventRecurrenceRuleNOVEMBER  Flag = 11
	FlagGuildScheduledEventRecurrenceRuleDECEMBER  Flag = 12
)

// Guild Template Object
// https://discord.com/developers/docs/resources/guild-template#guild-template-object
type GuildTemplate struct {
	Code                  string    `json:"code"`
	Name                  string    `json:"name"`
	Description           *string   `json:"description"`
	UsageCount            int       `json:"usage_count"`
	CreatorID             Snowflake `json:"creator_id"`
	Creator               *User     `json:"creator"`
	CreatedAt             Timestamp `json:"created_at"`
	UpdatedAt             Timestamp `json:"updated_at"`
	SourceGuildID         Snowflake `json:"source_guild_id"`
	SerializedSourceGuild *Guild    `json:"serialized_source_guild"`
	IsDirty               *bool     `json:"is_dirty"`
}

// Invite Object
// https://discord.com/developers/docs/resources/invite#invite-object
type Invite struct {
	Type                     Flag                 `json:"type"`
	Code                     string               `json:"code"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               *Flag                `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount *int                 `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   *int                 `json:"approximate_member_count,omitempty"`
	ExpiresAt                **Timestamp          `json:"expires_at,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// Invite Types
// https://discord.com/developers/docs/resources/invite#invite-object-invite-types
const (
	FlagInviteTypesGUILD    Flag = 0
	FlagInviteTypesGROUP_DM Flag = 1
	FlagInviteTypesFRIEND   Flag = 2
)

// Invite Target Types
// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
const (
	FlagInviteTargetTypeSTREAM               Flag = 1
	FlagInviteTargetTypeEMBEDDED_APPLICATION Flag = 2
)

// Invite Metadata Object
// https://discord.com/developers/docs/resources/invite#invite-metadata-object-invite-metadata-structure
type InviteMetadata struct {
	Uses      int       `json:"uses"`
	MaxUses   int       `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt Timestamp `json:"created_at"`
}

// Stage Instance Object
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
	ID                    Snowflake  `json:"id"`
	GuildID               Snowflake  `json:"guild_id"`
	ChannelID             Snowflake  `json:"channel_id"`
	Topic                 string     `json:"topic"`
	PrivacyLevel          Flag       `json:"privacy_level"`
	DiscoverableDisabled  bool       `json:"discoverable_disabled"`
	GuildScheduledEventID *Snowflake `json:"guild_scheduled_event_id"`
}

// Stage Instance Privacy Level
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
const (
	FlagStageInstancePrivacyLevelGUILD_ONLY Flag = 2
)

// Sticker Structure
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	ID          Snowflake  `json:"id"`
	PackID      *Snowflake `json:"pack_id,omitempty"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Tags        string     `json:"tags"`
	Type        Flag       `json:"type"`
	FormatType  Flag       `json:"format_type"`
	Available   *bool      `json:"available,omitempty"`
	GuildID     *Snowflake `json:"guild_id,omitempty"`
	User        *User      `json:"user,omitempty"`
	SortValue   *int       `json:"sort_value,omitempty"`
}

// Sticker Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
const (
	FlagStickerTypeSTANDARD Flag = 1
	FlagStickerTypeGUILD    Flag = 2
)

// Sticker Format Types
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
const (
	FlagStickerFormatTypePNG    Flag = 1
	FlagStickerFormatTypeAPNG   Flag = 2
	FlagStickerFormatTypeLOTTIE Flag = 3
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
	CoverStickerID *Snowflake `json:"cover_sticker_id,omitempty"`
	Description    string     `json:"description"`
	BannerAssetID  *Snowflake `json:"banner_asset_id,omitempty"`
}

// Subscription Object
// https://discord.com/developers/docs/resources/subscription#subscription-object
type Subscription struct {
	ID                 Snowflake   `json:"id"`
	UserID             Snowflake   `json:"user_id"`
	SKUIDs             []Snowflake `json:"sku_ids"`
	EntitlementIDs     []Snowflake `json:"entitlement_ids"`
	RenewalSKUIDs      []Snowflake `json:"renewal_sku_ids"`
	CurrentPeriodStart Timestamp   `json:"current_period_start"`
	CurrentPeriodEnd   Timestamp   `json:"current_period_end"`
	Status             Flag        `json:"status"`
	CanceledAt         *Timestamp  `json:"canceled_at"`
	Country            *string     `json:"country,omitempty"`
}

// Subscription Statuses
// https://discord.com/developers/docs/resources/subscription#subscription-statuses
const (
	FlagSubscriptionStatusACTIVE   Flag = 0
	FlagSubscriptionStatusENDING   Flag = 1
	FlagSubscriptionStatusINACTIVE Flag = 2
)

// User Object
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID               Snowflake              `json:"id"`
	Username         string                 `json:"username"`
	Discriminator    string                 `json:"discriminator"`
	GlobalName       *string                `json:"global_name"`
	Avatar           *string                `json:"avatar"`
	Bot              *bool                  `json:"bot,omitempty"`
	System           *bool                  `json:"system,omitempty"`
	MFAEnabled       *bool                  `json:"mfa_enabled,omitempty"`
	Banner           **string               `json:"banner,omitempty"`
	AccentColor      **int                  `json:"accent_color,omitempty"`
	Locale           *string                `json:"locale,omitempty"`
	Verified         *bool                  `json:"verified,omitempty"`
	Email            **string               `json:"email,omitempty"`
	Flags            *BitFlag               `json:"flag,omitempty"`
	PremiumType      *Flag                  `json:"premium_type,omitempty"`
	PublicFlags      *BitFlag               `json:"public_flag,omitempty"`
	AvatarDecoration **AvatarDecorationData `json:"avatar_decoration,omitempty"`
}

// User Flags
// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	FlagUserNONE                         BitFlag = 0
	FlagUserSTAFF                        BitFlag = 1 << 0
	FlagUserPARTNER                      BitFlag = 1 << 1
	FlagUserHYPESQUAD                    BitFlag = 1 << 2
	FlagUserBUG_HUNTER_LEVEL_1           BitFlag = 1 << 3
	FlagUserHYPESQUAD_ONLINE_HOUSE_ONE   BitFlag = 1 << 6
	FlagUserHYPESQUAD_ONLINE_HOUSE_TWO   BitFlag = 1 << 7
	FlagUserHYPESQUAD_ONLINE_HOUSE_THREE BitFlag = 1 << 8
	FlagUserPREMIUM_EARLY_SUPPORTER      BitFlag = 1 << 9
	FlagUserTEAM_PSEUDO_USER             BitFlag = 1 << 10
	FlagUserBUG_HUNTER_LEVEL_2           BitFlag = 1 << 14
	FlagUserVERIFIED_BOT                 BitFlag = 1 << 16
	FlagUserVERIFIED_DEVELOPER           BitFlag = 1 << 17
	FlagUserCERTIFIED_MODERATOR          BitFlag = 1 << 18
	FlagUserBOT_HTTP_INTERACTIONS        BitFlag = 1 << 19
	FlagUserACTIVE_DEVELOPER             BitFlag = 1 << 22
)

// Premium Types
// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	FlagPremiumTypeNONE         Flag = 0
	FlagPremiumTypeNITROCLASSIC Flag = 1
	FlagPremiumTypeNITRO        Flag = 2
	FlagPremiumTypeNITROBASIC   Flag = 3
)

// Avatar Decoration Data Structure
// https://discord.com/developers/docs/resources/user#avatar-decoration-data-object-avatar-decoration-data-structure
type AvatarDecorationData struct {
	Asset string    `json:"asset"`
	SKUID Snowflake `json:"sku_id"`
}

// User Connection Object
// https://discord.com/developers/docs/resources/user#connection-object-connection-structure
type Connection struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Revoked      *bool          `json:"revoked,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
	Verified     bool           `json:"verified"`
	FriendSync   bool           `json:"friend_sync"`
	ShowActivity bool           `json:"show_activity"`
	TwoWayLink   bool           `json:"two_way_link"`
	Visibility   Flag           `json:"visibility"`
}

// Visibility Types
// https://discord.com/developers/docs/resources/user#connection-object-visibility-types
const (
	FlagVisibilityTypeNONE     Flag = 0
	FlagVisibilityTypeEVERYONE Flag = 1
)

// Application Role Connection Structure
// https://discord.com/developers/docs/resources/user#application-role-connection-object
type ApplicationRoleConnection struct {
	PlatformName     *string           `json:"platform_name"`
	PlatformUsername *string           `json:"platform_user"`
	Metadata         map[string]string `json:"metadata"`
}

// Voice State Object
// https://discord.com/developers/docs/resources/voice#voice-state-object-voice-state-structure
type VoiceState struct {
	GuildID                 *Snowflake   `json:"guild_id,omitempty"`
	ChannelID               *Snowflake   `json:"channel_id"`
	UserID                  Snowflake    `json:"user_id"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              *bool        `json:"self_stream,omitempty"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp *Timestamp   `json:"request_to_speak_timestamp"`
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
	ID            Snowflake   `json:"id"`
	Type          Flag        `json:"type"`
	GuildID       **Snowflake `json:"guild_id,omitempty"`
	ChannelID     *Snowflake  `json:"channel_id"`
	User          *User       `json:"user,omitempty"`
	Name          *string     `json:"name"`
	Avatar        *string     `json:"avatar"`
	Token         *string     `json:"token,omitempty"`
	ApplicationID *Snowflake  `json:"application_id"`
	SourceGuild   *Guild      `json:"source_guild,omitempty"`
	SourceChannel *Channel    `json:"source_channel,omitempty"`
	URL           *string     `json:"url,omitempty"`
}

// Webhook Types
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
const (
	FlagWebhookTypeINCOMING        Flag = 1
	FlagWebhookTypeCHANNELFOLLOWER Flag = 2
	FlagWebhookTypeAPPLICATION     Flag = 3
)

// Bitwise Permission Flags
// https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
const (
	FlagBitwisePermissionCREATE_INSTANT_INVITE               BitFlag = 1 << 0
	FlagBitwisePermissionKICK_MEMBERS                        BitFlag = 1 << 1
	FlagBitwisePermissionBAN_MEMBERS                         BitFlag = 1 << 2
	FlagBitwisePermissionADMINISTRATOR                       BitFlag = 1 << 3
	FlagBitwisePermissionMANAGE_CHANNELS                     BitFlag = 1 << 4
	FlagBitwisePermissionMANAGE_GUILD                        BitFlag = 1 << 5
	FlagBitwisePermissionADD_REACTIONS                       BitFlag = 1 << 6
	FlagBitwisePermissionVIEW_AUDIT_LOG                      BitFlag = 1 << 7
	FlagBitwisePermissionPRIORITY_SPEAKER                    BitFlag = 1 << 8
	FlagBitwisePermissionSTREAM                              BitFlag = 1 << 9
	FlagBitwisePermissionVIEW_CHANNEL                        BitFlag = 1 << 10
	FlagBitwisePermissionSEND_MESSAGES                       BitFlag = 1 << 11
	FlagBitwisePermissionSEND_TTS_MESSAGES                   BitFlag = 1 << 12
	FlagBitwisePermissionMANAGE_MESSAGES                     BitFlag = 1 << 13
	FlagBitwisePermissionEMBED_LINKS                         BitFlag = 1 << 14
	FlagBitwisePermissionATTACH_FILES                        BitFlag = 1 << 15
	FlagBitwisePermissionREAD_MESSAGE_HISTORY                BitFlag = 1 << 16
	FlagBitwisePermissionMENTION_EVERYONE                    BitFlag = 1 << 17
	FlagBitwisePermissionUSE_EXTERNAL_EMOJIS                 BitFlag = 1 << 18
	FlagBitwisePermissionVIEW_GUILD_INSIGHTS                 BitFlag = 1 << 19
	FlagBitwisePermissionCONNECT                             BitFlag = 1 << 20
	FlagBitwisePermissionSPEAK                               BitFlag = 1 << 21
	FlagBitwisePermissionMUTE_MEMBERS                        BitFlag = 1 << 22
	FlagBitwisePermissionDEAFEN_MEMBERS                      BitFlag = 1 << 23
	FlagBitwisePermissionMOVE_MEMBERS                        BitFlag = 1 << 24
	FlagBitwisePermissionUSE_VAD                             BitFlag = 1 << 25
	FlagBitwisePermissionCHANGE_NICKNAME                     BitFlag = 1 << 26
	FlagBitwisePermissionMANAGE_NICKNAMES                    BitFlag = 1 << 27
	FlagBitwisePermissionMANAGE_ROLES                        BitFlag = 1 << 28
	FlagBitwisePermissionMANAGE_WEBHOOKS                     BitFlag = 1 << 29
	FlagBitwisePermissionMANAGE_GUILD_EXPRESSIONS            BitFlag = 1 << 30
	FlagBitwisePermissionUSE_APPLICATION_COMMANDS            BitFlag = 1 << 31
	FlagBitwisePermissionREQUEST_TO_SPEAK                    BitFlag = 1 << 32
	FlagBitwisePermissionMANAGE_EVENTS                       BitFlag = 1 << 33
	FlagBitwisePermissionMANAGE_THREADS                      BitFlag = 1 << 34
	FlagBitwisePermissionCREATE_PUBLIC_THREADS               BitFlag = 1 << 35
	FlagBitwisePermissionCREATE_PRIVATE_THREADS              BitFlag = 1 << 36
	FlagBitwisePermissionUSE_EXTERNAL_STICKERS               BitFlag = 1 << 37
	FlagBitwisePermissionSEND_MESSAGES_IN_THREADS            BitFlag = 1 << 38
	FlagBitwisePermissionUSE_EMBEDDED_ACTIVITIES             BitFlag = 1 << 39
	FlagBitwisePermissionMODERATE_MEMBERS                    BitFlag = 1 << 40
	FlagBitwisePermissionVIEW_CREATOR_MONETIZATION_ANALYTICS BitFlag = 1 << 41
	FlagBitwisePermissionUSE_SOUNDBOARD                      BitFlag = 1 << 42
	FlagBitwisePermissionCREATE_GUILD_EXPRESSIONS            BitFlag = 1 << 43
	FlagBitwisePermissionCREATE_EVENTS                       BitFlag = 1 << 44
	FlagBitwisePermissionUSE_EXTERNAL_SOUNDS                 BitFlag = 1 << 45
	FlagBitwisePermissionSEND_VOICE_MESSAGES                 BitFlag = 1 << 46
	FlagBitwisePermissionSEND_POLLS                          BitFlag = 1 << 49
	FlagBitwisePermissionUSE_EXTERNAL_APPS                   BitFlag = 1 << 50
)

// Permission Overwrite Types
const (
	FlagPermissionOverwriteTypeRole   Flag = 0
	FlagPermissionOverwriteTypeMember Flag = 1
)

// Role Object
// https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID           Snowflake `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         **string  `json:"icon,omitempty"`
	UnicodeEmoji **string  `json:"unicode_emoji,omitempty"`
	Position     int       `json:"position"`
	Permissions  string    `json:"permissions"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         *RoleTags `json:"tags,omitempty"`
	Flags        BitFlag   `json:"flags"`
}

// Role Tags Structure
// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID                *Snowflake `json:"bot_id,omitempty"`
	IntegrationID        *Snowflake `json:"integration_id,omitempty"`
	PremiumSubscriber    *string    `json:"premium_subscriber,omitempty"`
	SubscriptionListedID *Snowflake `json:"subscription_listing_id,omitempty"`
	AvailableForPurchase *string    `json:"available_for_purchase,omitempty"`
	GuildConnections     *string    `json:"guild_connections,omitempty"`
}

// Role Flags
// https://discord.com/developers/docs/topics/permissions#role-object-role-flags
const (
	IN_PROMPT BitFlag = 1 << 0
)

// Team Member Role Types
// https://discord.com/developers/docs/topics/teams#team-member-roles-team-member-role-types
const (
	FlagTeamMemberRoleTypeOwner     = ""
	FlagTeamMemberRoleTypeAdmin     = "admin"
	FlagTeamMemberRoleTypeDeveloper = "developer"
	FlagTeamMemberRoleTypeReadOnly  = "read_only"
)

// Team Object
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        *string       `json:"icon"`
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
	TeamID          Snowflake `json:"team_id"`
	User            *User     `json:"user"`
	Role            string    `json:"role"`
}

// Membership State Enum
// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
const (
	FlagMembershipStateEnumINVITED  Flag = 1
	FlagMembershipStateEnumACCEPTED Flag = 2
)

// Client Status Object
// https://discord.com/developers/docs/topics/gateway-events#client-status-object
type ClientStatus struct {
	Desktop *string `json:"desktop,omitempty"`
	Mobile  *string `json:"mobile,omitempty"`
	Web     *string `json:"web,omitempty"`
}

// Activity Object
// https://discord.com/developers/docs/topics/gateway-events#activity-object
type Activity struct {
	Name          string              `json:"name"`
	Type          Flag                `json:"type"`
	URL           **string            `json:"url,omitempty"`
	CreatedAt     int                 `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID *Snowflake          `json:"application_id,omitempty"`
	Details       **string            `json:"details,omitempty"`
	State         **string            `json:"state,omitempty"`
	Emoji         **Emoji             `json:"emoji,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      *bool               `json:"instance,omitempty"`
	Flags         BitFlag             `json:"flags,omitempty"`
	Buttons       []*Button           `json:"buttons,omitempty"`
}

// Activity Types
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-types
const (
	FlagActivityTypePlaying   Flag = 0
	FlagActivityTypeStreaming Flag = 1
	FlagActivityTypeListening Flag = 2
	FlagActivityTypeWatching  Flag = 3
	FlagActivityTypeCustom    Flag = 4
	FlagActivityTypeCompeting Flag = 5
)

// Activity Timestamps Struct
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-timestamps
type ActivityTimestamps struct {
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// Activity Emoji
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-emoji
type ActivityEmoji struct {
	Name     string     `json:"name"`
	ID       *Snowflake `json:"id,omitempty"`
	Animated *bool      `json:"animated,omitempty"`
}

// Activity Party Struct
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-party
type ActivityParty struct {
	ID   *string `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

// Activity Assets Struct
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-assets
type ActivityAssets struct {
	LargeImage *string `json:"large_image,omitempty"`
	LargeText  *string `json:"large_text,omitempty"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

// Activity Asset Image
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-asset-image
type ActivityAssetImage struct {
	ApplicationAsset string `json:"application_asset_id"`
	MediaProxyImage  string `json:"image_id"`
}

// Activity Secrets Struct
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-secrets
type ActivitySecrets struct {
	Join     *string `json:"join,omitempty"`
	Spectate *string `json:"spectate,omitempty"`
	Match    *string `json:"match,omitempty"`
}

// Activity Flags
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
const (
	FlagActivityINSTANCE                    BitFlag = 1 << 0
	FlagActivityJOIN                        BitFlag = 1 << 1
	FlagActivitySPECTATE                    BitFlag = 1 << 2
	FlagActivityJOIN_REQUEST                BitFlag = 1 << 3
	FlagActivitySYNC                        BitFlag = 1 << 4
	FlagActivityPLAY                        BitFlag = 1 << 5
	FlagActivityPARTY_PRIVACY_FRIENDS       BitFlag = 1 << 6
	FlagActivityPARTY_PRIVACY_VOICE_CHANNEL BitFlag = 1 << 7
	FlagActivityEMBEDDED                    BitFlag = 1 << 8
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
	FlagOAuth2ScopeRoleConnectionsWrite                  = "role_connections.write"
	FlagOAuth2ScopeRPC                                   = "rpc"
	FlagOAuth2ScopeRPCActivitiesWrite                    = "rpc.activities.write"
	FlagOAuth2ScopeRPCNotificationsRead                  = "rpc.notifications.read"
	FlagOAuth2ScopeRPCVoiceRead                          = "rpc.voice.read"
	FlagOAuth2ScopeRPCVoiceWrite                         = "rpc.voice.write"
	FlagOAuth2ScopeVoice                                 = "voice"
	FlagOAuth2ScopeWebhookIncoming                       = "webhook.incoming"
)
