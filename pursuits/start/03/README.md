# Question

Please `go run ./dont_cheat_sender.go` and `go run ./server.go`. This will spawn
a sender and a receiver. Please find the source IP of the sender without looking
at the code. Store the last byte of the address to the base 16 ("hex").

# Background

The sender and receiver send datagrams using UDP over IPv4. IPv4 provides a packet
structure with a source and destination IP (and other bits). UDP adds source and
destination port, payload (and other bits).

We are using (an extended) BSD socket API to request the operating system to provide
network services for us (send/receive/routing/congestion/...). The system calls in
questions are variants of:

* socket
* read/write
* sendmsg/recvmsg
* select/poll/epoll
* Original BPF for network filtering



* [RFC 791](https://datatracker.ietf.org/doc/html/rfc791)
** Look for what is (was) in an address. It is an identifier. What is a locator?
* [RFC 3493](https://datatracker.ietf.org/doc/html/rfc3493) IPv6 extensions to the BSD socket interface

