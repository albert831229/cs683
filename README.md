# cs683

## Write an internet worm

### 1. Scan
First generate a list of ip address range from 192.168.2.0 to 192.168.2.256. (Here I only generate up to 10 for the sake of testing)
Use socket package to scan through this list. For each ip address, check if it is running. If yes, try ssh to it.

### 2. SSH to it
Use paramiko package to connect to the running host with username and password in 'list' file.

### 3. Replicate
Copy the 'worm.py' and 'list' file to the host if it is not infected yet, then execute the worm program again. 

### 4. Safeguards
This script only scan the ip range from 192.168.2.0 to 192.168.2.10. And it will only attack(replicate)
the host if it has not been infected yet. A host is infected if it already has the file 'worm.py'. Finally, an infected host 
will keep looking for host that is not infected in this ip range. <br>

In my case, worm will first attack my raspberry pi, which is the only running machine within the given ip range, replicate 
itself, and then looking for other alive host. It will then ssh to my pi again. But this time it will not replicate itsef 
since my pi already contains 'worm.py'. Worm will simply skip this ip address. After trying all possible ip address in the given 
ip range, it will terminate itself, which prevents the widespread release of the worm.
