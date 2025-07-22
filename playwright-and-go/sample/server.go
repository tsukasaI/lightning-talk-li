package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files
	http.HandleFunc("/", handleHome)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Playwright Demo App</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 600px;
            margin: 50px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background-color: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            text-align: center;
        }
        h1 {
            color: #333;
            margin-bottom: 30px;
        }
        #text-display {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
            margin: 30px 0;
            padding: 20px;
            background-color: #f8f9fa;
            border-radius: 5px;
            border: 2px solid #dee2e6;
        }
        #change-button {
            background-color: #007bff;
            color: white;
            border: none;
            padding: 15px 30px;
            font-size: 16px;
            border-radius: 5px;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        #change-button:hover {
            background-color: #0056b3;
        }
        #change-button:active {
            transform: translateY(1px);
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Playwright Demo App</h1>
        <div id="text-display">Hello, World!</div>
        <button id="change-button">Change Text</button>
    </div>

    <script>
        const textDisplay = document.getElementById('text-display');
        const changeButton = document.getElementById('change-button');

        const messages = [
            "Hello, World!",
            "Welcome to Playwright Testing!",
            "Go + Playwright = Amazing!",
            "Browser automation is fun!",
            "Testing made easy!",
            "Click me again!"
        ];

        let currentIndex = 0;

        changeButton.addEventListener('click', function() {
            currentIndex = (currentIndex + 1) % messages.length;
            textDisplay.textContent = messages[currentIndex];
        });
    </script>
</body>
</html>
	`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
