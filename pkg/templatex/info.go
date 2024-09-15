package templatex

import (
	"reflect"
	"strings"
)

type Info struct {
	// Name of the function
	Name string
	// Description of the function
	Description string
	// Usage of the function
	Usage string
}

type SubFuncs struct {
	Name     string
	Type     reflect.Type
	Inner    uint8
	First    bool
	TypeName string
}

func (t *Template) FuncInfos() []Info {
	var infos []Info

	for name, fn := range t.funcs {
		// get function signature with reflect as string
		reflectFn := reflect.ValueOf(fn)
		reflectFnType := reflectFn.Type()

		description := extractDescription(name, reflectFnType)

		infos = append(infos, Info{
			Name:        name,
			Description: description,
		})
	}

	return infos
}

func extractDescription(name string, reflectFnTypeOrg reflect.Type) string {
	description := strings.Builder{}

	subFuncs := []SubFuncs{
		{Name: name, Type: reflectFnTypeOrg, Inner: 0},
	}

	infoMap := make(map[string]struct{})

	for len(subFuncs) > 0 {
		sub := subFuncs[0]
		subFuncs = subFuncs[1:]

		reflectFnType := sub.Type
		reflectFnParams := reflectFnType.NumIn()

		skipArg := 0
		if sub.Inner > 0 {
			// if subfunction, skip first parameter
			skipArg = 1

			switch {
			// case len(subFuncs) > 0 && sub.Inner > 1 && sub.First:
			// 	description.WriteString("\n" + strings.Repeat(" ", int(sub.Inner-1)) + "└─ ")
			case len(subFuncs) > 0:
				description.WriteString("\n" + strings.Repeat(" ", int(sub.Inner)) + "├─ ")
			default:
				description.WriteString("\n └─ ")
			}
		}

		reflectFnResults := reflectFnType.NumOut()

		// get function description
		description.WriteString(sub.Name + "(")
		if reflectFnParams > 0 {
			for i := skipArg; i < reflectFnParams; i++ {
				if reflectFnType.IsVariadic() && i == reflectFnParams-1 {
					description.WriteString("..." + reflectFnType.In(i).Elem().String())
					continue
				}

				description.WriteString(reflectFnType.In(i).String())
				if i < reflectFnParams-1 {
					description.WriteString(", ")
				}
			}
		}

		description.WriteString(")")

		if reflectFnResults > 0 {
			if reflectFnResults > 1 {
				description.WriteString(" (")
			} else {
				description.WriteString(" ")
			}

			for i := 0; i < reflectFnResults; i++ {
				out := reflectFnType.Out(i)
				outKind := out.Kind()

				description.WriteString(out.String())

				if strings.Contains(out.String(), sub.TypeName) {
					first := true
					if outKind == reflect.Struct || outKind == reflect.Pointer {
						if out.NumMethod() > 0 {
							for i := 0; i < out.NumMethod(); i++ {
								if !out.Method(i).IsExported() {
									continue
								}

								// prevent duplicate
								if _, ok := infoMap[out.Method(i).Name]; ok {
									continue
								}

								subFuncs = append([]SubFuncs{{
									Name:     out.Method(i).Name,
									Type:     out.Method(i).Type,
									Inner:    sub.Inner + 1,
									First:    first,
									TypeName: stringCut(out.String()),
								}}, subFuncs...)

								first = false

								infoMap[out.Method(i).Name] = struct{}{}
							}
						}
					}
				}

				if i < reflectFnResults-1 {
					description.WriteString(", ")
				}
			}

			if reflectFnResults > 1 {
				description.WriteString(")")
			}
		}
	}

	return description.String()
}

func stringCut(s string) string {
	i := strings.Index(s, ".")

	if i == -1 {
		return s
	}

	return s[:i]
}
