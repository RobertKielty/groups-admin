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
	client := groupsclient.NewGroupsClient("https://lists.cncf.io")
	emailPtr := flag.String("email", "", "groups.io email")
	passwordPtr := flag.String("pass", "", "groups.io password")
	flag.Parse()

	// Authenticate and get the token
	err := client.Authenticate(*emailPtr, *passwordPtr)
	if err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	// Get user ID for the existing user email
	userDetails, err := client.GetLoggedInUserDetails()
	if err != nil {
		fmt.Printf("Error getting user ID for %s: %v\n", emailPtr, err)
		return
	}
	fmt.Printf("userId of loggedInUser: %v\n", userDetails.ID)
	fmt.Printf("FullName of loggedInUser: %v\n", userDetails.FullName)
	// Get the list of subgroups where the existing user has Owner permissions
	loggedInUsersSubs, subscriptionCount, err := client.GetMemberInfoList()

	if err != nil {
		fmt.Printf("Error getting user groups for %s: %v\n", userDetails, err)
		return
	}

	if subscriptionCount == 0 {
		fmt.Printf("%s is not subscribed to any groups!\n", *emailPtr)
		return
	} else {
		fmt.Printf("%s is subscribed to %d groups, they are...\n", *emailPtr, subscriptionCount)
		for _, subscription := range loggedInUsersSubs {
			fmt.Printf("%s, ", subscription.GroupName)
		}
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

//func getUserGroups(apiToken, userID string) ([]string, error) {
//	url := fmt.Sprintf("%s/users/%s/groups", apiBaseURL, userID)
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Authorization", "Bearer "+apiToken)
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//
//		}
//	}(resp.Body)
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
//	}
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var groups []Group
//	err = json.Unmarshal(body, &groups)
//	if err != nil {
//		return nil, err
//	}
//
//	var groupIDs []string
//	for _, group := range groups {
//		if group.Role == "owner" {
//			groupIDs = append(groupIDs, group.ID)
//		}
//	}
//
//	return groupIDs, nil
//}
//
//func grantOwnerPermissions(apiToken, groupID, userID string) error {
//	url := fmt.Sprintf("%s/groups/%s/members/%s/role", apiBaseURL, groupID, userID)
//	payload := map[string]string{"role": "owner"}
//	jsonPayload, err := json.Marshal(payload)
//	if err != nil {
//		return err
//	}
//
//	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
//	if err != nil {
//		return err
//	}
//
//	req.Header.Set("Authorization", "Bearer "+apiToken)
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//	defer func(Body io.ReadCloser) error {
//		err := Body.Close()
//		if err != nil {
//			return err
//		}
//		return err
//	}(resp.Body)
//
//	if resp.StatusCode != http.StatusOK {
//		return fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
//	}
//
//	return nil
//}
