package vke

import "github.com/volcengine/terraform-provider-volcengine/common"

func BinaryJudgment(phase string, conditions interface{}, enables []string) (r []string, err error) {
	var (
		needBj bool
	)
	for _, enable := range enables {
		if phase == enable {
			needBj = true
		}
	}

	if !needBj {
		r = append(r, phase)
	} else {
		for _, condition := range conditions.([]interface{}) {
			var (
				t interface{}
			)
			t, err = common.ObtainSdkValue("Type", condition)
			if err != nil {
				return r, err
			}
			r = append(r, phase+"+"+t.(string))
		}
	}
	return r, err
}
