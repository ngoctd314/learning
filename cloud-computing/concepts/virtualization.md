# What is virtualization?

A set of techniques and tools to create virtual solution for processes that would typically require a physical platform. Virtualization allows the same host/computer/server to run multiple guest operating systems using virtualization software known as a hypervisor and easily move VMs between hosts.

Virtualization forms the backbone of cloud computing. Virtualizationis an invaluable resource for managing complex systems while reducing hardware complexity

## Why Do You Need Virtualization for Your Business?

Virtualization can reduce costs, make creating new environments easier, and allow you to streamline tasks (like backups) across a virtual ecosystem.

You are always looking for ways to be more efficient so you can provide a product or service to your customer at a lower cost. In previous years, in order to add a server to handle added tasks as your business grew, you would have to purchase more physical servers. Purchasing servers can be very expensive. In turn, this new server would also require additional power to run and a sysadmin or IT team to configure, install and administer it. 

Physical servers do not use all of their computing resources 100 percent of the time, which is where virtualization comes in as a cost benefit. By having multiple servers in a virtual environment, you would no longer need to add physical servers. Instead, you could use the unused resources from your existing servers to address your business growth needs.

## What are the Types of Virtualization?

**Single-Tenant Server Virtualization**

Setting up a new physical server requires a lengthy installation and configuration process where a technician would need to put the sever together, install the operating system, and configure all of the applications required for the site or workload to function. Although this can be speed up by directly installing a preconfigured disk image, it wouldn't be completely eliminated.

With virtualization, that process can be automated and simplified. 

You can automate the provisioning process completely by using what's known as a Type 1 hypervisor like VMware vCenter. A Type 1 hypervisor is a program that runs directly on the physical server (also known as bare metal hypervisor) and provides a virtual environments on which operating systems can run. Since the hypervisor is always running, even when no operating system is present, it can create a new virtual machine from an image with minutes with no direct intervention.

This is perfect for servers with only one tenant using the resources. The tenant (you, in this case) would be the only one that would have full access to the phsical environment. This means you would be able to create as many virtual machines as you want or need as long as the physical server has enough resources to allocate to the virtual machines.

## Multi-Tenant Server Virtualization

Now, a virtual machine might not use all of the resources of the physical machine it was created on. In that case, another benefit of virtualization is the ability to host several virtual machines in a single physical machine, thereby sharing the physical resources of the host.

In this way, your business could have its database, mail, and web servers hosted on a single, more powerful machine that still functions as if the three servers were separated, significantly reducing the complexity of supporting the system.

Virtual machines don't have to have different purposes. Load balancing is a common use of virtualization that involves multiple machines handling the same data. By distributing requests across multiple virtual servers, a website can efficiently handle higher volume of traffic than a single server could manage.
