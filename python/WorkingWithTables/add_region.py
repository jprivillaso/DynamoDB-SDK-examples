from __future__ import print_function # Python 2/3 compatibility
import boto3, pprint

#dynamodb = boto3.resource('dynamodb', region_name='us-west-2') # substitute your preferred region

client = boto3.client('dynamodb')

response = client.update_global_table(
    GlobalTableName='RetailDatabase',
    ReplicaUpdates=[
        {
            'Create': {'RegionName': 'us-west-2'}
        }
    ]
)

# Print the JSON response
pprint.pprint(response)
