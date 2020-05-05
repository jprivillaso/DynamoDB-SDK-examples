const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

const params = {
  TransactItems: [
    // Update the product status if the condition is met
    {
      Update: {
        TableName: "Products",
        Key: {
          ProductId: "B07JP6Z9PJ42",
        },
        UpdateExpression: "SET ProductStatus = :new_status",
        ConditionExpression: "ProductStatus = :expected_status",
        ExpressionAttributeValues: {
          ":new_status": "SOLD",
          ":expected_status": "IN_STOCK",
        },
        ReturnValuesOnConditionCheckFailure: "ALL_OLD",
      },
    },
    // Create the order if it doesn't already exist
    {
      Put: {
        TableName: "Orders",
        Item: {
          OrderId: "171-3549115-4111337",
          ProductId: "productKey",
          OrderStatus: "CONFIRMED",
          OrderTotal: "100",
        },
        ConditionExpression: "attribute_not_exists(OrderId)",
        ReturnValuesOnConditionCheckFailure: "ALL_OLD",
      },
    },
  ],
  ReturnConsumedCapacity: "TOTAL",
};

// Execute the actions defined as a single all-or-nothing operation
documentClient.transactWrite(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("transactWriteItem succeeded:", JSON.stringify(data, null, 2));
  }
});
