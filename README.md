# constrain-system
constrain-system by SICP

### How-to-Use
Use functional expression to represent 
```
9C = 5(F-32)
```
like this:
```go 
C, F := Variable("c"), Variable("f")
Equation(
    Multiplication(Params(Intermediate()), Params(Constant(9), C)).GetIntermediate(),
    Multiplication(Params(Intermediate()), Params(Constant(5),
        Addition(Params(Intermediate()), Params(F, Constant(-32))).GetIntermediate(),
    )).GetIntermediate()
)
```
