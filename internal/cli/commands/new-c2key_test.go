// Copyright 2020 Teserakt AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/teserakt-io/c2/pkg/pb"
	"github.com/teserakt-io/c2/internal/cli"
)

func newTestNewC2KeyCommand(clientFactory cli.APIClientFactory) cli.Command {
	cmd := NewNewC2KeyCommand(clientFactory)
	cmd.CobraCmd().SetOutput(ioutil.Discard)
	cmd.CobraCmd().DisableFlagParsing = true

	return cmd
}

func TestNewC2Key(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	c2Client := cli.NewMockC2Client(mockCtrl)

	c2ClientFactory := cli.NewMockAPIClientFactory(mockCtrl)
	c2ClientFactory.EXPECT().NewClient(gomock.Any()).AnyTimes().Return(c2Client, nil)

	t.Run("Execute forward expected request to the c2Client when passing a name", func(t *testing.T) {
		expectedRequest := &pb.NewC2KeyRequest{Force: true}

		c2Client.EXPECT().NewC2Key(gomock.Any(), expectedRequest).Return(&pb.NewC2KeyResponse{}, nil)
		c2Client.EXPECT().Close()

		cmd := newTestNewC2KeyCommand(c2ClientFactory)
		cmd.CobraCmd().Flags().Set("force", "1")
		err := cmd.CobraCmd().Execute()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Execute returns an error when --force is not present", func(t *testing.T) {
		cmd := newTestNewC2KeyCommand(c2ClientFactory)
		err := cmd.CobraCmd().Execute()
		if err == nil {
			t.Error("expected an error when no --force flag is present")
		}
	})
}
