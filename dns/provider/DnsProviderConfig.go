package provider


type DnsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#alias DnsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.0/docs#update DnsProvider#update}
	Update *DnsProviderUpdate `field:"optional" json:"update" yaml:"update"`
}

