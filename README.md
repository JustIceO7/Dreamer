# Dreamer (Stable Diffusion Discord Bot)
Discord Bot written in Go to generate images through Automatic1111 API from https://github.com/AUTOMATIC1111/stable-diffusion-webui.

## Setup
Clone this repository.

Install Go https://golang.org/dl/.

`.env` file is required within root directory with the following variables. `master` indicates your own Discord ID.
```
export discord_token=
export discord_app_id=
export master=
```

Discord Bot: `http://localhost:8080`.

Stable Diffusion: `http://localhost:7860`.

1. Create Discord token and place within `.env`.
2. Add the discord bot within your Discord server. It needs permissions to post messages, listen to reactions, mentioning anyone and uploading files.
3. Navigate into root directory and run `go run main.go`.
4. Ensure Automatic1111 is ran alongside this Discord bot with `--api` flag.

## Commands
To access the list of commands use `&help`.

## Go
