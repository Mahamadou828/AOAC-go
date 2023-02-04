package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Mahamadou828/AOAC/business/sys/aws"
	sdkaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Test struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Addresses []struct {
		Street     string `json:"street"`
		PostalCode string `json:"postalCode"`
	}
}

type Database struct {
	sess *session.Session
	svc  *dynamodb.DynamoDB
	env  string
}

type FindByIndexInput[T any] struct {
	KeyName   string
	KeyValue  string
	Index     string
	TableName string
	Dest      *[]T
	Limit     int64
	StartEK   string
}

type FindOneByIndexInput[T any] struct {
	KeyName   string
	KeyValue  string
	Index     string
	TableName string
	Dest      *T
}

func Open(client *aws.Client, env string) *Database {
	return &Database{
		sess: client.Sess,
		svc:  dynamodb.New(client.Sess),
		env:  env,
	}
}

// Delete Deletes a single item in a table by primary key.
func Delete(ctx context.Context, client *Database, tableName, id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: sdkaws.String(id),
			},
		},
		TableName: formatTableName(client.env, tableName),
	}

	if _, err := client.svc.DeleteItemWithContext(ctx, input); err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

// Find fetch all items of a table and unmarshal them into the dest parameter. dest should be a non nil pointer to an array
// Find will return the last Evaluated Key for pagination if all element weren't return
func Find[S ~[]T, T any](ctx context.Context, client *Database, tableName, startKey string, limit int64, dest *S) (string, error) {
	var lastEvalKey string

	input := &dynamodb.ScanInput{
		TableName: formatTableName(client.env, tableName),
		Limit:     sdkaws.Int64(limit),
	}

	if len(startKey) > 0 {
		input.SetExclusiveStartKey(map[string]*dynamodb.AttributeValue{
			"id": {
				S: sdkaws.String(startKey),
			},
		})
	}

	result, err := client.svc.ScanWithContext(ctx, input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return lastEvalKey, fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return lastEvalKey, fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return lastEvalKey, fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return lastEvalKey, fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return lastEvalKey, fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return lastEvalKey, fmt.Errorf(err.Error())
		}
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &dest); err != nil {
		log.Fatalf("%v", err)
	}

	if _, ok := result.LastEvaluatedKey["id"]; ok {
		lastEvalKey = *result.LastEvaluatedKey["id"].S
	}

	return lastEvalKey, nil
}

// FindByID return an item by the given ID
func FindByID[T any](ctx context.Context, client *Database, id string, tableName string, dest *T) error {
	input := &dynamodb.QueryInput{
		TableName: formatTableName(client.env, tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			fmt.Sprintf(":%s", "id"): {
				S: sdkaws.String(id),
			},
		},
		KeyConditionExpression: sdkaws.String(fmt.Sprintf("%s = :%s", "id", "id")),
	}

	result, err := client.svc.QueryWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}

	if len(result.Items) <= 0 {
		return fmt.Errorf(dynamodb.ErrCodeResourceNotFoundException)
	}

	item := result.Items[0]

	if err := dynamodbattribute.UnmarshalMap(item, &dest); err != nil {
		return fmt.Errorf("can't unmarshal dynamodb attribute %v", err)
	}

	return nil
}

// FindOneByIndex return an item by the given index, the index should be created on the table
// because we return the first find element
func FindOneByIndex[T any](ctx context.Context, client *Database, inp FindOneByIndexInput[T]) error {
	input := &dynamodb.QueryInput{
		TableName: formatTableName(client.env, inp.TableName),
		IndexName: sdkaws.String(inp.Index),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			fmt.Sprintf(":%s", inp.KeyName): {
				S: sdkaws.String(inp.KeyValue),
			},
		},
		KeyConditionExpression: sdkaws.String(fmt.Sprintf("%s = :%s", inp.KeyName, inp.KeyName)),
	}

	result, err := client.svc.QueryWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}

	item := result.Items[0]

	if err := dynamodbattribute.UnmarshalMap(item, &inp.Dest); err != nil {
		return fmt.Errorf("can't unmarshal dynamodb attribute %v", err)
	}

	return nil
}

// FindByIndex return a list of item that match the given index, the index should be created on the table
func FindByIndex[T any](ctx context.Context, client *Database, inp FindByIndexInput[T]) (string, error) {
	var lastEK string

	input := &dynamodb.QueryInput{
		TableName: formatTableName(client.env, inp.TableName),
		IndexName: sdkaws.String(inp.Index),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			fmt.Sprintf(":%s", inp.KeyName): {
				S: sdkaws.String(inp.KeyValue),
			},
		},
		Limit:                  sdkaws.Int64(inp.Limit),
		KeyConditionExpression: sdkaws.String(fmt.Sprintf("%s = :%s", inp.KeyName, inp.KeyName)),
	}

	if len(inp.StartEK) > 0 {
		input.SetExclusiveStartKey(map[string]*dynamodb.AttributeValue{
			"id": {
				S: sdkaws.String(inp.StartEK),
			},
			inp.KeyName: {
				S: sdkaws.String(inp.KeyValue),
			},
		})
	}

	result, err := client.svc.QueryWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return lastEK, fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return lastEK, fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return lastEK, fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return lastEK, fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return lastEK, fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return lastEK, fmt.Errorf(err.Error())
		}
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &inp.Dest); err != nil {
		return lastEK, fmt.Errorf("can't unmarshal dynamodb attribute %v", err)
	}

	if _, ok := result.LastEvaluatedKey["id"]; ok {
		lastEK = *result.LastEvaluatedKey["id"].S
	}

	return lastEK, nil
}

// UpdateOrCreate will create a new item if the provided item key does not exist. Otherwise, it will update the item.
func UpdateOrCreate[T interface{}](ctx context.Context, client *Database, tableName string, data T) error {
	item, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return fmt.Errorf("can't marshal data: %v into dynamodb attribute", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: formatTableName(client.env, tableName),
	}

	_, err = client.svc.PutItemWithContext(ctx, input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeTransactionConflictException:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeTransactionConflictException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf("%v: %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

// BatchWrite allow to save multiple item in one request. The number of items should not be more than 25 and
// each item should not be more than 400kb. For more details see:
// https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb@v1.44.108#DynamoDB.BatchWriteItemWithContext
func BatchWrite[T any](ctx context.Context, client *Database, tableName string, items []T) error {
	var body []*dynamodb.WriteRequest

	for _, item := range items {
		marshalInput, _ := dynamodbattribute.MarshalMap(item)
		rq := &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: marshalInput,
			},
		}

		body = append(body, rq)
	}

	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*formatTableName(client.env, tableName): body,
		},
	}

	//@todo implement loop until all items has been saved.
	_, err := client.svc.BatchWriteItemWithContext(ctx, input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				return fmt.Errorf("can't save items: %v, %v", dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				return fmt.Errorf("can't save items: %v, %v", dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				return fmt.Errorf("can't save items: %v, %v", dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				return fmt.Errorf("can't save items: %v, %v", dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				return fmt.Errorf("can't save items: %v, %v", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				return fmt.Errorf("can't save items: %v", aerr.Error())
			}
		}
	}

	return nil
}

func formatTableName(env, tableName string) *string {
	return sdkaws.String(fmt.Sprintf("%s-%s", env, tableName))
}
