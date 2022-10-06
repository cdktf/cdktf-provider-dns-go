package provider


type DnsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns#alias DnsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

