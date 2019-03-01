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
Storage channels involve data that is presented or written to a storage location to transmit information, ex. File system.
- ### Timing Channels
Timing channels send information in a way that involves manipulating the timing properties of a component of the system.

## Create a nearly undetectalbe convert channel with Tunnelshell
Tunnelshell is a light program that can only run on Linux. It operates in the similar client/server architecture like most Linux/Unix applications.  The beauty of Tunnelshell is that it supports variety of protocols including TCP, UDP, ICMP, and RawIP. Also, it is able to fragment packets to get past firewalls and intrusion detection systems.<br>

Steps to build an undetectable covert channel are as follows:
1. Fire up a Khali linux server and download Tunnelshell from https://packetstormsecurity.com/search/files/?q=Tunnelshell. 
2. We can upload Tunnelshell to the victim by TFTP it from my attacking systm.
2. Untar Tunnelshell and compile it by typing "make". To activate it on victim, type "./tunneld".
3. Connect to the tunnelshell daemon on the victim by typing "./tunnel -t frag victim/IP/address". When connect to Tunnelshell, it does not give us a command prompt, but rather a blank line. We can then type any Linux command and it returns output as if we were at a Linux prompt.
4. Try to detect Tunnelshell on the victim. Type "netstat" to show all current connections to the computer. Here, our tunnelshell will not display. However, if we list all running processes, our tunnelshell will appear on it unless we can embed it in a rootkit that will hide its process.

Reference: <br>
http://web.mit.edu/ha22286/www/papers/HST10.pdf <br>
https://null-byte.wonderhowto.com/how-to/hack-like-pro-create-nearly-undetectable-covert-channel-with-tunnelshell-0155704/
