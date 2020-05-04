const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

// Define the name of a user account to update. Note that in this example, we have to alias "name" using ExpressionAttributeNames as name is a reserved word in DynamoDB.
const params = {
  TableName: "RetailDatabase",
  Key: {
    pk: "jim.bob",
    sk: "metadata",
  },
  ExpressionAttributeNames: {
    "#n": "name",
  },
  UpdateExpression: "set #n = :nm",
  ExpressionAttributeValues: {
    ":nm": "Big Jim Bob",
  },
  ReturnValues: "UPDATED_NEW",
};

documentClient.update(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("UpdateItem succeeded:", JSON.stringify(data, null, 2));
  }
});
