# cpuinfo

```bash
lscpu
```

```bash
cat /proc/cpuinfo
```

```txt
processor       : 11
vendor_id       : GenuineIntel
cpu family      : 6
model           : 154
model name      : 12th Gen Intel(R) Core(TM) i7-1255U
stepping        : 4
microcode       : 0x430
cpu MHz         : 1447.393
cache size      : 12288 KB
physical id     : 0
siblings        : 12
core id         : 15
cpu cores       : 10
apicid          : 30
initial apicid  : 30
fpu             : yes
fpu_exception   : yes
cpuid level     : 32
wp              : yes
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf tsc_known_freq pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault epb ssbd ibrs ibpb stibp ibrs_enhanced tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid rdseed adx smap clflushopt clwb intel_pt sha_ni xsaveopt xsavec xgetbv1 xsaves split_lock_detect avx_vnni dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp hwp_pkg_req hfi umip pku ospke waitpkg gfni vaes vpclmulqdq rdpid movdiri movdir64b fsrm md_clear serialize arch_lbr ibt flush_l1d arch_capabilities
vmx flags       : vnmi preemption_timer posted_intr invvpid ept_x_only ept_ad ept_1gb flexpriority apicv tsc_offset vtpr mtf vapic ept vpid unrestricted_guest vapic_reg vid ple shadow_vmcs ept_mode_based_exec tsc_scaling usr_wait_pause
bugs            : spectre_v1 spectre_v2 spec_store_bypass swapgs eibrs_pbrsb
bogomips        : 5222.40
clflush size    : 64
cache_alignment : 64
address sizes   : 39 bits physical, 48 bits virtual
power management:
```

The information you provided is the output of CPU-related details from the `lscpu` or a similar command on a Linux system.

1. Processor information

vendor_id, cpu family, model, model name, stepping, microcode

2. Clock Speed

cpu MHz: 1447.393 MHz ~ 1.5*10^9 cycles per second

3. Cache Information

cache size

4. Core Information

siblings, cpu cores

5. Feature Flags

6. Cache Line Size

clflush size: 64 bytes

7. Address Sizes

- physical: 39 bits

This refers to the size of the physical memory address space that the processor can handle. In a system with a 39-bit physical address space, the processor can theoretically address up to 2^39 unique memory locations. This corresponds to 512 gigabytes (GB) physical memory.

Physical Address: Represents the actual location in the physical RAM (Random Access Memory) where data is stored. 

- virtual: 48 bits

This refer to the size of the virtual memory address space that the processor can handle. A 48-bit virtual address allow for 2^48 unique memory locations. This corresponds to 256TB of virtual memory.

Virtual Address: Represents the address that a program or process "thinks" it is using. Virtual memory allows the operating system  to use a larger virtual address space than the physical RAM available, and it involves mapping virtual addresses to physical addresses.
