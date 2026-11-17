// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	GRPCListenOn string
	MySQL        *MySQLConfig `json:",optional"`
}

type MySQLConfig struct {
	Host          string
	Port          int `json:",default=3306,range=[1:65535]"`
	User          string
	Database      string
	TLSSkipVerify bool `json:",optional"`
}
