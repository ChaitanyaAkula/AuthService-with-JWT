package config

type ServerConfig struct{
	Host string
	Port string
}

var serverConfig ServerConfig


func init(){
	serverConfig= ServerConfig{
		Host: getstringEmpty("AUTH_HOST"),
		Port: getstringEmpty("AUTH_PORT"),
	}
	
}

func GetServerConfig() ServerConfig{
	return serverConfig
} 