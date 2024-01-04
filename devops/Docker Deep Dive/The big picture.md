# The big picture

The idea of this chapter is to give you a quick big picture of what Docker is all about before we dive in deeper in later chapter.

We'll break this chapter into two:

- The Ops perspective
- The Dev perspective

The Ops Perspective section will download an image, start a new container, log in to the new container, run a command inside of it, and then destroy it.

The Dev Perspective section will pull some app-code from GitHub, inspect a Dockerfile, containerize the app, run it as a container.

These two sections will give you a good idea of what Docker is all about and how some of the major components fit together. It is recommended that you read both sections to get the dev and the ops perspective!

## The Ops perspective

When you install Docker, you get two major components:

- the Docker client
- the Docker daemon (sometimes called "server" or "engine")

The daemon implements the Docker Remote API.

In a default Linux installation, the client talks to the daemon via a local IPC/Unix socket at /var/run/docker.sock. You can test that the client and daemon are running and can talk to each other with the docker version command.

```bash
Client: Docker Engine - Community
 Version:           24.0.6
 API version:       1.43
 Go version:        go1.20.7

Server: Docker Engine - Community
 Engine:
  Version:          24.0.6
  API version:      1.43 (minimum version 1.12)
  Go version:       go1.20.7
```

**Images**

A good way to think of a Docker image is an object that contains an OS filesystem and an application. If you're a developer, you can think of an image as a class.

Run the docker image ls command on your Docker host.

```bash
docker image ls
```

We'll get into the details of where the image is stored and what's inside of it in later chapters. For now, it's enough to understand that an image contains enough of an operating system (OS). An image contains enough of an operating system (OS), as well as all the code and dependencies to run whatever application it's designed for.

It's also worth noting that each image gets its own unique ID. When working with the images you can refer to them using either IDs or names.

**Containers**

Now that we have an image pulled locally on our Docker host, we can use the docker container run command to launch a container from it. 

```bash
docker container run -it ubuntu:latest /bin/bash
```

docker container run tells the Docker daemon to start a new container. The -it flags tell the daemon to make the container interactive and to attach our current terminal to the shell of the container. 

Press Ctrl-PQ to exit the container without terminating it. This will land you back in the shell of your Docker host. You can verify this by looking at your shell prompt.

**Attaching to running containers**

You can attach your shell to running containers with the docker container exec command. As the container from the previous steps is still running.

Stop the container and kill it using the docker container stop and docker container rm commands.

## The Dev perspective

Containers are all about the apps!
