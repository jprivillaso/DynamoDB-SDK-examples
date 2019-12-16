import boto3
import json
aas_client = boto3.client('application-autoscaling')
iam_client = boto3.client('iam')

table_name = "Movies2"
min_capacity = 1
max_capacity = 100
read_target = 50
write_target = 50
cooldown_duration_sec = 150

role_name = '{}TableScalingRole'.format(table_name)
policy_name = '{}TableScalingPolicy'.format(table_name)

assume_role_policy_document = {"Version": "2012-10-17", "Statement": [{"Effect": "Allow","Principal": {"Service": [ "ec2.amazonaws.com" ]},"Action": ["sts:AssumeRole"]}]}
policy_document = {"Version": "2012-10-17", "Statement":[{"Effect": "Allow", "Action": ["dynamodb:DescribeTable", "dynamodb:UpdateTable", "cloudwatch:PutMetricAlarm", "cloudwatch:DescribeAlarms", "cloudwatch:GetMetricStatistics", "cloudwatch:SetAlarmState", "cloudwatch:DeleteAlarms"], "Resource": "*"}]}
create_role_response = iam_client.create_role(
    Path='/',
    RoleName=role_name,
    AssumeRolePolicyDocument=json.dumps(assume_role_policy_document),
    Description='Table Scaling Role for {}'.format(table_name),
    MaxSessionDuration=3600
)

create_policy_response = iam_client.create_policy(
    PolicyName=policy_name,
    Path='/',
    PolicyDocument=json.dumps(policy_document)
)

role_arn = create_role_response['Role']['Arn']
policy_arn = create_policy_response['Policy']['Arn']

response = iam_client.attach_role_policy(
    RoleName=role_name,
    PolicyArn=policy_arn
)

response = aas_client.register_scalable_target(
    ServiceNamespace='dynamodb',
    ResourceId='table/{}'.format(table_name),
    ScalableDimension='dynamodb:table:ReadCapacityUnits',
    MinCapacity=min_capacity,
    MaxCapacity=max_capacity,
    RoleARN=role_arn
)

response = aas_client.register_scalable_target(
    ServiceNamespace='dynamodb',
    ResourceId='table/{}'.format(table_name),
    ScalableDimension='dynamodb:table:WriteCapacityUnits',
    MinCapacity=min_capacity,
    MaxCapacity=max_capacity,
    RoleARN=role_arn
)

response = aas_client.put_scaling_policy(
    PolicyName='{}ScalingPolicy'.format(table_name),
    ServiceNamespace='dynamodb',
    ResourceId='table/{}'.format(table_name),
    ScalableDimension='dynamodb:table:ReadCapacityUnits',
    PolicyType='TargetTrackingScaling',
    TargetTrackingScalingPolicyConfiguration={
        'TargetValue': read_target,
        'PredefinedMetricSpecification': {'PredefinedMetricType': 'DynamoDBReadCapacityUtilization'},
        'ScaleOutCooldown': cooldown_duration_sec,
        'ScaleInCooldown': cooldown_duration_sec,
        'DisableScaleIn': True
    }
)

response = aas_client.put_scaling_policy(
    PolicyName='{}ScalingPolicy'.format(table_name),
    ServiceNamespace='dynamodb',
    ResourceId='table/{}'.format(table_name),
    ScalableDimension='dynamodb:table:WriteCapacityUnits',
    PolicyType='TargetTrackingScaling',
    TargetTrackingScalingPolicyConfiguration={
        'TargetValue': write_target,
        'PredefinedMetricSpecification': {'PredefinedMetricType': 'DynamoDBWriteCapacityUtilization'},
        'ScaleOutCooldown': cooldown_duration_sec,
        'ScaleInCooldown': cooldown_duration_sec,
        'DisableScaleIn': True
    }
)
