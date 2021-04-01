package connect

import (
	"strings"

	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/posener/complete"
)

const (
	dbeaverSynopsis = "Authorize a session against a target and invoke a Dbeaver client to connect to a database"
)

func dbeaverOptions(c *Command, set *base.FlagSets) {
	f := set.NewFlagSet("Dbeaver Options")

	f.StringVar(&base.StringVar{
		Name:       "style",
		Target:     &c.flagDbeaverStyle,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_STYLE",
		Completion: complete.PredictSet("dbeaver"),
		Default:    "dbeaver",
		Usage:      `Specifies how the CLI will attempt to invoke a Dbeaver client. This will also set a suitable default for -exec if a value was not specified. Currently-understood values are "dbeaver".`,
	})

	f.StringVar(&base.StringVar{
		Name:       "driver",
		Target:     &c.flagDbeaverConDriver,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_DRIVER",
		Completion: complete.PredictNothing,
		Default:    "postgres-jdbc",
		Usage:      `Specifies the driver for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "host",
		Target:     &c.flagDbeaverConHost,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_HOST",
		Completion: complete.PredictNothing,
		Usage:      `Specifies the host for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "port",
		Target:     &c.flagDbeaverConPort,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_PORT",
		Completion: complete.PredictNothing,
		Usage:      `Specifies the port of the database for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "database",
		Target:     &c.flagDbeaverConDb,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_DATABASE",
		Completion: complete.PredictNothing,
		Default:    "postgres",
		Usage:      `Specifies the database for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "name",
		Target:     &c.flagDbeaverConName,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_NAME",
		Completion: complete.PredictNothing,
		Default:    "postgres-sql",
		Usage:      `Specifies the name for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "user",
		Target:     &c.flagDbeaverConUser,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_USER",
		Completion: complete.PredictNothing,
		Usage:      `Specifies the database user for dbeaver connection.`,
	})

	f.StringVar(&base.StringVar{
		Name:       "password",
		Target:     &c.flagDbeaverConPassword,
		EnvVar:     "BOUNDARY_CONNECT_DBEAVER_PASSWORD",
		Completion: complete.PredictNothing,
		Usage:      `Specifies the database password for dbeaver connection.`,
	})

}

type dbeaverFlags struct {
	flagDbeaverStyle string
}

func (p *dbeaverFlags) defaultExec() string {
	return strings.ToLower(p.flagDbeaverStyle)
}

func (p *dbeaverFlags) buildArgs(c *Command, port, ip, addr string) []string {
	var args []string
	switch p.flagDbeaverStyle {
	case "dbeaver":
		{
			args = append(args, "-con")

			var subArgs []string
			subArgs = append(subArgs, "\"")

			subArgs = append(subArgs, "driver=", c.flagDbeaverConDriver)
			subArgs = append(subArgs, "|host=", c.flagDbeaverConHost)
			subArgs = append(subArgs, "|port=", c.flagDbeaverConPort)
			subArgs = append(subArgs, "|name=", c.flagDbeaverConName)
			subArgs = append(subArgs, "|database=", c.flagDbeaverConDb)
			subArgs = append(subArgs, "|user=", c.flagDbeaverConUser)

			if c.flagImpersonate {
				subArgs = append(subArgs, "|password=", c.impersonateCredentials)
			} else {
				subArgs = append(subArgs, "|password=", c.flagDbeaverConPassword)
			}

			subArgs = append(subArgs, "\"")

			args = append(args, strings.Join(subArgs, ""))
		}
	}
	return args
}
