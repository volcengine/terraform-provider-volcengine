package common

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
	sourceAclParam, err = SortAndStartTransJson((*call.SdkParam)[BypassParam].(map[string]interface{}))
	if err != nil {
		return false, err
	}
	ownerId, _ := ObtainSdkValue("Owner.ID", (*data)[BypassResponse])

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
	owner, _ := ObtainSdkValue("Owner", (*data)[BypassResponse])
	sourceAclParam["Owner"] = owner
	//merge public_acl
	mergeTosPublicAcl(d.Get("public_acl").(string), &sourceAclParam, ownerId.(string))

	(*call.SdkParam)[BypassParam] = sourceAclParam
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
