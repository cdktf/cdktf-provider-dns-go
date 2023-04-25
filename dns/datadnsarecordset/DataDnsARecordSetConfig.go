package datadnsarecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDnsARecordSetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Host to look up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/dns/3.3.2/docs/data-sources/a_record_set#host DataDnsARecordSet#host}
	Host *string `field:"required" json:"host" yaml:"host"`
}

