package config

type Server struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// captcha
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
