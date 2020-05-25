const AWS = require('aws-sdk');

AWS.config.update({ region: "us-west-2"});

const documentClient = new AWS.DynamoDB.DocumentClient();

const params = {
	TableName: 'Reply',
	KeyConditionExpression: '#id = :id and begins_with(#dt, :dt)',
	ExpressionAttributeNames: {
		'#id': 'Id',
		'#dt': 'ReplyDateTime'
	},
	ExpressionAttributeValues: {
		':id': "Amazon DynamoDB#DynamoDB Thread 1",
		':dt': "2015-09",
	},
}

documentClient.query(params, (err, data) => {
	if (err) {
		console.error(JSON.stringify(err, null, 2));
	} else {
		console.log("Query succeeded:", JSON.stringify(data, null, 2));
	}
})
