# Tower of Hanoi

![image](https://user-images.githubusercontent.com/28797741/137547086-5a4d159a-3c7a-4cfd-87d8-3abd3c957ffd.png)

## Premisa 1 - Direccion del movimiento

Si el total de discos es par:

    Disco par se mueve a la izquierda
    Disco inpar se mueve a la derecha
    
En caso contrario:

    Disco par se mueve a la derecha
    Disco inpar se mueve a la izquierda
    

## Premisa 2 - turno de mover

Un disco solo se debe mover cada 2<sup>(n-1)</sup> pasos, por ejemplo en un problema de 3 discos seria asi:
- 1 => 1 3 5 7
- 2 => 2 6
- 3 => 4
