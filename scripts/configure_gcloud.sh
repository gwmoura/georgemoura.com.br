#!/bin/bash

# Get the SDK tar and untar it.
TARFILE=google-cloud-sdk.tar.gz
wget https://dl.google.com/dl/cloudsdk/release/$TARFILE
tar xzf $TARFILE
rm $TARFILE

# Install the SDK
./google-cloud-sdk/install.sh \
  --usage-reporting false \
  --path-update false \
  --command-completion false

gcloud components update

# Set config.
gcloud config set disable_prompts True
gcloud config set project $PROJECT_ID
gcloud config set app/promote_by_default false
gcloud auth activate-service-account --key-file "$APPLICATION_CREDENTIALS"

# Diagnostic information.
gcloud info
