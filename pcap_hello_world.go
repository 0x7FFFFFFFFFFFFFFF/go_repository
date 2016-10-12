package main

import (
    "fmt"
    "os"
    "time"

    "github.com/davecgh/go-spew/spew"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

func main() {
    version := pcap.Version()
    fmt.Println(version)

    var devices = []pcap.Interface{}
    devices, _ = pcap.FindAllDevs()

    spew.Fdump(os.Stdout, devices)

    // Open a live device
    handle, _ := pcap.OpenLive("en0", int32(65535), false, -1*time.Second)
    defer handle.Close()
    // We can also open an offline file (dumps of WireShark, tcpdump etc are supported)
    // handle, _ := pcap.OpenOffline("dump.pcap")
    // defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    fmt.Println("======================================================")

    // Read one packet
    // packet, _ := packetSource.NextPacket()
    // spew.Fdump(os.Stdout, packet)

    // fmt.Println("======================================================")
    // fmt.Println(packet)

    // Read multiple packets
    for packet := range packetSource.Packets() {
        fmt.Println(packet)
    }
}
