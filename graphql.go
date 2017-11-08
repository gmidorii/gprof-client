package main

type Prof struct {
	Data struct {
		CPU struct {
			CacheSize int `json:"cache_size"`
			Cores     []struct {
				Percent float64 `json:"percent"`
			} `json:"cores"`
			Model     string `json:"model"`
			ModelName string `json:"model_name"`
		} `json:"cpu"`
		Disk struct {
			Io struct {
				ReadCount int `json:"read_count"`
			} `json:"io"`
			Usage struct {
				Free        int     `json:"free"`
				Path        string  `json:"path"`
				Total       int     `json:"total"`
				Used        int     `json:"used"`
				UsedPercent float64 `json:"used_percent"`
			} `json:"usage"`
		} `json:"disk"`
		File struct {
			Content     string `json:"content"`
			Name        string `json:"name"`
			UpdatedTime string `json:"updated_time"`
		} `json:"file"`
	} `json:"data"`
}
