# Gomplate Dependabot Demo

This is a demo showing how to use [gomplate](https://gomplate.ca) to template out a Go project.  Notably, the `go.mod` file used for templating is a completely valid `go.mod` file, which allows automated tools such as Dependabot to parse (and propose updates to) the template.

## Usage

1. [Install `gomplate`](https://docs.gomplate.ca/installing/).
2. Run `USER=farnsworth PROJECT=time-machine gomplate --input-dir=inDir --output-dir=outDir` to render the templates into a new `outDir` directory.

## Additional information

To make the demo simpler, the `_templater.tmpl` file hard-codes a relative path to the `outDir` directory. If you want to use this technique with the `--output-dir` configuration option, you'll need to adjust the hard-coded string to read the output path from a [data source](https://docs.gomplate.ca/usage/#--datasource-d) or [`context`](https://docs.gomplate.ca/usage/#--context-c) instead to ensure that the file is written to the correct directroy. See https://github.com/hairyhenderson/gomplate/issues/2057 for more information.
