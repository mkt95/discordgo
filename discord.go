// Discordgo - Discord bindings for Go
// Available at https://github.com/Bios-Marcel/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains high level helper functions and easy entry points for the
// entire discordgo package.  These functions are being developed and are very
// experimental at this point.  They will most likely change so please use the
// low level functions if that's a problem.

// Package discordgo provides Discord binding for Go
package discordgo

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// VERSION of DiscordGo, follows Semantic Versioning. (http://semver.org/)
const VERSION = "0.19.0"

// ErrMFA will be risen by New when the user has 2FA.
var ErrMFA = errors.New("account has 2FA enabled")

// NewWithToken creates a new Discord session and will use the given token
// for authorization.
func NewWithToken(userAgent, token string) (s *Session, err error) {
	// Create an empty Session interface.
	s = &Session{
		State:                  NewState(),
		Ratelimiter:            NewRatelimiter(),
		StateEnabled:           true,
		Compress:               true,
		ShouldReconnectOnError: true,
		ShardID:                0,
		ShardCount:             1,
		MaxRestRetries:         3,
		UserAgent:              userAgent,
		Token:                  token,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		sequence:               new(int64),
		LastHeartbeatAck:       time.Now().UTC(),
	}

	// The Session is now able to have RestAPI methods called on it.
	// It is recommended that you now call Open() so that events will trigger.

	return s, nil
}

// NewWithPassword creates a new Discord session and will sign in with the
// provided credentials.
//
// NOTE: While email/pass authentication is supported by DiscordGo it is
// HIGHLY DISCOURAGED by Discord. Please only use email/pass to obtain a token
// and then use that authentication token for all future connections.
// Also, doing any form of automation with a user (non Bot) account may result
// in that account being permanently banned from Discord.
func NewWithPassword(userAgent, username, password string) (s *Session, err error) {

	// Create an empty Session interface.
	s = &Session{
		State:                  NewState(),
		Ratelimiter:            NewRatelimiter(),
		StateEnabled:           true,
		Compress:               true,
		ShouldReconnectOnError: true,
		ShardID:                0,
		ShardCount:             1,
		MaxRestRetries:         3,
		UserAgent:              userAgent,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		sequence:               new(int64),
		LastHeartbeatAck:       time.Now().UTC(),
	}

	err = s.Login(username, password)
	if err != nil || s.Token == "" {
		if s.MFA {
			err = ErrMFA
		} else {
			err = fmt.Errorf("Unable to fetch discord authentication token. %v", err)
		}
		return
	}

	// The Session is now able to have RestAPI methods called on it.
	// It is recommended that you now call Open() so that events will trigger.

	return
}
