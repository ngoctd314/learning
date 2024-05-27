# Libvirt

The libvirt library is used to interface with many different virtualisation technologies. Before getting started with libvirt it is best to make sure your hardware supports the necessary virtualisation extensions for KVM. To check this, enter the following a terminal prompt: 

```txt
kvm-ok
```

A message will be printed informing you if your CPU does not support hardware virtualisation.

**Note:**

On many computers with processors supporting hardware-assisted virtualisation, it is necessary to first activate an option in the BIOS to enable it.

## Virtual networking

There are a few different ways to allow a virtual machine access to the external network.


