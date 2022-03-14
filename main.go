package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

fun main(){
	discord, err := discordgo.New("Bot " + "authentication token")
	fmt.Println("Ok")
}