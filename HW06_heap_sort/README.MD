### Выводы
1. Heap Sort - стабильное время независимо от входных данных
2. Insertion Sort - хорошо подходит для почти отсортированных данных
3. Варианты Shell Sort с разным выбором шагов, работают в целом неплохо,
но нужно дополнительное тестирование на разных N/seed
4. Quick Sort - самый "предсказуемый": дает либо хорошее, либо лучшее время в зависимости от входных данных

N=1000_random__BubbleSort_-16         	    2570	    448130 ns/op
N=1000_random__Insertion__-16         	    4323	    270761 ns/op
N=1000_random__ShellSort__-16         	   29492	     40660 ns/op
N=1000_random__Sedgewick__-16         	   47070	     25634 ns/op
N=1000_random__ShellFrank_-16         	   29785	     40225 ns/op
N=1000_random__HeapSort___-16         	   31239	     38682 ns/op
N=1000_random__QuickSort__-16         	   28503	     41619 ns/op

N=1000_digits__BubbleSort_-16         	    2713	    426828 ns/op
N=1000_digits__Insertion__-16         	    4647	    254137 ns/op
N=1000_digits__ShellSort__-16         	   67645	     17933 ns/op
N=1000_digits__Sedgewick__-16         	   82430	     14398 ns/op
N=1000_digits__ShellFrank_-16         	   65990	     18124 ns/op
N=1000_digits__HeapSort___-16         	   40491	     29800 ns/op
N=1000_digits__QuickSort__-16         	   83054	     14260 ns/op

N=1000_sorted__BubbleSort_-16         	    3609	    327508 ns/op
N=1000_sorted__Insertion__-16         	  595522	      2028 ns/op
N=1000_sorted__ShellSort__-16         	  110794	     10794 ns/op
N=1000_sorted__Sedgewick__-16         	  177708	      6790 ns/op
N=1000_sorted__ShellFrank_-16         	  100383	     11935 ns/op
N=1000_sorted__HeapSort___-16         	   34378	     34733 ns/op
N=1000_sorted__QuickSort__-16         	   42378	     27879 ns/op

N=1000_reverse_BubbleSort_-16         	    2186	    525854 ns/op
N=1000_reverse_Insertion__-16         	    2197	    534287 ns/op
N=1000_reverse_ShellSort__-16         	   95655	     12448 ns/op
N=1000_reverse_Sedgewick__-16         	  120480	      9992 ns/op
N=1000_reverse_ShellFrank_-16         	   90979	     13149 ns/op
N=1000_reverse_HeapSort___-16         	   37302	     32246 ns/op
N=1000_reverse_QuickSort__-16         	   60241	     19944 ns/op

N=10000_random__BubbleSort_-16        	      26	  41966000 ns/op
N=10000_random__Insertion__-16        	      43	  26562245 ns/op
N=10000_random__ShellSort__-16        	    1626	    731373 ns/op
N=10000_random__Sedgewick__-16        	    2188	    543878 ns/op
N=10000_random__ShellFrank_-16        	    1669	    716100 ns/op
N=10000_random__HeapSort___-16        	    1653	    683981 ns/op
N=10000_random__QuickSort__-16        	    1585	    763600 ns/op

N=10000_digits__BubbleSort_-16        	      27	  38730579 ns/op
N=10000_digits__Insertion__-16        	      49	  23982619 ns/op
N=10000_digits__ShellSort__-16        	    3672	    333275 ns/op
N=10000_digits__Sedgewick__-16        	    4588	    259334 ns/op
N=10000_digits__ShellFrank_-16        	    3766	    319694 ns/op
N=10000_digits__HeapSort___-16        	    2319	    522675 ns/op
N=10000_digits__QuickSort__-16        	    5476	    214625 ns/op

N=10000_sorted__BubbleSort_-16        	      37	  31176998 ns/op
N=10000_sorted__Insertion__-16        	   46550	     25769 ns/op
N=10000_sorted__Sedgewick__-16        	   13611	     88248 ns/op
N=10000_sorted__ShellFrank_-16        	    7280	    162084 ns/op
N=10000_sorted__HeapSort___-16        	    2475	    480427 ns/op
N=10000_sorted__QuickSort__-16        	    3505	    343858 ns/op

N=10000_reverse_BubbleSort_-16        	      20	  53031228 ns/op
N=10000_reverse_Insertion__-16        	      21	  53183044 ns/op
N=10000_reverse_ShellSort__-16        	    7939	    157289 ns/op
N=10000_reverse_Sedgewick__-16        	   10000	    109732 ns/op
N=10000_reverse_ShellFrank_-16        	    7078	    171060 ns/op
N=10000_reverse_HeapSort___-16        	    2358	    505633 ns/op
N=10000_reverse_QuickSort__-16        	    4963	    242260 ns/op

N=100000_random__BubbleSort_-16       	       1	10362053724 ns/op
N=100000_random__Insertion__-16       	       1	2752320196 ns/op
N=100000_random__ShellSort__-16       	     100	  10062297 ns/op
N=100000_random__Sedgewick__-16       	     169	   6993492 ns/op
N=100000_random__ShellFrank_-16       	     118	  10102808 ns/op
N=100000_random__HeapSort___-16       	     130	   9208047 ns/op
N=100000_random__QuickSort__-16       	     121	  10001699 ns/op

N=100000_digits__BubbleSort_-16       	       1	10338883829 ns/op
N=100000_digits__Insertion__-16       	       1	2486346855 ns/op
N=100000_digits__ShellSort__-16       	     313	   3868742 ns/op
N=100000_digits__Sedgewick__-16       	     428	   2834246 ns/op
N=100000_digits__ShellFrank_-16       	     332	   3431415 ns/op
N=100000_digits__HeapSort___-16       	     204	   5928635 ns/op
N=100000_digits__QuickSort__-16       	     490	   2483338 ns/op

N=100000_sorted__BubbleSort_-16       	       1	3170379438 ns/op
N=100000_sorted__Insertion__-16       	    3584	    334034 ns/op
N=100000_sorted__ShellSort__-16       	     589	   2401256 ns/op
N=100000_sorted__Sedgewick__-16       	     637	   1721439 ns/op
N=100000_sorted__ShellFrank_-16       	     477	   2709579 ns/op
N=100000_sorted__HeapSort___-16       	     199	   5870616 ns/op
N=100000_sorted__QuickSort__-16       	     316	   3819202 ns/op

N=100000_reverse_BubbleSort_-16       	       1	5498371689 ns/op
N=100000_reverse_Insertion__-16       	       1	5488965767 ns/op
N=100000_reverse_ShellSort__-16       	     513	   2621469 ns/op
N=100000_reverse_Sedgewick__-16       	     562	   1988664 ns/op
N=100000_reverse_ShellFrank_-16       	     564	   2585345 ns/op
N=100000_reverse_HeapSort___-16       	     205	   5849447 ns/op
N=100000_reverse_QuickSort__-16       	     372	   3088666 ns/op

### Домашнее задание
#### Алгоритм HeapSort

#### Цель:
Реализовать рассмотренные алгоритмы сортировки и сравнить их эффективность.

2 байта Реализовать алгоритм BubbleSort, SelectionSort или InsertionSort (два на выбор).
3 байта Реализовать алгоритм ShellSort с тремя вариантами выбора шагов.
3 байта Реализовать алгоритм HeapSort.
2 байта Простестировать алгоритмы, заполнить таблицу по результатам измерений, написать свой вывод сравнения, какой алгоритм лучше.

Файл с тестами приложен к заданию (563 мегабайта).
На первой строчке указан размер массива, на второй строчке через пробел перечислены элементы массива.
В файл результата записать числа из отсортированный массив в одну строчку через пробел.

Тестировать алгоритмы следует на массивах таких размеров:
1, 10, 100, 1.000, 10.000, 100.000, 1.000.000, 10.000.000 (этот по желанию)

И с различным характером данных:
а) random - массив их случайных чисел
б) digits - массив из случайных цифр
в) sorted - на 99% отсортированный массив
г) revers - обратно-отсортированный массив

Напишите, какие пункты вы сделали, сколько байт набрали и сколько ушло времени на каждый пункт.
Приложите ссылку на ваш код и на заполненную таблицу.

Критерии оценки:
2 байта - BubbleSort, SelectionSort или InsertionSort (два на выбор)
3 байта - ShellSort с тремя вариантами выбора шагов.
3 байта - HeapSort.
2 байта - Тестирование и вывод
