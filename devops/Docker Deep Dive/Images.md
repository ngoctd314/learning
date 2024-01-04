# Images

In this chapter we'll dive into Docker images. The aim of the game is to give you a solid understanding of what Docker images are and how to perform basic operations.

## Docker images - The TLDR

If you're a former VM admin you can think of Docker images as being like VM templates. A VM template is like a stopped VM - a Docker image is like a stopped container. If you're a developer you can think of them as being similar to classes.

You start by pulling images from an image registry. The most popular registry is Docker Hub, but others do exist. The pull operation downloads the image to your local Docker host where you can use it to start one or more Docker containers.

Images are made up of multiple layers that get stacked on top of each other and represented as a single object. Instead of the image is a cut-down operating system (OS) and all of the files and dependencies required to run an application. Because containers are intended to be fast and lightweight, images tend to be small.

## Docker images - The deep dive  

We've mentioned a couple of times already that images are like stopped containers (or classes if you're a developer). In fact, you can stop a container and create a new image from it. With this in mind, images are considered build-time constructs whereas containers are run-time constructs.

### Images and containers

We use the docker container run and docker service create commands to start one or more containers from a single image. However, once you've started a container from an image, the two constructs become dependent on each other and you cannot delete the image until the last container using it has been stopped and destroyed. 

### Images are usually small

The whole purpose of a container is to run an application or service. This means that the image a container is created from must contain all OS and application files required to run he app/service. However, containers are all about being fast and lightweight. This means that the images they're built from are usually small and stripped of all non-essential parts. 

For example, Docker images do not ship with 6 different shells for you to choose from - they usually ship with a single minimalist shell, or no shell at all. They also don't contain a kernel - all containers running on a Docker host share access to the host's kernel. For these reasons, we sometimes say images contain just enough os (usually just OS-related files and filesystem objects). 

The official Alpine Linux Docker image is about 4MB in size and is an extreme example of how small Docker images can be. That's not a typo! It really is about 4 megabytes! However, a more typical example might be something like the official Ubuntu Docker image which is currently about 120MB.

### Pulling images

The process of getting images onto a Docker host is called pulling. So, if you want the lastest Ubuntu image on your Docker host, you'd have to pull it. Use the commands below to pull some images and then check their sizes.

### Image registries

Docker images are stored in image registries. The most common registry is Docker Hub (https://hub.docker.com). Other registries exist, including 3rd party registries and secure on-premises registries.

Docker Hub also has the concept of official repositories and unofficial repositories.

As the name suggests, official repositories contain images that have been vetted by Docker, Inc. This mean they should contain up-to-date, high-quality code, that is secure, well-documented, and in-line with best practices.  

Unofficial repositories can be like the wild-west - you should not expect them to be safe, well-documented or built according to best practices. That's not saying everything in unofficial repositories is bad! There's some brilliant stuff in unoffical repositories. You just need to be very careful when getting software from the internet - even images from official repositories!

Most of the popular os and applications have their own official repositories on Docker Hub.

### Image naming and tagging

Addressing images from official repositories is as simple as giving the repository name and tag separated by a colon (:). The format for docker image pull when working with an image from an official repository is:

```txt
docker image pull <repository>:<tag>
```

First, if you do not specify an image tag after the repository name, Docker will assume you are referring to the image tagged as latest.

### Images and layers

A Docker image is just a bunch of loosely-connected read-only layers.

Docker takes care of stacking these layers and representing them as a single unified object.

Docker employs a storage driver that is responsible for stacking layers and presenting them as a single unified filesystem. Examples of storage drivers on Linux include AUFS, overlay2, devicemapper.

### Sharing image layers

Multiple images can, and do, share layers. This leads to efficiencies in space and performance.

Let's take a second look at the docker image pull command with the -a flag nigelpoulton/tu-demo repository.

As mentioned previously, Docker on Linux supports many different filesystems and storage drivers. Each is free to implement image layering, layer sharing, and copy-on-write behavior in its own way.

### Pulling images by digest

So far, we've shown you how to pull images by tag, and this is by far the most common way. But it has a problem - tags are mutable! This means it's possible to accidentally tag an image with an incorrect tag. Sometimes it's even possible to tag an image with the same tag as an existing, but different, image.

### Deleting Images

When you no longer need an image, you can delete it from your Docker host. With the docker image rm command.

Delete the images pulled in the previous steps with the docker image rm command. 

If the image you are trying to delete is in use by a running container you will not be able to delete it. Stop and delete any containers before trying the remove operation again.

A handy shortcut for cleaning up a system and deleting all images on a Docker host is to run the docker image rm command and pass it a list of all image IDs on the system by calling docker image ls with the -q flag. This is shown below.

```bash
docker image rm $(docker image ls -q) -f
```

## Images - The commands

- docker image pull is the command to download images. We pull images from repositories inside of remote registries. By default, images will be pulled from repositories on Docker Hub. This command will pull the image tagged as latest from the alpine repository on Docker Hub docker image.

- docker image ls list all of the images stored in your Docker host's local cache. To see the SHA256 digests of images add the --digests flag. 

- docker image inspect is a thing of beauty! It gives you all of the glorious details of image - layer data and metadata.

- docker image rm is the command to delete images.
