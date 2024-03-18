# Проектирование и дизайн приложения

## Приложение построено по принципу гексагональной архитектуры:

![Гексагональная архитектура](images/hexagonal-architecture-svg.webp)

### Доменные модели
Доменная модель представляет собой модель предметной области, которая описывает основные понятия, понимание и логику бизнес-сущности.

В доменные модели были вынесены основные сущности приложения:
- Пользователь
- Сессия 
- Секрет
  - Данные типа K/V
  - Данные типа файл
  - Данные банковских карт

#### Доменные сервисы
Правила и процессы, которые являются основой функционирования системы. Они описывают ключевые понятия, атрибуты и операции, связанные с предметной областью, которую система предназначена поддерживать.

### Сервисы приложения
### Варианты использования
### Порты и адаптеры

## Клиент-серверное взаимодействие
![Взаимодействие Клиент-Сервер](images/client-server-structure-2.png)

## Структура базы данных Серверного приложения
![Структура базы данных](images/db-structure.drawio.png)