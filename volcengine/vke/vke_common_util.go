package vke

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-volcengine/common"
)

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

func TransKubernetesConfig(resourceData *schema.ResourceData) map[string]interface{} {
	if value, ok := resourceData.GetOk("kubernetes_config"); ok {
		kubernetesConfig := make(map[string]interface{})
		if kubernetesArr, ok := value.([]interface{}); ok {
			if len(kubernetesArr) > 0 {
				kubernetesMap, ok := kubernetesArr[0].(map[string]interface{})
				if ok {
					if value, ok = kubernetesMap["labels"]; ok {
						labels := make([]interface{}, 0)
						if valueArr, ok := value.([]interface{}); ok {
							for _, v := range valueArr {
								label := make(map[string]interface{})
								if vMap, ok := v.(map[string]interface{}); ok {
									if l, ok := vMap["key"]; ok {
										label["Key"] = l
									}
									if l, ok := vMap["value"]; ok {
										label["Value"] = l
									}
								}
								if len(label) > 0 {
									labels = append(labels, label)
								}
							}
						}
						kubernetesConfig["Labels"] = labels
					}
					if value, ok = kubernetesMap["taints"]; ok {
						taints := make([]interface{}, 0)
						if valueArr, ok := value.([]interface{}); ok {
							for _, v := range valueArr {
								taint := make(map[string]interface{})
								if vMap, ok := v.(map[string]interface{}); ok {
									if l, ok := vMap["key"]; ok {
										taint["Key"] = l
									}
									if l, ok := vMap["value"]; ok {
										taint["Value"] = l
									}
									if l, ok := vMap["effect"]; ok {
										taint["Effect"] = l
									}
								}
								if len(taint) > 0 {
									taints = append(taints, taint)
								}
							}
						}
						kubernetesConfig["Taints"] = taints
					}
					if value, ok = kubernetesMap["cordon"]; ok {
						kubernetesConfig["Cordon"] = value
					}
				}
			}
		}
		return kubernetesConfig
	}
	return nil
}
