module github.com/volcengine/terraform-provider-vestack

go 1.12

require (
	github.com/fatih/color v1.7.0
	github.com/google/uuid v1.3.0
	github.com/hashicorp/hcl/v2 v2.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/stretchr/testify v1.7.0
	github.com/volcengine/volcstack-go-sdk v1.0.2
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
)

//以后正式开源后，直接在required的位置依赖github上的sdk 不在用这种gitlib的替换方式
//后续开发过程中可以用这种replace方式 但是replace在提交过程不要提交到版本库
replace github.com/volcengine/volcstack-go-sdk => code.byted.org/iaasng/volcstack-go-sdk v1.0.1-0.20220526125608-ba93ad6e0dd8
