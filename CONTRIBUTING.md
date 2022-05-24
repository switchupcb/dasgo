# Contribution

Dasgo is currently based on [contributor-led documentation](https://github.com/discord/discord-api-docs).

## Pull Requests

Pull requests must adhere to the following [specification](#specification).

## Specification

All types are named as defined by the source of truth; which is currently the contributor-led documentation.

### Flags

A flag is a [flag](https://discord.com/developers/docs/resources/application#application-object-application-flags), [type](https://discord.com/developers/docs/resources/channel#embed-object-embed-types), [key](https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-key), [level](https://discord.com/developers/docs/resources/guild#guild-object-verification-level) or any other option that Discord provides. All flags are denoted by `Flag` followed by their option name _(in singular form)_, then the option value. For example, `FlagUserSTAFF`, `FlagVerificationLevelHIGH`, `FlagMessageTypeDEFAULT`, etc.

_Flags that end in `Flags` should not include `Flags` (i.e `FlagMessageLOADING`)._

### Fields

Fields use data types that best represent their usage.

Snowflakes fields are defined using a `uint64` alias.  
- [Discord Snowflake](https://discord.com/developers/docs/reference#snowflakes)  
- [Twitter Snowflake](https://developer.twitter.com/en/docs/twitter-ids)  

Option fields are defined as a `Flag`, `BitFlag`, or [`CodeFlag`](https://discord.com/developers/docs/topics/opcodes-and-status-codes) to provide a separation of concerns.

### Tags and Pointers

Tags are applied to struct fields in order to marshal and unmarshal data. A `json` tag is applied to a field that pertains to a JSON key. A `url` tag is applied to a field that pertains to a URL Query String Parameter. `omitempty` and pointers are used to maintain correctness: A field with `omitempty` will **NOT** be included when it contains a zero value _(`nil` for pointers)_.

[Discord Nullable and Optional Resource Fields Documentation](https://discord.com/developers/docs/reference#nullable-and-optional-resource-fields) outlines the conditions when a field is optional or null. Optional fields are **NOT** required to be included in marshalled data _(i.e JSON/Query String)_. As a result, **optional fields should include an `omitempty` tag**. As a corollary do **NOT** include an `omitempty` tag when a field is required.

When a valid value of an optional field is equivalent to the zero value of a field, it will **NOT** be marshalled. For example, a field `int` with a value of `0` and an `omitempty` tag will **NOT** be marshalled. This is problematic because the developer will **NOT** be able to set the field to the zero value. As a result, **optional fields that CAN be equal to the Go Type's zero value should be pointers**.

Nullable fields are `nil` when marshalled. Therefore, **nullable fields should be pointers**. In contrast, an optional (`omitempty`) and nullable field (`*`) will cause conflict: For example, a field `*int` equivalent to `nil` with an `omitempty` tag will be **NOT** be included in marshalled data. This is problematic because the developer will **NOT** be able to set a field to null. As a result, **nullable and optional fields should be pointers _(without `omitempty`)_**.

As we can see, any of the above cases that require a solution to a problem places responsibility in the hands of the developer. Optional fields that **CAN** be equal to the Go Type's zero value might **NOT** be nullable, so the developer is expected to ensure that this doesn't occur. Nullable and optional fields will **ALWAYS** be included in marshalled data _(since `omitempty` is **NOT** applied)_, so the developer **must** handle these fields.

#### Examples

The optional non-nullable `type` field of an [Application Command](https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types) contains [Application Command Type](https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types) values which are **NEVER** equal to 0. In contrast, an unsigned integer — which represents the Go `FlagApplicationCommandType` — has a zero value equal to 0. Since `omitempty` is applied, an uninitialized unsigned integer _(which is equal to 0)_ will **NOT** be marshalled.

```go
type Application struct {
	// Type (optional, non-nullable) is NOT marshalled
	// when Type == 0 but IS when Type == 1.
	Type Flag `json:"type,omitempty"`
}
```

If `FlagApplicationCommandType` contains a flag that **CAN** be 0, we need to apply a pointer.

```go
type Application struct {
	// Type (optional, non-nullable) is NOT marshalled
	// when Type == nil but IS when Type == 0.
	Type *Flag `json:"type,omitempty"`
}
```

An optional non-nullable empty struct field _(i.e Author)_ should **NOT** create a struct with fields holding zero values. As a result, **structs** with `omitempty` should always be pointers.

```go
type Embed struct {
	// Author (optional, non-nullable) is NOT marshalled
	// when EmbedAuthor == nil but IS when EmbedAuthor == &{}.
	Author *EmbedAuthor `json:"author,omitempty"`
}
```

_An exception is provided to slices, channels, and maps, which — 99% of the time — should **NEVER** be pointers, but may contain pointers to types._

### Comments

Add comments to structs in the following format.
```go
// Name of Object (specified by header)
// Link to Discord API definition
type Name struct {
    // Discord Field Description.
    Example string `json:"example,omitempty"`
}
```

_Field comments are **NOT** required by the specification at the current time._

## Roadmap

Dasgo definitions are completed and in maintenance mode. However, one feature that may prove to be useful is adding comments for field-descriptions.