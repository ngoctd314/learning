# Database Design and Implementation

## 10. How to design a Database

## 11. How to create Databases, Tables, and Indexes

### 11.1. How to work with character sets and collations

When a column is defined with a string type such as CHAR or VARCHAR, MySQL stores a numeric value for each character. Then, it uses a character set to map the numeric values to the characters of the string.

The advantage of the utf8 character set is that it provides for all characters specified by the Unicode character set. The disadvantage of the utf8 character set is that it can use up to three bytes per character. This force MySQL to reserve three bytes per character for each character.

Every character set has a corresponding collation that determines how the characters within the set are stored.

**Two commonly used character sets**

|Name|Description|
|-|-|
|latin1|The latin1 character set uses one byte per character to provide for most characters in Western European languages|
|utf8|The utf8 character set uses one to three bytes per character to provide for all characters specified by the Unicode character set.|

### 11.2. How to work with storage engines

A storage engine determines how MySQL stores data and which database features are available to you. 

Two most commonly used storage engines: InnoDB and MyISAM.

Prior to MySQL 5.5, the MyISAM engine was the default storage engine. This engine supports some features that aren't supported by InnoDB tables. This engine supports some features that aren't supported by InnoDB tables, including full-text searches and spatial data types. However, the MyISAM engine doesn't support foreign keys, an important feature for maintaining referential integrity.

**Two commonly used storage engines**

|Name|Description|
|-|-|
|InnoDB|The default storage engine for MySQL 5.5 and later. This engine supports foreign keys and transactions|
|MyISAM|The default storage engine prior to MySQL 5.5. This engine support full-text searching and the spatial data types|

**Description**
- The storage engine determines how MySQL stores data and which database features are available to you.
- You can use multiple storage engines on the same server and within the same database

## 12. How to create views

As you've seen thoughout this book, SELECT queries can be complicated, particularly if they use multiple joins, subqueries, or complex functions. Because of that, you may want to save the queries you use regularly. One way to do that is to store the statement in a script. Another way is to create a view.

Unlike scripts, which are stored in files, views are stored as part of the database. As a result, they can be used by SQL programmers and by custom applications that have access to the database.

### 12.1. An introduction to Views

A view is a SELECT statement that's stored in the database as a database object. To create a view, you can use a CREATE VIEW statement. You can think of a view as virtual table that consits only of the rows and columns specified in its CREATE VIEW statement.

**Description**

- A view consits of a SELECT statement that's stored as an object in the database.
- Although a view behaves like a virtual table, it doesn't store any data. Instead, a view always refers back to its base tables.
- A view can also be referred to as a viewed table because it provides a view to the underlying base tables.

**Benefits of using views**

You can use views to limit the exposure of the tables in your database to external users and applications. 

You can also use views to restrict access to database. To do that, you include just the columns and rows you want a user or an application to have access to in the views. Then, you let the user or application access the data 

You can use views to hide the complexity of a SELECT statement.

Finally, when you create a view, you can allow data in the base table to be updated through the view. To do that, you use INSERT, UPDATE or DELETE statements to work with the view.

### 12.2. How to work with Views