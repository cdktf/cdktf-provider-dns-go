package provider


type DnsProviderUpdateGssapi struct {
	// The Kerberos realm or Active Directory domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.1/docs#realm DnsProvider#realm}
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// This or `password` is required if `username` is set, not supported on Windows.
	//
	// The path to a keytab file containing a key for `username`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.1/docs#keytab DnsProvider#keytab}
	Keytab *string `field:"optional" json:"keytab" yaml:"keytab"`
	// This or `keytab` is required if `username` is set. The matching password for `username`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.1/docs#password DnsProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user to authenticate as. If not set the current user session will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.1/docs#username DnsProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

