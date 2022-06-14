# Dasgo
Dasgo provides a direct representation of [Discord API Objects](https://discord.com/developers/docs/reference) in **Go**. 

_Dasgo is mentioned in the [Official Discord Developer Documentation](https://discord.com/developers/docs/topics/community-resources#api-types)._

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

## Libraries

Libraries that use Dasgo.

| Library                                      | Description                                    |
| :------------------------------------------- | :--------------------------------------------- |
| [Disgo](https://github.com/switchupcb/disgo) | Discord API Wrapper, Cache, and Shard Manager. |