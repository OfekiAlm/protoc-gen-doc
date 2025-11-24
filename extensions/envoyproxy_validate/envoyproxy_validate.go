package extensions

import (
	"encoding/json"
	"reflect"
	"strings"

	envoyvalidate "github.com/envoyproxy/protoc-gen-validate/validate"
	nanopbvalidate "github.com/pseudomuto/protoc-gen-doc/validate"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

// ValidateRule represents a single validator rule from the (validate.rules) method option extension.
type ValidateRule struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// ValidateExtension contains the rules set by the (validate.rules) method option extension.
// ValidateExtension contains the rules set by the (validate.rules) method option extension for
// envoyproxy/protoc-gen-validate.
type ValidateExtension struct {
	*envoyvalidate.FieldRules
	rules []ValidateRule // memoized so that we don't have to use reflection more than we need.
}

// MarshalJSON implements the json.Marshaler interface.
func (v ValidateExtension) MarshalJSON() ([]byte, error) { return json.Marshal(v.Rules()) }

// Rules returns the set of rules for this extension.
func (v ValidateExtension) Rules() []ValidateRule {
	if v.FieldRules == nil {
		return nil
	}
	if v.rules == nil {
		v.rules = flattenRules("", reflect.ValueOf(v.FieldRules))
	}
	return v.rules
}

// NanopbValidateExtension contains the rules for the custom nanopb validate.proto integration.
type NanopbValidateExtension struct {
	*nanopbvalidate.FieldRules
	rules []ValidateRule
}

// MarshalJSON implements json.Marshaler.
func (v NanopbValidateExtension) MarshalJSON() ([]byte, error) { return json.Marshal(v.Rules()) }

func (v NanopbValidateExtension) Rules() []ValidateRule {
	if v.FieldRules == nil {
		return nil
	}
	if v.rules == nil {
		v.rules = flattenRules("", reflect.ValueOf(v.FieldRules))
	}
	return v.rules
}

func flattenRules(prefix string, vv reflect.Value) (rules []ValidateRule) {
	vv = reflect.Indirect(vv)
	vt := vv.Type()
	switch vt.Kind() {
	case reflect.Struct:
	nextField:
		for i := 0; i < vt.NumField(); i++ {
			f := vt.Field(i)
			ft := f.Type
			fv := vv.Field(i)

			for ft.Kind() == reflect.Interface || ft.Kind() == reflect.Ptr {
				if fv.IsNil() {
					continue nextField
				}
				fv = fv.Elem()
				ft = fv.Type()
			}
			name := prefix
			if tag, ok := f.Tag.Lookup("protobuf"); ok {
				for _, opt := range strings.Split(tag, ",") {
					if strings.HasPrefix(opt, "name=") {
						if name != "" && !strings.HasSuffix(name, ".") {
							name += "."
						}
						name += strings.TrimPrefix(opt, "name=")
						break
					}
				}
			} else if _, ok := f.Tag.Lookup("protobuf_oneof"); !ok {
				continue nextField
			}
			rules = append(rules, flattenRules(name, fv)...)
		}
	case reflect.Slice:
		if vv.Len() == 0 {
			return nil
		}
		fallthrough
	default:
		rules = append(rules, ValidateRule{Name: prefix, Value: vv.Interface()})
	}
	return rules
}

func init() {
	// Original envoy extension
	extensions.SetTransformer("validate.rules", func(payload interface{}) interface{} {
		rules, ok := payload.(*envoyvalidate.FieldRules)
		if !ok {
			return nil
		}
		return ValidateExtension{FieldRules: rules}
	})

	// Custom nanopb extension
	extensions.SetTransformer("nanopb_validate.rules", func(payload interface{}) interface{} {
		rules, ok := payload.(*nanopbvalidate.FieldRules)
		if !ok {
			return nil
		}
		return NanopbValidateExtension{FieldRules: rules}
	})
}
