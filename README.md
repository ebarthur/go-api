# Go API

This is a simple API built using Go.

## Getting Started

To start the API, follow these steps:

1. Make sure you have Go installed on your machine. If not, you can download it from [here](https://golang.org/dl/).

2. Clone this repository to your local machine.

   ```bash
   git clone github.com/ebarthur/goapi.git
   ```

3. Navigate to the `cmd/api` directory in your terminal.

4. Run the following command to start the API:

   ```bash
   go run main.go
   ```

5. Once the API is running, you can access it at `http://localhost:8000`.

6. There is a middleware to get user coins. Find `internal/tools/mockdb.go` to get mock data.

7.  Attach username as a query param. eg: `http://localhost:8000/account/coins/?username=jason`

8. Add the AuthToken from the mock data to `Authorization` in your request header and make your get request.

9. Here is a link to the [Postman Documentation](https://www.postman.com/ebarthur/workspace/go-game-coin-api/request/34350550-e1395764-2ecc-4744-ba59-1f5cd45fe326?action=share&source=copy-link&creator=34350550&ctx=documentation)

10. Make sure the API server is on and running on your local machine else this will not work.

## API Endpoints

- `GET /account/coins`: Returns user coin balance

## Dependencies

- Run `go mod tidy` to install the dependencies called in the program

## Contributing

1. Fork the repository and clone it to your local machine.

2. Create a new branch for your feature or bug fix.

3. Make your changes and test them thoroughly.

4. Commit your changes with descriptive commit messages.

5. Push your changes to your fork.

6. Submit a pull request to the main repository.
