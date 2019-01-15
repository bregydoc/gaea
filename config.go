package gaea

type Config struct {
	UriConnection  string `yaml:"uri_connection"`
	UsersDBName    string `yaml:"users_db_name"`
	SessionsDBName string `yaml:"sessions_db_name"`
}
