package config

type Server struct {
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
