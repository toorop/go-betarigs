package betarigs

import (
	"strconv"
	"time"
)

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
	RentalDurations []struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"rental_durations"`
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

// pool represents a pool for rental
type Pool struct {
	Url            string `json:"url"`
	WorkerName     string `json:"worker_name"`
	WorkerPassword string `json:"worker_password"`
}

// Jtime 2014-07-06T17:39:16+0000"
const timeFormat = "2006-01-02T15:04:05+0000"

type jTime time.Time

func (t jTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(time.Time(t).Format(timeFormat))), nil
}

func (t *jTime) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	*(*time.Time)(t), err = time.Parse(timeFormat, q)
	return
}

func (t jTime) String() string { return time.Time(t).String() }

// Responses

// rentalJsonResponse represent a response to rental call
type RentalResponse struct {
	Id        int    `json:"id"`
	Status    string `json:"status"`
	CreatedAt jTime  `json:"created_at"`
	Duration  struct {
		InitialDuration struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"initial_duration"`
		CurrentDuration struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"initial_duration"`
	} `json:"current_duration"`
	//MiningPeriod []string `json:"mining_period"`
	Payment struct {
		Bitcoin struct {
			Price struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"price"`
			PaymentAddress string `json:"payment_address"`
		} `json:"bitcoin"`
	} `json:"payment"`
	Rig struct {
		Id        int `json:"id"`
		Algorithm struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"algorithm"`
		//Feedback []string `json:"feedback"`
		Pool Pool `json:"pool"`
	} `json:"rig"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
