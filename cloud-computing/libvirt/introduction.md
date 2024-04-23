# Introduction

The libvirt project:

- is a toolkit to manage virtualization platforms
- is accessible from C, Python, Perl, Go and more
- is licensed under open source licenses
- supports KVM, Hypervisor.frameword, QEMU, Xen...
- targets Linux, FreeBSD, Windows and macOS
- is used by many applications

Libvirt aims to support building and executing on multiple host OS platforms, as well as working with multiple hypervisors. This document outlines which platforms are targeted for each of these areas.

**Build targets**

- Linux, FreeBSD and macOS
 
 The project aims to support the most recent major version at all times. Support the previous major version will be dropped 2 years after the new major version is released or when the verdor drops support, whichever comes first. In this context, thrid-party efforts to extend the lifetime of a distro are not considered, even when they are endorsed by the vendor.

For the purposes of identifying supported software versions available on Linux, the project will look at CentOS, Debian, Fedora ... and Ubuntu LTS. Other distros will be assumed to ship similar sofware versions.

**Virtualization platforms**

For hypervisor drivers which execute locally (QEMU, LXC, VZ, libxl, etc), the set of supported operating system platforms listed above will inform choices as to the minimum required versions 3rd party libraries and hypervisor management APIs.
