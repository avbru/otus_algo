Время работы: Append(int) // добавление в конец массива

| N      | 10          | 100         | 1000         | 1000           |
|--------|-------------|-------------|--------------|----------------|
| single | 804.6 ns/op | 5841 ns/op  | 395144 ns/op | 29667181 ns/op |
| vector | 188.0 ns/op | 809.7 ns/op | 41439 ns/op  | 3098584 ns/op  |
| factor | 536.2 ns/op | 1597 ns/op  | 3373 ns/op   | 35998 ns/op    |

Время работы: PUT(idx,n) // элемент добавлялся в начало массива

| N      | 10          | 100        | 1000         | 1000           |
|--------|-------------|------------|--------------|----------------|
| single | 680.3 ns/op | 6968 ns/op | 453134 ns/op | 34125830 ns/op |
| vector | 222.8 ns/op | 5642 ns/op | 81511 ns/op  | 7418037 ns/op  |
| factor | 611.9 ns/op | 3602 ns/op | 50544 ns/op  | 3887296 ns/op  |


Get, Delete методы не отличаются для разных типов динамических массивов