package mxrecordset


type MxRecordSetMx struct {
	// The FQDN of the mail exchange, include the trailing dot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.2/docs/resources/mx_record_set#exchange MxRecordSet#exchange}
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// The preference for the record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.2/docs/resources/mx_record_set#preference MxRecordSet#preference}
	Preference *float64 `field:"required" json:"preference" yaml:"preference"`
}

