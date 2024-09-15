# Talos for Orange Pi 5

[![Build Talos Linux for Orange Pi 5](https://github.com/schneid-l/talos-orangepi5/actions/workflows/build-talos-opi5.yaml/badge.svg)](https://github.com/schneid-l/talos-orangepi5/actions/workflows/build-talos-opi5.yaml)

This repository provides Talos Linux support for the Orange Pi 5 and Orange Pi 5 Plus.

## Upstream dependencies

This repo uses upstream dependencies that are not in sync with the last Talos version.

I choose to stick to the mainline kernel to simplify the future updates and to have the latest features and fixes.

**⚠️ As this repo use the mainline kernel with patches, the full support of the Orange Pi 5 is not guaranteed (e.g. HDMI is not working).**

- Linux kernel: v6.10 (from [Collabora](https://gitlab.collabora.com/hardware-enablement/rockchip-3588/linux))
- Talos Linux: [v1.8.0-alpha.1](https://github.com/siderolabs/talos/tree/v1.8.0-alpha.1)

The best effort is made to keep the overlay in sync with the upstream dependencies.
This repository will be updated as soon as the new versions are available.

## Install

The images provided in this repository do not includes a bootloader as the Orange Pi 5 is equipped with a SPI flash that can be flashed with the bootloader.

### Flash the bootloader

The images provided in this repository are made to be booted with U-Boot or EDK2 UEFI firmware.

I provide U-Boot builds for Orange Pi 5 (and variants) in the [u-boot-orangepi5 repository](https://github.com/schneid-l/u-boot-orangepi5).
The informations to flash the bootloader are described in the repository README.

You can also flash [EDK2 UEFI firmware for Rockchip RK3588 platforms](https://github.com/edk2-porting/edk2-rk3588) (not tested).

The device tree are included in the image at the paths required by U-Boot and EDK2.

### Install Talos Linux

#### Install on a drive

The Talos image can be flashed on an SD card or a NVMe drive.

You can download the latest image from the [releases page](https://github.com/schneid-l/talos-orangepi5/releases).

The image can be flashed using [Etcher](https://www.balena.io/etcher/) on Windows, macOS, or Linux or using `dd` on Linux:

```bash
# Extract the image
zstd -d talos-orangepi5.raw.zst

# Flash the image
# Replace /dev/sdX with the destination device
# You can find the device with `lsblk` or `fdisk -l`
dd if=talos-orangepi5.raw of=/dev/sdX bs=4M status=progress
```

#### PXE Boot

**This repository does not provide a PXE server**, it is up to you to set up the PXE environment.

The [release page](https://github.com/schneid-l/talos-orangepi5/releases) provides the following files needed for PXE boot:

- `kernel-arm64` (the kernel)
- `initramfs-metal-arm64.xz` (the initramfs)
- `rk3588s-orangepi-5.dtb` and `rk3588-orangepi-5-plus.dtb` (the device tree blobs)

## Machine configuration

Use the `ghcr.io/schneid-l/talos-orangepi5` image instead of the upstream Talos Linux one.

```yaml
machine:
  install:
    disk: /dev/sda # replace with the device you want to install Talos on
    image: ghcr.io/schneid-l/talos-orangepi5:v1.1
    wipe: false
```

To upgrade you machine to the latest version with `talosctl`, you can use the following command:

```bash
talosctl upgrade --nodes <node-ip> \
      --image ghcr.io/schneid-l/talos-orangepi5:<version>
```

## Build

Clone the repository and build Talos Linux for Orange Pi 5:

```bash
git clone https://github.com/schneid-l/talos-orangepi5.git
cd talos-orangepi5
make
```

The image will be available in the `out` directory.

_The detail of all the build steps and parameters can be found in the [Makefile](Makefile)._

## License

This project is not affiliated with Xunlong, Orange Pi, Armbian, Collabora or Sidero Labs.

The code in this repository is licensed under the Mozilla Public License Version 2.0 to respect the Talos project license.

## Special thanks

- [Sidero Labs](https://www.siderolabs.com/) for the Talos project
- [Armbian](https://www.armbian.com/) for the initial work on the Orange Pi 5
- [Collabora](https://www.collabora.com/) for the kernel
- [@nberlee](https://github.com/nberlee) and [@pl4nty](https://github.com/pl4nty) for the initial work on other rk3588 devices for Talos and their help ❤️
