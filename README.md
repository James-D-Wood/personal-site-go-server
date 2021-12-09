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
├── internal                # utilities only needed in this app
|       
└──scripts                 # DB migrations
```

## Resources 

### Articles
These are entries to my blog/developer notebook and are comprised of article metadata and a markdown body.
Create and Update endpoints are auth protected, but Read endpoints are exposed. 

Try:
```sh
curl https://jameswood.dev/api/v1/articles
```