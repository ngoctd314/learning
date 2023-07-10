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


