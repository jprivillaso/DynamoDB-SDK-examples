from __future__ import print_function # Python 2/3 compatibility
import boto3, pprint

#dynamodb = boto3.resource('dynamodb', region_name='us-west-2') # substitute your preferred region

client = boto3.client('dynamodb')

response = client.create_global_table(
    GlobalTableName='RetailDatabase',
    ReplicationGroup=[
        {
            'RegionName': 'us-west-2', "RegionName": "us-west-1",
        },
    ]
)

# Print the JSON response
pprint.pprint(response)
