ip link add veth-host type veth peer name veth-cl

ip link set veth-cl netns client

ip addr add 192.168.11.1/24 dev veth-host
ip link set veth-host up

ip netns exec client ip addr add 192.168.11.2/24 dev veth-cl
ip netns exec client ip link set veth-cl up

ip netns exec client ip link set lo up

sysctl -w net.ipv4.ip_forward=1

iptables -t nat -A PREROUTING -p tcp --dport 8080 -j DNAT --to-destination 192.168.11.2:8080

iptables -t nat -A POSTROUTING -p tcp -d 192.168.11.2 --dport 8080 -j MASQUERADE
