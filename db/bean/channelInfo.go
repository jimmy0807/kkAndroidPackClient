package bean

type PackageApp struct {
	ApkName           string `json:"apk_name"`
	ApkLastUpdateTime string `json:"apk_last_update_time"`
	Status            string `json:"status"`
	ChannelName       string `json:"channel_name"`
	ChannelID         int64  `json:"channel_id"`
}
