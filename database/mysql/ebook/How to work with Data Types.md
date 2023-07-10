# How to work with data type

## 8.1. The data types

### 8.1.1. Overview
Five categories: **character**, **numeric**, **date and time**, **large object (LOB)** for storing images, sound, video, **spatial**.

### 8.1.2. The character types
CHAR and VARCHAR.

The character types

|Type|Bytes|Description|
|-|-|-|
|CHAR(M)|Mx3|Fixed-length strings of character data where M is the number of characters, between 0 and 255.|
|VARCHAR(M)|L+1|Variable-length strings of character data where M is the maximum number of characters, between 0 and 255.|

How the character types work with UTF-8

|Data type|Original value|Value stored|Bytes used|
|-|-|-|-|
|CHAR(2)|'CA'|'CA'|6|
|CHAR(10)|'CA'|'CA        '|30|
|VARCHAR(10)|'CA'|'CA'|3|

**Description**

- The CHAR type is used for fixed-length strings. A column with this type uses the same amount of storage for each value regardless of the actual length of the string.
- The VARCHAR type is used for variable-length strings. A column with this type uses a varying amount of storage for each value depending on the length of the string.
 
Although you typically store numeric values using numeric types, the character types by be a better choice for some numeric values. For example, you typically store zip codes, telephone numbers, and social security numbers in character columns even if they contain only numbers. That's because their values aren't used in numeric operations.


### 8.1.3. The integer types
The integer types

|Type|Bytes|Value ranges|
|-|-|-|
|BIGINT|8|-9BB -> 9BB|
|INT|4|-2B -> 2B|
|MEDIUMINT|3|-8M -> 8M|
|SMALLINT|2|-32K -> 32K|
|TYNYINT|1|-128 -> 128|

### 8.1.4. The fixed-point and floating-point types

You can use the DECIMAL type to store fixed-point numbers, which are numbers that have a fixed number of digits to the right of the decimal point.

In contrast to the DECIMAL type, the DOUBLE and FLOAT types store floating-point numbers. These data types provide for very large and very small numbers, but with a limited number of significant digits.

Because the precision of the integer types and the DECIMAL type is exact, these data types are considered exact numeric types. In constrast, the DOUBLE and FLOAT types are considered approximate numeric types because the may not represent a value exactly. 

**The fixed-point type**
|Type|Bytes|Description|
|-|-|-|
|DECIMAL(M, D)|Vary|Fixed-precision numbers where M specifies the maximum number of total digits and D specifies the number of digits to the right of the decimal. M can range from 1 to 65. D can range from 0 to 30 but can't be larger than M. The default is 0.|

**The floating-point types**
|Type|Bytes|Description|
|-|-|-|
|DOUBLE|8|Double-precision floating-point numbers from -1.79x10^308 to 1.79x10^308|
|FLOAT|4|Single-precision floating-point numbers from -3.4x10^38 to 3.4x10^38|

|Data type|Original value|Value stored|Bytes used|
|-|-|-|-|
|DECIMAL(9,2)|1.2|1.20|5|
|DOUBLE|1234.90|1234.90|8|
|FLOAT|1234.90|1234.90|4|

**Description**

- The DECIMAL type is considered an exact numeric type because its precision is exact.
- The DOUBLE and FLOAT types store floating-point numbers. These data types are considered approximate numeric data types because they may not represent a value exactly.

### 8.1.5. The date and time types

You can use DATE type to store a date without a time. You can use the TIME type to store a time without a date. And you can use either the DATETIME or TIMESTAMP types to store both a date and a time.

The problem with the TIMESTAMP type is that it can only store dates up to the year 2038. TIMESTAMP use 4 bytes, DATETIME use 8 bytes.

**The date and time types**
|Type|Bytes|Description|
|-|-|-|
|DATE|3|YYYY-MM-DD|
|TIME|3|HH:MM:SS|
|DATETIME|8|YYYY-MM-DD HH:MM:SS|
|TIMESTAMP|4|1970 - 2037|
|YEAR[(2|4)]|1|(19)70 - (20)69|

**Description**
- A column of TIMESTAMP type is automatically updated to the current date and time when a row is inserted or updated. If a table has multiple TIMESTAMP columns, only the first one is updated automatically.
- The TIMESTAMP type can only store dates up to the year 2038. This is known as the year 2038 problem, the Y2K38 problem, and the Unix Millennium bug. To fix this problem, use the DATETIME type instead of the TIMESTAMP type and update the value manually as needed.

### 8.1.6. The ENUM and SET types

The ENUM and SET types can be considered character data types since they allow you to restrict the values for a column to a limited set of strings. MySQL stores these values as integers, which reduces the number of bytes needed to store each string.


### 8.1.7. The large object types

## 8.2. How to convert data
