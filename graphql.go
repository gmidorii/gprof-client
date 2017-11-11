package main

type Prof struct {
	Data Data `json:"data"`
}

type Data struct {
	CPU  CPU  `json:"cpu"`
	Disk Disk `json:"disk"`
	File File `json:"file"`
}

type CPU struct {
	CacheSize int `json:"cache_size"`
	Cores     []struct {
		Percent float64 `json:"percent"`
	} `json:"cores"`
	Model     string `json:"model"`
	ModelName string `json:"model_name"`
}

type Disk struct {
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
}

type File struct {
	Content     string `json:"content"`
	Name        string `json:"name"`
	UpdatedTime string `json:"updated_time"`
}
