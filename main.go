package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"github.com/Shruty-Khullar/Lang_Translator/cli"
)
var wg sync.WaitGroup
var sourceLang string
var targetLang string
var sourceText string

func init() {
	
    //Flags: Command-line programs often accept flags or options from users to customize the command's execution. 
	//Command-line flags are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.
   //Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag word with a default value "foo" and a short description. This flag.String function returns a string pointer (not a string value); we’ll see how to use this pointer below.

	//Here we are declaring these 3 string var as flages with 1 parameter: pointer which will receive value returned by flag, the name of flag,its val,small description aboult flag
	flag.StringVar(&sourceLang,"s","en","Source language[en]")
	flag.StringVar(&targetLang,"t","fr","Target language[fr]")
	flag.StringVar(&sourceText,"st"," ","Text to Translate") 
}

func main() {
	//Once all flags are declared, call flag.Parse() to execute the command-line parsing.
    flag.Parse()
    
	//to see if no flags have been intialized we use no_of_flags falg.NFlags
	if flag.NFlag()==0 {
		fmt.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(1)
	}
    
	//Make a channel
	strchannel := make(chan string)
	wg.Add(1)  //add 1 process to waitgroup 

	//google lang api will interpret these 3 parameters with the mapped var
	reqbody := &cli.RequestBody {
		SourceLang : sourceLang,
		TargetLang : targetLang,
		SourceText : sourceText,
	}
    //In simple way it would have been like cli package has RequestBody strcut which is initializing the flags then it is using these flags to make a url request and recieve ans
	//But we will use channels and waitgroups 
	//Here Channel will the way by which gorountine will tell the main process that this is the result we got back
	go cli.RequestTranslate(reqbody,strchannel,&wg)   //now as we have aaded go infront of a funct call now its not a normal function call.It will be executed as separate process. So we need to add this in waitgroup so as to tell main fucntion to wait for 1 process
	//This function wont return anythig now anything. this will publish the result which it have fetched from the link and the (function call) go routine will fecth result frm channel

	processedstr:=strings.ReplaceAll(<-strchannel,"+"," ") //the result as fetched frm the link will have lots of + which we will replace using space
    fmt.Printf("%s\n",processedstr)
	close(strchannel)   //close channel

	wg.Wait()
}