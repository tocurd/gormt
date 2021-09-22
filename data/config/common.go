package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gopkg.in/yaml.v3"
)

// CfgBase base config struct
type CfgBase struct {
	// SerialNumber       string `json:"serial_number" yaml:"serial_number"`             // version.版本号
	// ServiceName        string `json:"service_name" yaml:"service_name"`               // service name .service名字
	// ServiceDisplayname string `json:"service_displayname" yaml:"service_displayname"` // display name .显示名
	// SerciceDesc        string `json:"sercice_desc" yaml:"sercice_desc"`               // sercice desc .service描述
	IsDev bool `json:"is_dev" yaml:"is_dev"` // Is it a development version?是否是开发版本
}

var Map = Config{
	CfgBase: CfgBase{
		IsDev: false,
	},
	Database:         map[string]DBInfo{},
	OutDir:           "./model",
	URLTag:           "json",
	Language:         "中 文",
	DbTag:            "gorm",
	Simple:           false,
	IsWEBTag:         true,
	IsWebTagPkHidden: false,
	// SingularTable: true,
	IsForeignKey:         true,
	IsOutSQL:             false,
	IsOutFunc:            false,
	IsGUI:                false,
	IsTableName:          true,
	TablePrefix:          "",
	SelfTypeDef:          make(map[string]string),
	WebTagType:           0,
	TableNames:           "",
	IsColumnName:         false,
	IsOutFileByTableName: false,
	Suffix:               "",
	OutFileName:          "model",
}

var configPath string

func init() {
	Map.SelfTypeDef["datatime"] = "time.Time"
	Map.SelfTypeDef["time"] = "time.Time"
	configPath = path.Join(tools.GetCurrentDirectory(), "config.yml") // 先找本程序文件夹
	if !tools.CheckFileIsExist(configPath) {                          // dont find it
		configPath = path.Join(tools.GetModelPath(), "config.yml")
		if !tools.CheckFileIsExist(configPath) {
			mylog.ErrorString("config.yml not exit. using default config")
		}
	}

	onInit()
	dev.OnSetDev(Map.IsDev)
}

func onInit() {
	err := InitFile(configPath)
	if err != nil {
		fmt.Println("Load config file error: ", err.Error())
		return
	}
}

// InitFile default value from file .
func InitFile(filename string) error {
	// if _, e := os.Stat(filename); e != nil {
	// 	fmt.Println("init default config file: ", filename)
	// 	if err := SaveToFile(); err == nil {
	// 		InitFile(filename)
	// 		return nil
	// 	} else {
	// 		fmt.Println("shit,fail", err)
	// 	}
	// 	// os.Exit(0)
	// }
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bs, &Map); err != nil {
		fmt.Println("read config file error: ", err.Error())
		return err
	}
	return nil
}

// GetServiceConfig Get service configuration information
// func GetServiceConfig() (name, displayName, desc string) {
// 	name = Map.ServiceName
// 	displayName = Map.ServiceDisplayname
// 	desc = Map.SerciceDesc
// 	return
// }

// GetIsDev is is dev
func GetIsDev() bool {
	return Map.IsDev
}

// SetIsDev is is dev
func SetIsDev(b bool) {
	Map.IsDev = b
}

// SaveToFile save config info to file
func SaveToFile() error {
	d, err := yaml.Marshal(Map)
	if err != nil {
		return err
	}
	tools.WriteFile(configPath, []string{
		string(d),
	}, true)
	return nil
}
