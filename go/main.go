package main

import (
	"bufio"
	"flag"
	"fmt"
	"main/groupsclient"
	"os"
	"strings"
)

func main() {
	baseUrl := flag.String("baseUrl", "", "base url of the groups.io server")
	emailPtr := flag.String("srcEmail", "", "groups.io email of the source user")
	passwordPtr := flag.String("srcPass", "", "groups.io password of the source user")
	cmdPtr := flag.String("cmd", "view", "Can be one of: getsubs, xferSubs, getUser")
	destEmailPtr := flag.String("destEmail", "", "email of user who will acquire your subscriptions and permissions on groups.io")

	flag.Parse()
	client := groupsclient.NewGroupsClient(*baseUrl)
	// Authenticate and get the token
	err := client.Authenticate(*emailPtr, *passwordPtr)
	if err != nil {
		fmt.Printf("main: client.Authenticate %+v", err)
		return
	}

	// Get user data associated with the user that we authorized the groups.io client.
	// For perms transfer this is the "source user", srcUser
	srcUser, err := client.GetAuthenticatedUser()
	if err != nil {
		fmt.Printf("main: Error getting user ID for %s: %v\n", *emailPtr, err)
		return
	}
	fmt.Printf("userId of loggedInUser: %v\n", srcUser.ID)
	fmt.Printf("FullName of loggedInUser: %v\n", srcUser.FullName)
	// Get the list of subgroups where the existing user has Owner permissions
	srcUsersSubs, subscriptionCount, err := client.GetMemberInfoList()

	if err != nil {
		fmt.Printf("main: Error getting user groups for %s: %v\n", srcUser.FullName, err)
		return
	}

	if subscriptionCount == 0 {
		fmt.Printf("main: %s is not subscribed to any groups!\n", *emailPtr)
		return
	} else {
		switch *cmdPtr {
		case "srcUserSubs":
			fullSummaryReport(emailPtr, subscriptionCount, srcUsersSubs)
		case "getUser":
			if len(*destEmailPtr) > 9 {
				targetUser, err := client.SearchMemberDetails(*destEmailPtr)
				if err != nil {
					fmt.Printf("main: switch case %s: SearchMemberDetails(%s) returned %v\n", *cmdPtr, *destEmailPtr, err)
					return
				}
				userReport(targetUser)
			} else {
				fmt.Printf("main: switch case %s: --destEmail not specified.\n", *cmdPtr)
			}

		case "xferSubs":
			targetUser, err := client.SearchMemberDetails(*destEmailPtr)
			if err != nil {
				fmt.Printf("Error running %s: %v\n", *cmdPtr, err)
				return
			}
			targetUserSubs, err := client.GrantOwnerPermsToGroupMember(*targetUser, srcUsersSubs)
			fmt.Printf("targetUserSubs %+v\n", targetUserSubs)
		default:
			fmt.Printf("main.go: unknown sub command %s\n", *cmdPtr)
		}
	}
}
func userReport(u interface{}) {
	fmt.Printf("User is %+v\n", u)
}

// fullSummaryReport reports on the logged-in user, showing their email, subs count and a list of their subscriptions
func fullSummaryReport(emailPtr *string, subscriptionCount int, loggedInUsersSubs []groupsclient.MemberInfo) {
	fmt.Printf("%s is subscribed to %d groups, they are...\n", *emailPtr, subscriptionCount)
	for _, subscription := range loggedInUsersSubs {
		fmt.Printf("%s, ", subscription.GroupName)
	}
}

// ContinuePrompt asks if user wants to continue
// if user enters y just return
// if user presses enter or n ir N then calls os.Exit(1)
func ContinuePrompt() {
	ok := YesNoPrompt("Do you want to continue? Defaults to No: [y/n] : ", false)
	if ok {
		fmt.Println("Continuing ...")
	} else {
		fmt.Println("Exiting ...")
		os.Exit(1)
	}
}

// YesNoPrompt asks yes/no questions using the label.
func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Printf("%s (%s) ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}
