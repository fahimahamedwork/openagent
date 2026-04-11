// OpenAgent - Ultra-lightweight personal AI agent
// License: MIT
//
// Copyright (c) 2026 OpenAgent contributors

package main

import (
        "fmt"
        "os"
        "time"

        "github.com/spf13/cobra"

        "github.com/fahimahamedwork/openagent/cmd/openagent/internal"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/agent"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/auth"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/cron"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/gateway"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/model"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/onboard"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/skills"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/status"
        "github.com/fahimahamedwork/openagent/cmd/openagent/internal/version"
        "github.com/fahimahamedwork/openagent/pkg/config"
        "github.com/fahimahamedwork/openagent/pkg/updater"
)

func NewOpenAgentCommand() *cobra.Command {
        short := fmt.Sprintf("%s openagent - Personal AI Assistant %s\n\n", internal.Logo, config.GetVersion())

        cmd := &cobra.Command{
                Use:     "openagent",
                Short:   short,
                Example: "openagent version",
        }

        cmd.AddCommand(
                onboard.NewOnboardCommand(),
                agent.NewAgentCommand(),
                auth.NewAuthCommand(),
                gateway.NewGatewayCommand(),
                status.NewStatusCommand(),
                cron.NewCronCommand(),
                skills.NewSkillsCommand(),
                model.NewModelCommand(),
                updater.NewUpdateCommand("openagent"),
                version.NewVersionCommand(),
        )

        return cmd
}

const (
        colorBlue = "\033[1;38;2;62;93;185m"
        colorRed  = "\033[1;38;2;213;70;70m"
        banner    = "\r\n" +
                colorBlue + "██████╗ ██╗ ██████╗ ██████╗ " + colorRed + " ██████╗██╗      █████╗ ██╗    ██╗\n" +
                colorBlue + "██╔══██╗██║██╔════╝██╔═══██╗" + colorRed + "██╔════╝██║     ██╔══██╗██║    ██║\n" +
                colorBlue + "██████╔╝██║██║     ██║   ██║" + colorRed + "██║     ██║     ███████║██║ █╗ ██║\n" +
                colorBlue + "██╔═══╝ ██║██║     ██║   ██║" + colorRed + "██║     ██║     ██╔══██║██║███╗██║\n" +
                colorBlue + "██║     ██║╚██████╗╚██████╔╝" + colorRed + "╚██████╗███████╗██║  ██║╚███╔███╔╝\n" +
                colorBlue + "╚═╝     ╚═╝ ╚═════╝ ╚═════╝ " + colorRed + " ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝\n " +
                "\033[0m\r\n"
)

func main() {
        fmt.Printf("%s", banner)

        tz_env := os.Getenv("TZ")
        if tz_env != "" {
                fmt.Println("TZ environment:", tz_env)
                zoneinfo_env := os.Getenv("ZONEINFO")
                fmt.Println("ZONEINFO environment:", zoneinfo_env)
                loc, err := time.LoadLocation(tz_env)
                if err != nil {
                        fmt.Println("Error loading time zone:", err)
                } else {
                        fmt.Println("Time zone loaded successfully:", loc)
                        time.Local = loc //nolint:gosmopolitan // We intentionally set local timezone from TZ env
                }
        }

        cmd := NewOpenAgentCommand()
        if err := cmd.Execute(); err != nil {
                os.Exit(1)
        }
}
