package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/beevik/etree"                  // known RCE CVE
	mssqldb "github.com/denisenkom/go-mssqldb" // known Critical CVE
	jwt "github.com/dgrijalva/jwt-go"          // vulnerable package
	"github.com/gin-gonic/gin"                 // known vulnerabilities
	"github.com/gorilla/websocket"             // known High CVE
	"github.com/labstack/echo"                 // known Critical CVE
	"github.com/miekg/dns"                     // known High CVE
	"github.com/sirupsen/logrus"               // known vulnerabilities
	"github.com/spf13/viper"                   // known Critical CVE
	"github.com/ulikunitz/xz"                  // known Critical CVE
	"golang.org/x/crypto/ssh"                  // has had critical CVEs
	"gopkg.in/yaml.v2"                         // known High CVE - CVE-2022-28948
)

var hardcodedSecret = "supersecretkey123" // hardcoded secret
var hardcodedPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4f5wg5l2hKsTeNem/V41fGnJm6gOdrj8ym3rFkEjWT2btPW9
uoYuJPRhBgYh8DPPQ8lT5ggOkPIkqAoVAYzYhjdGzMDPPV1EYpd8F9LB7FaYRbLa
4Q8ELhqT8gGlwPX7WKHPTPePR0xH7lf45xqTEAFPgmGqbEe9ggOFSqhgZ8d3LPV2
1PzYBbE6nE7bZ5NzQY9ZHW8bYjdvC6LPOTxvUvfyYzTsGnJ7e/0xQG8vzTSKa6C9
dQgKj8fzDSLEcDiU5MvN4Vp4TZb5Yr2iuQ5KzV8/4E5x5u2ZJMKe7GyJZf+ZwSgk
Xs0DqELWPdKOdoGPZnGkPIAXrr9iKr9+aFfv5wIDAQABAoIBACy3t5PHsKl7y3pQ
UcG5j0OIH8J8JnM9LZb5XvQn9y7OV5VJb8dQS6yUDzNcGKLOQI5FpMxV8t6NtRjc
K2PXbKTgXRJ2HT6pxZKj5lD5hzPnPvl4iN1jW5hQ7zGCY5xX2oqo6F7mSaE6JcFl
zWCLQ2GXJr6LZm9VUAe9Q5YjJXF8L8Jv6dQQGjFzNmjG3sVpvNjN7Y9TY3P5QnM6
pFjTkXOJGt7QxYNGzYZ1Qs9I2cVEJzF6fKkN2fq2JyKrZiQKnGVTjXKbvJ9Z1hb2
rQq8GJZdNdw1pXJZlQq4V1Z3B5CfMxGKRtE2aKJC7vLFkZbQPRG6s8z3cBV+TQaX
H7tRKCECgYEA+zqJ2tMKyJfgZ1Q3qQoZhY5bXDtJ1H2tS8qz8U1cHLf8iG7RdAo6
r+1ZsVNfQK5z3sK8V3tF7PZ8Zm8fKO7J6+3T8yLnzd3RZF7V8fL1G8wZ3D8YvP9B
vLlxL6WvEbYtO2dNYjW4mJl4fJfJlvJ3tV8lLjK5zY1Ec2F3r8QGS7sCgYEA5nUF
2UuBOXI4MzgW9C7gGqzQZ2cVBd1pFdJYzYVeZlYl1s8tQNqQ8uHsJ7HjzuOD8xFN
F9mFdFJhZkZGKmXz8wZK8+9jZ8xdKzJ5mOxJXqy6YRXq9yF6jVd8YjOGdZlH9qV5
7x8gF6cK3lZ3TJq9LYy2QgFHnJKt7Y8WvZtZ8HECgYEA1KQ4+5GBc3J3L6VcBtUj
UNvP7dK9v7zTkGJD5b8q6dE3mT1lKKzN3r8QDLf3mI7b7pLjGb8q7B5TmN8q3lZN
VH8q9Kv3J7xPQG7Z8fY3dQfZ8Hl6A5g9pW8Xp8qAO5a9qTqUjKv7JzKwJ5qw4zq5
D6nHh8pN7Q2q8XxF+ZtF7AsCgYEA2E5z3JzLdE7+8L5D8xMq8QD5XTlh8d4qdGJz
MKBLjh8T7bVP6GQbZ8qsGp8gd8nJ5zL6k8QU5qG6pYZK8J8pK9G8J8K6d8U5z8pN
8dVd8Y5xJ8zD8K9q8J5HgF8q5Y8zK8nU8mP8qJ8x6tJ8U8V5Q6fL8q5Qr8K6d8X5
4f1b6C8K4q8Q5aE4cCCgYAZ8oT7J9J8U5C8b8K4L8q5Y8zQ8pQJz8p6t5b8K4d8X5
4pK8Q5N8z8q5D8nU8mP8qJ8x6tJ8U8V5Q6fL8q5Qr8K6d8X5z8p6t5b8K4d8X54f1b
-----END RSA PRIVATE KEY-----` // hardcoded private key

var PORT = "5655"

func main() {
	// Explicit usage to ensure SBOM detection
	doc := etree.NewDocument()
	doc.SetRoot(doc.CreateElement("root"))
	_ = doc.WriteSettings

	xzReader := xz.Reader{}
	_ = xzReader

	dnsMsg := dns.Msg{}
	dnsMsg.SetQuestion("example.com.", dns.TypeA)

	_ = ssh.InsecureIgnoreHostKey()

	_ = mssqldb.Driver{}

	// Use yaml package
	yamlData := make(map[string]interface{})
	yaml.Unmarshal([]byte("key: value"), &yamlData)

	// Use logrus
	logger := logrus.New()
	logger.Info("test")

	// Use viper
	viper.SetDefault("test", "value")

	// Use websocket
	_ = websocket.Upgrader{}

	// Use gin
	gin.SetMode(gin.ReleaseMode)

	// Use echo
	e := echo.New()
	_ = e

	http.HandleFunc("/sbom", func(w http.ResponseWriter, r *http.Request) {
		vulns := []map[string]string{
			{"id": "CVE-2023-12345", "severity": "High", "description": "Example high severity vulnerability in dependency X."},
			{"id": "CVE-2022-54321", "severity": "Critical", "description": "Example critical vulnerability in dependency Y."},
			{"id": "CVE-2021-11111", "severity": "High", "description": "Another high severity vulnerability in dependency Z."},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"vulnerabilities": vulns,
		})
	})

	// Insecure endpoint using vulnerable JWT library and hardcoded secret
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" { // insecure: should be POST
			username := r.URL.Query().Get("username")
			if username == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Missing username"))
				return
			}
			// Create JWT token insecurely
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": username,
			})
			tokenString, _ := token.SignedString([]byte(hardcodedSecret))
			w.Write([]byte(tokenString))
		}
	})

	// Insecure file upload endpoint
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			file, _, err := r.FormFile("file")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("File upload error"))
				return
			}
			defer file.Close()
			content, _ := ioutil.ReadAll(file)
			_ = ioutil.WriteFile("/tmp/uploaded_file", content, 0644) // insecure: no validation
			w.Write([]byte("File uploaded"))
		}
	})

	// Insecure SSH key handling - hardcoded private key (vulnerability)
	_, err := ssh.ParsePrivateKey([]byte(hardcodedPrivateKey))
	if err != nil {
		log.Println("Failed to parse private key")
	}

	log.Printf("Server starting on port %s", PORT)
	log.Printf("Environment: %v", os.Environ())

	// Add a simple health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Starting HTTP server on :%s", PORT)
	server_err := http.ListenAndServe(":"+PORT, nil)
	if server_err != nil {
		log.Fatalf("Server failed to start: %v", server_err)
	}
}
