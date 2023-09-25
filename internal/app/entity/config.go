package entity

type Config struct {
    Server   ServerConfig
	Kvs RedisDB
	RateLimitRules []RateLimitRule 
}

type ServerConfig struct {
    Host string `json:"host"`
    Port string    `json:"port"`
	GoEnv string `json:"goenv"`
}
