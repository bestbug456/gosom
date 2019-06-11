# Gosom

Gosom is the memory efficent som implementation of [Kohonen Self Organising Map](https://ieeexplore.ieee.org/document/682362) (published in 1990) in pure Go. No external depdendences

## Gosom comparison with R library

For developing this library we analized the R library SOM. The 2 library split the wine dataset as follow

### R result

![r](/images/r.jpg)

### Gosom result

![r](/images/gosom.jpg)

As you can see excluding the initial randomness the 2 library split in the same way the data so you can assume we make the same error as R ;)

###### This library look like the library I always wanted. How can I use it?

Gosom is pretty straight forward library as you can see in the following example code we will initialise a new som and train it

``` Golang
    s, err := NewSom(5, 5, Rectangular, nil, 0, Bubble, false, 0, SumOfSquareDistance)
    if err != nil {
        panic("Error: %s\n", err.Error())
    }
    data := loadData()
    err = s.Train(data)
    if err != nil {
        panic("Error: %s\n", err.Error())
    }
    fmt.Printf("Codebooks: %+v\n", s.CodeBooks)
```

###### Ok this library seem esay to use but how about performance?

Sit down pick a cup of delicious coffee and watch the benchmark for the wine dataset:

``` Text
BenchmarkSomTrain-4             20      59128343 ns/op    142734051 B/op       35406 allocs/op
```

in less than 6 millisecond you have a complete train with the parameter you found in the example section, kinda cool isn't it?

###### This library seem very cool! Who create it?
This library was created by me (bestbug) and the best Data Scientist I've ever met [Franca Marinelli](https://www.linkedin.com/in/franca-marinelli-30b086126/ "Franca Marinelli")

###### If I found a bug or something wrong?
You can create a pull request to us (really appreciate) or if you don't have any idea about what is going wrong you can always open an issue here on GitHub!