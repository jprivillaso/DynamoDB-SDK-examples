const AWS = require('aws-sdk');

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

const params = {
	TableName: 'RetailDatabase',
	FilterExpression: '#y = :y',
	ExpressionAttributeNames: {
		'#y': 'year',
	},
	ExpressionAttributeValues: {
		':y': 2020,
	},
}

documentClient.query(params, (err, data) => {
	if (err) {
		console.error(JSON.stringify(err, null, 2));
	} else {
		console.log("Query succeeded:", JSON.stringify(data, null, 2));
	}
})
