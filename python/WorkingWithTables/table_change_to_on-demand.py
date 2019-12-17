from __future__ import print_function # Python 2/3 compatibility
import boto3, pprint

dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.Table("RetailDatabase") # Substitute your table name for RetailDatabase

table.update(
    BillingMode="PAY_PER_REQUEST"
)