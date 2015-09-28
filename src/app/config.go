package app

/*
	Application Configuration structure
*/

type Config struct {
	Env struct {
		DevMode bool
	}
	Server struct {
		Port        int
		Templates   string
		SessionKey  string
		AssetsProxy string
	}
	Db struct {
		CreateDB         bool
		ConnectionString string
		Dialect          string
	}
}
