# GO_backend_practice

**Ученик:** Глухов Сергей  
**Группа:** ЭФБО-07-26

## Практика № 1: Знакомство с Go
1. **Первая программа**
   - Файл: [`main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/main.go)
   - Описание: Базовая программа, печатающая приветствие.

2. **Карточка участника (переменные и типы)**
   - Файл: [`cmd/card/main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/cmd/card/main.go)
   - Описание: Демонстрация работы с константами (`const`), переменными разных типов (`string`, `int`, `bool`, `float64`), короткими объявлениями (`:=`) и форматированным выводом (`fmt.Printf`).

3. **Конвертер минут (самостоятельная работа)**
   - Файл: [`cmd/time-converter/main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/cmd/time-converter/main.go)
   - Описание: Программа для преобразования минут в часы и остаточные минуты.

## Практика №2: Условия, циклы, функции и коллекции
1. **Задача 1: Сводная информация о состоянии задачи**  
   - Файл: [`cmd/task-summary/main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/cmd/task-summary/main.go)
   - Описание: Программа генерирует список задач, добавляет новую задачу и отображает ее с нумерацией. Реализованы функции `statusLabel` и `countStatus`. Также реализован расчет общей статистики с помощью `map` и отображение выполненных задач.  

## Практика № 3: Структуры, методы и указатели
1. **Модель задачи и методы**
   - Файл: [`cmd/task-model/main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/cmd/task-model/main.go)
   - Описание: Структура `Task` с полями `ID`, `Title`, `Status`, `Priority`. Методы `Summary` (получатель по значению) и `Complete` (получатель-указатель `*Task`) для чтения и изменения состояния. Метод `SetPriority` с валидацией диапазона 1–3 и возвратом `bool`. Функция `findTaskByID` ищет задачу по ID и возвращает `&tasks[i]` либо `nil`, демонстрируя разницу между копией элемента в `range` и указателем на исходный элемент среза.

## Практика № 4: Интерфейсы, ошибки и консольное приложение
1. **Консольное приложение управления задачами**
   - Файл: [`cmd/task-cli/main.go`](https://github.com/Layokkoroche/GO_backend_practice/blob/main/go-practice-01/cmd/task-cli/main.go)
   - Описание: Интерфейс `TaskStore` (`Add`, `Get`, `Complete`, `Delete`, `List`) и реализация `MemoryStore` на срезе структур с CRUD-операциями. Валидация названия (`strings.TrimSpace`) и приоритета (1–3), удаление элемента среза по индексу, возврат копии среза через `append([]Task(nil), s.tasks...)`. Сентинельные ошибки `ErrNotFound` и `ErrInvalid`, обёртывание через `fmt.Errorf("...: %w", err)` и проверка через `errors.Is`. Интерактивное меню на едином `bufio.Scanner` с командами «добавить / список / завершить / удалить / выход», обработкой `strconv.Atoi`, `EOF` и неверного ввода. Завершение работы через `defer`.
