# URL Shortener Flask API


This is a solution to the URL shortener project on [roadmap.sh](https://roadmap.sh/projects/url-shortening-service)

This project built using Go, MySQL and used Gin Framework as a server. This service allow to create, retrieve, update, and delete shortened URLs, as well as the usage statistics of each shortened URL.

The method using to create the unique short code is creating a random string of length 4 contains of alphanumber. The possibility of uniqueness is 62^4.

## Features

- **Shorten URLs:** Generate a unique, shortened version of any URL.
- **Retrieve URLs:** Look up the all original URL.
- **Update URLs:** Modify the original URL associated with a shortened code.
- **Delete URLs:** Remove shortened URLs from the system.
- **View Stats:** Track how many times a shortened URL has been accessed.
- **Redirect:** Redirect to associated origin url

## Project Structure

- **`main.go`**: Entry point to start the application
- **`db`**: Contains initialization for connecting to database.
- **`models`**: Contain all data type
- **`routes`**: Contain function for handle the routing
- **`utils`**: Contain utilization function that used in this project

## Prerequisites

- Go 1.23.4
- MySQL

## Setup and Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/short-url.git
   cd short-url
   ```
   
2. **Create database:**
   Create the database from your own MySQL local. The database is in the `internal/db` folder

3. **Configure environment variables**
   ```
   # Configuration for Database
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
   ```
   
4. **Build the execution program:**
   ```
      make build
   ```

5. **Execute the program**
   ```
      make execute
   ```

## API Endpoints

### **POST /shorten**

Create a shortened URL.

- **Request:**
  - **Method:** POST
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "url": "http://example.com"
    }
    ```

- **Response:**
  - **Status Code:** 201 Created
  - **Body:**
    ```json
    {
      "message": "data created",
      "url": {
        "id": 1,
        "url": "https://www.example.com",
        "short_code": "QP47",
        "access_count": null,
        "created_at": "2025-01-09T08:05:27.742062+07:00",
        "updated_at": "2025-01-09T08:05:27.742063+07:00"
      }
    }
    ```

- **Description:** Generates a unique shortened URL for the provided original URL.

### **GET /shorten**

Retrieve all original shorten urls.

- **Request:**
  - **Method:** GET
  - **URL Parameter:** `short_code` (the unique shortened code)

- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    [
      {
        "id": 5,
        "url": "https://www.example.com",
        "short_code": "QP47",
        "access_count": 0,
        "created_at": "2025-01-09T01:05:28Z",
        "updated_at": "2025-01-09T01:05:28Z"
      }
    ]
    ```

- **Description:** Retrieves all the original URLs.

### **PUT /shorten/<short_code>**

Update the original URL associated with a shortened code.

- **Request:**
  - **Method:** PUT
  - **Content-Type:** `application/json`
  - **URL Parameter:** `short_code` (the unique shortened code)
  - **Body:**
    ```json
    {
      "url": "http://newexample.com"
    }
    ```

- **Response:**
  - **Status Code:** 201 CREATED
  - **Body:**
    ```json
    {
      "message": "data updated"
    }
    ```

- **Description:** Updates the original URL associated with the given shortened code.

### **DELETE /shorten/<short_code>**

Delete a shortened URL.

- **Request:**
  - **Method:** DELETE
  - **URL Parameter:** `short_code` (the unique shortened code)

- **Response:**
  - **Status Code:** 204 No Content
  - **Body:**
    ```json
    {
      "message": "data deleted"
    }
    ```

- **Description:** Deletes the shortened URL associated with the given shortened code.

### **GET /shorten/<short_code>/stats**

Get the usage statistics for a shortened URL.

- **Request:**
  - **Method:** GET
  - **URL Parameter:** `short_code` (the unique shortened code)

- **Response:**
  - **Status Code:** 200 OK
  - **Body:**
    ```json
    {
      "id": 5,
      "url": "https://www.example.com",
      "short_code": "QP47",
      "access_count": 0,
      "created_at": "2025-01-09T01:05:28Z",
      "updated_at": "2025-01-09T01:05:28Z"
    }
    ```

- **Description:** Retrieves the usage statistics, including the access count, for the given shortened URL.

### **GET /<short_code>**

Use the shortend link to redirect.

- **Request:**
  - **Method:** GET
  - **URL Parameter:** `short_code` (the unique shortened code)

- **Response:**
  - **Status Code:** 303 See Other

- **Description:** Redirect to associated origin URL.
