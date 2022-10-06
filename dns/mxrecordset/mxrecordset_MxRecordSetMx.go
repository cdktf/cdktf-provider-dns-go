package mxrecordset


type MxRecordSetMx struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#exchange MxRecordSet#exchange}.
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#preference MxRecordSet#preference}.
	Preference *float64 `field:"required" json:"preference" yaml:"preference"`
}

