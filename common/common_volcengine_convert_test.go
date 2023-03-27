package common

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func mockTestResource() *schema.Resource {
	return &schema.Resource{Schema: mockTestFieldsSchema()}
}

func mockTestResourceData() *schema.ResourceData {
	states := map[string]string{
		"availability_zone": "foo",
		"ports.#":           "3",
		"ports.0":           "1",
		"ports.1":           "2",
		"ports.2":           "5",
		"ingress.#":         "1",
		"ingress.0.from":    "8080",
	}
	resourceData, err := schema.InternalMap(mockTestFieldsSchema()).Data(&terraform.InstanceState{
		Attributes: states,
	}, nil)
	if err != nil {
		panic(err)
	}
	return resourceData
}

func mockTestFieldsSchema() map[string]*schema.Schema {
	fieldsSchema := map[string]*schema.Schema{
		"availability_zone": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"ports_empty": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"ingress": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"from": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
	}
	return fieldsSchema
}

func Test_ResourceDateToRequest_RequestConvertAll_ContentTypeJson_NoUpdate(t *testing.T) {
	resource := mockTestResource()
	resourceData := mockTestResourceData()

	data := []struct {
		Convert map[string]RequestConvert
		Result  map[string]interface{}
	}{
		{
			Convert: map[string]RequestConvert{},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo",
				"Ingress": []interface{}{map[string]interface{}{
					"from": 8080,
				}},
				"Ports": []interface{}{1, 2, 5},
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertJsonObject,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo",
				"Ingress.From":     8080,
				"Ports":            []interface{}{1, 2, 5},
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertJsonObjectArray,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo",
				"Ingress.1.From":   8080,
				"Ports":            []interface{}{1, 2, 5},
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertJsonObjectArray,
					TargetField: "TestIngress",
					NextLevelConvert: map[string]RequestConvert{
						"from": {
							TargetField: "TestFrom",
						},
					},
				},
				"ports": {
					Ignore: true,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone":       "foo",
				"TestIngress.1.TestFrom": 8080,
			},
		},
	}
	for _, element := range data {
		resp, err := ResourceDateToRequest(resourceData, resource, false, element.Convert, RequestConvertAll, ContentTypeJson)
		assert.Nil(t, err)
		assert.Equal(t, element.Result, resp)
	}
}

func Test_ResourceDateToRequest_RequestConvertAll_ContentTypeDefault_NoUpdate(t *testing.T) {
	resource := mockTestResource()
	resourceData := mockTestResourceData()

	data := []struct {
		Convert map[string]RequestConvert
		Result  map[string]interface{}
	}{
		{
			Convert: map[string]RequestConvert{},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo",
				"Ingress": []interface{}{map[string]interface{}{
					"from": 8080,
				}},
				"Ports": []interface{}{1, 2, 5},
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertListN,
				},
				"ports": {
					ConvertType: ConvertWithN,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo",
				"Ingress.1.From":   8080,
				"Ports.1":          1,
				"Ports.2":          2,
				"Ports.3":          5,
			},
		},
	}
	for _, element := range data {
		resp, err := ResourceDateToRequest(resourceData, resource, false, element.Convert, RequestConvertAll, ContentTypeDefault)
		assert.Nil(t, err)
		assert.Equal(t, element.Result, resp)
	}
}

func Test_ResourceDateToRequest_RequestConvertInConvert_ContentTypeDefault_NoUpdate(t *testing.T) {
	resource := mockTestResource()
	resourceData := mockTestResourceData()

	data := []struct {
		Convert map[string]RequestConvert
		Result  map[string]interface{}
	}{
		{
			Convert: map[string]RequestConvert{},
			Result:  map[string]interface{}{},
		},
		{
			Convert: map[string]RequestConvert{
				"availability_zone": {
					TargetField: "ZoneTest",
				},
			},
			Result: map[string]interface{}{
				"ZoneTest": "foo",
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertListN,
				},
				"ports": {
					ConvertType: ConvertWithN,
				},
			},
			Result: map[string]interface{}{
				"Ingress.1.From": 8080,
				"Ports.1":        1,
				"Ports.2":        2,
				"Ports.3":        5,
			},
		},
	}
	for _, element := range data {
		resp, err := ResourceDateToRequest(resourceData, resource, false, element.Convert, RequestConvertInConvert, ContentTypeDefault)
		assert.Nil(t, err)
		assert.Equal(t, element.Result, resp)
	}
}

func mockTestDiffResourceData() *schema.ResourceData {
	states := map[string]string{
		"availability_zone": "foo",
		"ports.#":           "3",
		"ports.0":           "1",
		"ports.1":           "2",
		"ports.2":           "5",
		"ingress.#":         "1",
		"ingress.0.from":    "8080",
	}
	diff := map[string]*terraform.ResourceAttrDiff{
		"availability_zone": {
			Old: "foo",
			New: "foo_new",
		},
		"ingress.0.from": {
			Old: "8080",
			New: "9999",
		},
	}
	resourceData, err := schema.InternalMap(mockTestFieldsSchema()).Data(&terraform.InstanceState{
		Attributes: states,
	}, &terraform.InstanceDiff{
		Attributes: diff,
	})
	if err != nil {
		panic(err)
	}
	return resourceData
}

func Test_ResourceDateToRequest_RequestConvertAll_ContentTypeDefault_Update(t *testing.T) {
	resource := mockTestResource()
	resourceData := mockTestDiffResourceData()

	data := []struct {
		Convert map[string]RequestConvert
		Result  map[string]interface{}
	}{
		{
			Convert: map[string]RequestConvert{},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo_new",
				"Ingress": []interface{}{map[string]interface{}{
					"from": 9999,
				}},
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertListUnique,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo_new",
				"Ingress.From":     9999,
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertListN,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo_new",
				"Ingress.1.From":   9999,
			},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					Ignore: true,
				},
				"ports": {
					ForceGet:    true,
					ConvertType: ConvertWithN,
				},
			},
			Result: map[string]interface{}{
				"AvailabilityZone": "foo_new",
				"Ports.1":          1,
				"Ports.2":          2,
				"Ports.3":          5,
			},
		},
	}
	for _, element := range data {
		resp, err := ResourceDateToRequest(resourceData, resource, true, element.Convert, RequestConvertAll, ContentTypeDefault)
		assert.Nil(t, err)
		assert.Equal(t, element.Result, resp)
	}
}

func Test_ResourceDateToRequest_RequestConvertIgnore_ContentTypeDefault_Update(t *testing.T) {
	resource := mockTestResource()
	resourceData := mockTestDiffResourceData()

	data := []struct {
		Convert map[string]RequestConvert
		Result  map[string]interface{}
	}{
		{
			Convert: map[string]RequestConvert{},
			Result:  map[string]interface{}{},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					ConvertType: ConvertListUnique,
				},
			},
			Result: map[string]interface{}{},
		},
		{
			Convert: map[string]RequestConvert{
				"ingress": {
					Ignore: true,
				},
				"ports": {
					ForceGet:    true,
					ConvertType: ConvertWithN,
				},
			},
			Result: map[string]interface{}{},
		},
	}
	for _, element := range data {
		resp, err := ResourceDateToRequest(resourceData, resource, true, element.Convert, RequestConvertIgnore, ContentTypeDefault)
		assert.Nil(t, err)
		assert.Equal(t, element.Result, resp)
	}
}

func mockTestEmptyResourceData() *schema.ResourceData {
	states := map[string]string{}
	resourceData, err := schema.InternalMap(mockTestFieldsSchema()).Data(&terraform.InstanceState{
		Attributes: states,
	}, nil)
	if err != nil {
		panic(err)
	}
	return resourceData
}

func Test_ResponseToResourceData(t *testing.T) {
	data := []struct {
		Request  map[string]interface{}
		Convert  map[string]ResponseConvert
		Response map[string]interface{}
	}{
		{
			Request: map[string]interface{}{
				"AvailabilityZone": "cn-shanghai",
			},
			Response: map[string]interface{}{
				"availability_zone": "cn-shanghai",
			},
		},
		{
			Request: map[string]interface{}{
				"Ports": []interface{}{1, 20, 192},
			},
			Response: map[string]interface{}{
				"ports": []interface{}{1, 20, 192},
			},
		},
		{
			Request: map[string]interface{}{
				"Data": map[string]interface{}{
					"Zone": "cn-shanghai",
				},
			},
			Convert: map[string]ResponseConvert{
				"Data.Zone": {
					TargetField: "availability_zone",
				},
			},
			Response: map[string]interface{}{
				"availability_zone": "cn-shanghai",
			},
		},
	}
	for _, element := range data {
		resourceData := mockTestEmptyResourceData()
		_, err := ResponseToResourceData(resourceData, mockTestResource(), element.Request, element.Convert)
		assert.Nil(t, err)

		for k, v := range element.Response {
			assert.Equal(t, v, resourceData.Get(k))
		}
	}
}
