# Monarch
Monarch is a command line tool for migrating SQL databases.

## TODO

(Rename a column)
(Change a column's data type)
(Change column's default value)
(Add a constraint)
(Remove a constraint)

## Usage
Create a table
`monarch generate migration create table [name]`

Drop a table
`monarch generate migration drop table [name]`

Rename a table
`monarch generate migration rename table [name] [newname]`

Add columns to a table
`monarch generate migration add column [ [colname:coltype] ... ] --table [tablename]`
`monarch generate migration add column givenName:varchar familyName:varchar --table users`

Drop columns from a table
`monarch generate migration drop column [ [colname] ... ] --table [tablename]`
`monarch generate migration drop column givenName familyName --table users`
