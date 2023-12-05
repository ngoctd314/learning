# Summary

The three most common performance-related requests we receive in our consulting practice are to find out whether a server is doing all of its optimally, to find out why a specific query is not executing quickly enough, and to troubleshoot mysterious intermittent incidents, which users commonly call "stalls", "pileups", or "freezes". This chapter is a direct response to those three types of requests. We'll show you tools and techniques to help you speed up a server's overall workload, speed up a single query, or troubleshoot and solve a problem when it's hard to observe, and you don't know what causes it or even how it manifests.

This might seem like a tall order, but it turns out that a simple method can show you the signal within the noise. That method is to focus on measuring what the server spends its time doing, and the technique that supports this is called profiling. In this chapter, we'll show you how to measure systems and generate profiles, and we'll show you how to profile your whole stack, from the application to the database server to individual queries.

