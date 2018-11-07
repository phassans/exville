package rocket

import "time"

const (
	apiPath = "api/v1"
	login   = "login"

	createUser = "users.create"
	deleteUser = "users.delete"
	infoUser   = "users.info"

	createChannel = "channels.create"
	deleteChannel = "channels.delete"
	infoChannel   = "channels.info"

	addUserToChannel      = "channels.invite"
	removeUserFromChannel = "channels.kick"
)

type (
	ErrorResponse struct {
		Status  string `json:"status"`
		Error   string `json:"error"`
		Message string `json:"message"`
	}

	AdminCredentials struct {
		AuthToken string `json:"authToken"`
		UserId    string `json:"userId"`
	}

	Err struct {
		Success   bool   `json:"success"`
		Error     string `json:"error"`
		ErrorType string `json:"errorType"`
	}
)

type (
	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserLoginResponse struct {
		Status string `json:"status"`
		Data   struct {
			UserID    string `json:"userId"`
			AuthToken string `json:"authToken"`
			Me        struct {
				ID     string `json:"_id"`
				Name   string `json:"name"`
				Emails []struct {
					Address  string `json:"address"`
					Verified bool   `json:"verified"`
				} `json:"emails"`
				Status           string   `json:"status"`
				StatusConnection string   `json:"statusConnection"`
				Username         string   `json:"username"`
				UtcOffset        int      `json:"utcOffset"`
				Active           bool     `json:"active"`
				Roles            []string `json:"roles"`
				Settings         struct {
					Preferences struct {
						EnableAutoAway              bool   `json:"enableAutoAway"`
						IdleTimeLimit               int    `json:"idleTimeLimit"`
						DesktopNotificationDuration int    `json:"desktopNotificationDuration"`
						AudioNotifications          string `json:"audioNotifications"`
						DesktopNotifications        string `json:"desktopNotifications"`
						MobileNotifications         string `json:"mobileNotifications"`
						UnreadAlert                 bool   `json:"unreadAlert"`
						UseEmojis                   bool   `json:"useEmojis"`
						ConvertASCIIEmoji           bool   `json:"convertAsciiEmoji"`
						AutoImageLoad               bool   `json:"autoImageLoad"`
						SaveMobileBandwidth         bool   `json:"saveMobileBandwidth"`
						CollapseMediaByDefault      bool   `json:"collapseMediaByDefault"`
						HideUsernames               bool   `json:"hideUsernames"`
						HideRoles                   bool   `json:"hideRoles"`
						HideFlexTab                 bool   `json:"hideFlexTab"`
						HideAvatars                 bool   `json:"hideAvatars"`
						SidebarGroupByType          bool   `json:"sidebarGroupByType"`
						SidebarViewMode             string `json:"sidebarViewMode"`
						SidebarHideAvatar           bool   `json:"sidebarHideAvatar"`
						SidebarShowUnread           bool   `json:"sidebarShowUnread"`
						SidebarShowFavorites        bool   `json:"sidebarShowFavorites"`
						SendOnEnter                 string `json:"sendOnEnter"`
						MessageViewMode             int    `json:"messageViewMode"`
						EmailNotificationMode       string `json:"emailNotificationMode"`
						RoomCounterSidebar          bool   `json:"roomCounterSidebar"`
						NewRoomNotification         string `json:"newRoomNotification"`
						NewMessageNotification      string `json:"newMessageNotification"`
						MuteFocusedConversations    bool   `json:"muteFocusedConversations"`
						NotificationsSoundVolume    int    `json:"notificationsSoundVolume"`
					} `json:"preferences"`
				} `json:"settings"`
			} `json:"me"`
		} `json:"data"`
	}

	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		User struct {
			ID        string    `json:"_id"`
			CreatedAt time.Time `json:"createdAt"`
			Services  struct {
				Password struct {
					Bcrypt string `json:"bcrypt"`
				} `json:"password"`
			} `json:"services"`
			Username string `json:"username"`
			Emails   []struct {
				Address  string `json:"address"`
				Verified bool   `json:"verified"`
			} `json:"emails"`
			Type         string    `json:"type"`
			Status       string    `json:"status"`
			Active       bool      `json:"active"`
			Roles        []string  `json:"roles"`
			UpdatedAt    time.Time `json:"_updatedAt"`
			Name         string    `json:"name"`
			CustomFields struct {
				Twitter string `json:"twitter"`
			} `json:"customFields"`
		} `json:"user"`
		Success bool `json:"success"`
	}

	DeleteUserRequest struct {
		UserId string `json:"userId"`
	}

	DeleteUserResponse struct {
		Success bool `json:"success"`
	}

	InfoUserRequest struct {
		Username string `json:"username"`
	}

	InfoUserResponse struct {
		User struct {
			ID        string `json:"_id"`
			Type      string `json:"type"`
			Status    string `json:"status"`
			Active    bool   `json:"active"`
			Name      string `json:"name"`
			UtcOffset int    `json:"utcOffset"`
			Username  string `json:"username"`
		} `json:"user"`
		Success bool `json:"success"`
	}

	ChannelCreateRequest struct {
		Name string `json:"name"`
	}

	ChannelCreateResponse struct {
		Channel struct {
			ID         string `json:"_id"`
			Name       string `json:"name"`
			Fname      string `json:"fname"`
			T          string `json:"t"`
			Msgs       int    `json:"msgs"`
			UsersCount int    `json:"usersCount"`
			U          struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
			} `json:"u"`
			CustomFields struct {
			} `json:"customFields"`
			Ts        time.Time   `json:"ts"`
			Ro        bool        `json:"ro"`
			SysMes    bool        `json:"sysMes"`
			Default   bool        `json:"default"`
			UpdatedAt time.Time   `json:"_updatedAt"`
			Lm        interface{} `json:"lm"`
		} `json:"channel"`
		Success bool `json:"success"`
	}

	ChannelErrorResponse struct {
		Success   bool   `json:"success"`
		Error     string `json:"error"`
		ErrorType string `json:"errorType"`
	}

	DeleteChannelRequest struct {
		RoomName string `json:"roomName"`
	}

	DeleteChannelResponse struct {
		Channel struct {
			ID        string   `json:"_id"`
			Name      string   `json:"name"`
			T         string   `json:"t"`
			Usernames []string `json:"usernames"`
			Msgs      int      `json:"msgs"`
			U         struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
			} `json:"u"`
			Ts time.Time `json:"ts"`
		} `json:"channel"`
		Success bool `json:"success"`
	}

	InfoChannelRequest struct {
		RoomName string `json:"roomName"`
	}

	InfoChannelResponse struct {
		Channel struct {
			ID        string    `json:"_id"`
			Ts        time.Time `json:"ts"`
			T         string    `json:"t"`
			Name      string    `json:"name"`
			Usernames []string  `json:"usernames"`
			Msgs      int       `json:"msgs"`
			Default   bool      `json:"default"`
			UpdatedAt time.Time `json:"_updatedAt"`
			Lm        time.Time `json:"lm"`
		} `json:"channel"`
		Success bool `json:"success"`
	}

	AddUserToChannelRequest struct {
		RoomId string `json:"roomId"`
		UserId string `json:"userId"`
	}

	AddUserToChannelResponse struct {
		Channel struct {
			ID        string    `json:"_id"`
			Ts        time.Time `json:"ts"`
			T         string    `json:"t"`
			Name      string    `json:"name"`
			Usernames []string  `json:"usernames"`
			Msgs      int       `json:"msgs"`
			UpdatedAt time.Time `json:"_updatedAt"`
			Lm        time.Time `json:"lm"`
		} `json:"channel"`
		Success bool `json:"success"`
	}

	RemoveUserFromChannelRequest struct {
		RoomId string `json:"roomId"`
		UserId string `json:"userId"`
	}

	RemoveUserFromChannelResponse struct {
		Channel struct {
			ID        string   `json:"_id"`
			Name      string   `json:"name"`
			T         string   `json:"t"`
			Usernames []string `json:"usernames"`
			Msgs      int      `json:"msgs"`
			U         struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
			} `json:"u"`
			Ts        time.Time `json:"ts"`
			Ro        bool      `json:"ro"`
			SysMes    bool      `json:"sysMes"`
			UpdatedAt time.Time `json:"_updatedAt"`
		} `json:"channel"`
		Success bool `json:"success"`
	}
)
