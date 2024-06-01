## Topologies

The layout of how a network communicates with different devices is called a topology.

## Connector

## Firewall

Can be either software or hardware.

Designed to prevent unauthorized access from entering a private networking.

Blocks unwanted traffic and permits wanted traffic.

**Stateful vs Stateless**

Firewalls inspect traffic in a couple of different ways

Stateful: Monitors all the connections and data streams that are passing through.

Stateless: Uses an ACL to allow or deny traffic. Only look into header.

## Network components

**Hub**

A hub is a device that has multiple ports that accepts ethernet connections from network devices. A hub is consider not to be intelligent because it does not filter any data or has any intelligence as to where data is supposed to be sent. When a data packet arrives at on of these ports, it is copied to all other ports, so all the devices on that hub sees that data packet. There are also two different types of hubs: passive and active. 

**Switch**

A switch is very similar to a hub. It's also a device that has multiple ports that accepts ethernet connections from network devices. But unlike a hub, a switch is intelligent. A switch can actually learn the physical addresses of the devices that are connected to it and it stores these addresses in a table. So when a data packet is sent to a switch, its directed only to the intended destination port. That's the major difference between a hub and a switch.

So as a result, switches are far more preferred over hubs, because they reduce any unnecessary traffic on the network.

|Hub|Switch|
|-|-|
|Layer 1 device, as per OSI model|Layer 2 device, as per OSI model|
|Has no memory|Has memory & stores MAC address table|
|Not an intelligent device|Intelligent device|
|Flood the networkd due to broadcasting|Do unicasting, multicasting  broadcasting|
|High Security risk device|Low security risk device|
|Less efficient| High efficient|
|Half Duplex|Full Duplex|

**Multilayer switch**

- Operates at layers 2 and 3 of the OSI model.
- Interprets layer 3 data similar to a router.

**Content switch**

- Can operate at layers  through 7 of the OSI model.
- Performs load balancing and filtering.

**Spanning Tree (protocol)**

The Spanning Tree Protocol allows for fault tolerance and prevents unnecessary traffic loops in the network.

Allows the switches to talk to each other to find in loops are happening in the network.

**Bridges**

Bridges are used to divide a network into separate collision domains.

A bridge to this network, it will reduce any unnecessary traffic between the two segments by filtering the data based on their MAC address. The bridge only allows data to crossover if it meets the required MAC address of the destination. Because a bridge keeps a record of all the MAC addresses of the NICs that are connected to it, and it will also block all data from crossing over if it fails to meet this requirement.

**Routers**

A router does exactly what its name implies. A router is a device that routes or forwards data from one network to another based on their IP address. When a data packet is received from the router, the router inspects the packet and determines if the packet was meant for its own network or if it's meant for another network. If the router determines that the data packet is meant for its own network, it receives it. But if it's not meant for its own network, it sends it off to another network. So a router is essentially the gateway for a network.

**Gateway**

A gateway can be defined as a device that joins two networks together. The interconnect networks with different or incompatible communication protocols. A gateway however does not change the data, it only changes the format of the data. So in simple terms, this illustration is showing how a gateway is transforming a simple signal into something totally different.

**NICs**

 Convert incoming serial data into parallel data (bit).

A NIC provides a constant dedicated connection to a network. NIC has an unique identifier, called a MAC address.

**Transceivers**

- A transceiver is a device that has both a transimitter and a receiver in the same package.

**WAPs (wireless access point)**

is a wireless hub that is used by wireless devices. It connected to a wired network and relays data between the wired network and the wireless device for communication purposes.

## MAC address

- A MAC address uniquely identifies each device on a network.
- So no two devices anywhere in the world have the same MAC address.
- The first 3 bytes of the MAC address identifies the manufacturer of the NIC.
- The last 3 bytes are unique number from that manufacturer.

## IP address

- A numeric address.
- It's an identifier for a computer or device on a network.
- Every device on a network to have an ip address to communicate.
- Consists of 2 parts: A network address and a host address.
- 2 types: IPv4 and IPv6

**Subnet MASK**

- IP address consists of 2 parts: A network address and a host address

**IP classes and ranges**

|Class|First Octet Address|Default Subnet Mask|
|-|-|-|
|A|1-126|255.0.0.0|
|B|128-191|255.255.0.0|
|C|192-223|255.255.255.0|

127 is reserved for loopback testing.

**Private IP**

- Not publicly registered.
- Cannot directly access the internet.

Example, let's say you have a small business and you need 10 public IP addresses so your employees can access the internet. Now you could contact your ISP and ask them for these additional IP addresses, but that would be very expensive and unnecessary. So that's where private IP addressing comes in. In private IP addressing, you can create these ten private IP addresses and just have one publicly registered IP address from your ISP. These ten private IPs would then be translated into the one public IP, so your employees can have access to the internet. This not only sasves money but it also helps prevent having a shortage of public IP addresses.

|Class|IP Range|Subnet Mask|
|-|-|-|
|A|10.0.0.0-10.255.255.255|255.0.0.0|
|B|172.1.0.0-172.31.255.255|255.255.0.0|
|C|192.168.0.0-192.168.255.255|255.255.255.0|

## Subnetting

The word subnet is short for subnetwork. Which means a smaller network within a larger one. Subnetting is basically breaking down a large network into smaller networks or subnets. Subnetting is basically breaking down a large network into smaller networks or subnets. It's mainly done to make your network more manageable. So for example, let's say you have a company with 3000 employees and yours ISP assigned you with a Class B IP address with a default subnet mask. So as you know from the previous lession, a Class B IP address will allow you approximately 65000 IPs for all your computers. Now you could all of your employees in one large network, and if you had a small business then this would be fine. But if you had a fairly large business, with for example 3000 computers, then this could be a problem because of traffic issues caused by so many broadcasts. And if a problem word to occur, it'll be very hard to pinpoint on one large network. Or in another scenario what if your business was scattered into three differnet geographical locations, then this would also be a problem. So a better way would be to break down your network into smaller ones or subnets.

Subnetting is basically done by changing the default subnet mask by borrowing some of the bits that were designated for hosts and using them to create subnets. So a default Class B subnet mask is 255.255.0.0. The first two octets are for the network, the last 2 octets for the network, the last 2 octets are designated for hosts. So let's say  we want to break down this network into three small ones. The fomual we would use is 2 to the n power-2: 2^n - 2, where n equals to the number of bits we need to borrow from the hsot portiion of the subnet mask. At least 3 subnets or larger.

2^n - 2

n = 2 -> 2^2 - 2 = 2 < 3 => fail
n = 3 -> 2^3 - 2 = 6 > 3 => OK

We need to borrow 3 bits from the hosts portition of the subnet mask to break up our network.

## IP Addressing Methods

Every computer on a network has to have an IP address for communication purposes and there are two ways that a computer can be assigned an IP address. It could be done either by using a dynamic IP, or a static IP. By using a dynamic IP, or a static IP. A dynamic IP is where a computer gets an IP address automatically from a DHCP server. DHCP stands for dynamic host configuration protocol.

## TCP (Transmission Control Protocol)

1. SYN -> 2. SYN ACK -> 3. ACK RECEIVED

Connection oriented protocol

## UDP (User Datagram Protocol)

A connectionless oriented protocol

## ARP (Address Resolution Protocol)

- Used to resolve IP addresses to MAC addresses.
- Computers use MAC addresses for communication between each other.
- Computers search their ARP cache first to find the target MAC address.
- If the MAC address is not in the ARP cache, the computer will broadcast a message asking for it.

## RARP (Reverse Address Resolution Protocol)

- RARP is just the opposite of ARP.
- RARP resolves MAC addresses to IP addresses.

## SCP (Secure Copy Protocol)

Uses secure shell to safeguard data as it's being trasferred over a network

## Ports

When data is sent over the internet to your computer, it needs to know how to accept it, and your computer accepts this data using ports, and these ports are categorized by 2 protocols, TCP and UDP.

Now a port is a logical connection that is used by programs to exchange information.

Ports are identified by a unique number.

Port number ranges from 0 to 65535.

## Networking service

### DNS

### NAT (Network Address Translation)

Translates a set of IP addresses to another set of IP addresses. 

- Private IP addresses to public.
- Public IP addresses to private.

Another version of network address translation is called PAT, which stands for port address translation

|IP|PORT|IP ADDRESS|
|-|-|-|
|67.158.212.121|4001|10.10.0.1|
|67.158.212.121|4002|10.10.0.2|
|67.158.212.121|4003|10.10.0.3|

Each computer in a private network is issued not only a unique IP address, but they are also issued a unique port number. This is done so that external data packets from the internet knows which computer on the private network it wants to talk to. So for example, if a device outside this network wanted to communicate with a computer on this private network, the IP address along with its port number would be translated by PAT to find the correct computer.

### SNAT (Static Network Address Translation)

SNAT stands for a static network address translation.

NAT translates a private network's IP address to a public IP address.

### Proxy service

### RDP (Remote Desktop Protocol)

### Broadcast

A single transmitter of data being received by multiple receivers.

### Unicast vs Multicast

Unicast - sent to a single destination.

Multicast - sent to multiple destinations at the same time.

## Routing protocols

### Loopback interface

A fake or virtual interface that is created on a router. Its not a physical interface, it's virtual, and this virtual interface is assigned an IP address of your choice.

Used for testing and administration purposes.

### Routing Table

A routing table is a file that contains set of rules that shows information on what path a data packet taks to its destination.

So as a data packet arrives at the router, the router looks at its routing table to find out where to forward the data packet along the best path to its destination.

So a basic routing table contains a network destination, which is an IP address of the final destination.

A network destination: The IP address of the final destination

Subnet mask: Determines which part of the IP address is the host and network portion.

Gateway: Tells the router which IP address the data packet should be forwarded to.

Interface: The outgoing IP address of the device that's sending the data.

Next hop: The IP address to which the IP address is forwarded to.

Metric: Determines the best route among multiple destinations.

### Routing protocols

TODO: READ again

Routing protocols collect information about the current network status and map out the best path for data packets to take to their specific destination.

**3 types of Routine Protocols**

- Distance Vector
- Link State
- Hybrid

## Wan Technologies

## Cloud & Virtualization


Ref: 

1. https://www.youtube.com/watch?v=VwN91x5i25g&list=PLBlnK6fEyqRgneraVKkEXrwyLVx2vJUvt
2. 
