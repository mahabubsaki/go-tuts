# Defer

Learn about Go's defer statement for cleanup and resource management:

## Topics Covered
- Defer statement basics
- Execution order (LIFO)
- Resource cleanup
- Error handling with defer
- Panic and recover with defer

## Key Concepts
- **LIFO Order**: Last In, First Out execution
- **Resource Management**: Automatic cleanup
- **Error Handling**: Cleanup on error paths
- **Panic Recovery**: Graceful error handling

## JavaScript Comparison
```javascript
// JavaScript: Manual cleanup with try/finally
try {
    const file = openFile();
    // work with file
} finally {
    file.close();
}

// Go: Automatic cleanup with defer
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()
// work with file
```

See `main.go` for detailed examples and patterns.
