package masrobo

import "time"

// IotDeviceInfo is the device information returned by the Open API.
type IotDeviceInfo struct {
	ID              int64          `json:"id"`
	DeviceID        string         `json:"device_id"`
	IotProductID    int64          `json:"iot_product_id"`
	ProductName     string         `json:"product_name"`
	CategoryCode    string         `json:"category_code"`
	UserID          int64          `json:"user_id"`
	DeviceName      string         `json:"device_name"`
	ActiveType      []string       `json:"active_type"`
	Status          int            `json:"status"`
	ActiveTime      time.Time      `json:"active_time"`
	DeactiveTime    *time.Time     `json:"deactive_time,omitempty"`
	ApplicationID   int64          `json:"application_id"`
	DeveloperID     int64          `json:"developer_id"`
	AppID           string         `json:"app_id"`
	ApplicationName string         `json:"application_name"`
	DeveloperName   string         `json:"developer_name"`
	ProductTitle    string         `json:"product_title"`
	DeviceData      map[string]any `json:"device_data"`
	DeviceDataTime  *time.Time     `json:"device_data_time,omitempty"`
	Screenshot      string         `json:"screenshot"`
	ScreenshotTime  *time.Time     `json:"screenshot_time,omitempty"`
	Services        map[string]any `json:"services,omitempty"`
	RelRole         int8           `json:"rel_role"`
	RelSource       string         `json:"rel_source"`
}

// LatestDeviceDataRecord is the latest data record returned by the Open API.
type LatestDeviceDataRecord struct {
	RawTopicName string    `json:"raw_topic_name"`
	Payload      any       `json:"payload,omitempty"`
	URL          string    `json:"url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// GetLatestDeviceDataResponse is the SDK response for the latest device data API.
type GetLatestDeviceDataResponse struct {
	ProductName string                  `json:"product_name"`
	DeviceID    string                  `json:"device_id"`
	TopicName   string                  `json:"topic_name"`
	Record      *LatestDeviceDataRecord `json:"record"`
}
