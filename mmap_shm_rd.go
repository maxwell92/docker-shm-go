//mmap_shm_rd.go
package main

/*
#include <stdio.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <fcntl.h>

#define SHMSZ 27

int shm_rd() {
	char c;
	char *shm = NULL;
	char *s = NULL;
	int fd;
	if ((fd = open("./shm.txt", O_RDONLY)) == -1) {
		return -1;
	}

	shm = (char *)mmap(shm, SHMSZ, PROT_READ, MAP_SHARED, fd, 0);

	if (!shm) {
		return -2;
	}

	close(fd);
	s = shm;
	int i=0;
	for(i=0;i<SHMSZ-1;i++) {
		printf("%c ",*(s+i));
	}
	printf("\n");

	return 0;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	i := C.shm_rd()
	if i != 0 {
		fmt.Println("Mmap Share Memory Read Error: ", i)
		return
	}
	fmt.Println("Mmap Share Memory Read OK")
}
