package tun

const (
	Name = "tun0"
	// MTU is the tunnel interface MTU. Use a standard Ethernet-sized MTU so a
	// single frame cannot monopolize the VSOCK byte stream for long periods.
	MTU       = 1500
	ProxyIP   = "10.0.0.1"
	EnclaveIP = "10.0.0.2"
)
