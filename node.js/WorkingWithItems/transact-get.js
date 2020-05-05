const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

// Define the partition keys for the two items we want to get.
const params = {
  TransactItems: [
    {
      Get: {
        TableName: "Products",
        Key: {
          ProductId: "B07JP6Z9PJ42",
        },
      },
    },
    {
      Get: {
        TableName: "Orders",
        Key: {
          OrderId: "171-3549115-4111337",
        },
      },
    },
  ],
  ReturnConsumedCapacity: "TOTAL",
};

documentClient.transactGet(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("transactGetItem succeeded:", JSON.stringify(data, null, 2));
  }
});
