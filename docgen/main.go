package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/volcengine/terraform-provider-vestack/common"
	ve "github.com/volcengine/terraform-provider-vestack/vestack"
)

const (
	docTPL = `---
subcategory: "{{.product}}"
layout: "{{.cloud_mark}}"
page_title: "{{.cloud_title}}: {{.name}}"
sidebar_current: "docs-{{.cloud_mark}}-{{.docType}}-{{.resource}}"
description: |-
  {{.description_short}}
---
# {{.name}}
{{.description}}
## Example Usage
{{.example}}
## Argument Reference
The following arguments are supported:
{{.arguments}}
{{if ne .attributes ""}}
## Attributes Reference
In addition to all arguments above, the following attributes are exported:
{{.attributes}}
{{end}}
{{if ne .import ""}}
## Import
{{.import}}
{{end}}
`
	idxTPL = `
<% wrap_layout :inner do %>
    <% content_for :sidebar do %>
        <div class="docs-sidebar hidden-print affix-top" role="complementary">
            <ul class="nav docs-sidenav">
                <li>
                    <a href="/docs/providers/index.html">All Providers</a>
                </li>
                <li>
                    <a href="/docs/providers/{{.cloud_mark}}/index.html">{{.cloud_title}} Provider</a>
                </li>
                {{range .Products}}
                <li>
                    <a href="#">{{.Name}}</a>
                    <ul class="nav">
                        {{ if .DataSources }}<li>
                            <a href="#">Data Sources</a>
                            <ul class="nav nav-auto-expand">{{range $Resource := .DataSources}}
                                <li>
                                    <a href="/docs/providers/{{$.cloud_mark}}/d/{{replace $Resource $.cloudPrefix ""}}.html">{{$Resource}}</a>
                                </li>{{end}}
                            </ul>
                        </li>{{ end }}
                        <li>
                            <a href="#">Resources</a>
                            <ul class="nav nav-auto-expand">{{range $Resource := .Resources}}
                                <li>
                                    <a href="/docs/providers/{{$.cloud_mark}}/r/{{replace $Resource $.cloudPrefix ""}}.html">{{$Resource}}</a>
                                </li>{{end}}
                            </ul>
                        </li>
                    </ul>
                </li>{{end}}
            </ul>
        </div>
    <% end %>
    <%= yield %>
<% end %>
`
)

const (
	cloudMark    = "vestack"
	cloudTitle   = "Vestack"
	cloudPrefix  = cloudMark + "_"
	docRoot      = "/website"
	providerRoot = "/vestack"
	exampleRoot  = "/example"
)

var (
	hclMatch  = regexp.MustCompile("(?si)([^`]+)?```(hcl)?(.*?)```")
	bigSymbol = regexp.MustCompile("([\u007F-\uffff])")
)

type Product struct {
	Name              string
	DataSource        string
	Resource          string
	DataSourcePath    string
	ResourcePath      string
	DataSourceExample string
	ResourceExample   string
}

var resourceKeys = map[string]string{
	"vpc": "VPC",
	"ecs": "ECS",
	"clb": "CLB",
	"eip": "EIP",
	"ebs": "EBS",
	"nat": "NAT",
}

type Products struct {
	Name        string
	DataSources []string
	Resources   []string
}

func replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func main() {
	var resource string
	flag.StringVar(&resource, "resource", "", "要生成的资源名称首字母小写")
	flag.Parse()
	provider := ve.Provider().(*schema.Provider)
	pwd, _ := os.Getwd()
	if strings.HasSuffix(pwd, "/docgen") {
		pwd = strings.Replace(pwd, "/docgen", "", -1)
	}
	// document for Index
	products := genIndex(pwd)

	for _, product := range products {
		if resource != "" && product.Resource != common.HumpToDownLine(resource) {
			continue
		}
		// document for DataSources
		if product.DataSource != "" {
			_dataSource := cloudMark + "_" + product.DataSource
			genDoc(pwd, product.Name, "data_source", product.DataSourcePath, _dataSource, provider.DataSourcesMap[_dataSource], product.DataSourceExample)
		}

		// document for Resources
		_resource := cloudMark + "_" + product.Resource
		genDoc(pwd, product.Name, "resource", product.ResourcePath, _resource, provider.ResourcesMap[_resource], product.ResourceExample)
	}
}

// genIndex generating index for resource
func genIndex(pwd string) (prods []Product) {
	providerPath := pwd + providerRoot
	provider := ve.Provider().(*schema.Provider)
	fmt.Println(pwd)
	message("generating erb from: %s\n", providerPath)
	files, _ := ioutil.ReadDir(providerPath)
	for _, file := range files {
		if file.IsDir() {
			p1 := providerPath + "/" + file.Name()
			if name, ok := resourceKeys[file.Name()]; ok {
				fs1, _ := ioutil.ReadDir(p1)
				for _, f1 := range fs1 {
					if f1.IsDir() {
						product := Product{Name: name}
						p2 := p1 + "/" + f1.Name()
						fs2, _ := ioutil.ReadDir(p2)
						for _, f2 := range fs2 {
							if strings.HasPrefix(f2.Name(), "data_source_") {
								product.DataSource = strings.Replace(strings.Replace(f2.Name(), "data_source_"+cloudMark+"_", "", -1), ".go", "", -1)
								if provider.DataSourcesMap[cloudMark+"_"+product.DataSource] == nil {
									product.DataSource = ""
								} else {
									product.DataSourcePath = p2 + "/" + f2.Name()
									product.DataSourceExample = pwd + exampleRoot + "/" + "data" + common.DownLineToHump(product.DataSource)
								}

							} else if strings.HasPrefix(f2.Name(), "resource_") {
								product.Resource = strings.Replace(strings.Replace(f2.Name(), "resource_"+cloudMark+"_", "", -1), ".go", "", -1)
								if provider.ResourcesMap[cloudMark+"_"+product.Resource] == nil {
									product.Resource = ""
								} else {
									product.ResourcePath = p2 + "/" + f2.Name()
									product.ResourceExample = pwd + exampleRoot + "/" + common.DownLineToHumpAndFirstLower(product.Resource)
								}
							}
						}
						prods = append(prods, product)
					}
				}

			}
		}
	}
	temp := make(map[string][]Product)
	for _, prod := range prods {
		if ps, ok := temp[prod.Name]; ok {
			temp[prod.Name] = append(ps, prod)
		} else {
			temp[prod.Name] = []Product{prod}
		}
	}
	var products []Products
	for k, v := range temp {
		pts := Products{
			Name: k,
		}
		for _, p := range v {
			if p.DataSource != "" {
				pts.DataSources = append(pts.DataSources, p.DataSource)
			}
			if p.Resource != "" {
				pts.Resources = append(pts.Resources, p.Resource)
			}
		}

		products = append(products, pts)
	}

	//sort
	sortProducts := make(map[string]Products)
	var (
		sortKeys     []string
		productsSort []Products
	)
	for _, _p := range products {
		sortProducts[_p.Name] = _p
		sortKeys = append(sortKeys, _p.Name)
	}
	sort.Strings(sortKeys)

	for _, k := range sortKeys {
		productsSort = append(productsSort, sortProducts[k])
	}

	data := map[string]interface{}{
		"cloud_mark":  cloudMark,
		"cloud_title": cloudTitle,
		"cloudPrefix": cloudPrefix,
		"Products":    productsSort,
	}

	filename := pwd + docRoot + "/" + fmt.Sprintf("%s.erb", cloudMark)
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		message("[FAIL!]open file %s failed: %s", filename, err)
		os.Exit(1)
	}

	defer fd.Close()

	tmpl := template.Must(template.New("t").Funcs(template.FuncMap{"replace": replace}).Parse(idxTPL))

	if err := tmpl.Execute(fd, data); err != nil {
		message("[FAIL!]write file %s failed: %s", filename, err)
		os.Exit(1)
	}

	message("[SUCC.]write doc to file success: %s", filename)
	return prods
}

// genDoc generating doc for data source and resource
func genDoc(pwd string, product, docType, filename, name string, resource *schema.Resource, example string) {
	if resource == nil {
		return
	}
	data := map[string]string{
		"product":           product,
		"name":              name,
		"docType":           strings.Replace(docType, "_", "", -1),
		"resource":          name[len(cloudMark)+1:],
		"cloud_mark":        cloudMark,
		"cloud_title":       cloudTitle,
		"example":           "",
		"description":       "",
		"description_short": "",
		"import":            "",
	}

	//filename := fmt.Sprintf("%s_%s_%s.go", docType, cloudMarkShort, data["resource"])
	//message("[START]get description from file: %s\n", filename)

	description, err := getFileDescription(filename)
	if err != nil {
		message("[FAIL!]get description failed: %s", err)
		os.Exit(1)
	}

	description = strings.TrimSpace(description)
	if description == "" && docType == "resource" {
		message("[FAIL!]description empty: %s\n", filename)
		os.Exit(1)
	}

	importPos := strings.Index(description, "Import\n")
	if importPos != -1 {
		data["import"] = strings.TrimSpace(description[importPos+7:])
		description = strings.TrimSpace(description[:importPos])
	}
	var (
		buf []byte
		hcl string
	)

	buf, err = ioutil.ReadFile(example + "/main.tf")
	if err != nil {
		message("[FAIL!]example usage error: %s\n", example)
	}
	hcl = "```hcl\n" + string(buf) + "\n" + "```"
	data["example"] = formatHCL(hcl)

	if docType == "resource" {
		data["description"] = "Provides a resource to manage " + common.DownLineToSpace(data["resource"])
		data["description_short"] = data["description"]
	} else {
		data["description"] = "Use this data source to query detailed information of " + common.DownLineToSpace(data["resource"])
		data["description_short"] = data["description"]
	}

	var (
		requiredArgs []string
		optionalArgs []string
		attributes   []string
		subStruct    []string
	)

	for k, v := range resource.Schema {
		if v.Description == "" {
			message("[FAIL!]description for '%s' is missing: %s\n", k, filename)
			os.Exit(1)
		} else {
			checkDescription(k, v.Description)
		}
		if docType == "data_source" && v.ForceNew {
			message("[FAIL!]Don't set ForceNew on data source: '%s'", k)
			os.Exit(1)
		}
		if v.Required && v.Optional {
			message("[FAIL!]Don't set Required and Optional at the same time: '%s'", k)
			os.Exit(1)
		}
		if v.Required {
			opt := "Required"
			if v.ForceNew {
				opt += ", ForceNew"
			}
			if v.Deprecated != "" {
				opt += ", **Deprecated**"
				v.Description = fmt.Sprintf("%s %s", v.Deprecated, v.Description)
			}
			requiredArgs = append(requiredArgs, fmt.Sprintf("* `%s` - (%s) %s", k, opt, v.Description))
			subStruct = append(subStruct, getSubStruct(0, k, v)...)
		} else if v.Optional {
			opt := "Optional"
			if v.ForceNew {
				opt += ", ForceNew"
			}
			if v.Deprecated != "" {
				opt += ", **Deprecated**"
				v.Description = fmt.Sprintf("%s %s", v.Deprecated, v.Description)
			}
			optionalArgs = append(optionalArgs, fmt.Sprintf("* `%s` - (%s) %s", k, opt, v.Description))
			subStruct = append(subStruct, getSubStruct(0, k, v)...)
		} else {
			attrs := getAttributes(0, k, v)
			if len(attrs) > 0 {
				attributes = append(attributes, attrs...)
			}
		}
	}

	sort.Strings(requiredArgs)
	sort.Strings(optionalArgs)
	sort.Strings(attributes)
	sort.Strings(subStruct)

	// remove duplicates
	if len(subStruct) > 0 {
		uniqSubStruct := make([]string, 0, len(subStruct))
		var i int
		for i = 0; i < len(subStruct)-1; i++ {
			if subStruct[i] != subStruct[i+1] {
				uniqSubStruct = append(uniqSubStruct, subStruct[i])
			}
		}
		uniqSubStruct = append(uniqSubStruct, subStruct[i])
		subStruct = uniqSubStruct
	}

	requiredArgs = append(requiredArgs, optionalArgs...)
	data["arguments"] = strings.Join(requiredArgs, "\n")
	if len(subStruct) > 0 {
		data["arguments"] += "\n" + strings.Join(subStruct, "\n")
	}
	data["attributes"] = strings.Join(attributes, "\n")
	if docType == "resource" {
		idAttribute := "* `id` - ID of the resource.\n"
		data["attributes"] = idAttribute + data["attributes"]
	}

	if docType == "resource" {
		filename = pwd + docRoot + "/docs/r/" + fmt.Sprintf("%s.html.markdown", data["resource"])
	} else {
		filename = pwd + docRoot + "/docs/d/" + fmt.Sprintf("%s.html.markdown", data["resource"])
	}

	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		message("[FAIL!]open file %s failed: %s", filename, err)
		os.Exit(1)
	}

	defer fd.Close()
	t := template.Must(template.New("t").Parse(docTPL))
	err = t.Execute(fd, data)
	if err != nil {
		message("[FAIL!]write file %s failed: %s", filename, err)
		os.Exit(1)
	}

	message("[SUCC.]write doc to file success: %s", filename)
}

// getAttributes get attributes from schema
func getAttributes(step int, k string, v *schema.Schema) []string {
	var attributes []string
	ident := strings.Repeat(" ", step*2)

	if v.Description == "" {
		return attributes
	} else {
		checkDescription(k, v.Description)
	}

	if v.Computed {
		if v.Deprecated != "" {
			v.Description = fmt.Sprintf("(**Deprecated**) %s %s", v.Deprecated, v.Description)
		}
		if _, ok := v.Elem.(*schema.Resource); ok {
			var listAttributes []string
			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				attrs := getAttributes(step+1, kk, vv)
				if len(attrs) > 0 {
					listAttributes = append(listAttributes, attrs...)
				}
			}
			var slistAttributes string
			sort.Strings(listAttributes)
			if len(listAttributes) > 0 {
				slistAttributes = "\n" + strings.Join(listAttributes, "\n")
			}
			attributes = append(attributes, fmt.Sprintf("%s* `%s` - %s%s", ident, k, v.Description, slistAttributes))
		} else {
			attributes = append(attributes, fmt.Sprintf("%s* `%s` - %s", ident, k, v.Description))
		}
	}

	return attributes
}

// getFileDescription get description from go file
func getFileDescription(fname string) (string, error) {
	fset := token.NewFileSet()

	parsedAst, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	if len(parsedAst.Comments) > 0 {
		return parsedAst.Comments[0].Text(), nil
	}

	return "", nil
}

// getSubStruct get sub structure from go file
func getSubStruct(step int, k string, v *schema.Schema) []string {
	var subStructs []string

	if v.Description == "" {
		return subStructs
	} else {
		checkDescription(k, v.Description)
	}

	var subStruct []string
	if v.Type == schema.TypeMap || v.Type == schema.TypeList || v.Type == schema.TypeSet {
		if _, ok := v.Elem.(*schema.Resource); ok {
			subStruct = append(subStruct, fmt.Sprintf("\nThe `%s` object supports the following:\n", k))
			var (
				requiredArgs []string
				optionalArgs []string
			)
			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				if vv.Required {
					opt := "Required"
					if vv.ForceNew {
						opt += ", ForceNew"
					}
					requiredArgs = append(requiredArgs, fmt.Sprintf("* `%s` - (%s) %s", kk, opt, vv.Description))
				} else if vv.Optional {
					opt := "Optional"
					if vv.ForceNew {
						opt += ", ForceNew"
					}
					optionalArgs = append(optionalArgs, fmt.Sprintf("* `%s` - (%s) %s", kk, opt, vv.Description))
				}
			}
			sort.Strings(requiredArgs)
			subStruct = append(subStruct, requiredArgs...)
			sort.Strings(optionalArgs)
			subStruct = append(subStruct, optionalArgs...)
			subStructs = append(subStructs, strings.Join(subStruct, "\n"))

			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				subStructs = append(subStructs, getSubStruct(step+1, kk, vv)...)
			}
		}
	}

	return subStructs
}

// formatHCL format HLC code
func formatHCL(s string) string {
	var rr []string

	s = strings.TrimSpace(s)
	m := hclMatch.FindAllStringSubmatch(s, -1)
	if len(m) > 0 {
		for _, v := range m {
			p := strings.TrimSpace(v[1])
			if p != "" {
				p = fmt.Sprintf("\n%s\n\n", p)
			}
			b := hclwrite.Format([]byte(strings.TrimSpace(v[3])))
			rr = append(rr, fmt.Sprintf("%s```hcl\n%s\n```", p, string(b)))
		}
	}

	return strings.TrimSpace(strings.Join(rr, "\n"))
}

// checkDescription check description format
func checkDescription(k, s string) {
	if s == "" {
		return
	}

	if strings.TrimLeft(s, " ") != s {
		message("[FAIL!]There is space on the left of description: '%s': '%s'", k, s)
		os.Exit(1)
	}

	if strings.TrimRight(s, " ") != s {
		message("[FAIL!]There is space on the right of description: '%s': '%s'", k, s)
		os.Exit(1)
	}

	if s[len(s)-1] != '.' && s[len(s)-1] != ':' {
		message("[FAIL!]There is no ending charset(. or :) on the description: '%s': '%s'", k, s)
		os.Exit(1)
	}

	if c := containsBigSymbol(s); c != "" {
		message("[FAIL!]There is unexcepted symbol '%s' on the description: '%s': '%s'", c, k, s)
		os.Exit(1)
	}

	for _, v := range []string{",", ".", ";", ":", "?", "!"} {
		if strings.Contains(s, " "+v) {
			message("[FAIL!]There is space before '%s' on the description: '%s': '%s'", v, k, s)
			os.Exit(1)
		}
	}
}

// containsBigSymbol returns the Big symbol if found
func containsBigSymbol(s string) string {
	m := bigSymbol.FindStringSubmatch(s)
	if len(m) > 0 {
		return m[0]
	}

	return ""
}

// message print color message
func message(msg string, v ...interface{}) {
	if strings.Contains(msg, "FAIL") {
		color.Red(fmt.Sprintf(msg, v...))
	} else if strings.Contains(msg, "SUCC") {
		color.Green(fmt.Sprintf(msg, v...))
	} else if strings.Contains(msg, "SKIP") {
		color.Yellow(fmt.Sprintf(msg, v...))
	} else {
		color.White(fmt.Sprintf(msg, v...))
	}
}
