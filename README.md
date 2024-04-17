#  Commit-AI

## Overview
Commit-AI is a CLI tool designed to automate the generation of meaningful commit messages for your Git repositories. Leveraging the power of OpenAI's GPT-3.5 Turbo, it analyzes your git diffs and crafts commit messages that are concise, descriptive, and follow best practices.

## Features
- **Automated Commit Messages**: Generate commit messages based on the changes in your git repository.
- **Customizable Prefixes**: Supports prefixes for different types of commits (feat, fix, refactor, chore).
- **Easy Integration**: Simple setup and usage within any Git repository.

## Requirements
- Go version 1.21.1 or higher.
- An OpenAI API key.

## Installation
1. Ensure you have Go installed on your system (version 1.21).
2. Clone the repository to your local machine.
3. Navigate to the cloned directory and run `go build && go install` to compile the application.

## Usage
To use Commit-AI, you need to provide your OpenAI API key using the `--apiKey` flag. Run the following command in your terminal:

```shell
commit-ai --apiKey YOUR_API_KEY
```

## Contributing
If you would like to contribute to the project, please fork the repository and make a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

