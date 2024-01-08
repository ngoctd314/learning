# Containers

## Docker containers - The TLDR

A container is the runtime instance of an image. In the same way that we can start a virtual machine (VM) from a virtual machine template, we start one or more containers from a single image. The big difference between a VM and a container is that containers are faster and more lightweight - instead of running a full-blown OS like a VM, containers share the OS/kernel with the host they're running on.  

![image to container](./assets/docker_image_containers.png)

The simplest way to start a container is with the docker container run command. The command can take a lot of arguments, but in its most basic form you tell it an image to use and a command to run:

```bash
docker container run <image> <command>
```

The -it flags used in the commands above will connect your current terminal window to the container's shell.

You can manually stop a container with the docker container stop command, and then restart it with docker container start. To get rid of a container forever you have to explicitly delete it using docker container rm.

That's the elevator pitch! Now let's get into the detail ...

## Docker containers - The deep dive

### Containers vs VMs

Containers and VMs both need a host to run on. We'll assume a single physical server that we need to run 4 business applications on.

In the VM model, the physical server is powered on and the hypervisor boots. Once the hypervisor boots it lays claim to all physical resources on the system such as CPU, RAM, storage, and NICs. The hypervisor then carves these hardware resources into virtual versions that look smell and feel exactly like the real thing. It then packages them into a software construct called virtual machine (VM). We then take those VMs and  install an os and application on each one. We said we had a single physical server and needed to run 4 applications, so we'd create 4 VMs, install 4 os, and then install 4 applications.

![vm 4 apps](./assets/vm_4apps.png)

Things are a bit different in the container model.

When the server is powered on, your chosen OS boots. In the Docker world this can be Linux or a modern version of Windows that has support for the container primitives in its kernel.

![docker 4 apps](./assets/docker_4apps.png)

containers perform OS virtualization - they carve up OS resources into virtual versions.

## The VM tax

## Checking the Docker daemon

## Starting a simple container

This simplest way to start a container is with the docker container run command.

The command below starts a simple container with will run a containerized version of Ubuntu Linux.

```sh
docker container run -it ubuntu:latest /bin/bash
```

When we hit Return, the Docker client made the appropriate API calls to the Docker daemon. The Docker daemon accepted the command and searched the Docker host's local cache to see if it already had a copy of the requested image. In this example it didn't, so it went to Docker Hub to see if it could find it there. It could, so it pulled it locally and stored it in its cache.

Once the image was pulled, the daemon created the container and executed the specified command inside of it.

If you look closely you'll see that your shell prompt has changed and you're now inside of the container.

Try executing some basic commands from inside of the container. You might notice that some commands do not work. This is because the images we used, like almost all container images, are highly optimized for containers. This means they don't have all of the normal commands and packages installed.  

## Container lifecycle

It's a common myth that containers can't persist data. They can!

A big part of the reason people think containers aren't good for persistent workloads, or persisting data, is because they're so good at non-persistent stuff. But being good at one thing doesn't mean you can't do other things.

You can stop, start, pause, and restart a container as many time as you want. And it'll all happen really fast. But the container and it's data will always be safe. It's not until you explicitly kill a container that you run any chance of losing its data. And even then, if you're storing container data in a volumne, that data's going to persist even after the container has gone.

## Stopping containers gracefully

Most containers in the Linux world will run a single process.

When you kill a running container with docker container rm <container> -f the container will be killed without warning. The procedure is quite violent - a bit like sneaking up behind the container it and shooting it in the back of the head. You're literally giving the container, and the process it's running, no chance to straighten its affairs before being killed.

However, the docker container stop command is far more polite (like pointing a gun to the containers head and saying "you've got 10 seconds to say any final words"). It gives the process inside of the container a heads-up that it's about to be stopped. Once the docker stop command returns, you can then delete the container with docker container rm.

The magic behind the scenes here can be explained with Linux/POSIX signals. Docker container stop sends a SIGTERM signal to the process with PID 1 inside of the container. As we just said, this gives the process a chance to clean things up and gracefully shut itself down. If it doesn't exit within 10 seconds it will receive a SIGKILL.

docker container rm <container> -f doesn't bother asking nicely with a SIGTERM, it just goes straight to the SIGKILL. Like we said a second ago, this is like creeping up from behind and smashing it over the head.

## Web server example

## Inspecting containers

The entries after "Cmd" show the command(s) that the container will run unless you override them with a different command as part of docker container run. If you remove all of the shell escapes in the example above, you get the following command /bin/sh -c "cd /src && node ./app.js". That's the default command a container based on this image will run.

## Tidying up

Here we're going to show you the simplest and quickest way to get rid of every running container on your Docker host. Be warned though, the procedure will forcible destroy all containers without giving them a chance to clean up. This should never be performed on production systems or systems running important containers.

Run the following command from the shell of your Docker host to delete all containers.

```sh
docker container rm $(docker container ls -aq) -f
```

## Containers - The commands

- docker container run is the command used to start new containers.
- Ctrl-PQ will detach your shell from the terminal of a container and leave the container running (UP) in the background
- docker container ls lists all containers in the running (UP) state. If you add the -a flag you will also see containers in the stopped (Exited) state.
- docker container exec lets you run a new process inside of a running container. It's useful for attaching the shell of your Docker host to a terminal inside of a running container.
