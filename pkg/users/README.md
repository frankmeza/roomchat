# `Users` package

## Overview of package

### `api.go`

`useConnectionsAPI` is a factory function that returns a Users API that is used inline for interacting with this module via cross-application function calls.  

The API methods do one of a few things:

- take care of all CRUD operations for `Users`, 
- and/or wrap private persistence level functions in `users/db.go`.

### `db.go`

The functions in this file are meant to stay private. They can still be accessed by other modules via a wrapper API method for that purpose.

The DB methods perform the actual CRUD operations with the persistence level, Postgres, via `gorm`.  

Generally, the functions:

- create a db connection,
- possibly construct a query,
- make the query,
- returning a result or error for each of these steps.

### `handlers.go`

The functions here are the same ones called in `index.go`. They are in charge of handling any incoming payload and translating it into a Go struct of some sort, and not necessarily an application model.

### `index.go`

`AddConnectionActions` adds the routes involved for interacting with this module via http.

### `types.go`

In here are the application models used in this module. Also in here are any necessary methods needed to inform Postgres about the JSON structure of any nested Go structs.