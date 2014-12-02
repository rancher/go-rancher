package cli

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func PrintTabular(obj interface{}) {
	listObj := obj.(map[string]interface{})
	listMapObj := make([]interface{}, 0)
	parseMap := false
	if _, ok := listObj["data"]; ok {
		switch listObj["data"].(type) {
		case []interface{}:
			listMapObj = listObj["data"].([]interface{})
		case map[string]interface{}:
			parseMap = true
		}
	}
	if parseMap {
		listMapObj = append(listMapObj, listObj)
	}
	for _, dataUnit := range listMapObj {
		tableData := make([][]string, 1)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Type", "Id", "Name", "Value"})
		mapUnit := dataUnit.(map[string]interface{})
		typeVal := mapUnit["type"].(string)
		idVal := mapUnit["id"].(string)
		for k, v := range mapUnit {
			if k == "type" || k == "id" {
				continue
			}
			keyVal := k
			var valVal string
			switch v.(type) {
			case int:
				valVal = strconv.Itoa(v.(int))
			case float64:
				valVal = strconv.FormatFloat(v.(float64), 'G', -1, 64)
			case string:
				valVal = v.(string)
			case bool:
				valVal = strconv.FormatBool(v.(bool))
			default:
				continue
			}
			if len(valVal) > 70 {
				valVal = valVal[:70]
			}
			tableData = append(tableData, []string{typeVal, idVal, keyVal, valVal})
		}
		table.AppendBulk(tableData)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.Render()
	}

}
