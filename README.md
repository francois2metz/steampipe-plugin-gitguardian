# Gitguardian plugin for Steampipe

Use SQL to query incidents from [Gitguardian][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/gitguardian

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/gitguardian.spc ~/.steampipe/config/gitguardian.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[gitguardian]: https://www.gitguardian.com/
