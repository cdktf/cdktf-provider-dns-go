package srvrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SrvRecordSetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#name SrvRecordSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// srv block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#srv SrvRecordSet#srv}
	Srv interface{} `field:"required" json:"srv" yaml:"srv"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#zone SrvRecordSet#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#id SrvRecordSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/srv_record_set#ttl SrvRecordSet#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

