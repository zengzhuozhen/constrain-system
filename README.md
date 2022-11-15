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
Multiplication(
    Params(Constant(9), C), Params(Constant(5),
    Addition(Params(Intermediate()), Params(F, Constant(-32))).GetIntermediate()))
)
```
##### Inner Structure
![img.png](doc/img.png)