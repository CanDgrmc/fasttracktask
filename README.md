# GoTask
 questions and answers are in config.yaml

## Duplicate answer 
http://localhost:3000/answer | POST
```json
{
	"UserId": 6,
	"QuizId": 1,
	"answers": [
			{
				"questionId": 1,
				"answerId" : 2
			},
			{
				"questionId": 1,
				"answerId" : 2
			},
			{
				"questionId": 3,
				"answerId" : 2
			},
			{
				"questionId": 4,
				"answerId" : 2
			}
		]
}
```

## Successful Answer 
http://localhost:3000/answer | POST
```json
{
	"UserId": 6,
	"QuizId": 1,
	"answers": [
			{
				"questionId": 1,
				"answerId" : 2
			},
			{
				"questionId": 2,
				"answerId" : 2
			},
			{
				"questionId": 3,
				"answerId" : 2
			},
			{
				"questionId": 4,
				"answerId" : 2
			}
		]
}
```

## GET 
http://localhost:3000/questions | GET
http://localhost:3000/answers/{questionId} | GET


