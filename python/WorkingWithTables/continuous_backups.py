# snippet-sourcedescription:[ Enable or disable continuous backups on a table. ]
# snippet-service:[dynamodb]
# snippet-keyword:[Python]
# snippet-keyword:[Amazon DynamoDB]
# snippet-keyword:[Code Sample]
# snippet-keyword:[ backups, DynamoDB, table ]
# snippet-sourcetype:[full-example]
# snippet-sourcedate:[ ]
# snippet-sourceauthor:[AWS]
# snippet-start:[dynamodb.python.codeexample.enable_continuous_backups]

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

from __future__ import print_function  # Python 2/3 compatibility
import boto3, pprint

dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.Table("MyTable")

response = table.meta.client.update_continuous_backups(
    TableName='MyTable',
    PointInTimeRecoverySpecification={
        'PointInTimeRecoveryEnabled': True # options are True|False
    }
)

# Print the JSON response
pprint.pprint(response)
