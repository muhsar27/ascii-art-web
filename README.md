# ASCII Art Web

## Description

ASCII Art Web is a web-based implementation of the ASCII Art project written in Go. It allows users to enter text through a graphical web interface and generate ASCII art using different banner styles.

The application runs an HTTP server that serves a web page where users can:

* Enter text
* Select an ASCII art banner style
* Submit the text to the server
* View the generated ASCII art result

Supported banner styles:

* standard
* shadow
* thinkertoy

The project demonstrates the use of:

* HTTP servers in Go
* HTML templates
* Form handling with POST requests
* File processing
* Error handling and HTTP status codes

---

## Authors

* Faith
* 

---

## Usage

### Prerequisites

* Go 1.20 or later installed

### Run the project

1. Clone the repository:

```bash
git clone <repository-url>
cd ascii-art-web
```

2. Start the server:

```bash
go run .
```

3. Open your browser and visit:

```text
http://localhost:8080
```

### Using the application

1. Enter text in the input field.
2. Select a banner style:

   * standard
   * shadow
   * thinkertoy
3. Click the **Generate** button.
4. The generated ASCII art will be displayed on the page.

---

## HTTP Endpoints

### GET /

Returns the main HTML page containing:

* Text input field
* Banner selection options
* Submit button

Response:

```http
HTTP/1.1 200 OK
```

### POST /ascii-art

Accepts form data containing:

* text
* banner

Generates ASCII art and returns the result.

Example form fields:

```text
text=Hello
banner=shadow
```

Response:

```http
HTTP/1.1 200 OK
```

---

## HTTP Status Codes

The application returns appropriate HTTP status codes:

| Status Code               | Description                                             |
| ------------------------- | ------------------------------------------------------- |
| 200 OK                    | Request completed successfully                          |
| 400 Bad Request           | Invalid request data or missing form values             |
| 404 Not Found             | Requested page, template, or banner file does not exist |
| 500 Internal Server Error | Unexpected server-side error                            |

---

## Project Structure

```text
ascii-art-web/
│
├── banners/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
│
├── templates/
│   ├── ascii.html
│   └── index.html
│
├── functions/
│   └── asciiartprint.go
│
├── main.go
└── README.md
```

---

## Implementation Details

### ASCII Art Generation Algorithm

1. Receive text input from the client.
2. Read the selected banner file.
3. Convert escaped newline characters (`\n`) into actual line breaks.
4. Split the input text into separate lines.
5. For each character:

   * Calculate its position in the ASCII table.
   * Locate the corresponding character representation in the banner file.
   * Append the matching ASCII art lines to the output.
6. Combine all generated lines into a final result.
7. Return the generated ASCII art to the user.

### Request Flow

```text
Browser
   │
   ▼
GET /
   │
   ▼
HTML Form
   │
   ▼
POST /ascii-art
   │
   ▼
Validate Input
   │
   ▼
Load Banner
   │
   ▼
Generate ASCII Art
   │
   ▼
Render Result
```

### Error Handling

The server validates:

* Request method
* Request path
* Input text
* Banner selection
* Template availability
* Banner file availability

Any failure returns the appropriate HTTP status code and error message.

---

## Features

* Web graphical interface
* Multiple banner support
* Form-based input submission
* HTTP status code handling
* Template rendering
* ASCII art generation using Go
* Clean project structure
