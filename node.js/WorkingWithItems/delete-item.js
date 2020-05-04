const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

// Notice the conditional expression where it will only delete if the postal code is the specified value
const params = {
  TableName: "RetailDatabase",
  Key: {
    pk: "jim.bob",
    sk: "metadata",
  },
  ConditionExpression: "shipaddr.pcode = :val",
  ExpressionAttributeValues: {
    ":val": "98260",
  },
};

documentClient.delete(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("DeleteItem succeeded:", JSON.stringify(data, null, 2));
  }
});
