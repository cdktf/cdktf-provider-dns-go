//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DnsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDnsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDnsProviderParameters(scope constructs.Construct, id *string, config *DnsProviderConfig) error {
	return nil
}

