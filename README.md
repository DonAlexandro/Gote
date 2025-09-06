# Gote

A simple and elegant Terminal User Interface (TUI) notes application written in Go using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework.

## Features

- **Terminal-based interface** - Clean and intuitive TUI for managing notes
- **SQLite storage** - Persistent storage using SQLite database
- **Create and edit notes** - Add new notes with titles and body content
- **Browse notes** - Navigate through your notes with keyboard shortcuts

## Installation

### Prerequisites

- Go 1.24.2 or higher
- SQLite3

### Build from source

1. Clone the repository:

   ```bash
   git clone https://github.com/DonAlexandro/Gote.git
   cd Gote
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the application:

   ```bash
   go build -o gote
   ```

4. Run the application:

   ```bash
   ./gote
   ```

## Usage

### Keyboard Shortcuts

#### List View (Main screen)

- `n` - Create a new note
- `↑/k` - Move selection up
- `↓/j` - Move selection down
- `Enter` - Edit selected note
- `q` or `Ctrl+C` - Quit application

#### Title View (New note)

- `Enter` - Confirm title and proceed to body editing
- `Esc` - Cancel and return to list view

#### Body View (Note editing)

- `Ctrl+S` - Save note and return to list view
- `Esc` - Cancel editing and return to list view

## Database

The application stores notes in a SQLite database file located at `./notes.db` relative to the executable. The database contains a single `notes` table with the following structure:

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - Common UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver
