 # Урок 5. Concurrency часть 2: основы типов данных из пакета sync
1. Напишите программу, которая запускает n потоков и дожидается завершения их всех
1. Реализуйте функцию для разблокировки мьютекса с помощью defer
1. Протестируйте производительность множества действительных чисел, безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение