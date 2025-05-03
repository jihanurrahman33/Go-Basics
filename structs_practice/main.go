package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs_practice/note"
	"example.com/structs_practice/todo"
)
func getNoteData()(string,string){
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title,content
}

type saver interface {
	Save() error 
	
} 

type outputtable interface{
	saver
	Display()
}
func outputData(data outputtable)error{
	data.Display()
	return saveData(data)
}

func saveData(data saver)error{
	data.Save()
	err := data.Save()
	if err != nil{
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Succesfully Saved note.")
	return nil
}

func main(){
	title,content:=getNoteData()
	todoText:=getUserInput("Todo text:")
	todo,err:=todo.New(todoText)
	if err !=nil{
		fmt.Println(err)
		return
	}
	
	userNote,err:=note.New(title,content)
	if err != nil{
		fmt.Println(err)
		return
	}

	err=outputData(todo)
	if err != nil{
		return
	}

	outputData(userNote)
	
	
}

func getUserInput(promt string)string{
	fmt.Printf("%v ",promt)
	reader:=bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	if err != nil{
		return ""
	}
	text =strings.TrimSuffix(text,"\n")
	text =strings.TrimSuffix(text,"\r")
	
	return text
}