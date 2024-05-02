# Chapter 3. Installing KVM Hypervisor, libvirt, and oVirt

This chapter provides you with an insight into the main topic of our book, which is the Kernel Virtual Machine (KVM) and its management tools, libvirt and oVirt. We will also learn how to do a complete installation of these tools from scratch using a basic deployment of CentOS 8.

## Getting accquainted with QEMU and libvirt

We started discussing KVM, QEMU, and various additional utilities that we can use to manage our KVM-based virtualization platform.

As a machine emulator, QEMU will be used so that we can create and run our virtual machines on any supported platform - be it as an emulator or virtualizer. We're going to focus our time on the second paradigm, which is using QEMU as a virtualizer. This means that we will be able to execute our virtual machine code directly on a hardware CPU below it, which means native or near-native performance and less overhead.

Bearing in mind that the overall KVM stack is built as a module, it shouldn't come as a surprise that QEMU also uses a modular approach. This has been a core principle in the Linux world for many years, which further boosts the efficiency of how we use our physical resources.

When we add libvirt as a management platform on top of QEMU, we get access to some cool new utilities such as the virsh command, which we can use to do virtual machine administration, virtual network administration, and a whole lot more. Some of the utilities that we've going to discuss later on in this book (for example, oVirt) use libvirt as a standardized set of libraries and utilities to make their GUI-magic possible - basically, they use libvirt as an API. There are other commands that we get access to for a variety of purposes. For example, we're going to use a command called virt-host-validate to check whether our server is compatible with KVM or not.

## Getting accquainted with oVirt

Bear in mind that most of the work that a sizeable percentage of Linux system administrators do is done via command-line utilities, libvirt, and KVM. They offer us a good set of tools to do everything that we need from the command line, as we're going to see in next part of this chapter. But also, we will get a hint as to what GUI-based administration can be like, as we're briefly going to discuss Virtual Machine Manager later in this chapter.

However, that still doesn't cover a situation in which you have loads of KVM-based hosts, hundreds of virtual machines, dozens of virtual networks interconnecting them, and a rack full of storage devices that you need to integrate with your KVM environment. Using the aforementioned utilities is just going to introduce you to a world of pain as you scale your environment out. The primary reason for this is rather simple - we still haven't introduced any kind of centralized software package for managing KVM - based environments. When we say centralized, we mean that in a literal sense - we need some kind of software solution that can connect to multiple hypervisors and manager all of their capabilities, including network, storage, memory, and CPU or, what we sometimes refer to as the four pillars of virtualization. This kind of software would preferably have some kind - well - we're all human. Quite a few of us perfer pictures to text, and interactivity to text-administrator only, especially at scale.

This is where oVirt project comes in. oVirt is an open source platform for the management of our KVM environment. It's a GUI-based tool that has a lot of moving parts in the background - the engine runs on Java-based WildFly server. Manage a KVM-based env from a centralized, web-based administration console.

### Virtual Machine Manager

The virtual machine manager, through the virt-manager package, provides a graphical user interface (GUI) for managing local and remote virtual machines. In addition to the virt-manager utility itself, the package also contains a collection of other helpful tools like virt-install, virt-clone and virt-viewer.

To install virt-manager, enter:

```sh
sudo apt install virt-manager
```

Since virt-manager requires a Graphical User Interface (GUI) environment we recommend installing it on a workstation or test machine instead of production server.

To connect to the local libvirt service, enter:

```sh
virt-manager
```

You can connect to the libvirt service running on another host by entering the following in a terminal prompt:

```sh
virt-manager -c qemu+ssh://virtnode1.mydomain.com/system
```

### Virtual Machine Viewer (virt-viewer)

The Virtual Machine Viewer application, through virt-viewer, allows you to connect to a virtual machine's console like virt-manager does, but reduced to the GUI functionality. virt-viewer requires a GUI to interface with the virtual machine.

virt-viewer can also connect to a remove host using SSH with key authentication:

```sh
virt-viewer -c qemu+ssh://virtnode1.mydomain.com/system <guestname>
```

### virt-install

virt-install is part of the virtinst package. It can help with installing classic ISO-based systems and provides a CLI for the most common options needed to do so.

### Manage VMs with virsh

There are several utilities available to manage virtual machines and libvirt. The virsh utility can be used from the command line. Some examples:

- To list running virtual machines:

```sh
virsh list
```

- To start a virtual machine:

```sh
virsh start <guestname>
```

- Similarly, to start a virtual machine at boot:

```sh
virsh autostart <guestname>
```

- Reboot a virtual machine with:

```sh
virsh reboot <guestname>
```

- To shut down a virtula machine you can do:

```sh
virsh shutdown <guestname>
```

- A CD-ROM device can be mounted in a virtual machine by entering:

```sh
virsh attach-disk <guestname> /dev/cdrom /media/cdrom
```

- To change the definition of a guest, virsh exposes the domain via:

```sh
virsh edit <guestname>
```

## Installing QEMU, libvirt, oVirt, virt-manager

After going through a basic installation of our server - selecting the installation profile assigning network configuration and root password, and adding additional users (if we need them) - we're faced with a system that we can't do virtualization with because it doesn't have all of the necessary utilities to run KVM virtual machines. So, the first thing that we're going to do is a simple installtion of the necessary modules and base applications so that we can check whether our server is compatible with KVM. So, log into your server as an administrative user and  issue the following command:

```bash
yum module install virt
dnf install qemu-img qemu-kvm libvirt libvirt-client virt-manager virt-install virt-viewer -y
```

We alos need to tell the kernel that we're going to use IOMMU. This is achieved by editing /etc/default/grub file, finding the GRUB_CMDLINE_LINUX and adding a statement at the end of this line:

```bash
intel_iommu=on
```

IOMMU (Input-Output Memory Management Unit) support in kernel. IOMMU is a feature of modern hardware that provides memory address translation and access control for devices that perform direct memory access (DMA). It's particularly important for virtualization, as it helps to isolate and secure devices accessed by virtual machine.

In your case, QEMU is suggesting that you enable IOMMU support in kernel by adding `intel_iommu=on` to the kernel command-line arguments. This argument specifically enables Intel VT-d (Intel Virtualization Technology for Directed I/O) support, which is Intel's implementation of IOMMU.

Don't forget add a single space before adding the line. Next step is reboot, so, we need to do:

```bash
systemctl reboot
```

By issuing these commands, we're installing all the necessary libraries and binaries to our KVM-based virtual machines, as well as to use virt- manager to manage our KVM virtualization server.

After that, let's check whether our host is compatible with all the necessary KVM requirements by issuing the following command:

```sh
virt-host-validate
```

This command goes through multipl tests to determine whether our server is compatible or not.

This shows that our server is ready for KVM. So, the next. step, now that all the necessary QEMU/libvirt utilities are installed, is to do some pre-flight checks to see whether everything that we installed was deployed correctly and works like it should. We will run the virsh net-list and virsh list commands to do this.

```sh
virsh net-list
```

This command checked whether our virtualization host has a correctly configured default virtual network switch/bridge (more about this in the next chapter)  

```txt
----------------------------------------
```
 Name   State   Autostart   Persistent


```sh
virsh list
```

This command checked whether we have any virtual machines running.

```txt
 Id   Name   State
--------------------
```

**Installing the first virtual machine in KVM**

We can now start using our KVM virtualization service for its primary purpose - to run virtual machines. Let's start by deploying a virtual machine on our host. For this purpose, we copied a CentOS 8.0 ISO file to our local folder called `/var/lib/libvirt/images`, which we're going to use to create our first virtual machine. We can do that from the command line by using the following command.

```bash
virt-install --name=ubuntu_test1 \
--vcpus=2 \
--ram=2048 \
--disk=/var/lib/libvirt/images/ubuntu_test1.qcow2,size=15 \
--cdrom=/var/lib/libvirt/images/ubuntu-20.04-amd64.iso \
--network=default \
--graphics vnc
```

There are some parameters here that might be a bit confusing. Let's start with the --os-variant parameter, which describes which guest os you want to install by using virt-install command. If you want to get a list of supported guest os, run the following command:

```osinfo-query os```

The --network parameter is related to our default virtual bridge (we mentioned this earlier). We definitely want our virtual machine to be network-connected, so we picked this parameter to make sure that it's network-connected out of the box.

After starting virt-install command, we should be presented with a VNC console window to follow along with the installation procedure. We can then select the language used, keyboard, time and date, and installation destination (click on the selected disk and press Done in the top-left corner). We can also active the network by going to Network & Host Name, clicking on the OFF button, selecting Done (which will then switch to the ON position), and connecting our virtual machine to the underlying network bidge (default). 

If all of this seems a bit like manual labor to you, we feel your pain. Imagine having to deploy dozens of virtual machines and clicking on all these settings. We're not in the 19th century anymore, so these must be an easier way to do this.

### Automating virtual machine installtion

By far, the simplest and the easiest way to do these things in a more automatic fashion would be to create and use something called a kickstart file. A kickstart file is basically a text configuration file that we use to configure all the deployment settings of our server, regardless of whether we're talking about a physical or virtual server. The only caveat is that kickstart files need to be pre-prepared and widely available - either on the network (web) or on local disk. There are other options that are supported, but these are the most commonly used ones.

For our purpose, we're going to use a kickstart file that's available on the network (via the web server).

When we installed our physical server, as part of the installation process (called anaconda), a file was saved in our /root directory called anaconda-ks.cfg. This is a kickstart file that contains the complete deployment configuration of our physical server, which we can then use as a basis to create a new kickstart file for our virtual machines.

...

### Installing oVirt

There are different methods of installing oVirt. We can either deploy is as a self-hosted engine (via the Cockpit web interface or CLI) or as a standalone application via package-based installation.

1. Install the oVirt engine for centralized management
2. Deploying oVirt agents on our CentOS 8-based hosts

### Starting a virtual machine using QEMU and libvirt

After the deployment process, we can start managing our virtual machines. We will use MasteringKVM01 and MasteringKVM02 as an example. Let's start them by using the virsh command, along with the start keyword:

```sh
virsh start ubuntu_test1
```

```txt
Domain 'ubuntu_test1' started
```

We can easily check their status by issuing a simple virsh list command:

```sh
virsh list
```

```txt
 Id   Name           State
------------------------------
 8    ubuntu_test1   running
```

If we want to gracefully shut down the ubuntu_test1 virtual machine, we can do so by using virsh shutdown (gracefull) command:

```sh
virsh shutdown ubuntu_test1
```

If we want to forcefully shut down the ubuntu_test1 virtual machine, we can do so by using the virsh destroy command:

```sh
virsh destroy ubuntu_test1
```

We just learned how to use the virsh command to manage virtual machine - start it and stop it - forcefully and gracefully. This will come in handy when we start extending our knowledge of using the virsh command in the following chaters, in which we're going to learn how to manage KVM networking and storage.

### Summary

In this chapter, we laid some basic groundwork and prerequisites for practically everything that we're going to do in the remaining chapters of this book. We learned how to install KVM and a libvirt stack. We also learned how to deploy oVirt as a GUI tool to manage our KVM hosts.

The next few chapters will take us in a more technical direction as we will cover networking and storage concepts. In order to do that, we will have to take a stop back and learn or review our previous knowledge about networking and storage as these are extremely important concepts for virtualization, and especially the cloud.

References:

- https://phoenixnap.com/kb/ubuntu-install-kvm
- https://ubuntu.com/server/docs/virtual-machine-manager
- https://ubuntu.com/server/docs/libvirt
