package cmd

// Transform data.
func transform(data *data) *data {
	data.fields = data.fields.
		exclude(args.fieldNamesToExclude).
		buildAccessor()
	return data
}
