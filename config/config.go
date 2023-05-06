package config

type Conf struct {
	Runmode string `mapstructure:"runmode" json:"runmode" yaml:"runmode"`

	Mysql struct {
		Dns    string `mapstructure:"dns" json:"dns" yaml:"dns"`
		Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	} `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	Http struct {
		Port int `mapstructure:"port" json:"port" yaml:"port"`
	} `mapstructure:"http" json:"http" yaml:"http"`
}
