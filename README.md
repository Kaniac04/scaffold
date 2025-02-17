# Scaffold CLI

A lightweight and efficient CLI tool to generate organized directory structures for various project types. This tool helps developers quickly scaffold industry-evolved project structures for Flask, Gin, and other frameworks.

## Features

- Quick project initialization with predefined structures
- Support for multiple project types and frameworks
- Easy-to-use command line interface
- View available project templates
- Preview directory structures before creation

## Installation

```bash
go install github.com/Kaniac04/Scaffold@latest
```

## Usage

### Basic Commands

```bash
scaffold [command] [flags]
```

### Available Commands

- `climb`: Initialize a new project with the desired structure
- `list`: Show available project types
- `view`: Show the directory structure for a project type
- `version`: Display the CLI version
- `help`: Show help message

### Command Details

#### Initialize a New Project

```bash
scaffold climb <project_name> --type <project_type>
```

This command creates a new project directory with the specified structure.

#### List Available Project Types

```bash
scaffold list
```

Shows all available project templates you can use.

#### View Project Structure

```bash
scaffold view --type <project_type>
```

Displays the directory structure for a specific project type before creation.

#### Get Version

```bash
scaffold version
```

Displays the current version of the CLI tool.

#### Get Help

```bash
scaffold help
```

Shows detailed help information about available commands and usage.

## Examples

Create a new Flask project:
```bash
scaffold climb myflaskapp --type flask
```

View Flask project structure:
```bash
scaffold view --type flask
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

Aayushmaan Garg
