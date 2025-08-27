package topic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"product-service/pkg/constants"
	"product-service/pkg/consul"

	"github.com/hashicorp/consul/api"
)

type TopicService interface {
	GetTopicByID(ctx context.Context, id string) (*Topic, error)
}

type topicService struct {
	client *callAPI
}

type callAPI struct {
	client       consul.ServiceDiscovery
	clientServer *api.CatalogService
}

var (
	topicServiceStr = "gallery-service"
)

func NewTopicService(client *api.Client) TopicService {
	mainServiceAPI := NewServiceAPI(client, topicServiceStr)
	return &topicService{
		client: mainServiceAPI,
	}
}

func NewServiceAPI(client *api.Client, serviceName string) *callAPI {
	sd, err := consul.NewServiceDiscovery(client, serviceName)
	if err != nil {
		fmt.Printf("Error creating service discovery: %v\n", err)
		return nil
	}

	service, err := sd.DiscoverService()
	if err != nil {
		fmt.Printf("Error discovering service: %v\n", err)
		return nil
	}

	if os.Getenv("LOCAL_TEST") == "true" {
		fmt.Println("Running in LOCAL_TEST mode — overriding service address to localhost")
		service.ServiceAddress = "localhost"
	}

	return &callAPI{
		client:       sd,
		clientServer: service,
	}
}

func (s *topicService) GetTopicByID(ctx context.Context, id string) (*Topic, error) {

    token, ok := ctx.Value(constants.TokenKey).(string)
    if !ok {
        return nil, fmt.Errorf("token not found in context")
    }

    topic, err := s.client.getTopicByID(id, token)

    if err != nil {
        if sc, ok := topic["status_code"].(float64); ok && int(sc) == 404 {
            return nil, nil 
        }
        return nil, err
    }

    innerData, ok := topic["data"].(map[string]interface{})
    if !ok || innerData == nil {
        return nil, nil
    }

    idVal, _ := innerData["id"].(string)
    if idVal == "" {
        return nil, nil
    }

    nameVal, _ := innerData["topic_name"].(string)

    return &Topic{
        ID:   idVal,
        Name: nameVal,
    }, nil
}

func (c *callAPI) getTopicByID(id string, token string) (map[string]interface{}, error) {

	endpoint := fmt.Sprintf("/api/v1/gateway/gallery/topics/%s", id)

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + token,
	}

	res, err := c.client.CallAPI(c.clientServer, endpoint, http.MethodGet, nil, header)
	if err != nil {
		fmt.Printf("Error calling API: %v\n", err)
		return nil, err
	}

	var topicData interface{}

	err = json.Unmarshal([]byte(res), &topicData)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	myMap := topicData.(map[string]interface{})

	return myMap, nil
}
