# Contribution

Dasgo is currently based on [contributor-led documentation](https://github.com/discord/discord-api-docs).

## Pull Requests

Pull requests must adhere to the following [specification](#specification).

### Specification

All types are named as defined by the source of truth; which is currently the [contributor-led documentation](https://github.com/discord/discord-api-docs).

#### Fields

Fields use data types that best represent their usage.

Snowflakes fields are defined using a `uint64` alias.  
- [Discord Snowflake](https://discord.com/developers/docs/reference#snowflakes)  
- [Twitter Snowflake](https://developer.twitter.com/en/docs/twitter-ids)  

Option fields are defined as a `Flag`, `BitFlag`, or [`CodeFlag`](https://discord.com/developers/docs/topics/opcodes-and-status-codes) to provide a separation of concerns between usage. A flag is a [flag](https://discord.com/developers/docs/resources/application#application-object-application-flags), [type](https://discord.com/developers/docs/resources/channel#embed-object-embed-types), [key](https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-key), [level](https://discord.com/developers/docs/resources/guild#guild-object-verification-level) or any other option that Discord provides. All flags are denoted by `Flag` followed by their option name, then the option value; For example, `FlagUserSTAFF`, `FlagVerificationLevelHIGH`, `FlagPremiumTierNONE`, etc.

#### Tags and Pointers

A JSON `omitempty` tag should **always** be applied to struct fields and maintain correctness with respect to its pointer. Apply a pointer to a field when the _Field's Discord Representation_ can **NOT** be equal to the Go Type's zero-value or required as `null`.

For example, [Application Command Types](https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types) are never equal to 0. In contrast, an unsigned integer — which represents the Go `FlagApplicationCommandType` — has a zero-value equal to 0. This means an uninitialized unsigned integer _(which is equal to 0)_ can be considered as a `nil` value for the `FlagTypeCommandApplication`. As a result, fields that reference `FlagTypeCommandApplication` should **NOT** be a pointer.

```go
type Application struct {
	Flags uint `json:"flags,omitempty"`
}
```

If this is not the case _(i.e [Verification Level](https://discord.com/developers/docs/resources/guild#guild-object-verification-level)),_ fields that reference the type **SHOULD** be a pointer.
 
```go
type Guild struct {
	VerificationLevel *uint `json:"verification_level,omitempty"`
}
```

An empty struct response should **NOT** create a struct with default values. As a result, **structs** with `omitempty` must always be pointers.

```go
type Embed struct {
	Author *EmbedAuthor `json:"author,omitempty"`
}
```

An exception is provided to slices, channels, and maps, which should **NEVER** be pointers, but may contain pointers to types.

#### Comments

Add comments to types and fields in the following format.
```go
// Name of Object (specified by header)
// Link to Discord API definition
type Name struct {
    // Description for field (provided by Discord).
    Example string `json:"example,omitempty"`
}
```

The following examples are provided.

##### Resources

```go
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
```

##### Requests

```go
// Edit Global Application Command
// PATCH/applications/{application.id}/commands/{command.id}
// https://discord.com/developers/docs/interactions/application-commands#edit-global-application-command
type EditGlobalApplicationCommand struct {
	CommandID                Snowflake
	Name                     string                      `json:"name,omitempty"`
	NameLocalizations        map[Flag]string             `json:"name_localizations,omitempty"`
	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Flag]string             `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission        bool                        `json:"default_permission,omitempty"`
}
```

##### Events

```go
// Thread Members Update
// https://discord.com/developers/docs/topics/gateway#thread-members-update
type ThreadMembersUpdate struct {
	ID             Snowflake       `json:"id,omitempty"`
	GuildID        Snowflake       `json:"guild_id,omitempty"`
	MemberCount    int             `json:"member_count,omitempty"`
	AddedMembers   []*ThreadMember `json:"added_members,omitempty"`
	RemovedMembers []Snowflake     `json:"removed_member_ids,omitempty"`
}
```

## Roadmap

Dasgo definitions are completed and in maintenance mode. However, one feature that may prove to be useful is adding comments for field-descriptions.