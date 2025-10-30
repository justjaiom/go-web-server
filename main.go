package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type WeatherResponse struct {
    CurrentWeather struct {
        Temperature float64 `json:"temperature"`
        Windspeed   float64 `json:"windspeed"`
    } `json:"current_weather"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    // Dallas, TX coordinates
    latitude := 32.7767
    longitude := -96.7970

    apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", latitude, longitude)

    resp, err := http.Get(apiURL)
    if err != nil {
        http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var weather WeatherResponse
    json.Unmarshal(body, &weather)

    // Return JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(weather)
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/api/weather", weatherHandler)

    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
