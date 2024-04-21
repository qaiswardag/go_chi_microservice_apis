# About

This repository contains a microservice application built using Go and Chi framework, designed to provide various APIs.

# Overview

This application serves as a platform for managing and providing APIs efficiently. Leveraging the Go programming language and the Chi framework, it offers a robust and scalable solution for building microservices.

# Features

The application is structured as a microservice, allowing for modular development and scalability.

# API Management

With a focus on providing APIs, the platform facilitates easy creation, deployment, and management of various endpoints.

# Go and Chi Framework

Built using Go programming language and the lightweight Chi framework, ensuring high performance and flexibility.

Once the application is running, access the APIs through the defined endpoints.

# Build

Build the app for Linux:
`GOOS=linux GOARCH=amd64 go build -o go_chi`

Run the app on server:
`./go_chi`

# Curl

Check Post request

Check the GET request using the following curl command
Show all http headers `-v`
`curl http://localhost:{PORT}/hello -v`

Check the POST request with the following curl command
`curl -X POST http://localhost:{PORT} -v`

# Contribution

Contributions are welcome! If you'd like to enhance the application or fix any issues, feel free to fork the repository and submit pull requests.

# License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as per the terms of the license.
