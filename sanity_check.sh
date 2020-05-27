#!/usr/bin/env bash

go test ./...

go install .

rm -rf migrations

monarch db reset
monarch g m create table users
monarch g m add column users givenName:varchar familyName:varchar
monarch g m rename column users givenName:firstName familyName:lastName
monarch g m drop column users lastName
monarch g m rename table users people
monarch g m drop table people
monarch db migrate
monarch db drop

rm -rf migrations
