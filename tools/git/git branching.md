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

## Branch management

The useful --merged and --no-merged options can filter this list to branches that you have or have not yet merged into the branch you're currently on. To see which branches are already merged into the branch you're on, you can run git branch --merged.

```bash
git branch --merged
iss53
*master
```

Try to delete branch with git branch -d

### Changing a branch name

```bash
git branch --move bad-branch-name corrected-branch-name
```

This replaces your bad-branch-name with corrected-branch-name, but this change is only local for now. To let others see the corrected branch on the remote, push it: 

```bash
git push --set-upstream origin corrected-branch-name
```

Notice that you're on the branch corrected-branch-name and it's available on the remote. However, the branch will the bad name is also still present but you can delete it by executing the following command:

```bash
git push origin --delete bad-branch-name
```

```bash
git branch --move master main
git push --set-upstream origin main
git branch --all
git push origin --delete master
```

Now you have a few more tasks in front of you to complete the transition: 
- Any projects that depend on this one will need to update their code and/or configuration
- Update any test-runner configuration files
- Adjust build and release scripts
- Redirect settings on your repo host for things like the repos's default branch, merge rules, and other things that match branch names.
- ...

## Branching Workflows

### Long-Running Branches

Git uses a simple three-way merge, merging from one branch into another multiple times over a long period is generally easy to do.

### Topic branches

Topic branches, however, are useful in projects of any size. A topic branch is short-lived branch that you create and use for a single particular feature or related work. This  is something you've likely never done with a VCS before because it's generally too expensive to create and merge branches. But in Git it's common to create, work on, merge, and delete branches several times a day. 

## Remote branches

Remote references are references(pointer) in your remote repositories, including branches, tags and so on. You can get a full list or remote references explicitly.

Nevertheless, a more common way is to take advantage of remote-tracking branches.

Remote-tracking branches are references to the state of remote branches.

**origin is not special**

Just like the branch name "master" does not have any special meaning in Git, neither does "origin". While "master" is the default name for a starting branch when you run git init which is the only reason it's widely used, "origin" is the default name for a remote when you run git clone. It you run git clone -o booyah instead, then you will have booyah/master as your default remote branch.

To synchronize your work with a given remote, you run a git fetch <remote> command (in our case, git fetch origin). This command looks up which server "origin" is, fetches any data from it that you don't yet have, and updates your local database, moving your origin/master pointer to its new, more up-to-date position. 

### Pushing

When you want to share a branch with the world, you need to push it up to a remote to which you have write access. Your local branches aren't automatically synchronized to the remotes you write to - you have to explicitly push the branches you want to share. 

### Tracking Branches

### Pulling

While the git fetch command will fetch all the changes on the server that you don't have yet, it will not modify your working directory at all. It will simply get the data for you and let you merge it yourself. However, there is a command called git pull which is essentially a git fetch immediately followed by a git merge in most cases. Git pull will look up what server and branch is tracking, fetch from that server and then try to merge in that remote branch. 

Generally it's better to simply use the fetch and merge commands explicitly as the magic of git pull can often be confusing.

### Deleting Remote Branches

Suppose you're done with a remote branch - say you and your collaborators are finished with a feature and have merged it into your remote's master branch. You can delete a remote branch using the --delete option to git push.

```bash
git push origin --delete serverfix
```

Basically all this does is to remove the pointer from the server. The Git server will generally keep the data there for a while until a garbage collection runs, so if it was accidentally deleted, it's often easy to recover.

## Rebasing

There are two main ways to integrate changes from one branch into another: the merge and the rebase. In this section you'll learn what rebasing is, how to do it, why it's pretty amazing tool, and in what cases you won't want to use it.

If you go back to earlier example from Basic Merging, you can see that you diverged your work and made commits on two different branches.

The easiest way to integrate the branches, as we've already covered, is the merge command. It performs a three-way merge between the two latest branch snapshots (C3 and C4) and the most rececent common ancestor of the (C2), creating a new snapshot (and commit). 

However, there is another way: you can take the patch of the change that was introduced in C4 and reapply it on top of C3. In git, this is called rebasing. With the rebase command, you can take all the changes  that were commited on one branch and reply them on different branch.

You would check out the experiment branch, and then rebase it onto the master branch as follows:

```bash
git checkout experiment
git rebase master
```

First, rewinding head to replay your work on top of it...
Applying: added staged command


