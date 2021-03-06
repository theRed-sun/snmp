package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snmp_server/mibs"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xsnmp"
	"snmp_server/xtask"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type snmpBatchData struct {
	DeviceType string                 `json:"dev_type"`
	ItemPath   string                 `json:"itempath"`
	ItemIDs    []int                  `json:"itemids"`
	OIDs       map[string]interface{} `json:"oids"`
}

type snmpBatchRequest struct {
	Token string        `json:"token"`
	Data  snmpBatchData `json:"data"`
}

type tItemInfo struct {
	ItemID   int    `json:"itemid"`
	ItemName string `json:"itemname"`
	Error    string `json:"error,omitempty"`
}

//supportBatch map value is default value
var supportBatch = map[string]interface{}{
	//for KN518
	"k518_language_mode":     0,
	"k518_http_mode":         0,
	"k518_http_port":         80,
	"k518_telnet_port":       23,
	"k518_sip_local_port":    5060,
	"k518_dtmf_relay_mode":   0,
	"k518_codec_type_list":   "3 1 0 4 2",
	"k518_hotline_enable":    0,
	"k518_hotline_number":    "690",
	"k518_hotline_account":   0,
	"k518_microphone_volume": 6,
	"k518_speaker_volume":    4,
	"k518_hookon_wait_time":  10,
	"k518_ring_style":        0,
	"k518_ring_volume":       7,
	//for KN519
	"k519_language_mode":     0,
	"k519_http_mode":         0,
	"k519_http_port":         80,
	"k519_telnet_port":       23,
	"k519_sip_local_port":    5060,
	"k519_dtmf_relay_mode":   0,
	"k519_codec_type_list":   "3 1 0 4 2",
	"k519_hotline_enable":    0,
	"k519_hotline_number":    "690",
	"k519_hotline_account":   0,
	"k519_video_mode":        0,
	"k519_video_paytype":     107,
	"k519_microphone_volume": 5,
	"k519_speaker_volume":    4,
	"k519_hookon_wait_time":  10,
	"k519_ring_style":        0,
	"k519_ring_volume":       7,
	//for KN618
	"k618_language_mode":     0,
	"k618_http_mode":         0,
	"k618_http_port":         80,
	"k618_sip_local_port":    5060,
	"k618_codec_type_list":   "3 1 0 4 2",
	"k618_hotline_enable":    0,
	"k618_hotline_number":    "690",
	"k618_hotline_account":   0,
	"k618_video_mode":        0,
	"k618_microphone_volume": 5,
	"k618_speaker_volume":    4,
	"k618_hookon_wait_time":  10,
	"k618_ring_style":        0,
	"k618_ring_volume":       7,
	//for all device
	"usl_reboot_device":          "reboot",
	"usl_set_default":            "default",
	"usl_ftp_server_ip":          "",
	"usl_ftp_server_port":        21,
	"usl_ftp_user_name":          "",
	"usl_ftp_user_passwd":        "",
	"usl_ftp_file_size":          "",
	"usl_ftp_soft_file_name":     "",
	"usl_ftp_save_cfg_file_name": "",
}

func snmpBatch(c *gin.Context) {
	var request snmpBatchRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	jdata, _ := json.MarshalIndent(&request, "", " ")
	seelog.Info("snmp_batch request:\n", string(jdata))
	if !authToken(request.Token, c) {
		return
	}
	var terminals []*model.Terminal
	var err error
	if len(request.Data.ItemIDs) > 0 {
		for _, id := range request.Data.ItemIDs {
			t, err := model.GetTerminalByID(id, xdb.Engine)
			if err == nil && t != nil {
				terminals = append(terminals, t)
				seelog.Infof("id %d terminal %v", id, t)
			}
		}
		seelog.Info("terminals:", terminals)
	} else {
		terminals, err = model.GetTerminalByPathAndType(request.Data.ItemPath, request.Data.DeviceType, xdb.Engine)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if len(request.Data.OIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "oids is empty", "data": gin.H{"suuports": supportBatch}})
		return
	}
	for key := range request.Data.OIDs {
		_, ok := supportBatch[key]
		if !ok {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": fmt.Sprintf("oid [%s] not support", key), "data": gin.H{"suuports": supportBatch}})
			return
		}
	}

	err = mibs.CheckSetPDUs(request.Data.OIDs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}

	rspOkItems := make([]*tItemInfo, 0, len(terminals))
	rspErrorItems := make([]*tItemInfo, 0, len(terminals))
	var waitGroup sync.WaitGroup
	for _, t := range terminals {
		//rspOkItems = append(rspOkItems, &tItemInfo{ItemID: t.ID, ItemName: t.Name})
		waitGroup.Add(1)
		go func(t *model.Terminal) {
			var ftpUpgrade *xtask.Upgrade
			var err error
			if _, ok := request.Data.OIDs["usl_ftp_soft_file_name"]; ok {
				ftpUpgrade, err = initUpgrade(t, c)
				if err != nil {
					rspErrorItems = append(rspErrorItems, &tItemInfo{ItemID: t.ID, ItemName: t.Name, Error: err.Error()})
					waitGroup.Done()
					return
				}
			}
			oids := copyOIDs(request.Data.OIDs)
			if _, ok := oids["usl_ftp_save_cfg_file_name"]; ok {
				oids["usl_ftp_save_cfg_file_name"] = fmt.Sprintf("%s/%s.cfg", t.NTID, time.Now().Format("20060102_150405"))
			}
			snmpResult, err := xsnmp.Default.Set(oids, 0, t.Remote())

			if err == nil {
				jdata, _ := json.Marshal(snmpResult)
				logInfo := &model.LogInfo{
					User:     getUsernameByToken(request.Token),
					NTID:     t.NTID,
					Event:    "snmp",
					SubEvent: "set_ok",
					Info:     string(jdata),
				}
				logInfo.Insert()
				rspOkItems = append(rspOkItems, &tItemInfo{ItemID: t.ID, ItemName: t.Name})
			} else {
				logInfo := &model.LogInfo{
					User:     getUsernameByToken(request.Token),
					NTID:     t.NTID,
					Event:    "snmp",
					SubEvent: "set_err",
					Info:     err.Error(),
				}
				logInfo.Insert()
				rspErrorItems = append(rspErrorItems, &tItemInfo{ItemID: t.ID, ItemName: t.Name, Error: err.Error()})
				if ftpUpgrade != nil {
					seelog.Warnf("ftp upgrade task Delete  %s for snmp set error", ftpUpgrade.NTID)
					ftpUpgrade.Delete(xdb.EngineTask)
				}
			}
			waitGroup.Done()
		}(t)
	}
	waitGroup.Wait()
	result := gin.H{
		"result":  0,
		"message": "OK",
		"data": gin.H{
			"ok_items":    rspOkItems,
			"error_items": rspErrorItems,
		},
	}
	c.JSON(http.StatusOK, result)
}

func copyOIDs(oids map[string]interface{}) map[string]interface{} {
	newoids := make(map[string]interface{})
	for key, value := range oids {
		newoids[key] = value
	}
	return newoids
}
