// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

var (
	kernelArgs = []string{
		"earlycon",
		"console=ttyS2,1500000n8",
		"console=tty1",
		"consoleblank=0",
		"sysctl.kernel.kexec_load_disabled=1",
		"talos.dashboard.disabled=1",
		"cgroup_enable=cpuset",
		"cgroup_memory=1",
		"cgroup_enable=memory",
		"swapaccount=1",
		"irqchip.gicv3_pseudo_nmi=0",
		"coherent_pool=2M",
		"pcie_aspm=off",
		"libata.force=noncq",
	}

	edk2DTBDest  = "/boot/EFI/dtb/base"
	uBootDTBDest = "/boot/EFI/dtb/rockchip"

	boardDTBNames = []string{
		"rk3588s-orangepi-5",
		"rk3588-orangepi-5-plus",
	}
)

func main() {
	adapter.Execute(&OrangePi5Installer{})
}

type OrangePi5Installer struct{}

type orangePi5ExtraOptions struct {
	Sata bool `json:"sata"`
}

func (i *OrangePi5Installer) GetOptions(extra orangePi5ExtraOptions) (overlay.Options, error) {
	return overlay.Options{
		Name:       "orangepi5",
		KernelArgs: kernelArgs,
	}, nil
}

func (i *OrangePi5Installer) Install(options overlay.InstallOptions[orangePi5ExtraOptions]) error {
	sataSuffix := ""
	if options.ExtraOptions.Sata {
		sataSuffix = "-sata"
	}

	for board := range boardDTBNames {
		if err := copyFiles(
			filepath.Join(options.ArtifactsPath, "/dtb/", boardDTBNames[board]+sataSuffix+".dtb"),
			filepath.Join(options.MountPrefix, uBootDTBDest, boardDTBNames[board]+".dtb"),
			filepath.Join(options.MountPrefix, edk2DTBDest, boardDTBNames[board]+".dtb")); err != nil {
			return err
		}
	}

	return nil
}

func copyFiles(src string, dst ...string) error {
	for _, d := range dst {
		if err := os.MkdirAll(filepath.Dir(d), 0o666); err != nil {
			return err
		}
		if err := copy.File(src, d); err != nil {
			return err
		}
	}
	return nil
}
