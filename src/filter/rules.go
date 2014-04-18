/*
  btcbot is a Bitcoin trading bot for HUOBI.com written
  in golang, it features multiple trading methods using
  technical analysis.

  Disclaimer:

  USE AT YOUR OWN RISK!

  The author of this project is NOT responsible for any damage or loss caused
  by this software. There can be bugs and the bot may not perform as expected
  or specified. Please consider testing it first with paper trading /
  backtesting on historical data. Also look at the code to see what how
  it's working.

  Weibo:http://weibo.com/bocaicfa
*/

package filter

import (
	"regexp"
)

func Rule(uri string) map[string]map[string]map[string]string {
	if rule, ok := rules[uri]; ok {
		return rule
	}
	for key, rule := range rules {
		reg := regexp.MustCompile(key)
		if reg.MatchString(uri) {
			return rule
		}
	}
	return nil
}

// 定义所有表单验证规则
var rules = map[string]map[string]map[string]map[string]string{
	// 用户注册验证规则
	"/account/register.json": {
		"username": {
			"require": {"error": "用户名不能为空！"},
			"regex":   {"pattern": `^\w*$`, "error": "用户名只能包含大小写字母、数字和下划线"},
			"length":  {"range": "4,20", "error": "用户名长度必须在%d个字符和%d个字符之间"},
		},
		"email": {
			"require": {"error": "邮箱不能为空！"},
			"email":   {"error": "邮箱格式不正确！"},
		},
		"passwd": {
			"require": {"error": "密码不能为空！"},
			"length":  {"range": "6,32", "error": "密码长度必须在%d个字符和%d个字符之间"},
		},
		"pass2": {
			"require": {"error": "确认密码不能为空！"},
			"compare": {"field": "passwd", "rule": "=", "error": "两次密码不一致"},
		},
	},
	// 发新帖
	"/topics/new.json": {
		"content": {
			"require": {"error": "内容不能为空！"},
			"length":  {"range": "3,1024", "error": "话题内容长度必须在%d个字符和%d个字符之间"},
		},
	},
	// 发回复
	`/comment/\d+\.json`: {
		"content": {
			"require": {"error": "内容不能为空！"},
			"length":  {"range": "3,1024", "error": "回复内容长度不能少于%d个字符"},
		},
	},

	// 发消息
	"/message/send.json": {
		"to": {
			"require": {"error": "必须指定发给谁"},
		},
		"content": {
			"require": {"error": "消息内容不能为空"},
		},
	},
	// 删除消息
	"/message/delete.json": {
		"id": {
			"require": {"error": "必须指定id"},
		},
		"msgtype": {
			"require": {"error": "必须指定消息类型"},
		},
	},
}
