## What is covert channel
Covert channel is a type of attack that creates a way to transfer information between two processes that are not supposed to allowed to happen. A covert channel can be used to extract information from or inplant imformation to an organization. A good example can be two students cheat on an exam while not being caught by teacher.

## General steps to communicate
1. Sender emits bits of data by manipulating a shared resource.
2. Receiver can infer the bits by monitoring a shared resource

## Types of covert channel
- ### Network
Send or receive information covertly by exploiting unused or nonsensitive fields of network packets, protocol specification and properties, or timing of packets.
- ### Hardware
Sender and receiver share information among processes within a computer.
- ### Operating System
Operating systems maintain large amounts of data about the system and contain many data structures such as current time or file system that can potentially be manipulated.
- ### Hybrid
A combination of different classes of channels. For instance, an attack may send information by sending packets within
certain time intervals, and the receiver may receive it by observing the state of the operating system. A hybrid channel is harder to reason about and defend against than a traditional covert channel.


## Classes of communication ways
- ### Storage Channels
- ### Timing Channels
- 


