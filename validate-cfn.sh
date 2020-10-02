#!/bin/bash

CLI_PROFILE=campsite

# Validate the CloudFormation template
aws cloudformation validate-template \
  --profile $CLI_PROFILE \
  --template-body file://main.yml