# matrigo
👩‍💻 Matrix library written in Go.

This code used to be in [github.com/fr3fou/gone](https://github.com/fr3fou/gone), but due to it being used in other places, I decided to extract it in its own repo.

 
## TODO

- [x] Randomize
- [x] Transpose
- [x] Scale
- [x] AddMatrix
- [x] Add
- [x] SubtractMatrix
- [x] Subtract
- [x] MultiplyMatrix
- [x] Multiply
- [x] Flatten
- [x] Unflatten
- [x] NewFromArray - makes a single row
- [x] Map
- [x] Fold
- [x] Methods to support chaining
- [x] Determinant
```go
n.Weights[i].
   Multiply(output).                         // weighted sum of the previous layer)
   Add(n.Layers[i+1].Bias).                  // bias
   Map(func(val float64, x, y int) float64 { // activation
       return n.Layers[i+1].Activator.F(val)
   })
```


## References

- <http://matrixmultiplication.xyz/>
- <https://www.khanacademy.org/math/precalculus/x9e81a4f98389efdf:matrices/x9e81a4f98389efdf:properties-of-matrix-addition-and-scalar-multiplication/a/properties-of-matrix-addition>
- <https://www.wikiwand.com/en/Matrix_(mathematics)>
