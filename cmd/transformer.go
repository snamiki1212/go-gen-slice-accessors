package cmd

import "slices"

// Transform data.
func transform(data *data, p pluralizer) *data {
	data.fields = data.fields.
		exclude(args.fieldNamesToExclude).
		buildAccessor(p)
	return data
}

// Build accessor name.
func (f *field) buildAccessor(p pluralizer) *field {
	f.Accessor = p.pluralize(f.Name)
	return f
}

// Build accessor names.
func (fs fields) buildAccessor(p pluralizer) fields {
	for i := range fs {
		fs[i].buildAccessor(p)
	}
	return fs
}

// Exclude fields by name.
func (fs fields) exclude(targets []string) fields {
	return slices.DeleteFunc(fs, func(f field) bool {
		return slices.Contains(targets, f.Name)
	})
}
