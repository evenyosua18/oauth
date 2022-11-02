package server

type ListServer struct {
	Debug string `yaml:"debug"`
	Grpc  Server `yaml:"grpc"`
}

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	TLS     bool   `yaml:"tls"`
	Timeout int    `yaml:"timeout"`
}
