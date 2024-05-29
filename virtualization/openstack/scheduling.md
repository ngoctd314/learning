# Scheduling

Compute uses the nova-scheduler service to determine how to dispatch compute requests. For example, the nova-scheduler service determines on which hosts a VM should launch. In the context of filters, the term host means a physical node that has a nova-compute service running on it. You can configure the scheduler through a variety of options.

By default, the scheduler_driver is configured as a filter scheduler, as described in the next section. In the default configuration, this scheduler considers hosts that meet all the following criteria:

- Have not been attempted for scheduling purposes (RetryFilter).
- Are in the requested availability zone (AvailabilityZoneFilter).
- Have sufficient RAM available (RamFilter).
- Have sufficient disk space available for root and ephemeral storage (DiskFilter).
- Can service the request (ComputeFilter).
- Satisfy the extra specs associated with the instance type (ComputeCa..)
- Satisfy any architecture, hypervisor type, or virtual machine mode properties specified on the instance's image properties 
...

The scheduler caches its list of available hosts; use the scheduler_driver_task_period option to specify how often the list is updated.

## Filter scheduler

The filter scheduler

(nova.scheduler.filter_scheduler.FilterScheduler) is the default scheduler for scheduling virtual machine instances. It supports filtering and weighting to make informed decisions on where a new instance should be created. 

### Filters

When the filter scheduler receives a request fro a resource, it first applies filters to determine which hosts are eligible for consideration when dispatching a resource. Filters are binary: either a host is accepted by the filter, or it is rejected. Hosts that are accepted by the filter are then processed by a different algorithm to decide which hosts to use for that request, described in the Weights section.

{1,2,3,4,5,6,7,8} -> filter -> {1,2,3,5,6} -> weighting -> {1,3,5}

**Filtering**

The scheduler_available_filters configuration option in nova.conf provides the Compute service with the list of the filters that are used by the scheduler. The default setting specifies all of the filter that are included with the Compute service:

```txt
scheduler_available_filters = nova.scheduler.filters.all_filters
```

**AggregateCoreFilter**

Filters host by CPU core numbers with a per-aggregate cpu_allocation_ratio value. 

**AggregateDiskFilter**

Fitlers host by disk configuration

### Weights

When resoucing instances, the filter scheduler filters and weights each  host in the list of acceptable hosts. Each time the scheduler selects a host, it virtually consumes resources on it, and subsequent selections are adjusted accordingly. This process is useful when the customer asks for the same large amount of instances, because weight is computed for each requested instance.

All weights are normalized before being summed up; the host with the largest weight is given the highest priority.

|Section|Option|Description|
|-|-|-|
|[DEFAULT]|ram_weight_multiplier|By default, the scheduler spreads instances across all hosts evenly. Set the ram_weight_multiplier option to a negative number if you prefer stacking instead of spreading. Use a floating-point value|
|[DEFAULT]|scheduler_host_subset_size|New instances are scheduled on a host that  is chosen randomly from a subsets of the N best hosts. This property defines the subset size ..|

...

