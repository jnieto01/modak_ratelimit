package entity

type RedisDB struct {
	Addr string `json:"addr"`
	Password string `json:"password"`
	Db int `json:"db"`
}

