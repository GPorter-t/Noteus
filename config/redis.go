package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`           // Redis DB
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`     // Redis 服务器地址
	Password string `mapstructure:"password" json:"password" yaml:""` // Redis 密码 == Auth
}
