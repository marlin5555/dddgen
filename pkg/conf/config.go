package config

// Config 配置接口
type Config interface {
	DBConfig() DBConfig
}

// DBConfig 数据库相关配置
type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Timeout  string `yaml:"timeout"`
	DBName   string `yaml:"db_name"`
}
type EmptyConfig struct {
}

func (e EmptyConfig) DBConfig() DBConfig {
	return DBConfig{}
}
