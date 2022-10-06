package txtrecordset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TxtRecordSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/txt_record_set#txt TxtRecordSet#txt}.
	Txt *[]*string `field:"required" json:"txt" yaml:"txt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/txt_record_set#zone TxtRecordSet#zone}.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/txt_record_set#id TxtRecordSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/txt_record_set#name TxtRecordSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/dns/r/txt_record_set#ttl TxtRecordSet#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

