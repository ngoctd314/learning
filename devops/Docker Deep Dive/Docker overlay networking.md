# Docker overlay networking

Container networking is increasingly important. Especially in production environments.

In this chapter we'll cover the fundamentals of native Docker overlay networking as implemented in a Docker overlay networking as implemented in a Docker swarm cluster.

## Docker overlay networking - The TLDR

In the real world it's vital that containers can communicate with each other reliably and securely, even when they're on different hosts on different networks. This is where overlay networking comes in to play. It allows you to create a flat secure layer 2 network spanning multiple hosts that containers can connect to. Containers on this network can then communicate directly.

Docker offers native overlay networking that is simple to configure and secure by default.

## Docker overlay networking - The deep dive
