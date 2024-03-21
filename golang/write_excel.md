
```
import "github.com/xuri/excelize/v2" 

func writeSheet(f *excelize.File, sheetName string, resourceList []*api.Resource) error {
	glog.Infof("writeSheet %v", sheetName)
	// Create a new sheet.
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	// set header
	for idx, header := range SheetHeaderList {
		colume := fmt.Sprintf("%c%v", 'A'+idx, 1)
		f.SetCellValue(sheetName, colume, header)
	}

	// row start
	row := 2
	for _, resource := range resourceList {
		for resourceName, resourceValue := range resource.Quota {
			f.SetCellValue(sheetName, fmt.Sprintf("A%v", row), resource.NodeName)
			f.SetCellValue(sheetName, fmt.Sprintf("B%v", row), resource.Region)
			f.SetCellValue(sheetName, fmt.Sprintf("C%v", row), resource.AZ)
			f.SetCellValue(sheetName, fmt.Sprintf("D%v", row), resource.Group)
			f.SetCellValue(sheetName, fmt.Sprintf("E%v", row), resource.ResourcePool)
			f.SetCellValue(sheetName, fmt.Sprintf("F%v", row), resource.Env)
			f.SetCellValue(sheetName, fmt.Sprintf("G%v", row), resourceName)
			f.SetCellValue(sheetName, fmt.Sprintf("H%v", row), resourceValue)

			resourceUsed := resource.Used[resourceName]
			resourceMargin := resource.Margin[resourceName]
			f.SetCellValue(sheetName, fmt.Sprintf("I%v", row), resourceUsed)
			f.SetCellValue(sheetName, fmt.Sprintf("J%v", row), resourceMargin)
			var rate float64
			if resourceValue != 0 {
				rate = resourceUsed / resourceValue
			}
			f.SetCellValue(sheetName, fmt.Sprintf("K%v", row), fmt.Sprintf("%.2f%%", rate*100))
			row++
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")
    return nil
}
```
