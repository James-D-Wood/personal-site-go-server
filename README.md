# My Backend Service

This API serves as a common backend to a handful of my projects and as a template for a RESTful API service for any future projects. Here I will experiment with different patterns for app configuration, unit and integration testing, CICD pipelines, etc.

## Directory Structure

```
.
├── .github
│   └── workflows
|       └── go.yml          # GitHub actions
├── app
│   ├── main.go             # entrypoint
│   └── personal-site-api
|       ├── cfg             # loads environment
|       ├── database        # database interface
|       ├── middleware      # cors, auth, logging
|       ├── resources       # REST resources (handlers, models, routes)
|       └── server          # loads environment
|
├── db
│   └── migrations          # migrations
|
└──  internal               # utilities only needed in this app
```

## Resources

### Metadata

### Deprecated

#### Articles

These are entries to my blog/developer notebook and are comprised of article metadata and a markdown body.
Create and Update endpoints are auth protected, but Read endpoints are exposed.

Try:

```sh
curl https://api.jameswood.dev/api/v1/articles
```

This was quickly replaced by https://notebook.james.codes/, a Docusaurus site hosted on GitHub pages for ease of deploy and better site organization/navigation.

## Development

### Run the Service Locally

Generate DB from migration files:

```sh
$ migrate -path db/migrations -database "postgres://localhost:5432/personal_site?sslmode=disable" up
```

Create a conf.yaml file with the following:

```yaml
postgresql:
  host: <hostname>
  user: <username>
  database: personal_site
  port: 5432
  password: <password>
```

```shell
go generate # build values for meta endpoint
go run .
```

### SQL Migrations

For migrations, I use [migrate](https://github.com/golang-migrate/migrate). To create a new migration the steps are:

```bash
migrate create -ext sql -dir db/migrations -seq [migration_name]
# populate up and down files
migrate -path db/migrations -database "postgres://localhost:5432/personal_site?sslmode=disable" up
```
