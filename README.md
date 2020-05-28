# Monarch
Monarch is a command line tool for migrating PostgreSQL databases.

## TODO

* Add a foreign key constraint
* Add a constraint (general)
* Remove a constraint
* Change a column's data type
* Change column's default value

## Usage

### Managing Databases

#### Basic Syntax
`monarch db command [ [argument] ... ]`

#### Create a database
Syntax:
`monarch db create`

Example: create a database with the name specificed in the config file.
```
monarch db create
```

#### Copy a database
Syntax:
`monarch db copy [sourcedbname] [targetdb]`

Example: copy a database from sourcedb to targetdb.
```
monarch db copy sourcedb targetdb
```

Note: copying a database can take a while depending on how much data is in the source db.

#### Drop a database
Syntax:
`monarch db drop`

Example: drop a database with the name specificed in the config file.
```
monarch db drop
```

#### Rename a database
Syntax:
`monarch db rename [name] [newname]`

Example: rename a database from olddb to newdb.
```
monarch db rename olddb newdb
```

Note: if you change the name of your database, you should also change the database name specificed in your config file.

#### Reset a database
Syntax:
`monarch db reset`

Example: drop and create a database with the name specificed in the config file.
```
monarch db reset
```

### Generating Migrations

#### Basic Syntax
`monarch generate migration [subcommand] [ [additional subcommands] ... ] [argument] [ [additional args] ... ] `

#### Using Shortcut Aliases
The command `generate` can be shortened to `g` and `migration` can be shortened to `m`.

For example, the following statement will generate a migration to create a table called my_table.
```
monarch g m create table MyTable
```

The following examples will use the shortened form of the `generate` and `migration` commands.

#### Create a table
Syntax:
`monarch generate migration create table [name]`

Example: generate a migration to create a table named `users`
```
monarch g m create table users
```

#### Create a table with columns
Syntax:
`monarch generate migration create table [name] [ [colName:type] ... ]`

Example: generate a migration to create a table named `users` with columns `given_name` and `family_name`
```
monarch g m create table users givenName:varchar familyName:varchar
```

#### Drop a table
Syntax:
`monarch generate migration drop table [name]`

Example: generate a migration to drop a table named `users`
```
monarch g m drop table users
```

#### Rename a table
Syntax:
`monarch generate migration rename table [name] [newname]`

Example: generate a migration to rename a table from `old_name` to `new_name`
```
monarch g m rename table oldName newName
```

#### Add columns to an existing table
Syntax:
`monarch generate migration add column [tableName] [ [colName:type] ... ]`

Example: generate a migration to add a column `email` to table `users`
```
monarch g m add column users email:varchar
```

#### Drop columns from a table
Syntax:
`monarch generate migration drop column [tableName] [ [colName] ... ]`

Example: generate a migration to drop column `email` from table `users`
```
monarch g m drop column users email
```

#### Rename columns in a table
Syntax:
`monarch generate migration rename column [tableName] [ [colName:newName] ... ]`

Example: generate a migration to rename column `given_name` to `first_name`
```
monarch g m rename column users givenName:firstName
```

#### Create an index on an existing table
Syntax:
`monarch generate migration create index [tableName] [colName]`

Example: generate a migration to create a default index on table `cars` column `color`
```
monarch g m create index cars color
```
