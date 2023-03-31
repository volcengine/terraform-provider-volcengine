package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
)

type Dispatcher struct {
	rateInfo *RateInfo
}

var defaultDispatcher *Dispatcher

func init() {
	defaultDispatcher = &Dispatcher{}
}

func DefaultDispatcher() *Dispatcher {
	return defaultDispatcher
}

type RateInfo struct {
	Create *Rate
	Read   *Rate
	Update *Rate
	Delete *Rate
	Data   *Rate
}
type Rate struct {
	Limiter   *rate.Limiter
	Semaphore *semaphore.Weighted
}

func NewRateLimitDispatcher(r *RateInfo) *Dispatcher {
	return &Dispatcher{
		rateInfo: r,
	}
}

func (d *Dispatcher) Create(resourceService ResourceService, resourceDate *schema.ResourceData, resource *schema.Resource) (err error) {
	defer func() {
		if d.rateInfo != nil && d.rateInfo.Create != nil && d.rateInfo.Create.Semaphore != nil {
			d.rateInfo.Create.Semaphore.Release(1)
		}
	}()
	if d.rateInfo != nil && d.rateInfo.Create != nil {
		ctx := context.Background()
		if d.rateInfo.Create.Limiter != nil {
			err = d.rateInfo.Create.Limiter.Wait(ctx)
			if err != nil {
				return err
			}
		}
		if d.rateInfo.Create.Semaphore != nil {
			err = d.rateInfo.Create.Semaphore.Acquire(ctx, 1)
			if err != nil {
				return err
			}
		}

	}
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
	defer func() {
		if d.rateInfo != nil && d.rateInfo.Update != nil && d.rateInfo.Update.Semaphore != nil {
			d.rateInfo.Update.Semaphore.Release(1)
		}
	}()
	if d.rateInfo != nil && d.rateInfo.Update != nil {
		ctx := context.Background()
		if d.rateInfo.Update.Limiter != nil {
			err = d.rateInfo.Update.Limiter.Wait(ctx)
			if err != nil {
				return err
			}
		}
		if d.rateInfo.Update.Semaphore != nil {
			err = d.rateInfo.Update.Semaphore.Acquire(ctx, 1)
			if err != nil {
				return err
			}
		}
	}
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
	defer func() {
		if d.rateInfo != nil && d.rateInfo.Read != nil && d.rateInfo.Read.Semaphore != nil {
			d.rateInfo.Read.Semaphore.Release(1)
		}
	}()
	if d.rateInfo != nil && d.rateInfo.Read != nil {
		ctx := context.Background()
		if d.rateInfo.Read.Limiter != nil {
			err = d.rateInfo.Read.Limiter.Wait(ctx)
			if err != nil {
				return err
			}
		}
		if d.rateInfo.Read.Semaphore != nil {
			err = d.rateInfo.Read.Semaphore.Acquire(ctx, 1)
			if err != nil {
				return err
			}
		}
	}
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
	defer func() {
		if d.rateInfo != nil && d.rateInfo.Delete != nil && d.rateInfo.Delete.Semaphore != nil {
			d.rateInfo.Delete.Semaphore.Release(1)
		}
	}()
	if d.rateInfo != nil && d.rateInfo.Delete != nil {
		ctx := context.Background()
		if d.rateInfo.Delete.Limiter != nil {
			err = d.rateInfo.Delete.Limiter.Wait(ctx)
			if err != nil {
				return err
			}
		}
		if d.rateInfo.Delete.Semaphore != nil {
			err = d.rateInfo.Delete.Semaphore.Acquire(ctx, 1)
			if err != nil {
				return err
			}
		}
	}
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
	defer func() {
		if d.rateInfo != nil && d.rateInfo.Data != nil && d.rateInfo.Data.Semaphore != nil {
			d.rateInfo.Data.Semaphore.Release(1)
		}
	}()
	if d.rateInfo != nil && d.rateInfo.Data != nil {
		ctx := context.Background()
		if d.rateInfo.Data.Limiter != nil {
			err = d.rateInfo.Data.Limiter.Wait(ctx)
			if err != nil {
				return err
			}
		}
		if d.rateInfo.Data.Semaphore != nil {
			err = d.rateInfo.Data.Semaphore.Acquire(ctx, 1)
			if err != nil {
				return err
			}
		}
	}
	info = resourceService.DatasourceResources(resourceDate, resource)
	condition, err = DataSourceToRequest(resourceDate, resource, info)
	if err != nil {
		return err
	}
	if info.ContentType == ContentTypeJson {
		condition = SortAndStartTransJson(condition)
	}
	switch info.ServiceCategory {
	case ServiceTos:
		condition, err = convertToTosParams(info.RequestConverts, condition)
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
	return ResponseToDataSource(resourceDate, resource, info, collection)
}
