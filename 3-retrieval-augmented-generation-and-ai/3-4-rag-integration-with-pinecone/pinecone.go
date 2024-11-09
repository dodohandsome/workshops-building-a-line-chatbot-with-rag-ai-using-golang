package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pinecone-io/go-pinecone/pinecone"
	"google.golang.org/protobuf/types/known/structpb"
)

var pc *pinecone.Client
var idxConnection *pinecone.IndexConnection

func InitPinecone(ctx context.Context) {
	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create Client: %v", err)
	}
	indexName := os.Getenv("INDEX_NAME")
	idxModel, err := pc.DescribeIndex(ctx, indexName)
	if err != nil {
		log.Fatalf("Failed to describe index \"%v\": %v", indexName, err)
	}

	idxConnection, err = pc.Index(pinecone.NewIndexConnParams{Host: idxModel.Host})
	if err != nil {
		log.Fatalf("Failed to create IndexConnection for Host %v: %v", idxModel.Host, err)
	}
}

func QueryVectors(ctx context.Context, vector []float64) ([]QueryMatch, error) {
	queryVector := ConvertToFloat32(vector)
	queryResponse, err := idxConnection.QueryByVectorValues(ctx, &pinecone.QueryByVectorValuesRequest{
		TopK:            1,
		Vector:          queryVector,
		IncludeMetadata: true,
	})
	if err != nil {
		return nil, err
	}

	var results []QueryMatch
	fmt.Println("Query Results:")
	for _, match := range queryResponse.Matches {
		metadata, err := ConvertQueryMetadata(match.Vector.Metadata)
		if err != nil {
			log.Fatalf("Error converting metadata: %v", err)
		}
		if match.Score >= 0.9 {
			results = append(results, QueryMatch{
				ID:       match.Vector.Id,
				Score:    match.Score,
				Metadata: metadata,
			})

		}

	}

	return results, nil
}

func ConvertToFloat32(input []float64) []float32 {
	output := make([]float32, len(input))
	for i, v := range input {
		output[i] = float32(v)
	}
	return output
}

func ConvertMetadata(metadata map[string]interface{}) (*structpb.Struct, error) {
	return structpb.NewStruct(metadata)
}
func ConvertQueryMetadata(metadata *structpb.Struct) (Metadata, error) {
	meta := Metadata{}
	for key, value := range metadata.GetFields() {
		switch key {
		case "cut":
			meta.Cut = value.GetStringValue()
		case "text":
			meta.Text = value.GetStringValue()
		}
	}
	return meta, nil
}
