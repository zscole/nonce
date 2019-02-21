package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// Build the network

func buildNetwork(nodes, start int) {

	fmt.Println("Cleaning Networks")

	for i := start; i < nodes+start; i++ {
		// remove networks
		CleanNetCMD := "docker network rm " + fmt.Sprintf("vlan%d", i)
		Cleanparts := strings.Fields(CleanNetCMD)
		Cleanhead := Cleanparts[0]
		Cleanparts = Cleanparts[1:len(Cleanparts)]

		out1, err1 := exec.Command(Cleanhead, Cleanparts...).Output()
		if err1 != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err1)
		}
		fmt.Printf("%s", out1)

		fmt.Printf("-------------Building Node%d Network-------------\n", i)

		cmd := "docker network create -d macvlan --subnet " + getSubnetByVlan(server, i) + " --gateway " + getGateway(server, i) + " -o parent=" + iface + "." + fmt.Sprintf("%d", i+101) + " " + fmt.Sprintf("vlan%d", i)

		parts := strings.Fields(cmd)
		head := parts[0]
		parts = parts[1:len(parts)]

		out2, err2 := exec.Command(head, parts...).Output()
		if err2 != nil {
			fmt.Println("error occured")
			fmt.Printf("%s", err2)
		}
		fmt.Printf("%s", out2)

		fmt.Println(getSubnetByVlan(server, i))
		fmt.Println(getGateway(server, i))
	}
}
