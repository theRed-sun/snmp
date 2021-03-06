package xtraphandler

import (
	"net"
	"snmp_server/mibs"
	"snmp_server/mibs/warning"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xtrap"
	"snmp_server/xwarning"
	"strings"

	"github.com/cihub/seelog"
	"github.com/soniah/gosnmp"
)

//DoWarningTest DoWarningTest
func DoWarningTest(msg mibs.IWarning) {
	doWarning(msg)
}

//DoWarning DoWarning
func DoWarning(msg mibs.IWarning) {
	doWarning(msg)
}

func doWarning(msg mibs.IWarning) {
	if msg.IsClear() {
		xwarning.ClearWarning(msg.GetNTID(), msg.WarningType(), xdb.EngineWarning)
		logInfo := &model.LogInfo{
			User:     "system",
			NTID:     msg.GetNTID(),
			Event:    "warning_clear",
			SubEvent: msg.WarningType(),
			Info:     "auto clear:" + msg.GetDemo(),
		}
		logInfo.Insert()
		return
	}

	w, has := xwarning.GetWarning(msg.GetNTID(), msg.WarningType(), xdb.EngineWarning)
	if !has {
		t, err := model.GetTerminalByNTID(msg.GetNTID(), xdb.Engine)
		if err != nil {
			seelog.Errorf("GetTerminalByNTID err:%v", err)
			return
		}
		if t == nil {
			if strings.EqualFold(msg.GetNTID(), "system_ntp") || strings.EqualFold(msg.GetNTID(), "system_disk") {
				t = &model.Terminal{
					ID:   0,
					NTID: msg.GetNTID(),
					Name: msg.GetNTID(),
					Path: "root",
				}
			} else {
				seelog.Errorf("GetTerminalByNTID nil")
				return
			}
		}

		w = &xwarning.Warning{
			TID:    t.ID,
			TName:  t.Name,
			NTID:   t.NTID,
			Path:   t.Path,
			WType:  msg.WarningType(),
			WValue: msg.WarningStatus(),
			WDemo:  msg.GetDemo(),
		}
		err = xwarning.InsertWarning(w, xdb.EngineWarning)
		if err != nil {
			seelog.Errorf("insert %s %s warning error:%v", w.NTID, w.WType, err)
		} else {
			seelog.Infof("insert %s %s warning ok", w.NTID, w.WType)
		}
		logInfo := &model.LogInfo{
			User:     "system",
			NTID:     msg.GetNTID(),
			Event:    "warning_occur",
			SubEvent: msg.WarningType(),
			Info:     msg.GetDemo(),
		}
		logInfo.Insert()
	} else {
		w.WDemo = msg.GetDemo()
		xwarning.UpdateWarning(w, xdb.EngineWarning)
		logInfo := &model.LogInfo{
			User:     "system",
			NTID:     msg.GetNTID(),
			Event:    "warning_occur",
			SubEvent: msg.WarningType(),
			Info:     msg.GetDemo(),
		}
		logInfo.Insert()
	}
}

func deviceMICWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(warning.DeviceMICWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)

}

func deviceSpeakerWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceSpeakerWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceLEDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {

	msg := new(warning.DeviceLEDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func deviceKeypWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceKeypWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func deviceQuickWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceQuickWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceCameraWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceCameraWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceLCDWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceLCDWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func deviceShellWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DeviceShellWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func powerWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.PowerWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func calloutWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.CalloutWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func registerWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.RegisterWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

func drycontactWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.DrycontactWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapMainBackupSwitchWarning = ".1.3.6.1.4.1.1800.50.1.19"
func mainBackupSwitchWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.MianBackupSwitchWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapNtpDropWarning = ".1.3.6.1.4.1.1800.50.1.20"
func netpDropWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.NtpDropWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapRecordingFailWarning = ".1.3.6.1.4.1.1800.50.1.21"
func recordingFaildWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.RecordingFailWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapMainPowerDropWarning = ".1.3.6.1.4.1.1800.50.1.22"
func mainPowerDropWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.MainPowerDropWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapBackupPowerDropWaring = ".1.3.6.1.4.1.1800.50.1.23"
func backupPowerDropWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.BackupPowerDropWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapMainNetDropWarning = ".1.3.6.1.4.1.1800.50.1.24"
func mainNetDropWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.MainNetDropWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapBackupNetDropWarning = ".1.3.6.1.4.1.1800.50.1.25"
func backupNetDropWarning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.BackupNetDropWarning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput1Warning = ".1.3.6.1.4.1.1800.50.1.26"
func usualInput1Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput1Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput2Warning = ".1.3.6.1.4.1.1800.50.1.27"
func usualInput2Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput2Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput3Warning = ".1.3.6.1.4.1.1800.50.1.28"
func usualInput3Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput3Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput4Warning = ".1.3.6.1.4.1.1800.50.1.29"
func usualInput4Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput4Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput5Warning = ".1.3.6.1.4.1.1800.50.1.30"
func usualInput5Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput5Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput6Warning = ".1.3.6.1.4.1.1800.50.1.31"
func usualInput6Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput6Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput7Warning = ".1.3.6.1.4.1.1800.50.1.32"
func usualInput7Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput7Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput8Warning = ".1.3.6.1.4.1.1800.50.1.33"
func usualInput8Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput8Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput9Warning = ".1.3.6.1.4.1.1800.50.1.34"
func usualInput9Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput9Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}

//const trapUsualInput10Warning = ".1.3.6.1.4.1.1800.50.1.35"
func usualInput10Warning(packet *gosnmp.SnmpPacket, remote *net.UDPAddr) {
	msg := new(warning.UsualInput10Warning)
	msg.FromSnmpPackage(packet)
	seelog.Info("recv Warning:", msg)
	doWarning(msg)
}
func init() {

	xtrap.RegisterHandler(trapTypeDeviceMICWarning, deviceMICWarning)
	xtrap.RegisterHandler(trapTypeDeviceSpeakerWaring, deviceSpeakerWarning)
	xtrap.RegisterHandler(trapTypeDeviceLEDWarning, deviceLEDWarning)
	xtrap.RegisterHandler(trapTypeDeviceKeypWarning, deviceKeypWarning)
	xtrap.RegisterHandler(trapTypeDeviceQuickWarning, deviceQuickWarning)
	xtrap.RegisterHandler(trapTypeDeviceCameraWarning, deviceCameraWarning)
	xtrap.RegisterHandler(trapTypeDeviceLCDWarning, deviceLCDWarning)
	xtrap.RegisterHandler(trapTypeDeviceShellWarning, deviceShellWarning)
	xtrap.RegisterHandler(traptypePowerWarning, powerWarning)
	xtrap.RegisterHandler(trapCalloutWarning, calloutWarning)
	xtrap.RegisterHandler(trapRegisterWarning, registerWarning)
	xtrap.RegisterHandler(trapDrycontactWarning, drycontactWarning)
	//const trapMainBackupSwitchWarning = ".1.3.6.1.4.1.1800.50.1.19"
	xtrap.RegisterHandler(trapMainBackupSwitchWarning, mainBackupSwitchWarning)
	//const trapNtpDropWarning = ".1.3.6.1.4.1.1800.50.1.20"
	xtrap.RegisterHandler(trapNtpDropWarning, netpDropWarning)
	//const trapRecordingFailWarning = ".1.3.6.1.4.1.1800.50.1.21"
	xtrap.RegisterHandler(trapRecordingFailWarning, recordingFaildWarning)
	//const trapMainPowerDropWarning = ".1.3.6.1.4.1.1800.50.1.22"
	xtrap.RegisterHandler(trapMainPowerDropWarning, mainPowerDropWarning)
	//const trapBackupPowerDropWaring = ".1.3.6.1.4.1.1800.50.1.23"
	xtrap.RegisterHandler(trapBackupPowerDropWaring, backupPowerDropWarning)
	//const trapMainNetDropWarning = ".1.3.6.1.4.1.1800.50.1.24"
	xtrap.RegisterHandler(trapMainNetDropWarning, mainNetDropWarning)
	//const trapBackupNetDropWarning = ".1.3.6.1.4.1.1800.50.1.25"
	xtrap.RegisterHandler(trapBackupNetDropWarning, backupNetDropWarning)
	//const trapUsualInput1Warning = ".1.3.6.1.4.1.1800.50.1.26"
	xtrap.RegisterHandler(trapUsualInput1Warning, usualInput1Warning)
	//const trapUsualInput2Warning = ".1.3.6.1.4.1.1800.50.1.27"
	xtrap.RegisterHandler(trapUsualInput2Warning, usualInput2Warning)
	//const trapUsualInput3Warning = ".1.3.6.1.4.1.1800.50.1.28"
	xtrap.RegisterHandler(trapUsualInput3Warning, usualInput3Warning)
	//const trapUsualInput4Warning = ".1.3.6.1.4.1.1800.50.1.29"
	xtrap.RegisterHandler(trapUsualInput4Warning, usualInput4Warning)
	//const trapUsualInput5Warning = ".1.3.6.1.4.1.1800.50.1.30"
	xtrap.RegisterHandler(trapUsualInput5Warning, usualInput5Warning)
	//const trapUsualInput6Warning = ".1.3.6.1.4.1.1800.50.1.31"
	xtrap.RegisterHandler(trapUsualInput6Warning, usualInput6Warning)
	//const trapUsualInput7Warning = ".1.3.6.1.4.1.1800.50.1.32"
	xtrap.RegisterHandler(trapUsualInput7Warning, usualInput7Warning)
	//const trapUsualInput8Warning = ".1.3.6.1.4.1.1800.50.1.33"
	xtrap.RegisterHandler(trapUsualInput8Warning, usualInput8Warning)
	//const trapUsualInput9Warning = ".1.3.6.1.4.1.1800.50.1.34"
	xtrap.RegisterHandler(trapUsualInput9Warning, usualInput9Warning)
	//const trapUsualInput10Warning = ".1.3.6.1.4.1.1800.50.1.35"
	xtrap.RegisterHandler(trapUsualInput10Warning, usualInput10Warning)
}
