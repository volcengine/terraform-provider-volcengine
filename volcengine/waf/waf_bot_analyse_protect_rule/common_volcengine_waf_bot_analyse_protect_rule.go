package waf_bot_analyse_protect_rule

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
)

func wafBotAnalyseProtectRuleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'ID:BotSpace:Host'")
	}
	id := items[0]
	botSpace := items[1]
	host := items[2]
	ccRuleIdIInt, err := strconv.Atoi(id)
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(" ID cannot convert to int ")
	}
	_ = d.Set("id", ccRuleIdIInt)
	_ = d.Set("host", host)
	_ = d.Set("bot_space", botSpace)

	return []*schema.ResourceData{d}, nil
}
