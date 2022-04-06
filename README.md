# Dasgo
Dasgo provides 1:1 type definitions for the [Discord API](https://discord.com/developers/docs/reference) in **Go** and is a submodule of the [Discord API Spec](https://github.com/switchupcb/discord-api-spec). As a result, it aims to provide a direct representation of Discord API Objects (by version).

## Using Dasgo

There are two ways to use `dasgo`.

### Alias

#### Import

Get a specific version of the `dasgo` dependency by specifying a branch or tag.
```
go get github.com/switchupcb/dasgo@latest
```

_Dasgo branches are referenced by API version (i.e v8, v9)._

#### Definition

Use an alias to rename the type in your project.

```go
// alias dasgo.User to DiscordUser in your application or library.
type DiscordUser dasgo.User
```

### Generation

Use copy/paste or code generation to add `dasgo` types to your repository.

## Specification

For more information, read the [Specification](CONTRIBUTING.md#specification).

## Contribution

You can contribute to this repository by reading [Contribution](CONTRIBUTING.md).

### License

The [BSD 3-Clause](LICENSE) is a permissive license.
