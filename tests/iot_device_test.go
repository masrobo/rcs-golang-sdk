package masrobo_test

import (
	"context"
	"os"
	"testing"

	"github.com/masrobo/rcs-golang-sdk/masrobo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// Config 从 config.yaml 读取的配置结构
type Config struct {
	BaseURL     string `yaml:"base_url"`
	AppID       string `yaml:"app_id"`
	AppKey      string `yaml:"app_key"`
	DeviceID    string `yaml:"device_id"`
	ProductName string `yaml:"product_name"`
}

// loadConfig 从 config.yaml 文件加载配置
func loadConfig() (*Config, error) {
	data, err := os.ReadFile("../config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// createController 使用 config.yaml 中的配置创建控制器
func createController(t *testing.T) (*masrobo.RcsController, *Config) {
	t.Helper()

	cfg, err := loadConfig()
	require.NoError(t, err, "读取 config.yaml 失败，请确保文件存在且格式正确")

	controller, err := masrobo.NewRcsController(masrobo.Config{
		BaseURL: cfg.BaseURL,
		AppID:   cfg.AppID,
		AppKey:  cfg.AppKey,
	})
	require.NoError(t, err, "创建 RcsController 失败")

	return controller, cfg
}

// TestGetDeviceInfo 获取设备信息，验证返回的关键字段
func TestGetDeviceInfo(t *testing.T) {
	controller, cfg := createController(t)

	result, err := controller.IotDevice.GetDeviceInfo(context.Background(), masrobo.DeviceInfoRequest{
		DeviceID: cfg.DeviceID,
	})
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, cfg.DeviceID, result.DeviceID)
	assert.Equal(t, cfg.ProductName, result.ProductName)
	assert.NotEmpty(t, result.DeviceName, "device_name 不应为空")
	assert.True(t, result.Status >= 0, "status 应大于等于 0")

	t.Logf("设备名称: %s", result.DeviceName)
	t.Logf("产品名称: %s", result.ProductName)
	t.Logf("设备状态: %d", result.Status)
}

// TestAddDevice 添加设备，验证返回二维码信息
func TestAddDevice(t *testing.T) {
	controller, cfg := createController(t)

	result, err := controller.IotDevice.AddDevice(context.Background(), masrobo.AddDeviceRequest{
		ProjectName: cfg.ProductName,
		DeviceID:    cfg.DeviceID,
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.NotEmpty(t, result.QrcodeUrl, "qrcode_url 不应为空")

	t.Logf("AddDevice 调用成功: qrcode_url=%s", result.QrcodeUrl)
}
