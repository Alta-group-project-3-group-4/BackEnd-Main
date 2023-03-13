package config

type AppConfig struct {
	DBUSERNAME        string
	DBPASS            string
	DBHOST            string
	DBPORT            string
	DBNAME            string
	JWTKEY            string
	keyid             string
	accesskey         string
	midtransserverkey string
}

func InitConfig() *AppConfig {
	return InitEnv()
}
