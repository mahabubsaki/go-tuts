# Web Server

Learn how to build web servers in Go:

## Topics Covered
- HTTP server basics
- Request handling and routing
- Middleware patterns
- Template rendering
- Static file serving
- RESTful API design
- Authentication and authorization
- Error handling in web applications

## Key Concepts
- **HTTP Package**: Go's built-in HTTP server
- **Request/Response**: HTTP request and response handling
- **Routing**: URL pattern matching and handlers
- **Middleware**: Cross-cutting concerns
- **Templates**: HTML template rendering

## JavaScript Comparison
```javascript
// JavaScript (Node.js/Express)
const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.send('Hello World!');
});

app.listen(3000);

// Go: Built-in HTTP server
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
})
http.ListenAndServe(":8080", nil)
```

See `main.go` for detailed examples and patterns.
