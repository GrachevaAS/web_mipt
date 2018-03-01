#Лок Петерсона

Если в методе Lock поменять местами строчки want[t] = true и victim = t, т. е.
'''
Lock(t): 
  victim = t
  want[t] = true
  while want[1 - t] and victim == t:
    // wait and retry
'''
то может произойти такая ситуация:
