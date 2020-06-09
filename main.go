package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"totp-proxy/proxy"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	defaultInterval = 30 * 60 // 30 minutes
	defaultSecret   = "NB2HI4B2F4XXM2LOMNSW45DDOVUS4Y3O"
)

func main() {

	var secret string
	var interval int64

	app := &cli.App{
		Name:      "totp-proxy",
		Usage:     "using time-based one-time password protect your website",
		ArgsUsage: "upstream",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "secret",
				Aliases:     []string{"s"},
				Value:       defaultSecret,
				Usage:       "base32 encoding totp secret",
				EnvVars:     []string{"SECRET"},
				Destination: &secret,
			},
			&cli.Int64Flag{
				Name:        "interval",
				Aliases:     []string{"i"},
				Value:       defaultInterval,
				Usage:       "totp token refresh interval",
				EnvVars:     []string{"INTERVAL"},
				Destination: &interval,
			},
		},
		Action: func(c *cli.Context) error {
			up := c.Args().Get(0)
			if up == "" {
				return errors.New("upstream is expected")
			}

			upstream, err := url.Parse(up)
			if err != nil {
				return errors.Wrap(err, "parse upstream")
			}

			if secret == "" {
				secret = defaultSecret
				log.Println("using default secret key:", secret)
			} else {
				log.Println("using secret key:", secret)
				log.Println("protect your secret key properly")
			}

			if interval == 0 {
				log.Println("using default interval:", defaultInterval)
			} else {
				log.Println("using interval:", interval)
			}

			server := proxy.NewServer(upstream, secret, interval)

			http.HandleFunc("/", server.ServeHTTP)
			log.Fatal(http.ListenAndServe(":8080", nil))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
