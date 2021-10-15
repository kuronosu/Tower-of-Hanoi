# Tower of Hanoi

Premisa 1 - Direccion del movimiento
si el total es par:
    disco par se mueve a la izquierda
    disco inpar se mueve a la derecha
en caso contrario
    disco par se mueve a la derecha
    disco inpar se mueve a la izquierda

Premisa 2:
un disco solo se debe mover cada 2**(n-1) pasos
por ejemplo en un problema de 3 discos seria asi:
    1 => 1 3 5 7
    2 => 2 6
    3 => 4
