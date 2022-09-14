package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	Title  string
	Domain string
}

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
		Run: func(_ *cobra.Command, _ []string) {
			// All-in-One mode
			cmds := map[string]func() error{
				"web": web,
			}
			errCh := make(chan error)
			for k, v := range cmds {
				log.Println("Launching component: " + k)
				go func() {
					// catch errors and nils, because nil means a service exited
					errCh <- v()
				}()
			}
			if err := <-errCh; err != nil {
				// NOTE: some services may need additional signaling for cleanup stuffs in the future
				log.Fatal(err)
			}
			os.Exit(0) // should never hit this exit unless somehow all services close without error
		},
		// SilenceUsage: true,
	}
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "YAML configuration file for burr")
	root.AddCommand(webCmd)

	//
	// Configuration
	//

	log.Printf("Reading configuration from %s...", cfgFile)
	cfgf, err := os.Open(cfgFile)
	if err != nil {
		log.Fatalf("unable to open %s: %v", cfgFile, err)
	}
	defer cfgf.Close()
	var cfg *Config
	if err := yaml.NewDecoder(cfgf).Decode(&cfg); err != nil {
		log.Fatalf("unable to parse YAML in %s: %v", cfgFile, err)
	}

	//
	// Run CLI
	//

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
