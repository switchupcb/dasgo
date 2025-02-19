# Contribution

Dasgo is based on [contributor-led documentation](https://github.com/discord/discord-api-docs).

## Pull Requests

Pull requests must adhere to the following [specification](#specification).

## Roadmap

Dasgo type definitions are completed and in maintenance mode.

Here are unimplemented features which enhance Dasgo.
- [Generate code based on discord-api-spec](https://github.com/switchupcb/dasgo/issues/2).
- Add type field comments.

# Specification

All types are named as defined by the source of truth; which is the contributor-led documentation.

## Flags

A flag is a [flag](https://discord.com/developers/docs/resources/application#application-object-application-flags), [type](https://discord.com/developers/docs/resources/channel#embed-object-embed-types), [key](https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-key), [level](https://discord.com/developers/docs/resources/guild#guild-object-verification-level) or any other option that Discord provides. 

All flags are denoted by `Flag` followed by their option name _(in singular form)_, then the option value, and typed with the respective Flag (`Flag`, `BitFlag`) type as defined in [`dasgo.go`](dasgo/dasgo.go). A flag that can't be represented by these definitions remains untyped.

```go
FlagVerificationLevelHIGH Flag = 3 
FlagMessageTypeDEFAULT Flag = 0
FlagUserSTAFF BitFlag = 1 << 0
```

_Flags that end in `Flags` should not include `Flags` (i.e `FlagMessageLOADING`)._

## Resources

### Comments

Add comments to resources in the following format.
```go
// Name of Object (specified by header)
// Link to Discord API definition
type Name struct {
    // Discord Field Description.
    Example string `json:"example,omitempty"`
}
```

_Field comments are **NOT** required by the specification._

### Fields

Fields use data types that best represent their usage.

Snowflakes fields are defined using a `uint64` alias.  
- [Discord Snowflake](https://discord.com/developers/docs/reference#snowflakes)  
- [Twitter Snowflake](https://developer.twitter.com/en/docs/twitter-ids)  

Option fields are defined as a `Flag`, `BitFlag`, or [`CodeFlag`](https://discord.com/developers/docs/topics/opcodes-and-status-codes) to provide a separation of concerns.

Fields that represent a slice of flags (`[]Flag`) are defined using the `Flags` type. This is required to differentiate a `[]Flag` from a `[]uint8` or `[]byte`, which marshals to an empty string (`""`) with zero length; as opposed to an empty array (`[]`).

### Tags and Pointers

**Here is a summary of the tags and pointers specification.**
- Normal fields are tagged with no options.
- Nullable fields must be defined with a pointer (`*`).
- Optional fields are defined with a pointer and tagged with `omitempty`.
  - _Exception:  `oauth2`, `ratelimit`_
- Nullable and Optional fields must be defined with a double pointer (`**`).

_Slices, channels, maps, and interfaces are pointers. Structs are always defined with pointers._

**Here is an explanation.**

Field Tags are applied to struct fields in order to marshal and unmarshal data.
- A `json` tag is applied to a field that pertains to a JSON key.
- A `url` tag is applied to a field that pertains to a URL Query String Parameter. 
- Tag options _(`-`, `omitempty`)_ and pointers are used to maintain correctness with the Discord API. 
  - A field with `omitempty` will **NOT** be marshalled when it contains a zero value _(`nil` for pointers)_.
  - A field with `-` will **NOT** be marshalled or unmarshalled: `-` is **ONLY** required when a respective tag _(e.g., `json`, `url`)_ is included in a struct.

_[Discord Nullable and Optional Resource Fields Documentation](https://discord.com/developers/docs/reference#nullable-and-optional-resource-fields) outlines the conditions when a field is optional and/or null._

#### How do you apply field tags?

**Optional fields must include an `omitempty` tag**.

What is the reason? Optional fields are **NOT** required to be included in marshalled data _(i.e JSON/Query String)_.

**Optional fields must be pointers**.

Here is the reason: A valid value _(e.g., `0`)_ of an optional field will **NOT** be marshalled when is equivalent to the zero value of a field.

Here is an example: A field `int` with a value of `0` and an `omitempty` tag will **NOT** be marshalled. 

Here is the problem with this example: The developer will **NOT** be able to set the field to the zero value.

**Nullable fields must be pointers**.

Here is the reason: `nil` pointers are null when marshalled.

**Nullable and optional fields must be double pointers**.

Here is the reason: An optional (`omitempty`) and nullable field (`*`) will cause conflict.

Here is an example: A field `*int` equivalent to `nil` with an `omitempty` tag will be **NOT** be included in marshalled data. 

Here is the problem with this example: The developer will **NOT** be able to set a field to null.

#### Field Tags and Pointer: Example

The optional non-nullable `type` field of an [Application Command](https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types) contains [Application Command Type](https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types) values which are **NEVER** equal to 0. In contrast, an unsigned integer — which represents the Go `FlagApplicationCommandType` — has a zero value equal to 0. Since `omitempty` is applied, an uninitialized unsigned integer _(which is equal to 0)_ will **NOT** be marshalled.

```go
type ApplicationCommand struct {
	// Type (optional, non-nullable) is NOT marshalled
	// when Type == 0 but IS when Type == 1.
	Type Flag `json:"type,omitempty"`
}
```

You need to apply a pointer when `FlagApplicationCommandType` contains a flag that **CAN** be 0.

```go
type ApplicationCommand struct {
	// Type (optional, non-nullable) is NOT marshalled
	// when Type == nil but IS when Type == 0.
	Type *Flag `json:"type,omitempty"`
}
```

**Structs** with `omitempty` should always be pointers.

Here is the reason: An optional non-nullable empty struct field _(i.e Author)_ should **NOT** create a struct with fields holding zero values.

```go
type Embed struct {
	// Author (optional, non-nullable) is NOT marshalled
	// when EmbedAuthor == nil but IS when EmbedAuthor == &{}.
	Author *EmbedAuthor `json:"author,omitempty"`
}
```

## Requests

Here is how to add a request to Dasgo.
1. Add the request's Endpoint Query String to the [endpoints.go](dasgo/endpoints.go) file _(in the correct position)_. 
2. Add the request's Object Representation to the [requests.go](dasgo/requests.go) file. I
3. Add the response's object representation to the [responses.go](dasgo/responses.go) file when the request involves a response, 

## Events

**Events should NOT contain `Op` fields**.

Here is the reason: Discord sends events through Gateway Payloads represented by the `GatewayPayload` struct. This object contains `Op` and `Data` fields that must be used by API Wrappers to marshal and unmarshal Gateway Payloads into Events (via `json.RawMessage`).