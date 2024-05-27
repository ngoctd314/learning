# Chapter 1. System Requirements

Virtualization is available with the KVM hypervisor for Red Hat Enterprise Linux 7 on the Intel 64 and AMD64 architectures. This chapter lists system requirements for running virtual machines, also referred to as VMs.

## 1.2. KVM hypervisor requirements

The KVM hypervisor requires:

- An intel processor with the Intel VT-x and Intel 64 virtualization extensions for x86-based systems.
- An AMD processor with the AMD-V and AMD64 virtualization extensions.

Virtualization extensions (Intel VT-x and AMD-V) are required for full virtualization.

# Chapter 2. Installing the virtualization packages

# Chapter 3. Creating a VM

# Chapter 4. Cloning virtual machine

# Chapter 5. KVM paravirtualized (virtio) driver

Paravirtualized drivers enhance the performance of guests, decreasing guest I/O latency and increasing throughput almost to bare-metal levels. It is recommended to use the paravirtualized drivers for fully virtualized guests running I/O-heavy tasks and applications.

Virtio drivers are KVM's paravirtualized device drivers, available for guest virtual machines running on KVM hosts. There drivers are included in the virtio package. The virtio package supports block (storage) devices and network interface controllers.

**NOTE**

PCI devices are limited by the virtualized system architecture.

## 5.1. Using KVM virtio drivers for existing storage devices

You can modify an existing hard disk device attached to the guest to use the virtio driver instead of the virtualized IDE driver. The example shown in this section edits libvirt configuration files. Note that the guest virtual machine does not need to be shut down to perform these steps, however the change will not be applied until the guest is completely but shut down and rebooted.

**Using KVM virtio drivers for existing devices**

```xml
<disk type='file' device='disk'>
    <driver name='qemu' type='qcow2'/>
    <source file='/var/lib/libvirt/images/ubuntu_test1.qcow2' index='1'/>
    <target dev='vda' bus='virtio'/>
    <address type='pci' domain='0x0000' bus='0x04' slot='0x00' function='0x0'/>
</disk>
```

- `dev='vda'`: This attribute specifies the device name that the os inside VM will see. In this case, `vda` typically indicates a virtual disk (the first one, hence 'a'). This naming convention follows traditional Linux naming schemes where 'vda' could be analogous to 'sda' for the first SCSI disk. Additional disks would typically be named 'vdb', 'vdc', etc.

- `bus='virtio'`: This attribute specifies the type of virtual bust this device is attached to. `virtio` is a paravirtualized bus, meaning it provides a simple and efficient abstraction of the actual hardware bus.

## 5.2. Using KVM virtio drivers for new storage devices

## 5.3. Using KVM virtio drivers for network interface devices

# Chapter 6. Network configuration

This chapter provides an introduction to the common networking configurations used by libvirt-based guest virtual machines.

Red Hat Enterprise Linux 7 supports the following networking setups for virtualization:

- virtual networks using Network Address Translation (NAT)
- directly allocated physical devices using PCI device assignment
- directly allocated virtual functions using PICe SR-IOV
- bridged networks

## 6.1. Network address translation (NAT) with libvirt

One of the most common methods for sharing network connections is to use Network Address Translation (NAT) forwarding

**Host Configuration**

# Chapter 7. Overcomiting with KVM

## 7.1. Introduction

The KVM hypervisor automatically overcommits CPUs and memory. This means that more virtulized CPUs and memory can be allocated to virtual machines that there are physical resources on the system. This is possible because most processes do not access 100% of their allocated resources all the time.

As a result, under-utilized virtualized servers or desktops can run on fewer hosts, which saves a number of system resources, with the net effect of less power, cooling and investment in server hardware.

## 7.2. Overcomiting memory

Guest virtual machines running on a KVM hypervisor do not have dedicated blocks of physical RAM assigned to them. Instead, each guest virtual machine functions as a Linux process where the host physical machine's Linux kernel allocates memory only when requested. In addition the host's memory manager can move the guest virtual machine's memory between its own physical memory and swap space.

Overcommitting requires allotting sufficient swap space on the host physical machine to accommodate all geust virtual machines as well as enough memory for the host physical machine's processes. As a basic rule, the host physical machine's operating system requires a maximum of 4 GB of memory along with a minimum of 4 GB of swap space.

**IMPORTANT**

Overcommitting is not an ideal solution for general memory issues. The recommended methods to deal with memory shortage are to allocate less memory per guest, add more physical memory to the host, or utilize swap space.

A virtual machine will run slower if it is swapped frequently. In addition, overcommitting can cause the system to run out of memory (OOM), which may lead to the Linux kernel shutting down important system processes. If you decide to overcommit memory, ensure sufficient testing is performed.

## 7.3. Overcommitting virtualized CPUs

The KVM hypervisor supports overcommitting virtualized CPUs (vCPUs). Virtualized CPUs can be overcommitted as far as load limits of guest virtual machines allow. Use caution when overcommitting vCPUs, as loads near 100% may cause dropped requests or unusable response times.

In Red Hat Enterprise Linux 7, it is possible to overcommit guests with more than one vCPU, known as symmetric multiprocessing (SMP) virtual machines. However, you may experience performance deterioration when running more cores on the virtual machine than are present on your physical CPU.

For example, a virtual machine with four vCPUs should not be run on a host machine with a dual core processor, but on a quad core host. Overcommitting SMP virtual machines beyond the physical number of processing cores causes significant performance degradation, due to programs getting less CPU time than required. In addition, it is not recommended to have more than 10 total allocated vCPUs per physical processor core. 

...

# Chapter 8. KVM Guest Timing Management

# Chapter 13. Managing Storage For Virtual Machines

The storage is attached to the virtual machine using paravirtualized or emulated block device drivers.

## 13.1. Storage Concepts

A storage pool is a quantity of storage set aside for use by guest vm. Storage pools are divided into storage volumes. Each storage volume is assigned to a guest vm as a block device on a guest bus.

Storage pools and volumes are managed using libvirt. With libvirt's remote protocol, it is possible to manage all aspects of a guest virtual machine's life cycle, as well as the configuration of the resources a management application, such as the Virtual Machine Manager, using libvirt can enable a user to perform all the required tasks for configuration the host physical machine for a guest virtual machine.

## 13.2. Using storage pools

### 13.2.1. Storage Pool Concepts

A storage pool is a file, directory, or storage device, managed by libvirt to provide storage to virtual machines. Storage pools are divided into storage volumes that store virtual machine images or are attached to virtual machines as additional storage. Multiple guests can share the same storage pool, allowing for better allocation of storage resources.

**Local storage pools**

Local storage pools are attached directly to the host server. They include local directories, directly attached disks, physical partitions, and LVM volume groups on local devices. Local storage pools are useful for development, testing, and small deployments that do not require migration or large numbers of virtual machines. Local storage pools may not be suitable for many production environments, because they cannot be used for live migration.

**Network (shared) storage pools**

Networked storage pools is include storage devices shared over a network using standard protocols. Networked storage is required when migrating virtual machines between hosts with virt-manager but it optional when migrating with **virsh**.

For more information on migrating virtual machines.

### 13.2.2. Creating Storage Pools

This section provides general instructions for creating storage pools using **virsh** and the **Virtual Machine Manager**. Using **virsh** enables you to specify all parameters, whereas using VMM.

## 13.3. Using Storage Volumes

### 13.3.1. Storage Volume Concepts

# Chapter 23. Manipulating The Domain XML

This chapter explains in detail the components of guest vm XML configuration files, also known as domain XML. In this chapter, the term domain refers to the root `<domain>` element required for all guest virtual machines. The domain XML has two attributes: type and id.type specifies the hypervisor used for running the domain. The allowed values are driver-specific, but include KVM and others. id is a unique integer identifier for the running guest virtual machine. Inactive machines have no id value. The sections in this chapter will describe the components of the domain XML. Addtional chapters in this manual may see this chapter when manipulation of the domain XML is required.

## 23.1. General Information And Metadata

This information is in this part of the domain XML:

```xml
<domain type='kvm' id='3'>
    <name>fv0</name>
    <uuid>4dea22b31d52d8f32516782e98ab3fa0</uuid>
    <title>A short description - title - of the domain</title>
    <description>A human readable description</description>
    <metadata>
        <app1:foo>
        <app2:bar>
    </metadata>
</domain>
```

|Element|Description|
|-|-|
|<name>|Assigns a name for the vm. This name should consist only of alpha-numeric characters and is required to be unique within the scope of a single host physical machine. It is often used to form the file name for storing the persistent configuration files.|
|<uuid>|Assigns a globally unique identifier for the virtual machine. The format must be RFC 4122-compliant, for example 3e3fce45-4f53-4fa7-bb32-
11f34168b82b. If omitted when defining or creating a new machine, a random UUID is generated. It is also possible to provide the UUID using a sysinfo specification.|
|<title>|Create space for a short description of the domain. The title should not contain any new lines.|
|<description>|Different from the title, this data is not used by libvirt.  It can contain any information the user choose to display|
|<metadata>|-|

## 23.2. Operating System Booting

There are a number of different ways to boot vms, including BIOS boot loader, host physical machine bool loader, direct kernel boot and container boot.

### 23.2.1. BIOS Boot Loader

Bootimg the BIOS is available for hypervisors supporting full virtualization. In this case, the BIOS has a boot order priority (floppy, hard disk, CD-ROM, network) determining where to locate the boot image. The `<os>` section of the domain XML contains the following information:

```xml
<os>
    <type>hvm</type>
    <boot dev='fd'/>
    <boot dev='hd'/>
    <boot dev='cdrom'/>
    <boot dev='network'/>
    <boot dev='yes'/>
    <smbios dev='sysinfo'/>
</os>
```

|Element|Description|
|-|-|
|<type>|Specifies the type of operating systems to be booted on the guest virtual machine. **hvm** indicates that the operating system is designed to run on bare metal and requires full virtualization. linux refers to an operating system that supports the KVM hypervisor guest ABI.|
|<boot>|Specifies the next boot device to consider with one of the following values: fd, hd, cdrom or network.|
|<bootmenu>|Determines whether or not to enable an interactive boot menu prompt on guest virtual machine start up. The enable attribute can be either yes or no. If not specified, the hypervisor default is used.|
|<smbios>|-|
|<bios>|-|

### 23.2.2. Direct Kernel Boot

### 23.2.3. Container Boot

## 23.3. SMBIOS system information

## 23.4. CPU Allocation

```xml
<domain>
    <vcpu placement='static'>2</vcpu>
</domain>
```

The `<vcpu>` element defines the maximum number of virtual CPUs allocated for the guest virtual machine operating system, which must be between 1 and the maximum number supported by the hypervisor. This element can contain an optional `cpuset` attribute, which is a common-separated list of physical CPU numbers that the domain process and virtual CPUs can be pinned to by default.

The <vcpu> element defines the maximum number of virtual CPUs allocated for the guest virtual machine operating system, which must be between 1 and the maximum number supported by the hypervisor.

**BIOS/UEFI Boot Order**

In a traditional BIOS or UEFI setup (for physical machines), you can specify the boot order to control the sequence of boot devices. Here's how it generally works:

- Floppy Device (`fd`): This option will try to boot from a floppy disk drive if available. It's a legacy option not commonly used in modern systems.
- Hard Disk (`hd`): Specifies botoing from one or more hard disk drives. This is the most common boot device for everyday use.
- Option Drive (`cdrom`): This allows the system to boot from an optical disc like a CD or DVD. Useful for installations or running live versions of operating systems.
- Network (`network`): Also known as PXE boot, this option allows the machine to boot from a network server using a network interface controller. This is commonly used in managed environments for installations and deployments.

## 23.18. Storage Pools

Although all storage pool back- ends share the same public APIs and XML format, they have varying levels of capabilities. Some may allow creation of volumes, others may only allow use of pre-existing volumes. Some may have constraints on volume size, or placement. 

### 23.18.1. Providing Metadata for the Storage Pool.

```xml
<pool type='iscsi'>
    <name>virtimages</name>
    <uuid>3e3fce45-4f53-4fa7-bb32-11f34168b82b</uuid>
    <allocation>10000000</allocation>
    <capacity>50000000</capacity>
    <available>40000000</available>
 </pool>
```

**Figure 23.79. General metadata tags**

The elements that are used in this example are explained:  

|Element|Description|
|-|-|
|<name>|Provides a name for the storage pool which must be unique to the host physical machine. This is mandatory when defining a storage pool.|
|<uuid>|Provides an identifier for the storage pool which must be globally unique. Although supplying the UUID is optional, if the UUId is not provided at the time the storage pool is created, a UUID will be automatically generated.|
|<allocation>|Provides the total storage allocation for the storage pool. This may be larger than the sum of the total allocation across all storage volumes due to the metadata overhead. This value is expressed in bytes. This element is read-only and the value should not be changed.|
|<capacity>|Provides the total storage capacity for the pool. Due to underlying device constraints, it may not be possible to use the full capacity for storage for storage volumes. This values is in bytes.|
|<available>|Provides the free space available for allocating new storage volumes in the storage pool. Due to underlying device constraints, it may not be possible to allocate the all of the free space to a single storage volume. This value is in bytes. This element is read-only and the value should not be changed.|

## 23.19. Storage Volumes

A storage volume will generally be either a file or a device node; since 1.2.0, an optional output-only attribute type lists the actual type (file, block, dir, network, or netdir).

### 23.19.1. General Metadata

The top section of the `<volume>` element contains information known as metadata as shown in this XML example:

```xml
<volume type='file'>
    <name>sparse.img</name>
    <key>/var/lib/libvirt/images/sparse.img</key>
    <allocation>0</allocation>
    <capacity unit="T">1</capacity>
</volume>
```

|Element|Description|
|-|-|
|<name>|Provides a name for the storage volume which is unique to the storage pool. This is mandatory when defining a storage volume|
|<key>|Provides an identifier for the storage volume which identifies in a single storage volume. In some cases it is possible to have two distinct keys identifying a single storage volume. This field cannot be set when creating a storage volume as it is always generated.|

## 23.21. A Sample Virtual Machine XML configuration

The following table shows a sample XML configuration of a guest virtual machine (VM), also referred to as domain XML, and explains the content of the configuration.

**A Sample Domain XML Configuration**

```xml
<domain type='kvm'>
    <name>TestGuest1</name>
    <uuid>ec6fbaa1-3eb4-49da-bf61-bb02fbec4967</uuid>
    <memory unit='KiB'>1048576</memory>
    <currentMemory unit='KiB'>1048576</currentMemory>
    <vcpu placement='static'>1048576</vcpu>
</domain>
```

This is a KVM called TestGuest1 with 1024 MiB allocated RAM.

```xml
<vcpu placement='static'>1</vcpu>
```

The guest VM has 1 allocated vCPU.

```xml
<os>
    <type arch='x86_64' machine='pc-i440fx-2.9'>hvm</type>
    <boot dev='hd'/>
</os>
```

The machine architecture is set to AMD64 and Intel 64 architecture, and uses the Intel 440FX machine type to determine feature compatibility. The OS is booted from the hard drive. 

```xml
<features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
</features>
```

In the context of vm configuration, particularly when using `libvirt` to manage VMs with KVM/QEMU, the `<features>` section of the VM's XML configuration file plays a crucial role in defining specific hardware features and their behaviors. The elements within `<features>` control advanced hardware settings that affect the functionality and performance of the virtual machine.

