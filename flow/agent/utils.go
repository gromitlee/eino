/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package agent

import (
	"errors"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

func ChatModelWithTools(model_ model.ChatModel, toolCallingModel model.ToolCallingChatModel,
	toolInfos []*schema.ToolInfo) (model.BaseChatModel, error) {

	if toolCallingModel != nil {
		if len(toolInfos) == 0 {
			return toolCallingModel, nil
		}
		return toolCallingModel.WithTools(toolInfos)
	}

	if model_ != nil {
		if len(toolInfos) == 0 {
			return model_, nil
		}
		err := model_.BindTools(toolInfos)
		if err != nil {
			return nil, err
		}

		return model_, nil
	}

	return nil, errors.New("no chat model provided")
}
