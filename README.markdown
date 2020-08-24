# datadog-to-terraform

A small utility to export Datadog dashboards and monitors to Terraform `.tf`
resource config files.

## Examples

```console
# to export a dashboard
$ DD_API_KEY=xxx DD_APP_KEY=xxx go run . dashoard mnr-gsq-2em


# to export a monitor
$ DD_API_KEY=xxx DD_APP_KEY=xxx go run . monitor 20625761
```

## hclencoder

This project uses my own fork of `rodaine/hclencoder` with support for
serializing a slice of structs as a series of blocks, instead of assigning an
array to an argument. Hopefully I'll get this upstreamed soon :D
