Scheduling works like so:

1. Scheduler gets a request spec from the super conductor, containing resource requirements. The super conductor operates at the top level of a deployment, as constrated with the cell conductor, which operates within a particular cell. 

2. Scheduler sends those requirements to placement.

3. Placement runs a query to determine the resource providers (in this case, compute nodes) that can satify those requirements.

4. Placement then constructs a data structure for each compute node as documented in the spec. The data structure contains summaries of the matching resource provider information for each compute node, along with the AllocationRequest that will be used to claim the requested resources if that compute node is selected. 

5. Placement returens this data structure to the Scheduler.

6. The Scheduler creates HostState objects for each compute node contained in the provider sumaries. These HostState objects contain the information about the host that will be used for sub-sequent filtering and weighing.

7. Since the request spec can specify one or more instances to be scheduled.
