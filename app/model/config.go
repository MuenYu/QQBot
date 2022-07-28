package model

type Config struct {
	System   *SystemConfig `yaml:"system"`
	Function []Function    `yaml:"function"`
}

type SystemConfig struct {
	QQ      int64  `yaml:"qq"`
	OPQ     string `yaml:"opq"`
	Admin   int64  `yaml:"admin"`
	HeadMsg string `yaml:"headmsg"`
	ErrMsg  string `yaml:"errmsg"`
}

type Function struct {
	Name      string         `yaml:"name"`
	Trigger   string         `yaml:"trigger"`
	Type      string         `yaml:"type"`
	Desc      string         `yaml:"desc"`
	AdminOnly bool           `yaml:"admin_only"`
	Config    map[string]any `yaml:"boot"`
}
