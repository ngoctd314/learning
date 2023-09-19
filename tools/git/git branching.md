# Git branching

## Branches in a Nutshell

Nearly every VCS has some form of branching support. Branching means you diverge from the main line of development and continue to do work without messing with that main line. In many VCS tools, this is a somewhat expensive process, often requiring you to create a new copy of your source code directory, which can take a long time for large projects.

Git doesn't store data as a series of changesets of differences, but instead as a series of snapshots.

When you make a commit, Git stores a commit object that contains a pointer to the snapshots of the content you staged. This object's also contains the author's name and email address, the message that you typed, and pointers to the commit or commits that directly came before this commit (its parent or parents)

## Creating a New Branch

What happens when you create a new branch? Well, doing so creates a new pointer for you to move around.

```bash
git branch testing
```
This creates a new pointer to the same commit you're currently on.

How does Git know what branch you're currently on? It keeps a special pointer called HEAD. In Git, this is a pointer to the local branch you're currently on. The git branch command only created a new branch - it didn't switch to that branch. 

### Switching branches

To switch to an existing branch, you run the git checkout command. Let's switch to the new testing branch:

```bash
git checkout testing
```

The HEAD branch moves forward when a commit is made

This is interesting, because now your testing branch has move forward, but your master branch still points to the commit you were on when you ran git checkout to switch branches. Let's switch back to the master branch:

```bash
git checkout master 
```
That command did two things. It moved the HEAD pointer back to point to the master branch, and it reverted the files in your working directory back to the snapshot that master points to.

**Creating a new branch and switching to it at the same time**

It's typical to create a new branch and want to switch to that new branch at the same time - this can be done in one operation with git checkout -b

From git version 2.23 onwards you can use git switch instead of git checkout to:
- Switch to an existing branch: git switch testing-branch
- Create a new branh and switch to it: git switch -c new-branch. The -c flag stands for create, you can also use the full flag --create
- Return to your previously checked out branch: git switch -

## Git branching - Basic branching and merging

