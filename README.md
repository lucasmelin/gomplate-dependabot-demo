# Gomplate Dependabot Demo

This is a demo showing how to use [gomplate](https://gomplate.ca) to template out a Go project.  Notably, the `go.mod` file used for templating is a completely valid `go.mod` file, which allows automated tools such as Dependabot to parse (and propose updates to) the template.

## Usage

1. [Install `gomplate`](https://docs.gomplate.ca/installing/).
2. Run `USER=farnsworth PROJECT=time-machine gomplate --input-dir=inDir --output-dir=outDir` to render the templates into a new `outDir` directory.

## Additional information

To make the demo simpler, the `_templater.tmpl` file hard-codes relative paths to the `inDir` and `outDir` directories. If you want to use this technique with the `--input-dir` and/or `--output-dir` configuration options, you'll need to adjust these hard-coded strings to read the input and output paths from a [data source](https://docs.gomplate.ca/usage/#--datasource-d) or [`context`](https://docs.gomplate.ca/usage/#--context-c) instead.
