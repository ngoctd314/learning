# Git basics

## Recording changes to the Repository

Each file in your working directory can be in one of two states: tracked or untracked.

Tracked files are files that were in the last snapshot, as well as any newly staged files; they can be unmodified, modified, or staged. In short, tracked files are files that Git knows about.

Untracked files are everything else - any files in your working directory that were not in your last snapshot and are not in your staging area. When you first clone a repository, all of your files will be tracked and unmodified because Git just checked them out and you haven't edited anything.

To view untracked file, use git status

### Tracking new files

In order to begin tracking a new file, you use the command git add. To begin tracking the README file, you can run this:

```bash
git add README
```

### Staging modified files

### Short status

```bash
git status -s
 M README
MM Makefile
A  lib/git.txt
M  account.txt
?? LICENSE.txt
```

There are two columns to the output - the left-hand column indicates the status of the staging area and the right-hand column indicates the status of the working tree.

### Ignoring files

.gitignore

### Viewing your staged and unstaged changes

If the git status command is too vauge for you - you want to know exactly what you changed, not just which files were changed - you can use git diff command. 

git diff compares what is in your working directory with what is in your staging area. The result tells you the changes you've made that you haven't yet staged.

If you want to see what you've staged what will go into your next commit, you can use git diff --staged. This command compares your staged changes to your last commit.

### Commiting your changes

You've created your first commit! You can see that the commit has given you some output about itself: which branch you committed, what SHA-1 checksum the commit has, how many files changed and statistics about lines added and removed in the commit.

The commit records the snapshot you set up in your staging area. Anything you didn't stage is still there modified; you can do another commit to add it to your history. Every time you perform a commit, you're recording a snapshot of your project that you can revert to or compare to later.

### Staging the Staging Area

Adding -a option to the git commit command makes Git automatically stage every file that is already tracked before doing the commit, letting you skip the git add part:

You don't have to run git add on the CONTRIBUTING.md file this case before you commit. That's because the -a flag includes all changed files. This is convenient, but be careful; sometimes this flag will cause you to include unwanted changes. 

### Removing files

To remove a file from Git, you have to remove it from your tracked files (more accurately, remove it from your staging area) and then commit. The git rm command does that, and also removes the file from your working directory so you don't see it as an untracked file the next time around.

### Moving files

If you want to rename a file in Git, you can run something like

```bash
git mv file_from file_to
```

```bash
mv file_from file_to
git rm file_from
git add file_to
```

## Viewing the commit history
