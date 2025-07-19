# this is a reference
# comes from this article: https://hackmd.io/@ranmJBMnTBajwNlW9v73TA/HyL2Y5Mwh#:~:text=Namespaces%20provide%20a%20way%20to,environment%20and%20from%20other%20containers.

# creating a network namespace
ip netns add client
ip netns add router
ip netns add server

ip netns

ip link

ip netns exec client ip link

# connecting network namespaces
ip link add dev veth-cl type veth peer name veth-rt-cl # create veth-cl with peer veth-rt-cl
ip link set dev veth-cl netns client # move veth-cl into netns 'client'
ip link set dev veth-rt-cl netns router # move the veth-rt-cl device into 'router' netns

# set up both devices
ip netns exec client ip link set dev veth-cl up
ip netns exec router ip link set dev veth-rt-cl up

# add ip addresses to both devices
ip netns exec client ip addr add 192.168.11.1 dev veth-cl
ip netns exec router ip addr add 192.168.11.2 dev veth-rt-cl


ip netns exec client ip route add 192.168.11.0/24 dev veth-cl
ip netns exec router ip route add 192.168.11.0/24 dev veth-rt-cl

ip netns exec client ping 192.168.11.2

# connecting router and server


# enable IP forwarding on the router

# default routes

#finally

# useful command for debugging