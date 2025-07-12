# Append

Learn about Go's append function for slice manipulation:

## Topics Covered
- append() function basics
- Growing slices dynamically
- Appending multiple elements
- Slice capacity and growth
- Performance considerations

## Key Concepts
- **Dynamic Growth**: Slices can grow as needed
- **Capacity Management**: Understanding slice capacity
- **Copy Behavior**: When slices are copied
- **Variadic Parameters**: Multiple element appending

## JavaScript Comparison
```javascript
// JavaScript: Array methods
let arr = [1, 2, 3];
arr.push(4);           // Add single element
arr.push(5, 6, 7);     // Add multiple elements
arr = arr.concat([8, 9]); // Concatenate arrays

// Go: append function
slice := []int{1, 2, 3}
slice = append(slice, 4)         // Add single element
slice = append(slice, 5, 6, 7)   // Add multiple elements
slice = append(slice, []int{8, 9}...) // Concatenate slices
```

See `main.go` for detailed examples and patterns.
