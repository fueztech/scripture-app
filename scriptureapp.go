package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func ourFeelings() []string {
	var variousMoods = []string{"Happy", "Sad", "Mad", "Anxious"}
	fmt.Println(variousMoods)

	return variousMoods
}

func greetUsers() {

	fmt.Printf("Welcome to our Helpful Scriptures App!")
	fmt.Printf("These are the selections you an choose from: %v\n", ourFeelings())
	fmt.Printf("Using one word, how are you feeling today about life?:")

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var moodToday string
	anxiousScriptures := []string{
		"Psalms 118:6",
		"Psalms 139:23",
		"Isaiah 30:15",
		"Isaiah 41:10,13",
		"Isaiah 43:2",
		"Matthew 6:25-34",
		"Phillipians 4:6,7",
		"1John 4:4",
	}
	// These scripture refrences help you appreciate life.
	comfortScriptures := []string{
		"1Samuel 1:10-16",
		"1Kings 19:10-19",
		"Psalms 27:10",
		"Psalms 119:76",
		"Isaiah 57:15",
		"Mark 10:29-31",
		"Romans 15:4",
		"1Corinthians 1:3,4",
	}

	happyScriptures := []string{
		"Deuteronomy 33:29",
		"nehemiah 8:10",
		"Psalms 112:1",
		"Psalms 119:111",
		"Psalms 34:8, 84:12",
		"Psalms 41:1",
		"Psalms 84:5",
		"Proverbs 23:15-16",
		"Proverbs 28:14",
		"Matthew 5:1-11",
		"Matthew 16:17",
		"Romans 4:7,8",
		"James 1:12",
		"1Peter 4:14",
		"Revelations 16:15",
		"Revelation 20:6",
	}

	angerScriptures := []string{
		"Psalms 4:4",
		"Proverbs 19:11",
		"Ephesians 4:26",
		"James 1:19",
	}

	greetUsers()

	for {

		fmt.Scan(&moodToday)

		isValidInput := len(moodToday) >= 3
		isValidEmotion := strings.Contains(moodToday, "Happy") || strings.Contains(moodToday, "Uncertain") || strings.Contains(moodToday, "Mad") || strings.Contains(moodToday, "Sad")

		if !isValidInput || !isValidEmotion {
			fmt.Println("You made an incorrect emotion selection, please try again.")
			continue
		}

		if moodToday == "Mad" {
			fmt.Println("Your Weekly Scripture:", angerScriptures[rand.Intn(len(angerScriptures))])
		} else if moodToday == "Uncertain" {
			fmt.Println("Your Weekly Scripture:", anxiousScriptures[rand.Intn(len(angerScriptures))])
		} else if moodToday == "Happy" {
			fmt.Println("Your Weekly Scripture:", happyScriptures[rand.Intn(len(happyScriptures))])
		} else if moodToday == "Sad" {
			fmt.Println("Your Weekly Scripture:", comfortScriptures[rand.Intn(len(comfortScriptures))])
		}
		break
	}
}
