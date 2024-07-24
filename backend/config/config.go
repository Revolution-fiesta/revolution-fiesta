package config

import (
	"strings"
)

const (
	Version        string = "0.0.1"
	Port           string = "8081"
	DatabaseUsr    string = ""
	DatabasePasswd string = ""
	DatabaseHost   string = ""
	DatabasePort   string = ""
	DatabaseName   string = ""
)

var (
	DatabaseUrl string
)

func init() {
	DatabaseUrl = getPgConnUrl(DatabaseUsr, DatabasePasswd, DatabaseHost, DatabasePort, DatabaseName)
}

func getPgConnUrl(usr, passwd, host, port, database string) string {
	builder := strings.Builder{}

	_, _ = builder.WriteString("postgres://")
	_, _ = builder.WriteString(usr)

	if passwd != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(passwd)
	}

	_, _ = builder.WriteRune('@')
	_, _ = builder.WriteString(host)

	if port != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(port)
	}

	_, _ = builder.WriteRune('/')
	_, _ = builder.WriteString(database)

	return builder.String()
}
