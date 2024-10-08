//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type ActorType struct {
	Name    string            `yaml:"name"`
	ID      string            `yaml:"id"`
	Unique  bool              `yaml:"unique"`
	Weight  int               `yaml:"weight"`
	Limit   int               `yaml:"limit"`
	Options map[string]string `yaml:"options,omitempty"`
}

type Config struct {
	ActorTypes []ActorType `yaml:"actor_types"`
}

func convertComment(comment string) string {
	lines := strings.Split(comment, "\n")
	var converted []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			converted = append(converted, "// "+strings.TrimPrefix(trimmed, "#"))
		}
	}
	return strings.Join(converted, "\n")
}

func main() {
	// 读取 YAML 文件
	yamlFile, err := ioutil.ReadFile("actor_types.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 提取顶级注释
	lines := strings.Split(string(yamlFile), "\n")
	var topLevelComment strings.Builder
	var actorTypesContent strings.Builder
	inTopLevelComment := true

	for _, line := range lines {
		if strings.TrimSpace(line) == "actor_types:" {
			inTopLevelComment = false
			actorTypesContent.WriteString(line + "\n")
			continue
		}

		if inTopLevelComment {
			topLevelComment.WriteString(line + "\n")
		} else {
			actorTypesContent.WriteString(line + "\n")
		}
	}

	// 解析 YAML
	var config Config
	err = yaml.Unmarshal([]byte(actorTypesContent.String()), &config)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// 提取各个 actor 的注释
	comments := make(map[string]string)
	var currentComment strings.Builder

	for i, line := range strings.Split(actorTypesContent.String(), "\n") {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "#") {
			currentComment.WriteString(line + "\n")
		} else if strings.HasPrefix(trimmedLine, "- name:") {
			actorName := strings.Trim(strings.TrimPrefix(trimmedLine, "- name:"), "\" ")
			comments[actorName] = currentComment.String()
			currentComment.Reset()
		} else if trimmedLine == "" && i < len(lines)-1 && strings.HasPrefix(strings.TrimSpace(lines[i+1]), "- name:") {
			currentComment.Reset()
		}
	}

	// 准备 Go 代码
	var code strings.Builder
	code.WriteString(`// Code generated by go generate; DO NOT EDIT.

package config

`)

	// 添加顶级注释
	code.WriteString(convertComment(topLevelComment.String()))
	code.WriteString("\nconst (\n")

	// 生成常量
	for _, actor := range config.ActorTypes {
		// 添加转换后的注释
		if comment, ok := comments[actor.Name]; ok && comment != "" {
			code.WriteString(convertComment(comment))
			code.WriteString("\n")
		}
		constName := "ACTOR_" + strings.ToUpper(actor.Name)
		code.WriteString(fmt.Sprintf("    %s = \"%s\"\n\n", constName, actor.Name))
	}

	code.WriteString(")\n")

	// 写入文件
	err = ioutil.WriteFile("actor_types.go", []byte(code.String()), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Println("actor_types.go has been generated successfully.")
}
