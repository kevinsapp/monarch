#!/usr/bin/env bash

go test ./...
if [ $? -eq 1 ]
then
  echo "Unit tests failed; Exiting." >&2
  exit 1
fi

go install .

rm -rf migrations

monarch db reset
monarch g m create table users givenName:varchar familyName:varchar
monarch g m add column users email:varchar phone:varchar
monarch g m rename column users givenName:firstName familyName:lastName
monarch g m drop column users phone
monarch g m rename table users people
monarch g m create table cars
monarch g m add column cars color:varchar
monarch g m create index cars color
monarch g m drop table people
monarch g m drop table cars
monarch db migrate
monarch db drop

rm -rf migrations
