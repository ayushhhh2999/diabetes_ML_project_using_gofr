# Diabetes Prediction API (Go (gofr) + Python)

This project is a RESTful API for predicting diabetes based on patient data. The backend is written in Go and uses the [gofr](https://github.com/gofr-dev/gofr) framework. It calls a Python script to run a machine learning model (e.g., scikit-learn) for the actual prediction.

---

## Features

- **REST API**: Accepts patient data as JSON and returns a diabetes prediction.
- **Go Backend**: Handles HTTP requests, validation, and process management.
- **Python ML Model**: Loads a trained model and scaler to make predictions.
- **Easy Integration**: Simple interface for frontend or other services.

---

## Project Structure


. ├── main.go # Go server code (API) ├── predict.py # Python script for prediction ├── model.pkl # Trained ML model (pickle) ├── scaler.pkl # Scaler for input normalization (pickle) ├── requirements.txt # Python dependencies ├── go.mod # Go module file └── README.md # Project documentation


---

## Requirements

- **Go** (v1.18+ recommended)
- **Python** (3.8+ recommended)
- **pip** (for Python dependencies)
- Trained model and scaler (`model.pkl`, `scaler.pkl`)

---

## Setup

### 1. Clone the Repository

```sh
git clone https://github.com/ayushhhh2999/diabetes_ML_project_using_gofr
cd diabetes-prediction-go-python

python -m venv venv
.\venv\Scripts\activate   # On Windows

pip install -r requirements.txt

go mod download
go run [main.go](http://localhost:8000)
The server will start (default: localhost:8000).

API Endpoints
Health Check
GET /
Returns: "Hello, World!"
Diabetes Prediction
POST /predict
Content-Type: application/json
Body Example:
{
  "pregnancies": 2,
  "glucose": 120,
  "bloodPressure": 70,
  "skinThickness": 20,
  "insulin": 79,
  "bmi": 25.6,
  "diabetespedigreeFunction": 0.351,
  "Age": 29
}
Response Example:
{
  "prediction": "The prediction result is: Positive"
}
```
How It Works
Client sends a POST request with patient data.

Go server parses the request and calls the Python script (predict.py) with the data as command-line arguments.

Python script loads the model and scaler, processes the input, and prints the prediction.

Go server captures the output and returns it as a JSON response.


Notes

Ensure model.pkl and scaler.pkl are present in the project directory.

The Go server expects Python to be available at .\\venv\\Scripts\\python.exe (Windows). Adjust the path if needed.

For deployment, consider using Docker for consistent environments.


Author

Ayush
