# Placement Usage

This is a REST API stack and data model used to track resource provider inventories and usages, along with different classes of resources. For example, a resource provider can be a compute node, a shared storage pool, or an IP allocation pool. The placement service tracks the inventory and usage of each provider. For example, an instance created on a compute  node may be a consumer of resources such as RAM and CPU from a compute node resource provider, disk from an external shared storage pool resource provider, disk from an external shared storage pool resource provider and IP addresses from an external IP pool resource provider.

The types of resources consumed are tracked as classes. The service provides a set of standard resource classes (for example DISK_GB, MEMORY_MB, and VCPU) and provides the ability to define custom resource classes as needed.



The placement service enables other projects to track their own resources. Those projects can register/delete their own resources to/from placement via the placement HTTP API.

The placement service originated in the Nova. As a result much of the functionality in placement was driven by nova's requirements. However, that functionality was designed to be sufficiently generic to be used by any service that needs to manage the selection and consumption of resources.

## How Nova Users Placement

Two processes, nova-compute, and nova-scheduler, host most of nova's interaction with placement.

## Placement Usage

### Tracking Resources

The placement service enables other projects to track their own resources. Those projects can register/delete their own resources to/from placement via the placement HTTP API.

The nova-scheduler is responsible for selecting a set of suitable destination hosts for a workload. It begins by formulating a request to placement for a list of allocation candidates. 

### Modeling with Provider Trees


