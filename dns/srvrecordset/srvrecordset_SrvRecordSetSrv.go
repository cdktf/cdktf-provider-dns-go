package srvrecordset


type SrvRecordSetSrv struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#port SrvRecordSet#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#priority SrvRecordSet#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#target SrvRecordSet#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#weight SrvRecordSet#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

