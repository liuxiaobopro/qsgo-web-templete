package config

type Conf struct {
	Runmode string `mapstructure:"runmode" json:"runmode" yaml:"runmode"`

	Mysql struct {
		Dns    string `mapstructure:"dns" json:"dns" yaml:"dns"`
		Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	} `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	Redis struct {
		Dns      string `mapstructure:"dns" json:"dns" yaml:"dns"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	} `mapstructure:"redis" json:"redis" yaml:"redis"`

	Http struct {
		Port int `mapstructure:"port" json:"port" yaml:"port"`
	} `mapstructure:"http" json:"http" yaml:"http"`
}
