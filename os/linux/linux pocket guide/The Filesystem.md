# The Filesystem

## Home Directories

User's personal files are often found in the directory `/home` (for ordinary user) or `/root` (for the superuser). Your home directory is typically `/home/<username>` (`/home/smith`, `/home/john`, etc).

## System Directores

A typical Linux system has tens of thousands of system directories.

Unless you're a system administrator, you'll rarely visit most sytem directories - but with a littel knowledge you can understand or guess their purposes. Their names often contain three parts

`/usr/local/share/emacs`

`/usr/local`: Directory scope
`/share`: category
`/emacs`: application

**Directory path part 1: category**

A category tells you the types of files found in a directory. For example, if the category is bin, you can be reasonably assured that the directory contains programs. Some common categories are as follows:

**Categories for programs**

bin: Programs (usually binary files)
sbin: Programs (usually binary files) intended to be run by the superuser
lib: Libraries of code used by programs

**Categories for documentation**
doc: Documentation
info: Documentation files for emacs's built-in help system
man: Documentation files (manual pages) displayed by the man program; the files are often compressed and are sprinkled with typesetting commands for man to interpret
share: Program-specific fields such as examples and installation instructions

## File Protections

A Linux system may have many users with login accounts. To maintain privary and security, must users can access oly some files one the system, not all.

**Who has permission?**

Every file and directory has an owner who has permission to do anything with it. Typically, the user who created a file is its owner, but ownership can be changed by the superuser.

Additionally, a predefined group of users may have permission to access a file. Groups are defined by the system administrator.

Finally, a file or directory can be opened to all users with login accounts on the system. You'll also see this set of users called the world or simply other.

**What kind of permission is granted?**

File owners, groups, and the world may each have permission to read, write (modify), and execute (run) particular files.

-rwxr-x---

|Position|Meaning|
|-|-|
|1|File type: - = file, d = directory, l = symbolic link, p = named pipe, c = character device, b = block device|
|2-4|Read,write, and execute permissions for the file's owner|
|5-7|Read, write, and execute permissions for the files group|
|8-10|Read, write, and execute permissions for all other users|

So -rwxr-x--- means a file that can be read, written and executed by the owner, read and executed by the group, and not accessed at all by other users.