Чем отличаются RWMutex от Mutex?

Ответ:

```

Мьютекс - это примитив синхронизации горутин.
Предназначен для предоставления разделяемого ресурса
только одной горутине в моменте премени.

RWMutex, в отличие от обычного Mutex, позволяет параллельно
читать разделяемый ресурс не блокируя его.
Для этого имеются доп. методы RLock() и RUnlock().
Применим в ситуациях, когда имеется мало писателей,
но много параллельных читателей.

```