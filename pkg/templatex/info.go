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
	TypeName string
	First    bool
	End      uint8
}

func (t *Template) FuncInfos() []Info {
	infos := make([]Info, 0, len(t.funcs))

	infoMap := make(map[string]struct{})
	for name, fn := range t.funcs {
		// get function signature with reflect as string
		reflectFn := reflect.ValueOf(fn)
		reflectFnType := reflectFn.Type()

		description := extractDescription(name, reflectFnType, infoMap)

		infos = append(infos, Info{
			Name:        name,
			Description: description,
		})
	}

	return infos
}

func extractDescription(name string, reflectFnTypeOrg reflect.Type, infoMap map[string]struct{}) string {
	description := strings.Builder{}

	subFuncs := []SubFuncs{
		{Name: name, Type: reflectFnTypeOrg, Inner: 0},
	}

	if infoMap == nil {
		infoMap = make(map[string]struct{})
	}

	for len(subFuncs) > 0 {
		sub := subFuncs[0]
		subFuncs = subFuncs[1:]

		reflectFnType := sub.Type
		reflectFnParams := reflectFnType.NumIn()

		lineDescription := strings.Builder{}

		innerLineStr := ""
		extensionLineStr := ""
		addFirstLine := false
		addLastLine := false
		addEndLine := false

		skipArg := 0
		if sub.Inner > 0 {
			// if subfunction, skip first parameter
			skipArg = 1

			switch {
			case sub.Inner > 1 && sub.First && len(subFuncs) > 0 && sub.Inner == subFuncs[0].Inner:
				innerLineStr = "\n" + strings.Repeat(" ", int(sub.End)) + strings.Repeat("│", max(int(sub.Inner)-1-int(sub.End), 0)) + "├─ "
			case sub.Inner > 1 && len(subFuncs) > 0 && sub.Inner != subFuncs[0].Inner:
				addLastLine = true
				extensionLineStr = "└─ "
				innerLineStr = "\n" + strings.Repeat("│", max(int(sub.Inner)-1, 0))
			case len(subFuncs) > 0:
				addFirstLine = true
				extensionLineStr = "├─ "
				innerLineStr = "\n" + strings.Repeat(" ", int(sub.End)) + strings.Repeat("│", max(int(sub.Inner)-1-int(sub.End), 0))
			default:
				addEndLine = true
				extensionLineStr = "└─ "
				innerLineStr = "\n" + strings.Repeat(" ", int(sub.End)) + strings.Repeat("│", max(int(sub.Inner)-1-int(sub.End), 0))
			}
		}

		reflectFnResults := reflectFnType.NumOut()
		lineDescription.WriteString(sub.Name + "(")
		if reflectFnParams > 0 {
			for i := skipArg; i < reflectFnParams; i++ {
				if reflectFnType.IsVariadic() && i == reflectFnParams-1 {
					lineDescription.WriteString("..." + reflectFnType.In(i).Elem().String())
					continue
				}

				lineDescription.WriteString(reflectFnType.In(i).String())
				if i < reflectFnParams-1 {
					lineDescription.WriteString(", ")
				}
			}
		}

		lineDescription.WriteString(")")

		first := true

		if reflectFnResults > 0 {
			if reflectFnResults > 1 {
				lineDescription.WriteString(" (")
			} else {
				lineDescription.WriteString(" ")
			}

			for i := range reflectFnResults {
				out := reflectFnType.Out(i)
				outKind := out.Kind()
				outStr := out.String()
				outStrTrim := strings.TrimPrefix(outStr, "*")

				lineDescription.WriteString(outStr)

				check := sub.TypeName
				if check == "" {
					check = sub.Name
				}

				if strings.Contains(outStr, check) {
					if sub.TypeName == "" {
						if _, ok := infoMap[outStrTrim]; ok {
							continue
						}

						infoMap[outStrTrim] = struct{}{}
					}

					var subM []SubFuncs
					if outKind == reflect.Struct || outKind == reflect.Pointer {
						for iMethod := range out.NumMethod() {
							outMethod := out.Method(iMethod)
							if !outMethod.IsExported() {
								continue
							}

							// prevent duplicate
							if _, ok := infoMap[outStrTrim+"."+outMethod.Name]; ok {
								continue
							}

							infoMap[outStrTrim+"."+outMethod.Name] = struct{}{}

							end := sub.End
							if addEndLine {
								end = sub.End + 1
							}

							subM = append(subM, SubFuncs{
								Name:     outMethod.Name,
								Type:     outMethod.Type,
								Inner:    sub.Inner + 1,
								TypeName: strCut(outStr),
								First:    first,
								End:      end,
							})

							first = false
						}
					}

					subFuncs = append(subM, subFuncs...)
				}

				if i < reflectFnResults-1 {
					lineDescription.WriteString(", ")
				}
			}

			if reflectFnResults > 1 {
				lineDescription.WriteString(")")
			}
		}

		if addFirstLine && !first {
			extensionLineStr = "├┐ "
		}

		if addLastLine && !first {
			extensionLineStr = "└┐ "
		}

		if addEndLine && !first {
			extensionLineStr = "└┐ "
		}

		description.WriteString(innerLineStr + extensionLineStr + lineDescription.String())
	}

	return description.String()
}

func strCut(s string) string {
	i := strings.Index(s, ".")

	if i == -1 {
		return s
	}

	return s[:i]
}
