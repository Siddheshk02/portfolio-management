# Portfolio Management Service

## Description
A simplified portfolio management service to manage user portfolios, track asset values, and provide basic analytics.

## Setup and Run
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/portfolio-management.git
    cd portfolio-management
    ```

3. Add and Update `.env` with your database credentials.
   ```
   User:<DB_USER>
   Password:<PASSWORD>
   ```

5. Run the application:
    ```sh
    go run cmd/main.go
    ```

## API Endpoints

### User Management
- `POST /register`: Register a new user.
- `POST /login`: Login a user.

### Portfolio Management
- `POST /portfolios`: Create a new portfolio.
- `GET /portfolios/{id}`: Get a portfolio by ID.
- `PUT /portfolios/{id}`: Update a portfolio by ID.
- `DELETE /portfolios/{id}`: Delete a portfolio by ID.

- `GET /portfolios/{id}/totalvalue`: Get Total Value of your Portfolio Assets
- `GET /portfolios/{id}/averagereturn`: Get Average Return on your Portfolio Assets

### Asset Management
- `POST /portfolios/{id}/assets`: Add an asset to a portfolio.
- `GET /portfolios/{id}/assets/{asset_id}`: Get an asset by ID.
- `PUT /portfolios/{id}/assets/{asset_id}`: Update an asset by ID.
- `DELETE /portfolios/{id}/assets/{asset_id}`: Delete an asset by ID.
