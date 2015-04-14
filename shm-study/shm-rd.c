#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include "chm_com.h"

int main(int argc, const char *argv[])
{
    int running = 1;
    void *shared_memory = (void *)0;
    struct shared_use_st *shared_stuff;
    int shmid;

    shmid = shmget((key_t)1234, sizeof(struct shared_use_st), 0666 | IPC_CREAT);
    if(shmid == -1) {
        fprintf(stderr, "shmget failed\n");
        exit(EXIT_FAILURE);
    }

    shared_memory = shmat(shmid, (void *)0,0);
    if(shared_memory == (void *)-1) {
        fprintf(stderr, "shmat failed\n");
        exit(EXIT_FAILURE);
    }

    printf("Memory attached at %X\n",(int)shared_memory);

    shared_stuff = (struct shared_use_st *)shared_memory;

    shared_stuff->written_by_you = 0;

    while(running) {
        if(shared_stuff->written_by_you) {
            printf("You wrote: %s",shared_stuff->some_text);
            sleep(1);
            shared_stuff->written_by_you=0;
            if(strncmp(shared_stuff->some_text,"end",3) == 0) {
                running = 0;
            }
        }
    }

    if(shmdt(shared_memory) == -1) {
        fprintf(stderr,"shmdt failerd!\n");
        exit(EXIT_FAILURE);
    }

    exit(EXIT_SUCCESS);
}
