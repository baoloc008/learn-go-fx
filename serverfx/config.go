package serverfx

type ServerConfig struct {
	listenAddr string
}

func newServerConfig() (*ServerConfig, error) {
	//return &ServerConfig{listenAddr: ":9620"}, errors.New("load config failed")
	return &ServerConfig{listenAddr: ":9620"}, nil
}
