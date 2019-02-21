package main

import "github.com/spf13/pflag"

var nodesPerCluster uint32 = 1

var (
	image       string
	start       int
	nodes       int
	clusterBits uint32
	nodeBits    uint32
	server      int
	iface       string
)

func main() {

	pflag.StringVarP(&image, "image", "i", "geth", "image")
	pflag.IntVarP(&start, "start", "p", 0, "starting position")
	pflag.IntVarP(&nodes, "nodes", "n", 10, "number of nodes")
	pflag.Uint32VarP(&clusterBits, "cluster-bits", "b", 14, "cluster bits")
	pflag.Uint32VarP(&nodeBits, "node-bits", "c", 2, "node bits")
	pflag.IntVarP(&server, "server", "s", 1, "server id")
	pflag.StringVarP(&iface, "iface", "I", "eno4", "interface")

	pflag.Parse()

	deleteContainer(start, nodes)
	buildNetwork(nodes, start)
	buildContainer(image, nodes)
	joinNetwork(start, nodes)
	startContainer(start, nodes)
}
