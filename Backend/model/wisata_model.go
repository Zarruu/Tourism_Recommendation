package model

type Wisata struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	CategoryName string  `json:"categoryName"`
	City         string  `json:"city"`
	TotalScore   float64 `json:"totalScore"`
	ReviewsCount int     `json:"reviewsCount"`
	Cluster      int     `json:"cluster"`
	ClusterName  string  `json:"clusterName"`
}