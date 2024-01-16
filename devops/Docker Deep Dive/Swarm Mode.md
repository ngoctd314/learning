# Swarm Model

Now that we know how to install Docker, pull images, and work with containers the next thing we need is a way to work with at all at scale. That's where orchestration and swarm mode comes into the picture.

As usual, we'll take a three-stage approach with a high-level explaination at the top, followed by a longer section with all the detail and some examples, and we'll finish things up with list of the main commands we learned.

The examples and outputs in this chapter will be from a Linux-based Swarm. However, all commands and features also work Docker on Windows.

## Swarm mode - The TLDR

It's one thing to follow along with the simple example in this book, but it's an entirely different thing running thousands of containers on tens or hundreds of Docker hosts! This is where orchestration comes into play!

At a high-level, orchestration is all about automating and simplifying the management of containerized applications at scale. Things like automatically rescheduling containers when nodes break, scaling things up when demand increases, and smoothly pushing updates and fixes into live production environments.

For the longest time orchestration like this was hard. Tools like Docker Swarm and Kubernetes were available, but they were complicated. Then along came Docker 1.12 and the new native swarm mode.

## Swarm mode - The deep dive

### Concepts and terminology

Swarm mode brought a load of changes and improvements to the way we manage containers at scale. At the heart of those changes is native clustering of Docker hosts that's deeply integrated into the Docker platform. We're not talking about something like K8s that's a separate tool requiring a highly skilled specialist to configure it on top of existing Docker infrastructures.
