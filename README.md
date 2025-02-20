# Ascii-Art-Web

Ascii-Art-Web is a web application that provides a graphical user interface (GUI) for interacting with ASCII banners, **Ascii-Art**. The website allows users to input text and select from different banner styles to generate stylized ASCII art.

The application runs a Go-based HTTP server that serves an interactive webpage where users can submit text and choose from various banners (shadow, standard, thinkertoy). The result is then displayed on a new page.

## Description

This project serves as a web version of your Ascii-Art project, enabling users to:

- Input text for conversion into ASCII art.
- Choose from three predefined ASCII banner styles: **shadow**, **standard**, and **thinkertoy**.
- Submit their text to the Go server, which will generate and return the ASCII art.
- View the result on a separate result page.

## Authors

    - azraji
    - ychatoua 

## Usage: How to Run

### Requirements

- Go 
- HTML, CSS

### Running the Server

1. Clone the repository:

   ```bash
   git clone https://learn.zone01oujda.ma/git/azraji/ascii-art-web-stylize.git
   ```

2. Navigate to the project directory:

   ```bash
   cd ascii-art-web-stylize
   ```

3. Build and run the Go server:

   ```bash
   cd main
   go run .
   ```

4. Visit `http://localhost:8080` in your web browser to access the application.

### Stopping the Server

To stop the server, use `Ctrl+C` in your terminal where the Go server is running.

## Implementation Details: Algorithm


### Banner Styles

The three available banner styles are:

1. **Shadow**
2. **Standard**
3. **Thinkertoy**

Each banner style transforms the input text into stylized ASCII art, and users can select their preferred style using a radio button.

### HTTP Status Codes

- **200 OK**: The request was successful, and the server returned a valid response.
- **404 Not Found**: The requested template or banner was not found.
- **400 Bad Request**: Invalid form input or missing data in the request.
- **500 Internal Server Error**: A server-side error occurred (e.g., unhandled exceptions).

## Instructions

### Project Structure

```plaintext
ascii-art-web/
├── main  
|     |               # Go server code
|     ├── templates/            # HTML templates
|     │   ├── index.html        # Main page template
|     │   └── result.html       # Result page template
|     ├── css/
|     │   ├── style.css        
|     │   └── styleresult.css       
|     ├── main.go
|     ├── shadow.txt
|     ├── standard.txt
|     └──  thinkertoy.txt
|
|
|── README.md             # Project documentation
|── func.go
└──go.mod

### HTTP Server (Go)

The server is written in Go, and the main logic for handling routes and generating the ASCII art is in `main.go`. This file sets up the routes, handles POST requests, and uses Go's `html/template` package to render the HTML templates.

### Templates

All HTML templates are located in the `templates` directory and are used to render dynamic content. The templates use Go's `{{ . }}` syntax for template variables.

