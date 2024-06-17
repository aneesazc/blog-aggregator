# Go Backend for Blog Aggregator

## Description

This is the backend service for the Blog Aggregator application, developed using Go. It provides APIs for managing users, feeds, and feed follows, and is deployed on Google Cloud Platform (GCP) using Cloud Run. A major feature of this project is the concurrent scraping of blogs from the followed feeds, ensuring that users receive the latest posts efficiently.

## Why?

The goal of this project is to create a robust, scalable, and efficient backend service for aggregating blog feeds. It aims to provide a seamless experience for users to register, manage their feeds, and follow updates from various blog sources. Concurrent scraping allows for timely updates and efficient handling of multiple feeds.

## Quick Start

### Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)

### Local Setup

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/blog-aggregator.git
   cd blog-aggregator
   ```

2. **Set Up Environment Variables**

   Create a `.env` file with the following content:

   ```env
   PORT="8080"
   DB_URL="your-database-url"
   ```

3. **Run the Application**

   ```sh
   go mod tidy
   go run main.go
   ```

4. **Build and Run with Docker**

   ```sh
   docker build -t blog-aggregator .
   docker run -p 8080:8080 blog-aggregator
   ```

### Deploy to GCP

1. **Build and Tag Docker Image**

   ```sh
   docker build -t gcr.io/YOUR_PROJECT_ID/blog-aggregator:v1 .
   ```

2. **Push Docker Image to GCR**

   ```sh
   docker push gcr.io/YOUR_PROJECT_ID/blog-aggregator:v1
   ```

3. **Deploy to Cloud Run**

   ```sh
   gcloud run deploy blog-aggregator \
       --image gcr.io/YOUR_PROJECT_ID/blog-aggregator:v1 \
       --platform managed \
       --region us-central1 \
       --allow-unauthenticated
   ```

## Usage

### Endpoints

- **Health Check**

  - `GET /v1/healthz`
  - Response: `{ "status": "ok" }`

- **User Registration**

  - `POST /v1/users`
  - Request Body: `{ "Name": "UserName" }`
  - Response: `{ "ID": "UUID", "CreatedAt": "Timestamp", "UpdatedAt": "Timestamp", "Name": "UserName", "Apikey": "ApiKey" }`

- **Create Feed**

  - `POST /v1/feeds`
  - Requires Authentication
  - Request Body: `{ "name": "Feed Name", "url": "Feed URL" }`
  - Response: `{ "ID": "UUID", "Name": "Feed Name", "Url": "Feed URL", "CreatedAt": "Timestamp", "UpdatedAt": "Timestamp", "UserID": "UUID" }`

- **Follow Feed**
  - `POST /v1/feed_follows`
  - Requires Authentication
  - Request Body: `{ "feed_id": "FeedID" }`
  - Response: `{ "id": "UUID", "feed_id": "FeedID", "user_id": "UserID", "created_at": "Timestamp", "updated_at": "Timestamp" }`

### Authentication

Authentication is handled using JWT tokens. The token is provided upon user registration and must be included in the `Authorization` header for authenticated endpoints.

## Contributing

1. **Fork the Repository**

   ```sh
   git clone https://github.com/yourusername/blog-aggregator.git
   cd blog-aggregator
   ```

2. **Create a Feature Branch**

   ```sh
   git checkout -b feature/your-feature-name
   ```

3. **Commit Your Changes**

   ```sh
   git commit -m "Description of your changes"
   ```

4. **Push to the Branch**

   ```sh
   git push origin feature/your-feature-name
   ```

5. **Create a Pull Request**

   Open a pull request on GitHub and provide a detailed description of your changes.
