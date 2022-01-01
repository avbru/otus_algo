### GET, DELETE O(1), PUT ->O(1) на больших N, накладные расходы на изменения размера
make hw10

|                      |              |
|----------------------|--------------|
| GET_N=1000_-16       | 25180 ns/op  |
| PUT_N=1000_-16       | 115061 ns/op |
| DELETE_N=1000_-16    | 23805 ns/op  |
| GET_N=10000_-16      | 34096 ns/op  |
| PUT_N=10000_-16      | 127143 ns/op |
| DELETE_N=10000_-16   | 30450 ns/op  |
| GET_N=100000_-16     | 32516 ns/op  |
| PUT_N=100000_-16     | 70342 ns/op  |
| DELETE_N=100000_-16  | 33590 ns/op  |
| GET_N=1000000_-16    | 26892 ns/op  |
| PUT_N=1000000_-16    | 50413 ns/op  |
| DELETE_N=1000000_-16 | 27813 ns/op  |

### Домашнее задание

#### Хэш-таблицы, часть I

- Реализовать хеш-таблицу, использующую метод цепочек
- Или реализовать хеш-таблицу с открытой адресацией
- дополнительно: реализовать "ленивое" удаление
- реализовать квадратичный пробинг

### Критерии оценки:

3 балла максимум за основное задание, 2 за дополнительное.