// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package raid

import (
	"encoding/json"
	"testing"

	"github.com/twitchdev/twitch-cli/internal/events"
	"github.com/twitchdev/twitch-cli/internal/models"
	"github.com/twitchdev/twitch-cli/test_setup"
)

var fromUser = "1234"
var toUser = "4567"

func TestEventSub(t *testing.T) {
	a := test_setup.SetupTestEnv(t)

	params := *&events.MockEventParameters{
		FromUserID:         fromUser,
		ToUserID:           toUser,
		Transport:          models.TransportWebhook,
		Trigger:            "raid",
		SubscriptionStatus: "enabled",
	}

	r, err := Event{}.GenerateEvent(params)
	a.Nil(err)

	var body models.SubEventSubResponse
	err = json.Unmarshal(r.JSON, &body)
	a.Nil(err)

	// write actual tests here (making sure you set appropriate values and the like) for eventsub
}

func TestFakeTransport(t *testing.T) {
	a := test_setup.SetupTestEnv(t)

	params := *&events.MockEventParameters{
		FromUserID:         fromUser,
		ToUserID:           toUser,
		Transport:          "fake_transport",
		Trigger:            "raid",
		SubscriptionStatus: "enabled",
	}

	r, err := Event{}.GenerateEvent(params)
	a.Nil(err)
	a.Empty(r)
}
func TestValidTrigger(t *testing.T) {
	a := test_setup.SetupTestEnv(t)

	r := Event{}.ValidTrigger("raid")
	a.Equal(true, r)

	r = Event{}.ValidTrigger("not_raid")
	a.Equal(false, r)
}

func TestValidTransport(t *testing.T) {
	a := test_setup.SetupTestEnv(t)

	r := Event{}.ValidTransport(models.TransportWebhook)
	a.Equal(true, r)

	r = Event{}.ValidTransport("noteventsub")
	a.Equal(false, r)
}
func TestGetTopic(t *testing.T) {
	a := test_setup.SetupTestEnv(t)

	r := Event{}.GetTopic(models.TransportWebhook, "trigger_keyword")
	a.NotNil(r)
}
