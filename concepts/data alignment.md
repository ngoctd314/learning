# Data alignment

https://www.youtube.com/watch?v=tyw6_B0-QZA

So you want to speed up your software so that is is capable of running even the most sophisticated audio algorithms in real time. How to achieve that? Well this question has multiple answers, but one of them is definitely data alignment, which is the topic of today's video.

Address alignment, datum alignment

Let's imagine that computer's memory consists of drawers. Each drawer is of the same size in bytes which is equal to a power of 2, and we can some how, and we can somehow choose which power it is.

For examle, the whole memory could be made up of 4-bytes blocks. We can sill address every byte, but only the first address in each block is said to be "aligned". That is address aligment.
