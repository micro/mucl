package generator

import (
	"fmt"

	"github.com/micro/mu/mucl"
)

// PluginMap is a map of plugin names to their respective plugin import options
var PluginMap = map[string]map[string]string{
	"broker": {
		"nats":     "github.com/micro/plugins/v5/broker/nats",
		"redis":    "github.com/micro/plugins/v5/broker/redis",
		"nsq":      "github.com/micro/mu-broker-nsq/v2",
		"rabbitmq": "github.com/micro/mu-broker-rabbitmq/v2",
		"mqtt":     "github.com/micro/mu-broker-mqtt/v2",
	},
	"registry": {
		"etcd":   "github.com/micro/plugins/v5/registry/etcd",
		"consul": "github.com/micro/plugins/v5/registry/consul",
	},
	"transport": {
		"nats":  "github.com/micro/mu-transport-nats/v2",
		"redis": "github.com/micro/mu-transport-redis/v2",
		"etcd":  "github.com/micro/mu-transport-etcd/v2",
	},
	"protocol": {
		"grpc":      "github.com/micro/mu-protocol-grpc/v2",
		"http":      "github.com/micro/mu-protocol-http/v2",
		"websocket": "github.com/micro/mu-protocol-websocket/v2",
	},
	"server": {
		"grpc":      "github.com/micro/mu-server-grpc/v2",
		"http":      "github.com/micro/mu-server-http/v2",
		"websocket": "github.com/micro/mu-server-websocket/v2",
	},
}

func GetPluginList(s *mucl.Service) []string {
	fmt.Println("GetPluginList")
	fmt.Println("broker", s.Broker())
	fmt.Println("registry", s.Registry())
	fmt.Println("transport", s.Transport())
	fmt.Println("protocol", s.Protocol())
	fmt.Println("server", s.Server())

	plugins := make([]string, 0)
	if s.Broker() != "" {
		key := s.Broker()
		fmt.Println("broker key", key)
		if val, ok := PluginMap["broker"][key]; ok {
			plugins = append(plugins, val)
		}
	}
	if s.Registry() != "" {
		key := s.Registry()
		fmt.Println("registry key", key)

		if val, ok := PluginMap["registry"][key]; ok {
			plugins = append(plugins, val)
		}
	}
	if s.Transport() != "" {
		key := s.Transport()
		fmt.Println("transport key", key)

		if val, ok := PluginMap["transport"][key]; ok {
			plugins = append(plugins, val)
		}
	}
	if s.Protocol() != "" {
		key := s.Protocol()
		fmt.Println("protocol key", key)

		if val, ok := PluginMap["protocol"][key]; ok {
			plugins = append(plugins, val)
		}
	}
	if s.Server() != "" {
		key := s.Server()
		fmt.Println("server key", key)
		if val, ok := PluginMap["server"][key]; ok {
			plugins = append(plugins, val)
		}
	}
	fmt.Println("plugins", plugins)
	return plugins
}
