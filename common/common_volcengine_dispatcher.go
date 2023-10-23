package common

import (
	"fmt"
	"time"

	re "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Dispatcher struct {
	//rateInfo *RateInfo
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{}
}

func DefaultDispatcher() *Dispatcher {
	return defaultDispatcher
}

func (d *Dispatcher) initDispatcher(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) {

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
	var callbacks []Callback
	if projectUpdateEnabled, ok := resourceService.(ProjectUpdateEnabled); ok {
		projectUpdateCallback := NewProjectService(resourceService.GetClient()).ModifyProject(projectUpdateEnabled.ProjectTrn(),
			resourceDate, resource, resourceService)
		callbacks = append(callbacks, projectUpdateCallback...)
	}
	callbacks = append(callbacks, resourceService.ModifyResource(resourceDate, resource)...)

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
	var (
		instance map[string]interface{}
		callErr  error
	)

	err = re.Retry(3*time.Minute, func() *re.RetryError {
		instance, callErr = resourceService.ReadResource(resourceDate, resourceDate.Id())
		if callErr != nil {
			if ResourceFlowLimitExceededError(callErr) {
				return re.RetryableError(callErr)
			} else {
				return re.NonRetryableError(fmt.Errorf("error on  reading  resource %q, %w", resourceDate.Id(), callErr))
			}
		} else {
			return nil
		}
	})

	if err != nil {
		return err
	}
	handlers := resourceService.WithResourceResponseHandlers(instance)
	if len(handlers) == 0 {
		resourceSpecial(resource, instance, nil)
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
		resourceSpecial(resource, data, convert)
		_, err = ResponseToResourceData(resourceDate, resource, data, convert)
		if err != nil {
			return err
		}
	}
	return err
}

func (d *Dispatcher) Delete(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	var (
		callbacks       []Callback
		unsubscribeInfo *UnsubscribeInfo
	)

	// 自动退订逻辑
	if unsubscribeEnabled, ok := resourceService.(UnsubscribeEnabled); ok {
		unsubscribeInfo, err = unsubscribeEnabled.UnsubscribeInfo(resourceDate, resource)
		if err != nil {
			return err
		}
	}

	if unsubscribeInfo != nil && unsubscribeInfo.NeedUnsubscribe {
		unsubscribeCallback := NewUnsubscribeService(resourceService.GetClient()).UnsubscribeInstance(unsubscribeInfo)
		callbacks = append(callbacks, unsubscribeCallback...)
	} else {
		callbacks = resourceService.RemoveResource(resourceDate, resource)
	}

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
		condition, err = SortAndStartTransJson(condition)
		if err != nil {
			return err
		}
	}
	switch info.ServiceCategory {
	case ServiceBypass:
		condition, err = convertToBypassParams(info.RequestConverts, condition)
		if err != nil {
			return err
		}
	default:
		break
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

	if info.EachResource != nil {
		collection, err = info.EachResource(collection, resourceDate)
		if err != nil {
			return err
		}
	}
	return ResponseToDataSource(resourceDate, resource, info, collection)
}
