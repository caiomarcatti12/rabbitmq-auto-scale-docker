package config

type Auth struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	VHost    string `yaml:"vhost"`
}

type Queue struct {
	Name          string `yaml:"name"`
	ContainerName string `yaml:"containerName"`
}

type Config struct {
	Auth   Auth    `yaml:"auth"`
	Queues []Queue `yaml:"queues"`
}
