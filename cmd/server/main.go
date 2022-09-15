package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

const version string = "v0.1"

func main() {
	fmt.Printf("Burr %s\n", version)
	fmt.Printf("Copyright © %d Matt Hamilton\n", time.Now().Year())
	fmt.Printf("Licensed under the GNU Affero Public License 3.0\n")

	//
	// CLI
	//

	var cfgFile string
	root := &cobra.Command{
		Use: "burr",
		RunE: func(_ *cobra.Command, _ []string) error {
			cfg, err := ConfigFromFile(cfgFile)
			if err != nil {
				return err
			}
			log.Printf("Launching all services for %s...", cfg.Domain)
			cmds := map[string]func(*Config) error{
				"web": web,
			}
			errCh := make(chan error)
			for k, v := range cmds {
				log.Println("Launching " + k)
				go func() {
					// catch errors and nils, because nil means a service exited
					errCh <- v(cfg)
				}()
			}
			if err := <-errCh; err != nil {
				// NOTE: some services may need additional signaling for cleanup stuffs in the future
				return err
			}
			return nil
		},
		// SilenceUsage: true,
	}
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "YAML configuration file for burr")

	// burr web
	root.AddCommand(webCmd)
	// burr worker -q push
	root.AddCommand(workerCmd)

	//
	// Run CLI
	//

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
