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
`monarch generate migration add column [tableName] [ [colName:type] ... ]`
`monarch generate migration add column users givenName:varchar familyName:varchar`

Drop columns from a table
`monarch generate migration drop column [tableName] [ [colName] ... ]`
`monarch generate migration drop column users givenName familyName`

Rename columns in a table
`monarch generate migration rename column [tableName] [ [colName:newName] ... ]`
`monarch generate migration rename column users givenName:firstName familyName:lastName`

## TODO

Change column's data type
