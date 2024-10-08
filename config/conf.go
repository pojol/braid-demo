package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//go:generate go run actor_types_gen.go

type ActorConfig struct {
	Name    string            `yaml:"name"`
	Unique  bool              `yaml:"unique"`
	Weight  int               `yaml:"weight"`
	Limit   int               `yaml:"limit"`
	Options map[string]string `yaml:"options,omitempty"`
}

type NodeConfig struct {
	ID     string                  `yaml:"id"`
	Weight string                  `yaml:"weight"`
	Actors []RegisteredActorConfig `yaml:"actors"`
}

type Config struct {
	Node NodeConfig `yaml:"node"`
}

type RegisteredActorConfig struct {
	Name    string            `yaml:"name"`
	Options map[string]string `yaml:"options,omitempty"`
}

type ActorTypes struct {
	ActorTypes []ActorConfig `yaml:"actor_types"`
}

func loadYAML(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

func ParseConfig(confPath, actorTypesPath string) (*NodeConfig, []ActorConfig, error) {
	// 读取配置文件
	configData, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read main config: %v", err)
	}

	// 读取 actor 类型文件
	actorTypesData, err := ioutil.ReadFile(actorTypesPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read actor types: %v", err)
	}

	return ParseConfigFromString(string(configData), string(actorTypesData))
}

func ParseConfigFromString(confData, actorTypesData string) (*NodeConfig, []ActorConfig, error) {

	// 解析 actor 类型
	var actorTypes ActorTypes
	if err := yaml.Unmarshal([]byte(actorTypesData), &actorTypes); err != nil {
		return nil, nil, fmt.Errorf("failed to parse actor types: %v", err)
	}

	// 创建 actor 类型映射，用于快速查找
	actorTypeMap := make(map[string]ActorConfig)
	for _, actorType := range actorTypes.ActorTypes {
		actorTypeMap[actorType.Name] = actorType
	}

	// 解析主配置
	var config Config
	if err := yaml.Unmarshal([]byte(confData), &config); err != nil {
		return nil, nil, fmt.Errorf("failed to parse main config: %v", err)
	}

	// 解析节点配置
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID = config.Node.ID
	}
	nodeWeight := os.Getenv("NODE_WEIGHT")
	if nodeWeight == "" {
		nodeWeight = config.Node.Weight
	}

	var parsedActors []ActorConfig
	for _, registeredActor := range config.Node.Actors {
		actorType, ok := actorTypeMap[registeredActor.Name]
		if !ok {
			return nil, nil, fmt.Errorf("actor %s is registered in conf.yml but not defined in actor_types.yml", registeredActor.Name)
		}

		actor := ActorConfig{
			Name:    actorType.Name,
			Unique:  actorType.Unique,
			Weight:  actorType.Weight,
			Limit:   actorType.Limit,
			Options: registeredActor.Options, // 使用 conf.yml 中的 options
		}
		parsedActors = append(parsedActors, actor)
	}

	nodeConfig := &NodeConfig{
		ID:     nodeID,
		Weight: nodeWeight,
		Actors: config.Node.Actors,
	}

	return nodeConfig, actorTypes.ActorTypes, nil
}
