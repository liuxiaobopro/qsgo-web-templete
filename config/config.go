package config

type Conf struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type Mysql struct {
	Dns string `mapstructure:"dns" json:"dns" yaml:"dns"`
}
