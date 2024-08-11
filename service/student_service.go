package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	t "github.com/marciomarinho/school-management-api-go/types"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *t.Student) error
	GetStudent(ctx context.Context, id string) (*t.Student, error)
	UpdateStudent(ctx context.Context, id string, student *t.Student) error
	DeleteStudent(ctx context.Context, id string) error
}

type StudentServiceImpl struct {
	db *dynamodb.Client
}

func NewStudentService(db *dynamodb.Client) StudentService {
	return &StudentServiceImpl{
		db: db,
	}
}

func (s *StudentServiceImpl) CreateStudent(ctx context.Context, student *t.Student) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Students"),
		Item: map[string]types.AttributeValue{
			"ID":    &types.AttributeValueMemberS{Value: student.ID},
			"Name":  &types.AttributeValueMemberS{Value: student.Name},
			"Class": &types.AttributeValueMemberS{Value: student.Class},
		},
	}

	_, err := s.db.PutItem(ctx, input)
	return err
}

func (s *StudentServiceImpl) GetStudent(ctx context.Context, id string) (*t.Student, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Students"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	result, err := s.db.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	student := &t.Student{
		ID:    result.Item["ID"].(*types.AttributeValueMemberS).Value,
		Name:  result.Item["Name"].(*types.AttributeValueMemberS).Value,
		Class: result.Item["Class"].(*types.AttributeValueMemberS).Value,
	}

	return student, nil
}

func (s *StudentServiceImpl) UpdateStudent(ctx context.Context, id string, student *t.Student) error {
	update := expression.Set(expression.Name("Name"), expression.Value(student.Name)).
		Set(expression.Name("Class"), expression.Value(student.Class))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Students"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = s.db.UpdateItem(ctx, input)
	return err
}

func (s *StudentServiceImpl) DeleteStudent(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Students"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	}

	_, err := s.db.DeleteItem(ctx, input)
	return err
}
