# Лок Петерсона

Если в методе Lock поменять местами строчки want[t] = true и victim = t, т.е.
```
Lock(t): 
  victim = t
  want[t] = true
  while want[1 - t] and victim == t:
    // wait and retry
```
то может произойти такая ситуация:
- 0й поток пишет в victim 0
- 1й поток пишет в victim 1
- 1й поток пишет в want[1] **true**
- 1й поток выходит из while, т.к. в wictim[0] пока **false**
- 0й поток пишет в want[0] true
- 0й поток выходит из while, т.к. victim == 1  

&#8594 оба потока выходят из Lock() и попадают в критическую секцию, всё сломалось
