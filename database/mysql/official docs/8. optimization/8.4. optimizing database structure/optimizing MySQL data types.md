# Optimizing MySQL Data Types

## Optimizing for Numeric Data

- For unique IDs or other values that can be represented as either strings or numbers, prefer numberic columns to string columns. Since large numeric values can be stored in fewer bytes than the corresponding strings, it is faster and takes less memory to transfer and compare them.
- If you are using numeric data, it is faster in many cases to access information from a database.
