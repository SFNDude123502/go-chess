# Go Chess

As it says in the README.md file, I am a high school student using a school-issued chromebook.      
Every month, the school finds the website that is not related to schoolwork with the most hours visited on by the computers and blocks it,       
Chess.com was one of those websites that we lost recently, now I do not care about the others in my school, but I did enjoy playing chess with my friend in 2nd period when we werent in the same classroom.      
         
     
The only logical next move is to make my own from scratch, for education purposes.      

      
      
I started by deciding what technology I would use for my server, I use Golang as my main language so I decided on the Gin Gonic http framework, I then got started on making a simple chess game in go, without the server.        
        
        
Step 1: Make a struct for a (blank)piece, then use composition to make structs which inherit the properties of piece && add specific ones       
Step 2: Make a [8][8]\*piece array, where in each sqare is a pointer to a piece, or a nil pointer if it is empty.      
Step 3: Make basic i/o functions, like printing+formatting the board and acceoting input, the input function works by asking for 2 strings, a location that holds the piece that you wish to move, and the location that you wish to move your piece to, and if it isnt valid, ask again      
Step 4: Validate input, by far the longest part. There are a few checks and steps to validate input:
* Starting Location has a piece
* Starting piece is not on enemies team
* Ending piece is not on your own team
* Move does not put king into check
* Move does not keep king in check
* Move is a valid location for the piece at starting move.        


Step 5: Make a master function where you input coords of piece and it will find out what type it is and then call the function specified for that type, ex: getAllOptions() will distribute pawns to getPawnOptions() and knights to getKnightOptions()
Step 6: make the get\*\*\*\*Options() functions, each class stores the movement info and the tests each possible movement in each direction, stopping at map borders and pieces in the way while ingoring all pieces with same color as attacking piece, then create an [][]int slice that keeps track of all valid coordinate pairs as it checks then returns it back to the input function, which checks to see if the [][]int slice contains the inputted coords      
Step 7: Make a function that finds every spot that the other team can attack then check to see if the king's coords is in that array  
Step 8: Implement check-checking, if at the begining and end of each player's turn it check to see if the king is in check, if it was before and is after, then you failed to move out of check and must redo your inputs, if it wasnt before and it was in check after then you moved into check and need to redo your inputs.
Step 9: Check every move your team can make, simulate all possible moves, if you are in check at start of round and the no matter what move you make you cant leave check at the end, its checkmate.        
Step 10: Make a function to reset the board so you can play again


Congrats! The intimidating part is done. Now its time for that part I thought was the easy part. 
Gin Gonic has a built-in function for serving html files, but the hard part was updating the page.  
Go has a built in templating laguage for html, but it only updates the page via http, meaning a page reload.  
I was planning on dealing with that later and I focused on the css and whatnot, here I decided that I would create a message box so players could chat while playing.

HTML-wise I made 1 grid, 2 forms, and 1 Unordered List. The 9x9 Grid included the chess board + line number/letters since the input is still typed out. Form 1 is the input for moves using the POST endpoint of /move, while Form 2 is the input to write new messages and uses the POST endpoint /message. The Unordered List is just there for displaying all the messages in chronilogical order with a scroll wheel to deal with overflow-y    

Now here is where I discover why tis part sucks. WebSockets.

Imagine playing chess with your friend who is in another building, and has to spam reload to see if you are afk, or just thinking for a while.  
We can't have that, so the better option to http is WebSockets, allowing a constant channel to send and recieve data from at any time, allowing for screen updates without reloading.  
I implimented the gorilla/websocket library and it worked perfectly on my computer at home.  
But I'm not running ths at home, so i loaded it up on github codespaces and "Websocket connection to xxxxxx.preview.github.dev Failed" pulled up in my console everytime I attempted to run it.  

Since then I still cant get it to work on my school laptop, I have tried other methods that don't cost money but none of them have worked, if you have an idea that I haven't tried yet  
* AWS, costs money
* Port Forwarding from home, accessing websited by direct IPs is blocked
* Github Pages, doesnt let you use your own http server, you can only input some html files
* Github Codespaces, websockets dont connect to the regular link
