# Adarsha School Website

This repository hosts the backend for the Adarsha School Website, a web application designed to provide information about the school, manage content, and potentially handle other school-related functionalities.

## Features

*   **Dynamic Content Serving:** Utilizes Go templates for rendering dynamic web pages.
*   **Modular Routing:** Employs the `go-chi/chi` router for organized and efficient API routing.
*   **CORS Support:** Integrated `go-chi/cors` for handling Cross-Origin Resource Sharing.
*   **Static File Serving:** Serves static assets such as images and videos (e.g., `bg-vid.mp4`, `logo.jpg`).
*   **Configurable Environment:** Uses a `.env` file for managing environment-specific configurations.

## Technologies Used

*   **Go:** The primary programming language for the backend.
*   **go-chi/chi/v5:** A lightweight, idiomatic, and composable router for building HTTP services in Go.
*   **go-chi/cors:** A flexible CORS middleware for `go-chi/chi`.
*   **HTML:** For structuring the web content, with Go templates for dynamic rendering.

## Setup and Installation

To get this project up and running on your local machine, follow these steps:

### Prerequisites

*   Go (version 1.24.6 or later recommended)
*   Git

### Clone the Repository

```bash
git clone https://github.com/guruorgoru/adarsha-server.git
cd adarsha-server
```

### Install Dependencies

```bash
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory of the project and populate it with the following environment variables:

```
PORT=8414
DB_URL="https://your-supabase-url.supabase.co"
DB_KEY="your-supabase-anon-key"
ADMIN_EMAIL="admin@example.com"
ADMIN_PASSWORD="your_admin_password"
COOKIE_NAME="admin_session"
```

**Note:** Replace `"your-supabase-url.supabase.co"`, `"your-supabase-anon-key"`, `"admin@example.com"`, and `"your_admin_password"` with your actual Supabase URL, API key, and desired admin credentials.

## Running the Application

To start the server, run the following command from the project root:

```bash
go run cmd/main.go
```

The server will typically start on the port specified in your `.env` file (default: `8414`). You can access the application in your web browser at `http://localhost:8414`.

## Project Structure

```
.
├── .air.toml             # Air configuration for live-reloading (development)
├── .env                  # Environment variables
├── go.mod                # Go module file
├── go.sum                # Go module checksums
├── cmd/
│   └── main.go           # Main application entry point
├── internal/
│   └── routes/
│       ├── handler.go    # HTTP handlers
│       └── router.go     # Chi router setup
├── static/
│   ├── bg-vid.mp4        # Background video
│   ├── favicon.ico       # Favicon
│   └── logo.jpg          # School logo
├── tmp/
│   └── main              # Temporary build output
└── views/
    ├── index.html        # Main HTML template
    └── partials/         # HTML partials for common components
        ├── footer.html
        ├── home-body.html
        └── navbar.html
```
