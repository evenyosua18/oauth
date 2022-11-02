package database

type ListDatabase struct {
	Oauth Database `yaml:"oauth"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Adapter  string `yaml:"adapter"`
}
