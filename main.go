// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/rand"
// 	"net/http"

// 	//"net/http"
// 	"time"

// 	"github.com/go-resty/resty/v2"
// )

// // ExponentialBackoffWithJitter calculates the backoff delay for each retry attempt, with jitter.
// func ExponentialBackoffWithJitter(attempt int) time.Duration {
// 	minWait := 1 * time.Second
// 	maxWait := 10 * time.Second

// 	// Exponential backoff with full jitter
// 	wait := time.Duration(math.Pow(2, float64(attempt))) * minWait
// 	jitter := time.Duration(rand.Int63n(int64(wait)))

// 	if wait+jitter > maxWait {
// 		return maxWait
// 	}

// 	return wait + jitter
// }

// // CallAPI makes an HTTP request with the specified settings and returns the response body.
// func CallAPI(url string, method string, headers map[string]string, params map[string]interface{}, cert tls.Certificate) (map[string]interface{}, error) {
// 	// Configure TLS settings
// 	tlsConfig := &tls.Config{
// 		MinVersion:         tls.VersionTLS12,
// 		InsecureSkipVerify: false, // Do not skip certificate verification
// 		Renegotiation:      tls.RenegotiateOnceAsClient,
// 		Certificates:       []tls.Certificate{cert}, // Client certificate for mutual TLS
// 	}

// 	// Configure HTTP transport settings
// 	tr := &http.Transport{
// 		TLSClientConfig:   tlsConfig,
// 		DisableKeepAlives: true, // Disable keep-alives to force new connections for each request
// 	}

// 	// Create a new Resty client with a custom transport
// 	client := resty.New().
// 		SetTransport(tr).
// 		SetTimeout(10 * time.Second). // Set request timeout
// 		SetRetryCount(5).             // Set retry count
// 		AddRetryCondition(func(r *resty.Response, err error) bool {
// 			// Retry on network errors or 5xx HTTP status codes
// 			return r.StatusCode() >= 500 || err != nil
// 		})

// 	// Add a custom backoff function
// 	client.SetRetryWaitTimeFunc(func(attempt int) time.Duration {
// 		return ExponentialBackoffWithJitter(attempt)
// 	})

// 	// Prepare the request
// 	request := client.R().
// 		SetHeaders(headers)

// 	// Set request parameters based on the HTTP method
// 	switch method {
// 	case "GET":
// 		stringParams := ConvertMapToStringMap(params)
// 		request.SetQueryParams(stringParams)
// 	case "POST", "PUT", "DELETE":
// 		request.SetBody(params)
// 	}

// 	// Execute the request and handle the response
// 	var responseBody map[string]interface{}
// 	var response *resty.Response
// 	var err error
// 	response, err := request.SetResult(&responseBody).Execute(method, url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return responseBody, nil
// }

// // ConvertMapToStringMap is a helper function to convert map[string]interface{} to map[string]string
// func ConvertMapToStringMap(params map[string]interface{}) map[string]string {
// 	stringParams := make(map[string]string)
// 	for key, value := range params {
// 		stringParams[key] = fmt.Sprintf("%v", value)
// 	}
// 	return stringParams
// }

// func main() {
// 	// Example usage of the CallAPI function
// 	url := "http://localhost:8081/v1/users/"
// 	method := "GET"
// 	headers := map[string]string{
// 		"Content-Type": "application/json",
// 	}
// 	params := map[string]interface{}{
// 		"param1": "value1",
// 		"param2": 123,
// 	}

// 	// Load client certificate
// 	cert, err := tls.LoadX509KeyPair("path/to/cert.pem", "path/to/key.pem")
// 	if err != nil {
// 		log.Fatalf("failed to load client certificate: %v", err)
// 	}

// 	response, err := CallAPI(url, method, headers, params, cert)
// 	if err != nil {
// 		log.Fatalf("API call failed: %v", err)
// 	}

// 	fmt.Printf("API response: %+v\n", response)
// }
/////////////////////////////////////////////////////////////////////////////////

// package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/go-resty/resty/v2"
// )

// // ExponentialBackoffWithJitter calculates the backoff delay for each retry attempt, with jitter.
// func ExponentialBackoffWithJitter(attempt int) time.Duration {
// 	minWait := 1 * time.Second
// 	maxWait := 10 * time.Second

// 	// Exponential backoff with full jitter
// 	wait := time.Duration(math.Pow(2, float64(attempt))) * minWait
// 	jitter := time.Duration(rand.Int63n(int64(wait)))

// 	if wait+jitter > maxWait {
// 		return maxWait
// 	}

// 	return wait + jitter
// }

// // CallAPI makes an HTTP request with the specified settings and returns the response body.
// func CallAPI(url string, method string, headers map[string]string, params map[string]interface{}, cert tls.Certificate) (map[string]interface{}, error) {
// 	// Configure TLS settings
// 	tlsConfig := &tls.Config{
// 		MinVersion:         tls.VersionTLS12,
// 		InsecureSkipVerify: false, // Do not skip certificate verification
// 		Renegotiation:      tls.RenegotiateOnceAsClient,
// 		Certificates:       []tls.Certificate{cert}, // Client certificate for mutual TLS
// 	}

// 	// Configure HTTP transport settings
// 	tr := &http.Transport{
// 		TLSClientConfig:   tlsConfig,
// 		DisableKeepAlives: true, // Disable keep-alives to force new connections for each request
// 	}

// 	// Create a new Resty client with a custom transport
// 	client := resty.New().
// 		SetTransport(tr).
// 		SetTimeout(10 * time.Second).          // Set request timeout
// 		SetRetryCount(5).                      // Set retry count
// 		SetRetryWaitTime(1 * time.Second).     // Set the minimum wait time between retries
// 		SetRetryMaxWaitTime(10 * time.Second). // Set the maximum wait time between retries
// 		AddRetryCondition(func(r *resty.Response, err error) bool {
// 			// Retry on network errors or 5xx HTTP status codes
// 			return r.StatusCode() >= 500 || err != nil
// 		})
// 		// SetRetryWaitTimeFunc(func(attempt int, response *resty.Response) time.Duration {
// 		// 	return ExponentialBackoffWithJitter(attempt)
// 		// })

// 	// Prepare the request
// 	request := client.R().
// 		SetHeaders(headers)

// 	// Set request parameters based on the HTTP method
// 	switch method {
// 	case "GET":
// 		stringParams := ConvertMapToStringMap(params)
// 		request.SetQueryParams(stringParams)
// 	case "POST", "PUT", "DELETE":
// 		request.SetBody(params)
// 	}

// 	// Define a variable to store the response body
// 	var responseBody map[string]interface{}

// 	// // Execute the request and handle the response
// 	// response, err := request.SetResult(&responseBody).Execute(method, url)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return responseBody, nil
// 	// Execute the request and handle the response with manual retry logic
// 	for attempt := 0; attempt <= client.RetryCount; attempt++ {
// 		response, err := request.SetResult(&responseBody).Execute(method, url)
// 		if err == nil && response.StatusCode() < 500 {
// 			return responseBody, nil
// 		}

// 		time.Sleep(ExponentialBackoffWithJitter(attempt))
// 	}

// 	return nil, fmt.Errorf("failed after %d attempts", client.RetryCount)
// }

// // ConvertMapToStringMap is a helper function to convert map[string]interface{} to map[string]string
// func ConvertMapToStringMap(params map[string]interface{}) map[string]string {
// 	stringParams := make(map[string]string)
// 	for key, value := range params {
// 		stringParams[key] = fmt.Sprintf("%v", value)
// 	}
// 	return stringParams
// }

// func main() {
// 	// Example usage of the CallAPI function
// 	url := "https://example.com/api"
// 	method := "POST" //"GET"
// 	headers := map[string]string{
// 		"Content-Type": "application/json",
// 	}
// 	// params := map[string]interface{}{
// 	// 	"param1": "value1",
// 	// 	"param2": 123,
// 	// }
// 	params := map[string]interface{}{
// 		"email":        "3246g11test288911@gmail.com",
// 		"password":     "fghjklhjgf",
// 		"name":         "sawerr",
// 		"check":        10,
// 		"created_time": time.Now().Format("15:04"),
// 	}

// 	// Load client certificate
// 	//cert, err := tls.LoadX509KeyPair("path/to/certs.pem", "path/to/key.pem")
// 	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
// 	if err != nil {
// 		log.Fatalf("failed to load client certificate: %v", err)
// 	}

// 	response, err := CallAPI(url, method, headers, params, cert)
// 	if err != nil {
// 		log.Fatalf("API call failed: %v", err)
// 	}

// 	fmt.Printf("API response: %+v\n", response)
// }
/////////////////////////////////////////////////////////////
/*
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// ExponentialBackoffWithJitter calculates the backoff delay for each retry attempt, with jitter.
func ExponentialBackoffWithJitter(attempt int) time.Duration {
	minWait := 1 * time.Second
	maxWait := 10 * time.Second

	// Exponential backoff with full jitter
	wait := time.Duration(math.Pow(2, float64(attempt))) * minWait
	jitter := time.Duration(rand.Int63n(int64(wait)))

	if wait+jitter > maxWait {
		return maxWait
	}

	return wait + jitter
}

// CallAPI makes an HTTP request with the specified settings and returns the response body.
func CallAPI(url string, method string, headers map[string]string, params map[string]interface{}, cert tls.Certificate) (map[string]interface{}, error) {
	// Configure TLS settings
	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: false, // Do not skip certificate verification
		Renegotiation:      tls.RenegotiateOnceAsClient,
		Certificates:       []tls.Certificate{cert}, // Client certificate for mutual TLS
	}

	// Configure HTTP transport settings
	tr := &http.Transport{
		TLSClientConfig:   tlsConfig,
		DisableKeepAlives: true, // Disable keep-alives to force new connections for each request
	}

	// Create a new Resty client with a custom transport
	client := resty.New().
		SetTransport(tr).
		SetTimeout(10 * time.Second).          // Set request timeout
		SetRetryCount(0).                      // Disable Resty's automatic retries
		SetRetryWaitTime(1 * time.Second).     // Set the minimum wait time between retries
		SetRetryMaxWaitTime(10 * time.Second). // Set the maximum wait time between retries
		AddRetryCondition(func(r *resty.Response, err error) bool {
			// Retry on network errors or 5xx HTTP status codes
			return r.StatusCode() >= 500 || err != nil
		})

	// Prepare the request
	request := client.R().
		SetHeaders(headers)

	// Set request parameters based on the HTTP method
	switch method {
	case "GET":
		stringParams := ConvertMapToStringMap(params)
		request.SetQueryParams(stringParams)
	case "POST", "PUT", "DELETE":
		request.SetBody(params)
	}

	// Define a variable to store the response body
	var responseBody map[string]interface{}
	var response *resty.Response
	var err error

	// Execute the request with manual retry logic
	for attempt := 0; attempt <= 5; attempt++ {
		fmt.Printf("Attempt %d...\n", attempt+1)
		response, err = request.SetResult(&responseBody).Execute(method, url)
		if err == nil && response.StatusCode() < 500 {
			return responseBody, nil
		}

		// Log the error and retry information
		fmt.Printf("Attempt %d failed: %v\n", attempt+1, err)
		if attempt < 5 {
			sleepDuration := ExponentialBackoffWithJitter(attempt)
			fmt.Printf("Retrying in %v...\n", sleepDuration)
			time.Sleep(sleepDuration)
		}
	}

	return nil, fmt.Errorf("failed after 5 attempts: %v", err)
}

// ConvertMapToStringMap is a helper function to convert map[string]interface{} to map[string]string
func ConvertMapToStringMap(params map[string]interface{}) map[string]string {
	stringParams := make(map[string]string)
	for key, value := range params {
		stringParams[key] = fmt.Sprintf("%v", value)
	}
	return stringParams
}

func main() {
	// Example usage of the CallAPI function
	url := "http://localhost:8081/v1/users/" //"https://example.com/api"
	method := "POST" // Example method
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]interface{}{
		"email":        "3246g11test288911@gmail.com",
		"password":     "fghjklhjgf",
		"name":         "sawerr",
		"check":        10,
		"created_time": time.Now().Format(time.RFC3339), // Using RFC3339 for standard time format
	}

	// Load client certificate
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	response, err := CallAPI(url, method, headers, params, cert)
	if err != nil {
		log.Fatalf("API call failed: %v", err)
	}

	fmt.Printf("API response: %+v\n", response)
}
*/
///////////////////////////////////////////////////////////////////////////

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

// ExponentialBackoffWithJitter calculates the backoff delay for each retry attempt, with jitter.
func ExponentialBackoffWithJitter(attempt int) time.Duration {
	minWait := 1 * time.Second
	maxWait := 10 * time.Second

	// Exponential backoff with full jitter
	wait := time.Duration(math.Pow(2, float64(attempt))) * minWait
	jitter := time.Duration(rand.Int63n(int64(wait)))
	log.Println("Jitter: ", jitter)

	if wait+jitter > maxWait {
		return maxWait
	}
	log.Println("Wait: ", wait)
	log.Println("wait + jitter: ", wait+jitter)
	return wait + jitter
}

// ResponseAPI defines the structure of the API response.
type ResponseAPI struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// Logger interface for logging with different severity levels.
type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

// logger implementation that uses the standard log package.
type logger struct {
	l *log.Logger
}

// createLogger initializes and returns a new logger instance.
func createLogger() *logger {
	return &logger{l: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)}
}

var _ Logger = (*logger)(nil)

// Errorf logs an error message.
func (l *logger) Errorf(format string, v ...interface{}) {
	l.output("ERROR RESTY "+format, v...)
}

// Warnf logs a warning message.
func (l *logger) Warnf(format string, v ...interface{}) {
	l.output("WARN RESTY "+format, v...)
}

// Debugf logs a debug message.
func (l *logger) Debugf(format string, v ...interface{}) {
	l.output("DEBUG RESTY "+format, v...)
}

// output formats and writes the log message.
func (l *logger) output(format string, v ...interface{}) {
	l.l.Printf(format, v...)
}

// CallAPI makes an HTTP request with the specified settings and returns the response body.
func CallAPI(url string, method string, headers map[string]string, params map[string]interface{}, cert tls.Certificate, logger Logger) (*ResponseAPI, error) {
	// Configure TLS settings
	tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: false, // Do not skip certificate verification
		Renegotiation:      tls.RenegotiateOnceAsClient,
		Certificates:       []tls.Certificate{cert}, // Client certificate for mutual TLS
	}

	// Configure HTTP transport settings
	tr := &http.Transport{
		TLSClientConfig:   tlsConfig,
		DisableKeepAlives: true, // Disable keep-alives to force new connections for each request
	}

	// Create a new Resty client with a custom transport
	client := resty.New().
		SetTransport(tr).
		SetTimeout(10 * time.Second).          // Set request timeout
		SetRetryCount(0).                      // Disable Resty's automatic retries
		SetRetryWaitTime(1 * time.Second).     // Set the minimum wait time between retries
		SetRetryMaxWaitTime(10 * time.Second). // Set the maximum wait time between retries
		AddRetryCondition(func(r *resty.Response, err error) bool {
			// Retry on network errors or 5xx HTTP status codes
			return r.StatusCode() >= 500 || err != nil
		})

	// Prepare the request
	request := client.R().
		SetHeaders(headers).
		SetDebug(true)

	// Set request parameters based on the HTTP method
	switch method {
	case "GET":
		stringParams := ConvertMapToStringMap(params)
		request.SetQueryParams(stringParams)
	case "POST", "PUT", "DELETE":
		request.SetBody(params)
	}

	// Define a variable to store the response body
	var responseBody ResponseAPI
	var response *resty.Response
	var err error

	// Execute the request with manual retry logic
	for attempt := 0; attempt <= 5; attempt++ {
		logger.Debugf("Attempt %d...\n", attempt+1)
		response, err = request.SetResult(&responseBody). //SetDebug(true).
									Execute(method, url)
		if err == nil && response.StatusCode() < 500 {
			logger.Debugf("Request succeeded on attempt %d\n", attempt+1)
			return &responseBody, nil
		}
		if response != nil {
			// Handle non-JSON responses (like HTML) here
			if response.StatusCode() == 404 {
				logger.Warnf("404 Not Found: %s\n", response.String())
			}
		}
		// Log the error and retry information
		logger.Errorf("Attempt %d failed: %v\n", attempt+1, err)
		if attempt < 5 {
			sleepDuration := ExponentialBackoffWithJitter(attempt)
			logger.Warnf("Retrying in %v...\n", sleepDuration)
			time.Sleep(sleepDuration)
		}
	}

	return nil, fmt.Errorf("failed after 5 attempts: %v", err)
}

// ConvertMapToStringMap is a helper function to convert map[string]interface{} to map[string]string
func ConvertMapToStringMap(params map[string]interface{}) map[string]string {
	stringParams := make(map[string]string)
	for key, value := range params {
		stringParams[key] = fmt.Sprintf("%v", value)
	}
	return stringParams
}

func main() {
	// Create a logger instance
	logger := createLogger()

	// Example usage of the CallAPI function
	//url := "https://example.com/api"
	url := "http://localhost:8081/v1/users/"
	method := "POST" // Example method
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]interface{}{
		"email":        "example@example.com",
		"password":     "examplepassword",
		"name":         "example",
		"check":        10,
		"created_time": time.Now().Format(time.RFC3339), // Using RFC3339 for standard time format
	}

	// Load client certificate
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		logger.Errorf("failed to load client certificate: %v", err)
	}

	response, err := CallAPI(url, method, headers, params, cert, logger)
	if err != nil {
		logger.Errorf("API call failed: %v", err)
	}

	fmt.Printf("API response: %+v\n", response)
}
