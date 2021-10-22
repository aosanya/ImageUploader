
### Deploy Step 1
Build the image using the command 
> docker build --tag amtest .  
<br />
amtest is an arbitrary name for the image

### Deploy Step 2
Run the image using the command
>docker run -d -p 8000:3000 amtest
<br />
8000 is a port of choice that is available on the host machine

### Test 
In the pages folder, load the index.html file. <br />
Note that it may be neccessary to update the port in the html file to reflect your choice of port in "Deploy Step 2". <br />

The default user id is "Sample Id 1". This can be changed using the "user id" input field.<br />

Choose a file and click upload.<br />

The uploaded user files can be listed by navigating to "localhost:8000/images/Sample Id 1"<br />