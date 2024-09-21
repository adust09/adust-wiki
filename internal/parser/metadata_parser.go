// internal/parser/metadata_parser.go
package parser

import (
	"gopkg.in/yaml.v3"
)

type ArticleMeta struct {
	Title  string   `yaml:"title"`
	Tags   []string `yaml:"tags"`
	Date   string   `yaml:"date"`
	Author string   `yaml:"author"`
}

func ParseFrontMatter(content []byte) (ArticleMeta, error) {
	var meta ArticleMeta
	err := yaml.Unmarshal(content, &meta)
	return meta, err
}
