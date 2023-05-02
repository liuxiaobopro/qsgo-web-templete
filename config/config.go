package config

type Conf struct {
	Mysql struct {
		Dns string `mapstructure:"dns" json:"dns" yaml:"dns"`
	} `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	Http struct {
		Port int `mapstructure:"port" json:"port" yaml:"port"`
	} `mapstructure:"http" json:"http" yaml:"http"`
}
