package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Message 表示一条聊天消息
type Message struct {
	Sender  string `json:"sender"`  // "user" 或 "ai"
	Content string `json:"content"` // 消息内容
}

type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 表示客户端发送的聊天请求
type ChatRequest struct {
	Message string `json:"message"`
}

// ChatResponse 表示返回给客户端的响应
type ChatResponse struct {
	Success bool   `json:"success"`
	Reply   string `json:"reply"`
}

type RequestBody struct {
	Model    string            `json:"model"`
	Messages []DeepSeekMessage `json:"messages"`
	Stream   bool              `json:"stream"`
}

type ResponseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

const (
	apiURL    = "https://api.deepseek.com/v1/chat/completions" // 请确认实际API地址
	apiKey    = "sk-02de10ad8d144e6f8bb46d47ac24f7f1"          // 替换为你的API密钥
	modelName = "deepseek-chat"                                // 确认实际模型名称

	apiURL2    = "https://ark.cn-beijing.volces.com/api/v3/chat/completions" // 请确认实际API地址
	apiKey2    = "1d323879-65a0-4c13-831c-97a3e21b76de"                      // 替换为你的API密钥
	modelName2 = "deepseek-v3-241226"                                        // 确认实际模型名称

)

// processAIMessage 模拟AI处理消息的逻辑
func ProcessAIMessage(message string) string {
	// 这里只是一个简单的示例，实际应用中应该调用AI API
	switch message {
	case "你好":
		return "你好！很高兴和你聊天。"
	case "你叫什么名字":
		return "我是一个AI助手，你可以叫我小A。"
	case "你会做什么":
		return "我可以回答你的问题，陪你聊天，提供各种信息。"
	default:
		return "我理解你说的是: " + message + "。这是一个有趣的话题！"
	}
}

func CallDeepseekAPI(payload RequestBody) (ResponseBody, error) {
	payload.Model = modelName

	// 打印请求日志
	//reqBody, _ := json.MarshalIndent(payload, "", "  ")
	//fmt.Printf("Request Body:\n%s\n", string(reqBody))

	jsonData, _ := json.Marshal(payload)

	// 创建HTTP请求
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	//fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("API返回错误: %s\n", body)
	}

	// 解析响应
	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("解析响应失败:", err)
	}

	if response.Error.Message != "" {
		fmt.Println("API错误:", response.Error.Message)
	}

	//if len(response.Choices) > 0 {
	//	assistantReply := response.Choices[0].Message.Content
	//	//fmt.Println("\nDeepSeek:", assistantReply)
	//	log.Println("\nDeepSeek:", assistantReply)
	//}
	return response, nil
}
