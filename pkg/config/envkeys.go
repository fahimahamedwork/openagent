// OpenAgent - Ultra-lightweight personal AI agent
// License: MIT
//
// Copyright (c) 2026 OpenAgent contributors

package config

import (
	"os"
	"path/filepath"

	"github.com/fahimahamedwork/openagent/pkg"
)

// Runtime environment variable keys for the openagent process.
// These control the location of files and binaries at runtime and are read
// directly via os.Getenv / os.LookupEnv. All openagent-specific keys use the
// OPENAGENT_ prefix. Reference these constants instead of inline string
// literals to keep all supported knobs visible in one place and to prevent
// typos.
const (
	// EnvHome overrides the base directory for all openagent data
	// (config, workspace, skills, auth store, …).
	// Default: ~/.openagent
	EnvHome = "OPENAGENT_HOME"

	// EnvConfig overrides the full path to the JSON config file.
	// Default: $OPENAGENT_HOME/config.json
	EnvConfig = "OPENAGENT_CONFIG"

	// EnvBuiltinSkills overrides the directory from which built-in
	// skills are loaded.
	// Default: <cwd>/skills
	EnvBuiltinSkills = "OPENAGENT_BUILTIN_SKILLS"

	// EnvBinary overrides the path to the openagent executable.
	// Used by the web launcher when spawning the gateway subprocess.
	// Default: resolved from the same directory as the current executable.
	EnvBinary = "OPENAGENT_BINARY"

	// EnvGatewayHost overrides the host address for the gateway server.
	// Default: "127.0.0.1"
	EnvGatewayHost = "OPENAGENT_GATEWAY_HOST"
)

func GetHome() string {
	homePath, _ := os.UserHomeDir()
	if openagentHome := os.Getenv(EnvHome); openagentHome != "" {
		homePath = openagentHome
	} else if homePath != "" {
		homePath = filepath.Join(homePath, pkg.DefaultOpenAgentHome)
	}
	if homePath == "" {
		homePath = "."
	}
	return homePath
}
