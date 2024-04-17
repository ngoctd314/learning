# Understanding Linux Virtualization

Virtualization is the technology that started a big technology shift toward IT consolidation.

## Linux virtualization and how it all started

Virtualization is a concept that creates virtualized resources and maps them to physical resources. This process can be done using specific hardware functionality (partitioning, via some kind of partition controller) or software functionality (hypervisor). So, as an example, if you have a physical PC-based server with 16 cores running a hypervisor, you can easily create one or more virtual machines with two cores each and start them up. Limits regarding how many virtual machines you can start is something that's vendor-based. In any case, hypervisor is going to be go-to guy that's going to try to manage that as efficiently as possible so that all of the virtual machine workloads get as much time on the CPU as possible.  

Virtualization solutions - Red Hat with KVM, Microsoft with Hyper-V, VMWare with ESXi, Oracle with Oracle VM,... This led to the development of various cloud solutions such as EC2, AWS, Office 365, Azure, vCloud...

Going back to October 2003, with all of the changes that were happening in the IT industry, there was one that was really important for this book and virtualization for Linux in general: the introduction of the first open source Hypervisor for x86 architecture, called Xen. 

Technically speaking, KVM uses a different, modular approach that transforms Linux kernels into fully function hypervisors for supported CPU architectures. When we say supported CPU architectures, we're talking about the basic requirement for KVM virtualization - CPUs need to support hardware virtualization extensions, known as AMD-V or Intel VT. To make things a bit easier, let's just say that you're really going to have to try very hard to find a modern CPU that doesn't support these extensions.

## Types of virtualization

There are various types of virtualization solutions, all of which are aimed at different use cases and are dependent on the fact that we've virtualization a different piece of the hardware or software stack. There are different types of virtualization in terms of how you're virtualizing - by partitioning, full virtualization, paravirtualization, hybrid virtualization, or container-based virtualization.

### Desktop virtualization (Virtual Desktop Infrastructure (VDI))

This is used by a lot of enterprise companies and offers huge advantages for a lot of scenarios because of the fact that users aren't dependent on a specific device that they're using to access their desktop system. They can connect from a mobile phone, tablet, or a computer, and the can usually connect to their virtualized desktop from anywhere as if they're sitting at their workplace and using a hardware computer. Benefits include easier, centralized management and monitoring much more simplified update workflows (you can update the base image for hundreds of virtual machines in a VDI solution)

https://youtu.be/weII6qT59mQ

### Server virtualization 

This is used by a vast majority of IT companies today. It offers good consolidation of server virtual machines versus physical servers, while offering many other operational advantages over regular, physical servers - easier to backup, more energy efficient, more freedom in terms of moving workloads from server to server, and more.

### Application virtualization

### Networking virtualization

More broader, cloud-based concept called Software-Defined Networking (SDN): This is a technology that creates virtula networks that are independent of the physical networking devices, such as switches. SDN is an extension of the network virtualization idea that can span across multiple sites, locations, or data centers. In terms of the concept of SDN, entire network configuration configuration is done in software, without you necessarily of network virtualization is how easy it is for you to manage complex networks that span multiple locations without having to do massive, physical network reconfiguration for all the physical devices on the network data path.

### Storage virtualization (Software-Defined Storage)

This is a technology that creates virtual storage devices out of pooled, physical storage devices that we can centrally manage as a single storage device. This means that we're creating some sort of abstraction layer that's going to isolate the internal functionality of storage devices from computers, applications, and other types of resources. SDS, as an extensions of that, decouples the storage software stack from the hardware it's running on by abstracting control and management planes from the underlying hardware, as well as offering different types of storage resources to virtual machines and applications 

If you take a look at these virtualization solutions and scale them up massively, that's when you realize that you're going to need various tools and solutions to efficiently manage the ever-growing infrastructure, hence the development of various automatization and orchestration tools.

If we're talking about how we're virtualizing a virtual machine as an object, there are different types of virtualization;

- Partitioning: This is type of virtualization in which a CPU is divided into different parts, and each part works as an individual system. This type of virtualization solution isolates a server into partitions, each of which can run a separate OS
- Full virtualization: In full virtualization, a virtual machine is used to simulate regular hardware while not being aware of the fact that it's virtualized. This is done for compatibility reasons - we don't have to modify the guest OS that we're going to run in a virtual machine. We can use a software - and hardware-based approach for this:
- Software-based: Used binary translation to virtualize the execution  of sensitive instruction sets while emulating hardware using software, which increases overhead and impacts scalability.
- Hardware-based: Removes binary translation from the equation while interfacing with a CPU's virtualization features (AMD-V, Intel V-T), which, in turn, means that instruction sets are being executed directly on the host CPU. This is what KVM does.
- Paravirtualization
- Hybrid virtualization
- Container-based virtualization: This is a type of application virtualization that use containers. A container is an object that packages an application and all its dependencies so that the application can be sscaled out and rapidly deployed without needing a virtual machine or a hypervisor. Keep in mind that there are technologies that can operate as both a hypervisor and a container host at the same time.
