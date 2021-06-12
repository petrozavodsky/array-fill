# Coding Interview Challenge (Fill Array) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/e733574e65df4212be7a9e980abe366e)](https://www.codacy.com/gh/petrozavodsky/array-fill/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=petrozavodsky/array-fill&amp;utm_campaign=Badge_Grade)
## Постановка задачи

Необходимо реализовать логику, которая заполняет массив 5x5 случайными уникальными целыми числами в диапазоне 0-100.

Имейте в виду, что эти требования могут измениться в будущем, поэтому вам следует разработать код, который может легко адаптироваться к новому набору правил.


## Solution is implemented using Go (1.16)

### Build & Run

```
make run
```

### Test
```
make test
```


### Build docker image

```
docker build -t array-fill .
```

### Run in docker 

```
docker build -t array-fill . && docker run array-fill
```
