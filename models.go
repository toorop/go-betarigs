package betarigs

// algo represents pair id->algo name as returned by GET algorithms
type AlgoId struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

// algorithm represents a GET algorithm JSON response
type Algorithm struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	MarketPrice struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"market_price"`
	RentedCapacity struct {
		Value uint64 `json:"value"`
		Unit  string `json:"unit"`
	} `json:"rented_capacity"`
	AvailableCapacity struct {
		Value uint64 `json:"value"`
		Unit  string `json:"unit"`
	} `json:"available_capacity"`
}

// rig represents a rig (Amazing !)
type Rig struct {
	Id            uint32 `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	DeclaredSpeed struct {
		Value uint32 `json:"value"`
		Unit  string `json:"unit"`
	} `json:"declared_speed"`
	Algorithm struct {
		Id   uint32 `json:"id"`
		Name string `json:"name"`
	} `json:"algorithm"`
	Price struct {
		Total struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"total"`
		PerSpeedUnit struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"per_speed_unit"`
	} `json:"price"`
	Stats struct {
		Total struct {
			Score             uint32 `json:"score"`
			NumberOfRentals   uint32 `json:"number_of_rentals"`
			NumberOfFeedbacks uint32 `json:"number_of_feedbacks"`
		} `json:"total"`
		Month struct {
			Score             uint32 `json:"score"`
			NumberOfRentals   uint32 `json:"number_of_rentals"`
			NumberOfFeedbacks uint32 `json:"number_of_feedbacks"`
		} `json:"month"`
	} `json:"stats"`
	Status struct {
		Available bool   `json:"available"`
		Label     string `json:"label"`
	} `json:"status"`
}
