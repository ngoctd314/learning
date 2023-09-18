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

```bash
git log 
```

git log lists the commits made in that repository in reverse chronological order; that is, the most recent commits show up first. As you can see, this command lists each commit with its SHA-1 checksum, the author's name and email, the data written, and the commit message.

One of the more helpful options is -p or --patch, which shows the difference (the patch output) introduced in each commit. 

```bash
git log -p
```

You can also use a series of summarizing options with git log. For example, if you want to see some abbreviated stats 

```bash
git log --stat
```

### Undoing things

At any stage, you may want to undo something. Here, we'll review a few basic tools for undoing changes that you've made. Be careful, because you can't always undo some of these undos. This is one of the few areas in Git where you may lose some work if you do it wrong.

One of the common undos takes place when you commit too early and possibly forget to add some files or you mess up your commit message. If you want to redo that commit, make the additional changes you forgot, stage them, and commit again user the --amend option:

```sh
git commit -m 'Initial commit'
git add forgotten_file
git commit --amend
```

### Unstaging a Staged File

The next two sections demonstrate how to work with your staging area and working directory changes. The nice part is that the command you use to determine the state of those two area also reminds you how to undo changes to them. For example, let's say you've changed two files and want to commit them as two separate changes, but you accidentally type git add * and stage them both. How can you unstage one of the two?

```sh
git add *
git status
```

Right below the "Changes to be commited" text, it says use git reset HEAD <file> to unstage.

```sh
git reset HEAD CONTRIBUTING.md
```

### Unmodifying a Modified File

What if you realize that you don't want to keep your changes to the CONTRIBUTING.md file? How can you easily unmodified it - revert it back to what it looked like when you last committed.

```sh
git checkout -- <file>
```

### Undoing things with git restore

Git version 2.23.0 introduced a new command: git restore. It's basically an alternative to git reset which we jsut covered. From Git version 2.23.0 onwards, Git will use git restore instead of git reset

#### Unstaging a Staged File with git restore

```bash
git restore --staged CONTRIBUTING.md
```

The CONTRIBUTING.md file is modified but once again unstaged

#### Unmodifying a Modified File with git restore

```bash
git restore CONTRIBUTING.md
```

## Working with Remotes

To be able to collaborate on any Git project, you need to know how to manage your remote repositories. Remote repositories are versions of your project that hosted are on the internet.


### Showing your remotes

```bash
git remote
```

Remote is the default name Git gives to the server you cloned from.

You can also specify -v, which shows you the URLs that Git has stored for the shortname to be used when reading and writing to that remote.

```bash
git remote -v
```

### Fetching and pullling from your Remotes

```bash
git fetch <remote>
```

The command goes out to that remote project and pulls down all the data from that remote project that you don't have yet. After you do this, you should have references to all the branches from that remote, which you can merge in or inspect at any time. Git fetch command only downloads the data to your local repository - it doesn't automatically merge it with any of your work or modify what you're currently working on.

If you current branch is set up to track a remote branch, you can use the git pull command to automatically fetch and then merge that remote branch into your current branch. This may be an easier or more comfortable workflow for you.

Running git pull generally fetches data from the server you originally cloned from and automatically tries to merge it into the code you're currently working on.

From Git version 2.27 onward, git pull will give a warning if the pull.rebase variable is not set. Git will keep warning you until you set the variable.

If you want the default behavior of Git (fast-forward if possible, else create a merge commit): git config --global pull.rebase "false" 

If you want rebase when pulling: git config --global pull.rebase "true"

### Pushing to Your remotes

#### Renaming and removing remotes

## Tagging

Git has the ability to tag specific points in a repository's history as being important. Typically, people use this functionality to mark release points(v1.0, v2.0 and so on).

## Git Aliases

```bash
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```

```bash
git config --global alias.unstage 'reset HEAD --'
```

```bash
git unstage fileA
git reset HEAD -- fileA
```

## Summary

At this point, you can do all the basic local Git operations - creating or cloning a repository, making changes, staging and commiting those changes, and viewing the history of all the changes the repository has been through.
