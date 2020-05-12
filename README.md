# Monarch
Monarch is a command line tool for migrating SQL databases.

## Usage
Create a table
`monarch generate migration create table [name]`

Drop a table
`monarch generate migration drop table [name]`

Rename a table
`monarch generate migration rename table [name] [newname]`

Add column to table
`monarch generate migration add column [colname:coltype] [...] --table [tablename]`
`monarch generate migration add column givenName:varchar familyName:varchar --table users`

## SQL

CREATE TABLE
DROP TABLE
(Rename a table)
(Add a column)
(Drop a column)
(Add a constraint)
(Remove a constraint)
(Change a column's data type)
(Raname a column)
(Change column's default value)
