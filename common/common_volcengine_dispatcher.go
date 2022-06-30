package common

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Dispatcher struct {
}

func (d *Dispatcher) Create(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	callbacks := resourceService.CreateResource(resourceDate, resource)
	var calls []SdkCall
	for _, callback := range callbacks {
		if callback.Err != nil {
			return callback.Err
		}
		err = callback.Call.InitWriteCall(resourceDate, resource, false)
		if err != nil {
			return err
		}
		calls = append(calls, callback.Call)
	}
	return CallProcess(calls, resourceDate, resourceService.GetClient(), resourceService)
}

func (d *Dispatcher) Update(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	callbacks := resourceService.ModifyResource(resourceDate, resource)
	var calls []SdkCall
	for _, callback := range callbacks {
		if callback.Err != nil {
			return callback.Err
		}
		err = callback.Call.InitWriteCall(resourceDate, resource, true)
		if err != nil {
			return err
		}
		calls = append(calls, callback.Call)
	}
	return CallProcess(calls, resourceDate, resourceService.GetClient(), resourceService)
}

func (d *Dispatcher) Read(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	instance, err := resourceService.ReadResource(resourceDate, resourceDate.Id())
	if err != nil {
		return err
	}
	handlers := resourceService.WithResourceResponseHandlers(instance)
	if len(handlers) == 0 {
		_, err = ResponseToResourceData(resourceDate, resource, instance, nil)
		return err
	}
	for _, handler := range handlers {
		var (
			data    map[string]interface{}
			convert map[string]ResponseConvert
		)
		data, convert, err = handler()
		if err != nil {
			return err
		}
		_, err = ResponseToResourceData(resourceDate, resource, data, convert)
		if err != nil {
			return err
		}
	}
	return err
}

func (d *Dispatcher) Delete(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	callbacks := resourceService.RemoveResource(resourceDate, resource)
	var calls []SdkCall
	for _, callback := range callbacks {
		if callback.Err != nil {
			return callback.Err
		}
		err = callback.Call.InitWriteCall(resourceDate, resource, true)
		if err != nil {
			return err
		}
		calls = append(calls, callback.Call)
	}
	return CallProcess(calls, resourceDate, resourceService.GetClient(), resourceService)
}

func (d *Dispatcher) Data(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	var (
		info       DataSourceInfo
		condition  map[string]interface{}
		collection []interface{}
	)
	info = resourceService.DatasourceResources(resourceDate, resource)
	condition, err = DataSourceToRequest(resourceDate, resource, info)
	if err != nil {
		return err
	}
	if info.ContentType == ContentTypeJson {
		condition = sortAndStartTransJson(condition)
	}
	collection, err = resourceService.ReadResources(condition)
	if err != nil {
		return err
	}
	if info.ExtraData != nil {
		collection, err = info.ExtraData(collection)
		if err != nil {
			return err
		}
	}
	return ResponseToDataSource(resourceDate, resource, info, collection)
}
