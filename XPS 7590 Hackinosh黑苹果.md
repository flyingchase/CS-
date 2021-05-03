# XPS 7590 Hackinosh黑苹果

##  参考

GitHub:

- https://github.com/gorquan/OC-XPS-7590
- https://github.com/Liyunlock/DELL-XPS-15-7590-4K-Touchscreen_OpenCore_Hackintosh
- https://github.com/romancin/Dell-XPS-7590-OpenCore
- https://github.com/kirainmoe/hasee-tongfang-macos/issues/33
- https://github.com/xxxzc/xps15-9570-macos/issues/40
- https://github.com/Drowningfish223/Xps-7590-BigSur
- https://github.com/stakeout55/Dell-XPS-7590-mac-OS-Big-Sur-11.1
- https://github.com/xxxzc/xps15-9570-macos/issues/69


$$
\sum_{i=1}^{k} \operatorname{JSD}\left[\alpha^{(i)}, \hat{\alpha}\right]+\frac{1}{k(k-1)} \sum_{i<j} \operatorname{JSD}\left[\alpha^{(i)}, \alpha^{(j)}\right]
$$


## 准备

### 硬件概况

| name | introduction      | name | introduction  |
| :--: | ----------------- | :--: | :-----------: |
| 型号 | XPS15-7590        | Cpu  |   i7-9750h    |
| 屏幕 | 4k 夏普触控屏     | 网卡 | Dw1820a(更换) |
| 固态 | SN720 512G        | 声卡 | ALC298(原生)  |
| 内存 | 48G(32+16) 英睿达 | Bios |               |

### Bios设置



### CFG解锁

显示器

PciRoot(0x0)/Pci(0x1C,0x0)/Pci(0x0,0x0)



bpr_probedelay=100 bpr_initialdelay=300 bpr_postresetdelay=300



## Bug



### 啰嗦模式



- OC引导Big Sur卡在`IOConsoleUsers: gIOScreenLockState 3, hs 0, bs 0 now` 注入苹果显示器EDID 48HZ (45 46字节替换为A6A6 再使用128计算最后一位字节 )
- 更新WEG DeviceP.. 删除注入EDID 





### 显示器描述文件



使用`ioreg -lw0 | grep IODisplayEDID | sed "/[^<]*</s///" | xxd -p -r | strings -6` 查看屏幕的生产型号



夏普屏幕安装对应的显示器描述文件, 位置`/Library/ColorSync/Profiles`

外接4k显示器 延迟高 卡顿: 重建缓存解决 `sudo kextchache -i\`



### 雷电3设备

- 



### 睡眠唤醒后重启

电池供电下睡眠唤醒后重启——>与TB3有关

bootargs启动参数详解:

 darkwake=1(唤醒 偏好设置中节能选项的小憩功能)

 igfxfw=2 (开启核显的最大频率)

keepsyms=1 (de bug使用)

swd_panic=1

| 参数                | 解释                                                         |
| :------------------ | :----------------------------------------------------------- |
| -v                  | 用于安装前期启动时显示代码界面，在安装macOS时，应添加此参数以获取明确的错误信息。 |
| -x                  | 安全启动模式，类似 Windows 的安全模式。此模式下 macOS 会尽可能少的加载 Kext 文件。 |
| -s                  | 单用户模式。这一模式将会启动终端模式，可以用这种方式修复你的系统。 |
| -f                  | 关闭 Kext 缓存模式，等于强制重建 Kext 缓存。                 |
| -l                  | 在系统日志中输出内存泄漏的相关记录。                         |
| arch=x86_64         | 该参数会强制 macOS 以 64 位内核模式启动，在 10.15 及以上没有什么作用。对应的是 arch=i386，将强制以 32 位模式启动。 |
| iog=0x0             | 此参数将强制 MacBook 机型在合盖后，接入外部显示器和键盘时系统保持开启状态；但同时，此参数会在接入外部显示器时关闭笔记本的内屏，这可能对保护屏幕以及省电有帮助。 |
| platform=X86PC      | 此参数将强制禁用 ACPI 电源管理。而 platform=ACPI 将强制启用 ACPI 电源管理。 |
| idlehalt=1          | 强制 CPU 进入低功耗模式。                                    |
| debug=0x100         | 此参数用于禁用五国图，把 Kernel Panic（内核崩溃）的相关数据直接输出在屏幕上，可用于禁止发生内核崩溃时自动重启，这将对排查错误有助益，这一参数还可以用于 Core Dump。其它可用值还有：0x200，这可以在内核崩溃后使用快捷键（C 继续、R 重启、K 进入 KDB）；0x400 可用于触发内核崩溃后自动进行 Core Dump；0x2000 将只生成并发送 Kernel Panic 日志，不包括完整的 Core Dump。除此之外还有很多其他值，但一般以上几个已经足够安装 macOS 时使用。 |
| keepsyms=1          | 此参数可以为 debug=0xN 提供更多错误信息。                    |
| dart=0              | 此参数会关闭 64 位硬件上的系统 PCI 地址映射器（DART）。DART 在拥有 2GB 以上物理内存的机器上是必需的，在默认情况下 DART 都是加载的。当使用 Clover 引导系统且 BIOS 无法关闭 VT-d 时可尝试此参数。 |
| darkwake=0          | 在拥有完全定制好的 USB 接口时完全不必使用此参数，除非你真的需要操控 HID Tickle 行为。darkwake 是 XNU 的一部分，XNU 是一个混合内核，是 Darwin 系统的一部分（macOS 和 iOS 均使用了 Darwin ）。因启动参数仅用于按位计算，所以可能的值有 0、1、2、3、256、257、258、259 等等以此类推，也因此 darkwake=8 实际等于 darkwake=0；darkwake=10 实际等于 darkwake=2，XNU 自 2782.1.97 起删除了这两个值（8 和 10），故这两个值在 Yosemite 及更高版本 macOS 中已失效。黑苹果建议关闭电能小憩，使用 pmset 命令调试休眠。如果实在需要使用，可尝试 darkwake=0 或 3。更多信息可参考[外网这篇文章](https://heipg.cn/link/aHR0cHM6Ly93d3cuaW5zYW5lbHltYWMuY29tL2ZvcnVtL3RvcGljLzM0MjAwMi1kYXJrd2FrZS1vbi1tYWNvcy1jYXRhbGluYS1ib290LWFyZ3MtZGFya3dha2U4LWRhcmt3YWtlMTAtYXJlLW9ic29sZXRlLw==)。 |
| nvda_drv=1          | 用于启用英伟达显卡驱动，包括开启 NVIDIA Web Driver           |
| nv_disable=1        | 关闭英伟达显卡驱动，请勿与 nvda_drv=1 同时使用。             |
| -no_compat_check    | 用于禁用 macOS 兼容性检查。例如，macOS 11.0 BigSur 不再支持 iMac 2014 年之前推出的机型，此时可使用此参数以禁止兼容性检查，以达到安装目的。 |
| kext-dev-mode=1     | 开启 Kext 开发模式，将允许加载未签名的 Kext。在 Yosemite 及更高版本 macOS 中，默认情况下出于安全原因，只会加载已签名的 Kext。此参数可以在 Yosemite 更改此设置，允许加载未签名的 Kext。在比 Yosemite 更新的 macOS 版本（El Capitan）中，引入了另一种安全机制，即系统完整性保护（SIP，也称为 Rootless），该系统会防止修改系统文件，加载未签名的 Kext 等。SIP 可以通过注入正确的 CSR NVRAM 变量来禁用，也可以通过恢复分区运行命令行禁用它。总之，在 10.11 及以后的系统中，已无需此参数。 |
| cpus=1              | CPU单核模式，用于限制系统中活动 CPU 的数量。苹果的开发者工具有一个选项用于启用或禁用系统中的一些 CPU，但你也可以通过这个参数指定要使用的 CPU 数量。在某些情况下，这也许有助于省电，或者你正在调试 X86 电源驱动。 |
| -xcpm               | 用于强制开启 xcpm 以实现 CPU 原生电源管理，一般用于较老架构的 CPU，例如 Ivy Bridge。 |
| npci=0x2000         | 此参数会禁用某些与 kIOPCIConfiguratorPFM64 相关的 PCI 调试，另一个相似的选择是 npci=0x3000，后者还会禁用与 gIOPCITunnelledKey 相关的调试。当卡在 PCI Start Configuration 时，应使用此参数，因为存在与 PCI 通道有关的 IRQ 冲突。 |
| npci=0x3000         | 同 npci=0x2000。                                             |
| -gux_no_idle        | 用于终止英特尔芯片的空闲模式（idle-mode）功能。              |
| slide=N             | 用于引导系统时分配系统内核在内存中的位置，Clover 在一排加号处卡住可以尝试 slide=0，[其它参考：Slide 值的说明](https://heipg.cn/tutorial/what-is-efi-file.html#关于-slide-值)。 |
| rootless=0          | 使用 Rootless 模式，请勿在 El Capitan 及更高版本的 macOS 上使用，因为从 El Capitan 起引入了 SIP（系统完整性保护）机制。一般情况下关闭 SIP 即可达成你的目的。 |
| -disablegfxfirmware | 在 WhateverGreen.kext 出现之前，该参数用于关闭苹果的 iGPU firmware 以正确驱动 Intel 核显，在 macOS 10.13 及更高版本中已不使用。 |

 

### 睡眠后蓝牙无法使用



## 三码更新

 





