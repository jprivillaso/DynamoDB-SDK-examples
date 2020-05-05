const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

// Define the partition keys and sort keys for the two items we want to get.
const params = {
  RequestItems: {
    RetailDatabase: {
      Keys: [
        {
          pk: "vikram.johnson@somewhere.com",
          sk: "metadata",
        },
        {
          pk: "jose.schneller@somewhere.com",
          sk: "metadata",
        },
      ],
      ConsistentRead: false, // For my user case, this data it is not changed often so why not get the reads at half price? Your use case might be different and need true.
    },
  },
  ReturnConsumedCapacity: "TOTAL",
};

documentClient.batchGet(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("BatchGetItem succeeded:", JSON.stringify(data, null, 2));
  }
});
