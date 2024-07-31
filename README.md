# Go with PostgREST

This repository demonstrates how to set up and use PostgREST to create a RESTful API from a PostgreSQL database.

## Features

- Automatic REST API generation from PostgreSQL database
- CRUD operations
- Secure access control

## Requirements

- PostgreSQL
- Docker (optional but recommended for easy setup)

## Getting Started

### Installation

1. **Clone the repository:**

    ```
    git clone https://github.com/muthukumar89uk/go-restapi-postgrest.git
    ```
  Click here to directly [download it](https://github.com/muthukumar89uk/go-restapi-postgrest/zipball/master). 

### Install dependencies:

          go mod tidy

### Run the Application
  1. Run the Server
   
       ```
          go run .
       ```   
  2. The server will start on `http://localhost:8080`.

### PostgREST Configuration

1. **Download PostgREST:**

    Download the latest release of PostgREST from the [official GitHub releases page](https://github.com/PostgREST/postgrest/releases).

2. **Create a Configuration File:**

    Create a `postgrest.conf` file with the following content:

    ```ini
    db-uri = "postgres://postgres:mysecretpassword@localhost:5432/mydatabase"
    db-schema = "api"
    db-anon-role = "postgres"
   ```
3. **Run PostgREST Using Docker Compose Yaml File**
   
   - Go to the docker-compose.yaml file directory then run the `docker-compose up` command.
   - This command will pull the necessary Docker images (if not already downloaded), create the containers, and start the services.

### Running PostgREST

1. **Start PostgREST:**

    ```
    ./postgrest postgrest.conf
    ```

2. **Access the API:**

    The API will be available at `http://localhost:3000/<table-name>`. You can test the endpoints using tools like curl or Postman.

    **Example Endpoints:**

    - `GET /tasks` - Retrieve all items
    - `GET /tasks/:id` - Retrieve an item by ID
    - `POST /tasks` - Create a new item
    - `PATCH /tasks/:id` - Update an existing item
    - `DELETE /tasks/:id` - Delete an item

    **Example cURL requests:**

    - Get all items:

        ```
        GET http://localhost:3000/items
        ```

    - Create a new item:

        ```
        POST http://localhost:3000/items
        ```

    - Update an existing item:

        ```
        PUT http://localhost:3000/tasks/:id
        ```

    - Delete an item:

        ```
        DELETE http://localhost:3000/tasks/:id
        ```
        
### Acknowledgements

- [PostgREST](https://postgrest.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
