package config

type Email struct {
	To       string `json:"to" mapstructure:"to" yaml:"to"`
	Port     int    `json:"port" mapstructure:"port" yaml:"port"`
	From     string `json:"from" mapstructure:"from" yaml:"from"`
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	IsSSL    bool   `json:"is-ssl" mapstructure:"is-ssl" yaml:"is-ssl"`
	Secret   string `json:"secret" mapstructure:"secret" yaml:"secret"`
	Nickname string `json:"nickname" mapstructure:"nickname" json:"nickname"`
}
