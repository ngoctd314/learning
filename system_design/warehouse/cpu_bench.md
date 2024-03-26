Total items: 10M

|Items của kho X|Physical CPU|Thời gian thực hiện|< 1s|
|-|-|-|-|
|1000|1|318.76766 ms|OK|
|1000|2|239.088106 ms|OK|
|1000|6|135.437861 ms|OK|
|5000|1|1.36215794 s|FAIL|
|5000|2|903.73861 ms|OK|
|5000|6|565.390196 ms|OK|
|8000|1|1.882444939 s|FAIL|
|8000|2|1.283444447 s|FAIL|
|8000|6|875.923163 ms|OK|
|10000|1|2.696601348 s|FAIL|
|10000|2|1.811389618 s|FAIL|
|10000|6|1.119023969 s|FAIL|

=> Với server chỉ với 6 CPUs thì solution này không đáp ứng được câu query đếm distinct items trong một kho có > 10000 items mà thời gian < 1s. Nếu item < 10K thì có thể đáp ứng được.

Tuy nhiên nếu có server đủ khỏe (tầm 24 physical CPUs) thì giải quyết các kho 10000 dự tính chỉ trong khoảng 0.3s.

**P/S** Benchmark trên physical CPUs, không phải logical.
