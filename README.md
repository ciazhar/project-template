# Setup
```bash

    # install db migration tool
    go get -v github.com/rubenv/sql-migrate/...
    
    # do database migration
    make migrate env=default
    
    # install dependency
    go mod tidy
    
    # running
    make run env=default
```