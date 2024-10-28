# Telegram Bot Workshop in Go

This project outlines the setup of a Telegram bot using Go. Follow the steps below to get your bot up and running.

## Prerequisites

Before starting, make sure you have Docker installed on your system. If you do not have it, visit [Docker's official site](https://www.docker.com/products/docker-desktop) to download and install.

## Project Setup

### Step 1: Prepare Environment Variables

Copy the example environment file to create your own environment variables file:

```bash
cp .env.example .env
```

### Step 2: Create Telegram Bot
Create a new bot by chatting with @BotFather on Telegram and following the instructions to get a token.

### Step 3: Add Bot Token
Once you have your bot's token, add it to your .env file under the key TELEGRAM_TOKEN:

plaintext

```bash
TELEGRAM_TOKEN=your_bot_token_here
```
Replace your_bot_token_here with the token you obtained from BotFather.

### Step 4: Start the Services
Run Docker Compose to start the services:

```bash


docker-compose up
```
### Step 5: Access the Web Interface
Navigate to localhost:8080 or 0.0.0.0:8080 on macOS in your web browser. Log in using the username and password that you specified in the .env file under DB_USER and DB_PASSWORD keys.

### Step 6: Initialize Database
Once logged in, add your desired entries to the quiz table in the database to manage your bot's quiz functionalities.

### Step 7: Run Your Bot
Now, you can interact with your bot in Telegram using the features you have set up.
[More information](https://round-bellflower-229.notion.site/Workshop-27-10-2024-12cb20b4840580b19f0cd79cb2d83787).
