package main

import (
	"fmt"
	"regexp"
	"strings"
)

var ClusterName string = "1111111111"
var ClusterAlias string = "ehr_oc_stage_auto"
var filters  = []string{"alias~=.*auto.*"}

func  filtersMatchCluster(filters []string) bool {
	fmt.Println(filters)
	for _, filter := range filters {
		if filter == ClusterName {
			return true
		}
		if filter == ClusterAlias {
			return true
		}
		if strings.HasPrefix(filter, "alias=") {
			// Match by exact cluster alias name
			alias := strings.SplitN(filter, "=", 2)[1]
			if alias == ClusterAlias {
				return true
			}
		} else if strings.HasPrefix(filter, "alias~=") {
			// Match by cluster alias regex

			aliasPattern := strings.SplitN(filter, "~=", 2)[1]
			fmt.Println(aliasPattern)
			if matched, _ := regexp.MatchString(aliasPattern, ClusterAlias); matched {
				return true
			}
		} else if filter == "*" {
			return true
		} else if matched, _ := regexp.MatchString(filter, ClusterName); matched && filter != "" {
			return true
		}
	}
	return false
}

func main()  {
	b := filtersMatchCluster(filters)
	fmt.Println(b)

	//bb, _ := regexp.MatchString(".*(?<!not_auto)$","ehr_oc_stage")
	//fmt.Println(bb)
	//matched, _ := regexp.MatchString(".*(?<!not_auto)$", "ehr_oc_stage")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString("(?s)(?!.*not_auto).*", "ehr_oc_stage")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString("^(?!.*not_auto$).*", "ehr_oc_stage")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString(".*[^n][^o][^t][_][a][u][t][o]$", "ehr_oc_stage")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString("^(?!.*not_auto$).*", "ehr_oc_stage")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString("^(?!.*not_auto$).*", "ehr_oc_stage_not_auto")
	//fmt.Println(matched)
	//
	//matched, _ = regexp.MatchString(".*auto$", "ehr_oc_stage_auto")
	//fmt.Println(matched)
	//matched, _ = regexp.MatchString("^((?!not_auto).)*$", "ehr_oc_stage")
	//fmt.Println(matched)


}
