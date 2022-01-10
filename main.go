package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/fiatjaf/lightningd-gjson-rpc/plugin"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment

	p := plugin.Plugin{
		Name:    "redis-publisher",
		Version: "v0.1",
		Options: []plugin.Option{
			{
				Name:        "test",
				Type:        "string",
				Description: "test string",
			},
		},
		Subscriptions: []plugin.Subscription{
			subscription("channel_opened"),
			subscription("channel_open_failed"),
			subscription("channel_state_changed"),
			subscription("connect"),
			subscription("disconnect"),
			subscription("invoice_payment"),
			subscription("invoice_creation"),
			subscription("warning"),
			subscription("forward_event"),
			subscription("sendpay_success"),
			subscription("sendpay_failure"),
			subscription("coin_movement"),
			subscription("openchannel_peer_sigs"),
		},
		Dynamic: true,
	}

	p.Run()
}

func subscription(kind string) plugin.Subscription {
	return plugin.Subscription{
		Type: kind,
		Handler: func(p *plugin.Plugin, params plugin.Params) {
			fmt.Println(params)
		},
	}
}
