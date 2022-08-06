package config

type Captcha struct {
	KeyLong int   `json:"key-long" yaml:"key-long" mapstructure:"key-long"`
	TimeOut int64 `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
}
