package main

import (
	"fmt"
	"os/exec"
	"strings"

	"gofr.dev/pkg/gofr"
)

// Patient matches the JSON structure from the Pydantic model (Go version)
type Patient struct {
	Pregnancies              int     `json:"pregnancies"`
	Glucose                  int     `json:"glucose"`
	BloodPressure            int     `json:"bloodPressure"`
	SkinThickness            int     `json:"skinThickness"`
	Insulin                  int     `json:"insulin"`
	BMI                      float64 `json:"bmi"`
	DiabetesPedigreeFunction float64 `json:"diabetespedigreeFunction"`
	Age                      int     `json:"Age"`
}

func main() {
	app := gofr.New()

	app.GET("/", func(ctx *gofr.Context) (any, error) {
		return "Hello, World!", nil
	})

	app.POST("/predict", func(ctx *gofr.Context) (any, error) {
		var patient Patient

		// Bind incoming JSON to Go struct
		if err := ctx.Bind(&patient); err != nil {
			return nil, fmt.Errorf("invalid request body: %v", err)
		}

		// Prepare arguments for Python script
		argv := []string{
			"predict.py",
			fmt.Sprintf("%d", patient.Pregnancies),
			fmt.Sprintf("%d", patient.Glucose),
			fmt.Sprintf("%d", patient.BloodPressure),
			fmt.Sprintf("%d", patient.SkinThickness),
			fmt.Sprintf("%d", patient.Insulin),
			fmt.Sprintf("%f", patient.BMI),
			fmt.Sprintf("%f", patient.DiabetesPedigreeFunction),
			fmt.Sprintf("%d", patient.Age),
		}

		// Execute Python script
		cmd := exec.Command(".\\venv\\Scripts\\python.exe", argv...)
		out, err := cmd.CombinedOutput()

		if err != nil {
			return nil, fmt.Errorf("error running Python script: %v, output: %s", err, string(out))
		}

		// Return JSON response
		return map[string]string{
			"prediction": fmt.Sprintf("The prediction result is: %s", strings.TrimSpace(string(out))),
		}, nil
	})

	fmt.Println("Server started successfully")
	app.Run()
}
