package templatex

import "reflect"

type Info struct {
	// Name of the function
	Name string
	// Description of the function
	Description string
	// Usage of the function
	Usage string
}

func FuncInfos(fns map[string]interface{}) []Info {
	var infos []Info

	for name, fn := range fns {
		// get function signature with reflect as string
		reflectFn := reflect.ValueOf(fn)
		reflectFnType := reflectFn.Type()
		// reflectFnName := reflectFnType.Name()
		reflectFnParams := reflectFnType.NumIn()
		reflectFnResults := reflectFnType.NumOut()

		// get function description
		var description string
		description = name + "("
		if reflectFnParams > 0 {
			for i := 0; i < reflectFnParams; i++ {
				description += reflectFnType.In(i).String()
				if i < reflectFnParams-1 {
					description += ", "
				}
			}
		}

		description += ")"

		if reflectFnResults > 0 {
			if reflectFnResults > 1 {
				description += " ("
			} else {
				description += " "
			}

			for i := 0; i < reflectFnResults; i++ {
				description += reflectFnType.Out(i).String()
				if i < reflectFnResults-1 {
					description += ", "
				}
			}

			if reflectFnResults > 1 {
				description += ")"
			}
		}

		infos = append(infos, Info{
			Name:        name,
			Description: description,
		})
	}

	return infos
}
