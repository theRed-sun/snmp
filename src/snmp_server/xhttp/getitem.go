package xhttp

import (
	"net/http"
	"snmp_server/model"
	"snmp_server/xdb"
	"snmp_server/xwarning"

	"github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

type getitemData struct {
	ItemID    int  `json:"itemid"`
	ItemType  int  `json:"itemtype"`
	Recursion bool `json:"recursion"`
}

type getitemRequest struct {
	Token string      `json:"token"`
	Data  getitemData `json:"data"`
}

type itemInfo struct {
	ItemID        int        `json:"itemid"`
	Parent        int        `json:"parent"`
	ItemName      string     `json:"itemname"`
	ItemPath      string     `json:"itempath"`
	ItemType      int        `json:"itemtype"`
	Status        int        `json:"status"`
	DevType       string     `json:"dev_type"`
	Warnings      int        `json:"warnings"`
	ImageURL      string     `json:"imageurl"`
	X             int        `json:"x"`
	Y             int        `json:"y"`
	ServiceStatus int        `json:"service_status"`
	Children      []itemInfo `json:"children,omitempty"`
}

func getChildren(parentID int) []itemInfo {
	items := make([]itemInfo, 0, 100)
	terminals, err := model.GetTerminalByParent(parentID, xdb.Engine)
	if err == nil {
		for _, terminal := range terminals {
			status := 0
			if terminal.IsOnline() {
				status = 1
			}
			warnings := 0
			if status == 1 {
				warnings = xwarning.Stats.GetCounts(terminal.ID)
			}
			item := itemInfo{
				ItemID:        terminal.ID,
				Parent:        terminal.Parent,
				ItemName:      terminal.Name,
				ItemPath:      terminal.Path,
				ItemType:      2,
				Status:        status,
				DevType:       terminal.Type,
				Warnings:      warnings,
				X:             terminal.X,
				Y:             terminal.Y,
				ServiceStatus: terminal.ServiceStatus,
			}
			items = append(items, item)
		}
	}

	zones, err := model.GetZonesByParent(parentID, xdb.Engine)
	if err == nil {
		for _, zone := range zones {
			item := itemInfo{
				ItemID:   zone.ID,
				Parent:   zone.Parent,
				ItemName: zone.Name,
				ItemPath: zone.Path,
				ItemType: 1,
				ImageURL: zone.ImageURL,
				X:        zone.X,
				Y:        zone.Y,
			}
			children := getChildren(zone.ID)
			if len(children) > 0 {
				item.Children = children
			}
			items = append(items, item)
		}
	}
	return items
}

func getitem(c *gin.Context) {
	seelog.Info("getitem.")
	var request getitemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}
	switch request.Data.ItemType {
	case 1:
		self, err := model.GetZoneByID(request.Data.ItemID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			return
		}
		selfInfo := itemInfo{
			ItemID:   self.ID,
			Parent:   self.Parent,
			ItemName: self.Name,
			ItemPath: self.Path,
			ItemType: 1,
			ImageURL: self.ImageURL,
			X:        self.X,
			Y:        self.Y,
		}
		if request.Data.Recursion {
			items := getChildren(request.Data.ItemID)
			result := gin.H{
				"result":  0,
				"message": "OK",
				"data": gin.H{
					"self":  selfInfo,
					"items": items,
				},
			}
			c.JSON(http.StatusOK, result)
			return
		}
		var items []itemInfo
		zones, err := model.GetZonesByParent(request.Data.ItemID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			break
		}
		for _, zone := range zones {
			item := itemInfo{
				ItemID:   zone.ID,
				Parent:   zone.Parent,
				ItemName: zone.Name,
				ItemPath: zone.Path,
				ItemType: 1,
				ImageURL: zone.ImageURL,
				X:        zone.X,
				Y:        zone.Y,
			}
			items = append(items, item)
		}

		terminals, _ := model.GetTerminalByParent(request.Data.ItemID, xdb.Engine)
		for _, terminal := range terminals {
			status := 0
			if terminal.IsOnline() {
				status = 1
			}
			warnings := 0
			if status == 1 {
				warnings = xwarning.Stats.GetCounts(terminal.ID)
			}
			item := itemInfo{
				ItemID:        terminal.ID,
				Parent:        terminal.Parent,
				ItemName:      terminal.Name,
				ItemPath:      terminal.Path,
				ItemType:      2,
				Status:        status,
				DevType:       terminal.Type,
				Warnings:      warnings,
				X:             terminal.X,
				Y:             terminal.Y,
				ServiceStatus: terminal.ServiceStatus,
			}
			items = append(items, item)
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"self":  selfInfo,
				"items": items,
			},
		}
		c.JSON(http.StatusOK, result)
	case 2:
		terminal, err := model.GetTerminalByID(request.Data.ItemID, xdb.Engine)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
			break
		}
		var items []itemInfo
		status := 0
		if terminal != nil {
			if terminal.IsOnline() {
				status = 1
			}
			warnings := 0
			if status == 1 {
				warnings = xwarning.Stats.GetCounts(terminal.ID)
			}
			item := itemInfo{
				ItemID:        terminal.ID,
				Parent:        terminal.Parent,
				ItemName:      terminal.Name,
				ItemPath:      terminal.Path,
				ItemType:      2,
				Status:        status,
				DevType:       terminal.Type,
				Warnings:      warnings,
				X:             terminal.X,
				Y:             terminal.Y,
				ServiceStatus: terminal.ServiceStatus,
			}
			items = append(items, item)
		}
		result := gin.H{
			"result":  0,
			"message": "OK",
			"data": gin.H{
				"items": items,
			},
		}
		c.JSON(http.StatusOK, result)

	case 3:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support user itemtype"})
	default:
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": "not support itemtype"})
	}
}

type getAllTerminalsRequest struct {
	Token string `json:"token"`
}

func getAllTerminals(c *gin.Context) {
	var request getAllTerminalsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	if !authToken(request.Token, c) {
		return
	}

	terminals, err := model.GetAllTerminals(xdb.Engine)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "message": err.Error()})
		return
	}
	tsinfos := make(map[string]string)
	for _, t := range terminals {
		tsinfos[t.Path+"."+t.Name] = t.NTID
	}

	result := gin.H{
		"result":  0,
		"message": "OK",
		"data":    tsinfos,
	}
	c.JSON(http.StatusOK, result)

}
