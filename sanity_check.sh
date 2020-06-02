#!/usr/bin/env bash

# Run package tests.
go test ./...
if [ $? -eq 1 ]
then
  echo "Package tests failed; Exiting." >&2
  exit 1
fi

# Compile and install program.
go install .

# Remove any leftover migrations and reset the database.
rm -rf migrations
monarch db reset

# Generate migrations
monarch g m create table users givenName:varchar familyName:varchar
monarch g m add column users email:varchar phone:varchar
monarch g m rename column users givenName:firstName familyName:lastName
monarch g m drop column users phone
monarch g m rename table users people
monarch g m create table cars
monarch g m add column cars make:varchar modelYear:smallint modelName:text color:text
monarch g m recast column cars modelName:varchar color:varchar
monarch g m create index cars make
monarch g m drop table people
monarch g m drop table cars

# Migrate schemas.
monarch db migrate

# Do cleanup.
monarch db drop
rm -rf migrations
