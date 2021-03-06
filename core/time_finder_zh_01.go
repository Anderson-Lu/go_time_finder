package core

import (
	"regexp"
	"strconv"

	util "github.com/Anderson-Lu/go_time_finder/util"
)

type TimeFinderZh01 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderZh01) GetName() string {
	return "[ZH-01]"
}

func (self *TimeFinderZh01) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{4})-([\d]{1,2})-([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{4})/([\d]{1,2})/([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{4})\.([\d]{1,2})\.([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2})时([\d]{1,2})分([\d]{1,2})秒?`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2})時([\d]{1,2})分([\d]{1,2})秒?`),
	}
	self.regexs = regexs
}

func NewTimeFindeZh01() *TimeFinderZh01 {
	self := &TimeFinderZh01{}
	self.init()
	return self
}

func (self *TimeFinderZh01) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		for _, value2 := range ts {
			if len(value2) != 7 {
				continue
			}
			year, _ := strconv.Atoi(value2[1])
			month, _ := strconv.Atoi(value2[2])
			day, _ := strconv.Atoi(value2[3])
			hour, _ := strconv.Atoi(value2[4])
			min, _ := strconv.Atoi(value2[5])
			sec, _ := strconv.Atoi(value2[6])
			timeInt := util.GetTime(year, month, day, hour, min, sec)
			tmp := &FinderResult{
				SourceStr: value2[0],
				ResultStr: value2[0],
				ResultUTC: timeInt,
			}
			ret = append(ret, *tmp)
			break
		}
	}
	return ret
}
