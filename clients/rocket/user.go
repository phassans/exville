package rocket

import (
	"encoding/json"
	"fmt"
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
)

func (c *client) Login(request UserLoginRequest) (UserLoginResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, login, AdminCredentials{})
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return UserLoginResponse{}, err
		}

		logger = logger.With().
			Int("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("login returned with error")
		return UserLoginResponse{}, fmt.Errorf("login returned with error: %s and code: %d", errResp.Message, errResp.Error)
	}

	// read response
	var resp UserLoginResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on UserLoginResponse")
		return UserLoginResponse{}, err
	}

	return resp, nil
}
