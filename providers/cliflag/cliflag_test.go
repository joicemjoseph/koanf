package cliflag

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestCliFlag(t *testing.T) {
	cliApp := cli.App{
		Name: "testing",
		Action: func(ctx *cli.Context) error {
			p := Provider(ctx, ".")
			x, err := p.Read()
			require.NoError(t, err)
			require.NotEmpty(t, x)

			fmt.Printf("x: %v\n", x)

			return nil
		},
		Flags: []cli.Flag{
			cli.HelpFlag,
			cli.VersionFlag,
			&cli.StringFlag{
				Name:    "test",
				Usage:   "test flag",
				Value:   "test",
				Aliases: []string{"t"},
				EnvVars: []string{"TEST_FLAG"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "x",
				Description: "yeah yeah testing",
				Action: func(ctx *cli.Context) error {
					p := Provider(ctx, "")
					x, err := p.Read()
					require.NoError(t, err)
					require.NotEmpty(t, x)
					return nil
				},
				Flags: []cli.Flag{
					cli.HelpFlag,
					cli.VersionFlag,
					&cli.StringFlag{
						Name:     "lol",
						Usage:    "test flag",
						Value:    "test",
						Required: true,
						EnvVars:  []string{"TEST_FLAG"},
					},
				},
			},
		},
	}

	x := append([]string{"testing", "--test", "gf", "x", "--lol", "dsf"}, os.Args...)
	err := cliApp.Run(append(x, os.Environ()...))
	require.NoError(t, err)
}
