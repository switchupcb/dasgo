# Dasgo

[![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge&logo=appveyor&logo=appveyor)](https://pkg.go.dev/github.com/switchupcb/dasgo)
[![License](https://img.shields.io/github/license/switchupcb/dasgo.svg?style=for-the-badge)](https://github.com/switchupcb/dasgo/blob/main/LICENSE)

Dasgo provides a direct representation of [Discord API Objects](https://discord.com/developers/docs/reference) in **Go**. 

**Dasgo is mentioned in the [Official Discord Developer Documentation](https://discord.com/developers/docs/topics/community-resources#api-types).**

## Using Dasgo

There are two ways to use `dasgo`.

### Alias

#### Import

Get a specific version of the `dasgo` library by specifying a tag or branch.

```
go get github.com/switchupcb/dasgo@v10
```

_Dasgo branches are referenced by API version (i.e `v10`)._

#### Definition

Use an alias to rename the type in your project.

```go
// alias dasgo.User to DiscordUser in your application or library.
type DiscordUser dasgo.User
```

### Generation

Use copy-paste or code generation to add `dasgo` types to your repository.

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
| [Disgo](https://github.com/switchupcb/disgo) | Discord API Wrapper, Shard Manager, and Cache. |