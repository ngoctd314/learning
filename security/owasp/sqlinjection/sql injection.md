# SQL Injection

If you are vulnerable to SQL Injection, attackers can run arbitrary commands against your database.

SQL injection attacks target websites that use an underlying SQL database and construct data queries to the database in an insecure fashion.

## Anatomy of a SQL Injection Attack

SQL injection attacks occur when the web server insecurely constructs the SQL statement it passes to the database driver. This allows the attacker to pass arguments via the HTTP request that cause the dirver to perform actions other than those the developer inteads.

The attacker might cause the database driver to run additional commands on the database.

If your website is vulnerable to SQL injection, an attackerr can often run arbitrary SQL statements against your database, allowing them to bypass authentication; read, download, and delete data at will; or even inject malicious JavaScript into the pages rendered to your users.

## Mitigation

### 1. Use Parameterized Statements

To protect against SQL injection attacks, your code needs to construct SQL strings using bind parameters. Bind parameters are placeholder characters that the database driver will safely replace with some supplied inputs.

SQL injection attacks use "control characters" that have special meaning in SQL statements to "jump out" of the context and change the whole semantics of the SQL statement.

A securely constructed SQL statment using bind parameters should look like 

```sql
-- Using bind parameters to protect against SQL injection 
sql = "SELECT * FROM users WHERE email = ? AND encrypted_password = ?"
```

This code constructs the SQL query in parameterized form using ? as the bind parameter. The code then binds the input values for each parameter to the statement, asking the database driver to insert the parameter values into the SQL statement while securely handling any control characters. Because the datbasae dirver makes sure not to terminate the SQL statement early, this SELECT statement will safely return no users, and the attack should fail. Parameterized statements ensure that the database driver treats all control character (such as ', --, and;) as an input to the SQL statement, rather than as part of the SQL statement. If you're not sure whether your website is using parameterized statements, go check immediately! SQL injection is probably the biggest risk your website will face.

### 2. Use ORM

ORMs use bind parameters under the hood, they protect against injection attacks in most cases. However, most ORMs also have backdoors that allows the developer to write raw SQL if needed.

### 3. Use Defense in Depth

You should always secure your website with redundancies. It's not enough to check your code line by line for vulnerabilities. You need to consider and enforce security at every level of the stack, allowing failures at one level to be mitigated by other strategies.

Consider how you secure your home. The most important defense is install locks on all doors and windows, but it also helps to have a burgalar alarm, security cameras, and maybe a large bad-tempered dog, in order to cover all eventualities.

When it comes to preventing SQL injection, defense in depth means using bind parameters, but also taking additional steps to minimize the harm in case an attacker still finds a ways to successfully execute injection attacks.

#### 3.1. Principle of Least Privilege

An additional way to mitigate injection attacks is to follow the principle of least privilege, which demands that every process and application run only with the permissions it needs to perform its permitted functions, and no more. This means that if an attacker injects code into your we server and compromises a particular software component, the damage they can do is limited to the actions permissible by that particular software component.

If your web server talks to a database, make sure account it uses to log into the database has limited permissions on the data. Most websites need to run only SQL statements that fall under the subset of SQL called the data manipulation language (DML), which includes the SELECT, INSERT, UPDATE, and DELETE statements we discussed earlier.

Narrowing the web server privileges to the minimal DML set reduces the harm an attacker can do if they discover a code vulnerability.

#### 3.2. Blind and Nonblind SQL injection

Hackers distinguish between blind and nonblind SQL injection attacks. If you website's error message leaks sensitive information to the client, like the message Unique constraint violated: this email address already exists in users table, this is a nonblind SQL attack. In this scenario, the attackers gets immediate feedback on their attempts to compromise your system.

If you keep your error messages to the client more generic, like the messages Could not find this username and password or An unexpected error occured, this is a blind SQL attack. This scenario means the attacker is effectively operating in the dark and has less to work with. Websites vulnerable to nonblind injection attacks are much easier to compromise, so avoid leaking information in error message.