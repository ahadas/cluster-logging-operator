package fluentd

import ()

func (conf *outputLabelConf) LegacySyslogPlugin() bool {
	if conf.Target.Secret != nil {
		return false
	}
	return true
}

func (conf *outputLabelConf) SyslogPlugin() string {
	if conf.LegacySyslogPlugin() {
		return "syslog_buffered"
	}
	return "remote_syslog"
}
