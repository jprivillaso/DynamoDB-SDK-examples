#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "Name of AWS profile"
    read awsProfile
elif [ $# -eq 1 ]; then
    awsProfile=$1
else
    echo Wrong number of arguments
    exit 1
fi

awsTempToken() {
    if [ ! -f ~/.aws/credentials ]; then
        echo "AWS credentials not configured!"
        return 1
    fi

    file="tmpCredentials.json"
    if [ -f "$file" ]; then
        rm ${file}
    fi

    AWS_TEMP_CREDS=$(aws sts assume-role --role-arn arn:aws:iam::544500146257:role/IdentityAccountAccessRole --role-session-name "RoleSession" --profile ${awsProfile})

    echo "{" >${file}
    echo "\"accessKeyId\"":\"$(echo $AWS_TEMP_CREDS | jq -r '.Credentials.AccessKeyId')\", >>${file}
    echo "\"secretAccessKey\"":\"$(echo $AWS_TEMP_CREDS | jq -r '.Credentials.SecretAccessKey')\", >>${file}
    echo "\"sessionToken\"":\"$(echo $AWS_TEMP_CREDS | jq -r '.Credentials.SessionToken')\", >>${file}
    echo "\"region\"":"\"eu-central-1\"" >>${file}
    echo "}" >>${file}

}

# retrieve temporary AWS credentials
awsTempToken
