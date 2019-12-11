
from __future__ import print_function # Python 2/3 compatibility
import boto3, pprint



client = boto3.client('autoscaling-plans')

client.