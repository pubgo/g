package models

import (
	"github.com/pubgo/g/xerror"
	"gopkg.in/yaml.v2"
)

type PlatformLink struct {
	ID        string `json:"id"`
	PublishAt int    `json:"publish_at"`
	UpdateAt  int    `json:"update_at"`
	URL       string `json:"url"`
}

type ResLink struct {
	Slug         string                  `json:"slug"`
	Tags         []string                `json:"tags"`
	Title        string                  `json:"title"`
	Weight       string                  `json:"weight"`
	Format       string                  `json:"format"`
	Abstract     string                  `json:"abstract"`
	Basket       string                  `json:"basket"`
	Score        int                     `json:"score"`
	Content      string                  `json:"content"`
	ContentHash  string                  `json:"content_hash"`
	PushAt       int                     `json:"push_at"`
	Group        string                  `json:"group"`
	ID           string                  `json:"id"`
	MinHash      string                  `json:"min_hash"`
	PlatformLink map[string]PlatformLink `json:"platform"`
	UpdateAt     int                     `json:"update_at"`
	PublishAt    int                     `json:"publish_at"`
}

type Matter struct {
	Slug     string
	Tags     []string `yaml:",flow"`
	KeyWords []string `yaml:",flow"`
	Title    string   `yaml:",omitempty"`
	Created  string   `yaml:",omitempty"`
	Modified string   `yaml:",omitempty"`

	Published  string   `yaml:",omitempty"`
	Categories []string `yaml:",flow,omitempty"`
	Toc        bool
	IsDraft    bool
	WordCount  int    `yaml:",omitempty"`
	Abstract   string `yaml:",omitempty"`
	Format     string `yaml:",omitempty"` // 数据格式，图片，json，yaml等
	Kind       string `yaml:",omitempty"` // 输出格式 ，展现出来的格式，列表，脑图等
	Type       string `yaml:",omitempty"` // 业务格式，逻辑格式，文章评论，爱心列表
	Media      map[string]struct {
		Format string `yaml:",omitempty"`
		URL    string `yaml:",omitempty"`
	} `yaml:",omitempty"`                 // 文章内部的资源
	Weight     int    `yaml:",omitempty"`
	ID         string `yaml:",omitempty"`
	Comments   bool   `yaml:",omitempty"`
	MinHash    string `yaml:",omitempty"`
	Namespaces string `yaml:",omitempty"`
	Path       string `yaml:",omitempty"`
	Platform   map[string]struct {
		ID        string `yaml:",omitempty"`
		Published string `yaml:",omitempty"`
		Updated   string `yaml:",omitempty"`
		URL       string `yaml:",omitempty"`
	} `yaml:",omitempty"`
}

type HugoMatter struct {
	Slug                      string
	Tags                      []string `yaml:",flow"`
	KeyWords                  []string `yaml:"keywords,flow"`
	Title                     string   `yaml:",omitempty"`
	Date                      string   `yaml:",omitempty"`
	Lastmod                   string   `yaml:",omitempty"`
	Author                    string   `yaml:",omitempty"`
	Audio                     string   `yaml:",omitempty"`
	PostMetaInFooter          bool     `yaml:",omitempty"`
	ContentCopyright          bool     `yaml:",omitempty"`
	Reward                    bool     `yaml:",omitempty"`
	Mathjax                   bool     `yaml:",omitempty"`
	MathjaxEnableSingleDollar bool     `yaml:",omitempty"`
	MathjaxEnableAutoNumber   bool     `yaml:",omitempty"`
	FlowchartDiagrams         struct {
		Enable  bool   `yaml:",omitempty"`
		Options string `yaml:",omitempty"`
	} `yaml:",omitempty"`
	SequenceDiagrams struct {
		Enable  bool   `yaml:",omitempty"`
		Options string `yaml:",omitempty"`
	} `yaml:",omitempty"`
	Published  string   `yaml:",omitempty"`
	Categories []string `yaml:",flow,omitempty"`
	Images []string `yaml:",flow,omitempty"`
	Toc        bool
	Draft      bool
	Headless      bool
	WordCount  int    `yaml:",omitempty"`
	Description   string `yaml:",omitempty"`
	Format     string `yaml:",omitempty"` // 数据格式，图片，json，yaml等
	Kind       string `yaml:",omitempty"` // 输出格式 ，展现出来的格式，列表，脑图等
	Layout       string `yaml:",omitempty"` // 输出格式 ，展现出来的格式，列表，脑图等
	Type       string `yaml:",omitempty"` // 业务格式，逻辑格式，文章评论，爱心列表
	Media      map[string]struct {
		Format string `yaml:",omitempty"`
		URL    string `yaml:",omitempty"`
	} `yaml:",omitempty"`                 // 文章内部的资源
	Weight     int    `yaml:",omitempty"`
	ID         string `yaml:",omitempty"`
	Comments   bool   `yaml:",omitempty"`
	MinHash    string `yaml:",omitempty"`
	Namespaces string `yaml:",omitempty"`
	Path       string `yaml:",omitempty"`
	Platform   map[string]struct {
		ID        string `yaml:",omitempty"`
		Published string `yaml:",omitempty"`
		Updated   string `yaml:",omitempty"`
		URL       string `yaml:",omitempty"`
	} `yaml:",omitempty"`
}

func (t *HugoMatter) String() string {
	_h := "---\n"
	_h += string(xerror.PanicBytes(yaml.Marshal(t)))
	_h += "---\n\n\n"
	return _h
}
