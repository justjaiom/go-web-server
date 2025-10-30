
## Features

- Fetches current weather data for any location (Dallas, TX by default)
- Dynamic JSON API endpoint: `/api/weather`
- Frontend updates automatically with live weather info
- Uses free Open-Meteo API (no API key required)






## Usage

### Landing Page
Visit `http://localhost:8080` to see:
- Your personal "About Me" information
- Live weather data for your location
- Responsive design that works on mobile and desktop

### API Endpoint

**GET** `/api/weather`

Returns live weather JSON:
```json
{
  "current_weather": {
    "temperature": 22.5,
    "windspeed": 10.2,
    "time": "2025-10-30T14:30"
  }
}
```

### Example API Call
```bash
curl http://localhost:8080/api/weather
```

---

## Customization

### Change Location

Edit the latitude and longitude in `main.go` in the `weatherHandler` function:
```go
// Dallas, TX (default)
latitude := 32.7767
longitude := -96.7970

// New York, NY
// latitude := 40.7128
// longitude := -74.0060
```

### Update About Me Content

Edit `static/index.html` to personalize:
- Your name and bio
- Profile picture
- Skills and interests
- Social media links
- Styling and colors

### Change Port

Modify the port in `main.go`:
```go
log.Println("Server starting on :8080")
log.Fatal(http.ListenAndServe(":8080", nil))
```

---

## Dependencies

- **Go standard library only** — no external packages needed
- Weather data provided by [Open-Meteo API](https://open-meteo.com/) (free, no API key required)

---

## Development

### Project Structure Breakdown

**main.go:**
- `main()` - Entry point, sets up routes and starts server
- `weatherHandler()` - Fetches weather data from Open-Meteo API
- File server setup for serving static HTML/CSS/JS

**static/index.html:**
- Landing page with personal information
- JavaScript to fetch and display weather data
- CSS styling for responsive design


## API Reference

### Open-Meteo Weather API

The server fetches data from:
```
https://api.open-meteo.com/v1/forecast?latitude={lat}&longitude={lon}&current_weather=true
```

**Parameters:**
- `latitude` - Geographic latitude
- `longitude` - Geographic longitude
- `current_weather=true` - Returns current weather conditions

**Response includes:**
- Temperature (°C)
- Wind speed (km/h)
- Weather code
- Time of measurement

---

## Future Additions

- [ ] Add multiple city support with dropdown selector
- [ ] Implement Go HTML templates for dynamic content
- [ ] Add 7-day weather forecast
- [ ] Include weather icons based on conditions
- [ ] Add unit toggle (Celsius/Fahrenheit)
- [ ] Implement caching to reduce API calls
- [ ] Add dark mode toggle


## Acknowledgments

- Weather data provided by [Open-Meteo](https://open-meteo.com/)
- Built with Go's standard `net/http` package

---

