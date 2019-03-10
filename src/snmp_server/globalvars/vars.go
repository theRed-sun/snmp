package globalvars

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"snmp_server/mibs"
	"snmp_server/mibs/warning"
	"snmp_server/model"
	"time"

	"github.com/cihub/seelog"

	"github.com/beevik/ntp"
)

//FTPDir ftp dir
var FTPDir = "/home/klsnmp/ftpfile/"

//FTPUser ftp User
var FTPUser = "uftp"

//FTPGroup ftp group
var FTPGroup = "uftp"

//FTPUID ftp user uid
var FTPUID = 0

//FTPGID ftp group  gid
var FTPGID = 0

//UpgradeDir 防止升级文件的目录
var UpgradeDir = "/home/klsnmp/upgrade"

//ShowSQL show sql for xorm
var ShowSQL = false

//StartTime 系统启动时间
var StartTime = time.Now()

var (
	//AppVersion app 版本号
	AppVersion = "0.0.1"
	//AppBuildTime 编译时间
	AppBuildTime = "2017-12-01T00:03:18+0800"
	//AppGitHash git hash
	AppGitHash = "undefined"
	//GoVersion go version
	GoVersion = "undefined"
)

//GetFTPChown chown for ftp user:group
func GetFTPChown() string {
	return fmt.Sprintf("%s:%s", FTPUser, FTPGroup)
}

//Configure config for snmp_server
type Configure struct {
	NTPDEnable bool   `json:"ntpd_enable"`
	NTPServer1 string `json:"ntp_server_1"`
	NTPServer2 string `json:"ntp_server_2"`
	WebPort    int    `json:"web_port"`
	SnmpPort   int    `json:"snmp_port"`

	doWarning func(mibs.IWarning)
}

//SetDoWarning set DoWarning callback
func (c *Configure) SetDoWarning(callback func(mibs.IWarning)) {
	c.doWarning = callback
}

//Exec  后端线程
func (c *Configure) Exec() {

	seelog.Infof("Start Configure, NTPDEnable %t", c.NTPDEnable)
	ntpErrorStatus := 0
	for {
		select {
		case <-time.After(60 * time.Second):
			seelog.Infof("NTP Enable %t", c.NTPDEnable)
			if c.NTPDEnable {
				time, err := ntp.Time(c.NTPServer1)
				if err == nil {
					if ntpErrorStatus > 0 {
						logInfo := &model.LogInfo{
							User:     "ntp",
							NTID:     "NA",
							Event:    "ntp",
							SubEvent: "sync_ok",
							Info:     "ntp recover",
						}
						logInfo.Insert()
						ntpWarning2 := &warning.NtpWarning{
							NTID:   "system_ntp",
							Clear:  1,
							Status: 1,
							Demo:   "ntp sync ok.",
						}
						c.doWarning(ntpWarning2)
					}
					ntpErrorStatus = 0
					seelog.Infof("get time (%s) from NTPServer1(%s)", time.Format("01/02/2006 15:04:05"), c.NTPServer1)
					cmd := exec.Command("date", "-s", time.Format("01/02/2006 15:04:05"))
					cmd.Run()
				} else {
					time, err = ntp.Time(c.NTPServer2)
					if err == nil {
						if ntpErrorStatus > 0 {
							logInfo := &model.LogInfo{
								User:     "ntp",
								NTID:     "NA",
								Event:    "ntp",
								SubEvent: "sync_ok",
								Info:     "ntp recover",
							}
							logInfo.Insert()

							ntpWarning2 := &warning.NtpWarning{
								NTID:   "system_ntp",
								Clear:  1,
								Status: 1,
								Demo:   "ntp sync ok.",
							}
							c.doWarning(ntpWarning2)
						}
						ntpErrorStatus = 0
						seelog.Infof("get time (%s) from NTPServer2(%s)", time.Format("01/02/2006 15:04:05"), c.NTPServer2)
						cmd := exec.Command("date", "-s", time.Format("01/02/2006 15:04:05"))
						cmd.Run()
					} else {
						ntpErrorStatus++
						if ntpErrorStatus == 1 {

							ntpWarning2 := &warning.NtpWarning{
								NTID:   "system_ntp",
								Clear:  0,
								Status: 1,
								Demo:   "ntp sync error.",
							}
							c.doWarning(ntpWarning2)
							logInfo := &model.LogInfo{
								User:     "ntp",
								NTID:     "NA",
								Event:    "ntp",
								SubEvent: "sync_err",
								Info:     err.Error(),
							}
							logInfo.Insert()
						}
						seelog.Errorf("get time error:%v", err)
					}
				}
			}
		}
	}
}

//Load load config form local
func (c *Configure) Load() {

	jdata, err := ioutil.ReadFile("configure.json")
	if err == nil {
		json.Unmarshal(jdata, c)
	} else {
		c.Save()
	}
}

//Save save to local
func (c *Configure) Save() error {

	jdata, _ := json.MarshalIndent(c, "", " ")
	err := ioutil.WriteFile("configure.json", jdata, os.ModePerm)

	return err
}

func (c *Configure) String() string {
	jdata, _ := json.Marshal(c)
	return string(jdata)
}

//Default default configure
var Default *Configure

func init() {

	Default = &Configure{
		NTPDEnable: false,
		WebPort:    9192,
		SnmpPort:   162,
	}

	Default.Load()
	go Default.Exec()
}
