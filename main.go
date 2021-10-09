package main

import (
	"context"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config in viper ", err)
		os.Exit(1)
	}

	client := godo.NewFromToken(viper.GetString("apiToken"))
	fmt.Printf("Here is client reference %#v\n", client)

	createRequest := &godo.DropletCreateRequest{
		Name:   viper.GetString("dropletName"),
		Region: viper.GetString("region"),
		Size:   viper.GetString("size"),
		Image: godo.DropletCreateImage{
			Slug: viper.GetString("image"),
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)
	if err != nil {
		fmt.Printf("Error creating Droplet %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Created new droplet ", newDroplet)
}
