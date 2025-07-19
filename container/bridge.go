package container

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"log"
)

// creates a veth pair and connects them to a bridge
func CreateVethPairAndBridge() (*netlink.Veth, *netlink.Bridge, error) {
	// Define veth pair names
	veth1Name := "veth1"
	veth2Name := "veth2"
	removeLinkIfExists(veth1Name)
	removeLinkIfExists(veth2Name)

	// Create a veth pair
	veth1 := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name: veth1Name,
		},
		PeerName: veth2Name,
	}

	if err := netlink.LinkAdd(veth1); err != nil {
		log.Fatalf("Failed to create veth pair: %v", err)
		return nil, nil, err
	}
	fmt.Printf("Created veth pair: %s <-> %s\n", veth1Name, veth2Name)

	// Bring up both interfaces
	if err := netlink.LinkSetUp(veth1); err != nil {
		log.Fatalf("Failed to bring up veth1: %v", err)
		return nil, nil, err
	}

	veth2, err := netlink.LinkByName(veth2Name)
	if err != nil {
		log.Fatalf("Failed to find veth2: %v", err)
		return nil, nil, err
	}
	if err := netlink.LinkSetUp(veth2); err != nil {
		log.Fatalf("Failed to bring up veth2: %v", err)
		return nil, nil, err
	}

	// Create a bridge
	bridgeName := "mpodbr"
	removeLinkIfExists(bridgeName)
	bridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{
			Name: bridgeName,
		},
	}

	if err := netlink.LinkAdd(bridge); err != nil {
		log.Fatalf("Failed to create bridge: %v", err)
		return nil, nil, err
	}
	fmt.Printf("Created bridge: %s\n", bridgeName)

	// Bring up the bridge
	if err := netlink.LinkSetUp(bridge); err != nil {
		log.Fatalf("Failed to bring up bridge: %v", err)
		return nil, nil, err
	}

	// Attach veth1 to the bridge
	if err := netlink.LinkSetMaster(veth1, bridge); err != nil {
		log.Fatalf("Failed to attach veth1 to bridge: %v", err)
		return nil, nil, err
	}
	fmt.Printf("Attached %s to bridge %s\n", veth1Name, bridgeName)
	return veth1, bridge, nil
}

func removeLinkIfExists(vethName string) error {
	link, err := netlink.LinkByName(vethName)
	if err != nil {
		return nil
	}

	linkRemovalErr := netlink.LinkDel(link)
	if linkRemovalErr != nil {
		return linkRemovalErr
	}

	return nil
}
