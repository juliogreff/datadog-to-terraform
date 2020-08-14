# datadog-to-terraform

A small utility to export Datadog dashboards to Terraform `.tf` resource config
file.

## Example

```console
$ DD_API_KEY=xxx DD_APP_KEY=xxx go run . mnr-gsq-2em
```

## hclencoder

This project uses my own fork of `rodaine/hclencoder` with support for
serializing a slice of structs as a series of blocks, instead of assigning an
array to an argument. Hopefully I'll get this upstreamed soon :D
