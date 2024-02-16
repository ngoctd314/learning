## Overview 

This page is an overview of Kubernetes.

Kubernetes is a portable, extensible, open source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation.

## Going back in time

**Traditional deployment era:** Early on, organizations ran applications on physical servers. There was no way to define resource boundaries for applications in a physical server, and this caused resource allocation issues.

**Virtualized deployment era:** As a solution, virtualization was introduced. It allows you to run multiple Virtualized Machines (VMs) on a single physical server's CPU. Virtualizeda allows applications to be isolated between VMs and provides a level of security as the information of one application cannot be freely accessed by another application.

**Container deployment era:** Containers are similar to VMs, but they have relaxed isolation properties to share the Operating System (OS) among the applications. Therefore, containers are considered lightweight. Similar to a VM, a container has its own filesystem, share of CPU, memory, process space, and more. As they are decoupled from the underlying infrastructure.

## Why you need Kubernetes and what it can do

Kubernetes provides you with:

- Service discovery and load balancing
- Storage orchestration
- Automated rollouts and rollbacks
- Automatic bin packing
- Self-healing
- Secret and configuration management
- Batch execution
- Horizontal scaling
- IPv4/IPv6 dual-stack
- Designed for extensibility

## What Kubernetes is not

Kubernetes is not a traditional, all-inclusive PaaS (Platform as a Service) system. Since Kubernetes operates at the container level rather than at the hardware level, it provides some generally applicable features common to PaaS offerings, such as development, scaling, load balancing, and lets monolithic, and these default isolations are optional and pluggable.

Kubernetes:

- Does not limit the types of applications supported. Kubernetes aims to support an extremely diverse variety of workloads, including stateless, stateful, and data-processing workloads. If an application can run in a container, it should run great on Kubernetes. 
- Does not deploy source code and does not build your application. CI, Delivery, and Deployment (CI/CD) workflows are determined by organization cultures and preferences as well as technical requirements.

## Objects in Kubernetes

**Required fields**

In the manifest (YAML or JSON file) for the Kubernetes object you want to create, you'll need to set values for the following fields:

- apiVersion: Which version of the Kubernetes API you're using to create this object
- kind: What kind of object you want to create
- metadata: Data that helps uniquely identify the object, including a name string, UID, and optional namespace
- spec: What state you desire for the object
