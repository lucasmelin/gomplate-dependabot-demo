package main

import "github.com/hashicorp/go-hclog"

func main() {
	hclog.Default().Info("{{ .Env.PROJECT }} was created by {{ .Env.USER }}")
}
