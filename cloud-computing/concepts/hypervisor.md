# Hypervisor

A hypervisor is what allows one machine to run multiple virtual machines. It's what allocates and controls the sharing of a machine's hardware.

## VM Disadvantage

- Consume a lot of disk space.
- Consume a lot of RAM and CPU power from the server.
- Slow to startup.
- Requires a license for each operating system.

## Question

1. Is any hardware can run hypervisor?


Most modern hardware is capable of running a hypervisor, but there are certain considerations to keep in mind:

- Virtualization Support: Hardware virtualization support is a key requirement for running a hypervisor efficiently. Most modern CPUs have hardware support for virtualization, such as Intel's VT-x (Virtualization Technology) or AMD's AMD-V (AMD virtualization). These features allow the hypervisor to run mutliple virtual machines with better performance and security.
- BIOS/UEFI Settings: Sometimes, you might need to enable virtualization support in the BIOS or UEFI firmware settings of your computer. This is often disabled by default for security reasons, so you may need to manually enable it.
- CPU Architecture: Different hypervisors have different requirements for CPU architecture. For example, some hypervisors may require a 64-bit CPU with specific virtualization extensions.
- Memory and Storage: Sufficient RAM and storage are also important for running a hypervisor and the virtual machines it hosts. Make sure your hardware meets the minimum requirements for the hypervisor and any virtual machines you plan to run. 

Red Hat switched to KVM in Red Hat Enterprise Linux 6 in 2010.

