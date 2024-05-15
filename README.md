# RSS feed aggregator

This project provides an API for managing users, RSS feeds, feed follows, and posts. It allows users to subscribe to RSS feeds and retrieve posts from those feeds. Below is a list of the available endpoints and their functionalities.

## Launching the project

1. Copy `.env.example` file and rename it to `.env`. Fill in the values in {{}}
2. Install [goose](https://github.com/pressly/goose#install)
3. Cd into `./sql/schema/` and run `goose postgres "postgres://{{username}}:@localhost:5432/{{database}}" up` (same as in .env, but without `?sslmode=disable`)
4. Go to the project root and run `go build -o out && ./out`

# RSS Feed Aggregator

This project provides an API for managing users, RSS feeds, feed follows, and posts. It allows users to subscribe to RSS feeds and retrieve posts from those feeds. Below is a list of the available endpoints and their functionalities.

## Base URL

The base URL for the API is `http://localhost` plus the port specified in your `.env` file. Ensure your `.env` file includes a `PORT` variable.

For example, if your `.env` file specifies `PORT=3000`, the `BASE_URL` would be `http://localhost:3000`

## Endpoints

### Get Readiness Status

Check if the service is ready to accept requests.

**GET** `{{BASE_URL}}/readiness`

### User Endpoints

Manage user accounts.

#### Create a New User

Create a new user account.

**POST** `{{BASE_URL}}/users`

**Request Body:**

```json
{
  "name": "John Doe"
}
```

#### Get User by API Key

Retrieve user details using the API key.

**GET** `{{BASE_URL}}/users`

**Headers:**

- `Authorization: {{API_KEY}}`

### Feed Endpoints

Manage RSS feeds.

#### Post a New Feed

Add a new RSS feed and automatically subscribe the user to it.

**POST** `{{BASE_URL}}/feeds`

**Headers:**

- `Content-Type: application/json`
- `Authorization: {{API_KEY}}`

**Request Body:**

```json
{
  "name": "Lane's blog",
  "url": "https://www.wagslane.dev/index.xml"
}
```

#### Get All Feeds

Retrieve a list of all RSS feeds.

**GET** `{{BASE_URL}}/feeds`

### Feed Follow Endpoints

Manage feed follows for users.

#### Create a Feed Follow

Subscribe a user to a specific feed.

**POST** `{{BASE_URL}}/feed_follows`

**Headers:**

- `Content-Type: application/json`
- `Authorization: {{API_KEY}}`

**Request Body:**

```json
{
  "feed_id": "{{FEED_ID}}"
}
```

#### Get User's Feed Follows

Retrieve a list of feeds the user is following.

**GET** `{{BASE_URL}}/feed_follows`

**Headers:**

- `Authorization: {{API_KEY}}`

#### Delete a Feed Follow

Unsubscribe a user from a specific feed.

**DELETE** `{{BASE_URL}}/feed_follows/{{FEED_FOLLOW_ID}}`

**Headers:**

- `Authorization: {{API_KEY}}`

### Post Endpoints

Manage posts retrieved from feeds.

#### Get User's Posts

Retrieve posts from feeds the user is following. The default limit is 10 posts.

**GET** `{{BASE_URL}}/posts?limit=2`

**Headers:**

- `Authorization: {{API_KEY}}`

## Usage

To use these endpoints, replace `{{BASE_URL}}` with `http://localhost:{{PORT}}`, where `{{PORT}}` is the port number specified in your `.env` file. Replace placeholders such as `{{API_KEY}}`, `{{FEED_ID}}`, and `{{FEED_FOLLOW_ID}}` with the appropriate values.
