package server

type ConfigServer struct {
	Debug       string `yaml:"debug"`
	ServiceName string `yaml:"serviceName"`
	Grpc        Server `yaml:"grpc"`
	Tracer      Tracer `yaml:"tracer"`
	Token       Token  `yaml:"token"`
}

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	TLS     bool   `yaml:"tls"`
	Timeout int    `yaml:"timeout"`
	MaxIdle string `yaml:"maxIdle"`
	MaxAge  string `yaml:"maxAge"`
}

type Tracer struct {
	Endpoint string `yaml:"endpoint"`
	Env      string `yaml:"env"`
}

type Token struct {
	Expired       string `yaml:"expired"`
	LengthExpired int    `yaml:"lengthExpiredToken"`
}
