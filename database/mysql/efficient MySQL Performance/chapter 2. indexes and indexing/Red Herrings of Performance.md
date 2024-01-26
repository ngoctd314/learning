# Red Herrings of Performance

Red herring is an idiom that refers to a distraction from a goal. When tracking down solutions to improve MySQL performance, two red herrings commonly distract engineers: faster hardware and MySQL tuning.

## Better, Faster Hardware!

When MySQL performance isn't acceptable, do not begin by scaling up (using better, faster hardware). It probably will help if you scale up significantly, but you learn nothing because it only proves what you already know: computers run faster on faster hardware. Better, faster hardware is a red herring of performance because you miss learning the real causes of, and solution to, slow performance.

## MySQL Tuning

In the television series Star Trek, engineers are able to modify the ship to increase power to engineers, weapons, shields, sensors, transporters, tractor beams - everything.

MySQL is more difficult to operate than a starship because no such modification are possible. But that does not stop engineers from trying.

**Tuning**

Tuning is adjusting MySQL system variables for research and development (R&D). It's laboratory work with specific goals and criteria. Benchmarking is common: adjusting system variables to measure the effect on performance. 

**Configuration**

Configuration is setting system variables to values that are appropriate for the hardware and environment.

**Optimizing**

Optimizing is improving MySQL performance by reducing the workload or making it more efficient
