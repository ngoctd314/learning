# Network Tutorial

## What is VLAN and Why VLAN

Reference: https://www.youtube.com/watch?v=_PPaArOxHhw&list=PLSNNzog5eydurp2zcB4xs6gdeeVW3cMOW

VLAN - Virtual Local Area Network.
VLAN is a logical network that can gropu devices/users regardless of their different physical locations.
A VLAN is created at the switch.

**Why VLAN**

1. Segmentation

We have 2 departments A and B. Both of them are connected to one switch. It means both departments belong to one network or one broadcast domain. One broadcast domain means when one computer broadcasts, everyone else can hear. This type of network design could cause two potential problems, at least.

+ Traffic congestion and poor performance.
+ Security concerns.

To address these two problems, here are three possible solutions:

+ We can buy more switches and routers.
+ We can create two VLANs with one switch like this. Physically these two departments share one physical switch but virtually they are in two different independent VLANs. We can also add one router so that two VLANs can talk to each other.
+ We can use a multilayer switch to create. We don't need a physical router in between. Instead, these two VLANs can talk to each other using inter-VLAN routing techniques.

2. The simplicity of network design and deployment

With VLANs, you're no longer confined to physical locations. Users in the same department can be located in different buildings or on differnet floors.

## VPN - Virtual Private Network

## VLAN: Static vs Dynamic

A static VLAN is a "port-based VLAN". This type of VLAN requires manual assignment of individual ports on a switch to a virtual network. Once ports are assigned, they are always associated with their pre-assigned VLANs. Static VLAN has nothing to do with devices. When we plug one device  into port 1, for example, the device will be on the Sales VLAN, and when we plug the same device into port 9, the device will become a member of the IT vlan. 

Dynamic VLAN (MAC based) is different. We define dynamic VLAN based on a device instead of port location. Dynamic VLAN is usually MAC-based, which means we define VLAN membership according to a device's MAC address. We can define VLAN membership based on the IP address of a device if we want to. Dynamic VLAN requires a central server called  the VLAN Member Policy Server, or simply VMPS server.

|MAC address|VLAN|
|-|-|
|AA|VLAN 10|
|BB|VLAN 10|
|CC|VLAN 20|
|DD|VLAN 30|
|EE|VLAN 30|

When a computer with the MAC address AA, is plugged into port 1, port 1 will communicate with the server and check against the database for its membership. In the database, computer AA is pre-assigned to VLAN 10. Suppose computer AA is plugged into port 5. The result would be the same. Port 5 will communicate with the server and find the pre-assigned VLAN to the computer AA. Thus, changing port location does not change a device's VLAN membership.

We can control devices on a central server instead of configuration each individual port of all switches. For a large organization, dynamic VLAN could simplify the network management.

No random device could be able to plug into a switch and get access to the network.

**In Summary**

Static VLAN requires manual assignment of ports to a VLAN. One port belong to one VLAN.

Dynamic VLAN is based on some characteristics of devices, normally MAC address. One port can be a member of multiple VLANs. Dynamic VLAN is flexible but more complex and many adminstrators still prefer static VLANs.

## IEEE 802.1q standard and VLAN tagging and trunking basics

Dot1q, VLAN tag, trunk, trunking, trunk port, and access port.

IEEE 802.1q, often referred to as Dot1q, is the networking standard that supports VLANs on the Ethernet network.

The standard defines a method of tagging traffic between two switches to tell which traffic belongs to which VLAN.

VLANs are local to each switch and VLAN info is not passed between switches. We need to link these two switches so that computers on the same VLAN could communicate. This link is called a trunk. A trunk or called Dot1q link, or simple trunk link, provides VLAN ID for frames traversing between switches.

A trunk can be configured between 2 switches or between a switch and a router. By default, a trunk can carry traffic from all VLANs to and from the switch. But can be configured to carry only specified VLAN traffic. The process of traversing different VLAN traffic over the trunk is called trunking.

The reason that a single link could become a trunk is not because of some special physical cable, but simple because of the configuration of switch ports on both ends of the link: trunk port. Trunk port is a Cisco term. In the non-Cisco world, it is called "tagged port". A trunk port would add a VLAN tag to an Ethernet frame to indicate what VLAN the frame belongs to.

Access Port / Untagged Port: Is a switch port that sends and expects to receive traffic with no VLAN tag. In other words, an access port carries traffic only for one VLAN. To make different VLAN traffic travel across two switches, we need a trunk. A trunk is created by configuration trunk ports. A trunk port sends or expects to receive tagged traffic. An access port sends or expects to receive untagged traffic.

A computer attached to an access port has no idea about VLAN and even its VLAN belonging. VLAN creation and management are switch's responsibilities.

When switch 1 gets the Ethernet frame from computer A, it knows the frame comes from a member of VLAN 10 and goes to another member of VLAN 10 on switch 2. Thus it needs to send the frame to switch 2 over the trunk.

For that purpose, switch 1 does one thing on its trunk port: inserting a VLAN tag into the Ethernet frame (FCS also changes due to the insertion).

The tag consists of 4 fields. x-x-C-VLAN identifier, which can be a number from 1 to 4094.

The process of adding VLAN tag into an Ethernet frame is called tagging or encapsulation. They mean the same thing.

The trunk in multiple-lane highway. Each lane is for each different VLAN traffic. When switch 2 gets the frame on its trunk port, it knows from the tag  that the frame should be forwarded to computer C on VLAN 10. Remmeber, an access port expects untagged frames. Thus switch 22 throws away the VLAN tag and delivers the original Ethernet frame to computer C.

**Summary**

## Default VLAN and Native VLAN

The default VLAN is VLAN 1, a default setting on Cisco switches and those of most other vendors. Unless we specifically assign an access port to a pariticular VLAN such as VLAN 10 or VLAN 10 the access port belongs to VLAN 1.

1. We cannot change the default VLAN. We cannot even delete the default VLAN. It is the default setting. It is VLAN 1.
2. VLAN 1 is never intended to be used as  a standard data VLAN. 

A native VLAN is a special VLAN whose traffic traverses on the 802.1q trunk without a VLAN tag. The native VLAN and management VLAN could be the same, but a good security pratice is they are not.

By default the native VLAN is VLAN 1. By we can change it to any number, such as VLAN 2, or 20, or 99 or whatever you like. It can be configured on the trunk port.

If computer E and computer F are not assigned to any VLAN, and thus they belong to VLAN 1 by default. When computer E sends a frame to computer F, it can travel over the trunk without any VLAN tag.

The native VLAN is a per trunk per switch configuration. A best security practice is to change the native VLAN to a different VLAN other than VLAN 1. The native VLAN should be the same on both ends of the trunk.

**Why Native VLAN?**

Native VLAN is one concept defined in 802.1q standard that was created for backward compatibility with old devices that don't support VLANs.

On an Ethernet network, all devices on the link must still be capable of communicating even if they do not speak the 802.1q protocol.

The native VLAN is used by the switch to carry specific control and management protocol traffic like CDP.

**What is the purpose?**

## InterVLAN Routing: 3 options

As we know, each VLAN is a different and isolated broadcast domain. Hosts on separate VLANs aren't able to communicate even if they are connected to the same switch.

However, communications between different VLANs are necessary in many cases. In other words, we must find ways to route the traffic between different VLANs. This is known as interVLAN routing.

3 options of InterVLAN routing:

1. Traditional: Using a router

We need a physical router and connect each VLAN to a different physical router interface. Each VLAN has its own physical link connected to each different physical router interface. With this approach, the router does not need to know about VLANs, and it just treats every single VLAN as every single different physical link.

This approach has a serious limitation: the number of physical interfaces on a router are limited and expensive. We need every cable going from the switch to the router for every single VLAN. This method works find with 2 or 3 VLANs.

2. Router on-a-stick

Router on-a-stick is a setup that consists of a router and a switch, which  are connected using one Ethernet link configured as an 802.1q trunk. The router must understand the concept of VLAN and the IEEE 802.1q standard.

With option 1, each cable going from the switch to the router is configured as an access link. With Router-on-a-stick InterVLAN Routing, we configure one single link as a trunk.

Physically, we have only one interface on the router, but many logical subinterfaces are created one subinterface per VLAN. Each subinterace is configured as a default gateway for each VLAN.

When computer 1 on Sales VLAN sends IP packet to computer 2 on Finance VLAN, the switch would deliver it to the Sales VLAN gateway, a logical subinterface on the router. Then the router checks its routing table, and forwards the IP packet to the computer 2 through the gateway of the Finance VLAN.

Physically there is only one link, but logically there are two separate links in this example.

3. Multilayer switch InterVLAN Routing

With the router-on-a-stick approach, we need a router and we need a switch but with a multilayer switch or sometimes we call a layer-3 switch.

Switching and routing are achieved inside one box.

Multilayer switches support Switch Virtual Interfaces or SVIs. SVIs are logical interfaces that can act as gateways and perform routing.

They behave like physical interfaces on a router. They have IP addresses associated with their VLANs and they are completely virtual. Similar to a logical subinterface on the router on-a-stick, one SVI can be created for each VLAN. There is one-to-one mapping between VLANs and their SVIs.

## VLAN Trunking Protocol (VTP)

VLAN Trunking Protocol (VTP) is a Cisco proprietary protocol whose primary goal is to manage all configured VLANs consistency across a switched network.

In the next few minutes. I will talk about why VTP and how it works.

VTP domain, 3 VTP modes, configuration revision number, 3 types of VTP messages and VTP pruning.

With VTP, we can create VLANs on one switch: VTP server, and all other switches - VTP clients, with synchronize themselves. VTP centralizes VLAN management. We can add, modify, or delete a VLAN on a VTP server then the server will distribute all these configurations to other VTP clients in the same VTP domain. 

A VTP domain consists of a group of interconnected switches. All switches in a domain share VLAN configuration details. On a large network, we might have several VTP domains. A router or layer 3 switch defines the boundary of each VTP domain. When configuring VTP for the first time, we must always assign a domain name. We must configure all switches in the VTP domain with the same domain name. Until the VTP domain name is specified, VLANs cannot be created or modified on a VTP server, and VLAN configuration is not propagated over the network.

A switch can be a member of only one VTP domain at a time. There are three VTP modes: VTP server, client and transparent. 

## Virtualization: VM and Hypervisor

Virtualization is the process of using special software on a physical machine - to create virtual machines. This special software is called "hypervisor". A virtual machine (VM) is called a "guest". There are 3 important points about virtual machines:

1. We can create and run as many virtual machines as we like. RAM is almost always the main limiting factor. Together, all the virtual machines share the same resources of the host. Yet, each virtual machine works independently.

2. A virtual machine is little more than a file sitting on a hard drive, but to users, a virtual machine appears and acts no differently from a physical computer.

3. A virtual machine can be configured to use not only a different operating system, but also a differnet type of CPU storage drive, or NIC than its host.

Not let's talk about the special software that creates and runs virtual machines: hypervisor.

**Advantages**

1. Saving money
2. Simplified Management
3. Threat Isolation
4. Backup and Recovery

**Disadvantages**

1. Compromised Performance
2. Complexity
3. Risk
4. License Cost

## Virtualization: Bridged, NAT, Host-only - Virtual machine connection types

Virtual machines can communicate with each other on the host as well as other physical machines on the physical network. To do so, virtual machines need two virtual things: virtual NICs (vNICs) and virtual switches/bridges. A virtual switch is a logically defined layer-2 device that passes frames between nodes. Virtual NICs of VMs are connected to the virtual ports on the virtual switch.

The basic principle of network communications is the same, whether they are virtual of physical. Virtual switches are just like physical switches are just like physical switches. Each virtual switch creates a separate broadcast domain. To connect two broadcast domains we need a layer-3 router.

Now let's talk about VM network connection. There are basically three network connection modes when we set up a virtual NIC for a virtual machine:

1. bridged connection

In the bridged connection mode, a VM connects to the physical network directly through the host physical NIC. The host's NIC is a bridge to all these three VMs. Just like their host and other physical computers, VMs obtain IP addressing information from a DHCP server on the physical network. When connected using bridged connection mode, a VM appears to other nodes as just another computer on the network. This example shows that VMs are connected to the network through the host's only NIC. However, we can install physical NICs for each virtual machine. Each virtual NIC gets a connection to its own dedicated physical NIC. In the bridged connection mode, these VM's IP addresses are visible and directly accessible by other computers on the network, thus good candidates for these VMs can be a mail server, a file server, or a web server.

2. NAT

The second connection mode is Network Address Translation, or NAT

3. host-only
