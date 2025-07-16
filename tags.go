package peg

import "reflect"

type configTags []reader

func (c *configTags) bind(cf configField, path ...string) {
	var val reflect.Value

	switch v := cf.val.(type) {
	case reflect.Value:
		val = v
	default:
		val = reflect.ValueOf(v)
	}

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

		nextCf := configField{
			val:          fieldVal,
			name:         field.Tag.Get("_name"),
			usage:        field.Tag.Get("_usage"),
			defaultValue: field.Tag.Get("_default"),
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

func (c *configTags) Read() {
	for _, r := range *c {
		r.read()
	}
}
