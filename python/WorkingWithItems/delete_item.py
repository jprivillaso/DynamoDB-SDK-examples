
from __future__ import print_function  # Python 2/3 compatibility
import boto3, json, decimal
from botocore.exceptions import ClientError

# Helper class to convert a DynamoDB item to JSON.
class DecimalEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, decimal.Decimal):
            if abs(o) % 1 > 0:
                return float(o)
            else:
                return int(o)
        return super(DecimalEncoder, self).default(o)


dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.Table('RetailDatabase')

# Delete an item, notice the conditional expression where it will only delete if the postal code is the specified value
try:
    response = table.delete_item(
        Key={
            'pk': "jim.bob",
            "sk": "metadata"
        },
        ConditionExpression="shipaddr.pcode = :val",
        ExpressionAttributeValues= {
            ":val": "98260"
        }
    )
except ClientError as e:
    if e.response['Error']['Code'] == "ConditionalCheckFailedException":
        print(e.response['Error']['Message'])
    else:
        raise
else:
    print("DeleteItem succeeded:")
    print(json.dumps(response, indent=4, cls=DecimalEncoder))