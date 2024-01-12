# RainbowUsers
This Microservice is for communication with the client version and **user login**.

## This version includes the following sections:
```
- Making a MakeFile.
- Continuous integration (ci.yml) .
- Writing the required query for the Postgres database.
- The sqlc workload automatically generates Golang language codes to communicate with the database (currently, Transaction variables have not been added).
- Writing automatic tests for methods related to CRUD operations on the Postgres database.
Soon, other stables including RestFul API and Transaction will be added to the microservice.
```

# sqlc

For further explanation about [sqlc](https://github.com/sqlc-dev/sqlc ) please read [here](https://github.com/sqlc-dev/sqlc ). 

## Usage

```yaml
version: '2'
plugins:
- name: golang
  wasm:
    url: https://downloads.sqlc.dev/plugin/sqlc-gen-go_1.2.0.wasm
    sha256: 965d73d22711eee3a210565e66f918b8cb831c5f5b612e680642a4a785dd1ca1
sql:
- schema: schema.sql
  queries: query.sql
  engine: postgresql
  codegen:
  - plugin: golang
    out: db
    options:
      package: db
      sql_package: pgx/v5
```



# Note!
The `db/query/account.sql` file contains generated Go files from SQLC. It should not be edited manually.</
To generate the Go files again run the command below in your terminal:
````
In terminal:
make sqlc
````



## These services are supposed to be included in an Enterprise and Scalability product, and you can find out about the added features by following me.


# Thank you
### Fakhredin Gholamizadeh
### fakhredin.gholamizadeh@gmail.com