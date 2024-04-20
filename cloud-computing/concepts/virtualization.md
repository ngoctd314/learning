# What is virtualization?

- https://www.youtube.com/watch?v=UBVVq-xz5i0

A set of techniques and tools to create virtual solution for processes that would typically require a physical platform. Virtualization allows the same host/computer/server to run multiple guest operating systems using virtualization software known as a hypervisor and easily move VMs between hosts.

Virtualization forms the backbone of cloud computing. Virtualizationis an invaluable resource for managing complex systems while reducing hardware complexity

## Why Do You Need Virtualization for Your Business?

Virtualization can reduce costs, make creating new environments easier, and allow you to streamline tasks (like backups) across a virtual ecosystem.

You are always looking for ways to be more efficient so you can provide a product or service to your customer at a lower cost. In previous years, in order to add a server to handle added tasks as your business grew, you would have to purchase more physical servers. Purchasing servers can be very expensive. In turn, this new server would also require additional power to run and a sysadmin or IT team to configure, install and administer it. 

Physical servers do not use all of their computing resources 100 percent of the time, which is where virtualization comes in as a cost benefit. By having multiple servers in a virtual environment, you would no longer need to add physical servers. Instead, you could use the unused resources from your existing servers to address your business growth needs.

## What are the Types of Virtualization?

## Single-Tenant Server Virtualization

Setting up a new physical server requires a lengthy installation and configuration process where a technician would need to put the sever together, install the operating system, and configure all of the applications required for the site or workload to function. Although this can be speed up by directly installing a preconfigured disk image, it wouldn't be completely eliminated.

With virtualization, that process can be automated and simplified. 

You can automate the provisioning process completely by using what's known as a Type 1 hypervisor like VMware vCenter. A Type 1 hypervisor is a program that runs directly on the physical server (also known as bare metal hypervisor) and provides a virtual environments on which operating systems can run. Since the hypervisor is always running, even when no operating system is present, it can create a new virtual machine from an image with minutes with no direct intervention.

This is perfect for servers with only one tenant using the resources. The tenant (you, in this case) would be the only one that would have full access to the phsical environment. This means you would be able to create as many virtual machines as you want or need as long as the physical server has enough resources to allocate to the virtual machines.

## Multi-Tenant Server Virtualization

Now, a virtual machine might not use all of the resources of the physical machine it was created on. In that case, another benefit of virtualization is the ability to host several virtual machines in a single physical machine, thereby sharing the physical resources of the host.

In this way, your business could have its database, mail, and web servers hosted on a single, more powerful machine that still functions as if the three servers were separated, significantly reducing the complexity of supporting the system.

Virtual machines don't have to have different purposes. Load balancing is a common use of virtualization that involves multiple machines handling the same data. By distributing requests across multiple virtual servers, a website can efficiently handle higher volume of traffic than a single server could manage.

## Desktop Virtualization

While virtualization is beneficial for Internet services and the cloud, that is by no means its only use case. Virtual machines on desktop computers are useful for many different purposes. Virtualization can be used to test applications developed for multiple os os a single machine or to run legacy applications that aren't compatible with newer os.

Virtual machines on desktop computers are useful for many different purposes. Virtualization can be used to test applications developed for multiple os on a single machine or to run legacy applications that aren't compatible with newer os.

Virtualization on desktops is accomplished through what's known as Type 2 hypervisor, which exists alongside the non-virtualized os that the computer uses typically (instead of having full control over all of the system's resources).

That said, when you run virtual machines on personal computers, it's also possible to turn things around and run personal computers as virtual machines instead. There is usually a single, central machine running one (or several) virtual computers in this type of setup.

Individual users connect through devices knonwn as thin clients to establish remote desktop connections without the full power and cost associated with a regular workstation. This simplifies IT costs and management in officices with many employees by centralizing individual workstations. Remote desktop is also every useful in shools and other similar organizations.

## Storage Virtualization

Data virtualization has a lot of benefits over directly storing data on individual hard drives. Data virtualization is:

- Accessible: Allows for multiple clients to access the same data (as any networked storage would) without having to manage the technical details of how data is stored.
- Safe: Administrators can quickly implement safety measures such as RAID mirroring and parity checks to minimize the chance of data loss during a drive value.
- Simple: Presents data to clients as if it was stored on a single value, simplifying access.

## Network Virtualization

We've discussed ways to virtualize servers, workstations, and network-attached storage, but what about virtualizing networks and the devices that connect them together? Many network devices that are used physically can be virtualization to minimize cost and make management easier:

- Load balancers, firewall, intrusion detection systems, wan

## Application Virtualization

Two popular solutions for application virtualization:

- Virtual Machine: Used by the Java runtime environment, virtual machines allow programs developed in Java to run regardless of the platform.
- Docker Containers: Docker is an open-source solution allowing developers to provide a defined set of dependencies for an application regardless of what is installed on the host machine without the need to emulate an entire operating system.

## Hypervisors

The software that creates and runs the virtualization is called a hypervisor

A hypervisor is what allows one machine to run multiple virtual machines. 

It allocates and controls the sharing of a machines resources.

- Storage space, RAM, CPUs

Hypervisors come in two different types. There's type 1 and type 2. Type 1 hypervisor is installed on bare metal hardware. Meaning that there is no existing operating system or any other software on the machine. The hypervisor is installed on empty, bare metal hardware. A type 2 hypervisor is installed and runs on top of existing operating system, such as Microsoft Windows, Linux, Unix and so on. So the operating system sits in between the machine and the hypervisor.

Type 1 hypervisor: Example VMware ESXi, Citrix XenServer, Microsoft Hyper-V

Type 2 hypervisor: runs on top of an existing os. Typically used on personal computers. 

## Virtualization Benefits

- Saves money on hardware and electricity.
- Saves money on floor space.
- Saves money on maintenance and management.
- Portability.
- Full computing capability.
- Disaster and recovery.
