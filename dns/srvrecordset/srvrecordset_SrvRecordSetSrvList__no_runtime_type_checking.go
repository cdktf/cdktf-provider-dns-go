//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package srvrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SrvRecordSetSrvList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SrvRecordSetSrvList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SrvRecordSetSrvList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SrvRecordSetSrvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SrvRecordSetSrvList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SrvRecordSetSrvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSrvRecordSetSrvListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

