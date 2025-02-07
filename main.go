package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Som-Kesharwani/shared-service/database"
	_ "github.com/Som-Kesharwani/shared-service/database"

	"github.com/Som-Kesharwani/shared-service/logger"
	_ "github.com/Som-Kesharwani/shared-service/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//Parse log level Info from command line
	/*logLevel := flag.Int("loglevel", 1, "an integer value (0-4)")
	flag.Parse()
	//calling the SetLogLevel with the comand-line argument
	logger.SetLogLevel(logger.Level(*logLevel), "Mylog.text")
	flag.Parse()*/
	//Calling the SetLogLevel with the command-line argument
	logger.Trace.Println("Main Started")
	loop()
	err := errors.New("text string")
	logger.Error.Println(err.Error())
	logger.Trace.Println("Main Completed")

	test := database.OpenCollection(database.Client, "test")
	test.InsertOne(context.Background(), bson.M{"name": "Som", "age": 25})
	test.Find(context.Background(), bson.M{"name": "som"})

}

func loop() {
	logger.Trace.Println("Loop startes")

	for i := 0; i < 10; i++ {
		logger.Info.Printf("Counter value is : %d", i)
	}
	logger.Warning.Printf("The counter variable is not being userd")
	logger.Trace.Println("Loop Completed!!")
	//arr1 := []int{1, 10, 100}
	//arr2 := []int{1000}
	//longestCommonPrefix(arr1, arr2)
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	hashMap := make(map[int]bool)

	for _, val := range arr1 {
		for val > 0 {
			if hashMap[val] {
				break
			}
			hashMap[val] = true
			val /= 10
		}
	}

	fmt.Println(hashMap)
	res := 0
	for _, val := range arr2 {
		for !hashMap[val] && val > 0 {
			val /= 10
		}
		tmp := len(strconv.Itoa(val))
		res = max(tmp, res)

	}

	return res

}
