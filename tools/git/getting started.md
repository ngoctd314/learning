# Getting started

## About Version Control

Version control is a system that records changes to a file or set of files over time so that you can recall specific versions later. VSC allows you to revert selected files back to a previous state, revert the entire project back to a previous state, compare changes over time, see who last modified something that might be causing a problem, who introduced an issue and when, and more.

### Local version control systems

Many people's version-control method of choice is to copy files into another directory (perhaps a timestamped directory if they're clever). This approach is very common because it is so simple, but it is also incredibly error prone. It is easy to forget which directory you're in and accidentally write to the wrong file or copy over files you don't mean to.

To deal with this issue, programmers long ago developed local VCSs that had a simple database that kept all the changes to files under revision control.

### Centralized version control systems

The next major issue that people encouter is that they need to collaborate with developers on other systems. Centralized Version Control Systems have a single server that contains all the versioned files, and a number of clients that check out files from that central place.

However this setup also has some serious downsides. The most obvious is the single point of failure that the centralized server represents.

### Distributed version control systems

This is where distributed version control systems step in. In a DVCS, clients don't just check out the latest snapshot of the files; rather, they fully mirror the repository, including its full history. 

## What is Git

### Snapshots, Not Differences

### Nearly every operation is local

### Git has integrity

### Git generally only adds data

### The three states

Git has three main state that your files can reside in: modified, staged, and committed:

- Modified means that you have changed the file but have not commited it to your database yet.
- Staged means that you have marked a modified file in its current version to go into your next commit snapshot.
- Committed means that the data is safely stored in your local database.

This leads us to the three main sections of a Git project: the working tree, the staging area, and the Git directory.

The working tree is a single checkout of one version of the project. These files are pulled out of the compressed database in the Git directory and placed on disk for you to use or modify.

The staging area is a file, generally contained in your Git directory, that stores information about what will go into your next commit.

The Git directory is where Git stores the metadata and object database for your project. This is the most important part of Git, and it is what is copied when you clone a repository from another computer.

The basic Git workflow goes something like this:

1. You modify files in your working tree

2. You selectively stage just those changes you want to be part of your next commit, which adds only those changes to the staging area.

3. You do a commit, which takes the files as they are in the staging and stores that snapshot permanently to your Git repository.

If a particular version of a file is in the Git directory, it's considered commited. If it has been modified and was added to the staging area, it is staged. And if it was changed since it was checkout but has either take advantage of them or skip the staged part entirely.

## First time git setup

Git comes with a tool called git config that lets you get and set configuration variables that control all aspects of how Git looks and operates. These variables can be stored in three different places:

1. [path]/etc/gitconfig

2. ~/.gitconfig or ~.config/git/config file: Values specific personally to you, the user. You can make Git read and write to this file specifically by passing --global option, and this affects all of the repository you work on your system.

3. config file in Git directory (that is .git/config) of whatever repository you're currently using: You can force Git to read from and write to this file with the --local option, but that is the fact the default. Unsurprisingly, you need to be located somewhere in a Git repository for this option to work properly.

### Your identity

The first thing you should do when you install Git is to set your username and email address. This is important because every Git commit uses this information, and it's immutably baked into the commits you start creating.

```bash
git config --global user.name "author"
git config --global user.email "author@gmail.com"
```

### Your editor

```bash
git config --global core.editor nvim
```

### Your default branch name

By default Git will create a branch called master when you create a new repository with git init. From Git version 2.28 onwards,  you can set a different name for the initial branch.

```bash
git config --global init.defaultBranch main
```

### Checking your settings

```bash
git config --list
```
