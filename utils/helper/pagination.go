package helper

import "math"

func Paginate(limit, offset, totalData int) map[string]interface{} {

	if offset < totalData {
		endIndex := offset + limit
		if endIndex > totalData {
			endIndex = totalData
		}
	}

	totalPage := 1
	if totalData > 0 {
		totalPage = int(math.Ceil(float64(totalData) / float64(limit)))
	}
	currentPage := int(math.Ceil(float64(offset+1) / float64(limit)))

	if currentPage > totalPage {
		currentPage = totalPage
	}

	pagination := map[string]interface{}{
		"current_limit":  limit,
		"current_offset": offset,
		"current_page":   currentPage,
		"total_data":     totalData,
		"total_page":     totalPage,
	}

	return pagination
}
