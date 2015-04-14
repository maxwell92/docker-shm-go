//mmap_shm_wr.go
package main

/*
#include <stdio.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <fcntl.h>

#define SHMSZ 27

int shm_wr()
{
	char c;
	char *shm = NULL;
	char *s = NULL;
	int fd;
	if((fd = open("./shm.txt", O_RDWR | O_CREAT, S_IRUSR| S_IWUSR)) == -1)
	{
		return -1;
	}

	lseek(fd, 500, SEEK_CUR);
	write(fd, "\0", 1);
	lseek(fd, 0, SEEK_SET);

	shm = (char *)mmap(shm, SHMSZ, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
	if(!shm)
	{
		return -2;
	}

	close(fd);
	s = shm;
	for (c='a';c<='z';c++)
	{
		*(s+(int)(c-'a')) = c;
	}
	return 0;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	i := C.shm_wr()
	if i != 0 {
		fmt.Println("Mmap Share Memory Create and Write Error:", i)
		return
	}
	fmt.Println("Mmap Share Memory Create and Write OK")
}
