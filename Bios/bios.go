package main

import (
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/klog/v2"

	"github.com/yumaojun03/dmidecode"
	bios "github.com/yumaojun03/dmidecode/parser/bios"
)

type BiosInf struct {
	Vendor                                 string               `json:"vendor"`
	BIOSVersion                            string               `json:"bios_version"`
	ReleaseDate                            string               `json:"release_date"`
	StartingAddressSegment                 uint16               `json:"starting_address_segment"`
	RuntimeSize                            bios.RuntimeSize     `json:"runtime_size"`
	RomSize                                bios.RomSize         `json:"rom_size"`
	Characteristics                        bios.Characteristics `json:"characteristics"`
	CharacteristicsExt1                    bios.Ext1            `json:"characteristics_ext_1"`
	CharacteristicsExt2                    bios.Ext2            `json:"characteristics_ext_2"`
	SystemBIOSMajorRelease                 byte                 `json:"system_bios_major_release"`
	SystemBIOSMinorRelease                 byte                 `json:"system_bios_minor_release"`
	EmbeddedControllerFirmwareMajorRelease byte                 `json:"firmware_major_release"`
	EmbeddedControllerFirmawreMinorRelease byte                 `json:"firmawre_minor_release"`
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	t := CheckBIOS()
	port := &BiosInf{}
	_ = json.Unmarshal([]byte(t), port)
	fmt.Println(port)
}

func CheckBIOS() string {
	dmi, err := dmidecode.New()
	checkError(err)
	infos, err := dmi.BIOS()
	checkError(err)
	fmt.Println(infos[0])

	bytes, err := json.Marshal(infos[0])
	if err != nil {
		klog.Errorf("err: %v", err)
	}
	return string(bytes)
}
