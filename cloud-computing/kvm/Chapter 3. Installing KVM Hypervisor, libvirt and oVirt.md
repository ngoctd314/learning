# Chapter 3. Installing KVM Hypervisor, libvirt, and oVirt

This chapter provides you with an insight into the main topic of our book, which is the Kernel Virtual Machine (KVM) and its management tools, libvirt and oVirt. We will also learn how to do a complete installation of these tools from scratch using a basic deployment of CentOS 8.

## Getting accquainted with QEMU and libvirt

As a machine emulator, QEMU will be used so that we can create and run our virtual machines on any supported platform - be it as an emulator or virtualizer. We're going to focus our time on the second paradigm, which is using QEMU as a virtualizer. This means that we will be able to execute our virtual machine code directly on a hardware CPU below it, which means native or near-native performance and less overhead.

Bearing in mind that the overall KVM stack is built as a module, it shouldn't come as a surprise that QEMU also uses a modular approach. This has been a core principle in the Linux world for many years, which further boosts the efficiency of how we use our physical resources.

When we add libvirt as a management platform on top of QEMU, we get access to some cool new utilities such as the virsh command, which we can use to do virtual machine administration, virtual network administration, and a whole lot more. Some of the utilities that we've going to discuss later on in this book (for example, oVirt) use libvirt as a standardized set of libraries and utilities to make their GUI-magic possible - basically, they use libvirt as an API. There are other commands that we get access to for a variety of purposes. For example, we're going to use a command called virt-host-validate to check whether our server is compatible with KVM or not.

## Getting accquainted with oVirt

Bear in mind that most of the work that a sizeable percentage of Linux system administrators do is done via command-line utilities, libvirt, and KVM. They offer us a good set of tools to do everything that we need from the command line, as we're going to see in next part of this chapter. But also, we will get a hint as to what GUI-based administration can be like, as we're briefly going to discuss Virtual Machine Manager later in this chapter.

However, that still doesn't cover a situation in which you have loads of KVM-based hosts, hundreds of virtual machines, dozens of virtual networks interconnecting them, and a rack full of storage devices that you need to integrate with your KVM environment. Using the aforementioned utilities is just going to introduce you to a world of pain as you scale your environment out. The primary reason for this is rather simple - we still haven't introduced any kind of centralized software package for managing KVM - based environments. When we say centralized, we mean that in a literal sense - we need some kind of software solution that can connect to multiple hypervisors and manager all of their capabilities, including network, storage, memory, and CPU or, what we sometimes refer to as the four pillars of virtualization. This kind of software would preferably have some kind - well - we're all human. Quite a few of us perfer pictures to text, and interactivity to text-administrator only, especially at scale.

This is where oVirt project comes in. oVirt is an open source platform for the management of our KVM environment. It's a GUI-based tool that has a lot of moving parts in the background - the engine runs on Java-based WildFly server. Manage a KVM-based env from a centralized, web-based administration console.

## Installing QEMU, libvirt, oVirt


