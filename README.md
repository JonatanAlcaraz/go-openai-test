# Chatbot Project in Go

This project is a simple chatbot implemented in Go that uses the OpenAI API to generate responses to user inputs.

## Requirements

- Go 1.16 or higher
- An OpenAI API key

## Installation

1. Clone the repository:

   ```bash
   git clone <REPOSITORY_URL>
   cd <REPOSITORY_NAME>
   ```

2. Install the dependencies:

   ```bash
   go get github.com/sashabaranov/go-openai
   go get github.com/joho/godotenv // Add this line
   ```

3. Create a `.env` file in the root of the project and add your API token:

   ```plaintext
   OPENAI_AUTH_TOKEN="your_api_key"
   ```


## Usage

To run the chatbot, use the following command:

```bash
go run main.go
```

Once the program is running, you can interact with the chatbot by typing your messages in the console. The chatbot will respond to your messages using the OpenAI API.

## Example Conversation

```
> Hello, how are you?
I'm fine, thank you. How about you?
> What can you do?
I can help you answer questions and hold a conversation.
```