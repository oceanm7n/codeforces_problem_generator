package main

import "os"

type ArgsMap map[string]*string

type Arg struct {
	exists bool
	value  string
}

func (this ArgsMap) GetArg(key string) (arg Arg) {
	v, b := this[key]
	if b {
		return Arg{b, *v}
	} else {
		return Arg{b, ""}
	}
}

func ReadArgs() ArgsMap {
	tmp := make(ArgsMap)
	var lastArg *string
	for _, v := range os.Args[1:] {
		if v[0] == '-' {
			t := ""
			tmp[v] = &t
			lastArg = tmp[v]
		} else {
			if lastArg == nil {
				t := ""
				tmp[v] = &t
			} else {
				if *lastArg == "" {
					*lastArg = v
				} else {
					*lastArg = *lastArg + " " + v
				}
			}
		}
	}
	return tmp
}
