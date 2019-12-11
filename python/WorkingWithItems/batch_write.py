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
from boto3.dynamodb.conditions import Key, Attr
from botocore.exceptions import ClientError

dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.Table("RetailDatabase")

with table.batch_writer() as batch:
    batch.put_item(
        Item={
            'pk': 'vikram.johnson@somewhere.com',
            'sk': 'metadata',
            'username': 'vikj',
            'first_name': 'Vikram',
            'last_name': 'Johnson',
            'name': 'Vikram Johnson',
            'age': 31,
            'address': {
                'road': '4328 Bakken Rd',
                'city': 'Greenbank',
                'pcode': 98253,
                'state': 'WA',
                'country': 'USA'
            }
        }
    )
    batch.put_item(
        Item={
            'pk': 'jose.schneller@somewhere.com',
            'sk': 'metadata',
            'username': 'joses',
            'first_name': 'Jose',
            'last_name': 'Schneller',
            'name': 'Jose Schneller',
            'age': 27,
            'address': {
                'road': '9531 Fish Rd',
                'city': 'Freeland',
                'pcode': 98249,
                'state': 'WA',
                'country': 'USA'
            }
        }
    )
    batch.put_item(
        Item={
            'pk': 'helga.ramirez@somewhere.com',
            'sk': 'metadata',
            'username': 'helgar',
            'first_name': 'Helga',
            'last_name': 'Ramirez',
            'name': 'Helga Ramirez',
            'age': 48,
            'address': {
                'road': '7243 Deer Lake Rd',
                'city': 'Clinton',
                'pcode': 98236,
                'state': 'WA',
                'country': 'USA'
            }
        }
    )