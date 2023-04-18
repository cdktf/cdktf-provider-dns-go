package provider


type DnsProviderUpdate struct {
	// The hostname or IP address of the DNS server to send updates to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#server DnsProvider#server}
	Server *string `field:"required" json:"server" yaml:"server"`
	// gssapi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#gssapi DnsProvider#gssapi}
	Gssapi *DnsProviderUpdateGssapi `field:"optional" json:"gssapi" yaml:"gssapi"`
	// Required if `key_name` is set.
	//
	// When using TSIG authentication, the algorithm to use for HMAC. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha256` or `hmac-sha512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#key_algorithm DnsProvider#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// The name of the TSIG key used to sign the DNS update messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#key_name DnsProvider#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Required if `key_name` is set A Base64-encoded string containing the shared secret to be used for TSIG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#key_secret DnsProvider#key_secret}
	KeySecret *string `field:"optional" json:"keySecret" yaml:"keySecret"`
	// The target UDP port on the server where updates are sent to. Defaults to `53`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#port DnsProvider#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// How many times to retry on connection timeout. Defaults to `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#retries DnsProvider#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Timeout for DNS queries.
	//
	// Valid values are durations expressed as `500ms`, etc. or a plain number which is treated as whole seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#timeout DnsProvider#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Transport to use for DNS queries.
	//
	// Valid values are `udp`, `udp4`, `udp6`, `tcp`, `tcp4`, or `tcp6`. Any UDP transport will retry automatically with the equivalent TCP transport in the event of a truncated response. Defaults to `udp`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#transport DnsProvider#transport}
	Transport *string `field:"optional" json:"transport" yaml:"transport"`
}

