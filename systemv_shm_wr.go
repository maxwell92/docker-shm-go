//systemv_shm_wr.go
package main

/*#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <stdio.h>

#define SHMSZ 27

int shm_wr() {
	char c;
	int shmid;
	key_t key;
	char *shm, *s;

	key = 5678;

	if((shmid = shmget(key, SHMSZ, IPC_CREAT | 0666)) < 0) {
		return -1;
	}

	if((shm = shmat(shmid, NULL, 0)) == (char *) -1) {
		return -2;
	}

	s = shm;
	for(c = 'a'; c <= 'z'; c++)
		*s++ = c;
	s = NULL;

	return 0;

}
*/
import "C"

import "fmt"

func main() {
	i := C.shm_wr()
	if i != 0 {
		fmt.Println("SystemV Share Memory Create and Write Error:", i)
		return
	}
	fmt.Println("SystemV Share Memory Create and Write OK")
}
