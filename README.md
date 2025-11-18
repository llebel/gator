# Gator

A CLI-based RSS feed aggregator that allows you to follow and read blog posts from multiple RSS feeds.

## Prerequisites

Before running Gator, you need to have the following installed:

- **Go** (version 1.24 or higher) - [Download and install Go](https://go.dev/doc/install)
- **PostgreSQL** - [Download and install PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the Gator CLI using `go install`:

```bash
go install github.com/llebel/gator@latest
```

This will download, compile, and install the `gator` binary to your `$GOPATH/bin` directory. Make sure `$GOPATH/bin` is in your system's `PATH`.

## Configuration

Gator requires a configuration file to connect to your PostgreSQL database. Create a `.gatorconfig.json` file in your home directory:

```bash
touch ~/.gatorconfig.json
```

Add the following content to the file:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` and `password` with your PostgreSQL credentials. The `current_user_name` field will be automatically set when you register or login.

## Database Setup

Before running Gator, you need to create the database and run migrations:

```bash
# Create the database
createdb gator

# Run migrations (if using a migration tool like goose)
# Or execute your SQL schema files directly
psql -d gator -f path/to/schema.sql
```

## Usage

### Available Commands

#### User Management

- **register** - Register a new user
  ```bash
  gator register <username>
  ```

- **login** - Login as an existing user
  ```bash
  gator login <username>
  ```

- **users** - List all registered users
  ```bash
  gator users
  ```

#### Feed Management

- **addfeed** - Add a new RSS feed (requires login)
  ```bash
  gator addfeed <feed_name> <feed_url>
  ```

- **feeds** - List all feeds in the system
  ```bash
  gator feeds
  ```

#### Following Feeds

- **follow** - Follow a feed (requires login)
  ```bash
  gator follow <feed_url>
  ```

- **following** - List feeds you're currently following (requires login)
  ```bash
  gator following
  ```

- **unfollow** - Unfollow a feed (requires login)
  ```bash
  gator unfollow <feed_url>
  ```

#### Reading Posts

- **browse** - Browse posts from feeds you follow (requires login)
  ```bash
  gator browse [limit]
  ```
  Optional `limit` parameter specifies how many posts to display (default: 2)

- **agg** - Aggregate/fetch new posts from all feeds
  ```bash
  gator agg <time_between_requests>
  ```
  Example: `gator agg 1m` fetches feeds every 1 minute

#### Utility

- **reset** - Reset the database (deletes all users and feeds)
  ```bash
  gator reset
  ```

## Example Workflow

```bash
# Register a new user
gator register john

# Add some RSS feeds
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
gator addfeed "Go Blog" https://go.dev/blog/feed.atom

# Follow a feed
gator follow https://blog.boot.dev/index.xml

# Start aggregating posts (in a separate terminal)
gator agg 5m

# Browse your feed
gator browse 10
```

## Development

This is a Boot.dev Blog Aggregator exercise project.
