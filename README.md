# Secondhand Marketplace API

RESTful API for a secondhand marketplace. Users can manage items, orders, users, and categories. The API supports authentication and authorization for different roles (users and admins).

## Features

- **User Management:**
  - Users can only view and update their own accounts.
  - Admins can CRUD all user accounts.

- **Item Management:**
  - Users can view, create, update, and delete their own items.
  - Admins can CRUD all items.

- **Order Management:**
  - Users can view, create, update, and delete their own orders.
  - Admins can CRUD all orders.

- **Category Management:**
  - Users can view all categories.
  - Admins can CRUD all categories.

## How to Run
### Prerequisites

- Golang
- MySQL

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/secondhand-marketplace-api.git
   cd secondhand-marketplace-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the `.env` file with your database and JWT settings:
   ```env
   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=secondhand_marketplace
   JWT_SECRET_KEY=your_jwt_secret
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## API Documentation

Access the documentation for each endpoints here:
> https://documenter.getpostman.com/view/40551639/2sAYJ6CKtx
