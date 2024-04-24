# Chapter 4. Libvirt Networking

Understanding how virtual networking works is really eassential for virtualization. It would be very hard to justify the costs associated with a scenario in which we didn't have virtual networking. Just imagine having multiple virtual machines on a virtualization host and buying network cards so that every single one of those virtual machines can have their own dedicated, physical network port. By implementing virtual networking, we're also consolidating networking in a much more manageable way, both from an administration and cost perspective.

This chapter provides you with an insight into the overall concept of virtualized networking and Linux-based networking concepts. We will also discuss physical and virtual networking concepts, try to compare them, and find similarities and differences between them. Also covered in this chapter is the concept of virtual switching for a per-host concept and spanned-across-hosts concept, as well as some more advanced topics. These topics include single-root input/output virtualization, which allows for a much more direct approach to hardware for certain scenarios. We will come back to some of the networking concepts later in this book as we start discussing cloud overlay networks. This is because the basic networking concepts aren't scalable enough for cloud environments.

- Understanding physical and virtual networking
- Using TAP/TUN
- Implementing Linux bridging
- Configuring Open vSwitch
- Understanding and configuring SR-IOV
- Understanding macvtap
- Let's get started!

## Understanding physical and virtual networking

Let's think about networking for a second. This is a subject that most system administrators nowadays understand pretty well. This might not up to the level many of us think we do, but still - if we were to try to find an area of system administration where we'd find the biggest common level of knowledge, it would be networking.

If we really understand physical networking, virtual networking is going to be a piece of cake for us. Spoiler alert: it's the same thing. If we don't, it's going to be exposed rather quickly.

A virtual switch is basically a software-based Layer 2 switch that you use to do two things:

- Hook up your virtual machines to it.
- Use its uplinks to connect them to physical server cards so that you can hook these physical network cards to a physical switch.

So, let's deal with why we need these virtual switches from the virtual machine perspective. As we mentioned earlier, we use a virtual switch to connect virtual machines to it. Why? Well, if we didn't have some kind of software object that sits in-between our physical network card and our virtual machine, we'd have a big problem - we would only connect virtual machines for which we have physical network ports to out physical network, and that would be intolerable. Imagine having 20 virtual machines on your server. This means that without a virtual switch, you'd have to have at least 20 physical network ports to connect to the physical network. On top of that, you'd actually use 20 physical ports on your physical switch as well, which would be a disaster.  

So, by introducing a virtual switch between a virtual machine and a physical network port, we're solving two problems at the same time - we're reducing the number of physical network adapters that we need per server, and we're reducing the number of physical switch ports that we need to use per server.

## Virtual networking

In order for that virtual switch to be able to connect to something on a virtual machine, we have to have an object to connect to - and that object is called a virtual network interface card, ofter referred to as a vNIC. Every time you configure a virtual machine with a virtual network card, you're giving it the ability to connect to a virtual switch that uses a physical network card as an uplink to a physical switch.

Of course, there are some potential drawbacks to this approach. For example, if you have 50 virtual machines connected to the same virtual switch that uses the same physical network card as an uplink and that uplink fails (due to a network card issue, cable issue, switch port issue, or switch issue), your 50 virtual machines won't have access to physical network. How do we get around this problem? By implementing a better desing and following the basic desing principles that we'd use on a physical network as well. Specifically, we'd use more than one physical uplink to the same virtual switch.

Linux has a lot of different types of networking interfaces, something like 20 different types, some of which are as follows:

- Bridge: Layer 2 interface for (virtual machine) networking.
- Bond: For combining network interfaces to a single interface (for balancing and failover reasons) into one logical interface.
- Team: Different to bonding, teaming doesn't create one logical interface, but can still do balancing and failover.
- MACVLAN: Creates multiple MAC addresses on a single physical interface (create subinterfaces) on Layer 2.
- IPVLAN: Unlike MACVLAN, IPVLAN uses the same MAC address and multiplexes on Layer 3.

...


