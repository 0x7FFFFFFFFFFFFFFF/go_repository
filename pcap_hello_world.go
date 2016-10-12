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
    handle, _ := pcap.OpenLive("lo0", int32(65535), false, -1*time.Second)
    defer handle.Close()

    // We can also open an offline file (dumps of WireShark, tcpdump etc are supported)
    // handle, _ := pcap.OpenOffline("dump.pcap")
    // defer handle.Close()

    // After we have a handle, we can set filters
    // Other types of filters include:
    // | Filter                                         | Description                                   |
    // |------------------------------------------------+-----------------------------------------------|
    // | 10.1.1.3                                       | To and from this ip address                   |
    // | 128.3/16                                       | Anything that's on this subnet                |
    // | port 53                                        | Traffic to and from this port (53 is for DNS) |
    // | host 8.8.8.8 and udp port 53                   | Self explanatory                              |
    // | net 199.16.156.0/22 and port 80                | Self explanatory                              |
    // | (port 80 or port 443) and not host 192.168.0.1 | Self explanatory                              |
    handle.SetBPFFilter("tcp and port 8880")
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
