package entrypoint

import core "github.com/Anderson-Lu/go_time_finder/core"

var finders []core.TimeFinder

func init() {
	RegistFinders(
		core.NewTimeFindeZh01(),
		core.NewTimeFindeZh02(),
		core.NewTimeFindeZh03(),
		core.NewTimeFindeZh04(),
		core.NewTimeFindeUs01(),
		core.NewTimeFindeUs02(),
		core.NewTimeFindeUs03(),
		core.NewTimeFindeUs04(),
		core.NewTimeFindeCo01(),
	)
}

func RegistFinders(fs ...core.TimeFinder) {
	for _, v := range fs {
		finders = append(finders, v)
	}
}

func FindTime(source string) []core.FinderResult {
	ret := []core.FinderResult{}
	for _, v := range finders {
		ret = v.Try(source)
		if len(ret) > 0 {
			break
		}
	}
	return ret
}
