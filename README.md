
# 🤖 Slack Bot with Wolfram Alpha Integration 🧠

This project is a **Slack bot** that allows users to ask questions and receive answers from **Wolfram Alpha**. The bot uses **Wit.ai** for Natural Language Processing (NLP) to understand user queries and extract structured data. Perfect for automating answers to complex questions! 🚀

---

## 🛠️ Prerequisites

Before setting up the project locally, ensure you have the following:

1. **Go** 
   Make sure Go is installed on your system. You can download it [here](https://golang.org/dl/)

2. **Slack** 💬  
   You need a Slack workspace and a Slack bot token. Create a new bot in your [Slack](https://app.slack.com/) workspace and obtain the bot token.

3. **Wit.ai** 🧠  
   Create a Wit.ai account and set up a new app. Obtain the [Wit](https://wit.ai/apps).ai token.

4. **Wolfram Alpha** 🔢  
   Create an account on [Wolfram](https://developer.wolframalpha.com/) Alpha and obtain an App ID.

5. **Environment Variables** 🔑  
   Create a `.env` file in the root directory of the project with the following variables:

   ```plaintext
   SLACK_BOT_TOKEN=your-slack-bot-token
   SLACK_APP_TOKEN=your-slack-app-token
   WIT_AI_TOKEN=your-wit-ai-token
   WOLFRAM_APP_ID=your-wolfram-app-id
   ```

---

## Set up environment variables 🔧  
   Create a `.env` file in the root directory and add your tokens as described in the prerequisites.

---

## 🏃 Running the Bot

1. **Start the bot** ▶️  
   Run the following command to start the bot:
   ```bash
   go run main.go
   ```

2. **Interact with the bot** 💬  
   In your Slack workspace, send a message to the bot in the following format:
   ```
   My question is - <your question>
   ```
   For example:
   ```
   My question is - What is the capital of Germany?
   ```

   The bot will process your question using Wit.ai, send it to Wolfram Alpha, and reply with the answer. 🎉

---

## 📦 Dependencies

This project uses the following Go packages:
- `github.com/joho/godotenv`: Loads environment variables from a `.env` file.
- `github.com/shomali11/slacker`: A Slack bot framework for Go.
- `github.com/wit-ai/wit-go`: A Go client for Wit.ai.
- `github.com/Edw590/go-wolfram`: A Go client for Wolfram Alpha.
- `github.com/tidwall/gjson`: A package for extracting values from JSON.

---

## ✨ Example Interaction

**User**:  
`My question is - What is the speed of light?`

**Bot**:  
`Wolfram answers: The speed of light in a vacuum is approximately 299,792 kilometers per second.`

---

Enjoy building and using your Slack bot! 🎉 If you have any questions or need further assistance, feel free to reach out. 😊

---
