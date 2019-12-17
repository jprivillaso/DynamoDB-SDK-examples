
from __future__ import print_function # Python 2/3 compatibility
import boto3

dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.create_table(
    TableName="RetailDatabase",  # Substitute your table name for RetailDatabase
    BillingMode="PAY_PER_REQUEST",
    KeySchema=[
        {
            'AttributeName': 'pk',
            'KeyType': 'HASH'  #Partition key
        },
        {
            'AttributeName': 'sk',
            'KeyType': 'RANGE'  #Sort key
        }
    ],
    AttributeDefinitions=[
        {
            'AttributeName': 'pk',
            'AttributeType': 'S'
        },
        {
            'AttributeName': 'sk',
            'AttributeType': 'S'
        },
    ],
)

table.meta.client.get_waiter('table_exists').wait(TableName='MyTable')
print('Table has been created, please continue to insert data.')