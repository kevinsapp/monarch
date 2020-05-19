# Monarch
Monarch is a command line tool for migrating SQL databases.

## TODO

* Add a foreign key constraint
* Add a constraint (general)
* Remove a constraint
* Change a column's data type
* Change column's default value

## Usage

### Generating Migrations
***Create a table***

Syntax
`monarch generate migration create table [name]`

Example: generate a migration to create a table named "users"
`monarch generate migration create table users`

***Create a table with columns***

Syntax
`monarch generate migration create table [name] [ [colName:type] ... ]`

Example: generate a migration to create a table named "users" with columns "given_name" and "family_name"
`monarch generate migration create table users`

Drop a table
`monarch generate migration drop table [name]`

Rename a table
`monarch generate migration rename table [name] [newname]`

Add columns to a table
`monarch generate migration add column [tableName] [ [colName:type] ... ]`
`monarch generate migration add column users givenName:varchar familyName:varchar`

Drop columns from a table
`monarch generate migration drop column [tableName] [ [colName] ... ]`
`monarch generate migration drop column users givenName familyName`

Rename columns in a table
`monarch generate migration rename column [tableName] [ [colName:newName] ... ]`
`monarch generate migration rename column users givenName:firstName familyName:lastName`
