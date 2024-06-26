# Project Online Shop

This project includes a client part and tests. Below are the instructions for running the client and executing tests.

## Cloning the Repo

### Step 1: Clone the repository
Open your terminal and execute the following command:
```bash
git clone https://github.com/dezween/Online_shop.git

### Step 2: Navigate to the main folder
cd Online_shop

## Running the Client

### Step 1: Navigate to the Client Directory

Open your terminal and execute the following command:
```bash
cd client


### Step 2: Start the Server

Once in the client directory, run the following command to start the server:

```bash
python -m http.server 8000





## Running the Server
### Step 1: Navigate to the Server Directory
Open your terminal and execute the following command:

```bash
cd go-shop
### Step 2: Start the Server

Once in the server directory, run the following command to start the server:

```bash
go run main.go


## Working with the API
### Get User by ID
To get a user by their ID, use the following endpoint:

```bash
GET http://localhost:8080/users/{id}

Replace {id} with the actual user ID.

### Example:

```bash
GET http://localhost:8080/users/14

Get All Users
To retrieve a list of all users, use the following endpoint:

```bash
GET http://localhost:8080/users

Register a User
To register a new user, use the following endpoint:

```bash
POST http://localhost:8080/users

Include the necessary user data in the request body.


## Running Tests
### Step 1: Navigate to the Project Directory
Open your terminal and execute the following command:

```bash
cd tests

## Step 2: Execute Tests

```bash
go test