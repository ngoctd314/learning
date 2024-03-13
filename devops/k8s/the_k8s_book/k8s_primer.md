# K8s primer

## K8s background

K8s is an application orchestrator.

**What is a containerized app**

A containerized is an app that runs in a container.

Before we had containers, applications ran on physical servers or in virtual machines.

**Kubernetes and Docker**

Kubernetes and Docker are complementary technologies. For example, it's common to develop your application with Docker and use Kubernetes to orchestrate them in production.

In this model, you write your code in your favorite languages, then use Docker to package it, test it, and ship it. But the final steps of deploying and running it is handled by K8s.

In fact, K8s has a couple of features that abstract the container runtime (make it interchangeable):

- The Container Runtime Interface (CRI) is an abstraction layer that standardizes the way 3rd-party container runtimes interface with K8s. It allows the container runtime code to exist outside of K8s, but interface with it in a supported and standardizes way.
- Runtime Classes is a new feature that was introduced in K8s 1.12 and promoted to beta in 1.14. It allows for different classes of runtimes. For example, the gVisor or Kata Containers runtimes might provide better workload isolation than the Docker and containerd runtimes.

**What about K8s vs Docker Swarm**

In 2016 and 2017 we had the orchestrator wars where Docker Swarm, Mesosphere DCOS, and Kubernetes competed to become the de-facto container orchestrator. To cut a long story short, Kubernetes won.
