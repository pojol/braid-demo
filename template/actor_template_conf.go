package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

//go:generate go run actor_template_gen.go

type ActorTemplateConfig struct {
	Name string `yaml:"name"`
}

type NodeConfig struct {
	ID     string                  `yaml:"id"`
	Weight string                  `yaml:"weight"`
	Port   string                  `yaml:"port"`
	Actors []RegisteredActorConfig `yaml:"actors"`
}

type Config struct {
	Node NodeConfig `yaml:"node"`
}

type RegisteredActorConfig struct {
	ID      string            `yaml:"id"`
	Name    string            `yaml:"name"`
	Weight  int               `yaml:"weight"`
	Unique  bool              `yaml:"unique"`
	Limit   int               `yaml:"limit"`
	Options map[string]string `yaml:"options,omitempty"`
}

type ActorTypes struct {
	ActorTypes []ActorTemplateConfig `yaml:"actor_templates"`
}

func ParseConfig(confPath, actorTypesPath string) (*NodeConfig, error) {
	// 读取配置文件
	configData, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read main config: %v", err)
	}

	// 读取 actor 类型文件
	actorTypesData, err := ioutil.ReadFile(actorTypesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read actor types: %v", err)
	}

	return ParseConfigFromString(string(configData), string(actorTypesData))
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func replaceEnvVars(content string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		key := match[1 : len(match)-1]
		return getEnvOrDefault(key, match)
	})
}

func ParseConfigFromString(confData, actorTypesData string) (*NodeConfig, error) {

	// 解析 actor 类型
	var actorTypes ActorTypes
	if err := yaml.Unmarshal([]byte(actorTypesData), &actorTypes); err != nil {
		return nil, fmt.Errorf("failed to parse actor types: %v", err)
	}

	// 创建 actor 类型映射，用于快速查找
	actorTypeMap := make(map[string]ActorTemplateConfig)
	for _, actorType := range actorTypes.ActorTypes {
		actorTypeMap[actorType.Name] = actorType
	}

	// 解析主配置
	var config Config

	confData = replaceEnvVars(confData)

	if err := yaml.Unmarshal([]byte(confData), &config); err != nil {
		return nil, fmt.Errorf("failed to parse main config: %v", err)
	}

	for _, registeredActor := range config.Node.Actors {
		_, ok := actorTypeMap[registeredActor.Name]
		if !ok {
			return nil, fmt.Errorf("actor %s is registered in node.yml but not defined in actor_template.yml", registeredActor.Name)
		}
	}

	nodeConfig := &NodeConfig{
		ID:     config.Node.ID,
		Weight: config.Node.Weight,
		Port:   config.Node.Port,
		Actors: config.Node.Actors,
	}

	return nodeConfig, nil
}
