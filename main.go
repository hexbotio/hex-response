package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/hexbotio/hex-plugin"
)

type HexResponse struct {
}

func (g *HexResponse) Perform(args hexplugin.Arguments) (resp hexplugin.Response) {
	resp = hexplugin.Response{
		Output:  args.Command,
		Success: true,
	}
	return resp
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"action": &hexplugin.HexPlugin{Impl: &HexResponse{}},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: hexplugin.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}
