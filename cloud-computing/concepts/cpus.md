# Basic concept about CPUs

## 64-bit CPUs

In computing, processors (CPUs) come in different architectures, including 32-bit and 64-bit. The transition to 64-bit CPUs took some time, with 64-bit architectures becoming more prevalent in the mid-2000s. Prior to this, most CPUs were 32-bit, meaning they processed data and memory addresses in 32-bit chunks.

## Hardware-assisted virtualization

Virtualization technology allows multiple virtual machines (VMs) to run on a single physical machine, each operating independently. Hardware-assisted virtualization, also known as hardware virtualization extensions, refers to features built into modern CPUs that enhance the performance and security of virtualization. These features help virtualization software (hypervisors) efficiently manage resources and isolate VMs from each other and the host system.

## 32-bit architectures and virtualization

Before the widespread adoption of 64-bits CPUs, there was little interest in running Hardware-assisted virtualization on 32-bit architectures. One primary reason for this was the limitation on memory addressing memory is 4GB (2^32). This limitation severely constrained the amount of memory that could be allocated to each virtual machine. 

## Memory limitation

The 4 GB memory limit posed a significant obstacle to virtualization. When running multiple VMs on a host system, each VM requires its own portition of memory. With the 4 GB limit, it was challenging to allocate sufficient memory to each VM, especially for resource-intensive applications. As a result, the scope and practically of using virtualization in 32-bit env were severely limited.


