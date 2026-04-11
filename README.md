<div align="center">
<img src="assets/logo.webp" alt="OpenAgent" width="140">

# OpenAgent

**Your personal AI assistant. Runs anywhere.**

Lightweight · Fast · Open Source

[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Discord](https://img.shields.io/badge/Discord-Community-4c60eb?style=flat&logo=discord&logoColor=white)](https://discord.gg/V4sAZ9XWpN)
</div>

---

**OpenAgent** is a lightweight AI assistant built in Go. It connects to 30+ LLM providers and 18+ chat platforms, runs on anything from a $10 single-board computer to a full server, and boots in under a second.

No bloat. No dependencies. Just a single binary.

## Why OpenAgent?

- **10MB RAM** — runs on hardware most AI tools can't
- **<1s boot** — even on a 0.6GHz processor
- **Single binary** — no runtime, no container, no hassle
- **30+ LLM providers** — OpenAI, Anthropic, Google, Ollama, and more
- **18+ chat platforms** — Telegram, Discord, WhatsApp, WeChat, Slack, and more
- **MCP support** — extend with any Model Context Protocol server
- **Smart routing** — simple queries go to cheap models, complex ones use the heavy lifters

## Quick Start

### Download

Grab the latest release for your platform from [GitHub Releases](https://github.com/fahimahamedwork/openagent/releases).

### Or build from source

```bash
git clone https://github.com/fahimahamedwork/openagent.git
cd openagent
make deps
make build
```

### First run

```bash
openagent onboard    # creates config and workspace
openagent gateway    # starts the gateway
```

That's it. OpenAgent will guide you through setting up a provider and a channel.

## WebUI Launcher

The easiest way to get started — especially on desktop:

```bash
openagent-launcher
# Opens at http://localhost:18800
```

Configure providers, channels, and chat — all from your browser.

## CLI Usage

```bash
openagent agent -m "What is 2+2?"    # one-shot question
openagent agent                       # interactive chat
openagent gateway                     # start gateway for chat apps
openagent status                      # check status
openagent model                       # view or switch model
openagent cron list                   # list scheduled jobs
openagent skills list                 # list installed skills
```

## Providers

OpenAgent supports 30+ LLM providers. Configure them in `model_list` using `protocol/model` format:

| Provider | Protocol | Key Required |
|----------|----------|-------------|
| OpenAI | `openai/` | Yes |
| Anthropic | `anthropic/` | Yes |
| Google Gemini | `gemini/` | Yes |
| OpenRouter | `openrouter/` | Yes |
| DeepSeek | `deepseek/` | Yes |
| Qwen | `qwen/` | Yes |
| Groq | `groq/` | Yes |
| Ollama | `ollama/` | No (local) |
| vLLM | `vllm/` | No (local) |
| Azure OpenAI | `azure/` | Yes |
| AWS Bedrock | `bedrock/` | AWS credentials |
| GitHub Copilot | `github-copilot/` | OAuth |

And many more — see [docs.openagent.io](https://docs.openagent.io) for the full list.

### Local models with Ollama

```json
{
  "model_list": [
    {
      "model_name": "local-llama",
      "model": "ollama/llama3.1:8b",
      "api_base": "http://localhost:11434/v1"
    }
  ]
}
```

## Channels

Talk to OpenAgent through your favorite messaging app:

| Channel | Setup | Protocol |
|---------|-------|----------|
| Telegram | Easy (bot token) | Long polling |
| Discord | Easy (bot token) | WebSocket |
| WhatsApp | Easy (QR scan) | Native / Bridge |
| WeChat | Easy (QR scan) | iLink API |
| Slack | Easy (bot token) | Socket Mode |
| QQ | Easy (AppID + Secret) | WebSocket |
| WeCom | Easy (QR login) | WebSocket |
| Feishu / Lark | Medium (App ID + Secret) | WebSocket |
| DingTalk | Medium (client credentials) | Stream |
| Matrix | Medium (homeserver + token) | Sync API |
| LINE | Medium (credentials + webhook) | Webhook |
| IRC | Medium (server + nick) | IRC protocol |

Plus more — see [docs.openagent.io](https://docs.openagent.io) for setup guides.

## MCP (Model Context Protocol)

Connect any MCP server to extend OpenAgent's capabilities:

```json
{
  "tools": {
    "mcp": {
      "enabled": true,
      "servers": {
        "filesystem": {
          "enabled": true,
          "command": "npx",
          "args": ["-y", "@modelcontextprotocol/server-filesystem", "/tmp"]
        }
      }
    }
  }
}
```

## Docker

```bash
git clone https://github.com/fahimahamedwork/openagent.git
cd openagent

# First run — generates config
docker compose -f docker/docker-compose.yml --profile launcher up

# Set your API keys, then start
vim docker/data/config.json
docker compose -f docker/docker-compose.yml --profile launcher up -d
```

## Hardware

OpenAgent runs on virtually anything:

- $10 Linux boards (RISC-V, ARM, MIPS, x86)
- Raspberry Pi (including Zero 2W)
- Android phones (via APK or Termux)
- Docker containers
- Any standard Linux/macOS/Windows machine

## Documentation

Full documentation is available at **[docs.openagent.io](https://docs.openagent.io)**.

## Contributing

PRs welcome. The codebase is intentionally small and readable.

## License

[MIT](LICENSE)
