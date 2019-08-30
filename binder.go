package resozyme

// BindTo binds objects to the resource.
func BindTo(resc Resource, args ...interface{}) Resource {
	for _, i := range args {
		resc.Bind(i)
	}
	return resc
}
