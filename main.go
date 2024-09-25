package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// Define the Airport struct
type Airport struct {
    Name     string `json:"name"`
    City     string `json:"city"`
    IATA     string `json:"iata"`
    ImageURL string `json:"image_url"`
}

// Define the AirportV2 struct with additional runway length information
type AirportV2 struct {
    Airport
    RunwayLength int `json:"runway_length"`
}

// Mock data for airports in Bangladesh
var airports = []Airport{
    {"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"},
    {"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"},
    {"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"},
}

// Mock data for airports in Bangladesh (with runway length for V2)
var airportsV2 = []AirportV2{
    {Airport{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"}, 3200},
    {Airport{"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"}, 2900},
    {Airport{"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"}, 2500},
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Status: OK"))
}

// Airports handler for the first endpoint
func Airports(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(airports)
}

// AirportsV2 handler for the second version endpoint
func AirportsV2(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(airportsV2)
}

// UpdateAirportImage handler for updating airport images
func UpdateAirportImage(w http.ResponseWriter, r *http.Request) {
    // Parse the request to get the airport name
    var request struct {
        Name string `json:"name"`
    }
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Find the airport by name
    var airport *Airport
    for i := range airports {
        if airports[i].Name == request.Name {
            airport = &airports[i]
            break
        }
    }

    if airport == nil {
        http.Error(w, "Airport not found", http.StatusNotFound)
        return
    }

    // Read the image from the request
    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Failed to read image", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Initialize AWS S3 client
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"), // Set your desired AWS region
    })
    if err != nil {
        http.Error(w, "Failed to create AWS session", http.StatusInternalServerError)
        return
    }
    svc := s3.New(sess)

    // Define the S3 bucket and the image key (file name)
    bucket := "bangladesh-airport-images" // Your S3 bucket name
    key := fmt.Sprintf("airport_images/%s.jpg", airport.IATA)

    // Upload image to S3
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   file,
        ACL:    aws.String("public-read"),
    })
    if err != nil {
        http.Error(w, "Failed to upload image to S3", http.StatusInternalServerError)
        return
    }

    // Update the airport's image URL
    airport.ImageURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)

    // Respond with the updated airport details
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(airport)
}

func main() {
    // Setup routes
    http.HandleFunc("/", HomePage)
    http.HandleFunc("/airports", Airports)
    http.HandleFunc("/airports_v2", AirportsV2)
    http.HandleFunc("/update_airport_image", UpdateAirportImage)

    // Start the server
    fmt.Println("Server starting on port :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}
