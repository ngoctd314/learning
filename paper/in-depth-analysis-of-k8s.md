# In-depth Analysis Of K8s Scheduling

https://zivukushingai.medium.com/in-depth-analysis-of-kubernetes-scheduling-abf08f949924

The key capbility of K8s is its ability to schedule containers onto cluster nodes automatically.

1. The scheduling process
2. Scheduling algorithms
3. Scheduler components
4. Scheduling in action

## What is K8s Scheduling

K8s scheduling refers to the automated procees of placing containerized workloads, known as Pods, onto nodes within a K8s cluster. The K8s scheduler is the core component that handles this placement process.

When a new Pod is created, it is initially unscheduled. The scheduler will then match the Pod to an appropriate node based on the Pod's resource requirements, hardware constraints, affinity rules, current node utilization, and other factors. The goal is to optimally place Pods for performance and high availability.
