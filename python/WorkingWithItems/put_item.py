# snippet-sourcedescription:[ check your table limits ]
# snippet-service:[dynamodb]
# snippet-keyword:[Python]
# snippet-keyword:[Amazon DynamoDB]
# snippet-keyword:[Code Sample]
# snippet-keyword:[ ]
# snippet-sourcetype:[full-example]
# snippet-sourcedate:[ ]
# snippet-sourceauthor:[AWS]
# snippet-start:[dynamodb.python.codeexample.check_limits]

#
#  Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
#  This file is licensed under the Apache License, Version 2.0 (the "License").
#  You may not use this file except in compliance with the License. A copy of
#  the License is located at
#
#  http://aws.amazon.com/apache2.0/
#
#  This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
#  CONDITIONS OF ANY KIND, either express or implied. See the License for the
#  specific language governing permissions and limitations under the License.
#
from __future__ import print_function # Python 2/3 compatibility
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

#Insert new user data into the system
try:
    response = table.put_item(
       Item={
            'pk': "jim.bob",
            'sk': "metadata",
            'name': "Jim Bob",
            'shipaddr': {
                'addr': "456 Nowhere Lane",
                'city': "Langely",
                'state': "WA",
                'pcode': "98260"
            },
            'email': "jim.bob@somewhere.com"
        }
    )
except ClientError as e:
    print(e.response['Error']['Message'])
else:
    print("PutItem succeeded:")
    print(json.dumps(response, indent=4, cls=DecimalEncoder))