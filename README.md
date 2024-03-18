# url-shortener
This project is a URL shortening service enabling to shorten long URLs into
eight-character codes - short URLs that redirect to original, longer ones.

### Technologies and Frameworks
The project uses the following technologies and frameworks:

- Go programming language
- Zap Logger for logging
- Redis for memory storage
- PostgreSQL for database storage
- Docker for containerization

### Installation and Running

```
make run_memory
```
Initiates the program with memory as the data storage option. 

```
make run_db
```
Initiates the program with a database as the data storage option.

Chosen command will start the containers and configure the necessary services.

### API Endpoints
The application provides the following API endpoints:

1. Alias Creating Endpoint (/):
- Endpoint: `http://127.0.0.1:8080/`
- Description: Allows users to add a URL for shortening. Users request a long URL through the body, and the service generates a unique eight-character code for it. If the same URL is added again, the service returns the same code.
- Method: POST
- Request:
- - url: The long URL to be shortened.
- Response:
- - Status Code: 201
- - Body: Eight-character code (/^[A-z0-9]{8}$/) representing the shortened URL.

2. Get Original URL Endpoint (/`<alias>`):
- Endpoint: `http://127.0.0.1:8080/<alias>`
- Description: Returns the original long URL associated with the provided eight-character code.
- Method: GET
- Parameters:
- - `<code>`: The eight-character code representing the shortened URL.
- Response:
- - Status Code: 200
- - Body: the original long URL.
