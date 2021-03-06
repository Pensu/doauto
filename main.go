package main

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in viper ", err)
	}

	client := godo.NewFromToken(viper.GetString("apiToken"))
	fmt.Printf("Here is client reference %v\n", client)

	createRequest := &godo.DropletCreateRequest{
		Name:   viper.GetString("dropletName"),
		Region: viper.GetString("region"),
		Size:   viper.GetString("size"),
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("We got an error %s\n", err)
		// return err
	}

	fmt.Println("Created new droplet ", newDroplet)
}
