- Thread 

    - Execution Unit 
    - Thread is created by OS 
    - User Space program asks the kernel throughh a wrapper mostly like clib/gclib/musl etc .. the kernel creates a thread and gives the handle to the userspace application
    - These threads are called are system or os threads
    - Thread contains, some stack memory, context information, some software registers those simulate to hardware registers.

    8 core processor --> it can run only 8 threads at a time
    - Multiplexing M:N

    Process
        - Code 
        Shared
        - Heap
        - Global Variables
        - Open files/sockets
        
        - Threads
            - Thread-1
                - Thread ID
                - Stack Memory(2mb)
                    - function parameters
                    - local variables
                    - return addresses
                    - saved registers
                - scheduling state
                - Thread-level storage
                - CPU Registers
                - Program Counter
                - Stack Pointer

            - Thread-2
                - Stack Memory(2mb)
                    - function parameters
                    - local variables
                    - return addresses
                    - saved registers
                - scheduling state
                - Thread-level storage
                - CPU Registers
                - Program Counter
                - Stack Pointer



- G M P model
- G goroutines
- M Number of threads
- P process

- How goroutines are scheduled 
