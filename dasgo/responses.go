package dasgo

import "time"

// List Public Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-public-archived-threads-response-body
type ListPublicArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// List Private Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-private-archived-threads-response-body
type ListPrivateArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// List Joined Private Archived Threads Response Body
// https://discord.com/developers/docs/resources/channel#list-joined-private-archived-threads-response-body
type ListJoinedPrivateArchivedThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
	HasMore bool            `json:"has_more"`
}

// List Active Guild Threads Response Body
// https://discord.com/developers/docs/resources/guild#list-active-guild-threads-response-body
type ListActiveGuildThreadsResponse struct {
	Threads []*Channel      `json:"threads"`
	Members []*ThreadMember `json:"members"`
}

// Get Guild Prune Count Response Body
// https://discord.com/developers/docs/resources/guild#get-guild-prune-count
type GetGuildPruneCountResponse struct {
	Pruned int `json:"pruned"`
}

// Modify Guild MFA Level Response
// https://discord.com/developers/docs/resources/guild#modify-guild-mfa-level
type ModifyGuildMFALevelResponse struct {
	Level Flag `json:"level"`
}

// List Nitro Sticker Packs Response
// https://discord.com/developers/docs/resources/sticker#list-nitro-sticker-packs
type ListNitroStickerPacksResponse struct {
	StickerPacks []*StickerPack `json:"sticker_packs"`
}

// Current Authorization Information Response Structure
// https://discord.com/developers/docs/topics/oauth2#get-current-authorization-information
type CurrentAuthorizationInformationResponse struct {
	Application *Application `json:"application"`
	Scopes      []int        `json:"scopes"`
	Expires     *time.Time   `json:"expires"`
	User        *User        `json:"user,omitempty"`
}

// Get Gateway Response
// https://discord.com/developers/docs/topics/gateway#get-gateway-example-response
type GetGatewayResponse struct {
	URL string `json:"url,omitempty"`
}

// Get Gateway Bot Response
// https://discord.com/developers/docs/topics/gateway#get-gateway-example-response
type GetGatewayBotResponse struct {
	URL               string            `json:"url"`
	Shards            *int              `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

// Redirect URL
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-redirect-url-example
type RedirectURL struct {
	Code  string `url:"code,omitempty"`
	State string `url:"state,omitempty"`

	// https://discord.com/developers/docs/topics/oauth2#advanced-bot-authorization
	GuildID     Snowflake `url:"guild_id,omitempty"`
	Permissions BitFlag   `url:"permissions,omitempty"`
}

// Access Token Response
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-access-token-response
type AccessTokenResponse struct {
	AccessToken  string        `json:"access_token,omitempty"`
	TokenType    string        `json:"token_type,omitempty"`
	ExpiresIn    time.Duration `json:"expires_in,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	Scope        string        `json:"scope,omitempty"`
}

// Redirect URI
// https://discord.com/developers/docs/topics/oauth2#implicit-grant-redirect-url-example
type RedirectURI struct {
	AccessToken string        `url:"access_token,omitempty"`
	TokenType   string        `url:"token_type,omitempty"`
	ExpiresIn   time.Duration `url:"expires_in,omitempty"`
	Scope       string        `url:"scope,omitempty"`
	State       string        `url:"state,omitempty"`
}

// Client Credentials Access Token Response
// https://discord.com/developers/docs/topics/oauth2#client-credentials-grant-client-credentials-access-token-response
type ClientCredentialsAccessTokenResponse struct {
	AccessToken string        `json:"access_token,omitempty"`
	TokenType   string        `json:"token_type,omitempty"`
	ExpiresIn   time.Duration `json:"expires_in,omitempty"`
	Scope       string        `json:"scope,omitempty"`
}

// Webhook Token Response
// https://discord.com/developers/docs/topics/oauth2#webhooks-webhook-token-response-example
type WebhookTokenResponse struct {
	TokenType    string        `json:"token_type,omitempty"`
	AccessToken  string        `json:"access_token,omitempty"`
	Scope        string        `json:"scope,omitempty"`
	ExpiresIn    time.Duration `json:"expires_in,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	Webhook      *Webhook      `json:"webhook,omitempty"`
}

// Extended Bot Authorization Access Token Response
// https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-access-token-response
type ExtendedBotAuthorizationAccessTokenResponse struct {
	TokenType    string        `json:"token_type,omitempty"`
	Guild        *Guild        `json:"guild,omitempty"`
	AccessToken  string        `json:"access_token,omitempty"`
	Scope        string        `json:"scope,omitempty"`
	ExpiresIn    time.Duration `json:"expires_in,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
}
