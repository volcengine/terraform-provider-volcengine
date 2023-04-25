package common

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	TosPath     = "PATH"
	TosDomain   = "DOMAIN"
	TosHeader   = "HEADER"
	TosParam    = "PARAM"
	TosUrlParam = "URL_PARAM"
	TosResponse = "RESPONSE"
	TosFilePath = "FILE_PATH"
)

func convertToTosParams(convert map[string]RequestConvert, condition map[string]interface{}) (result map[string]interface{}, err error) {
	if len(condition) > 0 {
		result = map[string]interface{}{
			TosDomain:   "",
			TosHeader:   make(map[string]string),
			TosPath:     []string{},
			TosParam:    condition,
			TosUrlParam: make(map[string]string),
		}

		if len(convert) > 0 {

			for k, v := range convert {
				k1 := DownLineToHump(k)
				if v.TargetField != "" {
					k1 = v.TargetField
				}
				if v.SpecialParam != nil {
					switch v.SpecialParam.Type {
					case DomainParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[TosDomain] = v1
							delete(condition, k1)
						}
					case UrlParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[TosUrlParam].(map[string]string)[k1] = v1.(string)
							delete(condition, k1)
						}
					case HeaderParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							if v1.(string) != "" {
								result[TosHeader].(map[string]string)[k1] = v1.(string)
							}
							delete(condition, k1)
						}
					case PathParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							temp := result[TosPath].([]string)
							temp = append(temp, strconv.Itoa(v.SpecialParam.Index)+":"+v1.(string))
							result[TosPath] = temp
							delete(condition, k1)
						}
					case FilePathParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[TosFilePath] = v1
							delete(condition, k1)
						}
					}
				}
			}
			//sort
			if v, ok := result[TosPath]; ok {
				temp := v.([]string)
				sort.Strings(temp)
				var afterSort []string
				for _, v1 := range temp {
					afterSort = append(afterSort, v1[strings.Index(v1, ":")+1:])
				}
				result[TosPath] = afterSort
			}
			//query
			result[TosParam] = condition
		}
	}
	return result, err
}

func mergeTosPublicAcl(acl string, param *map[string]interface{}, ownerId string) {
	if _, ok := (*param)["Grants"]; !ok {
		(*param)["Grants"] = []interface{}{}
	}
	vs := (*param)["Grants"].([]interface{})

	defer func() {
		(*param)["Grants"] = vs
	}()

	switch acl {
	case "private":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Id":   ownerId,
				"Type": "CanonicalUser",
			},
			"Permission": "FULL_CONTROL",
		}
		vs = append(vs, m)
		return
	case "public-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AllUsers",
				"Type":   "Group",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	case "public-read-write":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AllUsers",
				"Type":   "Group",
			},
			"Permission": "WRITE",
		}
		vs = append(vs, m)
		return
	case "authenticated-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Canned": "AuthenticatedUsers",
				"Type":   "Group",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	case "bucket-owner-read":
		m := map[string]interface{}{
			"Grantee": map[string]interface{}{
				"Id":   ownerId,
				"Type": "CanonicalUser",
			},
			"Permission": "READ",
		}
		vs = append(vs, m)
		return
	}
}

func BeforeTosPutAcl(d *schema.ResourceData, call SdkCall, data *map[string]interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	var sourceAclParam map[string]interface{}
	sourceAclParam, err = SortAndStartTransJson((*call.SdkParam)[TosParam].(map[string]interface{}))
	if err != nil {
		return false, err
	}
	ownerId, _ := ObtainSdkValue("Owner.ID", (*data)[TosResponse])

	grants, _ := ObtainSdkValue("Grants", sourceAclParam)
	if grants != nil {
		for _, grant := range grants.([]interface{}) {
			id, _ := ObtainSdkValue("Grantee.ID", grant)
			p, _ := ObtainSdkValue("Permission", grant)
			if id == ownerId && p == "FULL_CONTROL" {
				return false, fmt.Errorf("can not set FULL_CONTROL for owner")
			}
		}
	}

	//merge owner
	owner, _ := ObtainSdkValue("Owner", (*data)[TosResponse])
	sourceAclParam["Owner"] = owner
	//merge public_acl
	mergeTosPublicAcl(d.Get("public_acl").(string), &sourceAclParam, ownerId.(string))

	(*call.SdkParam)[TosParam] = sourceAclParam
	return true, nil
}

func ConvertTosAccountAcl() FieldResponseConvert {
	return func(i interface{}) interface{} {
		var accountAcl []interface{}
		owner, _ := ObtainSdkValue("Owner.ID", i)
		grants, _ := ObtainSdkValue("Grants", i)
		if grants != nil {
			for _, grant := range grants.([]interface{}) {
				permission, _ := ObtainSdkValue("Permission", grant)
				id, _ := ObtainSdkValue("Grantee.ID", grant)
				if id == nil {
					continue
				}
				if id == owner && permission == "FULL_CONTROL" {
					continue
				}
				g := map[string]interface{}{
					"AccountId":  id,
					"AclType":    "CanonicalUser",
					"Permission": permission,
				}
				accountAcl = append(accountAcl, g)
			}
		}
		return accountAcl
	}
}

func ConvertTosPublicAcl() FieldResponseConvert {
	return func(i interface{}) interface{} {
		owner, _ := ObtainSdkValue("Owner.ID", i)
		grants, _ := ObtainSdkValue("Grants", i)
		var (
			read  bool
			write bool
		)
		if grants != nil {
			for _, grant := range grants.([]interface{}) {
				id, _ := ObtainSdkValue("Grantee.ID", grant)
				canned, _ := ObtainSdkValue("Grantee.Canned", grant)
				t, _ := ObtainSdkValue("Grantee.Type", grant)
				permission, _ := ObtainSdkValue("Permission", grant)
				if canned != nil && canned.(string) == "AllUsers" && t.(string) == "Group" {
					if permission.(string) == "READ" {
						read = true
						continue
					} else if permission.(string) == "WRITE" {
						write = true
						continue
					}
				}

				if canned != nil && canned.(string) == "AuthenticatedUsers" && t.(string) == "Group" {
					if permission.(string) == "READ" {
						return "authenticated-read"
					}
					break
				}

				if id != nil && id.(string) == owner.(string) && t.(string) == "CanonicalUser" {
					if permission.(string) == "FULL_CONTROL" {
						return "private"
					} else if permission.(string) == "READ" {
						return "bucket-owner-read"
					}
					break

				}

			}
		}

		if read && !write {
			return "public-read"
		}
		if read && write {
			return "public-read-write"
		}
		return ""
	}
}

func TosAccountAclHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s-", m["account_id"]))
	buf.WriteString(fmt.Sprintf("%s", m["permission"]))
	return hashcode.String(buf.String())
}
