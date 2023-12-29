# HelloWellness: Golang Microservices WebServer for Hypertension Health
HelloWellness is a comprehensive Golang-based microservices WebServer designed to address the health needs of individuals dealing with hypertension. This project leverages Docker for containerization, enabling seamless deployment and scalability.

# Overview
The HelloWellness project is built with a microservices architecture, providing modularity and flexibility in development and deployment. Each microservice is responsible for specific functionalities related to hypertension health management.

# Features
User Management: Register and manage user profiles.

Technologies Used
Golang: The primary programming language for the WebServer and microservices.
Docker: Containerization for seamless deployment and scalability.
Gin: A lightweight Golang web framework for building the WebServer.
MySQL: Database management for storing user and health-related data.

# Getting Started
Follow these steps to get the HelloWellness project up and running on your local machine.

# Prerequisites
Docker installed on your machine.
Golang installed for local development.

# Installation
Clone the repository:
`git clone https://github.com/your-username/HelloWellness.git`

Navigate to the project directory:
`cd HelloWellness`

# INSTRUCTIONS

For the database, I am using MySQL in a container and running it with Docker Compose. Use `docker compose up/down` to start Docker Compose (create an equivalent for Windows if necessary).

You need to specify one of the following as an environment variable (For production, use your own environment variables information.):

MYSQL_ROOT_PASSWORD
MYSQL_ALLOW_EMPTY_PASSWORD
MYSQL_RANDOM_ROOT_PASSWORD

MYSQL_USER=hellowellness
MYSQL_ROOT_PASSWORD=123
MYSQL_DATABASE=hellowellness

To populate the Test database, use db.sql (execute it with MySQL Workbench after connecting to the database with user: root, password: 123).
After this, you can start the Gin server -> `go run main.go`
For the test, use curl -> `curl -X GET 0.0.0.0:8080/ping`
For the TEST, use -> `curl -X GET http://localhost:8081/api/v1/test`, as the port is defined as 8081 in the config.
To perform a POST TEST, use: `curl -X POST -H "Content-Type: application/json" -d '{"Name": "Test Product", "Email": "test@example.com", "Price": 9.99}' http://localhost:8081/api/v1/products`

MySQL container tips:

I have mapped in compose so that inside the project in `./mysql_container`, there is the data of the MySQL server.
You can check the logs in `./mysql_container/log/mysql.log`.
To display this file, check the "configure logging" section in db.sql.
Docker tips:

Docker has changed, at least on macOS; you need to run the app for the daemon to start running.
Run "brew install mysqlworkbench" to have the bench to connect to the database and run queries.
The content with the database files is locally mapped in the project in mysql_container (the folder is created when you run the container).
Go Gin hot reload:

Use a package called AIR -> `go install github.com/cosmtrek/air@latest`.
Install it using `install.sh`.
The binary is here on macOS -> `echo $(go env GOPATH)/bin/air`.
Add that to the path in your `~/.bashrc` or `~/.zshrc` to recognize the command in the terminal.

`alias air='$(go env GOPATH)/bin/air'`

Restart the terminal to load the new configuration.
The configuration for the air project is in .air.conf.
Then run it with this line:

`air -c .air.conf`

Voila! You have hot reload for Gin.


# Usage
Create a user account on the WebServer.
Log in and start monitoring health metrics.
Set reminders and receive alerts for medication and appointments.

# Contributing
Javier Charria Gómez - HelloWellness Lead Software Engineer
If you'd like to contribute to HelloWellness, please follow our Contribution Guidelines.

# License
This project is licensed under the MIT License - see the LICENSE file for details.

# Acknowledgments
The HelloWellness team acknowledges the importance of open-source contributions and collaboration.
Feel free to customize this README to fit the specific details of your project. Additionally, consider adding sections such as API documentation, testing instructions, and deployment guidelines to make the README even more informative.