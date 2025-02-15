package controller

import (
	"api_interes_compuesto/helpers"
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

type dataYear struct {
	Year    uint16  `json:"year"`
	Capital float64 `json:"capital"`
}

type dataMonth struct {
	Year           uint16  `json:"year"`
	Month          uint16  `json:"month"`
	MontfOfTheYear uint16  `json:"montfOfTheYear"`
	MonthName      string  `json:"monthName"`
	Capital        float64 `json:"capital"`
}

type body struct {
	Success   bool   `json:"success"`
	CodeError uint16 `json:"code"`
	Message   string `json:"message"`
}

type messageJsonYear struct {
	body `json:""`
	Data []dataYear `json:"data"`
}

type messageJsonMonth struct {
	body `json:""`
	Data []dataMonth `json:"data"`
}

func CalculateByYear(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Save query values
	anios, _ := strconv.ParseInt(query.Get("anios"), 10, 8)
	capitalInicial, _ := strconv.ParseFloat(query.Get("capitalInicial"), 64)
	interes, _ := strconv.ParseFloat(query.Get("interes"), 32)

	var response messageJsonYear

	response.body.Success = true
	response.body.CodeError = http.StatusOK
	response.body.Message = http.StatusText(http.StatusOK)

	for year := 1; year <= int(anios); year++ {
		response.Data = append(response.Data, dataYear{
			Year:    uint16(year),
			Capital: capitalInicial * math.Pow(1+interes/100, float64(year)),
		})
	}

	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error al obtener respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CalculateByMonth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Save query values
	anios, _ := strconv.ParseInt(query.Get("anios"), 10, 8)
	capitalInicial, _ := strconv.ParseFloat(query.Get("capitalInicial"), 64)
	interes, _ := strconv.ParseFloat(query.Get("interes"), 32)

	var response messageJsonMonth

	response.body.Success = true
	response.body.CodeError = http.StatusOK
	response.body.Message = http.StatusText(http.StatusOK)

	counterMonth := 1
	for year := 1; year <= int(anios); year++ {
		for month := 1; month <= 12; month++ {
			response.Data = append(response.Data, dataMonth{
				Year:           uint16(year),
				Month:          uint16(counterMonth),
				MontfOfTheYear: uint16(month),
				MonthName:      helpers.GetMonthName(month - 1),
				Capital:        capitalInicial * math.Pow(1+interes/100, float64(year*month)/(12*float64(year))),
			})
			counterMonth++
		}
		capitalInicial = capitalInicial * math.Pow(1+interes/100, float64(year*12)/(12*float64(year)))
	}

	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error al obtener respuesta", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
