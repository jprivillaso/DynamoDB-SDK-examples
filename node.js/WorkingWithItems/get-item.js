const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

const params = {
  TableName: "RetailDatabase",
  Key: {
    pk: "jim.bob",
    sk: "metadata",
  },
};

documentClient.get(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("GetItem succeeded:", JSON.stringify(data, null, 2));
  }
});
