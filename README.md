# monarch
A command line tool for migrating databases

## Usage
Create a table
monarch generate migration create table [name]
monarch generate migration create --table [name]

Drop a table
monarch generate migration drop table [name]
monarch generate migration drop --table [name]

Rename a table
monarch generate migration rename table [name] [newname]
monarch generate migration rename --table [name] [newname]

Add column to table
monarch generate migration add column [colname] --table [tablename]
monarch generate migration add --column [colname] --table [tablename]

ALTER TABLE products ADD COLUMN description text;


SQL

CREATE TABLE
DROP TABLE
(Rename a table)
(Add a column)
(Remove a column)
(Add a constraint)
(Remove a constraint)
(Change column's default value)
(Change a column's data type)
(Raname a column)
