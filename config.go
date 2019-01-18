package gaea

// TODO: PLEASEEEEEEEEEEEE
// Config ...
type Config struct {
	UriConnection  string `yaml:"uri_connection"`
	PeopleDBName   string `yaml:"users_db_name"`
	AccountsDBName string `yaml:"accounts_db_name"`
	SessionsDBName string `yaml:"sessions_db_name"`
}
