package entity

import (

)


type Response struct {
	IsAllowed bool `json:"isallowed"`
	Error DataError `json:"error"`
}


