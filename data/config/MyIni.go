package config

import (
	"fmt"
	"strings"

	"github.com/xxjwxc/public/tools"
)

// Config custom config struct
type Config struct {
	CfgBase              `yaml:"base"`
	Database             DBInfo            `yaml:"database"`
	OutDir               string            `yaml:"out_dir"`
	URLTag               string            `yaml:"url_tag"`  // url tag
	Language             string            `yaml:"language"` // language
	DbTag                string            `yaml:"db_tag"`   // 数据库标签（gormt,db）
	Simple               bool              `yaml:"simple"`
	IsWEBTag             bool              `yaml:"is_web_tag"`
	IsWebTagPkHidden     bool              `yaml:"is_web_tag_pk_hidden"` // web标记是否隐藏主键
	IsForeignKey         bool              `yaml:"is_foreign_key"`
	IsOutSQL             bool              `yaml:"is_out_sql"`
	IsOutFunc            bool              `yaml:"is_out_func"`
	IsGUI                bool              `yaml:"is_gui"` //
	IsTableName          bool              `yaml:"is_table_name"`
	IsNullToPoint        bool              `yaml:"is_null_to_point"` // null to porint
	TablePrefix          string            `yaml:"table_prefix"`     // 表前缀
	SelfTypeDef          map[string]string `yaml:"self_type_define"`
	OutFileName          string            `yaml:"out_file_name"`
	WebTagType           int               `yaml:"web_tag_type"`              // 默认小驼峰
	TableNames           string            `yaml:"table_names"`               // 表名（多个表名用","隔开）
	IsColumnName         bool              `yaml:"is_column_name"`            //是否输出列名
	IsOutFileByTableName bool              `yaml:"is_out_file_by_table_name"` //是否根据表名生成文件(多个表名生成多个文件)
}

// DBInfo mysql database information. mysql 数据库信息
type DBInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    // Port 端口号
	Username string // Username 用户名
	Password string // Password 密码
	Database string // Database 数据库名
	Type     int    // 数据库类型: 0:mysql , 1:sqlite , 2:mssql
}

// SetMysqlDbInfo Update MySQL configuration information
func SetMysqlDbInfo(info *DBInfo) {
	Map.Database = *info
}

// GetDbInfo Get configuration information .获取数据配置信息
func GetDbInfo() DBInfo {
	return Map.Database
}

// GetMysqlConStr Get MySQL connection string.获取mysql 连接字符串
func GetMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		Map.Database.Username,
		Map.Database.Password,
		Map.Database.Host,
		Map.Database.Port,
		Map.Database.Database,
	)
}

// SetOutDir Setting Output Directory.设置输出目录
func SetOutDir(outDir string) {
	Map.OutDir = outDir
}

// GetOutDir Get Output Directory.获取输出目录
func GetOutDir() string {
	if len(Map.OutDir) == 0 {
		Map.OutDir = "./model"
	}

	return Map.OutDir
}

// // SetSingularTable Set Disabled Table Name Plurals.设置禁用表名复数
// func SetSingularTable(b bool) {
// 	Map.SingularTable = b
// }

// // GetSingularTable Get Disabled Table Name Plurals.获取禁用表名复数
// func GetSingularTable() bool {
// 	return Map.SingularTable
// }

// GetSimple simple output.简单输出
func GetSimple() bool {
	return Map.Simple
}

// SetSimple simple output.简单输出
func SetSimple(b bool) {
	Map.Simple = b
}

// GetIsWEBTag json tag.json标记
func GetIsWEBTag() bool {
	return Map.IsWEBTag
}

// SetIsWEBTag json tag.json标记
func SetIsWEBTag(b bool) {
	Map.IsWEBTag = b
}

// GetIsWebTagPkHidden web tag是否隐藏主键
func GetIsWebTagPkHidden() bool {
	return Map.IsWebTagPkHidden
}

// GetIsForeignKey if is foreign key
func GetIsForeignKey() bool {
	return Map.IsForeignKey
}

// SetForeignKey Set if is foreign key.设置是否外键关联
func SetForeignKey(b bool) {
	Map.IsForeignKey = b
}

// SetIsOutSQL if is output sql .
func SetIsOutSQL(b bool) {
	Map.IsOutSQL = b
}

// GetIsOutSQL if is output sql .
func GetIsOutSQL() bool {
	return Map.IsOutSQL
}

// GetIsOutFunc if is output func .
func GetIsOutFunc() bool {
	return Map.IsOutFunc
}

// SetIsOutFunc if is output func .
func SetIsOutFunc(b bool) {
	Map.IsOutFunc = b
}

// GetIsGUI if is gui show .
func GetIsGUI() bool {
	return Map.IsGUI
}

// SetIsGUI if is gui show .
func SetIsGUI(b bool) {
	Map.IsGUI = b
}

// GetIsTableName if is table name .
func GetIsTableName() bool {
	return Map.IsTableName
}

// SetIsTableName if is table name .
func SetIsTableName(b bool) {
	Map.IsTableName = b
}

// GetURLTag get url tag.
func GetURLTag() string {
	if Map.URLTag != "json" && Map.URLTag != "url" {
		Map.URLTag = "json"
	}

	return Map.URLTag
}

// SetURLTag set url tag.
func SetURLTag(s string) {
	Map.URLTag = s
}

// GetLG get language tag.
func GetLG() string {
	if Map.Language != "English" && Map.Language != "中 文" {
		if tools.GetLocalSystemLang(true) == "en" {
			Map.Language = "English"
		} else {
			Map.Language = "中 文"
		}
	}

	return Map.Language
}

// SetLG set url tag.
func SetLG(s string) {
	Map.Language = s
}

// GetDBTag get database tag.
func GetDBTag() string {
	if Map.DbTag != "gorm" && Map.DbTag != "db" {
		Map.DbTag = "gorm"
	}

	return Map.DbTag
}

// SetDBTag get database tag.
func SetDBTag(s string) {
	Map.DbTag = s
}

// SetIsNullToPoint if with null to porint in struct
func SetIsNullToPoint(b bool) {
	Map.IsNullToPoint = b
}

// GetIsNullToPoint get if with null to porint in sturct
func GetIsNullToPoint() bool {
	return Map.IsNullToPoint
}

// SetTablePrefix set table prefix
func SetTablePrefix(t string) {
	Map.TablePrefix = t
}

// GetTablePrefix get table prefix
func GetTablePrefix() string {
	return Map.TablePrefix
}

// SetSelfTypeDefine 设置自定义字段映射
func SetSelfTypeDefine(data map[string]string) {
	Map.SelfTypeDef = data
}

// GetSelfTypeDefine 获取自定义字段映射
func GetSelfTypeDefine() map[string]string {
	return Map.SelfTypeDef
}

// SetOutFileName 设置输出文件名
func SetOutFileName(s string) {
	Map.OutFileName = s
}

// GetOutFileName 获取输出文件名
func GetOutFileName() string {
	return Map.OutFileName
}

// SetWebTagType 设置json tag类型
func SetWebTagType(i int) {
	Map.WebTagType = i
}

// GetWebTagType 获取json tag类型
func GetWebTagType() int {
	return Map.WebTagType
}

//GetTableNames get format tableNames by config. 获取格式化后设置的表名
func GetTableNames() string {
	var sb strings.Builder
	if Map.TableNames != "" {
		tableNames := Map.TableNames
		tableNames = strings.TrimLeft(tableNames, ",")
		tableNames = strings.TrimRight(tableNames, ",")
		if tableNames == "" {
			return ""
		}

		sarr := strings.Split(Map.TableNames, ",")
		if len(sarr) == 0 {
			fmt.Printf("tableNames is vailed, genmodel will by default global")
			return ""
		}

		for i, val := range sarr {
			sb.WriteString(fmt.Sprintf("'%s'", val))
			if i != len(sarr)-1 {
				sb.WriteString(",")
			}
		}
	}
	return sb.String()
}

//GetOriginTableNames get origin tableNames. 获取原始的设置的表名
func GetOriginTableNames() string {
	return Map.TableNames
}

//SetTableNames set tableNames. 设置生成的表名
func SetTableNames(tableNames string) {
	Map.TableNames = tableNames
}

//GetIsColumnName get  gen columnName config . 获取生成列名的config
func GetIsColumnName() bool {
	return Map.IsColumnName
}

//SetIsColumnName set gen ColumnName config. 设置生成列名的config
func SetIsColumnName(isColumnName bool) {
	Map.IsColumnName = isColumnName
}

//GetIsOutFileByTableName get  gen columnName config . 设置是否根据表名生成文件
func GetIsOutFileByTableName() bool {
	return Map.IsOutFileByTableName
}

//SetIsOutFileByTableName set gen ColumnName config. 设置是否根据表名生成文件
func SetIsOutFileByTableName(isOutFileByTableName bool) {
	Map.IsColumnName = isOutFileByTableName
}
