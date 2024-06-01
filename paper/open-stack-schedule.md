# OpenStack scheduling
 
## International Journal of Soft Computing and Engineering

https://www.ijrte.org/wp-content/uploads/papers/v8i2/B1481078219.pdf

The NOVA Scheduling process is divided into 3 Phases that include (1) Getting the current state of all the compute nodes: it will generate a list of hosts, (2) filtering phase will generate a list of suitable hosts by applying filters, (3) weighting phase will sort the hosts according to their weighted cost scores which are given by applying some cost functions. The sorted list of hosts is candidate to fulfil the user requests.

NOVA processes the User requests sequentially on the FIFO basic. All the user requests are queued and the requests are processed as they come in. No priorities are considered by the scheduler for stacking the requests in the queue. NOVA scheduler will queue the user requests only when the user requested resources are available. A response is sent to the user saying that the requested resources are not available. The details of all the available resources are stored in NOVA-database. In open stack the resources are arranged in partitions and the user is assigned with the desired resources. 

**Support of Filters by NOVA**

Open Stack supports various kinds of scheduler filters. A specific filter can be chosen by the implementer through customization of nova.config file. Each of the filter is designed to meet certain amount scope related scheduling the resources.

**NOVA scheduler algorithms / Steps**

The filter Scheduler is the default scheduler and is used most frequently. Filter Scheduler supports filtering and weighting to make informed decisions on where a new instance should be created and the kind of partitions that must be attached to the virtual machines.

Filter Scheduler iterates over all found compute nodes, evaluating each against a set of filters. The list of resulting hosts is ordered by weights. The Scheduler then chooses hosts for the requested number of instances, choosing the most weighted hsots. For a specific filter to succeed for a specific host,  the filter matches the user request against the state of the host plus some extra magic as defined by each filter.

If the Scheduler cannot find candidates for the next instance, then the user is informed of the lack hosts to run the VM instance. The Filter Scheduler has to be quite flexible so that it can support variety of filtering and weighting strategies. The user can also develop his/her own filtering strategy any make the system follow the strategy. However many filters have been already found and implemented that suits different filtering requirements. The way the scheduling is done: 

Conduter:

+ Build request spec object 1
+ Submit request to scheduler

Scheduler:

+ Submit resource requirements to placement

Process:

+ Query to determine rp representing computer nodes to satisfy the requirements.
+ Return list of resources provideers and their corresponding allocation to scheduler.

Scheduler:

+ Create host state object for each rp returned from placement
+ Filter and weigh results 

Process:

+ Create allocations against selected compute node

Scheduler:

+ Return list of selected hosts and alternatives along with their allocations to the conducter.

Conducter:

Submit list of suitable hosts to target cell.

1. Scheduler gets a request specification from the "super conductor" containing resource requirements. The "super conductor" operates at the top level of a deployment, as contrasted with the "cell conductor", which operates within a particular cell.

2. Scheduler sends those requirements to placement.

3. Placement runs a query to determine the resource providers (in this case, compute nodes) that can satisfy those requirements.

4. Placement then constructs a data structure for each compute node as documented in the specification. The data structure contains summaries of the matching resource provider information for each compute node, along with the Allocation Request that will be used to claim the requested resouces if that compute node is selected.

5. Placement returns this data structure to the Scheduler.

6. The scheduler creates HostState objects for each compute node contained in the provider summaries. The HostState objects contain the information about the host that will be used for subsequent filtering and weighing.

7. Since the request spec can specify one or more instances to be scheduled. The Scheduler repeates the next several steps for each requested instance.

8. Scheduler runs these HostState objects through the filters and weighers to further refine and rank the hosts to match the request.

9. Scheduler then selects the HostState at the top of the ranked list, and determines it's matching AllocationRequest from the data returned by Placement.

10. If the claim is not successful, that indicates that another process has consumed those resources, and the host is no longer able to satisfy the request. In that event, the Scheduler moves on the next host in the list, repeating the process until it is able to successfully claim the resources.

11. Once the Scheduler has found a host for which a successful claim has been made, it needs to select a number of "alternate" hosts. These are hosts from the ranked list that are in the same cell as the selected host, which can be used by the cell conductor in the event that the build on the selected host fails for some reason. The number of alternatives is determined by the configuration option scheduler.max_attempts.

...

**Configuring Filters**

Filters can be set by the users by setting the parameter filter_scheduler.available_filters assigned with the specific filter that must be used. The setting is done in the parameter file. All the filters that are configured can be used by Openstack during run time. With these settings, nova will use the Filter Scheduler that has been configured.

**Designing Weights**

Filter Scheduler uses the so-called weights during its work. A weigher is a way to select the best suitable host from a group of valid hosts by giving weights to all the hosts in the list. In order to prioritize one weigher against another, all the weighers have to define a multiplier that will be applied before computing the weight for a node.

All the weights are normalized beforehand so that the multiplier can be applied easily. Therefore the final weight for the object will be:

weight = w1_multiplier * norm(w1) + w2_multiplier * norm(w2)

A weigher should be a subclass of "weights. Base Host Wegher" and the can implemented both the "weight multiplier" and "weigh_object" methods or just implement the "weight "

**Weigher supported within NOVA**

The filter scheduler weighs hosts based on the config option "filter_scheduler.weight_classes" which defaults to "nova.scheduler.weights.all_weighers". Nova supports many other weighs that include RAMWeigher, CPUWeigher, DiskWeigher etc.

Filter Scheduler makes a local list of acceptable hosts by repeated filtering and weighing. Each time it chooses a host, it virtually consumes resources on it, so subsequent selections can adjust accordingly.

**Limitations of the current Nova Scheduler**

Many users want to do more advanced things with the scheduler, but the current architecture it not ready to support those use cases in a maintainable way. The following issues cannot be addressed with the current form of NOVA scheduler.

## K8s scheduler

- https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/
- https://romanglushach.medium.com/kubernetes-scheduling-understanding-the-math-behind-the-magic-2305b57d45b1

kube-scheduler is the default scheduler for K8s and runs as part of the control plane. kube-scheduler is designed so that, if you want and need to, you can write your own scheduling component and use that instead.

kube-scheduler selects an optimal mode to run newly created or not yet scheduled (unscheduled) pods. Since containers in pods - and pods themselves - can have different requirements, the scheduler filters out any node that don't meet a Pod's specific scheduling needs. Alternatively, the API lets you specify a node for a Pod when you create it, but this is unusal and is only done in special cases.

In a cluster, Nodes that meet the scheduling requirements for a Pod are called feasible nodes. If none of the nodes are suitable, the pod remains unscheduled until the scheduler is able to place it.  

The scheduler finds feasible Nodes for a Pod and then runs a set of functions to score the feasible Nodes and picks a Node with the highest score among the feasible ones to run the Pod. The scheduler then notifies the API server about this decision in a process called binding.

Factors that need to be taken into account for scheduling decisions include individual and collective resource requirements, hardware / software / policy constraints, affinity and anti-affinity specifications data locality, inter workload interference, and so on. 

**Node selection in kube-scheduler**

kube-scheduler selects a node for the pod in a 2-step operation:

1. Filtering
2. Scoring

The filtering step finds the set of Nodes where it's feasible to schedule the Pod. For example, the PodFitsResources filter checks whether a candidate Node has enough available resources to meet a Pod's specific resource requests. After this step, the node list contains any suitable Nodes; often, there will be more than one. If the list is empty, the Pod isn't (yet) schedulable.

In the scoring step, the scheduler ranks the remaining nodes to choose the most suitable Pod placement. The scheduler assigns a score to each Node that survived filtering, basing this score on the active scoring rules.

Finally, kube-scheduler assigns the Pod to the Node with the highest ranking. If there is more than one node with equal scores, kube-scheduler selects one of these at random.

There are two supported ways to configure the filtering and scoring behavior of the scheduler:

1. Scheduling Policies allow you to configure Predicates for filtering and Priorities for scoring.
2. Scheduling Profiles allow you to configure Plugins that implement different scheduling stages, including: QueueSort, Filter, Score, Bind, Reserve, Permit, and others.

Kubernetes scheduling is responsible for assigning pods (groups of one of more containers) to nodes in a cluster. The scheduler's primary objective is to optimize resource utilization and ensure that the application runs smoothly and efficiently. It takes into account factors such as hardware capabilities, available resources, quality of service (QoS), and affinity settings.

Efficient scheduling is crucial for maximizing resource utilization, improving application performance, and ensuring high availability in a Kubernetes cluster. By intelligently assigning pods (groups of containers) to nodes (individual machines in the cluster), K8s scheduling enables workload distribution, load balancing, and fault tolerance.

Kubernetes scheduling is composed of two main phases: filtering and scoring. In the filtering phase, K8s eliminates nodes that are not suitable for running a pod. For example, if a pod requests 2 GB of memory, K8s will filter out nodes that have less than 2GB of free memory. In the scoring phase, K8s assigns a score to each remaining node based on various criteria, such as resource utilization, pod affinity, and node affinity. The node with the highest score is selected as the best fit fort the pod.

The filtering and scoring phases are implemented by two types of components: predicates and priorities. Predicates are boolean functions that return true of false for each node. They are used to implement the filtering logic. Priorities are numeric functions that return a score between 0 and 10 for each node. They are used to implement the scoring logic.

K8s has a default set of predicates and pritorities that cover the most common use cases:

- **PodFitsResources:** checks if a node has enough CPU and memory to run a pod.
- **PodFitsHosts:** checks if a pod has a specific hostname requirement.
- **PodFitsHostPorts:** checks if a node has free ports for a pod's host port mappings.
- **PodSelectorMatches:** checks if a pod's node selector matches a node's labels.
- **NoDiskConflict:** checks if a pod's volume amounts conflict with any other pod's volume mounts on the same node.

### Understanding the Scheduling Algorithm

#### Node Selection

The first step in the K8s scheduling algorithm is node selection.
When a new pod is created, the scheduler evaluates the available nodes in the cluster and identifies suitable candidates for hosting the pod.

This evaluation is based on several factors, including:

- **resource availability:** ability of a computing system to allocate resources such as CPU, memory, storage, and network bandwidth to meet the demands of its workloads. It involves ensuring that there are sufficient resources available to run applications and services without bottlenecks or interruptions. Resource availability can be managed through techniques like resource scheduling, load balancing, and auto-scaling. 
- **node capacity:** maximum amount of resources that a single node in a distributed system can handle. It includes factors such as processing.
