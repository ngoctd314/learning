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

Follow these steps

1. Do some work on a website
2. Create a branch for a new user story you're working on
3. Do some work in that branch


Now you get call that there is an issue with the website, and you need to fix it immediately. With Git, you don't have to deploy your fix along with the iss53 changes you've made, and you don't have to put a log of effort into reverting those changes before you can work on applying your fix to what is in production. All you have to do is switch back to your master branch.

However, before you do that, note that if your working directory or staging area has uncommited changes that conflict with the branch you're checking out, Git won't let you switch branches. It's best to have a clean working state when you switch branches. It's best to have a clean working state when you switch branches. There are ways to get around this (namely, stashing and commit amending) that we'll cover later on. For now, let's assume you've commited all your changes, so you can switch back to your master branch:

```go
git checkout master
```

At this point, your project working directory is exactly the way it was before you started working on issue #53, and you can concentrate on your hotfix. This is an important point to remember: when you switch branches, Git reset your working directory to look like it did the last time you commited on that branch.

Switched to a new branch 'hotfix'
```bash
git checkout -b hotfix
```

```bash
vim index.html
git commit -a -m 'Fix broken email address'
```

[hotfix 1fb7853] Fix broken email address
1 file changed, 2 insertion(+)

You can run your tests, make sure the hotfix is what you want, and finally merge the hotfix branch back into your master branch to deploy to production.

```bash
git checkout master
git merge hotfix
```

Updating f42c576 .. 3a0874c
Fast-forward
indexhtml | 2++
1 file changed, 2 insertion (+)

You'll notice the phrase "fast-forward" in that merge. Because the commit C4 pointed to by the branch hotfix you merged is was directly ahead of the commit C2 you're on, Git simply moves the pointer forward. To phrase that another way, when you try to merge one commit with a commit that can be reached by following the first commits history, Git simplies things by moving the pointer forward because there is no divergent work to merge together - this is called fast-forward.

After your super-important fix is deployed, you're ready to switch back to the work you were doing before you were interrupted. However, first you'll delete the hotfix branch, because you no longer need it - the master branch points at the same place. You can delete it with the -d option to git branch:

```bash
git branch -d hotfix
```

```bash
git checkout iss53 
```
