package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the container name from the HOSTNAME environment variable
	containerName := os.Getenv("HOSTNAME")

	// Get some additional container information
	goVersion := runtime.Version()

        // Get the CUSTOM_ENV environment variable that you set with -e
        demoEnv := os.Getenv("DEMO_ENV")

	// Get the CUSTOM_ENV environment variable that you set with -e
        secretEnv := os.Getenv("DEMO_SECRET")
        
	// Get the PROJECT environment variable that you set with -e
        projectEnv := os.Getenv("PROJECT")

	// Get the CUSTOM_ENV environment variable that you set with -e
        envEnv := os.Getenv("ENV")

	// Create an HTML page with the information
	html := fmt.Sprintf(`
	<html>
        <head>
        <title>Container Info</title>
        <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Encode+Sans:wght@400;700&display=swap"> 
        <style>
            .container {
                font-family: 'Encode Sans', sans-serif;
                width: 300px;
                padding: 20px;
                border: 1px solid #ccc;
                box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                margin: 50px auto;
                background-color: #fff;
                text-align: center;
            }

            .header {
                font-family: 'Encode Sans', sans-serif;
                background-color: #7D00FF; /* Change the color as desired */
                color: #fff;
                padding: 10px;
                text-align: center;
                font-weight: bold;
                border-top-left-radius: 10px; /* Rounded top-left corner */
                border-top-right-radius: 10px; /* Rounded top-right corner */
                border-bottom-left-radius: 10px; /* Rounded top-left corner */
                border-bottom-right-radius: 10px;
            }

            .logo {
                height: 50px;
                widht: 40px;
                margin-top: 10px;
                margin-bottom: 20px;

            }
        </style>
        </head>
        <body>
        <div class="container">
            <img class="logo" src="https://www.shamanops.com/assets/img2/logo_shamanops.svg" alt="Logo Image">
            <div class="header">Container Information</div>
            <p><strong>PROJECT</strong><br> %s</p>
            <p><strong>ENV</strong><br> %s</p>   
            <p><strong>LANGUAGE VERSION</strong><br> %s</p>  
            <p><strong>CONTAINER NAME</strong><br> %s</p>
    	    <p><strong>DEMO ENVIRONMENT</strong><br> %s</p>
            <p><strong>DEMO SECRET</strong><br> %s</p>
	    <img class="logo" src="https://www.shamanops.com/assets/img2/golang.png" alt="logo-golang">
        </div>
        </body>
        </html>

	`, projectEnv, envEnv, goVersion, containerName, demoEnv, secretEnv)

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

