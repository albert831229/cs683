# Create a polyglot file
Run with:<br>
```shell
./gen_poly.py --out magic.pdf --in "doc.pdf" --message monalisa_joconde_ascii.txt --zip winnie.jpeg --header sample.nes
```
- ```--out``` is the path/name of the output file, which will be a PDF file.
- ```--in``` is the path/name of the input file, which is your original PDF file.
- ```--message``` is the plaintext message that should appear when the resulting file will be opened in a hex editor, or directly ```cat``` in a terminal.
- ```--zip``` is the list of files you want to zip and append to the original PDF file.
- ```--header``` is the file that is to be included at the beginning of the file, before the PDF itself.
Running the command generates a new PDF file, ```magic.pdf``` which can be opend with any PDF reader and will show exactly the same as the original PDF.<br><br>
Then, we can unzip this file by ```unzip magic.pdf```. It will show the original PDF file and the files we just appended. Here, a picture of winnie will appear.<br><br>
Reference:<br>
https://github.com/perfaram/pdf-zip-nes-polyglot
