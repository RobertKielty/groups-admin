package groupsclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TokenResponse structure to hold the JSON token response
type TokenResponse struct {
	Token string `json:"token"`
}

// GroupsClient struct to hold client configuration
type GroupsClient struct {
	BaseURL string
	Token   string
	Client  *http.Client
}
type User struct {
	ID                      int    `json:"id"`
	Object                  string `json:"object"`
	Created                 string `json:"created"`
	Updated                 string `json:"updated"`
	Email                   string `json:"email"`
	FullName                string `json:"full_name"`
	UserName                string `json:"user_name"`
	Timezone                string `json:"timezone"`
	Status                  string `json:"status"`
	ProfilePhotoURL         string `json:"profile_photo_url"`
	PostPref                string `json:"post_pref"`
	PerPagePref             string `json:"per_page_pref"`
	AllowFacebookLogin      bool   `json:"allow_facebook_login"`
	AllowGoogleLogin        bool   `json:"allow_google_login"`
	AllowSsoLogin           bool   `json:"allow_sso_login"`
	CsrfToken               string `json:"csrf_token"`
	TwoFactorEnabled        bool   `json:"two_factor_enabled"`
	RecoveryCodes           string `json:"recovery_codes"`
	DontMungeMessageID      bool   `json:"dont_munge_message_id"`
	AboutMe                 string `json:"about_me"`
	AboutFormat             string `json:"about_format"`
	Location                string `json:"location"`
	Website                 string `json:"website"`
	TimePref                string `json:"time_pref"`
	DatePref                string `json:"date_pref"`
	MondayStart             bool   `json:"monday_start"`
	ProfilePrivacy          string `json:"profile_privacy"`
	DefaultMessageView      string `json:"default_message_view"`
	TopicsSortDir           string `json:"topics_sort_dir"`
	TopicSortDir            string `json:"topic_sort_dir"`
	MessagesSortDir         string `json:"messages_sort_dir"`
	ExpandedMessagesSortDir string `json:"expanded_messages_sort_dir"`
	SearchSort              string `json:"search_sort"`
	SearchSortDir           string `json:"search_sort_dir"`
	PhotosOrderBy           string `json:"photos_order_by"`
	PhotosSortDir           string `json:"photos_sort_dir"`
	AlbumOrderBy            string `json:"album_order_by"`
	AlbumSortDir            string `json:"album_sort_dir"`
	DefaultCalendarView     string `json:"default_calendar_view"`
	DefaultHashtagView      string `json:"default_hashtag_view"`
	DefaultRsvpView         string `json:"default_rsvp_view"`
	HomePage                string `json:"home_page"`
}
type GroupData struct {
	ID                  int       `json:"id"`
	Object              string    `json:"object"`
	Created             string    `json:"created"`
	Updated             string    `json:"updated"`
	UserID              int       `json:"user_id"`
	GroupID             int       `json:"group_id"`
	GroupName           string    `json:"group_name"`
	Status              string    `json:"status"`
	PostStatus          string    `json:"post_status"`
	EmailDelivery       string    `json:"email_delivery"`
	MessageSelection    string    `json:"message_selection"`
	AutoFollowReplies   bool      `json:"auto_follow_replies"`
	MaxAttachmentSize   string    `json:"max_attachment_size"`
	ApprovedPosts       int       `json:"approved_posts"`
	ModStatus           string    `json:"mod_status"`
	PendingMsgNotify    string    `json:"pending_msg_notify"`
	PendingSubNotify    string    `json:"pending_sub_notify"`
	SubNotify           string    `json:"sub_notify"`
	StorageNotify       string    `json:"storage_notify"`
	SubGroupNotify      string    `json:"sub_group_notify"`
	MessageReportNotify string    `json:"message_report_notify"`
	AccountNotify       string    `json:"account_notify"`
	ModPermissions      string    `json:"mod_permissions"`
	OwnerMsgNotify      string    `json:"owner_msg_notify"`
	ChatNotify          string    `json:"chat_notify"`
	PhotoNotify         string    `json:"photo_notify"`
	FileNotify          string    `json:"file_notify"`
	WikiNotify          string    `json:"wiki_notify"`
	DatabaseNotify      string    `json:"database_notify"`
	Email               string    `json:"email"`
	UserStatus          string    `json:"user_status"`
	UserName            string    `json:"user_name"`
	Timezone            string    `json:"timezone"`
	FullName            string    `json:"full_name"`
	AboutMe             string    `json:"about_me"`
	Location            string    `json:"location"`
	Website             string    `json:"website"`
	ProfilePrivacy      string    `json:"profile_privacy"`
	DontMungeMessageID  bool      `json:"dont_munge_message_id"`
	UseSignature        bool      `json:"use_signature"`
	UseSignatureEmail   bool      `json:"use_signature_email"`
	Signature           string    `json:"signature"`
	Color               string    `json:"color"`
	CoverPhotoURL       string    `json:"cover_photo_url"`
	IconURL             string    `json:"icon_url"`
	NiceGroupName       string    `json:"nice_group_name"`
	SubsCount           int       `json:"subs_count"`
	MostRecentMessage   time.Time `json:"most_recent_message"`
	Perms               struct {
		Object                          string `json:"object"`
		ArchivesVisible                 bool   `json:"archives_visible"`
		PollsVisible                    bool   `json:"polls_visible"`
		MembersVisible                  bool   `json:"members_visible"`
		ChatVisible                     bool   `json:"chat_visible"`
		CalendarVisible                 bool   `json:"calendar_visible"`
		FilesVisible                    bool   `json:"files_visible"`
		DatabaseVisible                 bool   `json:"database_visible"`
		PhotosVisible                   bool   `json:"photos_visible"`
		WikiVisible                     bool   `json:"wiki_visible"`
		MemberDirectoryVisible          bool   `json:"member_directory_visible"`
		HashtagsVisible                 bool   `json:"hashtags_visible"`
		GuidelinesVisible               bool   `json:"guidelines_visible"`
		SubgroupsVisible                bool   `json:"subgroups_visible"`
		OpenDonationsVisible            bool   `json:"open_donations_visible"`
		SponsorVisible                  bool   `json:"sponsor_visible"`
		ManageSubgroups                 bool   `json:"manage_subgroups"`
		DeleteGroup                     bool   `json:"delete_group"`
		DownloadArchives                bool   `json:"download_archives"`
		DownloadEntireGroup             bool   `json:"download_entire_group"`
		DownloadMembers                 bool   `json:"download_members"`
		ViewActivity                    bool   `json:"view_activity"`
		CreateHashtags                  bool   `json:"create_hashtags"`
		ManageHashtags                  bool   `json:"manage_hashtags"`
		ManageIntegrations              bool   `json:"manage_integrations"`
		ManageGroupSettings             bool   `json:"manage_group_settings"`
		MakeModerator                   bool   `json:"make_moderator"`
		ManageMemberSubscriptionOptions bool   `json:"manage_member_subscription_options"`
		ManagePendingMembers            bool   `json:"manage_pending_members"`
		RemoveMembers                   bool   `json:"remove_members"`
		BanMembers                      bool   `json:"ban_members"`
		ManageGroupBilling              bool   `json:"manage_group_billing"`
		ManageGroupPayments             bool   `json:"manage_group_payments"`
		EditArchives                    bool   `json:"edit_archives"`
		ManagePendingMessages           bool   `json:"manage_pending_messages"`
		InviteMembers                   bool   `json:"invite_members"`
		CanPost                         bool   `json:"can_post"`
		CanVote                         bool   `json:"can_vote"`
		ManagePolls                     bool   `json:"manage_polls"`
		ManagePhotos                    bool   `json:"manage_photos"`
		ManageMembers                   bool   `json:"manage_members"`
		ManageCalendar                  bool   `json:"manage_calendar"`
		ManageChats                     bool   `json:"manage_chats"`
		ViewMemberDirectory             bool   `json:"view_member_directory"`
		ManageFiles                     bool   `json:"manage_files"`
		ManageWiki                      bool   `json:"manage_wiki"`
		ManageSubscription              bool   `json:"manage_subscription"`
		PublicPage                      bool   `json:"public_page"`
		SubPage                         bool   `json:"sub_page"`
		ModPage                         bool   `json:"mod_page"`
	} `json:"perms"`
	ExtraMemberData []struct {
		ColID          int       `json:"col_id"`
		ColType        string    `json:"col_type"`
		Text           string    `json:"text,omitempty"`
		Checked        bool      `json:"checked,omitempty"`
		Date           time.Time `json:"date,omitempty"`
		Time           time.Time `json:"time,omitempty"`
		StreetAddress1 string    `json:"street_address1,omitempty"`
		StreetAddress2 string    `json:"street_address2,omitempty"`
		City           string    `json:"city,omitempty"`
		State          string    `json:"state,omitempty"`
		Zip            string    `json:"zip,omitempty"`
		Country        string    `json:"country,omitempty"`
		Title          string    `json:"title,omitempty"`
		URL            string    `json:"url,omitempty"`
		Desc           string    `json:"desc,omitempty"`
		ImageName      string    `json:"image_name,omitempty"`
	} `json:"extra_member_data"`
}
type MemberInfoList struct {
	Object        string `json:"object"`
	TotalCount    int    `json:"total_count"`
	StartItem     int    `json:"start_item"`
	EndItem       int    `json:"end_item"`
	HasMore       bool   `json:"has_more"`
	NextPageToken int    `json:"next_page_token"`
	SortField     string `json:"sort_field"`
	SecondOrder   string `json:"second_order"`
	Query         string `json:"query"`
	SortDir       string `json:"sort_dir"`
	// groupData
	Data []GroupData `json:"data"`
}

// NewGroupsClient function to initialize the GroupsClient
func NewGroupsClient(baseURL string) *GroupsClient {
	return &GroupsClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

// Authenticate method to get and store the API token
func (c *GroupsClient) Authenticate(email, password string) error {
	formData := url.Values{
		"email":    {email},
		"password": {password},
		"token":    {"true"},
	}

	groupsApiLoginUrl := fmt.Sprintf("%s/api/v1/login", c.BaseURL)
	resp, err := c.Client.Post(groupsApiLoginUrl, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close() // ignoring error

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return err
	}

	c.Token = tokenResponse.Token
	return nil
}

// doRequest method to make authenticated HTTP requests
func (c *GroupsClient) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), body)

	if err != nil {
		return nil, err
	}

	// Add the token to the Authorization header using basic auth format
	req.SetBasicAuth(c.Token, "")
	return c.Client.Do(req)
}

// GetMemberInfoList method to get member info list with pagination
func (c *GroupsClient) GetMemberInfoList() ([]GroupData, int, error) {
	// First call should not include the page_token parameter
	objectLimit := 100
	subCount := 0
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/getsubs?limit=%d", objectLimit), nil)

	if err != nil {
		return nil, subCount, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, subCount, fmt.Errorf("GetMemberInfoList: first call to getsubs, received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, subCount, err
	}

	var memberInfoList MemberInfoList
	if err := json.Unmarshal(body, &memberInfoList); err != nil {
		return nil, 0, err
	}
	resp.Body.Close() // ignoring error
	subCount = memberInfoList.TotalCount
	allSubscriptions := make([]GroupData, 0, subCount)
	allSubscriptions = append(allSubscriptions, memberInfoList.Data[:]...)
	nextPageToken := memberInfoList.NextPageToken
	hasMore := memberInfoList.HasMore

	// ref https://groups.io/api#pagination
	for hasMore != false {
		endpoint := fmt.Sprintf("/api/v1/getsubs?limit=100&page_token=%d", nextPageToken)
		forLoopResponse, err := c.doRequest("GET", endpoint, nil)

		if err != nil {
			return nil, subCount, err
		}

		if forLoopResponse.StatusCode != http.StatusOK {
			return nil, subCount, fmt.Errorf("received non-200 response code: %d", forLoopResponse.StatusCode)
		}

		body, err := io.ReadAll(forLoopResponse.Body)
		if err != nil {
			return nil, subCount, err
		}

		var subscription MemberInfoList
		if err := json.Unmarshal(body, &subscription); err != nil {
			return nil, subCount, err
		}

		allSubscriptions = append(allSubscriptions, subscription.Data[:]...)
		nextPageToken = subscription.NextPageToken
		hasMore = subscription.HasMore
		forLoopResponse.Body.Close() // ignoring error
	}

	return allSubscriptions, subCount, nil
}

// GetLoggedInUserDetails method to get user details from the API
func (c *GroupsClient) GetLoggedInUserDetails() (*User, error) {
	resp, err := c.doRequest("GET", "/api/v1/getuser", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // ignoring error

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loggedInUserDetails User
	if err := json.Unmarshal(body, &loggedInUserDetails); err != nil {
		return nil, err
	}

	return &loggedInUserDetails, nil
}

// GrantOwnerPermsToUser assigns the OwnerRole to user for each group in groups
// returns number of groups that were updated or error
func (c *GroupsClient) GrantOwnerPermsToUser(user User, groups []GroupData) (int, error) {
	var groupsUpdated int = 0
	var err error = nil
	fmt.Printf("Making %v on %d groups\n", user, len(groups))

	return groupsUpdated, err
}
