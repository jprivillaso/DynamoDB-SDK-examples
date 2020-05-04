const AWS = require("aws-sdk");

AWS.config.update({ region: "us-west-2" });

const documentClient = new AWS.DynamoDB.DocumentClient();

// Define new user data to insert into the system
const params = {
  TableName: "RetailDatabase",
  Item: {
    pk: "jim.bob@somewhere.com",
    sk: "metadata",
    name: "Jim Bob",
    first_name: "Jim",
    last_name: "Bob",
    address: {
      road: "456 Nowhere Lane",
      city: "Langely",
      state: "WA",
      pcode: "98260",
      country: "USA",
    },
    username: "jbob",
  },
};

documentClient.put(params, (err, data) => {
  if (err) {
    console.error(JSON.stringify(err, null, 2));
  } else {
    console.log("PutItem succeeded:", JSON.stringify(data, null, 2));
  }
});
