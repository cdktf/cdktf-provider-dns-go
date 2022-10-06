package mxrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MxRecordSetConfig struct {
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
	// mx block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#mx MxRecordSet#mx}
	Mx interface{} `field:"required" json:"mx" yaml:"mx"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#zone MxRecordSet#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#id MxRecordSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#name MxRecordSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/mx_record_set#ttl MxRecordSet#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

