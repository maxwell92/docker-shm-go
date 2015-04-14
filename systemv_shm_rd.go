//systemv_shm_rd.go

package main

/*
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdio.h>

#define SHMSZ 27

int shm_rd() {
	char c;
	int shmid;
	key_t key;
	char *shm,*s;

	key = 5678;

	if ((shmid = shmget(key,SHMSZ,0666)) < 0) {
		return -1;
	}

	if((shm = shmat(shmid, NULL, 0)) == (char *) -1) {
		return -2;
	}

	s = shm;

	int i=0;
	for(i=0;i<SHMSZ-1;i++)
		printf("%c ",*(s+i));
	printf("\n");
	s = NULL;

	return 0;
}
*/
import "C"

import "fmt"

func main() {
	i := C.shm_rd()
	if i != 0 {
		fmt.Println("SystemV Share Memory Create and Read Error:", i)
		return
	}
	fmt.Println("SystemV Share Memory Create and Read OK")
}
