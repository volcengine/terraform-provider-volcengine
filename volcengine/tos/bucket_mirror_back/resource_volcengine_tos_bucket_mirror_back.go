package tos_bucket_mirror_back

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketMirrorBack can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_mirror_back.default bucket_name
```

*/

func ResourceVolcengineTosBucketMirrorBack() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketMirrorBackCreate,
		Read:   resourceVolcengineTosBucketMirrorBackRead,
		Update: resourceVolcengineTosBucketMirrorBackUpdate,
		Delete: resourceVolcengineTosBucketMirrorBackDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the TOS bucket.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The mirror_back rules of the bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The ID of the mirror_back rule.",
						},
						"redirect": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The redirect configuration of the mirror_back rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The type of redirect.",
									},
									"fetch_source_on_redirect": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to fetch source on redirect.",
									},
									"pass_query": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to pass query parameters.",
									},
									"follow_redirect": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to follow redirects.",
									},
									"mirror_header": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The mirror header configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pass_all": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to pass all headers.",
												},
												"pass": {
													Type:        schema.TypeSet,
													Optional:    true,
													Set:         schema.HashString,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The headers to pass.",
												},
												"remove": {
													Type:        schema.TypeSet,
													Optional:    true,
													Set:         schema.HashString,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "The headers to remove.",
												},
												"set": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "The mirror header configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The key of the header.",
															},
															"value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value of the header.",
															},
														},
													},
												},
											},
										},
									},
									"public_source": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The public source configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"source_endpoint": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "The source endpoint.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"primary": {
																Type:        schema.TypeSet,
																Optional:    true,
																Set:         schema.HashString,
																Elem:        &schema.Schema{Type: schema.TypeString},
																Description: "The primary endpoints.",
															},
															"follower": {
																Type:        schema.TypeSet,
																Optional:    true,
																Set:         schema.HashString,
																Elem:        &schema.Schema{Type: schema.TypeString},
																Description: "The follower endpoints.",
															},
														},
													},
												},
												"fixed_endpoint": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether the endpoint is fixed.",
												},
											},
										},
									},
									"transform": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "The transform configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"with_key_prefix": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The key prefix to add.",
												},
												"with_key_suffix": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The key suffix to add.",
												},
												"replace_key_prefix": {
													Type:        schema.TypeList,
													Optional:    true,
													MaxItems:    1,
													Description: "The replace key prefix configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key_prefix": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The key prefix to replace.",
															},
															"replace_with": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value to replace with.",
															},
														},
													},
												},
											},
										},
									},
									"fetch_header_to_meta_data_rules": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The fetch header to metadata rules.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"source_header": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The source header.",
												},
												"meta_data_suffix": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The metadata suffix.",
												},
											},
										},
									},
									"private_source": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The private source configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"source_endpoint": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "The source endpoint.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"primary": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The primary endpoints.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"endpoint": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The endpoint.",
																		},
																		"bucket_name": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The bucket name.",
																		},
																		"credential_provider": {
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "The credential provider.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"role": {
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "The role.",
																					},
																				},
																			},
																		},
																	},
																},
															},
															"follower": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The follower endpoints.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"endpoint": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The endpoint.",
																		},
																		"bucket_name": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The bucket name.",
																		},
																		"credential_provider": {
																			Type:        schema.TypeList,
																			MaxItems:    1,
																			Optional:    true,
																			Description: "The credential provider.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"role": {
																						Type:        schema.TypeString,
																						Optional:    true,
																						Description: "The role.",
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"fetch_source_on_redirect_with_query": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to fetch source on redirect with query.",
									},
								},
							},
						},
						"condition": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The condition of the mirror_back rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_code": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Error code for triggering the source re-fetch function.",
									},
									"key_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The prefix of the object name that matches the source object.",
									},
									"key_suffix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The suffix of the object name that matches the source object.",
									},
									"allow_host": {
										Type:        schema.TypeSet,
										Optional:    true,
										Set:         schema.HashString,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Only when a specific domain name is supported will the origin retrieval be triggered.",
									},
									"http_method": {
										Type:        schema.TypeSet,
										Optional:    true,
										Set:         schema.HashString,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The type of request that triggers the re-sourcing process.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineTosBucketMirrorBackCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketMirrorBackService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketMirrorBack())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_mirror_back %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketMirrorBackRead(d, meta)
}

func resourceVolcengineTosBucketMirrorBackRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketMirrorBackService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketMirrorBack())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_mirror_back %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketMirrorBackUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketMirrorBackService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketMirrorBack())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_mirror_back %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketMirrorBackRead(d, meta)
}

func resourceVolcengineTosBucketMirrorBackDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketMirrorBackService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketMirrorBack())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_mirror_back %q, %s", d.Id(), err)
	}
	return nil
}
