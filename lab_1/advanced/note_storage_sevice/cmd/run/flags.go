package run

import cli "github.com/urfave/cli/v2"

var cmdFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "db-host",
		Usage:   "db host",
		EnvVars: []string{"DB_HOST"},
		Value:   "db:5432",
	},
	&cli.StringFlag{
		Name:    "db-user",
		Usage:   "db user",
		EnvVars: []string{"DB_USER"},
		Value:   "postgres",
	},
	&cli.StringFlag{
		Name:    "db-password",
		Usage:   "db password",
		EnvVars: []string{"DB_PASSWORD"},
		Value:   "password",
	},
	&cli.StringFlag{
		Name:    "db-dbname",
		Usage:   "db dbname",
		EnvVars: []string{"DB_DBNAME"},
		Value:   "postgres",
	},
}
