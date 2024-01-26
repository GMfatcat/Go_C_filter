## Go_C_filter

Goal:

Test for the speed to do convolution in pure go , goroutine, and Cgo.

Test with max pooling instead of convolution

Small JPG and Large JPG

```
Max pooling process in c completed in: 10050800
Max pooling process in go completed in: 3557900
Max pooling process in goroutine completed in: 517400
1000 times average for small image:
C avg: 2.042432ms
Go avg: 3.540069ms
Goroutine avg: 613.049µs
1000 times average for large image:
C avg: 51.926401ms
Go avg: 71.978682ms
Goroutine avg: 10.886355ms
Max pooling process completed.
```

What's Next?

1. find out how to use cgo in go mods(not use in main.go) -> 2024/1/26 done
2. Try to speed up go version by goroutine -> 2024/1/26 done

```
// make sure to get gcc compiler
go run pooling.go
```



