package peg

import "reflect"

type configTags []source

func (c *configTags) bind(cf *configField, path ...string) {
	val := cf.val()

	t := val.Type()

	if t.Kind() != reflect.Struct {
		val = val.Elem()
		t = t.Elem()
	}

	if cf.name != "" {
		path = append(path, cf.name)
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		if field.Tag == "" {
			continue
		}

		fieldVal := val.Field(i)

		nextCf := &configField{
			val:          func() reflect.Value { return fieldVal },
			name:         field.Tag.Get("peg.name"),
			usage:        field.Tag.Get("peg.usage"),
			defaultValue: field.Tag.Get("peg.default"),
			required:     field.Tag.Get("peg.required"),
		}

		if field.Type.Kind() == reflect.Struct {
			c.bind(nextCf, path...)
			continue
		}

		for _, r := range *c {
			r.bind(nextCf, path...)
		}
	}
}

func (c *configTags) Read() (err error) {
	for _, r := range *c {
		if err = r.read(); err != nil {
			break
		}
	}

	return
}
