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

### Articles
These are entries to my blog/developer notebook and are comprised of article metadata and a markdown body.
Create and Update endpoints are auth protected, but Read endpoints are exposed. 

Try:
```sh
curl https://api.jameswood.dev/api/v1/articles
```

## Development

### SQL Migrations
For migrations, I use [migrate](https://github.com/golang-migrate/migrate). To create a new migration the steps are:
```bash
migrate create -ext sql -dir db/migrations -seq [migration_name]
# populate up and down files
migrate -path db/migrations -database "postgres://localhost:5432/personal_site?sslmode=disable" up
``` 