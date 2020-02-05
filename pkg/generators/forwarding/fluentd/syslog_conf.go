package fluentd

import (
	"strings"
)

func (conf *outputLabelConf) LegacySyslogPlugin() bool {
	if conf.Target.Secret != nil {
		return false
	}
	if conf.protocolSet() {
		return false
	}
	return true
}

func (conf *outputLabelConf) protocolSet() bool {
	endpoint := conf.Target.Endpoint
	return strings.Index(endpoint, protocolSeparator) != -1
}

func (conf *outputLabelConf) SyslogProtocol() string {
	if conf.protocolSet() {
		endpoint := conf.Target.Endpoint
		return strings.Split(endpoint, protocolSeparator)[0]
	}
	return "tcp"
}

func (conf *outputLabelConf) SyslogPlugin() string {
	if conf.LegacySyslogPlugin() {
		return "syslog_buffered"
	}
	return "remote_syslog"
}
