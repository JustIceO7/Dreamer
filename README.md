# Dreamer (Stable Diffusion Discord Bot)
Discord Bot written in Golang to generate images through Automatic1111 API from https://github.com/AUTOMATIC1111/stable-diffusion-webui.

## Queue System

This bot uses a queue system to manage image generation tasks. All requests are processed in the order they are received.

## Setup
1. Clone this repository:
```
git clone https://github.com/JustIceO7/Dreamer.git
```

2. Install Golang https://golang.org/dl/ (requires Golang version >1.23.1).

3. Create a .env file in the root directory with:
```
export discord_token=your_discord_token
export discord_app_id=your_app_id
export master=your_discord_id
```
**Note:** Users listed as `master` in `.env` have priority in the queue.

4. Ensure Automatic1111 is running with the `--api flag`.

5. Add the bot to your Discord server with permissions to post messages, upload files, react, and mention users.

6. Run the bot:
```
go run main.go
```

Discord Bot: `http://localhost:8080`.

Stable Diffusion: `http://localhost:7860`.
## Commands
`&help` – Shows all available commands.

`/generate <prompt>` - Generates an image from the given text prompt.

## Go
