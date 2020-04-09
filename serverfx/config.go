package serverfx

type ServerConfig struct {
	listenAddr string
}

func NewServerConfig() (*ServerConfig, error) {
	//return &ServerConfig{listenAddr: ":9620"}, errors.New("load config failed")
	return &ServerConfig{listenAddr: ":9620"}, nil
}
