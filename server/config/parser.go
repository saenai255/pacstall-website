package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"pacstall.dev/webserver/log"

	"github.com/BurntSushi/toml"
)

const defaultConfigPath = "./webserver.toml"

var configPath = flag.String("config", defaultConfigPath, fmt.Sprintf("Path to configuration file. Default: %s", defaultConfigPath))
var IsProduction = false

func Load() {
	flag.Parse()
	cfg := loadConfig()
	setFeatureFlags(cfg.FeatureFlags)
	setLogging(cfg.Logging)
	setTCPServer(cfg.TCPServer)
	setPacstallPrograms(cfg.PacstallPrograms)
	IsProduction = cfg.Production

	log.Info.Printf("Loaded configuration from %v", *configPath)
	log.Info.Printf("Running in %v mode", func() string {
		mode := "development"
		if IsProduction {
			mode = "production"
		}
		return mode
	}())
	log.Debug.Printf("Feature flags configuration: %#v", FeatureFlags)
	log.Debug.Printf("Logging configuration: %#v", Logging)
	log.Debug.Printf("Server configuration: %#v", TCPServer)
	log.Debug.Printf("Pacstall Programs configuration: %#v", PacstallPrograms)
}

func loadConfig() tomlConfiguration {
	data := tomlConfiguration{}
	bytes, err := os.ReadFile(*configPath)
	if err != nil {
		log.Error.Fatalf("Could not read file '%s'\n%v", *configPath, err)
	}

	if err = toml.Unmarshal(bytes, &data); err != nil {
		log.Error.Fatalf("Could not parse file '%s'\n%v", *configPath, err)
	}

	validate(data)
	return data
}

func prettify(data tomlConfiguration) string {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Warn.Fatalf(err.Error())
	}
	return string(out)
}

func validate(data tomlConfiguration) {
	config_error := false

	defer func() {
		if config_error {
			os.Exit(1)
		}
	}()

	if data.PacstallPrograms.Path == "" {
		log.Error.Printf("Configuration file '%s' is missing required attribute `pacstall_programs.path\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.TempDir == "" {
		log.Error.Printf("Configuration file '%s' is missing required attribute `pacstall_programs.tmp_dir\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.UpdateInterval == 0 {
		log.Error.Printf("Configuration file '%s' is missing required attribute `pacstall_programs.update_interval\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.MaxOpenFiles == 0 {
		log.Error.Printf("Configuration file '%s' is missing required attribute `pacstall_programs.max_open_files\n`", *configPath)
		config_error = true

	}

	if data.TCPServer.Port == 0 {
		log.Error.Printf("Configuration file '%s' is missing required attribute `tcp_server.port\n`", *configPath)
		config_error = true

	}

	if data.TCPServer.PublicDir == "" {
		log.Error.Printf("Configuration file '%s' is missing required attribute `tcp_server.public_dir\n`", *configPath)
		config_error = true

	}
}
