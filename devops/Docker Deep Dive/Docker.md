# Docker

When somebody says "Docker" they can be referring to any of at least three things:

1. Docker, Inc. the company.
2. Docker the container runtime and orchestration technology.
3. Docker the open source project.

Docker is software that runs on Linux and Windows. It creates, manages and orchestrates containers. The software is developed in the open as part of the Moby open-source project on github.

Throughout this book we'll use the term "Docker, Inc." when referring to Docker the company. All other uses of the term "Docker" will refer to the technology or the open-source project.

## The Docker runtime and orchestration engine

When most technologists talk about Docker, they're referring to the Docker Engine.

The Docker Engine can be download from the Docker website or built from source from GitHub. It's available on Linux and Windows, with open-source and commercially supported offerings. 

## The docker open-source project (Moby)

The term "Docker" is also used to refer to the open-source Docker project. This is the set of tools that get combined into things like the Docker daemon and client you can download and install from docker.com. 

## The container ecosystem

One of the core philosophies at Docker, Inc. is often referred to as Batteries included but removable.

This is a way of saying you can swap out a lot of the native Docker stuff and replace it with stuff from 3rd parties. A good example of this is the networking stack. The core Docker product ships with built-in networking. But the networking stack is pluggable meaning you can rip out the native Docker networking and replace it with something else from a 3rd party.

## The open container initiative (OCI)
